package blockchain

import (
	"api-openlive/blockchain/create_event"
	"api-openlive/blockchain/create_event_sec"
	common2 "api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	"context"
	"math/big"
)

func (s Synchronization) createEvent(log *create_event.CreateEventJoinEvent, ctx context.Context) {
	s.logger.Infof("CreateEvent Nft %v User %v", log.NftId.String(), log.User.Hex())
	joined, _ := s.marketSvc.Store.CheckItemJoinEvent(log.NftId.String())
	if joined {
		s.logger.Infof("CreateEvent Item: %v exist", log.NftId.String())
		return
	}
	users, err := s.userSvc.GetUsersByLogCreateEvent(log.User)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqEvent := marketplacemodel.UpdateItemEvent{}
	reqEvent.Nft = log.NftId.String()
	reqEvent.ContractAddress = log.ContractAddress.String()
	reqEvent.UserId = users.Id
	oldItem, _ := s.marketSvc.Store.FindItem(&marketplacemodel.Item{UserId: reqEvent.UserId, Nft: reqEvent.Nft, ContractAddress: reqEvent.ContractAddress})
	reqEvent.ItemId = oldItem.Id
	reqEvent.Revenue = "0"
	reqEvent.UserDisplay = users.DisplayName
	reqEvent.UserAvatar = users.Avatar
	reqEvent.ItemName = oldItem.Title
	reqEvent.WalletAddress = users.WalletAddress
	if err := s.marketSvc.JoinEventNft(reqEvent); err != nil {
		s.logger.Errorf("Update join event nft: %v", err.Error())
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
	s.logger.Infof("Join event successfully")
}

func (s Synchronization) cancelEvent(log *create_event.CreateEventOutEvent, ctx context.Context) {
	s.logger.Infof("CancelEvent Nft %v User %v", log.NftId.String(), log.User.Hex())
	joined, _ := s.marketSvc.Store.CheckItemJoinEvent(log.NftId.String())
	if !joined {
		s.logger.Infof("CancelEvent Item: %v not join event", log.NftId.String())
		return
	}
	users, err := s.userSvc.GetUsersByLogCreateEvent(log.User)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqEvent := marketplacemodel.UpdateItemEvent{}
	reqEvent.Nft = log.NftId.String()
	reqEvent.ContractAddress = log.ContractAddress.String()
	reqEvent.UserId = users.Id
	reqEvent.UserDisplay = users.DisplayName
	reqEvent.UserAvatar = users.Avatar
	reqEvent.WalletAddress = users.WalletAddress

	if err := s.marketSvc.CancelEventNft(reqEvent); err != nil {
		s.logger.Errorf("Update cancel event nft: %v", err.Error())
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
	s.logger.Infof("Cancel event successfully")
}

func (s Synchronization) createEventSec(log *create_event_sec.CreateEventJoinEvent, ctx context.Context) {
	s.logger.Infof("CreateEvent Nft %v User %v", log.NftId.String(), log.User.Hex())
	joined, _ := s.marketSvc.Store.CheckItemJoinEvent(log.NftId.String())
	if joined {
		s.logger.Infof("CreateEvent Item: %v exist", log.NftId.String())
		return
	}
	users, err := s.userSvc.GetUsersByLogCreateEvent(log.User)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqEvent := marketplacemodel.UpdateItemEvent{}
	reqEvent.Nft = log.NftId.String()
	reqEvent.ContractAddress = log.ContractAddress.String()
	reqEvent.UserId = users.Id
	oldItem, _ := s.marketSvc.Store.FindItem(&marketplacemodel.Item{UserId: reqEvent.UserId, Nft: reqEvent.Nft, ContractAddress: reqEvent.ContractAddress})
	reqEvent.ItemId = oldItem.Id
	reqEvent.Revenue = log.FundToRef.String()
	reqEvent.UserDisplay = users.DisplayName
	reqEvent.UserAvatar = users.Avatar
	reqEvent.ItemName = oldItem.Title
	reqEvent.WalletAddress = users.WalletAddress
	if err := s.marketSvc.JoinEventNft(reqEvent); err != nil {
		s.logger.Errorf("Update join event nft: %v", err.Error())
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
	s.logger.Infof("Join event successfully")
}

func (s Synchronization) cancelEventSec(log *create_event_sec.CreateEventOutEvent, ctx context.Context) {
	s.logger.Infof("CancelEvent Nft %v User %v", log.NftId.String(), log.User.Hex())
	joined, _ := s.marketSvc.Store.CheckItemJoinEvent(log.NftId.String())
	if !joined {
		s.logger.Infof("CancelEvent Item: %v not join event", log.NftId.String())
		return
	}
	users, err := s.userSvc.GetUsersByLogCreateEvent(log.User)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	reqEvent := marketplacemodel.UpdateItemEvent{}
	reqEvent.Nft = log.NftId.String()
	reqEvent.ContractAddress = log.ContractAddress.String()
	reqEvent.UserId = users.Id
	reqEvent.UserDisplay = users.DisplayName
	reqEvent.UserAvatar = users.Avatar
	reqEvent.WalletAddress = users.WalletAddress

	if err := s.marketSvc.CancelEventNft(reqEvent); err != nil {
		s.logger.Errorf("Update cancel event nft: %v", err.Error())
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
	s.logger.Infof("Cancel event successfully")
}
