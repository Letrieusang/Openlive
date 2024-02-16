package usermodel

import (
	"api-openlive/common"
	"time"
)

type User struct {
	Id              int64     `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Email           string    `gorm:"type:VARCHAR(255)" json:"email"`
	TmpEmail        string    `gorm:"type:VARCHAR(255)" json:"tmp_email"`
	IsVerifiedEmail int       `gorm:"type:TINYINT(2) default 0" json:"is_verified_email"`
	VerifyToken     string    `gorm:"type:TEXT" json:"verify_token"`
	Phone           string    `gorm:"type:VARCHAR(16)" json:"phone"`
	Avatar          string    `gorm:"type:TEXT" json:"avatar"`
	Cover           string    `gorm:"type:TEXT" json:"cover"`
	DisplayName     string    `gorm:"type:VARCHAR(255)" json:"display_name"`
	Op              int64     `gorm:"type:INT(11)" json:"op"`
	OpSpend         int64     `gorm:"type:INT(11)" json:"op_spend"`
	OpClaimed       int64     `gorm:"type:INT(11)" json:"op_claimed"`
	OpBackup        string    `gorm:"type:VARCHAR(255)" json:"op_backup"`
	Description     string    `gorm:"type:TEXT" json:"description" json:"description"`
	Status          int64     `gorm:"type:INT(2)" json:"status,omitempty"`
	WalletAddress   string    `gorm:"uniqueIndex;type:VARCHAR(255)" json:"wallet_address" validate:"required"`
	FbLink          string    `gorm:"type:TEXT" json:"fb_link"`
	TwitterLink     string    `gorm:"type:TEXT" json:"twitter_link"`
	InsLink         string    `gorm:"type:TEXT" json:"ins_link"`
	LinkInLink      string    `gorm:"type:TEXT" json:"link_in_link"`
	Type            int64     `gorm:"type:INT(11)" json:"type"`
	Ref             string    `gorm:"type:VARCHAR(255)" json:"ref"`
	ExpireMail      time.Time `gorm:"type:TIMESTAMP NULL DEFAULT NULL" json:"expire_mail"`
	//common.SQLModel `swaggerignore:"true"`
}

type UserDetail struct {
	Id            int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Email         string `gorm:"type:VARCHAR(255)" json:"email"`
	Phone         string `gorm:"type:VARCHAR(16)" json:"phone"`
	Avatar        string `gorm:"type:TEXT" json:"avatar"`
	Cover         string `gorm:"type:TEXT" json:"cover"`
	DisplayName   string `gorm:"type:VARCHAR(255)" json:"display_name"`
	Description   string `gorm:"type:TEXT" json:"description"`
	WalletAddress string `gorm:"uniqueIndex;type:VARCHAR(255)" json:"wallet_address" validate:"required"`
	FbLink        string `gorm:"type:TEXT" json:"fb_link"`
	TwitterLink   string `gorm:"type:TEXT" json:"twitter_link"`
	InsLink       string `gorm:"type:TEXT" json:"ins_link"`
	LinkInLink    string `gorm:"type:TEXT" json:"link_in_link"`
	Type          int64  `gorm:"type:INT(11)" json:"type"`
	Ref           string `json:"ref"`
	Op            int64  `json:"op"`
	OpSpend       int64  `json:"op_spend"`
	OpClaimed     int64  `json:"op_claimed"`
	OpBackup      string `json:"op_backup"`
}

type UserDetailAddTimeJoinRef struct {
	Id            int64     `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Email         string    `gorm:"type:VARCHAR(255)" json:"email"`
	Phone         string    `gorm:"type:VARCHAR(16)" json:"phone"`
	Avatar        string    `gorm:"type:TEXT" json:"avatar"`
	Cover         string    `gorm:"type:TEXT" json:"cover"`
	DisplayName   string    `gorm:"type:VARCHAR(255)" json:"display_name"`
	Description   string    `gorm:"type:TEXT" json:"description"`
	WalletAddress string    `gorm:"uniqueIndex;type:VARCHAR(255)" json:"wallet_address" validate:"required"`
	FbLink        string    `gorm:"type:TEXT" json:"fb_link"`
	TwitterLink   string    `gorm:"type:TEXT" json:"twitter_link"`
	InsLink       string    `gorm:"type:TEXT" json:"ins_link"`
	LinkInLink    string    `gorm:"type:TEXT" json:"link_in_link"`
	Type          int64     `gorm:"type:INT(11)" json:"type"`
	Ref           string    `json:"ref"`
	TimeJoinRef   time.Time `json:"time_join_ref"`
}

