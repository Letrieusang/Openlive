package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	marketplacestorage "api-openlive/module/market_place/storage"
	usermodel "api-openlive/module/user/model"
	userstorage "api-openlive/module/user/storage"
	"api-openlive/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"math/big"
	"time"
)

type MarketPlaceController struct {
	Store     marketplacestorage.Storage
	UserStore userstorage.Storage
}

func (i MarketPlaceController) CreateMarketPlace(rq marketplacemodel.CreateMarketPlaceRequest, createContract string) error {
	userFilter := usermodel.UserFilter{
		WalletAddress: rq.WalletAddress,
	}
	user, _ := i.UserStore.GetUserByCondition(userFilter)
	itemFilter := marketplacemodel.ItemFilter{
		Nft:             rq.Nft,
		ContractAddress: createContract,
	}
	item, _ := i.Store.GetItemByID(itemFilter)
	marketInfo := marketplacemodel.MarketPlace{
		TransactionId:    rq.TransactionId,
		SellerId:         user.Id,
		ContractAddress:  rq.ContractAddress,
		ItemId:           item.Id,
		Status:           common.MARKET_PLACE_STATUS_SELLING,
		Price:            rq.Price,
		IdOnMarket:       rq.IdOnMarket,
		SellType:         rq.SellType,
		AuctionStartTime: rq.AuctionStartTime,
		AuctionEndTime:   rq.AuctionEndTime,
	}
	if marketInfo.SellType != common.MARKET_PLACE_SELL_TYPE_AUCTION && marketInfo.SellType != common.MARKET_PLACE_SELL_TYPE_SELL {
		marketInfo.SellType = common.MARKET_PLACE_SELL_TYPE_SELL
	}
	if marketInfo.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
		if len(rq.StartPrice) < 1 {
			return errors.New("start price invalid")
		}
		marketInfo.Price = rq.StartPrice
		marketInfo.StartPrice = rq.StartPrice
		if len(marketInfo.Price) < 1 {
			marketInfo.Price = rq.StartPrice
		}
	}
	err := i.Store.InsertMarketPlace(&marketInfo)
	if err != nil {
		return err
	}

	// notification
	listFollowingUser, err := i.UserStore.ListFollowingUser(&usermodel.UserFollow{FollowerId: user.Id})
	if err != nil {
		return err
	}
	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      item.Title,
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"wallet_address": user.WalletAddress,
	})

	listNoti := make([]*usermodel.Notification, 0, len(listFollowingUser))
	for _, v := range listFollowingUser {
		notification := &usermodel.Notification{
			UserId:             v.UserId,
			LastTimeRead:       0,
			NotificationType:   common.NOTIFICATION_LISTED,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		listNoti = append(listNoti, notification)
	}
	activityType := common.LISTED_ITEM
	notifyType := common.SELL
	if rq.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
		activityType = common.AUCTION_ITEM
		notifyType = common.NOTIFICATION_AUCTION
	}
	listNoti = append(listNoti, &usermodel.Notification{
		UserId:             user.Id,
		LastTimeRead:       0,
		NotificationType:   int64(notifyType),
		Status:             common.UNREAD,
		NotificationObject: base58.Encode(notiObj),
	})

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return err
		}
	}

	metadata, _ := json.Marshal(map[string]string{
		"price": rq.Price,
	})

	err = i.UserStore.InsertActivities(&usermodel.Activities{
		UserId:       user.Id,
		ItemId:       item.Id,
		MetaData:     string(metadata),
		ActivityType: int64(activityType),
		CollectionId: item.CollectionId,
	})
	if err != nil {
		return err
	}
	i.BuildUserDailyEvent(user.Id, common.DAILY_EVENT_TYPE_SELL_NFT)
	return nil
}

