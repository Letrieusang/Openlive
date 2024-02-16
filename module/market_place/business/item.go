package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"encoding/json"
	"errors"
	"math/big"
	"sort"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (i MarketPlaceController) ListItemWithConditions(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]*marketplacemodel.ItemResponse, error) {
	// check is admin for show item hide
	if rq.PresentUserId == rq.UserId {
		rq.IsHide = -1
	}
	if rq.CurrentUserId > 0 {
		rq.PresentUserId = rq.CurrentUserId
	}
	if rq.PresentUserId > 0 {
		user, err := i.UserStore.GetUserDetailById(rq.PresentUserId)
		if err != nil {
			return nil, err
		}
		if user.Type >= 6 {
			rq.IsHide = -1
		}
	}

	//rq.Stts = append(rq.Stts, common.ITEM_NEW)
	guestId := rq.UserId
	if rq.IsOwned == 1 {
		rq.Stts = append(rq.Stts, common.ITEM_NEW, common.ITEM_ON_SALE, common.ITEM_READY_ON_SELL)
	}

	var tmpMarketFilter marketplacemodel.MarketPlaceFilter
	tmpMarketFilter.UserId = rq.UserId
	tmpMarketFilter.Status = common.MARKET_PLACE_STATUS_SELLING

	listMarket, err := i.Store.ListMarketPlace(tmpMarketFilter, nil)
	if err != nil {
		return nil, err
	}

	var listMarketItemIds []int64
	var listMarketItemMainIds = make(map[int64]int64)
	mapItem := make(map[int64]*marketplacemodel.MarketPlace)
	for _, v := range listMarket {
		mapItem[v.ItemId] = v
		if _, ok := listMarketItemMainIds[v.Item.MainItemId]; ok {
			listMarketItemIds = append(listMarketItemIds, v.ItemId)
			continue
		}
		if v.Item.MainItemId == 0 {
			listMarketItemMainIds[v.ItemId] = v.ItemId
		} else {
			listMarketItemMainIds[v.Item.MainItemId] = v.Item.MainItemId
		}
	}
	rq.OrNotListItemIds = listMarketItemIds
	listItems, err := i.Store.GetItemWithConditionNew(rq, paging)
	if err != nil {
		return nil, err
	}
	if len(listItems) < 1 {
		return nil, nil
	}
	var listItemsId []int64
	var ListOwnerId []int64
	for _, v := range listItems {
		listItemsId = append(listItemsId, v.Id)
		ListOwnerId = append(ListOwnerId, v.UserId)
		ListOwnerId = append(ListOwnerId, v.ItemCreatorId)
	}

	// like
	mapLiked := make(map[int64]bool, 100)
	mapLikeItem := make(map[int64][]usermodel.UserDetail, 100)
	if rq.UserId > 0 {
		// list liked item
		listLiked, err := i.Store.ListLikedItem(listItemsId, rq.UserId)
		if err != nil {
			return nil, err
		}
		for _, v := range listLiked {
			mapLiked[v.ItemId] = true
		}

	}

	likes, err := i.Store.ListLikes(listItemsId)
	if err != nil {
		return nil, err
	}

	listUserId := make([]int64, 0, len(likes))
	for _, v := range likes {
		listUserId = append(listUserId, v.UserId)
	}

	usersLike, err := i.UserStore.GetListUserById(listUserId)
	if err != nil {
		return nil, err
	}

	owners, err := i.UserStore.GetListUserById(ListOwnerId)
	if err != nil {
		return nil, err
	}
	mapOwner := make(map[int64]usermodel.UserDetail, len(owners))
	for _, v := range owners {
		mapOwner[v.Id] = v
	}

	for _, v := range likes {
		if rq.CurrentUserId == v.UserId {
			mapLiked[v.ItemId] = true
		}
		for _, j := range usersLike {
			if v.UserId == j.Id && len(mapLikeItem[v.ItemId]) <= common.LIMIT_USER_LIKE {
				mapLikeItem[v.ItemId] = append(mapLikeItem[v.ItemId], j)
				continue
			}
		}
	}

	result := make([]*marketplacemodel.ItemResponse, 0, len(listItems))
	for _, v := range listItems {
		var tmpResult marketplacemodel.ItemResponse
		v.Image = utils.RebuildUrlPath(v.Image)
		if v.IsLocking == 1 && v.UserId != guestId {
			v.Code = ""
		}
		tmpResult.ItemDetail = v
		if mapItem[v.Id] != nil {
			tmpResult.Price = mapItem[v.Id].Price
			tmpResult.Seller = mapItem[v.Id].Seller
			tmpResult.IdOnMarket = mapItem[v.Id].IdOnMarket
			tmpResult.MarketContractAddress = mapItem[v.Id].ContractAddress
			tmpResult.SellType = mapItem[v.Id].SellType
			tmpResult.AuctionEndTime = mapItem[v.Id].AuctionEndTime
			tmpResult.AuctionStartTime = mapItem[v.Id].AuctionStartTime
		}

		userLikes := make([]usermodel.UserDetail, 0, 100)
		if mapLikeItem[v.Id] != nil {
			userLikes = mapLikeItem[v.Id]
		}
		tmpResult.UserLikes = userLikes
		tmpResult.IsLike = mapLiked[v.Id]
		tmpResult.IsLike = mapLiked[v.Id]
		tmpResult.Owner = mapOwner[v.UserId]
		tmpResult.ItemCreator = mapOwner[v.ItemCreatorId]

		result = append(result, &tmpResult)
	}

	return result, nil
}