type ExploreUserResponse struct {
	UserDetail
	TotalFollow int64 `json:"total_follow"`
}

type TopSellerResponse struct {
	UserDetail
	TotalAmount         string  `json:"total_amount"`
	TotalNFTSold        int64   `json:"total_nft_sold"`
	USDTVolume          float64 `json:"usdt_volume"`
	PercentVolume       float64 `json:"percent_volume"`
	PercentTotalNftSold float64 `json:"percent_total_nft_sold"`
}

type TopCreatorResponse struct {
	UserDetail
	TotalCreated        int64   `json:"total_created"`
	TotalCollection     int64   `json:"total_collection"`
	PercentTotalCreated float64 `json:"percent_total_created"`
}

type TopCreatorFilter struct {
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
}

type UserTransactionSoldDetail struct {
	Id                int64  `json:"id,omitempty"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Avatar            string `json:"avatar"`
	Cover             string `json:"cover"`
	DisplayName       string `json:"display_name"`
	Description       string `json:"description"`
	WalletAddress     string `json:"wallet_address"`
	FbLink            string `json:"fb_link"`
	TwitterLink       string `json:"twitter_link"`
	InsLink           string `json:"ins_link"`
	LinkInLink        string `json:"link_in_link"`
	TransactionSoldId int64  `json:"transactionSoldId"`
}

func (User) TableName() string { return "users" }

type Notification struct {
	Id                 int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	Status             int64  `gorm:"type:INT(2)" json:"status"`
	UserId             int64  `gorm:"type:INT(11)" json:"user_id"`
	NotificationType   int64  `gorm:"type:INT(11)" json:"notification_type"`
	NotificationObject string `gorm:"type:LONGTEXT" json:"notification_object"`
	LastTimeRead       int64  `gorm:"type:INT(11)" json:"last_time_read"`
	common.SQLModel    `swaggerignore:"true"`
}

func (Notification) TableName() string { return "notifications" }

type NotificationSetting struct {
	Id                  int64 `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	UserId              int64 `gorm:"type:INT(11)" json:"user_id"`
	ItemSold            int64 `gorm:"type:INT(11)" json:"item_sold"`
	BiddingActivities   int64 `gorm:"type:INT(11)" json:"bidding_activities"`
	AuctionExpiration   int64 `gorm:"type:INT(11)" json:"auction_expiration"`
	OutBid              int64 `gorm:"type:INT(11)" json:"out_bid"`
	NewLetter           int64 `gorm:"type:INT(11)" json:"new_letter"`
	FollowingActivities int64 `gorm:"type:INT(11)" json:"following_activities"`
	LikeAndFollow       int64 `gorm:"type:INT(11)" json:"like_and_follow"`
	AuctionWin          int64 `gorm:"type:INT(11)" json:"auction_win"`
	YourActivities      int64 `gorm:"type:INT(11)" json:"your_activities"`
	common.SQLModel     `swaggerignore:"true"`
}

func (NotificationSetting) TableName() string { return "notification_settings" }

type CMSUser struct {
	Id              int    `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	LoginName       string `gorm:"type:VARCHAR(60) NOT NULL" json:"username" validate:"required"`
	Password        string `gorm:"type:VARCHAR(60) NOT NULL" json:"password" validate:"required"`
	Email           string `gorm:"type:VARCHAR(255)" json:"email" validate:"required"`
	Phone           string `gorm:"type:VARCHAR(16)" json:"phone" validate:"required"`
	Role            string `gorm:"type:VARCHAR(16)" json:"role" validate:"required"`
	Avatar          string `json:"avatar"`
	common.SQLModel `swaggerignore:"true"`
}

func (CMSUser) TableName() string { return "cms_users" }

type CMSUserLoginInformation struct {
	CMSUserId int
	Role      string
	Password  string
}

type RefUser struct {
	Id              int     `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	UserId          int64   `gorm:"type:INT(11)" json:"user_id"`
	UserRefId       int64   `gorm:"type:INT(11)" json:"user_ref_id"`
	Profit          float64 `gorm:"type:DOUBLE" json:"profit"`
	RefCode         string  `gorm:"type:VARCHAR(255)" json:"ref_code"`
	common.SQLModel `swaggerignore:"true"`
}