func (i MarketPlaceController) UpdateMarketPlaceCms(rq *marketplacemodel.UpdateMarketPlaceRequest) error {
	var update = make(map[string]interface{})
	if rq.IsHot >= 0 {
		update["is_hot"] = rq.IsHot
	}
	if rq.DisplayOrder > 0 {
		update["display_order"] = rq.DisplayOrder
	}
	if err := i.Store.UpdateMarketPlaceCms(rq.Id, update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) ListItemMarketPlace(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.ListMarketPlaceWithCollectionNameResponse, error) {
	guestId := rq.UserId
	result, err := i.Store.ListMarketPlaceRebuild(rq, paging)
	if err != nil {
		return nil, err
	}

	listItemId := make([]int64, 0, len(result))
	var listAuctionId []int64
	for _, v := range result {
		listItemId = append(listItemId, v.ItemId)
		if v.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
			listAuctionId = append(listAuctionId, v.ItemId)
		}
	}

	var listAuctionBid []marketplacemodel.AuctionBid
	if len(listAuctionId) > 0 {
		var filterAuctionBid marketplacemodel.GetAuctionFilter
		filterAuctionBid.MarketPlaceIds = listAuctionId
		filterAuctionBid.OrderBy = "auction_bids.id desc"
		filterAuctionBid.GroupBy = "market_place_id"
		var errBid error
		listAuctionBid, errBid = i.Store.GetAuctionBid(filterAuctionBid, nil)
		if errBid != nil {
			return nil, errBid
		}
	}

	mapLiked := make(map[int64]bool, 100)
	mapLikeItem := make(map[int64][]usermodel.UserDetail, 100)
	if rq.UserId > 0 {
		// list liked item
		listLiked, err := i.Store.ListLikedItem(listItemId, rq.UserId)
		if err != nil {
			return nil, err
		}
		for _, v := range listLiked {
			mapLiked[v.ItemId] = true
		}

	}

	likes, err := i.Store.ListLikes(listItemId)
	if err != nil {
		return nil, err
	}

	listUserId := make([]int64, 0, len(likes))
	for _, v := range likes {
		listUserId = append(listUserId, v.UserId)
	}

	users, err := i.UserStore.GetListUserById(listUserId)
	if err != nil {
		return nil, err
	}

	for _, v := range likes {
		for _, j := range users {
			j.Avatar = utils.RebuildUrlPath(j.Avatar)
			j.Cover = utils.RebuildUrlPath(j.Cover)
			if v.UserId == j.Id && len(mapLikeItem[v.ItemId]) <= common.LIMIT_USER_LIKE {
				mapLikeItem[v.ItemId] = append(mapLikeItem[v.ItemId], j)
				continue
			}
		}
	}

	market := make([]*marketplacemodel.ListMarketPlaceWithCollectionNameResponse, 0, len(result))
	for _, v := range result {
		userLikes := make([]usermodel.UserDetail, 0, 100)
		if v.Item.IsLocking == 1 && v.Item.UserId != guestId {
			v.Item.Code = ""
		}
		v.Item.Image = utils.RebuildUrlPath(v.Item.Image)
		v.Seller.Avatar = utils.RebuildUrlPath(v.Seller.Avatar)
		v.Seller.Cover = utils.RebuildUrlPath(v.Seller.Cover)
		if mapLikeItem[v.ItemId] != nil {
			userLikes = mapLikeItem[v.ItemId]
		}

		var tmpAuction marketplacemodel.AuctionBid
		for _, bid := range listAuctionBid {
			if bid.MarketPlaceId == v.Id {
				bid.UserInfo.Avatar = utils.RebuildUrlPath(bid.UserInfo.Avatar)
				bid.UserInfo.Cover = utils.RebuildUrlPath(bid.UserInfo.Cover)
				tmpAuction = bid
			}
		}

		market = append(market, &marketplacemodel.ListMarketPlaceWithCollectionNameResponse{
			MarketPlaceWithCollectionName: v,
			AuctionBid:                    tmpAuction,
			IsLike:                        mapLiked[v.ItemId],
			UserLikes:                     userLikes,
		})
	}

	return market, nil
}

func (i MarketPlaceController) ListMostLikeItemMarketPlace(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.ListMarketPlaceResponse, error) {

	result, err := i.Store.ListMostLikeMarketPlace(paging.Limit)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	listItemId := make([]int64, 0, len(result))
	for _, v := range result {
		listItemId = append(listItemId, v.ItemId)
	}

	mapLiked := make(map[int64]bool, 100)
	mapLikeItem := make(map[int64][]usermodel.UserDetail, 100)
	if rq.UserId > 0 {
		// list liked item
		listLiked, err := i.Store.ListLikedItem(listItemId, rq.UserId)
		if err != nil {
			return nil, err
		}
		for _, v := range listLiked {
			mapLiked[v.ItemId] = true
		}

	}

	likes, err := i.Store.ListLikes(listItemId)
	if err != nil {
		return nil, err
	}

	listUserId := make([]int64, 0, len(likes))
	for _, v := range likes {
		listUserId = append(listUserId, v.UserId)
	}

	users, err := i.UserStore.GetListUserById(listUserId)
	if err != nil {
		return nil, err
	}

	for _, v := range likes {
		for _, j := range users {
			if v.UserId == j.Id && len(mapLikeItem[v.ItemId]) <= common.LIMIT_USER_LIKE {
				mapLikeItem[v.ItemId] = append(mapLikeItem[v.ItemId], j)
				continue
			}
		}
	}

	market := make([]*marketplacemodel.ListMarketPlaceResponse, 0, len(result))
	for _, v := range result {
		userLikes := make([]usermodel.UserDetail, 0, 100)
		v.Item.Image = utils.RebuildUrlPath(v.Item.Image)
		if mapLikeItem[v.ItemId] != nil {
			userLikes = mapLikeItem[v.ItemId]
		}
		market = append(market, &marketplacemodel.ListMarketPlaceResponse{
			MarketPlace: v,
			IsLike:      mapLiked[v.ItemId],
			UserLikes:   userLikes,
		})
	}

	return market, nil
}

