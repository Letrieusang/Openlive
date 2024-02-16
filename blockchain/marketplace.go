package blockchain

import (
	"api-openlive/blockchain/marketplace"
	common2 "api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

func (s Synchronization) buyItem(log *marketplace.MarketplaceBuyMarketItem, ctx context.Context) {
	s.logger.Infof("Create Buy: %v", log.Raw.TxHash.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		IdOnMarket: log.IdOnMarket.Int64(),
		Status:     common2.MARKET_PLACE_STATUS_SOLDED,
	}
	total, _ := s.marketSvc.Store.CountMarketPlace(rqFilter)
	if total > 0 {
		s.logger.Infof("Create buy Item: %v exist", log.Raw.TxHash.Hex())
		return
	}

	rq := marketplacemodel.CreateMarketPlaceSoldRequest{
		WalletAddress:   log.Owner.Hex(),
		IdOnMarket:      log.IdOnMarket.Int64(),
		TransactionId:   log.Raw.TxHash.Hex(),
		ContractAddress: log.NftContract.Hex(),
		Price:           log.Price.String(),
		MarketContract:  log.Raw.Address.String(),
	}
	err := s.marketSvc.CreateMarketPlaceSold(rq)
	if err != nil {
		s.logger.Errorf("Buy Item err: %v", err.Error())
	}
	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Infof("Header Block Error: ", err.Error())
	}
	finalizeBlock.SetString(header.Number.String(), 10)
	s.logger.Infof("Header Block Number: ", header.Number)
	newBlock := marketplacemodel.BlockChainBlock{BlockNumber: finalizeBlock.Uint64(), Status: common2.BLOCK_STATUS_WS_SUCCESS, IsBlockProcess: common2.IS_NOT_BLOCK_PROCESS}
	errCreateBlock := s.marketSvc.Store.AddWsBlock(newBlock)
	if errCreateBlock != nil {
		s.logger.Infof("Create Block Error: ", errCreateBlock.Error())
	}
}

func (s Synchronization) createSellItem(log *marketplace.MarketplaceCreateMarketItem, ctx context.Context) {
	s.logger.Infof("Create sell Item: %v", log.Raw.TxHash.Hex())
	fmt.Printf("Create sell Item: %v", log.Raw.TxHash.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		ContractAddress: log.Raw.Address.Hex(),
		IdOnMarket:      log.IdOnMarket.Int64(),
		Status:          common2.MARKET_PLACE_STATUS_SELLING,
	}
	total, _ := s.marketSvc.Store.CountMarketPlace(rqFilter)
	if total > 0 {
		s.logger.Infof("Create sell Item: %v exist", log.Raw.TxHash.Hex())
		fmt.Printf("Create sell Item: %v exist", log.Raw.TxHash.Hex())
		return
	}

	rq := marketplacemodel.CreateMarketPlaceRequest{
		WalletAddress:   log.Seller.String(),
		Nft:             log.TokenId.String(),
		IdOnMarket:      log.IdOnMarket.Int64(),
		TransactionId:   log.Raw.TxHash.Hex(),
		ContractAddress: log.Raw.Address.Hex(),
		Price:           log.Price.String(),
		SellType:        common2.MARKET_PLACE_SELL_TYPE_SELL,
	}
	err := s.marketSvc.CreateMarketPlace(rq, log.NftContract.Hex())
	if err != nil {
		s.logger.Errorf("Create Sell Err: %v", err.Error())
		fmt.Printf("Create Sell Err: %v", err.Error())
		return
	}

	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Infof("Header Block Error: ", err.Error())
		fmt.Printf("Header Block Error: ", err.Error())
		return
	}
	finalizeBlock.SetString(header.Number.String(), 10)
	s.logger.Infof("Header Block Number: ", header.Number)
	newBlock := marketplacemodel.BlockChainBlock{BlockNumber: finalizeBlock.Uint64(), Status: common2.BLOCK_STATUS_WS_SUCCESS, IsBlockProcess: common2.IS_NOT_BLOCK_PROCESS}
	errCreateBlock := s.marketSvc.Store.AddWsBlock(newBlock)
	if errCreateBlock != nil {
		s.logger.Infof("Create Block Error: ", errCreateBlock.Error())
	}

}

func (s Synchronization) cancelSellItem(log *marketplace.MarketplaceCancelSell, ctx context.Context) {
	s.logger.Infof("Cancel sell Item: %v", log.Raw.TxHash.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		IdOnMarket: log.IdOnMarket.Int64(),
	}
	total, _ := s.marketSvc.Store.CountMarketPlace(rqFilter)
	if total < 1 {
		s.logger.Infof("Cancel sell Item: %v exist", log.Raw.TxHash.Hex())
		return
	}
	err := s.marketSvc.CancelSellItem(&marketplacemodel.CancelMarketplace{
		Nft:        log.TokenId.String(),
		IdOnMarket: log.IdOnMarket.Int64(),
	})
	if err != nil {
		s.logger.Errorf("Cancel sell err: %v", err.Error())
	}
	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Infof("Header Block Error: ", err.Error())
	}
	finalizeBlock.SetString(header.Number.String(), 10)
	s.logger.Infof("Header Block Number: ", header.Number)
	newBlock := marketplacemodel.BlockChainBlock{BlockNumber: finalizeBlock.Uint64(), Status: common2.BLOCK_STATUS_WS_SUCCESS, IsBlockProcess: common2.IS_NOT_BLOCK_PROCESS}
	errCreateBlock := s.marketSvc.Store.AddWsBlock(newBlock)
	if errCreateBlock != nil {
		s.logger.Infof("Create Block Error: ", errCreateBlock.Error())
	}
}

