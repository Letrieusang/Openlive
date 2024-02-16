package usermodel

type UserFilter struct {
	Id            int64  `json:"id"`
	CurrentId     int64  `json:"current_id"`
	WalletAddress string `json:"wallet_address"`
	LoginName     string `json:"login_name"`
	Email         string `json:"email"`
	Ref           string `json:"ref"`
	IsVerified    int    `json:"is_verified"`
}

type ExploreUserFilter struct {
	OrderBy string `json:"order_by" form:"order_by"`
	Search  string `json:"search" form:"search"`
}

type NotificationFilter struct {
	UserId     int64  `json:"user_id"`
	Status     int64  `json:"status" form:"status"`
	Type       string `json:"type" form:"type"`
	CustomType []int  `json:"custom_type"`
}

type ActivitiesFilter struct {
	UserId          int64  `gorm:"type:INT(11)" json:"user_id"`
	ItemId          int64  `gorm:"type:INT(11)" json:"item_id"`
	FollowId        int64  `gorm:"type:INT(11)" json:"follow_id"`
	MetaData        string ` gorm:"type:TEXT" json:"meta_data"`
	CollectionId    int64  `gorm:"type:INT(11)" json:"collection_id"`
	ActivityType    string `json:"activity_type"`
	Nft             string `json:"nft" form:"nft"`
	ContractAddress string `json:"contract_address" form:"contract_address"`
	Time            int64  `json:"time"`
}

type RefNetworkFilter struct {
	UserId           int64   `json:"id" form:"id"`
	TimeFilter       int     `json:"time" form:"time"`
	SortUserName     int     `json:"sort_user_name" form:"sort_user_name"`
	SortEventVolume  int     `json:"sort_event_volume" form:"sort_event_volume"`
	SortMarketVolume int     `json:"sort_market_volume" form:"sort_market_volume"`
	StartTimeStamp   int64   `json:"start_time_stamp" `
	UserIds          []int64 `json:"user_ids"`
}

type RefNetworkResult struct {
	UserDetailAddTimeJoinRef
	InviterId    string              `json:"inviter_id"`
	RefLink      string              `json:"ref_link"`
	ListUserRefs []ListRefUserResult `json:"list_user_refs"`
	TotalEarning string              `json:"total_earning"`
	TotalMember  int64               `json:"total_member"`
}

type ListRefUserResult struct {
	Id                int64  `json:"id"`
	Avatar            string `json:"avatar"`
	DisplayName       string `json:"display_name"`
	WalletAddress     string `json:"wallet_address"`
	EventVolume       string `json:"event_volume"`
	EventTotalNFT     int64  `json:"event_total_nft"`
	MarketTotalVolume string `json:"market_total_volume"`
	MarketVolume      string `json:"market_volume"`
}

type DailyEventFilter struct {
	Id               int64  `json:"id" from:"id"`
	EventName        string `json:"event_name" from:"event_name"`
	EventType        int    `json:"event_type" from:"event_type"`
	RewardType       int    `json:"reward_type" from:"reward_type"`
	RewardOp         int    `json:"reward_op" from:"reward_op"`
	Status           int    `json:"status" from:"status"`
	NumberOfTask     int    `json:"number_of_task" from:"number_of_task"`
	StartDateTime    string `json:"start_date_time" from:"start_date_time"`
	EndDateTime      string `json:"end_date_time" from:"end_date_time"`
	OrderBy          string `json:"order_by" from:"order_by"`
	UserId           int64  `json:"user_id" from:"user_id"`
	DailyEventUserId int64  `json:"daily_event_user_id" from:"daily_event_user_id"`
	CurrentTime      string `json:"current_time" from:"current_time"`
	CurrentTimeStart string `json:"current_time_start" from:"current_time_start"`
	CurrentTimeEnd   string `json:"current_time_end" from:"current_time_end"`
	DailyEventId     int64  `json:"daily_event_id" from:"daily_event_id"`
	IsHistory        int    `json:"is_history" from:"is_history"`
}

type RewardFilter struct {
	RewardId         int64  `json:"reward_id" from:"reward_id"`
	UserId           int64  `json:"user_id" from:"user_id"`
	RewardName       string `json:"reward_name" from:"reward_name"`
	UserOp           int64  `json:"user_op"  from:"reward_name"`
	IsAvailable      int    `json:"is_available" from:"is_available"`
	IsPastReward     int    `json:"is_past_reward" from:"is_past_reward"`
	IsGetClaimed     int    `json:"is_get_claimed" from:"is_get_claimed"`
	UserRewardStatus int    `json:"user_reward_status" from:"user_reward_status"`
	UserRewardUserId int64  `json:"user_reward_user_id" from:"user_reward_user_id"`
	UserRewardId     int64  `json:"user_reward_id" from:"user_reward_id"`
	OrderBy          string `json:"order_by" from:"order_by"`
}