func (i MarketPlaceController) ListVotedItemMarketPlace(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.MarketPlace, error) {
	var listId []int64
	// list id voted item
	result, err := i.Store.ListVoting(&marketplacemodel.VoteFilter{
		Cols:     []string{"market_place_id", "id"},
		Distinct: []string{"market_place_id"},
	})
	if err != nil {
		return nil, err
	}
	for _, v := range result {
		listId = append(listId, v.MarketPlaceId)
	}

	// set list item voted id
	rq.Ids = listId
	listMarket, err := i.Store.ListMarketPlace(rq, paging)
	for index, _ := range listMarket {
		listMarket[index].Item.Image = utils.RebuildUrlPath(listMarket[index].Item.Image)
	}
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return listMarket, nil
}

func (i MarketPlaceController) DetailSellItem(contractAddress, nft string, userId int64) (*marketplacemodel.DetailItemResponse, error) {

	result, err := i.Store.GetMarketPlaceByItemId(&marketplacemodel.MarketPlaceFilter{
		Nft:             nft,
		ContractAddress: contractAddress,
		Status:          common.MARKET_PLACE_STATUS_SELLING,
	})
	if err != nil {
		return nil, err
	}
	if result.Id < 1 {
		return nil, errors.New("Record not found!")
	}

	isLiked := false
	if userId > 0 {
		isLiked, err = i.Store.IsLikedItem(result.ItemId, userId)
		if err != nil {
			return nil, err
		}
	}

	result.Item.Image = utils.RebuildUrlPath(result.Item.Image)
	response := &marketplacemodel.DetailItemResponse{
		IsLike:           isLiked,
		Id:               result.Id,
		TransactionId:    result.TransactionId,
		ContractAddress:  result.ContractAddress,
		IdOnMarket:       result.IdOnMarket,
		ItemId:           result.ItemId,
		Item:             result.Item.ConvertToDetail(),
		SellerId:         result.SellerId,
		Seller:           result.Seller,
		Price:            result.Price,
		Status:           result.Status,
		TotalVote:        result.TotalVote,
		TotalLike:        result.Item.TotalLike,
		Description:      result.Description,
		SellType:         result.SellType,
		StartPrice:       result.StartPrice,
		AuctionEndTime:   result.AuctionEndTime,
		AuctionStartTime: result.AuctionStartTime,
		UserCreatedId:    result.UserCreatedId,
		UserUpdatedId:    result.UserUpdatedId,
		CreatedAt:        result.CreatedAt,
		UpdatedAt:        result.UpdatedAt,
		PackId:           result.PackId,
	}

	if result.Item.NumOfCopies > 1 {
		var countItemCopyFilter marketplacemodel.CountItemFilter
		countItemCopyFilter.Status = common.ITEM_ON_SALE
		if result.Item.MainItemId > 0 {
			countItemCopyFilter.MainItemId = result.Item.MainItemId
		} else {
			countItemCopyFilter.MainItemId = result.Item.Id
		}
		response.AvailableCopy, _ = i.Store.CountAvailableItemCopied(countItemCopyFilter)
		if result.Item.MainItemId == 0 {
			response.AvailableCopy++
		}
	}

	// check code
	if userId == result.Item.UserId {
		response.Item.Code = result.Item.Code
	}

	//get created user
	createdUser, errCreatedUser := i.UserStore.GetCreatedUserByItemId(result.Item.Id)
	if errCreatedUser != nil {
		return nil, errCreatedUser
	}
	response.ItemCreator = createdUser

	if result.Item.CollectionId > 0 {
		collection, errColl := i.Store.DetailCollection(result.Item.CollectionId)
		if errColl != nil {
			return nil, errColl
		}
		response.Collection = collection
	}
	if result.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
		var listAuctionBid []marketplacemodel.AuctionBid
		var listAuctionId []int64
		listAuctionId = append(listAuctionId, result.Id)
		var filterAuctionBid marketplacemodel.GetAuctionFilter
		filterAuctionBid.MarketPlaceIds = listAuctionId
		filterAuctionBid.OrderBy = "auction_bids.id desc"
		filterAuctionBid.IsGetUser = 1
		var errBid error
		listAuctionBid, errBid = i.Store.GetAuctionBid(filterAuctionBid, nil)
		if errBid != nil {
			return nil, errBid
		}
		for index, _ := range listAuctionBid {
			listAuctionBid[index].UserInfo.Avatar = utils.RebuildUrlPath(listAuctionBid[index].UserInfo.Avatar)
			listAuctionBid[index].UserInfo.Cover = utils.RebuildUrlPath(listAuctionBid[index].UserInfo.Cover)
		}
		response.AuctionBid = listAuctionBid
	}

	//update view
	_ = i.Store.UpdateViewItem(result.ItemId)

	return response, nil
}

