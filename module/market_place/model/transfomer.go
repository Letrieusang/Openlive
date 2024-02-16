package marketplacemodel

import (
	usermodel "api-openlive/module/user/model"
	"time"
)

type ItemResponse struct {
	ItemDetail
	Seller                usermodel.User         `json:"seller"`
	Owner                 usermodel.UserDetail   `json:"owner"`
	ItemCreator           usermodel.UserDetail   `json:"item_creator"`
	Price                 string                 `json:"price"`
	IsLike                bool                   `json:"is_like"`
	IdOnMarket            int64                  `json:"id_on_market"`
	UserLikes             []usermodel.UserDetail `json:"user_likes"`
	MarketContractAddress string                 `json:"market_contract_address"`
	SellType              int                    `json:"sell_type"`
	AuctionStartTime      string                 `json:"auction_start_time"`
	AuctionEndTime        string                 `json:"auction_end_time"`
}

type ItemDetailResponse struct {
	*ItemDetail
	IsOldEvent    int                  `json:"is_old_event"`
	AvailableCopy int64                `json:"available_copy"`
	ItemCreator   usermodel.User       `json:"item_creator"`
	Price         string               `json:"price"`
	Owner         usermodel.UserDetail `json:"owner"`
	AuctionBid    []AuctionBid         `json:"auction_bid"`
}

