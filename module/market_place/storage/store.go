package marketplacestorage

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strings"
	"time"
)

type MarketPlaceStore struct {
	db *gorm.DB
}

func NewMarketPlaceStore(db *gorm.DB) *MarketPlaceStore {
	return &MarketPlaceStore{db: db}
}

func (m MarketPlaceStore) InsertMarketPlace(marketPlace *marketplacemodel.MarketPlace) error {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if err := tx.Table(marketplacemodel.MarketPlace{}.TableName()).Create(&marketPlace).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(marketplacemodel.Item{}.TableName()).Where("id = ?", marketPlace.ItemId).Update("status", common.ITEM_ON_SALE).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m MarketPlaceStore) InsertVoteItem(in *marketplacemodel.Vote) error {
	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if err := tx.Table(marketplacemodel.Vote{}.TableName()).Create(&in).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(marketplacemodel.MarketPlace{}.TableName()).
		Where("id = ?", in.MarketPlaceId).
		UpdateColumn("total_vote", gorm.Expr("total_vote + ?", 1)).Error; err != nil {
		return err
	}

	return tx.Commit().Error

}

func (m MarketPlaceStore) InsertLikeItem(in *marketplacemodel.Like) error {
	var out *marketplacemodel.Like
	err := m.db.Table(marketplacemodel.Like{}.TableName()).
		Where("item_id = ?", in.ItemId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return err.Error
	}

	if out.Id != 0 {
		if out.IsDeleted == common.ACTIVE {
			return common.ErrUserLiked
		}
		in.Id = out.Id
	}

	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	result := tx.Table(marketplacemodel.Like{}.TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_deleted"}),
	}).Create(in)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	var found marketplacemodel.Item
	if err := tx.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", in.ItemId).
		Find(&found).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", in.ItemId).
		Updates(map[string]interface{}{
			"total_like": gorm.Expr("total_like + ?", 1),
			"updated_at": found.UpdatedAt,
		}).Error; err != nil {
		return err
	}

	return tx.Commit().Error

}

func (m MarketPlaceStore) DisLikeItem(in *marketplacemodel.Like) error {
	var out *marketplacemodel.Like
	err := m.db.Table(marketplacemodel.Like{}.TableName()).
		Where("item_id = ?", in.ItemId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return common.ErrUserNotLiked
		}
		return err.Error
	}

	if out.IsDeleted == common.DELETED {
		return common.ErrUserDisLiked
	}

	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if out.IsDeleted == common.ACTIVE {
		if err := tx.Table(marketplacemodel.Like{}.TableName()).
			Where("id = ?", out.Id).
			Updates(map[string]interface{}{
				"is_deleted": common.DELETED,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}

		var found marketplacemodel.Item
		if err := tx.Table(marketplacemodel.Item{}.TableName()).
			Where("id = ?", in.ItemId).
			Find(&found).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Table(marketplacemodel.Item{}.TableName()).
			Where("id = ?", in.ItemId).
			Updates(map[string]interface{}{
				"total_like": gorm.Expr("total_like - ?", 1),
				"updated_at": found.UpdatedAt,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}

func (m MarketPlaceStore) InsertMarketPlaceSold(marketPlace *marketplacemodel.MarketPlaceSold) error {
	return m.db.Table(marketplacemodel.MarketPlaceSold{}.TableName()).Create(&marketPlace).Error
}

func (m MarketPlaceStore) InsertTransactionLog(logs *marketplacemodel.TransactionLog) error {
	return m.db.Table(marketplacemodel.TransactionLog{}.TableName()).Create(&logs).Error
}

func (m MarketPlaceStore) UpdateMarketPlaceStatus(market *marketplacemodel.MarketPlace) error {
	return m.db.Table(marketplacemodel.MarketPlace{}.TableName()).Where("id = ?", market.Id).Updates(&market).Error
}

func (m MarketPlaceStore) GetMarketPlaceByItemId(marketPlace *marketplacemodel.MarketPlaceFilter) (*marketplacemodel.MarketPlace, error) {
	data := &marketplacemodel.MarketPlace{}
	err := m.buildMarketCommonQuery(*marketPlace).Joins("Item").Joins("Seller").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m MarketPlaceStore) ListTransactionLogs(marketPlace marketplacemodel.TransactionLogFilter) ([]*marketplacemodel.TransactionLog, error) {
	data := make([]*marketplacemodel.TransactionLog, 0, 1000)
	err := m.buildTransactionLogQuery(marketPlace).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m MarketPlaceStore) ListTransactionLogsWithTime(marketPlace marketplacemodel.TransactionLogFilter, paging *common.PagingWithTime) ([]*marketplacemodel.TransactionLog, error) {
	data := make([]*marketplacemodel.TransactionLog, 0, 1000)
	tx := m.buildTransactionLogQuery(marketPlace)
	tx.Where("transaction_logs.created_at BETWEEN ? AND ?", paging.CustomStart, paging.End)
	err := tx.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m MarketPlaceStore) ListMarketPlace(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.MarketPlace, error) {
	result := make([]*marketplacemodel.MarketPlace, 0, 1000)
	tx := m.buildListMarketPlaceCommonQuery(rq)

	if paging != nil {
		errCount := tx.Where("market_places.deleted_at IS NULL").Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	err := tx.
		Joins("Item").
		Joins("Seller").
		Joins("LEFT JOIN collections as c ON c.id = Item.collection_id").
		Where("Item.is_hide = ?", common.ITEM_UNHIDE).
		Order("id desc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) ListMarketPlaceWithTime(rq marketplacemodel.MarketPlaceFilter, paging *common.PagingWithTime) ([]*marketplacemodel.MarketPlace, error) {
	result := make([]*marketplacemodel.MarketPlace, 0, 1000)
	tx := m.buildListMarketPlaceCommonQuery(rq)

	if paging != nil {
		errCount := tx.Where("market_places.deleted_at IS NULL").Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	err := tx.
		Joins("Item").
		Joins("Seller").
		Joins("LEFT JOIN collections as c ON c.id = Item.collection_id").
		Where("Item.is_hide = ?", common.ITEM_UNHIDE).
		Where("market_places.created_at BETWEEN ? AND ?", paging.CustomStart, paging.End).
		Order("id desc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) ListMarketPlaceSold(rq marketplacemodel.MarketPlaceSoldFilter) ([]*marketplacemodel.MarketPlaceSold, error) {
	result := make([]*marketplacemodel.MarketPlaceSold, 0, 1000)
	tx := m.buildMarketPlaceSoldQuery(rq)
	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) buildMarketPlaceSoldQuery(filter marketplacemodel.MarketPlaceSoldFilter) *gorm.DB {
	tx := m.db.Table(marketplacemodel.MarketPlaceSold{}.TableName())
	if len(filter.ItemIds) > 0 {
		tx.Where("item_id IN ?", filter.ItemIds)
	}
	if len(filter.MIds) > 0 {
		tx.Where("market_place_id IN ?", filter.MIds)
	}
	return tx
}

func (i MarketPlaceStore) UpdateMarketPlaceCms(id int64, update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (m MarketPlaceStore) ListMarketPlaceRebuild(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.MarketPlaceWithCollectionName, error) {
	result := make([]*marketplacemodel.MarketPlaceWithCollectionName, 0, 1000)
	tx := m.buildListMarketPlaceCommonQuery(rq)

	if paging != nil {
		errCount := tx.Where("market_places.deleted_at IS NULL").Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	query := tx.Select("market_places.*, c.name as collection_name").
		Joins("Item").
		Joins("Seller").
		Joins("LEFT JOIN collections as c ON c.id = Item.collection_id").
		Where("Item.is_hide = ?", common.ITEM_UNHIDE).
		Order("market_places.display_order DESC")
	if rq.IsTrending > 0 {
		query.Where("Item.is_trending = ?", rq.IsTrending)
	}
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) ListMostLikeMarketPlace(limit int) ([]*marketplacemodel.MarketPlace, error) {
	result := make([]*marketplacemodel.MarketPlace, 0, 1000)
	tx := m.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Where("market_places.status = ?", common.ITEM_STATUS_ON_MARKET).
		Order("Item.total_like DESC")

	err := tx.
		Joins("Item").
		Joins("Seller").Where("Item.collection_id <> 0 OR Item.collection_id IS NOT NULL").
		Limit(limit).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) ListVoting(rq *marketplacemodel.VoteFilter) ([]*marketplacemodel.Vote, error) {
	var result []*marketplacemodel.Vote
	tx := m.buildVotingCommonQuery(rq)
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) InsertStaking(rq *marketplacemodel.Staking) (*marketplacemodel.Staking, error) {
	err := m.db.Table(marketplacemodel.Staking{}.TableName()).Create(&rq).Error
	if err != nil {
		return nil, err
	}
	return rq, nil
}

func (m MarketPlaceStore) FindStaking(rq *marketplacemodel.Staking) (*marketplacemodel.Staking, error) {
	var out *marketplacemodel.Staking
	err := m.db.Table(marketplacemodel.Staking{}.TableName()).Where(&rq).Take(&out)
	if err.Error != nil {
		return nil, err.Error
	}
	return out, nil
}

func (m MarketPlaceStore) UpdateStaking(rq *marketplacemodel.Staking) error {
	err := m.db.Table(marketplacemodel.Staking{}.TableName()).Where("id = ?", rq.Id).Updates(&rq).Error
	if err != nil {
		return err
	}
	return nil
}

func (m MarketPlaceStore) IsStakingByUser(rq *marketplacemodel.Staking) (bool, error) {
	var has bool
	err := m.db.Table(marketplacemodel.Staking{}.TableName()).
		Select("count(*) > 0").
		Where("user_id = ?", rq.UserId).
		Where("status = ?", common.STAKING).
		Take(&has).Error
	if err != nil {
		return false, err
	}

	if !has {
		return false, err
	}

	return true, nil
}

func (m MarketPlaceStore) IsExistMarketPlace(rq *marketplacemodel.MarketPlace) (bool, error) {
	var has bool
	err := m.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Select("count(*) > 0").
		Where("id = ?", rq.Id).
		Take(&has).Error
	if err != nil {
		return false, err
	}

	if !has {
		return false, err
	}

	return true, nil
}

func (m MarketPlaceStore) buildMarketCommonQuery(filter marketplacemodel.MarketPlaceFilter) *gorm.DB {
	tx := m.db.Table(marketplacemodel.MarketPlace{}.TableName())
	if filter.Id != 0 {
		tx.Where("market_places.id = ?", filter.Id)
	}
	if filter.ItemId != 0 {
		tx.Where("market_places.item_id = ?", filter.ItemId)
	}
	if filter.IdOnMarket > 0 {
		tx.Where("market_places.id_on_market = ?", filter.IdOnMarket)
	}
	if filter.Status > 0 {
		tx.Where("market_places.status = ?", filter.Status)
	}
	if filter.ContractAddress != "" {
		tx.Where("BINARY Item.contract_address = ?", filter.ContractAddress)
	}
	if filter.Nft != "" {
		tx.Where("Item.nft = ?", filter.Nft)
	}
	if filter.MarketContract != "" {
		tx.Where("BINARY market_places.contract_address = ?", filter.MarketContract)
	}
	if filter.SellType > 0 {
		tx.Where("sell_type = ?", filter.SellType)
	}
	return tx
}

func (m MarketPlaceStore) buildTransactionLogQuery(filter marketplacemodel.TransactionLogFilter) *gorm.DB {
	tx := m.db.Table(marketplacemodel.TransactionLog{}.TableName())
	if len(filter.Cols) > 0 {
		tx.Select(filter.Cols)
	}
	if len(filter.ItemIds) > 0 {
		tx.Where("item_id IN ?", filter.ItemIds)
	}
	return tx
}

func (m MarketPlaceStore) buildVotingCommonQuery(filter *marketplacemodel.VoteFilter) *gorm.DB {
	tx := m.db.Table(marketplacemodel.Vote{}.TableName())
	if len(filter.Cols) > 0 {
		tx.Select(filter.Cols)
	}
	if len(filter.Distinct) > 0 {
		tx.Distinct(filter.Distinct)
	}
	return tx
}

func (m MarketPlaceStore) buildListMarketPlaceCommonQuery(rq marketplacemodel.MarketPlaceFilter) *gorm.DB {
	tx := m.db.Table(marketplacemodel.MarketPlace{}.TableName())
	if rq.Status != 0 {
		tx.Where("market_places.status = ?", rq.Status)
	}
	if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.IsHot)) {
		tx.Order(fmt.Sprintf("c.is_trendings %v", strings.ToUpper(rq.IsHot)))
	}
	if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.IsHot)) {
		tx.Order(fmt.Sprintf("c.is_hot %v", strings.ToUpper(rq.IsHot)))
	}
	if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Price)) {
		tx.Order(fmt.Sprintf("market_places.price *1 %v", strings.ToUpper(rq.Price)))
	}
	if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.CreatedAt)) {
		tx.Order(fmt.Sprintf("market_places.created_at %v", strings.ToUpper(rq.CreatedAt)))
	}
	if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.TotalLike)) {
		tx.Order(fmt.Sprintf("Item.total_like %v", strings.ToUpper(rq.TotalLike)))
	}
	if len(rq.ItemIds) > 0 {
		tx.Where("market_places.item_id IN ?", rq.ItemIds)
	}
	if rq.CategoryId != 0 {
		subQuery := m.db.Table(marketplacemodel.Item{}.TableName()).Where("category_id = ?", rq.CategoryId).Select("id").Group("id")
		tx.Where("market_places.item_id IN (?) ", subQuery)
	}
	if rq.CollectionId != 0 {
		subQuery := m.db.Table(marketplacemodel.Item{}.TableName()).Where("collection_id = ?", rq.CollectionId).Select("id").Group("id")
		tx.Where("market_places.item_id IN (?) ", subQuery)
	}
	if rq.PriceMin > 0 {
		tx.Where("market_places.price >= ?", rq.PriceMin)
	}
	if rq.PriceMax > 0 {
		tx.Where("market_places.price <= ?", rq.PriceMax)
	}
	if len(rq.IsHot) > 0 {
		tx.Where("market_places.is_hot = ?", rq.IsHot)
	}
	if rq.Creator != 0 {
		status := 0
		if rq.Creator == common.VIP_USER {
			status = 1
		}
		if rq.Creator == common.BASIC_USER {
			status = 2
		}
		subQuery := m.db.Table(usermodel.User{}.TableName()).Where("status = ?", status).Select("id").Group("id")
		tx.Where("market_places.seller_id IN (?) ", subQuery)
	}
	if len(rq.ItemIds) > 0 {
		tx.Where("market_places.item_id IN ?", rq.ItemIds)
	}

	if rq.IsDiscover == 1 {
		subQuery := m.db.Table(marketplacemodel.Item{}.TableName()).
			Joins("INNER JOIN collections as c ON c.id = items.collection_id").
			Where("c.is_hold = ?", 1).Select("items.id").Group("id")
		tx.Where("NOT market_places.item_id IN (?) ", subQuery)
	}

	if rq.SellType > 0 {
		tx.Where("market_places.sell_type = ?", rq.SellType)
	}

	return tx
}