func (i MarketPlaceController) CreateMarketPlaceSold(rq marketplacemodel.CreateMarketPlaceSoldRequest) error {
	userFilter := usermodel.UserFilter{
		WalletAddress: rq.WalletAddress,
	}
	user, _ := i.UserStore.GetUserByCondition(userFilter)
	marketPlace, err := i.Store.GetMarketPlaceByItemId(&marketplacemodel.MarketPlaceFilter{
		IdOnMarket:      rq.IdOnMarket,
		ContractAddress: rq.ContractAddress,
		MarketContract:  rq.MarketContract,
	})
	if err != nil {
		return err
	}
	if marketPlace.Status != common.MARKET_PLACE_STATUS_SELLING {
		return common.ErrItemNotPricing
	}
	marketInfo := marketplacemodel.MarketPlaceSold{
		TransactionId:   rq.TransactionId,
		BuyerId:         user.Id,
		ContractAddress: rq.ContractAddress,
		MarketPlaceId:   marketPlace.Id,
		Status:          2,
		Price:           rq.Price,
	}
	err = i.Store.InsertMarketPlaceSold(&marketInfo)
	if err != nil {
		return err
	}
	//save log
	transactionLog := marketplacemodel.TransactionLog{
		TransactionId:   rq.TransactionId,
		BuyerId:         user.Id,
		ContractAddress: rq.ContractAddress,
		Price:           rq.Price,
		ItemId:          marketPlace.ItemId,
		SellerId:        marketPlace.SellerId,
	}
	if err = i.Store.InsertTransactionLog(&transactionLog); err != nil {
		return err
	}

	// notification
	listFollowingUser, err := i.UserStore.ListFollowingUser(&usermodel.UserFollow{FollowerId: user.Id})
	if err != nil {
		return err
	}
	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      marketPlace.Item.Title,
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"price":          rq.Price,
		"wallet_address": user.WalletAddress,
	})
	activityType := common.PURCHASED_ITEM
	notifyType := common.BUY
	otherNotifyType := common.PURCHASED
	if rq.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
		activityType = common.AUCTION_CLAIM_ITEM
		notifyType = common.NOTIFICATION_AUCTION_CLAIM_NFT
		otherNotifyType = common.NOTIFICATION_AUCTION_OTHER_PURCHASED

	}
	listNoti := make([]*usermodel.Notification, 0, len(listFollowingUser))
	for _, v := range listFollowingUser {
		notification := &usermodel.Notification{
			UserId:             v.UserId,
			LastTimeRead:       0,
			NotificationType:   int64(otherNotifyType),
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		listNoti = append(listNoti, notification)
	}

	listNoti = append(listNoti, []*usermodel.Notification{
		{
			UserId:             marketPlace.SellerId,
			LastTimeRead:       0,
			NotificationType:   common.ORTHER_USER_SELL,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		},
		{
			UserId:             user.Id,
			LastTimeRead:       0,
			NotificationType:   int64(notifyType),
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		},
	}...)

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return err
		}
	}

	metadata, _ := json.Marshal(map[string]string{
		"price": rq.Price,
	})
	activities := []*usermodel.Activities{
		{
			UserId:         marketPlace.SellerId,
			ItemId:         marketPlace.ItemId,
			MetaData:       string(metadata),
			ActivityType:   common.SOLD_ITEM,
			ActivityUserId: user.Id,
		},
		{
			UserId:         user.Id,
			ItemId:         marketPlace.ItemId,
			MetaData:       string(metadata),
			ActivityType:   int64(activityType),
			ActivityUserId: marketPlace.SellerId,
			CollectionId:   marketPlace.Item.CollectionId,
		},
	}

	if err = i.UserStore.InsertActivities(activities...); err != nil {
		return err
	}

	marketPlace.Status = 2
	if err = i.Store.UpdateMarketPlaceStatus(marketPlace); err != nil {
		return err
	}
	if err = i.Store.UpdateItemUserId(marketPlace.ItemId, user.Id, 0, common.ITEM_STATUS_ON_INVENTORY); err != nil {
		return err
	}
	i.BuildUserDailyEvent(user.Id, common.DAILY_EVENT_TYPE_BUY_NFT)
	return nil
}