type DetailItemResponse struct {
	ItemCreator      usermodel.User `json:"item_creator"`
	IsLike           bool           `json:"is_like"`
	Collection       *Collection    `json:"collection"`
	Id               int64          `json:"id,omitempty"`
	TransactionId    string         `json:"transaction_id" validate:"required"`
	ContractAddress  string         `json:"contract_address" validate:"required"`
	IdOnMarket       int64          `json:"id_on_market" validate:"required"`
	ItemId           int64          `json:"item_id" validate:"required"`
	Item             ItemDetail     `json:"item,omitempty"`
	SellerId         int64          `json:"seller_id" validate:"required"`
	Seller           usermodel.User `json:"seller,omitempty"`
	Price            string         `json:"price" validate:"required"`
	Status           int64          `json:"status" validate:"required"`
	PackId           int64          `json:"pack_id"`
	TotalVote        int64          `json:"total_vote"`
	TotalLike        int64          `json:"total_like"`
	Description      string         `json:"description"`
	SellType         int            `json:"sell_type"`
	StartPrice       string         `json:"start_price"`
	AuctionStartTime string         `json:"auction_start_time"`
	AuctionEndTime   string         `json:"auction_end_time"`
	AuctionBid       []AuctionBid   `json:"auction_bid"`
	UserCreatedId    int64          `json:"created_user_id"`
	UserUpdatedId    int64          `json:"updated_user_id"`
	AvailableCopy    int64          `json:"available_copy"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type ListMarketPlaceResponse struct {
	*MarketPlace
	IsLike    bool                   `json:"is_like"`
	UserLikes []usermodel.UserDetail `json:"user_likes"`
}

type ListMarketPlaceWithCollectionNameResponse struct {
	*MarketPlaceWithCollectionName
	AuctionBid `json:"auction_bid"`
	IsLike     bool                   `json:"is_like"`
	UserLikes  []usermodel.UserDetail `json:"user_likes"`
}

type ListItemEventResponse struct {
	MetaData MetaItemEventResult    `json:"meta_data"`
	Data     []*ListItemEventResult `json:"data"`
}

type MetaItemEventResult struct {
	TotalUser      int64                      `json:"total_user"`
	TotalNft       int64                      `json:"total_nft"`
	ListCategories []CountItemEventByCategory `json:"list_categories"`
}

type CountItemEventByCategory struct {
	CategoryId   int64  `json:"category_id"`
	CategoryName string `json:"category_name"`
	TotalNft     int64  `json:"total_nft"`
}

type ListItemEventResult struct {
	Id                int64  `json:"id,omitempty"`
	UserId            int64  `json:"user_id"`
	Status            int64  `json:"status"`
	TransactionId     string `json:"transaction_id"`
	ContractAddress   string `json:"contract_address"`
	Title             string `json:"title"`
	Nft               string `json:"nft"`
	Image             string `json:"image"`
	Description       string `json:"description"`
	CategoryId        int64  `json:"category_id"`
	IsApproval        int64  `json:"is_approval"`
	View              int64  `json:"view"`
	UserAvatar        string `json:"user_avatar"`
	UserName          string `json:"user_name"`
	UserWalletAddress string `json:"user_wallet_address"`
	TotalLike         int64  `json:"total_like"`
	CollectionName    string `json:"collection_name"`
	IsJoinEvent       uint8  `json:"is_join_event"`
	IsLike            bool   `json:"is_like"`
	NumOfCopies       int64  `json:"num_of_copies"`
}

type CheckCanJoinEventResponse struct {
	Status bool `json:"status"`
}

type ExploreStatisticResponse struct {
	Total       int64 `json:"total"`
	Art         int64 `json:"art"`
	Game        int64 `json:"game"`
	Music       int64 `json:"music"`
	Photography int64 `json:"photography"`
	Sport       int64 `json:"sport"`
	Video       int64 `json:"video"`
}

type CategoryStatistic struct {
	Category
	Total int64 `json:"total"`
}

type ActivitiesUserResponse struct {
	Id           int64                `json:"id"`
	UserId       int64                `json:"user_id"`
	ItemId       int64                `json:"item_id"`
	FollowId     int64                `json:"follow_id"`
	MetaData     map[string]string    `json:"meta_data"`
	CollectionId int64                `json:"collection_id"`
	ActivityType int64                `json:"activity_type"`
	User         usermodel.UserDetail `json:"user,omitempty"`
	Follow       usermodel.UserDetail `json:"follow,omitempty"`
	ActivityUser usermodel.UserDetail `json:"activity_user,omitempty"`
	Item         Item                 `json:"item,omitempty"`
	Collection   Collection           `json:"collection,omitempty"`
	CreatedAt    time.Time            `json:"created_at,omitempty"`
}

type ItemHistoryResponse struct {
	Id           int64                `json:"id"`
	UserId       int64                `json:"user_id"`
	ItemId       int64                `json:"item_id"`
	MetaData     map[string]string    `json:"meta_data"`
	ActivityType int64                `json:"activity_type"`
	User         usermodel.UserDetail `json:"user,omitempty"`
	Item         Item                 `json:"item,omitempty"`
}

type ListRefInfoResult struct {
	UserId    int64     `json:"user_id" from:"user_id"`
	ItemId    int64     `json:"item_id" from:"item_id"`
	Profit    float64   `json:"profit,omitempty" from:"profit"`
	Price     string    `json:"price,omitempty" from:"price"`
	CreatedAt time.Time `json:"created_at" from:"created_at"`
}

type CollectionExploreResponse struct {
	Id            int64   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Logo          string  `json:"Logo"`
	FeaturedImage string  `json:"featured_image"`
	Banner        string  `json:"banner"`
	FloorPrice    float64 `json:"floor_price"`
	Volume        float64 `json:"volume"`
	TotalItem     int64   `json:"total_item"`
	TotalOwner    int64   `json:"total_owner"`
}

type PropertiesValueResponse struct {
	Name     string `json:"name"`
	TotalNft int64  `json:"total_nft"`
}

type PropertiesResponse struct {
	Name    string                    `json:"name"`
	Total   int64                     `json:"total"`
	Element []PropertiesValueResponse `json:"element"`
}

type AuctionBidResponse struct {
	Id            int64          `json:"id,omitempty"`
	MarketPlaceId int64          `json:"market_place_id" validate:"required"`
	UserId        int64          `json:"user_id" validate:"required"`
	Price         string         `json:"image" validate:"required"`
	UserInfo      usermodel.User `json:"user_info,omitempty"`
}
