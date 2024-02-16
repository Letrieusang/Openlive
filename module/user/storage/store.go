package userstorage

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
	"database/sql"
	"fmt"
	"gorm.io/gorm/clause"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (s UserStore) CreateUser(user *usermodel.User) error {
	return s.db.Table(usermodel.User{}.TableName()).Create(&user).Error
}

func (s UserStore) CreateRefUser(user *usermodel.RefUser) error {
	return s.db.Table(usermodel.RefUser{}.TableName()).Create(&user).Error
}

func (s UserStore) UpdateUser(user *usermodel.User) error {
	return s.db.Table(usermodel.User{}.TableName()).Where("id = ?", user.Id).Updates(&user).Error
}

func (s UserStore) UpdateUserOp(userId, op, opClaimed int64, opHash string) error {
	return s.db.Table(usermodel.User{}.TableName()).Where("id = ?", userId).
		Update("op", op).Update("op_backup", opHash).
		Update("op_claimed", opClaimed).Error
}

func (s UserStore) GetUserByCondition(filter usermodel.UserFilter) (usermodel.User, error) {
	userInfo := usermodel.User{}
	tx := s.buildUserCommonQuery(filter).Find(&userInfo)
	if tx.Error != nil {
		return userInfo, tx.Error
	}
	return userInfo, nil
}

func (s UserStore) CheckExistUser(rq *usermodel.CreateUserRequest) (bool, error) {
	var has bool
	if err := s.db.Table(usermodel.User{}.TableName()).Select("count(id) > 0").Where(
		"BINARY wallet_address = ?", rq.WalletAddress,
	).Limit(1).Take(&has).Error; err != nil {
		return false, err
	}
	return has, nil
}

func (s UserStore) CheckCanAddRefUser(rq *usermodel.CreateUserRequest) (bool, error) {
	var has bool
	query := s.db.Table(usermodel.RefUser{}.TableName()).Select("count(id) > 0").
		Where("user_id = ?", rq.UserId)
	if err := query.Limit(1).Take(&has).Error; err != nil {
		return false, err
	}
	return has, nil
}

func (s UserStore) IsExistUserById(id int64) (bool, error) {
	var has bool
	if err := s.db.Table(usermodel.User{}.TableName()).Select("count(id) > 0").Where(
		"id = ?", id,
	).Limit(1).Take(&has).Error; err != nil {
		return false, err
	}

	return has, nil

}

func (s UserStore) buildUserCommonQuery(filter usermodel.UserFilter) *gorm.DB {
	tx := s.db.Table(usermodel.User{}.TableName())
	if filter.Id != 0 {
		tx.Where("id = ?", filter.Id)
	}
	if len(filter.WalletAddress) > 0 {
		tx.Where("BINARY wallet_address = ?", filter.WalletAddress)
	}
	if len(filter.LoginName) > 0 {
		tx.Where("login_name = ?", filter.LoginName)
	}
	if len(filter.Email) > 0 {
		tx.Where("email = ?", filter.Email)
	}
	if len(filter.Ref) > 0 {
		tx.Where("ref = ?", filter.Ref)
	}
	return tx
}

func (s UserStore) GetWalletAddressById(id int64) (string, error) {
	var walletAddress string
	err := s.db.Model(&usermodel.User{}).Where("id = ?", id).Select("wallet_address").Limit(1).Find(&walletAddress).Error
	return walletAddress, err
}

func (s UserStore) GetIdByWalletAddress(walletAddress string) (int64, error) {
	var id int64
	err := s.db.Model(&usermodel.User{}).Where("BINARY wallet_address", walletAddress).Limit(1).Select("id").Find(&id).Error
	return id, err
}

func (s UserStore) GetListUserById(listUserId []int64) ([]usermodel.UserDetail, error) {
	var users []usermodel.UserDetail
	err := s.db.Model(&usermodel.User{}).Where("id IN ?", listUserId).Find(&users).Error
	return users, err
}

func (s UserStore) GetUserDetailById(id int64) (usermodel.UserDetail, error) {
	var user usermodel.UserDetail
	err := s.db.Model(&usermodel.User{}).Where("id = ?", id).Find(&user).Error
	return user, err
}

func (s UserStore) GetUserTransaction(mkId []int64) ([]usermodel.UserTransactionSoldDetail, error) {
	var users []usermodel.UserTransactionSoldDetail
	err := s.db.Model(&usermodel.User{}).Select("users.*, ms.id as transactionSoldId").
		Joins("INNER JOIN market_place_solds as ms ON ms.buyer_id = users.id").
		Where("ms.market_place_id IN ?", mkId).Find(&users).Error
	return users, err
}

func (s UserStore) InsertNotification(noti *usermodel.Notification) error {
	return s.db.Table(usermodel.Notification{}.TableName()).Create(&noti).Error
}