func (i MarketPlaceController) ChangeItemStatus(req marketplacemodel.ChangeMPStatusRequest) (*marketplacemodel.MarketPlace, error) {
	marketPlace, err := i.Store.GetMarketPlaceByItemId(&marketplacemodel.MarketPlaceFilter{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if marketPlace.Id == 0 {
		return nil, common.ErrNotFoundMarketPlaceId
	}
	marketPlace.Status = req.Status

	err = i.Store.UpdateMarketPlaceStatus(marketPlace)
	if err != nil {
		return nil, err
	}

	return marketPlace, nil
}

func (i MarketPlaceController) VoteItem(req marketplacemodel.VoteItemRequest) (*marketplacemodel.Vote, error) {
	// check user is staking
	isStaking, err := i.Store.IsStakingByUser(&marketplacemodel.Staking{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	if !isStaking {
		return nil, common.ErrUserNotStaking
	}

	existMarketPlace, err := i.Store.IsExistMarketPlace(&marketplacemodel.MarketPlace{Id: req.MarketPlaceId})
	if err != nil {
		return nil, err
	}

	if !existMarketPlace {
		return nil, common.ErrNotFoundMarketPlaceId
	}

	// vote req
	voteReq := &marketplacemodel.Vote{
		TransactionId:   req.TransactionId,
		ContractAddress: req.ContractAddress,
		MarketPlaceId:   req.MarketPlaceId,
		UserId:          req.UserId,
	}
	err = i.Store.InsertVoteItem(voteReq)
	if err != nil {
		return nil, err
	}
	return voteReq, err
}

func (i MarketPlaceController) LikeItem(rq marketplacemodel.LikeItemRequest) error {
	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	if err != nil {
		return err
	}

	if item.Id == 0 {
		return errors.New("Invalid item")
	}

	likeRq := &marketplacemodel.Like{
		ItemId:    rq.Id,
		UserId:    rq.UserId,
		IsDeleted: common.ACTIVE,
	}
	err = i.Store.InsertLikeItem(likeRq)
	if err != nil {
		return err
	}

	user, err := i.UserStore.GetUserByCondition(usermodel.UserFilter{Id: rq.UserId})
	if err != nil {
		return err
	}

	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      item.Title,
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"item_image":     item.Image,
		"wallet_address": user.WalletAddress,
	})

	exist, err := i.UserStore.IsExistNotification(&usermodel.Notification{
		UserId:             item.UserId,
		NotificationType:   common.LIKED,
		NotificationObject: base58.Encode(notiObj),
	})

	if !exist && item.UserId != user.Id {
		notification := &usermodel.Notification{
			UserId:             item.UserId,
			LastTimeRead:       0,
			NotificationType:   common.LIKED,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		err = i.UserStore.InsertNotification(notification)
		if err != nil {
			return nil
		}
	}

	err = i.UserStore.InsertActivities(&usermodel.Activities{
		UserId:       user.Id,
		ItemId:       item.Id,
		MetaData:     "",
		ActivityType: common.LIKE_ITEM,
	})
	if err != nil {
		return err
	}
	i.BuildUserDailyEvent(rq.UserId, common.DAILY_EVENT_TYPE_LIKE)
	return nil
}

func (i MarketPlaceController) DisLikeItem(rq marketplacemodel.LikeItemRequest) error {
	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	if err != nil {
		return err
	}

	if item.Id == 0 {
		return errors.New("Invalid item")
	}

	likeRq := &marketplacemodel.Like{
		ItemId:    rq.Id,
		UserId:    rq.UserId,
		IsDeleted: common.ACTIVE,
	}
	err = i.Store.DisLikeItem(likeRq)
	if err != nil {
		return err
	}

	err = i.UserStore.InsertActivities(&usermodel.Activities{
		UserId:       rq.UserId,
		ItemId:       item.Id,
		MetaData:     "",
		ActivityType: common.UNLIKE_ITEM,
	})
	if err != nil {
		return err
	}

	return nil
}

func (i MarketPlaceController) StakingItem(rq *marketplacemodel.StakingRequest) (*marketplacemodel.Staking, error) {
	user, err := i.UserStore.GetUserByCondition(usermodel.UserFilter{WalletAddress: rq.WalletAddress})
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, common.ErrNotFoundUser
	}

	req := &marketplacemodel.Staking{
		TransactionId:   rq.TransactionId,
		ContractAddress: rq.ContractAddress,
		UserId:          user.Id,
		Amount:          rq.Amount,
		InterestRate:    rq.InterestRate,
		TimeDeposit:     rq.TimeDeposit,
		Status:          common.STAKING,
	}
	result, err := i.Store.InsertStaking(req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i MarketPlaceController) UpdateStatusStaking(rq *marketplacemodel.UpdateStakingStatusRequest) (*marketplacemodel.Staking, error) {
	result, err := i.Store.FindStaking(&marketplacemodel.Staking{Id: rq.Id})
	if err != nil {
		return nil, err
	}

	result.Status = rq.Status
	err = i.Store.UpdateStaking(result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (i MarketPlaceController) CancelSellItem(rq *marketplacemodel.CancelMarketplace) error {
	err := i.Store.CancelSellItem(rq)
	if err != nil {
		return err
	}

	marketFilter := marketplacemodel.MarketPlaceFilter{
		IdOnMarket: rq.IdOnMarket,
		Status:     common.ITEM_STATUS_ON_MARKET,
		Unscoped:   true,
	}
	if len(rq.ContractAddress) > 0 {
		marketFilter.ContractAddress = rq.ContractAddress
	}
	if rq.SellType > 0 {
		marketFilter.SellType = rq.SellType
	}
	market, errItem := i.Store.GetMarketPlaceByIdOnMarket(&marketFilter)
	if errItem != nil {
		return errItem
	}

	user, err := i.UserStore.GetUserDetailById(market.SellerId)
	if err != nil {
		return err
	}

	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: market.ItemId})
	if err != nil {
		return err
	}

	// notification
	listFollowingUser, err := i.UserStore.ListFollowingUser(&usermodel.UserFollow{FollowerId: user.Id})
	if err != nil {
		return err
	}
	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      item.Title,
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"wallet_address": user.WalletAddress,
	})

	listNoti := make([]*usermodel.Notification, 0, len(listFollowingUser))
	for _, v := range listFollowingUser {
		notification := &usermodel.Notification{
			UserId:             v.UserId,
			LastTimeRead:       0,
			NotificationType:   common.NOTIFICATION_UNLISTED,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		listNoti = append(listNoti, notification)
	}
	activityType := common.UNLISTED_ITEM
	notifyType := common.CANCEL
	if rq.SellType == common.MARKET_PLACE_SELL_TYPE_AUCTION {
		activityType = common.AUCTION_CANCEL
		notifyType = common.NOTIFICATION_AUCTION_CANCEL

	}
	listNoti = append(listNoti, &usermodel.Notification{
		UserId:             user.Id,
		LastTimeRead:       0,
		NotificationType:   int64(notifyType),
		Status:             common.UNREAD,
		NotificationObject: base58.Encode(notiObj),
	})

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return err
		}
	}

	err = i.UserStore.InsertActivities(&usermodel.Activities{
		UserId:       user.Id,
		ItemId:       item.Id,
		ActivityType: int64(activityType),
		CollectionId: item.CollectionId,
	})
	if err != nil {
		return err
	}

	return nil
}