func (i MarketPlaceStore) GetItemWithCondition(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.Item, error) {
	var result []marketplacemodel.Item

	tx := i.buildItemCommonQuery(rq)

	if paging != nil {
		if err := tx.Count(&paging.Total).Error; err != nil {
			return nil, err
		}
		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	if len(rq.Price) > 0 || len(rq.Favourite) > 0 {
		tx.Joins("LEFT JOIN market_places ON items.id = market_places.item_id").
			Select("items.*, market_places.price")
		if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Favourite)) {
			tx.Order(fmt.Sprintf("items.total_like %v", rq.Favourite))
		}

		if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Price)) {
			tx.Order(fmt.Sprintf("market_places.price %v", rq.Price))
		}

	}
	if len(rq.View) > 0 && utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.View)) {
		tx.Order(fmt.Sprintf("view %v", rq.View))
	}
	if len(rq.Created) > 0 && utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Created)) {
		tx.Order(fmt.Sprintf("created_at %v", rq.Created))
	}
	if len(rq.OrderBy) > 0 {
		tx.Order(rq.OrderBy)
	}

	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) GetItemWithConditionNew(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.ItemDetail, error) {
	var result []marketplacemodel.ItemDetail

	tx := i.buildItemCommonQuery(rq)

	if paging != nil {
		if err := tx.Count(&paging.Total).Error; err != nil {
			return nil, err
		}
		if paging.Limit != -1 {
			tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
		}
	}

	if rq.IsOwned == 1 {
		tx.Joins("LEFT JOIN transaction_logs ON transaction_logs.id = items.id").Order("transaction_logs.created_at DESC")
	}

	if len(rq.Price) > 0 || len(rq.Favourite) > 0 {
		tx.Joins("LEFT JOIN market_places ON items.id = market_places.item_id").
			Select("items.*, market_places.price")
		if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Favourite)) {
			tx.Order(fmt.Sprintf("items.total_like %v", rq.Favourite))
		}

		if utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Price)) {
			tx.Order(fmt.Sprintf("market_places.price %v", rq.Price))
		}

	}
	if len(rq.View) > 0 && utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.View)) {
		tx.Order(fmt.Sprintf("view %v", rq.View))
	}
	if len(rq.Created) > 0 && utils.InStringSlice(utils.OrderBy, strings.ToUpper(rq.Created)) {
		tx.Order(fmt.Sprintf("created_at %v", rq.Created))
	}
	if len(rq.OrderBy) > 0 {
		tx.Order(rq.OrderBy)
	}

	tx.Group("items.id")

	if err := tx.Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) GetItemWithConditionNewWithMarket(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.ItemDetail, error) {
	var result []marketplacemodel.ItemDetail

	tx := i.buildItemCommonQuery(rq)
	tx.Joins("INNER JOIN market_places ON items.id = market_places.item_id")
	tx.Select("items.*, item_creators.user_id as item_creator_id, market_places.id as market_place_id, id_on_market")
	tx.Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SELLING)
	if paging != nil {
		if err := tx.Count(&paging.Total).Error; err != nil {
			return nil, err
		}
		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	tx.Group("items.id")

	if err := tx.Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) GetItemByID(rq marketplacemodel.ItemFilter) (*marketplacemodel.Item, error) {
	var result *marketplacemodel.Item
	tx := i.buildItemCommonQuery(rq)
	if err := tx.Limit(1).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListItems(rq marketplacemodel.ItemFilter) ([]marketplacemodel.Item, error) {
	var result []marketplacemodel.Item
	tx := i.buildItemCommonQuery(rq)
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) GetListItemById(listUserId []int64) ([]marketplacemodel.Item, error) {
	var result []marketplacemodel.Item
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id IN ?", listUserId).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) buildItemCommonQuery(filter marketplacemodel.ItemFilter) *gorm.DB {
	tx := i.db.Table(marketplacemodel.Item{}.TableName()).Select("items.*, item_creators.user_id as item_creator_id").
		Joins("INNER JOIN item_creators ON items.id = item_creators.item_id")
	tx.Where("items.deleted_at IS NULL")
	if filter.IsTrending >= 0 {
		tx.Where("is_trending = ?", filter.IsTrending)
	}
	if filter.Id != 0 {
		tx.Where("items.id = ?", filter.Id)
	}
	if len(filter.Ids) > 0 {
		tx.Where("items.id IN ?", filter.Ids)
	}
	if filter.CategoryID != 0 {
		tx.Where("category_id = ?", filter.CategoryID)
	}
	if filter.Nft != "" {
		tx.Where("nft = ?", filter.Nft)
	}
	if filter.Status > 0 {
		tx.Where("items.status = ?", filter.Status)
	}
	if filter.CollectionId != 0 {
		tx.Where("collection_id = ?", filter.CollectionId)
	}

	if filter.ContractAddress != "" {
		tx.Where("BINARY contract_address = ?", filter.ContractAddress)
	}
	if filter.NotCollection == 1 {
		tx.Where("collection_id = ? or collection_id IS NULL", 0)
	}
	if len(filter.Stts) > 0 {
		tx.Where("items.status IN ?", filter.Stts)
	}
	if len(filter.CollectionIds) > 0 {
		tx.Where("collection_id IN ?", filter.CollectionIds)
	}
	if len(filter.OrNotListItemIds) > 0 {
		tx.Where("(NOT items.id IN ?)", filter.OrNotListItemIds)
	}
	/*if filter.IsApproval > 0 {
		isApprove := filter.IsApproval
		if isApprove != common.APPROVE {
			isApprove = 0
		}
		tx.Where("is_approval = ?", isApprove)
	}*/
	if filter.IsJoinEvent == 1 {
		tx.Where("is_join_event = ?", common.ITEM_ON_JOIN_EVENT)
	}

	if filter.IsCreated == 1 {
		tx.Where("item_creators.user_id = ?", filter.UserId)
		tx.Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC)
		if filter.IsOwned == 0 {
			filter.UserId = 0
		}
	}

	if filter.IsOwned == 1 {
		tx.Where("items.user_id = ?", filter.UserId)
		tx.Order("updated_at DESC")
		filter.UserId = 0
	}

	if filter.IsFavourite == 1 {
		subQuery := i.db.Table(marketplacemodel.Like{}.TableName()).
			Where("likes.user_id = ?", filter.UserId).
			Where("is_deleted = ?", common.ACTIVE).
			Select("item_id").Group("likes.item_id")
		tx.Where("items.id IN (?)", subQuery).Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC)
		filter.UserId = 0
	}

	if filter.IsSold == 1 {
		subQuery := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
			Where("seller_id = ?", filter.UserId).
			Where("status = ?", common.MARKET_PLACE_STATUS_SOLDED).
			Select("item_id").Group("market_places.item_id")
		tx.Where("items.id IN (?)", subQuery).Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC)
		filter.UserId = 0
	}

	if filter.UserId != 0 {
		tx.Where("items.user_id = ?", filter.UserId)
	}

	if len(filter.Search) > 0 {
		tx.Where("title LIKE ?", "%"+filter.Search+"%")
	}

	if filter.IsGetMulti == 1 {
		tx.Where("(items.main_item_id = ? OR items.id = ?)", filter.IsGetMultiId, filter.IsGetMultiId)
	}

	if filter.IsNotGetMulti == 1 {
		//tx.Where("(items.main_item_id IS NULL OR items.main_item_id = 0)")
	}

	if filter.IsHide >= 0 {
		tx.Where("items.is_hide = ?", filter.IsHide)
	}

	tx.Where("items.deleted_at IS NULL")
	return tx
}