func (i MarketPlaceController) DetailItem(filter marketplacemodel.ItemFilter) (*marketplacemodel.ItemDetailResponse, error) {
	info, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{
		Nft:             filter.Nft,
		ContractAddress: filter.ContractAddress,
		IsTrending:      filter.IsTrending,
	})
	if err != nil {
		return nil, err
	}
	if info.Id == 0 {
		return nil, common.ErrNotFoundItem
	}

	//get number of copy available
	var availableCount int64 = 0
	if info.NumOfCopies > 1 {
		var countItemCopyFilter marketplacemodel.CountItemFilter
		countItemCopyFilter.Status = common.ITEM_NEW
		if info.MainItemId > 0 {
			countItemCopyFilter.MainItemId = info.MainItemId
		} else {
			countItemCopyFilter.MainItemId = info.Id
		}
		availableCount, _ = i.Store.CountAvailableItemCopied(countItemCopyFilter)
		if info.MainItemId == 0 {
			availableCount++
		}
	}

	marketPlace, err := i.Store.GetMarketPlaceByItemId(&marketplacemodel.MarketPlaceFilter{ItemId: info.Id, Status: common.MARKET_PLACE_STATUS_SELLING})
	if err != nil {
		return nil, err
	}
	//get created user
	createdUser, errCreatedUser := i.UserStore.GetCreatedUserByItemId(info.Id)
	if errCreatedUser != nil {
		return nil, errCreatedUser
	}

	// get owner
	owner, errOwner := i.UserStore.GetUserDetailById(info.UserId)
	if errOwner != nil {
		return nil, errOwner
	}

	totalLike, err := i.Store.CountLikeById(info.Id)
	if err != nil {
		return nil, err
	}

	infoDetail := info.ConvertToDetail()
	// show code for owner
	if filter.UserId == info.UserId {
		infoDetail.Code = info.Code
	}

	infoDetail.IsLike, _ = i.Store.IsLikeItem(filter.UserId, info.Id)
	infoDetail.TotalLike = totalLike

	infoDetail.Collection = &marketplacemodel.Collection{}
	if info.CollectionId > 0 {
		collection, errColl := i.Store.DetailCollection(info.CollectionId)
		if errColl != nil {
			return nil, errColl
		}
		infoDetail.Collection = collection
	}

	//update view
	_ = i.Store.UpdateViewItem(info.Id)

	infoDetail.Image = utils.RebuildUrlPath(infoDetail.Image)

	var filterAuctionBid marketplacemodel.GetAuctionFilter
	filterAuctionBid.ItemId = infoDetail.Id
	filterAuctionBid.OrderBy = "auction_bids.id desc"
	filterAuctionBid.IsGetUser = 1
	listAuctionBid, errBid := i.Store.GetAuctionBid(filterAuctionBid, nil)
	if errBid != nil {
		return nil, errBid
	}
	for index, _ := range listAuctionBid {
		listAuctionBid[index].UserInfo.Avatar = utils.RebuildUrlPath(listAuctionBid[index].UserInfo.Avatar)
		listAuctionBid[index].UserInfo.Cover = utils.RebuildUrlPath(listAuctionBid[index].UserInfo.Cover)
	}

	result := &marketplacemodel.ItemDetailResponse{
		ItemDetail:    &infoDetail,
		AvailableCopy: availableCount,
		ItemCreator:   createdUser,
		Price:         marketPlace.Price,
		Owner:         owner,
		IsOldEvent:    -1,
		AuctionBid:    listAuctionBid,
	}
	if info.IsJoinEvent == 1 {
		//check event
		result.IsOldEvent = 1
		if exist, _ := i.Store.CheckOldItemEvent(info.Id); !exist {
			result.IsOldEvent = 0
		}
	}

	return result, nil
}

func (i MarketPlaceController) ItemCreator(rq *marketplacemodel.CreateItemRequest, userId int64, options map[string]interface{}) (marketplacemodel.Item, error) {
	var result marketplacemodel.Item
	cover := utils.DESTINATION + utils.GetFileName(rq.Cover)
	var image string
	if rq.IsAirDrop != 1 {
		image = utils.DESTINATION + utils.GetFileName(rq.Image)

		if len(rq.Image) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(rq.Image, cast.ToString(options["host"]), ""), image); err != nil {
				return result, err
			}
		}
		if image != utils.DESTINATION {
			image = cast.ToString(options["host"]) + image
		}
	} else {
		image = rq.Image
	}

	if len(rq.Cover) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Image, cast.ToString(options["host"]), ""), cover); err != nil {
			return result, err
		}
	}

	user, errUser := i.UserStore.GetUserDetailById(userId)
	if errUser != nil {
		return result, errUser
	}

	isExistCategory, errItem := i.Store.IsExistCategory(&marketplacemodel.Category{Id: rq.CategoryId})
	if errItem != nil {
		return result, errItem
	}
	if !isExistCategory {
		return result, common.ErrNotFoundCate
	}

	if rq.IsLocking && len(rq.Code) == 0 {
		return result, common.ErrNFTCodeRequired
	}

	item := marketplacemodel.Item{
		UserId:          userId,
		Nft:             rq.Nft,
		TransactionId:   rq.TransactionId,
		ContractAddress: rq.ContractAddress,
		Status:          common.ITEM_CREATE_WAITING_SYNC,
		Title:           rq.Title,
		Image:           image,
		Description:     rq.Description,
		CategoryId:      rq.CategoryId,
		IsApproval:      common.UN_APPROVE,
		NumOfCopies:     rq.NumOfCopies,
		IsLocking:       cast.ToInt64(rq.IsLocking),
	}

	if rq.IsLocking {
		item.Code = rq.Code
	}

	if len(rq.Property) > 0 {
		for _, v := range rq.Property {
			if len(v.Name) == 0 || cast.ToString(v.Value) == "" {
				return result, common.ErrNFTPropertiesRequired
			}
		}
		binData, _ := json.Marshal(rq.Property)
		item.Properties = string(binData)
	}

	if cover != utils.DESTINATION {
		item.Cover = cast.ToString(options["host"]) + cover
	}

	if rq.CollectionId > 0 {
		isExistCollection, errCol := i.Store.IsExistCollection(&marketplacemodel.Collection{Id: rq.CollectionId})
		if errCol != nil {
			return result, errCol
		}
		if !isExistCollection {
			return result, common.ErrNotFoundCol
		}
		item.CollectionId = rq.CollectionId
	}
	result = item
	if item.NumOfCopies > 1 {
		var listItems []marketplacemodel.Item
		for i := 0; i < int(item.NumOfCopies)-1; i++ {
			listItems = append(listItems, item)
		}
		itemId, err := i.Store.InsertItems(item, listItems)
		if err != nil {
			return result, err
		}
		item.Id = itemId
		result.Id = itemId
	} else {
		itemId, err := i.Store.InsertItem(item)
		if err != nil {
			return result, err
		}
		item.Id = itemId
		result.Id = itemId
	}

	// notification
	listFollowingUser, err := i.UserStore.ListFollowingUser(&usermodel.UserFollow{FollowerId: user.Id})
	if err != nil {
		return result, err
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
			NotificationType:   common.MINTED,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		listNoti = append(listNoti, notification)
	}

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return result, err
		}
	}

	err = i.UserStore.InsertActivities(&usermodel.Activities{
		UserId:       item.UserId,
		ItemId:       item.Id,
		MetaData:     "",
		ActivityType: common.MINTED_ITEM,
		CollectionId: item.CollectionId,
	})
	if err != nil {
		return result, err
	}

	return result, nil

}