func (i MarketPlaceController) CollectionDetail(collectionId int64, userId int64) (marketplacemodel.CollectionDetail, error) {
	collection, err := i.Store.DetailCollection(collectionId)
	if err != nil {
		return marketplacemodel.CollectionDetail{}, err
	}

	user, err := i.UserStore.GetUserDetailById(collection.UserId)
	if err != nil {
		return marketplacemodel.CollectionDetail{}, err
	}

	likes, err := i.Store.ListCollectionLike(marketplacemodel.CollectionLikeFilter{
		CollectionId: collectionId,
		IsDeleted:    common.ACTIVE,
		Cols:         []string{"id", "collection_id", "user_id"},
	})
	if err != nil {
		return marketplacemodel.CollectionDetail{}, err
	}

	isLike := false
	userLikesId := make([]int64, 0, len(likes))
	for _, v := range likes {
		if userId == v.UserId {
			isLike = true
		}
		userLikesId = append(userLikesId, v.UserId)

	}

	userLikes, err := i.UserStore.GetListUserById(userLikesId)
	if err != nil {
		return marketplacemodel.CollectionDetail{}, err
	}

	items, _ := i.Store.GetItemWithCondition(marketplacemodel.ItemFilter{CollectionId: collectionId, IsNotGetMulti: 1, IsApproval: 1, IsHide: -1, IsTrending: -1, Stts: []int64{common.ITEM_ON_SALE, common.ITEM_NEW, common.ITEM_READY_ON_SELL}}, nil)
	collection.Logo = utils.RebuildUrlPath(collection.Logo)
	collection.Banner = utils.RebuildUrlPath(collection.Banner)
	collection.FeaturedImage = utils.RebuildUrlPath(collection.FeaturedImage)
	result := collection.ConvertToDetail()
	result.FloorPrice = "0"
	result.TotalItem = len(items)
	result.Items = items
	listOwner := make(map[int64]int64)
	//get list items
	for _, item := range items {
		listOwner[item.UserId] = item.UserId
	}

	listMarket, _ := i.Store.ListMarketPlace(marketplacemodel.MarketPlaceFilter{CollectionId: collectionId, Status: common.MARKET_PLACE_STATUS_SELLING}, nil)
	var totalPrice int64
	var floorPrice int64
	floorPrice = 0
	mod := big.NewInt(1000000000000000000)
	for _, market := range listMarket {
		distance, _ := new(big.Int).SetString(market.Price, 10)
		price := new(big.Int).Div(distance, mod).Int64()
		if floorPrice == 0 {
			floorPrice = price
		}
		if floorPrice > price {
			floorPrice = price
		}
	}
	listMarketSold, _ := i.Store.ListMarketPlace(marketplacemodel.MarketPlaceFilter{CollectionId: collectionId, Status: common.MARKET_PLACE_STATUS_SOLDED}, nil)
	for _, market := range listMarketSold {
		distance, _ := new(big.Int).SetString(market.Price, 10)
		price := new(big.Int).Div(distance, mod).Int64()
		totalPrice += price
	}
	result.FloorPrice = fmt.Sprint(floorPrice)
	result.TotalVolume = fmt.Sprint(totalPrice)
	result.TotalOwner = len(listOwner)
	result.Seller = &user
	result.TotalLike = int64(len(likes))
	result.IsLiked = isLike
	result.UserLikes = userLikes

	//update view
	_ = i.Store.UpdateViewCollection(collection.Id)

	return result, nil
}