func (s UserStore) BatchInsertNotification(noti []*usermodel.Notification) error {
	return s.db.Table(usermodel.Notification{}.TableName()).CreateInBatches(&noti, 100).Error
}

func (s UserStore) IsExistNotification(noti *usermodel.Notification) (bool, error) {
	var exist bool
	err := s.db.Table(usermodel.Notification{}.TableName()).
		Where("notification_object = ?", noti.NotificationObject).
		Where("user_id = ?", noti.UserId).
		Where("notification_type = ?", noti.NotificationType).
		Select("COUNT(id) > 0").Take(&exist).Error

	if err != nil {
		return false, err
	}

	return exist, err
}

func (s UserStore) ListNotification(noti *usermodel.NotificationFilter, paging *common.Paging) ([]*usermodel.Notification, error) {
	var result []*usermodel.Notification
	tx := s.buildNotificationCommonQuery(noti)

	if paging != nil {
		errCount := tx.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}

	if err := tx.Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserStore) UpdateNotification(updator, selector *usermodel.Notification) error {
	err := s.db.Table(usermodel.Notification{}.TableName()).
		Where(&selector).Updates(&updator).Error
	if err != nil {
		return err
	}
	return nil
}

func (s UserStore) UpdateNotificationSetting(updator, selector *usermodel.NotificationSetting) error {
	update := utils.ConvertStructToMapInterface(updator)
	delete(update, "id")
	err := s.db.Table(usermodel.NotificationSetting{}.TableName()).
		Where(&selector).Updates(&update).Error
	if err != nil {
		return err
	}
	return nil
}

func (s UserStore) CreateNotificationSetting(user *usermodel.NotificationSetting) error {
	return s.db.Table(usermodel.NotificationSetting{}.TableName()).Create(&user).Error
}

