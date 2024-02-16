package common

const (
	USER_OTP_LENG         = 6
	USER_OTP_DURATION     = 3
	EXPIRED_TOKEN         = 180
	EXPIRED_REFRESH_TOKEN = 7 * 24 * 120
	USER_TYPE             = "customer"
	CMS_USER_TYPE         = "cms_user"
	HASH_PASSWORD         = "OPVMARKET@123456789!"
)

const (
	// common
	DELETED = 1
	ACTIVE  = 0

	//item approve
	APPROVE    = 1
	UN_APPROVE = 0

	//user
	VIP_USER   = 6
	BASIC_USER = 2

	// item status
	ITEM_ON_SALE                        = 1
	ITEM_NEW                            = 2
	ITEM_REJECT                         = 3
	ITEM_PENDING                        = 4
	ITEM_READY_ON_SELL                  = 5
	ITEM_CREATE_WAITING_SYNC            = 9
	ITEM_TRANSFER_RECEIVER_ID_NOT_FOUND = 11

	// hide
	ITEM_HIDE   = 1
	ITEM_UNHIDE = 0

	// item join event
	ITEM_NOT_JOIN_EVENT = 0
	ITEM_ON_JOIN_EVENT  = 1

	// item on market status
	ITEM_STATUS_ON_MARKET    = 1
	ITEM_STATUS_ON_INVENTORY = 2

	// market place status
	MARKET_PLACE_STATUS_SELLING = 1
	MARKET_PLACE_STATUS_SOLDED  = 2

	//market place sold status
	MARKET_PLACE_SOLD_STATUS_BUYED  = 2
	MARKET_PLACE_SOLD_STATUS_REJECT = 1

	//market type
	MARKET_PLACE_SELL_TYPE_SELL    = 1
	MARKET_PLACE_SELL_TYPE_AUCTION = 2

	COUNT_MIN_NFT_SHOW_COLLECTION = -1

	// staking status
	STAKING = 1
	DONE    = 2

	AVATAR     = "avatar"
	COVER      = "cover"
	ITEM       = "item"
	COLLECTION = "collection"
	ITEM_IMAGE = "image"
	// collection
	COLLECTION_NEW     = 1
	COLLECTION_PUBLIC  = 2
	COLLECTION_PENDING = 3
	IS_HOT             = 1

	//like
	LIMIT_USER_LIKE = 6

	//ref
	REF_LINK_URL_DEV  = "https://market.openlivenft.io/connect-wallet"
	REF_LINK_URL_PROD = "https://openlivenft.io/connect-wallet"
	DEFAULT_PROFIT    = 0.1

	//setting
	IS_VERIFY_SELL_NFT     = 1
	IS_NOT_VERIFY_SELL_NFT = 0

	ITEM_JOIN_EVENT_STATUS   = 1
	ITEM_CANCEL_EVENT_STATUS = 2

	EVENT_NFT_BASE_VALUE = 5

	ITEM_JOIN_EVENT_PROFIT        = 20
	ITEM_JOIN_EVENT_STATUS_ACTIVE = 1

	ISO8601     = "2006-01-02"
	ISO8601TIME = "2006-01-02 15:04:05"
)

// blockchain
const (
	BLOCK_STATUS_NEW        = 1
	BLOCK_STATUS_WS_SUCCESS = 2
	BLOCK_STATUS_WS_FAIL    = 3
	BLOCK_STATUS_SUCCESS    = 4

	IS_NOT_BLOCK_PROCESS = 1
	IS_BLOCK_PROCESS     = 2
)

// notification
const (
	NOTIFICATION_TRANSACTION = 1
	NOTIFICATION_COLLECTION  = 2
	NOTIFICATION_NFT         = 3
	NOTIFICATION_LIKE        = 4

	READ   = 1
	UNREAD = 2
)

const (
	FOLLOWING_USER_ACTIVITIES = iota + 1
	FOLLOWER_USER_ACTIVITIES
	AUCTION_ACTIVITIES
	OTHER_USER_ACTIVITIES
	SELL_ACTIVITIES
	USERS_ACTIVITIES
)

