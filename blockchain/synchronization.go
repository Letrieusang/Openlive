package blockchain

import (
	"api-openlive/blockchain/create_event"
	"api-openlive/blockchain/create_event_sec"
	"api-openlive/blockchain/marketplace"
	"api-openlive/blockchain/opv_multiple_factory"
	"api-openlive/blockchain/transfer"
	common2 "api-openlive/common"
	"api-openlive/config"
	market_svc "api-openlive/module/market_place/business"
	marketplacemodel "api-openlive/module/market_place/model"
	user_svc "api-openlive/module/user/business"
	"api-openlive/monitor"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
	"time"
)

type Synchronization struct {
	cfg       config.Blockchain
	marketSvc *market_svc.MarketPlaceController
	userSvc   *user_svc.UserController
	logger    monitor.Logger
	c         *Client
	cache     *redis.Client
}

func NewSynchronization(cfg config.Blockchain, marketSvc *market_svc.MarketPlaceController, userSvc *user_svc.UserController, cache *redis.Client) *Synchronization {
	return &Synchronization{
		marketSvc: marketSvc,
		cfg:       cfg,
		userSvc:   userSvc,
		logger:    monitor.NewLogger("blc.out"),
		c:         NewClientHttp(),
		cache:     cache,
	}
}

func (s Synchronization) StartSync(isCronjob int) {
	if isCronjob == 1 {
		s.cronjobSync()
	} else {
		s.logger.Infof("--------------------------------------------------========================----------------------------------")
		s.logger.InforfWithCapture("Start Sync first time \nSC 1: %v\nSC 2: %v\n", s.cfg.GetOpvMarketplaceAddressFirst(), s.cfg.GetOpvMarketplaceAddressSecond())
		go s.autoRestartSync(s.cfg.GetOpvMarketplaceAddressFirst())
		go s.autoRestartSync(s.cfg.GetOpvMarketplaceAddressSecond())
		go s.autoRestartSyncTransfer()
		go s.autoRestartSyncEvent()
	}
}

func (s Synchronization) autoRestartSyncTransfer() {
	errC := make(chan error, 1)
	i := 0
	for {
		errC <- s.syncDataTransfer()
		err := <-errC
		i++
		if err != nil {
			s.logger.ErrofWithCapture("Sync transfer Time: %v\nError when sync %v", i, err.Error())
		} else {
			s.logger.Errorf("Sync transfer Time: %v Error is nil with time: %v", i)
		}
	}
}

func (s Synchronization) autoRestartSyncEvent() {
	errC := make(chan error, 1)
	i := 0
	for {
		errC <- s.syncDataEvent()
		err := <-errC
		i++
		if err != nil {
			s.logger.ErrofWithCapture("Sync event Time: %v\nError when sync %v", i, err.Error())
		} else {
			s.logger.Errorf("Sync event Time: %v Error is nil with time: %v", i)
		}
	}
}

func (s Synchronization) autoRestartSync(scAddress common.Address) {
	errC := make(chan error, 1)
	i := 0
	for {
		errC <- s.syncDataOfSC(scAddress)
		err := <-errC
		i++
		if err != nil {
			s.logger.ErrofWithCapture("Contract Address: %v\nTime: %v\nError when sync %v", scAddress.Hex(), i, err.Error())
		} else {
			s.logger.Errorf("Contract Address: %v, Time: %v Error is nil with time: %v", scAddress.Hex(), i)
		}
	}
}

