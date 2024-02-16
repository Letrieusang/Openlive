package userstorage

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
)

type Storage interface {
	CreateUser(user *usermodel.User) error
	CreateRefUser(user *usermodel.RefUser) error
	UpdateUser(user *usermodel.User) error
	UpdateUserOp(userId, op, opClaimed int64, opHash string) error
	GetWalletAddressById(id int64) (string, error)
	GetIdByWalletAddress(walletAddress string) (int64, error)
	GetUserByCondition(filter usermodel.UserFilter) (usermodel.User, error)
	CheckExistUser(rq *usermodel.CreateUserRequest) (bool, error)
	CheckCanAddRefUser(rq *usermodel.CreateUserRequest) (bool, error)
	IsExistUserById(id int64) (bool, error)
	GetListUserById(listUserId []int64) ([]usermodel.UserDetail, error)
	GetUserDetailById(id int64) (usermodel.UserDetail, error)
	InsertNotification(noti *usermodel.Notification) error
	BatchInsertNotification(noti []*usermodel.Notification) error
	IsExistNotification(noti *usermodel.Notification) (bool, error)
	UpdateNotification(updator, selector *usermodel.Notification) error
	UpdateNotificationSetting(updator, selector *usermodel.NotificationSetting) error
	CreateNotificationSetting(user *usermodel.NotificationSetting) error
	FindNotificationSetting(selector *usermodel.NotificationSetting) (*usermodel.NotificationSetting, error)
	ListNotification(noti *usermodel.NotificationFilter, paging *common.Paging) ([]*usermodel.Notification, error)
	GetLoginInformation(username string) (usermodel.CMSUserLoginInformation, error)
	CreateUserCMS(user *usermodel.CMSUser) error
	HasLoginName(loginName string) (bool, error)
	GetRefUserByUserId(userId int64) (*usermodel.RefUser, error)
	GetCreatedUserByItemId(itemId int64) (usermodel.User, error)
	GetUserWithDistinctWalletAddresses(walletAddresses []string) ([]usermodel.User, error)
	GetUserWithWalletAddresses(walletAddresses string) (*usermodel.User, error)
	CountUserItemEvent() (int64, error)
	FollowUser(in *usermodel.UserFollow) error
	UnFollowUser(in *usermodel.UserFollow) error
	ListFollowingUser(noti *usermodel.UserFollow) ([]*usermodel.UserFollow, error)
	SearchUser(rq *usermodel.SearchRequest) ([]*usermodel.UserSearchResponse, error)
	SearchUserWd(rq *usermodel.SearchRequest) ([]*usermodel.UserSearchResponse, error)
	SearchCollection(rq *usermodel.SearchRequest) ([]*usermodel.CollectionSearchResponse, error)
	SearchNFT(rq *usermodel.SearchRequest, paging *common.Paging) ([]*usermodel.NFTSearchResponse, error)
	SearchMarketPlaceNft(itemId []int64, rq *usermodel.SearchRequest, paging *common.Paging) ([]*usermodel.NFTSearchResponse, error) // search
	GetListSearchHistory(userId int64) ([]*usermodel.UserSearch, error)
	InsertSearchHistory(search *usermodel.UserSearch) error

	CountFollowing(userId int64) (int64, error)
	CountFollowers(userId int64) (int64, error)
	CheckFollowing(userId, targetUserId int64) (int64, error)
	ListFollow(userId []int64, currentUser int64) ([]*usermodel.ListFollowQueryResponse, error)
	InsertActivities(activity ...*usermodel.Activities) error
	ExploreUser(in usermodel.ExploreUserFilter) ([]usermodel.ExploreUserResponse, error)
	CountAvailableItemCopied(filter marketplacemodel.CountItemFilter) (int64, error)

	// activities
	ListActivities(filter usermodel.ActivitiesFilter) ([]*usermodel.Activities, error)

	//ref
	GetListUserRef(filter usermodel.RefNetworkFilter) ([]usermodel.UserDetailAddTimeJoinRef, error)
	GetUserF0Ref(userId int64) (usermodel.UserDetail, error)

	//system setting
	SystemSettingBatchSave(requests []usermodel.SystemSetting) error
	DeleteSystemSetting(id int64) error
	GetAllSystemSetting() ([]usermodel.SystemSetting, error)

	//bank income
	GetBankIncome(filter usermodel.BankIncomeFilter) (usermodel.BankIncome, error)
	UpdateBankIncome(id int64, bankIncome usermodel.BankIncome) error

	//daily event
	SystemDailyEventCreate(requests usermodel.DailyEvent) error
	SystemDailyEventUpdate(id int64, requests usermodel.DailyEvent) error
	SystemDailyEventDelete(id int64) error
	SystemDailyEventDetail(req usermodel.DailyEventFilter) (usermodel.DailyEvent, error)
	SystemDailyEventList(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.DailyEvent, error)
	SystemUserDailyEventList(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.UserDailyEventExInfo, error)
	SystemUserDailyEventCreate(req usermodel.UserDailyEvent) error
	SystemUserDailyEventCount(req usermodel.DailyEventFilter) (int64, error)

	//reward
	ListRewardDepart(paging *common.Paging) ([]usermodel.RewardDepartment, error)
	ListReward(paging *common.Paging) ([]usermodel.Reward, error)
	ListRewardDepartment(req usermodel.RewardFilter, paging *common.Paging) ([]usermodel.RewardDepartment, error)
	DetailReward(req usermodel.RewardFilter) (usermodel.Reward, error)
	SaveRewardDepartment(requests *usermodel.RewardDepartment) error
	UpdateRewardDepartment(id int64, requests usermodel.RewardDepartment) error
	SaveReward(requests *usermodel.Reward) error
	UpdateReward(id int64, requests usermodel.Reward) error
	SaveUserReward(requests *usermodel.UserReward, totalClaim, opSpend, op int64) error
	ListUserReward(req usermodel.RewardFilter) ([]usermodel.UserReward, error)

	//verify mail
	GetVerifyToken(userId int64) (usermodel.User, error)
}