func (i MarketPlaceController) LikeCollection(rq marketplacemodel.LikeCollectionRequest) error {

	_, err := i.Store.FindCollection(&marketplacemodel.Collection{Id: rq.Id})
	if err != nil {
		return err
	}

	likeRq := &marketplacemodel.CollectionLike{
		CollectionId: rq.Id,
		UserId:       rq.UserId,
		IsDeleted:    common.ACTIVE,
	}
	err = i.Store.InsertLikeCollection(likeRq)
	if err != nil {
		return err
	}

	//user, err := i.UserStore.GetUserByCondition(usermodel.UserFilter{Id: rq.UserId})
	//if err != nil {
	//	return err
	//}
	//
	//item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	//if err != nil {
	//	return err
	//}

	//notiObj, _ := json.Marshal(map[string]string{
	//	"item_name": item.Title,
	//	"user_name": user.DisplayName,
	//})
	//
	//exist, err := i.UserStore.IsExistNotification(&usermodel.Notification{
	//	UserId:             rq.UserId,
	//	NotificationType:   common.NOTIFICATION_LIKE,
	//	NotificationObject: base58.Encode(notiObj),
	//})
	//
	//if !exist && item.UserId != user.Id {
	//	notification := &usermodel.Notification{
	//		UserId:             rq.UserId,
	//		LastTimeRead:       0,
	//		NotificationType:   common.NOTIFICATION_LIKE,
	//		Status:             common.UNREAD,
	//		NotificationObject: base58.Encode(notiObj),
	//	}
	//	err = i.UserStore.InsertNotification(notification)
	//	if err != nil {
	//		return nil
	//	}
	//}
	i.BuildUserDailyEvent(rq.UserId, common.DAILY_EVENT_TYPE_LIKE)
	return nil
}

func (i MarketPlaceController) DisLikeCollection(rq marketplacemodel.LikeCollectionRequest) error {
	_, err := i.Store.FindCollection(&marketplacemodel.Collection{Id: rq.Id})
	if err != nil {
		return err
	}

	likeRq := &marketplacemodel.CollectionLike{
		CollectionId: rq.Id,
		UserId:       rq.UserId,
		IsDeleted:    common.ACTIVE,
	}
	err = i.Store.DisLikeCollection(likeRq)
	if err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) CollectionItemProperties(collectionId int64) ([]marketplacemodel.PropertiesResponse, error) {
	items, err := i.Store.CollectionItemProperties(collectionId)
	if err != nil {
		return nil, err
	}

	tmpResult := make(map[string]map[string]int64)

	result := make([]marketplacemodel.PropertiesResponse, 0, 100)
	properties := make([]marketplacemodel.Property, 0, 1)

	for _, v := range items {
		err = json.Unmarshal([]byte(v.Properties), &properties)
		if err != nil {
			continue
		}

		existName := make(map[string]bool)
		existValue := make(map[string]bool)
		for _, k := range properties {
			if len(k.Name) == 0 || k.Value == nil {
				continue
			}
			if tmpResult[k.Name] == nil {
				tmpResult[k.Name] = make(map[string]int64)
			}
			value := k.Value.(string)
			if !existName[k.Name] && !existValue[value] {
				tmpResult[k.Name][value] += 1
				existName[k.Name] = true
				existValue[k.Name] = true
			}

		}
	}

	for k, v := range tmpResult {
		element := make([]marketplacemodel.PropertiesValueResponse, 0, len(v))
		for value, count := range v {
			element = append(element, marketplacemodel.PropertiesValueResponse{
				Name:     value,
				TotalNft: count,
			})
		}
		result = append(result, marketplacemodel.PropertiesResponse{
			Name:    k,
			Total:   int64(len(v)),
			Element: element,
		})
	}

	return result, nil
}