func (i MarketPlaceStore) UpdateItemStatus(id []int64, status int) error {
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id IN ?", id).
		Updates(map[string]interface{}{
			"status":      status,
			"is_approval": common.APPROVE,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateItemCms(id int64, update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateItemSync(items []marketplacemodel.Item) error {
	tx := i.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	for _, item := range items {
		if err := tx.Table(marketplacemodel.Item{}.TableName()).Where("id = ?", item.Id).
			Updates(map[string]interface{}{
				"status":           item.Status,
				"nft":              item.Nft,
				"contract_address": item.ContractAddress,
				"transaction_id":   item.TransactionId,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (i MarketPlaceStore) UpdateItemUserId(id, userId, colectionId int64, status int) error {
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).Where("id = ?", id).
		Update("user_id", userId).
		Update("status", status).
		Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) InsertCollection(rq *marketplacemodel.Collection) error {
	if err := i.db.Table(marketplacemodel.Collection{}.TableName()).Create(rq).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) CheckNameCollection(rq *marketplacemodel.Collection) (bool, error) {
	var has bool
	if err := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Where("name = ?", rq.Name).
		Select("count(id) > 0").
		Limit(1).
		Take(&has).Error; err != nil {
		return false, err
	}
	return has, nil
}

func (i MarketPlaceStore) InsertCategory(rq *marketplacemodel.Category) error {
	if err := i.db.Table(marketplacemodel.Category{}.TableName()).Create(rq).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateCategory(rq *marketplacemodel.Category) error {
	if err := i.db.Table(marketplacemodel.Category{}.TableName()).Where("id = ?", rq.Id).Updates(rq).Error; err != nil {
		return err
	}
	return nil

}
func (i MarketPlaceStore) DeleteCategory(rq *marketplacemodel.Category) error {
	if err := i.db.Table(marketplacemodel.Category{}.TableName()).Delete(rq).Error; err != nil {
		return err
	}
	return nil
}
func (i MarketPlaceStore) ListCategories(rq *marketplacemodel.CategoryFilter) ([]*marketplacemodel.Category, error) {
	var result []*marketplacemodel.Category
	if err := i.db.Table(marketplacemodel.Category{}.TableName()).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListCollections(rq *marketplacemodel.CollectionFilter, paging *common.Paging) ([]*marketplacemodel.Collection, error) {
	var result []*marketplacemodel.Collection
	subQuery := fmt.Sprintf("(select count(items.id) from items where collection_id = collections.id and items.status = %v And items.deleted_at IS NULL)", common.ITEM_STATUS_ON_MARKET)
	query := i.db.Table(marketplacemodel.Collection{}.TableName())
	query.Where("collections.deleted_at IS NULL")
	if rq.IsHot >= 0 {
		query.Where("is_hot = ?", rq.IsHot)
	}
	if rq.IsTrending >= 0 {
		query.Where("is_trending = ?", rq.IsTrending)
	}
	if rq.Status > 0 {
		query.Where("status = ?", rq.Status)
	}
	if rq.UserId > 0 {
		query.Where("user_id = ?", rq.UserId)
	} else if rq.IsCms != 1 {
		query.Where(subQuery + ">" + fmt.Sprint(common.COUNT_MIN_NFT_SHOW_COLLECTION))
	}
	if len(rq.Search) > 0 {
		query.Where("name LIKE ?", "%"+rq.Search+"%")
	}
	if len(rq.OrderBy) > 0 {
		query.Order(rq.OrderBy)
	}
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return nil, nil
		}
		if paging.Limit != -1 {
			query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
		}
	}
	if err := query.Select(fmt.Sprintf("*, %v AS count_items", subQuery)).Order("is_hot desc, count_items desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListCollectionsByCondition(rq *marketplacemodel.CollectionFilter, paging *common.Paging) ([]*marketplacemodel.Collection, error) {
	var result []*marketplacemodel.Collection
	query := i.db.Table(marketplacemodel.Collection{}.TableName())
	if len(rq.UserIds) > 0 {
		query.Where("user_id IN ?", rq.UserIds)
	}
	if len(rq.Cols) > 0 {
		query.Select(rq.Cols)
	}
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return nil, nil
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListCollectionByIds(ids []int64) ([]*marketplacemodel.Collection, error) {
	var result []*marketplacemodel.Collection
	query := i.db.Table(marketplacemodel.Collection{}.TableName())
	query.Where("id IN ?", ids)
	query.Where("collections.deleted_at IS NULL")
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ItemCreator(rq *marketplacemodel.ItemCreator) error {
	if err := i.db.Table(marketplacemodel.ItemCreator{}.TableName()).Create(rq).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) ItemCreators(rq []marketplacemodel.ItemCreator) error {
	if err := i.db.Table(marketplacemodel.ItemCreator{}.TableName()).Create(&rq).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) InsertItem(rq marketplacemodel.Item) (int64, error) {
	tx := i.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n", r)
			tx.Rollback()
		}
	}()

	err := tx.Table(marketplacemodel.Item{}.TableName()).Create(&rq).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return 0, err
	}
	errCreator := i.ItemCreator(&marketplacemodel.ItemCreator{
		ItemId: rq.Id,
		UserId: rq.UserId,
	})
	if errCreator != nil {
		tx.Rollback()
		log.Println(errCreator)
		return 0, errCreator
	}
	tx.Commit()
	return rq.Id, nil
}

func (i MarketPlaceStore) InsertItems(item marketplacemodel.Item, rq []marketplacemodel.Item) (int64, error) {
	tx := i.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n", r)
			tx.Rollback()
		}
	}()
	errMain := tx.Table(marketplacemodel.Item{}.TableName()).Create(&item).Error
	if errMain != nil {
		tx.Rollback()
		return 0, errMain
	}
	for index, _ := range rq {
		if rq[index].Id > 0 {
			rq[index].Id = 0
		}
		rq[index].MainItemId = item.Id
	}

	err := tx.Table(marketplacemodel.Item{}.TableName()).Create(&rq).Error
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return 0, err
	}
	var itemCreators []marketplacemodel.ItemCreator
	var tmpMainCreator marketplacemodel.ItemCreator
	tmpMainCreator.ItemId = item.Id
	tmpMainCreator.UserId = item.UserId
	itemCreators = append(itemCreators, tmpMainCreator)
	for _, subItem := range rq {
		var tmpCreator marketplacemodel.ItemCreator
		tmpCreator.ItemId = subItem.Id
		tmpCreator.UserId = subItem.UserId
		itemCreators = append(itemCreators, tmpCreator)
	}

	errCreator := i.ItemCreators(itemCreators)
	if errCreator != nil {
		tx.Rollback()
		log.Println(errCreator)
		return 0, errCreator
	}
	tx.Commit()
	return item.Id, nil
}

func (i MarketPlaceStore) IsExistCategory(rq *marketplacemodel.Category) (bool, error) {
	var has bool
	if err := i.db.Table(marketplacemodel.Category{}.TableName()).Select("count(id) > 0").Where(
		"id = ?", rq.Id,
	).Limit(1).Take(&has).Error; err != nil {
		return false, err
	}

	return has, nil
}

func (i MarketPlaceStore) IsExistCollection(rq *marketplacemodel.Collection) (bool, error) {
	var has bool
	if err := i.db.Table(marketplacemodel.Collection{}.TableName()).Select("count(id) > 0").Where(
		"id = ?", rq.Id,
	).Limit(1).Take(&has).Error; err != nil {
		return false, err
	}

	return has, nil
}

func (i MarketPlaceStore) IsLikedItem(itemId int64, userId int64) (bool, error) {
	var has bool
	if err := i.db.Table(marketplacemodel.Like{}.TableName()).Select("count(id) > 0").Where(
		"item_id = ? and user_id = ?", itemId, userId,
	).Where("is_deleted = ?", common.ACTIVE).Limit(1).Take(&has).Error; err != nil {
		return false, err
	}

	return has, nil
}

func (i MarketPlaceStore) BatchInsertItem(items []marketplacemodel.Item) (int64, error) {
	result := i.db.Create(&items)
	return result.RowsAffected, result.Error
}

func (i MarketPlaceStore) GetByUserIdAndNftIds(userId int64, tokenId []string, contractAddress string) (map[string]bool, error) {
	var tokenIds []string
	err := i.db.Model(&marketplacemodel.Item{}).Where("user_id", userId).
		Where("contract_address", contractAddress).
		Where("nft IN ?", tokenId).Select("nft").Find(&tokenIds).Error
	if err != nil {
		return nil, err
	}
	mId := make(map[string]bool)
	for _, id := range tokenIds {
		mId[id] = true
	}
	return mId, nil
}

func (i MarketPlaceStore) BuyItem(sellerId int64, ownerId int64, nft string) error {
	tx := i.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	item := marketplacemodel.Item{
		UserId: ownerId,
		Status: common.ITEM_NEW,
	}
	result := tx.Model(&marketplacemodel.Item{}).Where("user_id", sellerId).Where("nft", nft).Updates(&item)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected != 1 {
		tx.Rollback()
		return common.ErrUpdateOwner
	}
	return nil
}

func (i MarketPlaceStore) UpdateItemStatusByUserIdAndNft(status int, userId int64, nft string) error {
	return i.db.Model(&marketplacemodel.Item{}).Where("user_id", userId).Where("nft", nft).Update("status", status).Error
}

func (i MarketPlaceStore) CancelSellItem(in *marketplacemodel.CancelMarketplace) error {
	marketFilter := marketplacemodel.MarketPlaceFilter{
		IdOnMarket: in.IdOnMarket,
		Status:     common.ITEM_STATUS_ON_MARKET,
	}
	market, errItem := i.GetMarketPlaceByIdOnMarket(&marketFilter)
	if errItem != nil {
		return errItem
	}

	if market.Id == 0 {
		return common.ErrNotFoundItem
	}

	tx := i.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if txx := tx.Model(&marketplacemodel.MarketPlace{}).Where("id = ?", market.Id).Where("status", common.ITEM_STATUS_ON_MARKET).Update("deleted_at", time.Now()); txx.Error != nil {
		tx.Rollback()
		return txx.Error
	}
	if txx := tx.Model(&marketplacemodel.Item{}).Where("id = ?", market.ItemId).Update("status", common.ITEM_NEW); txx.Error != nil {
		tx.Rollback()
		return txx.Error
	}
	return tx.Commit().Error
}

func (i MarketPlaceStore) DetailCollection(collectionId int64) (*marketplacemodel.Collection, error) {
	data := &marketplacemodel.Collection{}
	err := i.db.Table(marketplacemodel.Collection{}.TableName()).Where("id = ?", collectionId).Find(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i MarketPlaceStore) GetCollectionByCondition(rq *marketplacemodel.CanCreateCollectionRequest) (marketplacemodel.Collection, error) {
	data := marketplacemodel.Collection{}
	query := i.db.Table(marketplacemodel.Collection{}.TableName())
	if len(rq.Name) > 0 {
		query.Where("name = ?", rq.Name)
	}
	query.Where("collections.deleted_at IS NULL")
	err := query.Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (i MarketPlaceStore) GetCollectionItems(listCollectionId []int64) ([]marketplacemodel.Item, error) {
	var dataItems []marketplacemodel.Item
	errItem := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("collection_id IN ?", listCollectionId).
		Where("status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Find(&dataItems).Error
	if errItem != nil {
		return nil, errItem
	}
	return dataItems, nil
}

func (i MarketPlaceStore) AddBannerCms(update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.AddBannerCmsRequest{}.TableName()).
		Create(update).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) ListBannerCms() ([]marketplacemodel.AddBannerCmsRequest, error) {
	var data []marketplacemodel.AddBannerCmsRequest
	if err := i.db.Table(marketplacemodel.AddBannerCmsRequest{}.TableName()).
		Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (i MarketPlaceStore) GetListTransaction(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]marketplacemodel.ShortInfoMarketPlace, error) {
	var result []marketplacemodel.ShortInfoMarketPlace
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Select("market_places.*, ms.id as market_place_sold_updated_at, ms.buyer_id as buyer_id, "+
			"ms.updated_at as market_place_sold_updated_at, ms.transaction_id as sold_transaction_id").
		Joins("INNER JOIN market_place_solds as ms ON ms.market_place_id = market_places.id").
		Where("ms.status = ?", common.MARKET_PLACE_SOLD_STATUS_BUYED)
	query.Where("market_places.deleted_at IS NULL")
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return nil, nil
		}

		if paging.Limit == 0 {
			paging.Limit = int(paging.Total)
		}

		if paging.Page == 0 {
			paging.Page = 1
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	updateString := "ms.updated_at "
	if strings.ToLower(rq.SortBy) == "asc" {
		updateString += strings.ToLower(rq.SortBy)
	} else {
		updateString += "desc"
	}

	err := query.Order(updateString).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) TopSeller(in []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TransactionLogTopSeller, error) {
	var result []*marketplacemodel.TransactionLogTopSeller
	query := i.db.Table(marketplacemodel.TransactionLog{}.TableName()).
		Where("created_at BETWEEN ? AND ?", paging.CustomStart, paging.End)

	query.Select("seller_id,price, item_id, created_at")
	query.Where("seller_id IN ?", in)

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListUserTopSeller(paging *common.PagingWithTime) ([]int64, error) {
	var result []int64
	query := i.db.Table(marketplacemodel.TransactionLog{}.TableName()).Select("seller_id").
		Where("created_at BETWEEN ? AND ?", paging.Start, paging.End).Group("seller_id")
	if err := query.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if paging != nil {
		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i MarketPlaceStore) CountItemSell(userId []int64) ([]*marketplacemodel.TransactionLogTopSeller, error) {
	var result []*marketplacemodel.TransactionLogTopSeller
	query := i.db.Table(marketplacemodel.TransactionLog{}.TableName()).
		Group("seller_id").
		Where("seller_id IN ?", userId)
	query.Select("seller_id, COUNT(item_id) as total_nft")

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) TopCreator(ids []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TopCreatorResponse, error) {
	var result []*marketplacemodel.TopCreatorResponse
	query := i.db.Table(marketplacemodel.ItemCreator{}.TableName()).
		Where("created_at BETWEEN ? AND ?", paging.CustomStart, paging.End).
		Where("user_id IN ?", ids)

	if err := query.Select("user_id, created_at").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) TopCreatorUserId(paging *common.PagingWithTime) ([]int64, error) {
	var result []int64
	query := i.db.Table(marketplacemodel.ItemCreator{}.TableName()).Select("user_id").
		Group("user_id").Where("created_at BETWEEN ? AND ?", paging.Start, paging.End)

	if err := query.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, err
}

func (i MarketPlaceStore) TopVolume(ids []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TopVolume, error) {
	var result []*marketplacemodel.TopVolume
	query := i.db.Table(marketplacemodel.TransactionLog{}.TableName()).
		Joins("LEFT JOIN items ON transaction_logs.item_id = items.id").
		Joins("INNER JOIN collections ON collections.id = items.collection_id").
		Where("transaction_logs.created_at BETWEEN ? AND ?", paging.CustomStart, paging.End).
		Where("collections.id IN ?", ids)
	query.Where("transaction_logs.deleted_at IS NULL")

	if err := query.Select("collections.*, transaction_logs.price as sold_price, transaction_logs.created_at as trans_created").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) TopVolumeCollectionId(paging *common.PagingWithTime) ([]int64, error) {
	var result []int64
	query := i.db.Table(marketplacemodel.TransactionLog{}.TableName()).
		Joins("LEFT JOIN items ON transaction_logs.item_id = items.id").
		Joins("INNER JOIN collections ON collections.id = items.collection_id").
		Where("transaction_logs.created_at BETWEEN ? AND ?", paging.Start, paging.End).
		Group("collections.id")
	query.Where("transaction_logs.deleted_at IS NULL")
	if err := query.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	err := query.Select("collections.id").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (i MarketPlaceStore) MapFloorPrice(collectionIds []int64) ([]*marketplacemodel.MapFloorPrice, error) {
	var result []*marketplacemodel.MapFloorPrice
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Joins("INNER JOIN items ON market_places.item_id = items.id").
		Joins("RIGHT JOIN collections ON collections.id = items.collection_id").
		Where("collections.id IN ?", collectionIds).
		Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SELLING).
		Group("collections.id")
	query.Where("market_places.deleted_at IS NULL")
	if err := query.Select("collections.*, " + "MIN(market_places.price * 1) as floor_price").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListLikedItem(itemId []int64, userId int64) ([]marketplacemodel.Like, error) {
	var result []marketplacemodel.Like
	if err := i.db.Table(marketplacemodel.Like{}.TableName()).
		Where("item_id IN ? and user_id = ?", itemId, userId).
		Where("is_deleted = ?", common.ACTIVE).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListLikes(itemId []int64) ([]marketplacemodel.Like, error) {
	var result []marketplacemodel.Like
	if err := i.db.Table(marketplacemodel.Like{}.TableName()).
		Where("item_id IN ?", itemId).
		Where("is_deleted = ?", common.ACTIVE).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) UpdateCollection(updator, selector *marketplacemodel.Collection) error {
	err := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Where(&selector).Updates(&updator).Error
	if err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) FindCollection(selector *marketplacemodel.Collection) (*marketplacemodel.Collection, error) {
	var result *marketplacemodel.Collection
	err := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Where(&selector).Take(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) CollectionItemProperties(collectionId int64) ([]*marketplacemodel.Item, error) {
	var result []*marketplacemodel.Item
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Select("properties, id").
		Where("collection_id = ?", collectionId).
		Where("status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Where("properties IS NOT NULL AND properties <> ?", "").
		Where("properties <> ?", "null").
		Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) UpdateItem(updator, selector *marketplacemodel.Item) error {
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where(&selector).Updates(&updator).Error
	if err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) FindItem(selector *marketplacemodel.Item) (*marketplacemodel.Item, error) {
	var result *marketplacemodel.Item
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where(&selector).Take(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) InsertLikeCollection(in *marketplacemodel.CollectionLike) error {
	var out *marketplacemodel.CollectionLike
	err := m.db.Table(marketplacemodel.CollectionLike{}.TableName()).
		Where("collection_id = ?", in.CollectionId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return err.Error
	}

	if out.Id != 0 {
		if out.IsDeleted == common.ACTIVE {
			return common.ErrUserLiked
		}
		in.Id = out.Id
	}

	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	result := tx.Table(marketplacemodel.CollectionLike{}.TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_deleted"}),
	}).Create(in)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error

}

func (m MarketPlaceStore) DisLikeCollection(in *marketplacemodel.CollectionLike) error {
	var out *marketplacemodel.CollectionLike
	err := m.db.Table(marketplacemodel.CollectionLike{}.TableName()).
		Where("collection_id = ?", in.CollectionId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return common.ErrUserNotLiked
		}
		return err.Error
	}

	if out.IsDeleted == common.DELETED {
		return common.ErrUserDisLiked
	}

	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if out.IsDeleted == common.ACTIVE {
		if err := tx.Table(marketplacemodel.CollectionLike{}.TableName()).
			Where("id = ?", out.Id).
			Updates(map[string]interface{}{
				"is_deleted": common.DELETED,
			}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error

}

func (i MarketPlaceStore) CountCollectionLike(id []int64) ([]*marketplacemodel.CollectionLikeResponse, error) {
	var result []*marketplacemodel.CollectionLikeResponse
	query := i.db.Table(marketplacemodel.CollectionLike{}.TableName()).
		Select("id, collection_id, COUNT(id) as total_like").
		Group("id").Where("collection_id IN ?", id).Where("is_deleted = ?", common.ACTIVE)

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListCollectionLike(filter marketplacemodel.CollectionLikeFilter) ([]*marketplacemodel.CollectionLike, error) {
	var result []*marketplacemodel.CollectionLike
	//query := i.db.Table(marketplacemodel.CollectionLike{}.TableName()).
	//	Select("id, collection_id, user_id").
	//	Where("collection_id IN ?", id).Where("is_deleted = ?", common.ACTIVE)
	query := i.buildCollectionLikeCommonQuery(filter)

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) buildCollectionLikeCommonQuery(filter marketplacemodel.CollectionLikeFilter) *gorm.DB {
	query := i.db.Table(marketplacemodel.CollectionLike{}.TableName())
	query.Where("collection_likes.deleted_at IS NULL")
	if filter.Id > 0 {
		query.Where("id = ?", filter.Id)
	}
	if len(filter.Ids) > 0 {
		query.Where("id IN ?", filter.Ids)
	}
	if filter.UserId > 0 {
		query.Where("user_id = ?", filter.UserId)
	}
	if len(filter.Cols) > 0 {
		query.Select(filter.Cols)
	}
	if len(filter.CollectionIds) > 0 {
		query.Where("collection_id IN ?", filter.CollectionIds)
	}
	if filter.CollectionId > 0 {
		query.Where("collection_id = ?", filter.CollectionId)
	}
	return query.Where("is_deleted = ?", filter.IsDeleted).Order("id DESC")
}

func (i MarketPlaceStore) BatchUpdateCollectionItem(ids []int64, updater map[string]interface{}) error {
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id IN ?", ids).
		Updates(updater).Error
	if err != nil {
		return err
	}

	return nil
}

func (i MarketPlaceStore) CountLikeById(itemId int64) (int64, error) {
	var result int64
	if err := i.db.Table(marketplacemodel.Like{}.TableName()).Select("COUNT(*)").
		Where("item_id = ?", itemId).
		Where("is_deleted = ?", common.ACTIVE).Find(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) IsLikeItem(currentUserId, itemId int64) (int64, error) {
	var result int64
	if err := i.db.Table(marketplacemodel.Like{}.TableName()).
		Where("item_id = ?", itemId).
		Where("user_id = ?", currentUserId).
		Where("is_deleted = ?", common.ACTIVE).Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (m MarketPlaceStore) GetMarketPlaceByIdOnMarket(marketPlace *marketplacemodel.MarketPlaceFilter) (*marketplacemodel.MarketPlace, error) {
	data := &marketplacemodel.MarketPlace{}
	tx := m.buildMarketCommonQuery(*marketPlace)
	if marketPlace.Unscoped {
		tx.Unscoped()
	}
	err := tx.Joins("Item").Joins("Seller").Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i MarketPlaceStore) UpdateViewItem(itemId int64) error {
	_, err := i.GetItemByID(marketplacemodel.ItemFilter{Id: itemId})
	if err != nil {
		return err
	}

	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", itemId).
		Updates(map[string]interface{}{
			"view": gorm.Expr("view + ?", 1),
		}).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateViewCollection(id int64) error {
	if err := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Where("id = ?", id).
		UpdateColumn("view", gorm.Expr("view + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (m MarketPlaceStore) UpdateOwnerOfItem(request marketplacemodel.UpdateItemOwnerRequest) error {
	result := m.db.Table(marketplacemodel.Item{}.TableName()).Where("user_id = ? AND nft = ? AND BINARY contract_address = ?", request.SenderId, request.Nft, request.ContractAddress)
	if request.ReceiverId > 0 {
		result.Update("user_id", request.ReceiverId)
	} else {
		result.Update("status", common.ITEM_TRANSFER_RECEIVER_ID_NOT_FOUND)
	}
	result.Limit(1)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("Update owner of nft")
	}
	return nil
}

func (i MarketPlaceStore) GetSetting() (*marketplacemodel.Setting, error) {
	var result *marketplacemodel.Setting
	err := i.db.Table(marketplacemodel.Setting{}.TableName()).Limit(1).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m MarketPlaceStore) UpdateItemJoinEvent(request marketplacemodel.UpdateItemEvent) error {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	result := tx.Table(marketplacemodel.Item{}.TableName()).Where("user_id = ? AND nft = ? AND BINARY contract_address = ?", request.UserId, request.Nft, request.ContractAddress).
		Update("is_join_event", common.ITEM_ON_JOIN_EVENT).
		Update("time_join_event", time.Now().Unix()).
		Limit(1)
	if result.Error != nil || result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Update owner of nft")
	}
	var itemEventHistory marketplacemodel.ItemJoinEventHistory
	itemEventHistory.ItemId = request.ItemId
	itemEventHistory.UserJoinEventId = request.UserId
	itemEventHistory.Status = common.ITEM_JOIN_EVENT_STATUS_ACTIVE
	itemEventHistory.Profit = 0.2
	itemEventHistory.Revenue = request.Revenue
	if err := tx.Table(marketplacemodel.ItemJoinEventHistory{}.TableName()).Create(&itemEventHistory).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (m MarketPlaceStore) UpdateItemCancelEvent(request marketplacemodel.UpdateItemEvent) error {
	result := m.db.Table(marketplacemodel.Item{}.TableName()).Where("user_id = ? AND nft = ? AND BINARY contract_address = ?", request.UserId, request.Nft, request.ContractAddress).
		Update("is_join_event", common.ITEM_NOT_JOIN_EVENT).
		Limit(1)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("Update owner of nft")
	}
	return nil
}

func (m MarketPlaceStore) GetListItemByEvent(rq marketplacemodel.ItemEventFilter, paging *common.Paging) ([]*marketplacemodel.ListItemEventResult, error) {
	var listItem []*marketplacemodel.ListItemEventResult
	result := m.BuildCountItemQuery()
	if rq.UserId > 0 {
		result.Where("items.user_id = ?", rq.UserId)
	}
	if rq.CategoryId > 0 {
		result.Where("items.category_id = ?", rq.CategoryId)
	}

	if len(rq.Search) > 0 {
		result.Where("items.title LIKE ?", "%"+rq.Search+"%")
	}

	result.Group("items.id")
	if err := result.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if len(rq.SortBy) > 0 {
		result.Order(rq.SortBy)
	}
	if paging.Limit > 0 {
		result.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	findError := result.Order("id DESC").Find(&listItem).Error
	if findError != nil {
		return nil, errors.New("Record not found")
	}
	return listItem, nil
}

func (m MarketPlaceStore) CountListItemByEvent() (int64, error) {
	var total int64
	result := m.BuildCountItemQuery()
	result.Group("items.id")
	err := result.Count(&total).Error
	return total, err
}

func (m MarketPlaceStore) BuildCountItemQuery() *gorm.DB {
	result := m.db.Table(marketplacemodel.Item{}.TableName()).
		Select("items.*, users.avatar as user_avatar, users.display_name as user_name, users.wallet_address as user_wallet_address, "+
			"count(CASE WHEN likes.is_deleted =0 THEN 1 END) as total_like, collections.`name` as collection_name").
		Joins("INNER JOIN users ON items.user_id = users.id").
		Joins("LEFT JOIN likes ON items.id = likes.item_id").
		Joins("LEFT JOIN collections ON items.collection_id = collections.id").
		Where("items.is_join_event = ?", common.ITEM_ON_JOIN_EVENT).
		Where("items.is_hide = ?", common.ITEM_UNHIDE).
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Where("items.deleted_at IS NULL")
	return result
}

func (m MarketPlaceStore) CountItemEventByCate(cateId int64) (int64, error) {
	var total int64
	err := m.db.Table(marketplacemodel.Item{}.TableName()).
		Joins("INNER JOIN categories ON items.category_id = categories.id").
		Where("categories.id = ?", cateId).
		Where("items.is_join_event = ?", common.ITEM_ON_JOIN_EVENT).
		Where("items.deleted_at IS NULL").
		Count(&total).Error
	return total, err
}

func (i MarketPlaceStore) CheckItemCanJoinEvent(contractAddress string, nft string, userId int) (bool, error) {
	var result int64
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("user_id = ?", userId).
		Where("binary contract_address = ?", contractAddress).
		Where("nft = ?", nft).
		Where("status IN ?", []int{common.ITEM_NEW, common.ITEM_READY_ON_SELL}).
		Where("is_join_event = ?", common.ITEM_NOT_JOIN_EVENT).Count(&result).Error; err != nil {
		return false, err
	}
	return result > 0, nil
}

func (i MarketPlaceStore) CheckItemJoinEvent(nft string) (bool, error) {
	var result int64
	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("nft = ?", nft).
		Where("is_join_event = ?", common.ITEM_ON_JOIN_EVENT).Count(&result).Error; err != nil {
		return false, err
	}
	return result > 0, nil
}

func (i MarketPlaceStore) AddWsBlock(block marketplacemodel.BlockChainBlock) error {
	var total int64
	errCount := i.db.Table(marketplacemodel.BlockChainBlock{}.TableName()).Where("block_number = ?", block.BlockNumber).Count(&total).Error
	if errCount != nil {
		return errCount
	}
	if total > 0 {
		return nil
	}
	err := i.db.Table(marketplacemodel.BlockChainBlock{}.TableName()).Create(&block).Error
	return err
}
func (i MarketPlaceStore) GetLatestBlock() (marketplacemodel.BlockChainBlock, error) {
	var block marketplacemodel.BlockChainBlock
	err := i.db.Table(marketplacemodel.BlockChainBlock{}.TableName()).
		Where("is_block_process = ?", common.IS_BLOCK_PROCESS).
		Where("status = ?", common.BLOCK_STATUS_SUCCESS).Order("block_number desc").Limit(1).Find(&block).Error
	if err != nil {
		return block, err
	}
	return block, nil
}

func (i MarketPlaceStore) CountMarketPlace(rq marketplacemodel.MarketPlaceFilter) (int64, error) {
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName())
	if rq.IdOnMarket > 0 {
		query.Where("id_on_market = ?", rq.IdOnMarket)
	}
	if rq.Status > 0 {
		query.Where("status = ?", rq.Status)
	}
	if len(rq.MarketContract) > 0 {
		query.Where("contract_address = ?", rq.MarketContract)
	}
	if rq.SellType > 0 {
		query.Where("sell_type = ?", rq.SellType)
	}
	query.Where("deleted_at IS NULL")
	var total int64
	err := query.Count(&total).Error

	if err != nil {
		return total, err
	}
	return total, nil
}

func (i MarketPlaceStore) InsertBlock(block []marketplacemodel.BlockChainBlock) error {
	err := i.db.Create(&block)
	return err.Error
}

func (i MarketPlaceStore) GetWsBlock() ([]marketplacemodel.BlockChainBlock, error) {
	var block []marketplacemodel.BlockChainBlock
	err := i.db.Table(marketplacemodel.BlockChainBlock{}.TableName()).
		Where("is_block_process = ?", common.IS_NOT_BLOCK_PROCESS).
		Where("status = ?", common.BLOCK_STATUS_WS_SUCCESS).Order("block_number desc").Find(&block).Error
	if err != nil {
		return block, err
	}
	return block, nil
}

func (i MarketPlaceStore) UpdateWsBlockSuccess(block []marketplacemodel.BlockChainBlock, listIds []uint64) error {
	err := i.db.Table(marketplacemodel.BlockChainBlock{}.TableName()).
		Where("id IN ?", listIds).
		Updates(block).Error
	return err
}

func (i MarketPlaceStore) CountItemOnSaleByUserid(user usermodel.UserDetail) (int64, error) {
	var result int64
	query := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("user_id = ?", user.Id).
		Where("status = ?", common.ITEM_STATUS_ON_MARKET)
	if user.Type < 6 {
		query.Where("is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("deleted_at IS NULL")
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CountCollectionByUserid(user usermodel.UserDetail) (int64, error) {
	var result sql.NullInt64
	query := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Select("COUNT(id)").
		Where("user_id = ?", user.Id).
		Where("deleted_at IS NULL")

	if err := query.Take(&result).Error; err != nil {
		return 0, err
	}
	return result.Int64, nil
}

func (i MarketPlaceStore) CountNFTByUserid(userId int64, status []int) (int64, error) {
	var result int64
	query := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("user_id = ?", userId).
		Where("status IN ?", status).
		Where("deleted_at IS NULL")

	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CountOwnerNFTByUserId(user usermodel.UserDetail, status []int) (int64, error) {
	var result int64
	query := i.db.Table(marketplacemodel.Item{}.TableName()).Joins("INNER JOIN item_creators ON items.id = item_creators.item_id")
	query.Where("items.user_id = ?", user.Id).Where("items.status IN ?", status).
		Where("(items.main_item_id IS NULL OR items.main_item_id = 0)")
	if user.Type < 6 {
		query.Where("items.is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("items.deleted_at IS NULL")
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CountNFTCreatedByUserid(user usermodel.UserDetail) (int64, error) {
	var result int64
	query := i.db.Table(marketplacemodel.ItemCreator{}.TableName()).
		Where("item_creators.user_id = ?", user.Id).
		Joins("INNER JOIN items ON items.id = item_creators.item_id").
		Where("(items.main_item_id IS NULL OR items.main_item_id = 0)").
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC)
	if user.Type < 6 {
		query.Where("items.is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("items.deleted_at IS NULL")
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CountLikedNFTByUserid(user usermodel.UserDetail) (int64, error) {
	var result sql.NullInt64
	query := i.db.Table(marketplacemodel.Like{}.TableName()).
		Select("COUNT(likes.id)").
		Joins("INNER JOIN items ON items.id = likes.item_id").
		Where("(items.main_item_id IS NULL OR items.main_item_id = 0)").
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Where("likes.user_id = ?", user.Id).
		Where("likes.is_deleted = ?", common.ACTIVE)

	if user.Type < 6 {
		query.Where("items.is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("items.deleted_at IS NULL")
	if err := query.Take(&result).Error; err != nil {
		return 0, err
	}
	return result.Int64, nil
}

func (i MarketPlaceStore) CountNFTEventByUserid(user usermodel.UserDetail) (int64, error) {
	var result sql.NullInt64
	query := i.db.Table(marketplacemodel.Item{}.TableName()).
		Select("COUNT(id)").
		Where("user_id = ?", user.Id).
		Where("is_join_event = ?", common.ITEM_ON_JOIN_EVENT)

	if user.Type < 6 {
		query.Where("items.is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("items.deleted_at IS NULL")
	if err := query.Take(&result).Error; err != nil {
		return 0, err
	}
	return result.Int64, nil
}

func (i MarketPlaceStore) CountNFTSoldByUserid(user usermodel.UserDetail) (int64, error) {
	var result int64
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Joins("INNER JOIN items ON items.id = market_places.item_id").
		Joins("INNER JOIN item_creators ON items.id = item_creators.item_id").
		Where("seller_id = ?", user.Id).
		Where("(items.main_item_id IS NULL OR items.main_item_id = 0)").
		Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SOLDED).Group("market_places.item_id")

	if user.Type < 6 {
		query.Where("items.is_hide = ?", common.ITEM_UNHIDE)
	}
	query.Where("items.deleted_at IS NULL")
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CollectionStats(itemIds []int64, lastTime int64) ([]marketplacemodel.MarketPlace, error) {
	var result []marketplacemodel.MarketPlace
	now := time.Now()
	last := now.AddDate(0, 0, int(lastTime*-1))

	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Where("item_id IN ?", itemIds).
		Where("updated_at >= ?", last)

	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) UpdateCollectionCms(id int64, update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.Collection{}.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) HideNftCms(id int64) error {
	var exist bool
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", id).
		Where("is_hide = ?", common.ITEM_HIDE).Select("COUNT(id) > 0").Take(&exist).Error
	if err != nil {
		return err
	}

	if exist {
		return common.ErrNftHided
	}

	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_hide": common.ITEM_HIDE,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UnHideNftCms(id int64) error {
	var exist bool
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", id).
		Where("is_hide = ?", common.ITEM_UNHIDE).Select("COUNT(id) > 0").Take(&exist).Error
	if err != nil {
		return err
	}

	if exist {
		return common.ErrNftUnHided
	}

	if err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_hide": common.ITEM_UNHIDE,
		}).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) NFTStatistic() (int64, error) {
	var total int64
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("is_hide = ?", common.ITEM_UNHIDE).
		Where("status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Where("is_join_event = ?", common.ITEM_ON_JOIN_EVENT).
		Select("COUNT(id)").
		Take(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (i MarketPlaceStore) CountTotalNft() (int64, error) {
	var total int64
	err := i.db.Table(marketplacemodel.Item{}.TableName()).
		Where("is_hide = ?", common.ITEM_UNHIDE).
		Where("status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Select("COUNT(id)").
		Take(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (i MarketPlaceStore) ExploreStatistic(total *int64) ([]*marketplacemodel.CategoryStatistic, error) {
	var result []*marketplacemodel.CategoryStatistic

	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Joins("INNER JOIN items ON market_places.item_id = items.id").
		Joins("INNER JOIN categories ON items.category_id = categories.id").
		Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SELLING).
		Where("market_places.deleted_at IS NULL").
		Group("categories.id")

	subQuery := i.db.Table(marketplacemodel.Item{}.TableName()).
		Joins("INNER JOIN collections as c ON c.id = items.collection_id").
		Where("c.is_hold = ?", 1).Select("items.id").Group("id")
	query.Where("NOT market_places.item_id IN (?) ", subQuery)

	if err := query.Select("categories.id, categories.name, COUNT(categories.id) as total").Find(&result).Error; err != nil {
		return nil, err
	}

	queryTotal := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Joins("INNER JOIN items ON market_places.item_id = items.id").
		Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SELLING).
		Where("market_places.deleted_at IS NULL").
		Where("items.category_id <> 0").
		Where("NOT market_places.item_id IN (?) ", subQuery)

	if err := queryTotal.Count(total).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (i MarketPlaceStore) GetItemJoinEventHistory(filter usermodel.RefNetworkFilter) ([]marketplacemodel.ListRefInfoResult, error) {
	query := i.db.Table(marketplacemodel.ItemJoinEventHistory{}.TableName()).
		Select("user_join_event_id as user_id, item_id, profit, revenue as price, created_at").
		Where("user_join_event_id IN ?", filter.UserIds).Where("status = ?", common.ITEM_JOIN_EVENT_STATUS)
	if filter.StartTimeStamp > 0 {
		query.Where("created_at >= ?", filter.StartTimeStamp)
	}
	var result []marketplacemodel.ListRefInfoResult
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) CheckOldItemEvent(itemId int64) (bool, error) {
	query := i.db.Table(marketplacemodel.ItemJoinEventHistory{}.TableName()).
		Where("item_id = ?", itemId).Where("status = ?", common.ITEM_JOIN_EVENT_STATUS)
	var result int64
	if err := query.Count(&result).Error; err != nil {
		return false, err
	}
	if result > 0 {
		return false, nil
	}
	return true, nil
}

func (i MarketPlaceStore) GetRefMarketPlace(filter usermodel.RefNetworkFilter) ([]marketplacemodel.ListRefInfoResult, error) {
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Select("seller_id as user_id, item_id, price, created_at").
		Where("seller_id IN ?", filter.UserIds).Where("status = ?", common.MARKET_PLACE_STATUS_SOLDED)
	if filter.StartTimeStamp > 0 {
		query.Where("created_at >= ?", filter.StartTimeStamp)
	}
	var result []marketplacemodel.ListRefInfoResult
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) CountAvailableItemCopied(filter marketplacemodel.CountItemFilter) (int64, error) {
	query := i.db.Table(marketplacemodel.Item{}.TableName())
	if filter.Id > 0 {
		query.Where("items.id = ?", filter.Id)
	}
	if filter.MainItemId > 0 {
		query.Where("items.main_item_id = ?", filter.MainItemId)
	}
	if filter.Status > 0 {
		query.Where("items.status = ?", filter.Status)
	}
	if len(filter.Stts) > 0 {
		query.Where("items.status IN (?)", filter.Stts)
	}
	if filter.CheckMarket > 0 {
		query.Joins("INNER JOIN market_places ON items.id = market_places.item_id")
		query.Where("market_places.status = ?", common.MARKET_PLACE_STATUS_SELLING)
		if filter.MarketUserId > 0 {
			query.Where("market_places.user_id = ?", filter.MarketUserId)
		}
	}
	var result int64
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i MarketPlaceStore) CollectionExplore(filter marketplacemodel.ExploreCollectionFilter, paging *common.Paging) ([]*marketplacemodel.CollectionExploreResponse, error) {
	var result []*marketplacemodel.CollectionExploreResponse
	tx := i.db.Table(marketplacemodel.Collection{}.TableName())
	tx.Where("collections.deleted_at IS NULL")
	if len(filter.OrderBy) > 0 {
		tx.Order(filter.OrderBy)
	}

	if filter.TotalItemMin > 0 {
		tx.Having("total_item >= ?", filter.TotalItemMin)
	}

	if filter.TotalItemMax > 0 {
		tx.Having("total_item <= ?", filter.TotalItemMax)
	}

	if filter.FloorPriceMin > 0 {
		tx.Having("floor_price >= ?", filter.FloorPriceMin)
	}

	if filter.FloorPriceMax > 0 {
		tx.Having("floor_price <= ?", filter.FloorPriceMax)
	}

	if len(filter.Search) > 0 {
		tx.Where("collections.name LIKE ?", "%"+filter.Search+"%")
	}

	stts := [3]int{common.ITEM_ON_SALE, common.ITEM_READY_ON_SELL, common.ITEM_NEW}
	tx.Joins("LEFT JOIN items ON items.collection_id = collections.id").
		Joins("LEFT JOIN market_places ON items.id = market_places.item_id").
		Select("collections.*, "+
			"COUNT(DISTINCT CASE WHEN items.status IN ? AND (items.main_item_id IS NULL OR items.main_item_id = 0) THEN items.id END) as total_item, "+
			"COUNT(DISTINCT CASE WHEN items.status IN ? AND (items.main_item_id IS NULL OR items.main_item_id = 0) THEN items.user_id END) as total_owner, "+
			"MIN(CASE WHEN market_places.status = ? THEN market_places.price * 1 / ? END) as floor_price, "+
			"SUM(CASE WHEN market_places.status = ? THEN market_places.price * 1 / ? END) as volume",
			stts,
			stts,
			common.MARKET_PLACE_STATUS_SELLING,
			1000000000000000000,
			common.MARKET_PLACE_STATUS_SOLDED,
			1000000000000000000,
		).Group("collections.id")

	if paging != nil {
		errCount := i.db.Raw("SELECT COUNT(temp_result.id) FROM (?) temp_result", tx).Scan(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return result, nil
		}
	}

	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListBrands(rq *marketplacemodel.BrandFilter, paging *common.Paging) ([]*marketplacemodel.Brand, error) {
	var result []*marketplacemodel.Brand
	subQuery := fmt.Sprintf("(select count(brand_image.id) from brand_image where brand_id = brands.id)")
	query := i.db.Table(marketplacemodel.Brand{}.TableName())
	if rq.IsHot >= 0 {
		query.Where("is_hot = ?", rq.IsHot)
	}
	if rq.Id > 0 {
		query.Where("id = ? ", rq.Id)
	}
	query.Where("deleted_at IS NULL")
	query.Order("display_order ASC")
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return nil, nil
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	if err := query.Select(fmt.Sprintf("*, %v AS count_items", subQuery)).Order("count_items desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) GetBrandImage(ListBrandsImage []int64) ([]marketplacemodel.BrandImage, error) {
	var dataItems []marketplacemodel.BrandImage
	errItem := i.db.Table(marketplacemodel.BrandImage{}.TableName()).
		Where("brand_id IN ?", ListBrandsImage).
		Where("deleted_at IS NULL").
		Order("display_order ASC").
		Find(&dataItems).Error
	if errItem != nil {
		return nil, errItem
	}
	return dataItems, nil
}

func (i MarketPlaceStore) UpdateBrandCms(id int64, update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.Brand{}.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateBrandImageCms(id int64, update map[string]interface{}) error {
	if err := i.db.Table(marketplacemodel.BrandImage{}.TableName()).
		Where("id = ?", id).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (m MarketPlaceStore) CreateBrandCms(brand marketplacemodel.Brand, brandImages []marketplacemodel.BrandImage) error {
	tx := m.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if err := tx.Table(marketplacemodel.Brand{}.TableName()).Create(&brand).Error; err != nil {
		tx.Rollback()
		return err
	}

	for index, _ := range brandImages {
		brandImages[index].BrandId = brand.Id
	}
	if err := tx.Table(marketplacemodel.BrandImage{}.TableName()).Create(&brandImages).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (i MarketPlaceStore) CreateBrandImageCms(create *marketplacemodel.BrandImageDetail) error {
	return i.db.Table(marketplacemodel.BrandImage{}.TableName()).Create(create).Error
}

func (i MarketPlaceStore) DeleteBrandImage(id int64) error {
	query := i.db.Table(marketplacemodel.BrandImage{}.TableName()).Where("id = ?", id)
	if err := query.Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) DeleteBrand(id int64) error {
	query := i.db.Table(marketplacemodel.Brand{}.TableName()).Where("id = ?", id)
	if err := query.Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) ExploreNft(filter marketplacemodel.ExploreNftFilter, paging *common.Paging) ([]*marketplacemodel.ExploreNftQueryResponse, error) {
	var result []*marketplacemodel.ExploreNftQueryResponse
	tx := i.db.Table(marketplacemodel.Item{}.TableName()).
		Joins("LEFT JOIN market_places ON items.id = market_places.item_id").
		Joins("LEFT JOIN likes ON items.id = likes.item_id").
		Select("items.*, "+
			"COUNT(CASE WHEN likes.is_deleted = ? AND likes.user_id = ? THEN 1 END) > 0  as is_like, "+
			"SUM(CASE WHEN market_places.status = ? THEN market_places.price * 1.0 / ? END) as price",
			common.ACTIVE,
			filter.UserId,
			common.MARKET_PLACE_STATUS_SELLING,
			1000000000000000000.0).
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Group("items.id")

	if filter.PriceMax > 0 {
		tx.Having("price <= ?", filter.PriceMax)
	}

	if filter.PriceMin > 0 {
		tx.Having("price >= ?", filter.PriceMin)
	}

	if filter.CollectionId > 0 {
		tx.Where("items.collection_id = ?", filter.CollectionId)
	}

	if filter.CategoryId > 0 {
		tx.Where("items.category_id = ?", filter.CategoryId)
	}

	if len(filter.OrderBy) > 0 {
		tx.Order(filter.OrderBy)
	}

	// properties
	//?properties="Size:M","Size:L"
	if len(filter.Properties) > 0 {
		tx.Where("properties IS NOT NULL AND properties <> ?", "").Where("properties <> ?", "null")
		tx.Having(utils.BuildFilterJson(filter.Properties, "properties"))
	}

	if len(filter.Search) > 0 {
		tx.Where("items.title LIKE ?", "%"+filter.Search+"%")
	} else {
		tx.Where("market_places.id IS NOT NULL")
		tx.Where("market_places.status <> ?", common.MARKET_PLACE_STATUS_SOLDED)
	}

	if paging != nil {
		errCount := i.db.Raw("SELECT COUNT(temp_result.id) FROM (?) temp_result", tx).Scan(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) SaveAuctionBid(bid marketplacemodel.AuctionBid) error {
	query := i.db.Table(marketplacemodel.AuctionBid{}.TableName())
	if err := query.Create(&bid).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) GetAuctionBid(filter marketplacemodel.GetAuctionFilter, paging *common.Paging) ([]marketplacemodel.AuctionBid, error) {
	query := i.db.Table(marketplacemodel.AuctionBid{}.TableName())
	if filter.IsGetUser > 0 {
		query.Joins("UserInfo")
	}
	if filter.ItemId > 0 {
		query.Joins("INNER JOIN market_places ON market_places.id = auction_bids.market_place_id")
		query.Where("market_places.item_id = ?", filter.ItemId)
	} else if len(filter.ItemIds) > 0 {
		query.Joins("INNER JOIN market_places ON market_places.id = auction_bids.market_place_id")
		query.Where("market_places.item_id IN ?", filter.ItemIds)
	}
	if filter.Id > 0 {
		query.Where("auction_bids.id = ?", filter.Id)
	}
	if filter.MarketPlaceId > 0 {
		query.Where("market_place_id = ?", filter.MarketPlaceId)
	}
	if len(filter.MarketPlaceIds) > 0 {
		query.Where("market_place_id IN (?)", filter.MarketPlaceIds)
	}
	if len(filter.OrderBy) > 0 {
		query.Order(filter.OrderBy)
	}
	if len(filter.GroupBy) > 0 {
		query.Group(filter.GroupBy)
	}
	var result []marketplacemodel.AuctionBid
	if paging != nil {
		errCount := query.Where("auction_bids.deleted_at IS NULL").Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ListReport(paging *common.Paging) ([]*marketplacemodel.ReportTableDetail, error) {
	query := i.db.Table(marketplacemodel.ReportTableDetail{}.TableName())
	var result []*marketplacemodel.ReportTableDetail
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return nil, nil
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i MarketPlaceStore) ChangeReport(rq *marketplacemodel.ChangeReportRequest) error {
	query := i.db.Table(marketplacemodel.ReportTableDetail{}.TableName())
	if rq.Id > 0 {
		query.Where("table_reports.id = ? ", rq.Id)
		if rq.IsReject > 0 {
			query.Update("table_reports.is_reject", rq.IsReject)
		}
		if rq.IsHide > 0 {
			query.Update("table_reports.is_hide", rq.IsHide)
		}
		if len(rq.RejectMessage) > 0 {
			query.Update("table_reports.reject_message", rq.RejectMessage)
		}
		if len(rq.HideMessage) > 0 {
			query.Update("table_reports.reject_message", rq.HideMessage)
		}
	}
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) RequestReport(rq *marketplacemodel.ReportTableDetail) error {
	query := i.db.Table(marketplacemodel.ReportTableDetail{}.TableName())
	if err := query.Create(rq).Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) UpdateMarketPlacePack(idOnMarket []int64, transaction string, sellType int, packId int64) error {
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).Where("id_on_market IN ?", idOnMarket).
		Where("sell_type = ?", sellType).Where("transaction_id = ?", transaction).
		Where("status = ?", common.MARKET_PLACE_STATUS_SELLING).Where("deleted_at IS NULL").
		Update("pack_id", packId)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (i MarketPlaceStore) CountMarketPlacePack(idOnMarket []int64, transaction string, sellType int) (int64, error) {
	query := i.db.Table(marketplacemodel.MarketPlace{}.TableName()).Where("id_on_market IN ?", idOnMarket).
		Where("sell_type = ?", sellType).Where("transaction_id = ?", transaction).
		Where("status = ?", common.MARKET_PLACE_STATUS_SELLING).Where("deleted_at IS NULL")
	var result int64
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}
