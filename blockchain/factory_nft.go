package blockchain

import (
	"api-openlive/blockchain/factory"
	"api-openlive/blockchain/opv_multiple_factory"
	common2 "api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (s Synchronization) mintNft(log *opv_multiple_factory.OpvMultipleFactoryLogBatchMint, ctx context.Context) {
	if log == nil {
		return
	}
	s.logger.Infof("Mint Nft %v User %v", log.FromId.String(), log.User.Hex())
	users, err := s.userSvc.GetUsersByLogCreateNft(log.User)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	var itemFilter marketplacemodel.ItemFilter
	itemFilter.Nft = ""
	itemFilter.ContractAddress = ""
	itemFilter.Id = log.UniqueId.Int64()
	itemFilter.UserId = users.Id
	itemFilter.Status = common2.ITEM_CREATE_WAITING_SYNC
	itemFilter.OrderBy = "items.id asc"
	if log.ToId.Int64() == 0 {
		log.ToId = log.FromId
	}
	if log.ToId.Int64() > log.FromId.Int64() {
		itemFilter.IsGetMulti = 1
		itemFilter.IsGetMultiId = log.UniqueId.Int64()
		itemFilter.Id = 0
	}

	items, err := s.marketSvc.Store.GetItemWithCondition(itemFilter, nil)
	if err != nil {
		s.logger.Infof("CreateItem Item: %v not found", log.FromId)
		return
	}
	if len(items) < 1 {
		s.logger.Infof("CreateItem Item: %v not found", log.FromId)
		return
	}
	header, err := s.c.client.HeaderByNumber(ctx, nil)
	i := log.FromId.Int64()

	for index, _ := range items {
		if i > log.ToId.Int64() {
			break
		}
		items[index].Status = common2.ITEM_NEW
		items[index].ContractAddress = log.Raw.Address.Hex()
		items[index].Nft = fmt.Sprint(i)
		items[index].TransactionId = header.TxHash.Hex()
		i++
	}
	if errUpdate := s.marketSvc.Store.UpdateItemSync(items); errUpdate != nil {
		s.logger.Infof("CreateItem Update Item: %v Error", log.FromId)
		return
	}
	s.userSvc.BuildUserDailyEvent(users.Id, common2.DAILY_EVENT_TYPE_MINT_NFT)
	//insert to blockchainblock
	finalizeBlock := big.NewInt(0)
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
	s.logger.Infof("Join event successfully")
}

func (s Synchronization) selfTransferNft(log *factory.FactoryTransfer, ctx context.Context) {
	s.selfTransferNftProcess(log.TokenId.String(), log.Raw.Address.Hex(), log.From, log.To, ctx)
}
func (s Synchronization) selfMultiTransferNft(log *opv_multiple_factory.OpvMultipleFactoryTransfer, ctx context.Context) {
	s.selfTransferNftProcess(log.TokenId.String(), log.Raw.Address.Hex(), log.From, log.To, ctx)
}

func (s Synchronization) selfTransferNftProcess(tokenId, contractAddress string, from, to common.Address, ctx context.Context) {
	s.logger.Infof("Self Transfer Nft %v From %v To %v", tokenId, from.Hex(), to.Hex())
	users, err := s.userSvc.GetUsersByLogTransfer(from, to)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqUpdate := marketplacemodel.UpdateItemOwnerRequest{}
	for _, u := range users {
		if u.WalletAddress == from.Hex() {
			reqUpdate.SenderId = u.Id
			continue
		}
		if u.WalletAddress == to.Hex() {
			reqUpdate.ReceiverId = u.Id
			continue
		}
	}
	if reqUpdate.SenderId == 0 {
		s.logger.Errorf("Sender or receiver is invalid")
		return
	}
	reqUpdate.Nft = tokenId
	reqUpdate.ContractAddress = contractAddress
	if err := s.marketSvc.TransferNft(reqUpdate); err != nil {
		s.logger.Errorf("Update owner: %v", err.Error())
		return
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
	s.logger.Infof("Update Owner successfully")
}
