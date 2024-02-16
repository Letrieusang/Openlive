package blockchain

import (
	"api-openlive/blockchain/marketplace"
	common2 "api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"math/big"
)

func (s Synchronization) registerAuction(log *marketplace.MarketplaceCreateAuction, ctx context.Context) {
	if log == nil {
		return
	}
	s.logger.Infof("Auction Nft %v User %v", log.TokenId.String(), log.Seller.Hex())
	fmt.Printf("Auction Nft %v User %v", log.TokenId.String(), log.Seller.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		MarketContract: log.Raw.Address.Hex(),
		IdOnMarket:     log.IdOnMarket.Int64(),
		Status:         common2.MARKET_PLACE_STATUS_SELLING,
	}
	total, _ := s.marketSvc.Store.CountMarketPlace(rqFilter)
	if total > 0 {
		s.logger.Infof("Create Auction Item: %v exist", log.Raw.TxHash.Hex())
		fmt.Printf("Create Auction Item: %v exist", log.Raw.TxHash.Hex())
		return
	}
	rq := marketplacemodel.CreateMarketPlaceRequest{
		WalletAddress:    log.Seller.String(),
		Nft:              log.TokenId.String(),
		IdOnMarket:       log.IdOnMarket.Int64(),
		TransactionId:    log.Raw.TxHash.Hex(),
		ContractAddress:  log.Raw.Address.Hex(),
		Price:            log.MinBid.String(),
		StartPrice:       log.MinBid.String(),
		SellType:         common2.MARKET_PLACE_SELL_TYPE_AUCTION,
		AuctionStartTime: log.StartTime.String(),
		AuctionEndTime:   log.EndTime.String(),
	}
	err := s.marketSvc.CreateMarketPlace(rq, log.NftContract.Hex())
	if err != nil {
		s.logger.Errorf("Create Auction Err: %v", err.Error())
		fmt.Printf("Create Auction Err: %v", err.Error())
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
		fmt.Printf("Create Block Error: ", errCreateBlock.Error())
	}
}

func (s Synchronization) auctionBid(log *marketplace.MarketplaceNewBid, ctx context.Context) {
	if log == nil {
		return
	}
	s.logger.Infof("Auction Nft Bid %v User %v", log.TokenId.String(), log.Bidder.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		MarketContract: log.Raw.Address.Hex(),
		IdOnMarket:     log.IdOnMarket.Int64(),
		SellType:       common2.MARKET_PLACE_SELL_TYPE_AUCTION,
		Status:         common2.MARKET_PLACE_STATUS_SELLING,
	}
	auction, err := s.marketSvc.Store.GetMarketPlaceByIdOnMarket(&rqFilter)
	if err != nil || auction.Id < 1 {
		s.logger.Infof("Auction : %v not exist", log.Raw.TxHash.Hex())
		return
	}
	users, err := s.userSvc.GetUsersByLogCreateNft(log.Bidder)
	if err != nil {
		s.logger.Errorf("Get Users error %v", err.Error())
		return
	}
	rq := marketplacemodel.AuctionBid{
		Price:         log.Price.String(),
		MarketPlaceId: auction.Id,
		UserId:        users.Id,
	}

	errAuctionBid := s.marketSvc.Store.SaveAuctionBid(rq)
	if errAuctionBid != nil {
		s.logger.Errorf("Create auction bid Err: %v", err.Error())
	}
	userAvatar := utils.RebuildUrlPath(users.Avatar)
	s.AuctionNotification(auction.Item.Title, users.DisplayName, userAvatar, users.WalletAddress, users.Id, common2.NOTIFICATION_AUCTION_BID)
	metadata, _ := json.Marshal(map[string]string{
		"price": rq.Price,
	})
	s.AuctionActivity(string(metadata), users.Id, auction.Item.Id, auction.Item.CollectionId, common2.BID_ITEM)

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

func (s Synchronization) auctionClaimNft(log *marketplace.MarketplaceClaimNft, ctx context.Context) {
	if log == nil {
		return
	}
	s.logger.Infof("Auction Nft %v User %v", log.TokenId.String(), log.Seller.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		MarketContract: log.Raw.Address.Hex(),
		IdOnMarket:     log.IdOnMarket.Int64(),
		SellType:       common2.MARKET_PLACE_SELL_TYPE_AUCTION,
		Status:         common2.MARKET_PLACE_STATUS_SELLING,
	}
	auction, err := s.marketSvc.Store.GetMarketPlaceByIdOnMarket(&rqFilter)
	if err != nil || auction.Id < 1 {
		s.logger.Infof("Auction : %v not exist", log.Raw.TxHash.Hex())
		return
	}

	rq := marketplacemodel.CreateMarketPlaceSoldRequest{
		WalletAddress:   log.Buyer.Hex(),
		IdOnMarket:      log.IdOnMarket.Int64(),
		TransactionId:   log.Raw.TxHash.Hex(),
		ContractAddress: log.NftContract.Hex(),
		Price:           log.Price.String(),
		MarketContract:  log.Raw.Address.String(),
		SellType:        common2.MARKET_PLACE_SELL_TYPE_AUCTION,
	}
	errMkSold := s.marketSvc.CreateMarketPlaceSold(rq)
	if errMkSold != nil {
		s.logger.Errorf("Buy Item err: %v", errMkSold.Error())
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

func (s Synchronization) auctionCancel(log *marketplace.MarketplaceCancelSellAuction, ctx context.Context) {
	if log == nil {
		return
	}
	s.logger.Infof("Auction Nft %v User %v", log.TokenId.String(), log.Seller.Hex())
	rqFilter := marketplacemodel.MarketPlaceFilter{
		MarketContract: log.Raw.Address.Hex(),
		IdOnMarket:     log.IdOnMarket.Int64(),
		SellType:       common2.MARKET_PLACE_SELL_TYPE_AUCTION,
		Status:         common2.MARKET_PLACE_STATUS_SELLING,
	}
	total, _ := s.marketSvc.Store.CountMarketPlace(rqFilter)
	if total < 1 {
		s.logger.Infof("Auction : %v not exist", log.Raw.TxHash.Hex())
		return
	}

	err := s.marketSvc.CancelSellItem(&marketplacemodel.CancelMarketplace{
		Nft:             log.TokenId.String(),
		IdOnMarket:      log.IdOnMarket.Int64(),
		ContractAddress: log.Raw.Address.Hex(),
		SellType:        common2.MARKET_PLACE_SELL_TYPE_AUCTION,
	})
	if err != nil {
		s.logger.Errorf("Cancel aution err: %v", err.Error())
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

func (s Synchronization) AuctionNotification(title, displayName, avatar, walletAddress string, userId, notifyType int64) {
	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      title,
		"user_name":      displayName,
		"user_avatar":    avatar,
		"wallet_address": walletAddress,
	})
	var listNoti []*usermodel.Notification
	listNoti = append(listNoti, &usermodel.Notification{
		UserId:             userId,
		LastTimeRead:       0,
		NotificationType:   notifyType,
		Status:             common2.UNREAD,
		NotificationObject: base58.Encode(notiObj),
	})

	s.userSvc.Store.BatchInsertNotification(listNoti)
}

func (s Synchronization) AuctionActivity(metadata string, userId, itemId, collectionId, activityType int64) {
	s.userSvc.Store.InsertActivities(&usermodel.Activities{
		UserId:       userId,
		ItemId:       itemId,
		MetaData:     metadata,
		ActivityType: activityType,
		CollectionId: collectionId,
	})
}