func (s Synchronization) syncDataOfSC(scAddress common.Address) (errR error) {
	errC := make(chan error, 1)
	defer func() {
		if r := recover(); r != nil {
			s.logger.PanicfWithCapture("Panic when sync %v %v", scAddress.Hex(), r)
		}
	}()

	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c, err := ethclient.DialContext(timeout, s.cfg.GetWsUrl())
	if err != nil {
		s.logger.Errorf("Dial ws ... %v\n", err)
		return err
	}
	s.logger.Infof("Start Sync: ........... %v", scAddress.Hex())
	marketplaceFilter, err := marketplace.NewMarketplaceFilterer(scAddress, c)
	if err != nil {
		s.logger.Errorf("Create maketplace filter: %v\n", err)
		return err
	}

	timeout, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	createSellLogC := make(chan *marketplace.MarketplaceCreateMarketItem)
	createSellSub, err := marketplaceFilter.WatchCreateMarketItem(&bind.WatchOpts{Context: ctx}, createSellLogC, nil)

	if err != nil {
		s.logger.Errorf("Watch Create Sell Item: %v\n", err)
		return err
	}

	timeout, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	buyItemLogC := make(chan *marketplace.MarketplaceBuyMarketItem)
	buyItemSub, err := marketplaceFilter.WatchBuyMarketItem(&bind.WatchOpts{Context: ctx}, buyItemLogC, nil)
	if err != nil {
		s.logger.Errorf("Watch User Buy Item: %v\n", err)
		return err
	}

	timeout, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	cancelLogC := make(chan *marketplace.MarketplaceCancelSell)
	cancelSub, err := marketplaceFilter.WatchCancelSell(&bind.WatchOpts{Context: ctx}, cancelLogC, nil)
	if err != nil {
		s.logger.Errorf("Watch Cancel sell: %v\n", err)
		return err
	}

	go func() {
		for {
			select {
			case err := <-createSellSub.Err():
				s.logger.Errorf("Subscribe Create Sell Item")
				errC <- err
			case createSellLog := <-createSellLogC:
				go s.createSellItem(createSellLog, ctx)
			case err := <-buyItemSub.Err():
				s.logger.Errorf("Subscribe Buy Item")
				errC <- err
			case buyItemLog := <-buyItemLogC:
				go s.buyItem(buyItemLog, ctx)
			case err := <-cancelSub.Err():
				s.logger.Errorf("Subscribe Open Pack")
				errC <- err
			case cancelLog := <-cancelLogC:
				go s.cancelSellItem(cancelLog, ctx)
			}
		}
	}()
	errR = <-errC
	return errR
}

func (s Synchronization) syncDataTransfer() error {
	errC := make(chan error, 1)
	defer func() {
		if r := recover(); r != nil {
			s.logger.PanicfWithCapture("Panic when sync %v %v", s.cfg.GetOpvNftTransferAddress(), r)
		}
	}()

	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c, err := ethclient.DialContext(timeout, s.cfg.GetWsUrl())
	if err != nil {
		s.logger.Errorf("Dial ws ... %v\n", err)
		return err
	}
	transferFilter, err := transfer.NewTransferFilterer(s.cfg.GetOpvNftTransferAddress(), c)
	if err != nil {
		s.logger.Errorf("Create maketplace filter: %v\n", err)
		return err
	}

	timeout, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	nftTransferLog := make(chan *transfer.TransferTransferNFT)
	nftTransferSub, err := transferFilter.WatchTransferNFT(&bind.WatchOpts{Context: ctx}, nftTransferLog)
	if err != nil {
		s.logger.Errorf("Watch User Buy Item: %v\n", err)
		return err
	}

	go func() {
		for {
			select {
			case err := <-nftTransferSub.Err():
				s.logger.Errorf("Subscribe NFT transfer")
				errC <- err
			case nftTransfer := <-nftTransferLog:
				s.transferNft(nftTransfer, ctx)
			}
		}
	}()
	errR := <-errC
	return errR
}

