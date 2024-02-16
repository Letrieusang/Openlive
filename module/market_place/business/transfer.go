package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"encoding/json"
)

func (i MarketPlaceController) TransferNft(req marketplacemodel.UpdateItemOwnerRequest) error {
	err := i.Store.UpdateOwnerOfItem(req)
	if err != nil {
		return err
	}

	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{
		Nft:             req.Nft,
		ContractAddress: req.ContractAddress,
	})
	if err != nil {
		return err
	}

	metadata, _ := json.Marshal(map[string]string{
		"is_owner": "true",
	})
	if req.ReceiverId > 0 {
		activities := []*usermodel.Activities{
			{
				UserId:         req.SenderId,
				ItemId:         item.Id,
				ActivityType:   common.TRANSFERED_ITEM,
				ActivityUserId: req.ReceiverId,
			},
		}
		tmpActivity := usermodel.Activities{
			UserId:         req.ReceiverId,
			ItemId:         item.Id,
			ActivityType:   common.TRANSFERED_ITEM,
			ActivityUserId: req.SenderId,
			MetaData:       string(metadata),
		}
		activities = append(activities, &tmpActivity)
		err = i.UserStore.InsertActivities(activities...)
		if err != nil {
			return err
		}
	} else {
		activities := []*usermodel.Activities{
			{
				UserId:       req.SenderId,
				ItemId:       item.Id,
				ActivityType: common.TRANSFERED_ITEM,
			},
		}
		err = i.UserStore.InsertActivities(activities...)
		if err != nil {
			return err
		}
	}

	return nil
}

/*func (i MarketPlaceController) CreateActivity(req marketplacemodel.CreateActivityInfo) error {
	item, err := i.Store.GetItemByID(marketplacemodel.ItemFilter{
		Nft:             req.Nft,
		ContractAddress: req.ContractAddress,
	})
	if err != nil {
		return err
	}

	metadata, _ := json.Marshal(map[string]string{
		"is_owner": "true",
	})

	activities := []*usermodel.Activities{
		{
			UserId:         req.SenderId,
			ItemId:         item.Id,
			ActivityType:   req.Type,
			ActivityUserId: req.ReceiverId,
		},
	}
	if req.ReceiverId > 0 {
		tmpActivity := usermodel.Activities{
			UserId:         req.ReceiverId,
			ItemId:         item.Id,
			ActivityType:   req.Type,
			ActivityUserId: req.SenderId,
			MetaData:       string(metadata),
		}
		activities = append(activities, &tmpActivity)
	}

	err = i.UserStore.InsertActivities(activities...)
	if err != nil {
		return err
	}

	return nil
}*/