func (i MarketPlaceController) CreateCollections(rq *marketplacemodel.CreateCollectionRequest, options map[string]interface{}) (*marketplacemodel.Collection, error) {
	exist, err := i.Store.CheckNameCollection(&marketplacemodel.Collection{Name: rq.Name})
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("collection name has exist")
	}

	banner := utils.DESTINATION + utils.GetFileName(rq.Banner)
	featured := utils.DESTINATION + utils.GetFileName(rq.FeaturedImage)
	logo := utils.DESTINATION + utils.GetFileName(rq.Logo)

	if len(rq.Banner) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Banner, cast.ToString(options["host"]), ""), banner); err != nil {
			return nil, err
		}
	}

	if len(rq.FeaturedImage) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.FeaturedImage, cast.ToString(options["host"]), ""), featured); err != nil {
			return nil, err
		}
	}

	if len(rq.Logo) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Logo, cast.ToString(options["host"]), ""), logo); err != nil {
			return nil, err
		}
	}

	if rq.IsHold && (rq.TimePublicSale <= rq.TimeStartHold || rq.TimePublicSale <= rq.TimeEndHold) {
		return nil, errors.New("Incorrect time")
	}

	listItemId := make([]int64, 0, 100)
	items, err := i.Store.GetItemWithCondition(marketplacemodel.ItemFilter{Ids: rq.ItemIds, UserId: rq.UserId}, nil)
	if err != nil {
		return nil, err
	}

	if len(rq.ItemIds) == 0 {
		items = make([]marketplacemodel.Item, 0, 1)
	}

	for _, v := range items {
		if v.CollectionId != 0 {
			return nil, errors.New("please select nft hasn't collection")
		}
		listItemId = append(listItemId, v.Id)
	}

	if rq.IsHold {
		var minOpv float64
		var err error
		if minOpv, err = strconv.ParseFloat(rq.MinOpvHolding, 64); err != nil {
			return nil, errors.New("Nin opv need to hold incorrect")
		}
		if minOpv < 1 {
			return nil, errors.New("Nin opv need to hold incorrect")
		}

		if len(items) < 1 {
			return nil, errors.New("required min 1 items for create collection")
		}
	}

	linkBin, _ := json.Marshal(rq.Links)
	collection := &marketplacemodel.Collection{
		Name:            rq.Name,
		Description:     rq.Description,
		UserId:          rq.UserId,
		IsHold:          rq.IsHold,
		TimeEndHold:     rq.TimeEndHold,
		TimeStartHold:   rq.TimeStartHold,
		TimePublicSale:  rq.TimePublicSale,
		Status:          common.COLLECTION_NEW,
		Banner:          cast.ToString(options["host"]) + banner,
		FeaturedImage:   cast.ToString(options["host"]) + featured,
		Logo:            cast.ToString(options["host"]) + logo,
		ContractAddress: rq.ContractAddress,
		IsHot:           0,
		Links:           string(linkBin),
	}
	if rq.IsHold {
		collection.MinOpvHolding = rq.MinOpvHolding
	}

	//get user info
	userDetail, errUser := i.UserStore.GetUserDetailById(rq.UserId)
	if errUser != nil {
		return nil, errors.New("Incorrect user")
	}
	if userDetail.Type >= common.VIP_USER {
		collection.IsHot = common.IS_HOT
	}
	err = i.Store.InsertCollection(collection)
	if err != nil {
		return nil, err
	}

	// batch item update
	if len(listItemId) > 0 {
		err = i.Store.BatchUpdateCollectionItem(listItemId, map[string]interface{}{
			"collection_id": collection.Id,
		})
		if err != nil {
			return nil, err
		}
	}

	// notification
	listFollowingUser, err := i.UserStore.ListFollowingUser(&usermodel.UserFollow{FollowerId: rq.UserId})
	if err != nil {
		return nil, err
	}
	notiObj, _ := json.Marshal(map[string]string{
		"collection_name": collection.Name,
		"user_name":       userDetail.DisplayName,
		"user_avatar":     userDetail.Avatar,
		"wallet_address":  userDetail.WalletAddress,
	})

	listNoti := make([]*usermodel.Notification, 0, len(listFollowingUser))
	for _, v := range listFollowingUser {
		notification := &usermodel.Notification{
			UserId:             v.UserId,
			LastTimeRead:       0,
			NotificationType:   common.CREATED_COLLECTION,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		listNoti = append(listNoti, notification)
	}

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return nil, err
		}
	}

	return collection, nil
}

func (i MarketPlaceController) CanCreateCollections(rq *marketplacemodel.CanCreateCollectionRequest) (bool, error) {
	collection, err := i.Store.GetCollectionByCondition(rq)
	if err != nil {
		return false, err
	}
	if collection.Id > 0 {
		return false, errors.New("collection name has exist")
	}

	return true, nil
}

func (i MarketPlaceController) ListCollections(filter marketplacemodel.CollectionFilter, paging *common.Paging) ([]*marketplacemodel.CollectionDetail, error) {
	data, err := i.Store.ListCollections(&filter, paging)
	if err != nil {
		return nil, err
	}
	var listUserId []int64
	var listCollectionId []int64
	for _, collection := range data {
		listUserId = append(listUserId, collection.UserId)
		listCollectionId = append(listCollectionId, collection.Id)
	}

	countLike, errCount := i.Store.CountCollectionLike(listCollectionId)
	if errCount != nil {
		return nil, errCount
	}
	mapCount := make(map[int64]int64, len(countLike))
	for _, v := range countLike {
		mapCount[v.CollectionId] = v.TotalLike
	}

	likes, errLike := i.Store.ListCollectionLike(marketplacemodel.CollectionLikeFilter{
		CollectionIds: listCollectionId,
		IsDeleted:     common.ACTIVE,
		Cols:          []string{"id", "collection_id", "user_id"},
	})
	if errLike != nil {
		return nil, errLike
	}

	mapLiked := make(map[int64]bool, len(likes))
	listUserLikeId := make([]int64, 0, len(likes))
	for _, v := range likes {
		if filter.UserIdLike == v.UserId {
			mapLiked[v.CollectionId] = true
		}
		listUserLikeId = append(listUserLikeId, v.UserId)
	}

	dataUserLike, errUserLike := i.UserStore.GetListUserById(listUserLikeId)
	if errUserLike != nil {
		return nil, errUserLike
	}
	mapUserLike := make(map[int64][]usermodel.UserDetail, len(dataUserLike))
	for _, v := range likes {
		for _, j := range dataUserLike {
			if v.UserId == j.Id && len(mapUserLike[v.CollectionId]) <= common.LIMIT_USER_LIKE {
				mapUserLike[v.CollectionId] = append(mapUserLike[v.CollectionId], j)
				continue
			}
		}
	}

	var dataItems []marketplacemodel.Item
	if filter.NotGetItem != 1 {
		var errItem error
		dataItems, errItem = i.Store.GetCollectionItems(listCollectionId)
		if errItem != nil {
			return nil, errItem
		}
	}
	dataUser, errUser := i.UserStore.GetListUserById(listUserId)
	if errUser != nil {
		return nil, errUser
	}

	var details []*marketplacemodel.CollectionDetail
	for _, collection := range data {
		collection.Banner = utils.RebuildUrlPath(collection.Banner)
		collection.FeaturedImage = utils.RebuildUrlPath(collection.FeaturedImage)
		collection.Logo = utils.RebuildUrlPath(collection.Logo)
		detail := collection.ConvertToDetail()
		if len(dataItems) > 0 {
			for _, item := range dataItems {
				if collection.Id == item.CollectionId {
					item.Image = utils.RebuildUrlPath(item.Image)
					detail.Items = append(detail.Items, item)
				}
			}
		}
		for _, user := range dataUser {
			if collection.UserId == user.Id {
				user.Cover = utils.RebuildUrlPath(user.Cover)
				user.Avatar = utils.RebuildUrlPath(user.Avatar)
				detail.Seller = &user
				detail.Owner = &user
				break
			}
		}

		detail.TotalLike = mapCount[detail.Id]
		detail.UserLikes = mapUserLike[detail.Id]
		detail.IsLiked = mapLiked[detail.Id]
		details = append(details, &detail)
	}

	return details, nil
}