func (s Synchronization) addPackSale(log *marketplace.MarketplaceNewPack, ctx context.Context) {
	var idOnMarket []int64
	for _, id := range log.IdsOnMarket {
		idOnMarket = append(idOnMarket, id.Int64())
	}
	if isGo := s.BuildAddPackRedis(idOnMarket, log.Pack.Int64(), common2.MARKET_PLACE_SELL_TYPE_SELL, log.Raw.TxHash.Hex(), true); !isGo {
		return
	}
	s.addPackProcess(ctx, idOnMarket, log.Pack.Int64(), common2.MARKET_PLACE_SELL_TYPE_SELL, log.Raw.TxHash.Hex())
}

func (s Synchronization) addPackSaleAuction(log *marketplace.MarketplaceNewAuctionPack, ctx context.Context) {
	var idOnMarket []int64
	for _, id := range log.StorageIdOnMarket {
		idOnMarket = append(idOnMarket, id.Int64())
	}
	if isGo := s.BuildAddPackRedis(idOnMarket, log.PackId.Int64(), common2.MARKET_PLACE_SELL_TYPE_AUCTION, log.Raw.TxHash.Hex(), true); !isGo {
		return
	}
	s.addPackProcess(ctx, idOnMarket, log.PackId.Int64(), common2.MARKET_PLACE_SELL_TYPE_AUCTION, log.Raw.TxHash.Hex())
}

func (s Synchronization) addPackSaleReBuild(ctx context.Context) {
	newPackInfoJson, errGetCache := s.cache.Get(context.Background(), "addPackProcess").Result()
	if errGetCache != nil {
		return
	}
	var newPackInfos []marketplacemodel.NewPackInfo
	err := json.Unmarshal([]byte(newPackInfoJson), &newPackInfos)
	fmt.Println("??????????????????????????????????????")
	fmt.Println(len(newPackInfos))
	if err != nil {
		return
	}
	var checkNewPackInfos = newPackInfos
	for index, newPackInfo := range newPackInfos {
		if isGo := s.BuildAddPackRedis(newPackInfo.IdOnMarket, newPackInfo.PackId, newPackInfo.SellType, newPackInfo.Transaction, false); !isGo {
			continue
		}
		if isOk := s.addPackProcess(ctx, newPackInfo.IdOnMarket, newPackInfo.PackId, newPackInfo.SellType, newPackInfo.Transaction); isOk {
			checkNewPackInfos = RemoveIndex(checkNewPackInfos, index)
		}
	}
	fmt.Println("2222222222222222222222222222222222222222222")
	if len(checkNewPackInfos) < 1 {
		if stt := s.cache.Del(context.Background(), "addPackProcess"); stt.Err() != nil {
			fmt.Println("3333333333333333333333333333333")
			fmt.Println(stt.Err().Error())
			return
		}
	} else {
		tmpNewPackInfoJson, _ := json.Marshal(checkNewPackInfos)
		if stt := s.cache.Set(context.Background(), "addPackProcess", string(tmpNewPackInfoJson), 10*time.Minute); stt.Err() != nil {
			return
		}
	}
}

func (s Synchronization) addPackProcess(ctx context.Context, idOnMarket []int64, packId int64, sellType int, transaction string) bool {
	s.logger.Infof("Create sell Item: %v", transaction)

	err := s.marketSvc.Store.UpdateMarketPlacePack(idOnMarket, transaction, sellType, packId)
	if err != nil {
		s.logger.Errorf("Update Pack Sell Err: %v", err.Error())
		return false
	}

	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	if err != nil {
		s.logger.Infof("Header Block Error: ", err.Error())
		return false
	}
	finalizeBlock.SetString(header.Number.String(), 10)
	s.logger.Infof("Header Block Number: ", header.Number)
	newBlock := marketplacemodel.BlockChainBlock{BlockNumber: finalizeBlock.Uint64(), Status: common2.BLOCK_STATUS_WS_SUCCESS, IsBlockProcess: common2.IS_NOT_BLOCK_PROCESS}
	errCreateBlock := s.marketSvc.Store.AddWsBlock(newBlock)
	if errCreateBlock != nil {
		s.logger.Infof("Create Block Error: ", errCreateBlock.Error())
		return false
	}
	return true
}

func (s Synchronization) BuildAddPackRedis(idOnMarket []int64, packId int64, sellType int, transaction string, isBuildRedis bool) bool {
	isGo := true
	//count
	total, errTotal := s.marketSvc.Store.CountMarketPlacePack(idOnMarket, transaction, sellType)
	if errTotal != nil || int(total) != len(idOnMarket) {
		isGo = false
		if !isBuildRedis {
			return isGo
		}

		newPackInfo := marketplacemodel.NewPackInfo{
			IdOnMarket:  idOnMarket,
			PackId:      packId,
			SellType:    sellType,
			Transaction: transaction,
		}
		var newPackInfos []marketplacemodel.NewPackInfo
		newPackInfoJson, errGetCache := s.cache.Get(context.Background(), "addPackProcess").Result()
		if errGetCache == nil {
			err := json.Unmarshal([]byte(newPackInfoJson), &newPackInfos)
			if err != nil {
				return isGo
			}

		}
		newPackInfos = append(newPackInfos, newPackInfo)
		tmpNewPackInfoJson, _ := json.Marshal(newPackInfos)
		if stt := s.cache.Set(context.Background(), "addPackProcess", string(tmpNewPackInfoJson), 10*time.Minute); stt.Err() != nil {
			return isGo
		}
		return isGo
	}
	return isGo
}

func RemoveIndex(s []marketplacemodel.NewPackInfo, index int) []marketplacemodel.NewPackInfo {
	return append(s[:index], s[index+1:]...)
}