func (RefUser) TableName() string { return "ref_users" }

type UserFollow struct {
	Id              int   `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	UserId          int64 `gorm:"type:INT(11)" json:"user_id"`
	FollowerId      int64 `gorm:"type:INT(11)" json:"follower_id"`
	TimeStartFollow int64 `gorm:"type:INT(11)" json:"time_start_follow"`
	IsDeleted       int64 `gorm:"type:INT(11)" json:"is_deleted"`
	common.SQLModel `swaggerignore:"true"`
}

func (UserFollow) TableName() string { return "user_follows" }

type UserSearch struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	UserId          int64  `gorm:"type:INT(11)" json:"user_id"`
	Search          string `gorm:"type:TEXT" json:"search"`
	common.SQLModel `swaggerignore:"true"`
}

func (UserSearch) TableName() string { return "user_searches" }

type Activities struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	UserId          int64  `gorm:"type:INT(11)" json:"user_id"`
	ItemId          int64  `gorm:"type:INT(11)" json:"item_id"`
	FollowId        int64  `gorm:"type:INT(11)" json:"follow_id"`
	MetaData        string ` gorm:"type:TEXT" json:"meta_data"`
	CollectionId    int64  `gorm:"type:INT(11)" json:"collection_id"`
	ActivityType    int64  `gorm:"type:INT(11)" json:"activity_type"`
	ActivityUserId  int64  `gorm:"type:INT(11)" json:"activity_user_id"`
	common.SQLModel `swaggerignore:"true"`
}

func (Activities) TableName() string { return "activities" }

type SystemSetting struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	Type            int    `gorm:"type:TINYINT(2) NOT NULL default 1" json:"type" validate:"required"`
	Image           string `gorm:"type:TEXT" json:"image"`
	Link            string `gorm:"type:TEXT" json:"link"`
	DisplayOrder    int    `gorm:"type:INT(10) NOT NULL default 0" json:"display_order"`
	Description     string `gorm:"type:TEXT" json:"description"`
	common.SQLModel `swaggerignore:"true"`
}

func (SystemSetting) TableName() string { return "system_settings" }

type DefaultSetting struct {
	LeftBanner  []SystemSetting `json:"left_banner"`
	RightBanner []SystemSetting `json:"right_banner"`
	TopBanner   []SystemSetting `json:"top_banner"`
	TextBanner  []SystemSetting `json:"text_banner"`
}

type BankIncome struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	TransactionId   int64  `gorm:"type:INT(11) NOT NULL" json:"transaction_id" validate:"required"`
	Amount          string `gorm:"type:VARCHAR(255)" json:"amount" validate:"required"`
	Status          int    `gorm:"type:TINYINT(4) NOT NULL default 1" json:"status"`
	BankId          string `gorm:"type:VARCHAR(60)" json:"bank_id"`
	Sign            string `gorm:"type:TEXT" json:"sign"`
	Description     string `gorm:"type:LONGTEXT" json:"description"`
	common.SQLModel `swaggerignore:"true"`
}

func (BankIncome) TableName() string { return "bank_incomes" }

type PaymentNotificationRequest struct {
	Status      int    `json:"status" from:"status"`
	Amount      string `json:"amount" from:"amount"`
	OrderId     string `json:"order_id" from:"order_id"`
	PaymentTime string `json:"time" from:"time"`
	Sign        string `json:"sign" from:"sign"`
	SecretKey   string `json:"secret_key" from:"secret_key"`
}

type BankIncomeFilter struct {
	Status        int    `json:"status" from:"status"`
	Amount        string `json:"amount" from:"amount"`
	TransactionId string `json:"transaction_id" from:"transaction_id"`
}

type DailyEvent struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	EventName       string `gorm:"type:VARCHAR(255) NOT NULL" json:"event_name" validate:"required"`
	EventType       int    `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"event_type"`
	RewardType      int    `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"reward_type"`
	RewardOp        int64  `gorm:"type:INT(11) NOT NULL DEFAULT 1" json:"reward_op" validate:"required"`
	Status          int    `gorm:"type:TINYINT(4) NOT NULL default 1" json:"status"`
	NumberOfTask    int    `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"number_of_task"`
	StartDateTime   string `gorm:"type:VARCHAR(255)" json:"start_date"`
	StartDateEnd    string `gorm:"type:VARCHAR(255)" json:"start_date_end"`
	Description     string `gorm:"type:LONGTEXT" json:"description"`
	DisplayOrder    int    `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"display_order"`
	Thumbnail       string `gorm:"type:VARCHAR(255)" json:"thumbnail"`
	common.SQLModel `swaggerignore:"true"`
}

func (DailyEvent) TableName() string { return "daily_events" }

type UserDailyEvent struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	UserId          int64  `gorm:"type:INT(11)" json:"user_id"`
	DailyEventId    int64  `gorm:"type:INT(11)" json:"daily_event_id"`
	RewardOp        int64  `gorm:"type:INT(11) NOT NULL DEFAULT 1" json:"reward_op" validate:"required"`
	ActivityDate    string `gorm:"type:VARCHAR(255)" json:"activity_date"`
	Description     string `gorm:"type:LONGTEXT" json:"description"`
	common.SQLModel `swaggerignore:"true"`
}

func (UserDailyEvent) TableName() string { return "user_daily_events" }

type UserDailyEventExInfo struct {
	Id                  int64  `json:"id"`
	UserId              int64  `json:"user_id"`
	DailyEventId        int64  `json:"daily_event_id"`
	RewardOp            int64  `json:"reward_op" validate:"required"`
	ActivityDate        string `json:"activity_date"`
	Description         string `json:"description"`
	EventName           string `json:"event_name" validate:"required"`
	EventRewardOp       int64  `json:"event_reward_op" validate:"required"`
	NumberOfTask        int    `json:"number_of_task"`
	NumberOfCurrentTask int    `json:"number_of_current_task"`
	RewardType          int    `json:"reward_type"`
	EventType           int    `json:"event_type"`
}

type Reward struct {
	Id                  int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	RewardDepartmentId  int64  `gorm:"type:INT(11) NOT NULL" json:"reward_department_id" validate:"required"`
	Status              int    `gorm:"type:TINYINT(4) NOT NULL default 1" json:"status"`
	Price               int64  `gorm:"type:INT(11) NOT NULL DEFAULT 0" json:"price"`
	RewardName          string `gorm:"type:VARCHAR(255)" json:"reward_name" validate:"required"`
	Description         string `gorm:"type:LONGTEXT" json:"description"`
	DisplayOrder        int    `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"display_order"`
	RewardStartDateTime string `gorm:"type:VARCHAR(255)" json:"reward_start_date_time"`
	RewardEndDateTime   string `gorm:"type:VARCHAR(255)" json:"reward_end_date_time"`
	ExpiredDateTime     string `gorm:"type:VARCHAR(255)" json:"expired_date_time"`
	TotalReward         int64  `gorm:"type:INT(11) NOT NULL default -1" json:"total_reward"`
	TotalRewarded       int64  `gorm:"type:INT(11)" json:"total_rewarded"`
	Image               string `gorm:"type:VARCHAR(255)" json:"image"`
	common.SQLModel     `swaggerignore:"true"`
}

func (Reward) TableName() string { return "rewards" }

type RewardDepartment struct {
	Id                   int64    `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	RewardDepartmentName string   `gorm:"type:VARCHAR(255) NOT NULL" json:"reward_department_name" validate:"required"`
	Status               int      `gorm:"type:TINYINT(4) NOT NULL default 1" json:"status"`
	RewardDepartmentType int      `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"reward_department_type"`
	Description          string   `gorm:"type:LONGTEXT" json:"description"`
	DisplayOrder         int      `gorm:"type:TINYINT(4) NOT NULL DEFAULT 1" json:"display_order"`
	Thumbnail            string   `gorm:"type:VARCHAR(255)" json:"thumbnail"`
	Rewards              []Reward `json:"rewards,omitempty"`
	common.SQLModel      `swaggerignore:"true"`
}