func (i MarketPlaceController) ListTransaction(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]marketplacemodel.ShortInfoMarketPlace, error) {
	listTransactions, errTrans := i.Store.GetListTransaction(rq, paging)
	if errTrans != nil {
		return nil, errTrans
	}

	if listTransactions == nil {
		return nil, nil
	}

	var listUserId []int64
	var listItemId []int64
	for _, tr := range listTransactions {
		listUserId = append(listUserId, tr.BuyerId)
		listUserId = append(listUserId, tr.SellerId)
		listItemId = append(listItemId, tr.ItemId)
	}

	listUsers, errUser := i.UserStore.GetListUserById(listUserId)
	if errUser != nil {
		return nil, errUser
	}

	listItems, errItem := i.Store.GetListItemById(listItemId)
	if errItem != nil {
		return nil, errItem
	}

	for index, tr := range listTransactions {
		for _, u := range listUsers {
			u.Avatar = utils.RebuildUrlPath(u.Avatar)
			u.Cover = utils.RebuildUrlPath(u.Cover)
			if tr.BuyerId == u.Id {
				listTransactions[index].Buyer = u
			}
			if tr.SellerId == u.Id {
				listTransactions[index].Seller = u
			}
		}
		for _, it := range listItems {
			it.Image = utils.RebuildUrlPath(it.Image)
			it.Cover = utils.RebuildUrlPath(it.Cover)
			if tr.ItemId == it.Id {
				listTransactions[index].Item = it
			}
		}
	}

	return listTransactions, nil
}

func (i MarketPlaceController) UpdateItem(rq *marketplacemodel.UpdateItemRequest) (*marketplacemodel.Item, error) {
	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	if err != nil {
		return nil, err
	}

	if item.UserId != rq.UserId {
		return nil, common.ErrNotPermissionUpdateItem
	}

	if item.Status == common.ITEM_ON_SALE && rq.CollectionId > 0 && rq.CollectionId != item.CollectionId {
		return nil, common.ErrCantNotUpdateNft
	}

	if item.IsApproval == common.APPROVE && rq.Status == common.ITEM_PENDING {
		return nil, common.ErrCantUpdateItemStatus
	}

	if rq.CollectionId != 0 {
		collection, err := i.Store.FindCollection(&marketplacemodel.Collection{Id: rq.CollectionId})
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, err
		}

		if collection == nil {
			return nil, common.ErrNotFoundCol
		}

		if collection.UserId != rq.UserId {
			return nil, common.ErrNotPermissionUpdate
		}
	}

	if rq.Status == common.ITEM_PENDING && item.Status == common.ITEM_NEW {
		//case verify sell
		if setting, errSetting := i.Store.GetSetting(); errSetting == nil {
			if setting.IsVerifySellNft == common.IS_NOT_VERIFY_SELL_NFT {
				rq.Status = common.ITEM_READY_ON_SELL
				rq.IsApproved = common.APPROVE
			}
		} else {
			rq.IsApproved = item.IsApproval
		}
	}

	if err := i.Store.UpdateItem(&marketplacemodel.Item{
		CollectionId: rq.CollectionId,
		Status:       rq.Status,
		IsApproval:   rq.IsApproved,
	}, &marketplacemodel.Item{Id: rq.Id}); err != nil {
		return nil, err
	}

	item, err = i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	if err != nil {
		return nil, err
	}

	return item, nil

}