func (s Synchronization) syncDataEvent() error {
	errC := make(chan error, 1)
	defer func() {
		if r := recover(); r != nil {
			s.logger.PanicfWithCapture("Panic when sync %v %v", s.cfg.GetOpvNftTransferAddress(), r)
		}
	}()

	ctx := context.Background()
	timeout, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c, err := ethclient.DialContext(timeout, s.cfg.GetWsUrl())
	if err != nil {
		s.logger.Errorf("Dial ws ... %v\n", err)
		return err
	}
	eventFilter, err := create_event.NewCreateEventFilterer(s.cfg.GetOpvNftCreateEventAddress(), c)
	if err != nil {
		s.logger.Errorf("Create create event filter: %v\n", err)
		return err
	}
	eventFilterSec, err := create_event_sec.NewCreateEventFilterer(s.cfg.GetOpvNftCreateEventSecAddress(), c)
	if err != nil {
		s.logger.Errorf("Create create event filter: %v\n", err)
		return err
	}

	timeout, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	nftCreateEventLog := make(chan *create_event.CreateEventJoinEvent)
	nftTransferSub, err := eventFilter.WatchJoinEvent(&bind.WatchOpts{Context: ctx}, nftCreateEventLog)

	nftCancelEventLog := make(chan *create_event.CreateEventOutEvent)
	nftCancelEventSub, err := eventFilter.WatchOutEvent(&bind.WatchOpts{Context: ctx}, nftCancelEventLog)

	nftCreateEventSecLog := make(chan *create_event_sec.CreateEventJoinEvent)
	nftTransferSubSec, err := eventFilterSec.WatchJoinEvent(&bind.WatchOpts{Context: ctx}, nftCreateEventSecLog)

	nftCancelEventSecLog := make(chan *create_event_sec.CreateEventOutEvent)
	nftCancelEventSubSec, err := eventFilterSec.WatchOutEvent(&bind.WatchOpts{Context: ctx}, nftCancelEventSecLog)
	if err != nil {
		s.logger.Errorf("Watch User Buy Item: %v\n", err)
		return err
	}

	go func() {
		for {
			select {
			case err := <-nftTransferSub.Err():
				s.logger.Errorf("Create Event NFT Error")
				errC <- err
			case nftCreateEvent := <-nftCreateEventLog:
				go s.createEvent(nftCreateEvent, ctx)
			case err := <-nftCancelEventSub.Err():
				s.logger.Errorf("Cancel Event NFT error")
				errC <- err
			case nftCancelEent := <-nftCancelEventLog:
				go s.cancelEvent(nftCancelEent, ctx)
			case err := <-nftTransferSubSec.Err():
				s.logger.Errorf("Create Event NFT Error")
				errC <- err
			case nftCreateEvent := <-nftCreateEventSecLog:
				go s.createEventSec(nftCreateEvent, ctx)
			case err := <-nftCancelEventSubSec.Err():
				s.logger.Errorf("Cancel Event NFT error")
				errC <- err
			case nftCancelEent := <-nftCancelEventSecLog:
				go s.cancelEventSec(nftCancelEent, ctx)
			}
		}
	}()
	errR := <-errC
	return errR
}

func (s Synchronization) cronjobSync() {
	isProcess, errGetCache := s.cache.Get(context.Background(), "cronjob_sync").Result()
	if errGetCache == nil && isProcess == "1" {
		fmt.Println("-----------------------------------------------------------")
		fmt.Println("Process still running")
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		return
	}
	if stt := s.cache.Set(context.Background(), "cronjob_sync", "1", 10*time.Minute); stt.Err() != nil {
		return
	}
	ctx := context.Background()
	blockNumber, err := s.marketSvc.Store.GetLatestBlock()
	if err != nil {
		log.Printf("Error when get block number %v\n", err.Error())
		return
	}
	blockNumber.BlockNumber++
	if blockNumber.BlockNumber < 17252049 {
		blockNumber.BlockNumber = 17252049
	}

	blockNumber.BlockNumber = 24459345
	startBlock := int64(blockNumber.BlockNumber)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Printf("Can not get new block header: %v\n", err)
		return
	}
	toBlock := header.Number.Int64()

	if startBlock > toBlock || toBlock <= 0 {
		fmt.Println("Error: From < To")
		return
	}
	var tmp int64
	for toBlock-startBlock > 0 {
		tmp = toBlock
		if toBlock-startBlock > 500 {
			tmp = startBlock + 500
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(startBlock),
			ToBlock:   big.NewInt(tmp),
			Addresses: []common.Address{
				s.cfg.GetOpvNftCreateEventAddress(),
				s.cfg.GetOpvNftCreateEventSecAddress(),
				s.cfg.GetOpvMarketplaceAddressFirst(),
				s.cfg.GetOpvMarketplaceAddressSecond(),
				s.cfg.GetOpvMarketplaceAddressThird(),
				s.cfg.GetOpvNftTransferAddress(),
				s.cfg.GetOpvFactoryAddress(),
				s.cfg.GetOpvCreateNftAddress(),
				s.cfg.GetOpvAuctionAddress(),
			},
		}
		ctx := context.Background()
		logs, err := s.c.client.FilterLogs(ctx, query)
		log.Printf("Filter from Block: %v To %v\n", startBlock, tmp)
		log.Println(s.cfg.GetOpvNftCreateEventSecAddress())
		if err != nil {
			log.Printf("Block: %v To %v, Filter logs error %v\n", startBlock, tmp, err.Error())
			return
		}

		s.atomicFilterLog(logs, ctx)
		startBlock = tmp
	}
	if stt := s.cache.Del(context.Background(), "cronjob_sync"); stt.Err() != nil {
		return
	}
}