type RewardDepartmentDto struct {
	Id                   int64    `json:"id"`
	RewardDepartmentName string   `json:"reward_department_name" validate:"required"`
	Status               int      `json:"status"`
	RewardDepartmentType int      `json:"reward_department_type"`
	Description          string   `json:"description"`
	DisplayOrder         int      `json:"display_order"`
	Thumbnail            string   `json:"thumbnail"`
	Rewards              []Reward `gorm:"migration" json:"rewards,omitempty"`
}
type CreateRewardDepartmentRequest struct {
	Id                   int64  `json:"id"`
	RewardDepartmentName string `json:"reward_department_name" validate:"required"`
	Status               int    `json:"status"`
	RewardDepartmentType int    `json:"reward_department_type"`
	Description          string `json:"description"`
	DisplayOrder         int    `json:"display_order"`
	Thumbnail            string `json:"thumbnail"`
}

func (c CreateRewardDepartmentRequest) ConvertToMainStruct() RewardDepartment {
	var result = RewardDepartment{
		Id:                   c.Id,
		RewardDepartmentName: c.RewardDepartmentName,
		Status:               c.Status,
		RewardDepartmentType: c.RewardDepartmentType,
		Description:          c.Description,
		DisplayOrder:         c.DisplayOrder,
		Thumbnail:            c.Thumbnail,
	}
	return result
}