func (i MarketPlaceController) UpdateCollection(rq *marketplacemodel.UpdateCollectionRequest, options map[string]interface{}) (*marketplacemodel.Collection, error) {
	collection, err := i.Store.FindCollection(&marketplacemodel.Collection{Id: rq.Id, UserId: rq.UserId})
	if err != nil {
		return nil, err
	}

	if collection.Id == 0 {
		return nil, common.ErrNotFoundCol
	}

	if collection.UserId != rq.UserId {
		return nil, common.ErrNotPermissionUpdate
	}

	collectionUpdate := &marketplacemodel.Collection{
		Name:        rq.Name,
		Description: rq.Description,
	}

	if rq.Links != nil {
		linksBin, _ := json.Marshal(rq.Links)
		collectionUpdate.Links = string(linksBin)
	}

	if len(rq.MinOpvHolding) > 0 {
		collectionUpdate.MinOpvHolding = rq.MinOpvHolding
	}

	if len(rq.ItemIds) > 0 {
		listItemId := make([]int64, 0, 100)
		items, err := i.Store.GetItemWithCondition(marketplacemodel.ItemFilter{Ids: rq.ItemIds, UserId: rq.UserId}, nil)
		if err != nil {
			return nil, err
		}

		if len(rq.ItemIds) == 0 {
			items = make([]marketplacemodel.Item, 0, 1)
		}

		for _, v := range items {
			if v.CollectionId != 0 {
				return nil, errors.New("please select nft hasn't collection")
			}
			listItemId = append(listItemId, v.Id)
		}

		// batch item update
		if len(listItemId) > 0 {
			err = i.Store.BatchUpdateCollectionItem(listItemId, map[string]interface{}{
				"collection_id": collection.Id,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	if len(rq.Banner) > 0 {
		banner := utils.DESTINATION + utils.GetFileName(rq.Banner)
		if err := utils.MoveFile(strings.ReplaceAll(rq.Banner, cast.ToString(options["host"]), ""), banner); err != nil {
			return nil, err
		}
		collectionUpdate.Banner = cast.ToString(options["host"]) + banner
	}

	if len(rq.FeaturedImage) > 0 {
		featured := utils.DESTINATION + utils.GetFileName(rq.FeaturedImage)
		if err := utils.MoveFile(strings.ReplaceAll(rq.FeaturedImage, cast.ToString(options["host"]), ""), featured); err != nil {
			return nil, err
		}
		collectionUpdate.FeaturedImage = cast.ToString(options["host"]) + featured
	}

	if len(rq.Logo) > 0 {
		logo := utils.DESTINATION + utils.GetFileName(rq.Logo)
		if err := utils.MoveFile(strings.ReplaceAll(rq.Logo, cast.ToString(options["host"]), ""), logo); err != nil {
			return nil, err
		}
		collectionUpdate.Logo = cast.ToString(options["host"]) + logo
	}

	errUpdate := i.Store.UpdateCollection(collectionUpdate, &marketplacemodel.Collection{Id: collection.Id})
	if errUpdate != nil {
		return nil, errUpdate
	}

	collection, err = i.Store.FindCollection(&marketplacemodel.Collection{Id: rq.Id, UserId: rq.UserId})
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func (i MarketPlaceController) UpdateCollectionCms(rq *marketplacemodel.UpdateCollectionCmsRequest) error {
	var update = make(map[string]interface{})
	if rq.IsHot >= 0 {
		update["is_hot"] = rq.IsHot
	}
	if rq.IsTrending >= 0 {
		update["is_trending"] = rq.IsTrending
	}
	if rq.DisplayOrder >= 0 {
		update["display_order"] = rq.DisplayOrder
	}
	if err := i.Store.UpdateCollectionCms(rq.Id, update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) DeleteBrandImage(id int64) error {
	if err := i.Store.DeleteBrandImage(id); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) DeleteBrand(id int64) error {
	if err := i.Store.DeleteBrand(id); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) ListBrands(filter marketplacemodel.BrandFilter, paging *common.Paging) ([]*marketplacemodel.BrandDetail, error) {
	data, err := i.Store.ListBrands(&filter, paging)
	if err != nil {
		return nil, err
	}
	var ListBrandsImage []int64
	for _, brands := range data {
		ListBrandsImage = append(ListBrandsImage, brands.Id)
	}

	var dataBrands []marketplacemodel.BrandImage
	var errBrands error
	dataBrands, errBrands = i.Store.GetBrandImage(ListBrandsImage)
	if errBrands != nil {
		return nil, errBrands
	}

	var details []*marketplacemodel.BrandDetail
	for _, brands := range data {
		brands.BackgroundImage = utils.RebuildUrlPath(brands.BackgroundImage)
		detail := brands.ConvertToDetail()
		if len(dataBrands) > 0 {
			for _, BrandImage := range dataBrands {
				if brands.Id == BrandImage.BrandId {
					BrandImage.ItemBackgroundImage = utils.RebuildUrlPath(BrandImage.ItemBackgroundImage)
					BrandImage.ItemAvatarImage = utils.RebuildUrlPath(BrandImage.ItemAvatarImage)
					BrandImageDetail := BrandImage.ConvertToDetail()
					detail.Images = append(detail.Images, BrandImageDetail)
				}
			}
		}
		details = append(details, &detail)
	}

	return details, nil
}

func (i MarketPlaceController) UpdateBrandCms(rq *marketplacemodel.UpdateBrandCmsRequest, options map[string]interface{}) error {
	var update = make(map[string]interface{})
	if strings.Contains(rq.BackgroundImage, "/tmp/") {
		backgroundimage := utils.DESTINATION + utils.GetFileName(rq.BackgroundImage)
		if len(rq.BackgroundImage) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(rq.BackgroundImage, cast.ToString(options["host"]), ""), backgroundimage); err != nil {
				return err
			}
		}
		rq.BackgroundImage = cast.ToString(options["host"]) + backgroundimage
	}
	if rq.Ishot >= 0 {
		update["is_hot"] = rq.Ishot
	}
	if rq.DisplayOrder >= 0 {
		update["display_order"] = rq.DisplayOrder
	}
	update["brand_title"] = rq.BrandTitle
	update["description"] = rq.Description
	update["background_image"] = rq.BackgroundImage
	if err := i.Store.UpdateBrandCms(rq.Id, update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) AddBannerCms(rq *marketplacemodel.AddBannerCmsRequest) error {
	var update = make(map[string]interface{})
	if len(rq.Title) >= 0 {
		update["title"] = rq.Title
	}
	if len(rq.BlueTitle) >= 0 {
		update["blue_title"] = rq.BlueTitle
	}
	if len(rq.Line1) >= 0 {
		update["line1"] = rq.Line1
	}
	if len(rq.Line2) >= 0 {
		update["line2"] = rq.Line2
	}
	if len(rq.Line3) >= 0 {
		update["line3"] = rq.Line3
	}
	if len(rq.Line4) >= 0 {
		update["line4"] = rq.Line4
	}
	if err := i.Store.AddBannerCms(update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) ListBanner() ([]*marketplacemodel.BannerDetail, error) {
	data, err := i.Store.ListBannerCms()
	if err != nil {
		return nil, err
	}
	var details []*marketplacemodel.BannerDetail
	for _, banner := range data {
		detail := banner.ConvertToDetail()
		details = append(details, &detail)
	}
	return details, nil
}

func (i MarketPlaceController) UpdateItemStatusCMS(rq *marketplacemodel.UpdateItemStatusRequest) error {
	if err := i.Store.UpdateItemStatus(rq.Id, int(rq.Status)); err != nil {
		return err
	}
	return nil

}

func (i MarketPlaceController) UpdateItemCMS(rq *marketplacemodel.UpdateItemCmsRequest) error {
	var update = make(map[string]interface{})
	if rq.Status > 0 {
		update["status"] = rq.Status
	}
	if rq.IsTrending >= 0 {
		update["is_trending"] = rq.IsTrending
	}
	if rq.DisplayOrder >= 0 {
		update["display_order"] = rq.DisplayOrder
	}
	if err := i.Store.UpdateItemCms(rq.Id, update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) ListItemEvent(rq marketplacemodel.ItemEventFilter, paging *common.Paging) (*marketplacemodel.ListItemEventResponse, error) {
	//if rq.UserId > 0 {
	//	user, err := i.UserStore.GetUserDetailById(rq.UserId)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	if user.Type >= 6 {
	//		rq.IsHide = -1
	//	}
	//}

	if len(rq.SortBy) == 0 {
		rq.SortBy = "items.time_join_event DESC"
	}
	result, err := i.Store.GetListItemByEvent(rq, paging)
	if err != nil {
		return nil, err
	}
	listItemId := make([]int64, 0, len(result))
	for index, v := range result {
		result[index].Image = utils.RebuildUrlPath(result[index].Image)
		result[index].UserAvatar = utils.RebuildUrlPath(result[index].UserAvatar)
		listItemId = append(listItemId, v.Id)
	}
	response := marketplacemodel.ListItemEventResponse{
		Data: result,
	}

	if rq.IsShowAll == 1 {
		return &response, nil
	}
	//count user
	var metaData marketplacemodel.MetaItemEventResult
	totalUser, errCountUser := i.UserStore.CountUserItemEvent()
	if errCountUser != nil {
		return nil, errCountUser
	}

	//count nft
	if rq.CategoryId < 1 {
		metaData.TotalNft = paging.Total
	} else {
		totalNft, errCountNft := i.Store.CountListItemByEvent()
		if errCountNft != nil {
			return nil, errCountNft
		}
		metaData.TotalNft = totalNft
	}

	metaData.TotalUser = totalUser
	//get list categories
	var cateFilter *marketplacemodel.CategoryFilter
	listCategories, errListCate := i.Store.ListCategories(cateFilter)
	if errCountUser != nil {
		return nil, errListCate
	}
	var totalItemByCates marketplacemodel.CountItemEventByCategory
	for _, cate := range listCategories {
		totalItems, errByCate := i.Store.CountItemEventByCate(cate.Id)
		if errByCate != nil {
			return nil, errByCate
		}
		totalItemByCates.CategoryId = cate.Id
		totalItemByCates.CategoryName = cate.Name
		totalItemByCates.TotalNft = totalItems
		metaData.ListCategories = append(metaData.ListCategories, totalItemByCates)
	}
	response.MetaData = metaData

	mapLiked := make(map[int64]bool)
	mapLikeItem := make(map[int64][]usermodel.UserDetail)
	if rq.CurrentUserId > 0 {
		// list liked item
		listLiked, err := i.Store.ListLikedItem(listItemId, rq.CurrentUserId)
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
			j.Cover = utils.RebuildUrlPath(j.Cover)
			j.Avatar = utils.RebuildUrlPath(j.Avatar)
			if v.UserId == j.Id && len(mapLikeItem[v.ItemId]) <= common.LIMIT_USER_LIKE {
				mapLikeItem[v.ItemId] = append(mapLikeItem[v.ItemId], j)
				continue
			}
		}
	}

	for index, v := range response.Data {
		/*userLikes := make([]usermodel.UserDetail, 0, 100)
		if mapLikeItem[v.Id] != nil {
			userLikes = mapLikeItem[v.Id]
		}*/
		response.Data[index].IsLike = mapLiked[v.Id]
		//response.Data[index].UserLikes = userLikes
	}
	return &response, errCountUser
}

func (i MarketPlaceController) CheckItemCanJoinEvent(walletAddress string, nft string, userId int) (bool, error) {
	listItems, err := i.Store.CheckItemCanJoinEvent(walletAddress, nft, userId)
	return listItems, err

}

func (i MarketPlaceController) CmsStats(collectionId int64, time int64) (marketplacemodel.CollectionStatsResponse, error) {
	listItems, err := i.Store.ListItems(marketplacemodel.ItemFilter{CollectionId: collectionId})
	if err != nil {
		return marketplacemodel.CollectionStatsResponse{}, err
	}
	itemIds := make([]int64, 0, len(listItems))
	for _, v := range listItems {
		itemIds = append(itemIds, v.Id)
	}

	stats, err := i.Store.CollectionStats(itemIds, time)
	if err != nil {
		return marketplacemodel.CollectionStatsResponse{}, err
	}

	totalNFTSold := int64(0)
	mapTotalNFTSoldPerDay := make(map[string]int64)
	totalPriceSold := int64(0)
	mapNftSoldPrice := make(map[string]int64)

	mod := big.NewInt(1000000000000000000)
	resultDetail := make([]marketplacemodel.StatsDetail, 0, 1000)

	if len(stats) == 0 {
		return marketplacemodel.CollectionStatsResponse{
			AvgPrice:  0,
			AvgVolume: 0,
			Detail:    resultDetail,
		}, nil
	}

	for _, v := range stats {
		distance, _ := new(big.Int).SetString(v.Price, 10)
		price := new(big.Int).Div(distance, mod).Int64()

		if v.Status == common.MARKET_PLACE_STATUS_SOLDED {
			totalPriceSold += price
			totalNFTSold += 1
			mapNftSoldPrice[v.UpdatedAt.Format(common.ISO8601)] += price
			mapTotalNFTSoldPerDay[v.UpdatedAt.Format(common.ISO8601)] += 1
		}

	}

	totalPrice := float64(0)
	for k, v := range mapTotalNFTSoldPerDay {
		vPerDay := float64(mapNftSoldPrice[k]) / float64(v)
		if v == 0 {
			vPerDay = 0
		}
		e := marketplacemodel.StatsDetail{
			Time:        k,
			TotalNft:    v,
			TotalVolume: float64(mapNftSoldPrice[k]),
			AvgPrice:    vPerDay,
		}
		totalPrice += vPerDay
		resultDetail = append(resultDetail, e)
	}
	sort.SliceStable(resultDetail, func(i, j int) bool {
		return resultDetail[i].Time < resultDetail[j].Time
	})

	result := marketplacemodel.CollectionStatsResponse{
		AvgPrice:  0,
		AvgVolume: 0,
	}

	result.Detail = resultDetail

	if len(mapTotalNFTSoldPerDay) > 0 {
		result.AvgPrice = totalPrice / float64(len(mapTotalNFTSoldPerDay))
		result.AvgVolume = float64(totalPriceSold) / float64(len(mapTotalNFTSoldPerDay))
	}

	return result, nil

}

func (i MarketPlaceController) ExploreStatistic() (*marketplacemodel.ExploreStatisticResponse, error) {
	total := int64(0)
	data, err := i.Store.ExploreStatistic(&total)
	if err != nil {
		return nil, err
	}

	result := &marketplacemodel.ExploreStatisticResponse{
		Total: total,
	}

	for _, v := range data {
		switch v.Id {
		case common.ART:
			result.Art = v.Total
		case common.GAME:
			result.Game = v.Total
		case common.PHOTOGRAPHY:
			result.Photography = v.Total
		case common.MUSIC:
			result.Music = v.Total
		case common.VIDEO:
			result.Video = v.Total
		case common.SPORT:
			result.Sport = v.Total
		}

	}

	return result, nil

}

func (u MarketPlaceController) ItemHistory(filter usermodel.ActivitiesFilter) ([]marketplacemodel.ActivitiesUserResponse, error) {
	item, err := u.Store.GetItemByID(marketplacemodel.ItemFilter{Nft: filter.Nft, ContractAddress: filter.ContractAddress})
	if err != nil {
		return nil, err
	}

	if item.Id == 0 {
		return nil, common.ErrNotFoundItem
	}

	filter.ItemId = item.Id
	activities, err := u.UserStore.ListActivities(filter)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0, len(activities))
	for _, v := range activities {
		if v.UserId > 0 {
			userIds = append(userIds, v.UserId)
		}
		if v.ActivityUserId > 0 {
			userIds = append(userIds, v.ActivityUserId)
		}
	}

	mapUser := make(map[int64]usermodel.UserDetail)
	if len(userIds) > 0 {
		users, err := u.UserStore.GetListUserById(userIds)
		if err != nil {
			return nil, err
		}
		for _, v := range users {
			mapUser[v.Id] = v
		}
	}

	result := make([]marketplacemodel.ActivitiesUserResponse, 0, len(activities))
	for _, v := range activities {
		data := make(map[string]string)
		_ = json.Unmarshal([]byte(v.MetaData), &data)
		result = append(result, marketplacemodel.ActivitiesUserResponse{
			Id:           v.Id,
			UserId:       v.UserId,
			ItemId:       v.ItemId,
			FollowId:     v.FollowId,
			MetaData:     data,
			CollectionId: v.CollectionId,
			ActivityType: v.ActivityType,
			ActivityUser: mapUser[v.ActivityUserId],
			Item:         *item,
			User:         mapUser[v.UserId],
			CreatedAt:    v.CreatedAt,
		})
	}

	return result, nil

}

func (u MarketPlaceController) CollectionHistory(filter usermodel.ActivitiesFilter) ([]marketplacemodel.ActivitiesUserResponse, error) {
	collection, err := u.Store.FindCollection(&marketplacemodel.Collection{Id: filter.CollectionId})
	if err != nil {
		return nil, err
	}

	if collection.Id == 0 {
		return nil, common.ErrNotFoundItem
	}

	activities, err := u.UserStore.ListActivities(filter)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0, len(activities))
	itemIds := make([]int64, 0, len(activities))
	for _, v := range activities {
		if v.UserId > 0 {
			userIds = append(userIds, v.UserId)
		}

		if v.ActivityUserId > 0 {
			userIds = append(userIds, v.ActivityUserId)
		}

		if v.ItemId > 0 {
			itemIds = append(itemIds, v.ItemId)
		}
	}

	mapItem := make(map[int64]marketplacemodel.Item)
	if len(itemIds) > 0 {
		items, err := u.Store.ListItems(marketplacemodel.ItemFilter{Ids: itemIds})
		if err != nil {
			return nil, err
		}
		for _, v := range items {
			mapItem[v.Id] = v
		}
	}

	mapUser := make(map[int64]usermodel.UserDetail)
	if len(userIds) > 0 {
		users, err := u.UserStore.GetListUserById(userIds)
		if err != nil {
			return nil, err
		}
		for _, v := range users {
			mapUser[v.Id] = v
		}
	}

	result := make([]marketplacemodel.ActivitiesUserResponse, 0, len(activities))
	for _, v := range activities {
		data := make(map[string]string)
		_ = json.Unmarshal([]byte(v.MetaData), &data)
		result = append(result, marketplacemodel.ActivitiesUserResponse{
			Id:           v.Id,
			UserId:       v.UserId,
			ItemId:       v.ItemId,
			FollowId:     v.FollowId,
			MetaData:     data,
			CollectionId: v.CollectionId,
			ActivityType: v.ActivityType,
			Item:         mapItem[v.ItemId],
			User:         mapUser[v.UserId],
			ActivityUser: mapUser[v.ActivityUserId],
			Collection:   *collection,
			CreatedAt:    v.CreatedAt,
		})
	}

	return result, nil

}

func (i *MarketPlaceController) NFTHide(rq marketplacemodel.HideNftRequest) error {
	userInfo, errInfo := i.UserStore.GetUserDetailById(rq.UserId)
	if rq.IsAdmin != 1 {
		if errInfo != nil {
			return errInfo
		}
		if userInfo.Id == 0 {
			return common.ErrNotFoundUser
		}
		if userInfo.Type < 6 {
			return common.ErrNotPermissionHideItem
		}
	}

	err := i.Store.HideNftCms(rq.Id)
	if err != nil {
		return err
	}

	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{Id: rq.Id})
	if err != nil {
		return err
	}

	user, err := i.UserStore.GetUserDetailById(item.UserId)
	if err != nil {
		return err
	}

	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      item.Title,
		"item_image":     item.Image,
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"wallet_address": user.WalletAddress,
	})

	listNoti := []*usermodel.Notification{
		{
			UserId:             rq.UserId,
			LastTimeRead:       0,
			NotificationType:   common.HIDE,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		},
		{
			UserId:             user.Id,
			LastTimeRead:       0,
			NotificationType:   common.HIDE_NFT,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		},
	}

	if len(listNoti) > 0 {
		err = i.UserStore.BatchInsertNotification(listNoti)
		if err != nil {
			return err
		}
	}

	return nil

}
func (i *MarketPlaceController) NFTUnHide(rq marketplacemodel.HideNftRequest) error {
	if rq.IsAdmin != 1 {
		userInfo, errInfo := i.UserStore.GetUserDetailById(rq.UserId)
		if errInfo != nil {
			return errInfo
		}
		if userInfo.Id == 0 {
			return common.ErrNotFoundUser
		}
		if userInfo.Type < 6 {
			return common.ErrNotPermissionHideItem
		}
	}
	err := i.Store.UnHideNftCms(rq.Id)
	if err != nil {
		return err
	}
	return nil
}

func (i *MarketPlaceController) NFTEventStatistic() (map[string]int64, error) {
	totalNft, err := i.Store.CountTotalNft()
	if err != nil {
		return nil, err
	}

	totalNftJoinEvent, err := i.Store.NFTStatistic()
	if err != nil {
		return nil, err
	}

	return map[string]int64{
		"total_nft":       totalNft,
		"total_nft_event": totalNftJoinEvent,
	}, nil
}

func (i *MarketPlaceController) ExploreCollection(filter marketplacemodel.ExploreCollectionFilter, paging *common.Paging) ([]*marketplacemodel.CollectionExploreResponse, error) {
	res, err := i.Store.CollectionExplore(filter, paging)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (i *MarketPlaceController) ExploreNft(filter marketplacemodel.ExploreNftFilter, paging *common.Paging) ([]*marketplacemodel.ExploreNftResponse, error) {
	res, err := i.Store.ExploreNft(filter, paging)
	if err != nil {
		return nil, err
	}

	result := make([]*marketplacemodel.ExploreNftResponse, 0, len(res))

	for _, v := range res {
		tmp := &marketplacemodel.ExploreNftResponse{
			Price:      v.Price,
			ItemDetail: v.ConvertToDetail(),
		}
		tmp.IsLike = v.IsLike
		result = append(result, tmp)
	}

	return result, nil
}

func (i MarketPlaceController) UpdateBrandImageCms(rq *marketplacemodel.UpdateBrandImageCmsRequest, options map[string]interface{}) error {
	var update = make(map[string]interface{})
	if strings.Contains(rq.ItemBackgroundImage, "/tmp/") {
		backgroundimage := utils.DESTINATION + utils.GetFileName(rq.ItemBackgroundImage)
		if len(rq.ItemBackgroundImage) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(rq.ItemBackgroundImage, cast.ToString(options["host"]), ""), backgroundimage); err != nil {
				return err
			}
		}
		rq.ItemBackgroundImage = cast.ToString(options["host"]) + backgroundimage
	}
	if strings.Contains(rq.ItemAvatarImage, "/tmp/") {
		backgroundimage := utils.DESTINATION + utils.GetFileName(rq.ItemAvatarImage)
		if len(rq.ItemAvatarImage) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(rq.ItemAvatarImage, cast.ToString(options["host"]), ""), backgroundimage); err != nil {
				return err
			}
		}
		rq.ItemAvatarImage = cast.ToString(options["host"]) + backgroundimage
	}
	if rq.DisplayOrder >= 0 {
		update["display_order"] = rq.DisplayOrder
	}
	update["brand_name"] = rq.BrandName
	update["link"] = rq.Link
	update["item_background_image"] = rq.ItemBackgroundImage
	update["item_avatar_image"] = rq.ItemAvatarImage
	if err := i.Store.UpdateBrandImageCms(rq.Id, update); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) CreateBrandCms(rq *marketplacemodel.CreateBrandCmsRequest, options map[string]interface{}) error {
	backgroundimage := utils.DESTINATION + utils.GetFileName(rq.BackgroundImage)
	if len(rq.BackgroundImage) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.BackgroundImage, cast.ToString(options["host"]), ""), backgroundimage); err != nil {
			return err
		}
	}
	brands := marketplacemodel.Brand{
		BrandTitle:      rq.BrandTitle,
		IsHot:           rq.IsHot,
		Description:     rq.Description,
		BackgroundImage: cast.ToString(options["host"]) + backgroundimage,
		DisplayOrder:    rq.DisplayOrder,
	}
	var results []marketplacemodel.BrandImage
	for _, value := range rq.BrandImages {
		itemavatarimage := utils.DESTINATION + utils.GetFileName(value.ItemAvatarImage)
		itembackgroundimage := utils.DESTINATION + utils.GetFileName(value.ItemBackGroundImage)
		if len(value.ItemAvatarImage) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(value.ItemAvatarImage, cast.ToString(options["host"]), ""), itemavatarimage); err != nil {
				return err
			}
		}
		if len(value.ItemBackGroundImage) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(value.ItemBackGroundImage, cast.ToString(options["host"]), ""), itembackgroundimage); err != nil {
				return err
			}
		}
		result := marketplacemodel.BrandImage{
			BrandName:           value.BrandName,
			ItemAvatarImage:     cast.ToString(options["host"]) + itemavatarimage,
			Link:                value.Link,
			ItemBackgroundImage: cast.ToString(options["host"]) + itembackgroundimage,
			DisplayOrder:        value.DisplayOrder,
		}
		results = append(results, result)
	}
	if err := i.Store.CreateBrandCms(brands, results); err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceController) CreateBrandImageCms(rq *marketplacemodel.CreateBrandImageCmsRequest, options map[string]interface{}) (*marketplacemodel.BrandImageDetail, error) {
	itembackgroundimage := utils.DESTINATION + utils.GetFileName(rq.ItemBackGroundImage)
	if len(rq.ItemBackGroundImage) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.ItemBackGroundImage, cast.ToString(options["host"]), ""), itembackgroundimage); err != nil {
			return nil, err
		}
	}
	itemavatarimage := utils.DESTINATION + utils.GetFileName(rq.ItemAvatarImage)
	if len(rq.ItemBackGroundImage) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.ItemBackGroundImage, cast.ToString(options["host"]), ""), itemavatarimage); err != nil {
			return nil, err
		}
	}
	create := &marketplacemodel.BrandImageDetail{
		BrandName:           rq.BrandName,
		Link:                rq.Link,
		BrandId:             rq.BrandId,
		ItemBackgroundImage: cast.ToString(options["host"]) + itembackgroundimage,
		ItemAvatarImage:     cast.ToString(options["host"]) + itemavatarimage,
		DisplayOrder:        rq.DisplayOrder,
	}
	if err := i.Store.CreateBrandImageCms(create); err != nil {
		return nil, err
	}
	return create, nil
}

func (i MarketPlaceController) ListItemToken(rq marketplacemodel.ItemFilter, paging *common.Paging) (map[int64]string, error) {
	var listItems []marketplacemodel.ItemDetail
	var err error
	if rq.IsBuy == 1 {
		listItems, err = i.Store.GetItemWithConditionNewWithMarket(rq, paging)
	} else {
		listItems, err = i.Store.GetItemWithConditionNew(rq, paging)
	}
	if err != nil {
		return nil, err
	}
	if paging.Limit != len(listItems) {
		return nil, errors.New("Not enough item")
	}
	result := make(map[int64]string)
	for _, item := range listItems {
		if rq.IsBuy == 1 {
			result[item.MarketPlaceId] = item.IdOnMarket
		} else {
			result[item.Id] = item.Nft
		}
	}

	return result, nil
}