func (s Synchronization) atomicFilterLog(logs []types.Log, ctx context.Context) {
	eventFilter, err := create_event.NewCreateEventFilterer(s.cfg.GetOpvNftCreateEventAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create create event filter: %v\n", err)
		panic(err)
	}
	eventFilterSec, err := create_event_sec.NewCreateEventFilterer(s.cfg.GetOpvNftCreateEventSecAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create create event filter: %v\n", err)
		panic(err)
	}
	transferFilter, err := transfer.NewTransferFilterer(s.cfg.GetOpvNftTransferAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create maketplace filter: %v\n", err)
		panic(err)
	}
	/*marketplaceFilter, err := marketplace.NewMarketplaceFilterer(s.cfg.GetOpvMarketplaceAddressFirst(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create maketplace filter: %v\n", err)
		panic(err)
	}*/
	marketplaceFilterSec, err := marketplace.NewMarketplaceFilterer(s.cfg.GetOpvMarketplaceAddressSecond(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create maketplace filter: %v\n", err)
		panic(err)
	}

	/*factoryFilter, err := factory.NewFactoryFilterer(s.cfg.GetOpvFactoryAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create factory filter: %v\n", err)
		panic(err)
	}*/

	createNftFilter, err := opv_multiple_factory.NewOpvMultipleFactory(s.cfg.GetOpvCreateNftAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create factory filter: %v\n", err)
		panic(err)
	}

	/*auctionFilter, err := auction.NewAuctionFilterer(s.cfg.GetOpvAuctionAddress(), s.c.client)
	if err != nil {
		s.logger.Errorf("Create factory filter: %v\n", err)
		panic(err)
	}*/

	//get list processed ws block
	wsBlocks, _ := s.marketSvc.Store.GetWsBlock()
	rebuildWsBlock := make(map[uint64]marketplacemodel.BlockChainBlock)
	if len(wsBlocks) > 0 {
		for _, tmpWsBlock := range wsBlocks {
			rebuildWsBlock[tmpWsBlock.BlockNumber] = tmpWsBlock
		}
	}

	var insertBlock []marketplacemodel.BlockChainBlock
	var updateBlock []marketplacemodel.BlockChainBlock
	var listIdUpdate []uint64
	for _, vlog := range logs {
		log.Printf("ProcessBloc: %d", vlog.BlockNumber)
		fmt.Println("GetOpvMarketplaceSellTopic: ", s.cfg.GetOpvMarketplaceSellTopic())
		fmt.Println("GetOpvMarketplaceAddressFirst: ", s.cfg.GetOpvMarketplaceAddressFirst())
		fmt.Println("GetOpvMarketplaceAddressSecond: ", s.cfg.GetOpvMarketplaceAddressSecond())
		fmt.Println("GetOpvMarketplaceAddressThird: ", s.cfg.GetOpvMarketplaceAddressThird())
		if _, ok := rebuildWsBlock[vlog.BlockNumber]; ok {
			log.Printf("Processed: %d", vlog.BlockNumber)
			tmpBlock := rebuildWsBlock[vlog.BlockNumber]
			tmpBlock.Status = common2.BLOCK_STATUS_SUCCESS
			tmpBlock.IsBlockProcess = common2.IS_BLOCK_PROCESS
			updateBlock = append(updateBlock, tmpBlock)
			listIdUpdate = append(listIdUpdate)
		} else {
			isGotProcess := false
			for _, topic := range vlog.Topics {
				switch topic {
				case s.cfg.GetOpvMarketplaceBuyTopic():
					log.Printf("Block: %v, Buy : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					buy, _ := marketplaceFilterSec.ParseBuyMarketItem(vlog)
					go s.buyItem(buy, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvMarketplaceSellTopic():
					log.Printf("Block: %v, SEll : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					sell, _ := marketplaceFilterSec.ParseCreateMarketItem(vlog)
					fmt.Println("GetOpvMarketplaceSellTopic: 11111111111111111111111", sell)
					go s.createSellItem(sell, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvMarketplaceCancelSellTopic():
					log.Printf("Block: %v, Cancel Sell : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					cancelSell, _ := marketplaceFilterSec.ParseCancelSell(vlog)
					go s.cancelSellItem(cancelSell, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNftCreateEventTopic():
					log.Printf("Block: %v, Event : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					event, _ := eventFilter.ParseJoinEvent(vlog)
					s.createEvent(event, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNFTCancelEventTopic():
					log.Printf("Block: %v, Cancel Event : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					cancelEvent, _ := eventFilter.ParseOutEvent(vlog)
					s.cancelEvent(cancelEvent, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNftCreateEventSecTopic():
					log.Printf("Block: %v, Event Sec : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					event, _ := eventFilterSec.ParseJoinEvent(vlog)
					log.Println(event)
					s.createEventSec(event, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNFTCancelEventSecTopic():
					log.Printf("Block: %v, Cancel Event Sec : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					cancelEvent, _ := eventFilterSec.ParseOutEvent(vlog)
					s.cancelEventSec(cancelEvent, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNFTTransferTopic():
					log.Printf("Block: %v, Transfer : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					transfer, _ := transferFilter.ParseTransferNFT(vlog)
					s.transferNft(transfer, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNftMintTopic():
					log.Printf("Block: %v, Mint : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					mintLog, _ := createNftFilter.ParseLogBatchMint(vlog)
					s.mintNft(mintLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvAuctionTopic():
					log.Printf("Block: %v, Auction : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					auctionLog, _ := marketplaceFilterSec.ParseCreateAuction(vlog)
					fmt.Println("111111111111111111111111111111111111111111111")
					fmt.Println(auctionLog)
					s.registerAuction(auctionLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvAuctionBidTopic():
					log.Printf("Block: %v, Auction BID : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					auctionLog, _ := marketplaceFilterSec.ParseNewBid(vlog)
					s.auctionBid(auctionLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvAuctionClaimNftTopic():
					log.Printf("Block: %v, Auction Claim nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					auctionLog, _ := marketplaceFilterSec.ParseClaimNft(vlog)
					s.auctionClaimNft(auctionLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvAuctionCancelTopic():
					log.Printf("Block: %v, Auction Claim nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					auctionLog, _ := marketplaceFilterSec.ParseCancelSellAuction(vlog)
					s.auctionCancel(auctionLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNewPackTopic():
					log.Printf("Block: %v, Auction Claim nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					tmpLog, _ := marketplaceFilterSec.ParseNewPack(vlog)
					s.addPackSale(tmpLog, ctx)
					isGotProcess = true
					break
				case s.cfg.GetOpvNewPackAuctionTopic():
					log.Printf("Block: %v, Auction Claim nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
					tmpLog, _ := marketplaceFilterSec.ParseNewAuctionPack(vlog)
					s.addPackSaleAuction(tmpLog, ctx)
					isGotProcess = true
					break
					/*case s.cfg.GetOpvTransferTopic():
						log.Printf("Block: %v, Transfer nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
						auctionLog, _ := factoryFilter.ParseTransfer(vlog)
						s.selfTransferNft(auctionLog, ctx)
						isGotProcess = true
						break
					case s.cfg.GetOpvMultiTransferTopic():
						log.Printf("Block: %v, Transfer nft : %v\n", vlog.BlockNumber, vlog.TxHash.Hex())
						auctionLog, _ := createNftFilter.ParseTransfer(vlog)
						s.selfMultiTransferNft(auctionLog, ctx)
						isGotProcess = true
						break*/
				}
			}
			if isGotProcess {
				var tmpBlock marketplacemodel.BlockChainBlock
				tmpBlock.BlockNumber = vlog.BlockNumber
				tmpBlock.Status = common2.BLOCK_STATUS_SUCCESS
				tmpBlock.IsBlockProcess = common2.IS_BLOCK_PROCESS
				insertBlock = append(insertBlock, tmpBlock)
			}
		}
	}
	//rerun add block
	s.addPackSaleReBuild(ctx)
	if len(insertBlock) > 0 {
		s.marketSvc.Store.InsertBlock(insertBlock)
	}
	log.Printf("Update ErrorBlock: ")
	log.Printf("Count listIdUpdate: %d", len(listIdUpdate))
	if len(listIdUpdate) > 0 {
		log.Printf("Update number of record: %d", len(listIdUpdate))
		err := s.marketSvc.Store.UpdateWsBlockSuccess(updateBlock, listIdUpdate)
		if err != nil {
			log.Printf("Update Error: %s", err.Error())
		}
	}
}