func (s UserStore) FindNotificationSetting(selector *usermodel.NotificationSetting) (*usermodel.NotificationSetting, error) {
	var result *usermodel.NotificationSetting
	query := s.db.Table(usermodel.NotificationSetting{}.TableName()).
		Where(&selector)
	err := query.Take(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (s UserStore) buildNotificationCommonQuery(noti *usermodel.NotificationFilter) *gorm.DB {
	tx := s.db.Table(usermodel.Notification{}.TableName())
	tx.Where("notifications.deleted_at IS NULL")
	if noti.UserId > 0 {
		tx.Where("user_id = ?", noti.UserId)
	}
	if noti.Status > 0 {
		tx.Where("status = ?", noti.Status)
	}
	if len(noti.Type) > 0 {
		arrType := strings.Split(noti.Type, ",")
		tx.Where("notification_type IN ?", arrType)
	}
	if len(noti.CustomType) > 0 {
		tx.Where("notification_type IN ?", noti.CustomType)
	}
	return tx
}

func (s UserStore) GetLoginInformation(username string) (usermodel.CMSUserLoginInformation, error) {
	var infor usermodel.CMSUserLoginInformation
	tx := s.db.Table(usermodel.CMSUser{}.TableName()).Where("login_name = ?", username).Find(&infor)
	return infor, tx.Error
}

func (s UserStore) HasLoginName(loginName string) (bool, error) {
	var check bool
	tx := s.db.Table(usermodel.CMSUser{}.TableName()).Where("login_name = ?", loginName).Find(&check)
	if tx.Error != nil {
		return check, tx.Error
	}
	return check, nil
}

func (s UserStore) CreateUserCMS(user *usermodel.CMSUser) error {
	return s.db.Table(usermodel.CMSUser{}.TableName()).Create(user).Error
}

func (s UserStore) GetRefUserByUserId(userId int64) (*usermodel.RefUser, error) {
	var infor *usermodel.RefUser
	tx := s.db.Table(usermodel.RefUser{}.TableName()).
		Where("user_id = ?", userId).Limit(1).Find(&infor)
	return infor, tx.Error
}

func (s UserStore) GetCreatedUserByItemId(itemId int64) (usermodel.User, error) {
	var users usermodel.User
	tx := s.db.Table(usermodel.User{}.TableName()).Select("users.*").
		Joins("Inner Join item_creators as ic ON ic.user_id = users.id").
		Where("ic.item_id = ?", itemId).Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}

func (s UserStore) GetUserWithDistinctWalletAddresses(walletAddresses []string) ([]usermodel.User, error) {
	var users []usermodel.User
	err := s.db.Model(&usermodel.User{}).Distinct("wallet_address").Where("BINARY wallet_address IN ?", walletAddresses).Select("id,wallet_address").Find(&users).Error
	return users, err
}

func (s UserStore) GetUserWithWalletAddresses(walletAddresses string) (*usermodel.User, error) {
	var users usermodel.User
	err := s.db.Model(&usermodel.User{}).Distinct("wallet_address").Where("BINARY wallet_address = ?", walletAddresses).Select("id,wallet_address").Find(&users).Error
	return &users, err
}

func (s UserStore) CountUserItemEvent() (int64, error) {
	var total int64
	err := s.db.Model(&usermodel.User{}).
		Joins("INNER JOIN items ON users.id = items.user_id").
		Where("items.is_join_event = ?", common.ITEM_ON_JOIN_EVENT).Group("users.id").Count(&total).Error
	return total, err
}

func (m UserStore) FollowUser(in *usermodel.UserFollow) error {
	var out *usermodel.UserFollow
	err := m.db.Table(usermodel.UserFollow{}.TableName()).
		Where("follower_id = ?", in.FollowerId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return err.Error
	}

	if out.Id != 0 {
		if out.IsDeleted == common.ACTIVE {
			return common.ErrUserFollowed
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

	result := tx.Table(usermodel.UserFollow{}.TableName()).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_deleted", "time_start_follow"}),
	}).Create(in)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error

}

func (m UserStore) UnFollowUser(in *usermodel.UserFollow) error {
	var out *usermodel.UserFollow
	err := m.db.Table(usermodel.UserFollow{}.TableName()).
		Where("follower_id = ?", in.FollowerId).
		Where("user_id", in.UserId).
		Take(&out)
	if err.Error != nil {
		if err.Error == gorm.ErrRecordNotFound {
			return common.ErrUserNotFollowed
		}
		return err.Error
	}

	if out.IsDeleted == common.DELETED {
		return common.ErrUserUnFollowed
	}

	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if out.IsDeleted == common.ACTIVE {
		if err := tx.Table(usermodel.UserFollow{}.TableName()).
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

func (m UserStore) ListFollowingUser(noti *usermodel.UserFollow) ([]*usermodel.UserFollow, error) {
	var result []*usermodel.UserFollow
	ss := m.buildFollowingCommonQuery(noti)
	err := ss.Where("is_deleted = ?", common.ACTIVE).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserStore) buildFollowingCommonQuery(noti *usermodel.UserFollow) *gorm.DB {
	tx := s.db.Table(usermodel.UserFollow{}.TableName())
	if noti.UserId > 0 {
		tx.Where("user_id = ?", noti.UserId)
	}
	if noti.FollowerId > 0 {
		tx.Where("follower_id = ?", noti.FollowerId)
	}
	return tx
}

func (s UserStore) SearchUser(rq *usermodel.SearchRequest) ([]*usermodel.UserSearchResponse, error) {
	var result []*usermodel.UserSearchResponse
	tx := s.db.Table(usermodel.User{}.TableName()).
		Select("users.id, users.avatar, users.display_name, users.type,users.wallet_address")
	if len(rq.Value) > 0 {
		tx.Where("users.display_name LIKE ?", "%"+rq.Value+"%")
	}
	err := tx.Find(&result).Error
	if err != nil || len(result) < 1 {
		return nil, err
	}

	var listUserId []int64
	for _, user := range result {
		listUserId = append(listUserId, user.Id)
	}

	var tmpFollow []usermodel.UserFollowSearchResponse
	txFollow := s.db.Table(usermodel.UserFollow{}.TableName()).
		Joins("INNER JOIN users ON users.id=user_follows.follower_id").
		Select("count(user_follows.id) as user_follow_id, users.id as user_id").
		Where("users.id IN ?", listUserId).
		Where("user_follows.is_deleted = ?", common.ACTIVE).
		Group("users.id")
	errFollow := txFollow.Find(&tmpFollow).Error
	if errFollow != nil {
		return nil, err
	}

	for index, user := range result {
		for _, follow := range tmpFollow {
			if user.Id == follow.UserId {
				result[index].Followers = follow.UserFollowId
				break
			}
		}
	}

	return result, nil

}

func (s UserStore) SearchUserWd(rq *usermodel.SearchRequest) ([]*usermodel.UserSearchResponse, error) {
	var result []*usermodel.UserSearchResponse
	tx := s.db.Table(usermodel.User{}.TableName()).
		Select("users.id, users.avatar, users.display_name, users.type,users.wallet_address")
	if len(rq.Value) > 0 {
		tx.Where("users.wallet_address LIKE ?", "%"+rq.Value+"%")
	}
	err := tx.Find(&result).Error
	if err != nil || len(result) < 1 {
		return nil, err
	}

	var listUserId []int64
	for _, user := range result {
		listUserId = append(listUserId, user.Id)
	}

	var tmpFollow []usermodel.UserFollowSearchResponse
	txFollow := s.db.Table(usermodel.UserFollow{}.TableName()).
		Joins("INNER JOIN users ON users.id=user_follows.follower_id").
		Select("count(user_follows.id) as user_follow_id, users.id as user_id").
		Where("users.id IN ?", listUserId).
		Where("user_follows.is_deleted = ?", common.ACTIVE).
		Group("users.id")
	errFollow := txFollow.Find(&tmpFollow).Error
	if errFollow != nil {
		return nil, err
	}

	for index, user := range result {
		for _, follow := range tmpFollow {
			if user.Id == follow.UserId {
				result[index].Followers = follow.UserFollowId
				break
			}
		}
	}

	return result, nil

}

func (s UserStore) SearchCollection(rq *usermodel.SearchRequest) ([]*usermodel.CollectionSearchResponse, error) {
	var result []*usermodel.CollectionSearchResponse
	tx := s.db.Table(marketplacemodel.Collection{}.TableName())
	if len(rq.Value) > 0 {
		tx.Where("name LIKE ?", "%"+rq.Value+"%")
	}
	tx.Select("id, logo, name, featured_image, view")
	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserStore) SearchNFT(rq *usermodel.SearchRequest, paging *common.Paging) ([]*usermodel.NFTSearchResponse, error) {
	var result []*usermodel.NFTSearchResponse
	tx := s.db.Table(marketplacemodel.Item{}.TableName())
	tx.Joins("LEFT JOIN collections ON items.collection_id = collections.id").
		Joins("LEFT JOIN market_places ON items.id = market_places.item_id").
		Select("items.id, items.view, items.user_id, items.nft, items.contract_address, "+
			"items.image, items.title, items.is_join_event, items.transaction_id, items.status, items.category_id, items.num_of_copies, items.status, items.main_item_id, "+
			"COAlESCE(SUM(CASE WHEN (market_places.deleted_at IS NULL AND market_places.status = ?) THEN market_places.price*1 END),0) as price, "+
			"collections.name as collection_name, items.total_like", common.MARKET_PLACE_STATUS_SELLING).
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC).
		Where("items.is_hide = ?", common.ITEM_UNHIDE).
		Where("items.deleted_at IS NULL").Group("items.id").
		Group("market_places.item_id")
	tx.Where("(items.main_item_id IS NULL OR items.main_item_id = 0)")
	if len(rq.Value) > 0 {
		tx.Where("title LIKE ?", "%"+rq.Value+"%")
	}
	if len(rq.CategoryIds) > 0 {
		temp := strings.Split(rq.CategoryIds, ",")
		tx.Where("items.category_id IN ?", temp)
	}
	if len(rq.Status) > 0 {
		tx.Where("(market_places.status <> ? OR market_places.status IS NULL)", common.MARKET_PLACE_STATUS_SOLDED).
			Where("market_places.deleted_at IS NULL").Where("market_places.sell_type IS NOT NULL")
		temps := strings.Split(rq.Status, ",")
		tx.Where("items.status IN ?", temps)
	}
	if rq.IsJoinEvent > 0 {
		tx.Where("items.is_join_event = ?", rq.IsJoinEvent)
	}
	if rq.PriceMin > 0 {
		tx.Where("(price/1000000000000000000) >= ?", rq.PriceMin)
	}
	if rq.PriceMax > 0 {
		tx.Where("(price/1000000000000000000) <= ?", rq.PriceMax)
	}
	if rq.CollectionId > 0 {
		tx.Where("items.collection_id = ?", rq.CollectionId)
	}
	if len(rq.OrNotListItemIds) > 0 {
		tx.Where("(NOT items.id IN ?)", rq.OrNotListItemIds)
	}
	for v, _ := range rq.Properties {
		tx.Where("items.properties LIKE ?", "%"+rq.Properties[v].Name+"%")
		temps := fmt.Sprintf("%v", rq.Properties[v].Value)
		tx.Where("items.properties LIKE ?", "%"+temps+"%")
	}
	if rq.OrderBy == 1 {
		tx.Order("market_places.updated_at DESC").Order("items.updated_at DESC")
	} else if rq.OrderBy == 2 {
		tx.Order("price DESC")
	} else if rq.OrderBy == 3 {
		tx.Order("price ASC")
	} else if rq.OrderBy == 4 {
		tx.Order("market_places.auction_end_time ASC, market_places.auction_start_time DESC").
			Where("market_places.auction_end_time > ?", time.Now().Unix())
		rq.SellType = common.MARKET_PLACE_SELL_TYPE_AUCTION
	} else if rq.OrderBy == 5 {
		tx.Order("items.total_like DESC")
	}
	if rq.SellType > 0 {
		tx.Where("market_places.sell_type = ?", rq.SellType)
	}
	if paging != nil && rq.SellType == 0 && rq.PriceMin == 0 {
		errCount := tx.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return nil, nil
		}
		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserStore) CountAvailableItemCopied(filter marketplacemodel.CountItemFilter) (int64, error) {
	query := s.db.Table(marketplacemodel.Item{}.TableName())
	if filter.Id > 0 {
		query.Where("id = ?", filter.Id)
	}
	if filter.MainItemId > 0 {
		query.Where("main_item_id = ?", filter.MainItemId)
	}
	if filter.Status > 0 {
		query.Where("status = ?", filter.Status)
	}
	if len(filter.Stts) > 0 {
		query.Where("status IN (?)", filter.Stts)
	}
	var result int64
	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (s UserStore) SearchMarketPlaceNft(itemId []int64, rq *usermodel.SearchRequest, paging *common.Paging) ([]*usermodel.NFTSearchResponse, error) {
	var result []*usermodel.NFTSearchResponse
	tx := s.db.Table(marketplacemodel.MarketPlace{}.TableName()).
		Joins("INNER JOIN items ON items.id = market_places.item_id").
		Select("market_places.price, market_places.auction_start_time, market_places.auction_end_time, "+
			"market_places.item_id as id, market_places.pack_id as pack_id, items.main_item_id as main_item_id, items.user_id as user_id ").
		Where("(market_places.status <> ? OR market_places.status IS NULL)", common.MARKET_PLACE_STATUS_SOLDED).
		Where("market_places.deleted_at IS NULL").Where("items.deleted_at IS NULL")
	if len(itemId) > 0 {
		tx.Where("market_places.item_id IN ?", itemId)
	}
	if rq.PriceMin > 0 {
		tx.Where("(market_places.price/1000000000000000000) >= ? ", rq.PriceMin)
	}
	if rq.PriceMax > 0 {
		tx.Where("(market_places.price/1000000000000000000) <= ?", rq.PriceMax)
	}
	if rq.SellType > 0 {
		tx.Where("market_places.sell_type = ?", rq.SellType)
	}
	tx.Order("market_places.created_at DESC").Group("market_places.item_id")
	if paging != nil && (rq.SellType != 0 || rq.PriceMin != 0) {
		errCount := tx.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return nil, nil
		}
		tx.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s UserStore) GetListSearchHistory(userId int64) ([]*usermodel.UserSearch, error) {
	var result []*usermodel.UserSearch
	tx := s.db.Table(usermodel.UserSearch{}.TableName())
	tx.Where("user_id = ?", userId).Order("id desc").Limit(6)
	if err := tx.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s UserStore) InsertSearchHistory(search *usermodel.UserSearch) error {
	return s.db.Table(usermodel.UserSearch{}.TableName()).Create(search).Error
}

func (i UserStore) CountFollowing(userId int64) (int64, error) {
	var result sql.NullInt64
	query := i.db.Table(usermodel.UserFollow{}.TableName()).
		Select("COUNT(id)").
		Where("user_id = ?", userId).
		Where("is_deleted = ?", common.ACTIVE)

	if err := query.Take(&result).Error; err != nil {
		return 0, err
	}
	return result.Int64, nil
}

func (i UserStore) CountFollowers(userId int64) (int64, error) {
	var result sql.NullInt64
	query := i.db.Table(usermodel.UserFollow{}.TableName()).
		Select("COUNT(id)").
		Where("follower_id = ?", userId).
		Where("is_deleted = ?", common.ACTIVE)

	if err := query.Take(&result).Error; err != nil {
		return 0, err
	}
	return result.Int64, nil
}

func (i UserStore) CheckFollowing(userId, targetUserId int64) (int64, error) {
	var result int64
	query := i.db.Table(usermodel.UserFollow{}.TableName()).
		Where("user_id = ?", userId).
		Where("follower_id = ?", targetUserId).
		Where("is_deleted = ?", common.ACTIVE)

	if err := query.Count(&result).Error; err != nil {
		return 0, err
	}
	return result, nil
}

func (i UserStore) ListFollow(userId []int64, currentUser int64) ([]*usermodel.ListFollowQueryResponse, error) {
	rawQuery := fmt.Sprintf("user_follows.*, COUNT(user_id) as total_followers, COUNT(CASE WHEN user_follows.user_id = %v THEN 1 END) > 0 as is_following", currentUser)
	var result []*usermodel.ListFollowQueryResponse
	query := i.db.Table(usermodel.UserFollow{}.TableName()).
		Where("is_deleted = ?", common.ACTIVE).
		Where("follower_id IN ?", userId).
		Group("follower_id").
		Select(rawQuery)
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i UserStore) ListActivities(filter usermodel.ActivitiesFilter) ([]*usermodel.Activities, error) {
	var result []*usermodel.Activities
	query := i.BuildActivitiesCommonQuery(filter)
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i UserStore) BuildActivitiesCommonQuery(filter usermodel.ActivitiesFilter) *gorm.DB {
	query := i.db.Table(usermodel.Activities{}.TableName()).
		Joins("LEFT JOIN items on items.id = activities.item_id").
		Where("items.status <> ?", common.ITEM_CREATE_WAITING_SYNC)
	query.Where("activities.deleted_at IS NULL")
	if filter.UserId > 0 {
		query.Where("activities.user_id = ?", filter.UserId)
	}
	if filter.ItemId > 0 {
		query.Where("activities.item_id = ?", filter.ItemId)
	}
	if len(filter.ActivityType) > 0 {
		typeAc := strings.Split(filter.ActivityType, ",")
		query.Where("activities.activity_type IN ?", typeAc)
	}
	if filter.Time > 0 {
		last := time.Now().AddDate(0, 0, -1*int(filter.Time))
		query.Where("activities.created_at >= ?", last)
	}
	if filter.CollectionId > 0 {
		query.Where("activities.collection_id = ?", filter.CollectionId)
	}
	return query.Order("id DESC")
}

func (i UserStore) InsertActivities(activity ...*usermodel.Activities) error {
	if err := i.db.Table(usermodel.Activities{}.TableName()).Create(&activity).Error; err != nil {
		return err
	}
	return nil
}

func (i UserStore) GetListUserRef(filter usermodel.RefNetworkFilter) ([]usermodel.UserDetailAddTimeJoinRef, error) {
	query := i.db.Table(usermodel.User{}.TableName()).
		Select("users.*, ref_users.created_at as time_join_ref").
		Joins("LEFT JOIN ref_users ON users.id = ref_users.user_id").
		Where("(ref_users.user_ref_id = ? OR users.id = ?)", filter.UserId, filter.UserId)
	var result []usermodel.UserDetailAddTimeJoinRef
	if err := query.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (i UserStore) GetUserF0Ref(userId int64) (usermodel.UserDetail, error) {
	query := i.db.Table(usermodel.User{}.TableName()).
		Joins("INNER JOIN ref_users ON users.id = ref_users.user_ref_id").
		Where("ref_users.user_id = ?", userId)
	var result usermodel.UserDetail
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (m UserStore) SystemSettingBatchSave(requests []usermodel.SystemSetting) error {
	tx := m.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	result := tx.Table(usermodel.SystemSetting{}.TableName()).Create(requests)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	return tx.Commit().Error

}

func (i UserStore) DeleteSystemSetting(id int64) error {
	query := i.db.Table(usermodel.SystemSetting{}.TableName()).
		Where("id = ?", id)
	if err := query.Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (i UserStore) GetAllSystemSetting() ([]usermodel.SystemSetting, error) {
	var result []usermodel.SystemSetting
	query := i.db.Table(usermodel.SystemSetting{}.TableName()).Order("display_order DESC")
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (i UserStore) GetBankIncome(filter usermodel.BankIncomeFilter) (usermodel.BankIncome, error) {
	var result usermodel.BankIncome
	query := i.db.Table(usermodel.BankIncome{}.TableName())
	BankIncomeQueryBuilder(query, filter)
	if err := query.Find(&result).Error; err != nil {
		return result, err
	}
	return result, nil
}

func BankIncomeQueryBuilder(query *gorm.DB, filter usermodel.BankIncomeFilter) {
	if filter.Status > 0 {
		query.Where("status = ?", filter.Status)
	}
	if len(filter.TransactionId) > 0 {
		query.Where("transaction_id = ?", filter.Status)
	}
}

func (i UserStore) UpdateBankIncome(id int64, bankIncome usermodel.BankIncome) error {
	query := i.db.Table(usermodel.BankIncome{}.TableName()).Where("id = ?", id).Updates(bankIncome)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (s UserStore) ExploreUser(in usermodel.ExploreUserFilter) ([]usermodel.ExploreUserResponse, error) {
	var result []usermodel.ExploreUserResponse
	tx := s.db.Table(usermodel.User{}.TableName()).
		Joins("LEFT JOIN user_follows ON users.id = user_follows.follower_id").
		Select("COUNT(CASE WHEN user_follows.is_deleted = 0 then user_follows.id END) as total_follow, users.*").
		Group("users.id")

	if len(in.OrderBy) > 0 {
		tx.Order(in.OrderBy)
	}

	if len(in.Search) > 0 {
		tx.Where("users.display_name LIKE ? OR users.wallet_address LIKE ?", "%"+in.Search+"%", "%"+in.Search+"%")
	}

	err := tx.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i UserStore) SystemDailyEventCreate(requests usermodel.DailyEvent) error {
	err := i.db.Table(usermodel.DailyEvent{}.TableName()).Create(&requests).Error
	return err
}

func (i UserStore) SystemDailyEventUpdate(id int64, requests usermodel.DailyEvent) error {
	err := i.db.Table(usermodel.DailyEvent{}.TableName()).Updates(&requests).Where("id = ?", id).Error
	return err
}

func (i UserStore) SystemDailyEventDelete(id int64) error {
	err := i.db.Table(usermodel.DailyEvent{}.TableName()).Where("id = ?", id).Update("deleted_at", time.Now()).Error
	return err
}

func (i UserStore) SystemDailyEventDetail(req usermodel.DailyEventFilter) (usermodel.DailyEvent, error) {
	query := i.db.Table(usermodel.DailyEvent{}.TableName())
	DailyEventQueryBuild(query, req)
	var result usermodel.DailyEvent
	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (i UserStore) SystemDailyEventList(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.DailyEvent, error) {
	query := i.db.Table(usermodel.DailyEvent{}.TableName())
	DailyEventQueryBuild(query, req)
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
	var result []usermodel.DailyEvent
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func DailyEventQueryBuild(query *gorm.DB, req usermodel.DailyEventFilter) {
	query.Where("daily_events.deleted_at IS NULL")
	if req.Id > 0 {
		query.Where("daily_events.id = ?", req.Id)
	}
	if len(req.EventName) > 0 {
		query.Where("daily_events.event_name LIKE %?%", req.EventName)
	}
	if req.Status > 0 {
		query.Where("daily_events.status = ?", req.Status)
	}
	if req.RewardType > 0 {
		query.Where("daily_events.reward_type = ?", req.RewardType)
	}
	if req.RewardOp > 0 {
		query.Where("daily_events.reward_op = ?", req.RewardOp)
	}
	if req.NumberOfTask > 0 {
		query.Where("daily_events.number_of_task = ?", req.NumberOfTask)
	}
	if len(req.StartDateTime) > 0 {
		query.Where("daily_events.start_date_time = ?", req.StartDateTime)
	}
	if len(req.EndDateTime) > 0 {
		query.Where("daily_events.end_date_time = ?", req.EndDateTime)
	}
	if req.EventType > 0 {
		query.Where("daily_events.event_type = ?", req.EventType)
	}
	if req.DailyEventUserId > 0 {
		query.Where("daily_events.user_id = ?", req.UserId)
	}
	if req.UserId > 0 {
		query.Where("user_daily_events.user_id = ?", req.UserId)
	}
	if len(req.CurrentTimeStart) > 0 && len(req.CurrentTimeEnd) > 0 {
		query.Where("(user_daily_events.activity_date >= ?", req.CurrentTimeStart).
			Where("user_daily_events.activity_date <= ?)", req.CurrentTimeEnd)
	}
	if len(req.OrderBy) > 0 {
		query.Order(req.OrderBy)
	}
}

func (i UserStore) SystemUserDailyEventList(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.UserDailyEventExInfo, error) {
	query := i.db.Table(usermodel.UserDailyEvent{}.TableName()).Joins("INNER JOIN daily_events ON daily_events.id = user_daily_events.daily_event_id")
	DailyEventQueryBuild(query, req)
	if paging != nil && paging.Limit > 0 {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return nil, nil
		}
		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	query.Select("user_daily_events.*, event_name, event_type ,daily_events.reward_op as event_reward_op, number_of_task, reward_type")
	var result []usermodel.UserDailyEventExInfo
	err := query.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (i UserStore) SystemUserDailyEventCreate(req usermodel.UserDailyEvent) error {
	err := i.db.Table(usermodel.UserDailyEvent{}.TableName()).Create(&req).Error
	return err
}

func (i UserStore) SystemUserDailyEventCount(req usermodel.DailyEventFilter) (int64, error) {
	query := i.db.Table(usermodel.UserDailyEvent{}.TableName())
	DailyEventQueryBuild(query, req)
	var result int64
	errCount := query.Count(&result).Error
	return result, errCount
}

func (i UserStore) ListRewardDepartment(req usermodel.RewardFilter, paging *common.Paging) ([]usermodel.RewardDepartment, error) {
	query := i.db.Table(usermodel.RewardDepartment{}.TableName())
	RewardQueryBuilder(query, req)
	var result []usermodel.RewardDepartment
	if paging != nil && paging.Limit > 0 {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}

		if paging.Total == 0 {
			return result, nil
		}

		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	query.Preload("Rewards")
	if len(req.OrderBy) < 1 {
		query.Order("reward_departments.display_order desc")
	}
	errCount := query.Find(&result).Error
	return result, errCount
}

func (i UserStore) DetailReward(req usermodel.RewardFilter) (usermodel.Reward, error) {
	query := i.db.Table(usermodel.Reward{}.TableName())
	RewardQueryBuilder(query, req)
	var result usermodel.Reward
	errCount := query.Find(&result).Error
	return result, errCount
}

func RewardQueryBuilder(query *gorm.DB, req usermodel.RewardFilter) {
	if req.RewardId > 0 {
		query.Where("rewards.id = ?", req.RewardId)
	}
	if len(req.RewardName) > 0 {
		query.Where("rewards.reward_name LIKE %?%", req.RewardName)
	}
	if req.UserRewardStatus > 0 {
		query.Where("user_rewards.status = ?", req.UserRewardStatus)
	}
	if req.IsAvailable > 0 {
		query.Where("rewards.price <= ?", req.UserOp)
	}
	if req.IsAvailable > 0 {
		query.Where("(rewards.status = ? OR rewards.reward_end_date_time < ?)", common.REWARD_DEPARTMENT_STATUS_NOT_AVALIABLE, time.Now().Unix())
	}
	if req.UserRewardUserId > 0 {
		query.Where("user_rewards.user_id = ?", req.UserRewardUserId)
	}
	if req.UserRewardId > 0 {
		query.Where("user_rewards.id = ?", req.UserRewardId)
	}
	if req.IsGetClaimed > 0 {
		query.Joins("INNER JOIN rewards ON rewards.reward_department_id = reward_departments.id")
		query.Joins("INNER JOIN user_rewards ON user_rewards.reward_id = rewards.id")
		query.Where("user_rewards.user_id = ?", req.UserId)
	} else if req.UserId > 0 {
		query.Joins("INNER JOIN user_rewards ON user_rewards.reward_id = rewards.id")
		query.Where("user_rewards.user_id = ?", req.UserId)
	}
	if len(req.OrderBy) > 0 {
		query.Order(req.OrderBy)
	}
}

func (i UserStore) SaveRewardDepartment(requests *usermodel.RewardDepartment) error {
	err := i.db.Table(usermodel.RewardDepartment{}.TableName()).Create(&requests).Error
	return err
}

func (i UserStore) UpdateRewardDepartment(id int64, requests usermodel.RewardDepartment) error {
	err := i.db.Table(usermodel.RewardDepartment{}.TableName()).Updates(&requests).Where("id = ?", id).Error
	return err
}

func (i UserStore) SaveReward(requests *usermodel.Reward) error {
	err := i.db.Table(usermodel.Reward{}.TableName()).Create(&requests).Error
	return err
}

func (i UserStore) UpdateReward(id int64, requests usermodel.Reward) error {
	err := i.db.Table(usermodel.Reward{}.TableName()).Updates(&requests).Where("id = ?", id).Error
	return err
}

func (i UserStore) SaveUserReward(requests *usermodel.UserReward, totalClaim, opSpend, op int64) error {
	tx := i.db.Session(&gorm.Session{CreateBatchSize: 100}).Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Printf("run time panic: %v\n\n", r)
			tx.Rollback()
		}
	}()

	if err := tx.Table(usermodel.UserReward{}.TableName()).Create(&requests).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table(usermodel.Reward{}.TableName()).Where("id = ?", requests.RewardId).
		Update("total_rewarded", totalClaim).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table(usermodel.User{}.TableName()).Where("id = ?", requests.UserId).
		Update("op", op).Update("op_spend", opSpend).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (i UserStore) ListUserReward(req usermodel.RewardFilter) ([]usermodel.UserReward, error) {
	query := i.db.Table(usermodel.UserReward{}.TableName())
	RewardQueryBuilder(query, req)
	query.Preload("Reward")
	var result []usermodel.UserReward
	errCount := query.Find(&result).Error
	return result, errCount
}

func (i UserStore) ListReward(paging *common.Paging) ([]usermodel.Reward, error) {
	query := i.db.Table(usermodel.Reward{}.TableName())
	var result []usermodel.Reward
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return result, nil
		}
		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	errCount := query.Find(&result).Error
	return result, errCount
}

func (i UserStore) ListRewardDepart(paging *common.Paging) ([]usermodel.RewardDepartment, error) {
	query := i.db.Table(usermodel.RewardDepartment{}.TableName())
	var result []usermodel.RewardDepartment
	if paging != nil {
		errCount := query.Count(&paging.Total).Error
		if errCount != nil {
			return nil, errCount
		}
		if paging.Total == 0 {
			return result, nil
		}
		query.Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit)
	}
	errCount := query.Find(&result).Error
	return result, errCount
}

func (s UserStore) GetVerifyToken(userId int64) (usermodel.User, error) {
	var result usermodel.User
	query := s.db.Table(usermodel.User{}.TableName())
	query.Where("users.id = ?", userId)
	err := query.Find(&result).Error
	if err != nil {
		return result, err
	}
	return result, nil
}