const (
	// following activities
	NOTIFICATION_LISTED = iota + 1
	NOTIFICATION_UNLISTED
	PURCHASED
	MINTED
	CREATED_COLLECTION
	CREAT_AUCTION

	//auction activities
	AUCTION_WIN
	AUCTION_BEATEN
	AUCTION_ACCEPTED

	// follow activities
	FOLLOWED

	//orther ther activities
	LIKED

	// sell activities
	ORTHER_USER_SELL

	// sell activities
	SELL
	BUY
	CANCEL
	BID
	USER_CREATE_AUCTION
	JOIN_EVENT
	REMOVE_FROM_EVENT
	HIDE
	BURN

	// admin activities
	HIDE_NFT
	BURN_NFT

	//activity auction
	NOTIFICATION_AUCTION
	NOTIFICATION_AUCTION_CANCEL
	NOTIFICATION_AUCTION_BID
	NOTIFICATION_AUCTION_CLAIM_NFT
	NOTIFICATION_AUCTION_OTHER_PURCHASED
)

var MapSetting = map[string][]int{
	"item_sold":            {ORTHER_USER_SELL},
	"bidding_activities":   {},
	"auction_expiration":   {},
	"out_bid":              {},
	"new_letter":           {},
	"following_activities": {NOTIFICATION_LISTED, NOTIFICATION_UNLISTED, PURCHASED, MINTED, CREATED_COLLECTION, CREAT_AUCTION},
	"like_and_follow":      {FOLLOWED, LIKED},
	"auction_win":          {},
	"your_activities":      {SELL, BUY, CANCEL, BID, USER_CREATE_AUCTION, JOIN_EVENT, REMOVE_FROM_EVENT},
}

const (
	MINTED_ITEM = iota + 1
	LISTED_ITEM
	AUCTION_ITEM
	UNLISTED_ITEM
	PURCHASED_ITEM
	SOLD_ITEM
	BID_ITEM
	LIKE_ITEM
	UNLIKE_ITEM
	FOLLOW_USER
	UNFOLLOW_USER
	TRANSFERED_ITEM
	AUCTION_CANCEL
	AUCTION_CLAIM_ITEM
)

const (
	ART = iota + 1
	GAME
	PHOTOGRAPHY
	MUSIC
	VIDEO
	SPORT
)

// system setting
const (
	SYSTEM_SETTING_TYPE_TOP_BANNER   = 1
	SYSTEM_SETTING_TYPE_LEFT_BANNER  = 2
	SYSTEM_SETTING_TYPE_RIGHT_BANNER = 3
	SYSTEM_SETTING_TYPE_TEXT_BANNER  = 4
)

// Bank income
const (
	BANK_INCOME_STATUS_WAITNG       = 1
	BANK_INCOME_STATUS_NOT_SEND_OPV = 2
	BANK_INCOME_STATUS_CANCEL       = 3
	BANK_INCOME_STATUS_SUCCESS      = 4

	WITHDRAW_STATUS_CREATE_TX_FAILED = 6
	WITHDRAW_STATUS_NEW              = 1
	WITHDRAW_STATUS_PROCESSING       = 2
	WITHDRAW_STATUS_APPROVED         = 3
	WITHDRAW_STATUS_REJECT           = 4
	WITHDRAW_STATUS_FAILED           = 5

	BANK_INCOME_SECRET_KEY = "ua!uE2sSc4GGP8cFK5QmD9Nc%Koa*#vw%@ZTR^eEyksPTys#&Zp&FWBykkTba4ZC"
)

// Daily event
const (
	DAILY_EVENT_TYPE_MINT_NFT   = 1
	DAILY_EVENT_TYPE_SELL_NFT   = 2
	DAILY_EVENT_TYPE_BUY_NFT    = 3
	DAILY_EVENT_TYPE_JOIN_EVENT = 4
	DAILY_EVENT_TYPE_LIKE       = 5
	DAILY_EVENT_TYPE_SHARE      = 6
	DAILY_EVENT_TYPE_LOGIN      = 7
	DAILY_EVENT_TYPE_SWAP       = 8
)

// Daily event
const (
	REWARD_DEPARTMENT_STATUS_AVALIABLE     = 1
	REWARD_DEPARTMENT_STATUS_NOT_AVALIABLE = 2

	USER_REWARD_STATUS_NEW  = 1
	USER_REWARD_STATUS_USED = 2
)

// verify mail
const (
	EMAIL_NOT_VERIFIED = 1
	EMAIL_VERIFIED     = 2
)