func (RewardDepartment) TableName() string { return "reward_departments" }

type UserReward struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id"`
	UserId          int64  `gorm:"type:INT(11) NOT NULL" json:"user_id" validate:"required"`
	RewardId        int64  `gorm:"type:INT(11) NOT NULL" json:"reward_id" validate:"required"`
	Reward          Reward `gorm:"migration" json:"reward,omitempty"`
	UserRewardTitle string `gorm:"type:VARCHAR(255)" json:"user_reward_title"`
	Status          int    `gorm:"type:TINYINT(4) NOT NULL default 1" json:"status"`
	Price           int64  `gorm:"type:INT(11) NOT NULL DEFAULT 0" json:"price"`
	ExpiredDateTime string `gorm:"type:VARCHAR(255)" json:"expired_date_time"`
	Description     string `gorm:"type:LONGTEXT" json:"description"`
	common.SQLModel `swaggerignore:"true"`
}

func (UserReward) TableName() string { return "user_rewards" }

type RewardDto struct {
	Id                  int64  `json:"id"`
	RewardDepartmentId  int64  `json:"reward_department_id"`
	Status              int    `json:"status"`
	Price               int64  `json:"price"`
	RewardName          string `json:"reward_name"`
	Description         string `json:"description"`
	DisplayOrder        int    `json:"display_order"`
	RewardStartDateTime string `json:"reward_start_date_time"`
	RewardEndDateTime   string `json:"reward_end_date_time"`
	Image               string `json:"image"`
	TotalReward         int64  `json:"total_reward"`
	TotalRewarded       int64  `json:"total_rewarded"`
	ClaimTime           string `json:"claim_time"`
}

func (r Reward) ConvertToDto() RewardDto {
	var result = RewardDto{
		Id:                  r.Id,
		RewardDepartmentId:  r.RewardDepartmentId,
		Status:              r.Status,
		Price:               r.Price,
		RewardName:          r.RewardName,
		Description:         r.Description,
		DisplayOrder:        r.DisplayOrder,
		RewardStartDateTime: r.RewardStartDateTime,
		RewardEndDateTime:   r.RewardEndDateTime,
		TotalReward:         r.TotalReward,
		TotalRewarded:       r.TotalRewarded,
		Image:               r.Image,
	}
	return result
}

type ClaimReward struct {
	RewardId int64 `json:"reward_id"`
	UserId   int64 `json:"user_id"`
}

type SendVerifyEmailRequest struct {
	UserId      int64  `json:"user_id"`
	TmpEmail    string `json:"tmp_email" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
}

type VerifyEmailRequest struct {
	VerifyToken string `json:"verify_token" validate:"required"`
	UserId      int64  `json:"user_id" validate:"required"`
}

type ExchangeOPV struct {
	OPV float64 `json:"opv" validate:"required"`
}