func (i MarketPlaceController) ItemDetailBid(itemId int64) ([]marketplacemodel.BidItemDetailResponse, error) {
	bids, err := i.Store.GetAuctionBid(marketplacemodel.GetAuctionFilter{ItemId: itemId}, nil)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0, len(bids))
	for _, v := range bids {
		userIds = append(userIds, v.UserId)
	}

	users, err := i.UserStore.GetListUserById(userIds)
	if err != nil {
		return nil, err
	}

	mapUsers := make(map[int64]usermodel.UserDetail)
	for _, v := range users {
		mapUsers[v.Id] = v
	}

	result := make([]marketplacemodel.BidItemDetailResponse, 0, len(bids))
	for _, v := range bids {
		price, _ := new(big.Float).SetString(v.Price)
		priceFloat, _ := price.Float64()
		result = append(result, marketplacemodel.BidItemDetailResponse{
			Id:            v.Id,
			MarketPlaceId: v.MarketPlaceId,
			UserId:        v.UserId,
			Price:         v.Price,
			UserInfo:      mapUsers[v.UserId],
			USDTPrice:     utils.GetConvertedNow(priceFloat),
		})
	}

	return result, nil
}

func (u MarketPlaceController) BuildUserDailyEvent(userId int64, dailyEventType int) {
	var req usermodel.DailyEventFilter
	//get user
	user, errU := u.UserStore.GetUserDetailById(req.UserId)
	if errU != nil {
		return
	}
	req.EventType = dailyEventType
	dailyEvent, err := u.UserStore.SystemDailyEventDetail(req)
	if err != nil {
		return
	}
	timeCurrent := fmt.Sprint(time.Now().Year()) + "-" + fmt.Sprint(int(time.Now().Month())) + "-" + fmt.Sprint(time.Now().Day())
	timeCurrentStart, _ := time.Parse(common.ISO8601TIME, timeCurrent+" 00:00:00")
	timeCurrentEnd, _ := time.Parse(common.ISO8601TIME, timeCurrent+" 23:59:59")
	req.UserId = userId
	req.EventType = 0
	req.DailyEventId = dailyEvent.Id
	req.CurrentTimeStart = fmt.Sprint(timeCurrentStart.Unix())
	req.CurrentTimeEnd = fmt.Sprint(timeCurrentEnd.Unix())
	total, _ := u.UserStore.SystemUserDailyEventCount(req)

	var userDailyEvent usermodel.UserDailyEvent
	userDailyEvent.UserId = userId
	userDailyEvent.RewardOp = dailyEvent.RewardOp
	userDailyEvent.DailyEventId = dailyEvent.Id
	userDailyEvent.ActivityDate = fmt.Sprint(time.Now().Unix())

	if int(total) >= dailyEvent.NumberOfTask {
		return
	}
	u.UserStore.SystemUserDailyEventCreate(userDailyEvent)
	//update user
	u.UserStore.UpdateUserOp(userId, user.Op+dailyEvent.RewardOp, user.OpClaimed+dailyEvent.RewardOp, "")
}

func (i MarketPlaceController) ListReports(paging *common.Paging) ([]*marketplacemodel.ReportTableResponse, error) {
	data, err := i.Store.ListReport(paging)
	if err != nil {
		return nil, err
	}
	var details []*marketplacemodel.ReportTableResponse
	for _, reports := range data {
		detail := reports.ConvertToDetail()

		infoItem, errItem := i.Store.GetItemByID(marketplacemodel.ItemFilter{
			Id: reports.ItemId,
		})
		if errItem != nil {
			return nil, err
		}

		infoUser, errUser := i.UserStore.GetUserDetailById(reports.UserId)
		if errUser != nil {
			return nil, errUser
		}

		detail.UserDetail = infoUser
		detail.ItemDetail = *infoItem
		details = append(details, &detail)
	}
	return details, nil
}

func (i MarketPlaceController) ChangeReport(rq *marketplacemodel.ChangeReportRequest) error {
	if err := i.Store.ChangeReport(rq); err != nil {
		return err
	}
	return nil
}
func (i MarketPlaceController) RequestReport(rq *marketplacemodel.RequestReport) error {
	create := &marketplacemodel.ReportTableDetail{
		UserId:        rq.UserId,
		ItemId:        rq.ItemId,
		ReportReason:  rq.ReportReason,
		ReportMessage: rq.ReportMessage,
	}
	if err := i.Store.RequestReport(create); err != nil {
		return err
	}
	return nil
}
