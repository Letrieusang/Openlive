package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
)

func (i MarketPlaceController) JoinEventNft(req marketplacemodel.UpdateItemEvent) error {
	err := i.Store.UpdateItemJoinEvent(req)
	if err != nil {
		return err
	}
	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      req.ItemName,
		"user_name":      req.UserDisplay,
		"user_avatar":    req.UserAvatar,
		"wallet_address": req.WalletAddress,
	})

	listNoti := &usermodel.Notification{
		UserId:             req.UserId,
		LastTimeRead:       0,
		NotificationType:   common.JOIN_EVENT,
		Status:             common.UNREAD,
		NotificationObject: base58.Encode(notiObj),
	}

	err = i.UserStore.InsertNotification(listNoti)
	if err != nil {
		return err
	}
	i.BuildUserDailyEvent(req.UserId, common.DAILY_EVENT_TYPE_JOIN_EVENT)
	return nil
}

func (i MarketPlaceController) CancelEventNft(req marketplacemodel.UpdateItemEvent) error {
	oldItem, err := i.Store.FindItem(&marketplacemodel.Item{UserId: req.UserId, Nft: req.Nft, ContractAddress: req.ContractAddress})
	if err != nil {
		return err
	}

	err = i.Store.UpdateItemCancelEvent(req)
	if err != nil {
		return err
	}

	notiObj, _ := json.Marshal(map[string]string{
		"item_name":      oldItem.Title,
		"user_name":      req.UserDisplay,
		"user_avatar":    req.UserAvatar,
		"wallet_address": req.WalletAddress,
	})

	listNoti := &usermodel.Notification{
		UserId:             req.UserId,
		LastTimeRead:       0,
		NotificationType:   common.REMOVE_FROM_EVENT,
		Status:             common.UNREAD,
		NotificationObject: base58.Encode(notiObj),
	}

	err = i.UserStore.InsertNotification(listNoti)
	if err != nil {
		return err
	}

	return nil
}
