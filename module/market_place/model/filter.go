package marketplacemodel

type ItemFilter struct {
	Id               int64   `json:"id" form:"id"`
	CategoryID       int64   `json:"category_id" form:"category_id"`
	Nft              string  `json:"nft" form:"nft"`
	Ids              []int64 `json:"ids" form:"ids"`
	UserId           int64   `json:"user_id" form:"user_id"`
	ContractAddress  string  `json:"contract_address" form:"contract_address"`
	Status           int64   `json:"status" form:"status"`
	CollectionId     int64   `json:"collection_id" form:"collection_id"`
	CurrentUserId    int64   `json:"current_user_id" form:"current_user_id"`
	PresentUserId    int64   `json:"present_user_id" form:"present_user_id"`
	NotCollection    int64   `json:"not_collection" form:"not_collection"`
	Price            string  `json:"price" form:"price"`
	View             string  `json:"view" form:"view"`
	Favourite        string  `json:"favourite" form:"favourite"`
	Created          string  `json:"created" form:"created"`
	Stts             []int64 `json:"stts"`
	IsApproval       int64   `json:"is_approval" form:"is_approval"`
	IsJoinEvent      uint8   `json:"is_join_event" form:"is_join_event"`
	IsFavourite      int64   `json:"is_favourite" form:"is_favourite"`
	IsCreated        int64   `json:"is_created" form:"is_created"`
	IsOwned          int64   `json:"is_owned" form:"is_owned"`
	IsSold           int64   `json:"is_sold" form:"is_sold"`
	CollectionIds    []int64 `json:"collection_ids"`
	Search           string  `json:"search" form:"search"`
	OrderBy          string  `json:"order_by" form:"order_by"`
	IsGetMulti       int     `json:"is_get_multi"`
	IsNotGetMulti    int     `json:"is_not_get_multi"`
	IsGetMultiId     int64   `json:"is_get_multi_id"`
	IsTrending       int64   `json:"is_trending" form:"is_trending"`
	IsHide           int     `json:"is_hide"`
	IsBuy            int     `json:"is_buy"`
	OrNotListItemIds []int64 `json:"or_list_item_ids"`
	OwnerId          int64   `json:"owner_id" form:"owner_id"`
}

type ItemGetListIdFilter struct {
	ItemId   int64 `json:"item_id" form:"item_id"`
	Number   int   `json:"number" form:"number"`
	IsBuy    int   `json:"is_buy" form:"is_buy"`
	SellerId int64 `json:"seller_id"  form:"seller_id"`
}

type CountItemFilter struct {
	Id           int64 `json:"id" form:"id"`
	MainItemId   int64 `json:"main_item_id" form:"main_item_id"`
	Status       int   `json:"status" form:"status"`
	Stts         []int `json:"stts"`
	CheckMarket  int   `json:"check_market" form:"check_market"`
	MarketUserId int64 `json:"market_user_id" form:"check_market"`
}

type CategoryFilter struct {
}

type CollectionFilter struct {
	IsHot      int `json:"is_hot" form:"is_hot"`
	IsTrending int `json:"is_trending" form:"is_trending"`
	UserId     int64
	UserIdLike int64    `json:"user_id" form:"user_id"`
	Status     int64    `json:"status" form:"status"`
	UserIds    []int64  `json:"user_ids"`
	Cols       []string `json:"cols"`
	Search     string   `json:"search" form:"search"`
	OrderBy    string   `json:"order_by" form:"order_by"`
	NotGetItem int      `json:"not_get_item" form:"not_get_item"`
	IsCms      int      `json:"is_cms" form:"is_cms"`
}

type TransactionLogFilter struct {
	ItemIds []int64  `json:"item_ids"`
	Cols    []string `json:"cols"`
}

type MarketPlaceFilter struct {
	Id              int64    `json:"id" form:"id"`
	Status          int64    `json:"status" form:"status"`
	ItemId          int64    `json:"item_id" form:"item_id"`
	CategoryId      int64    `json:"category_id" form:"category_id"`
	CollectionId    int64    `json:"collection_id" form:"collection_id"`
	IdOnMarket      int64    `json:"id_on_market" form:"id_on_market"`
	SortBy          string   `json:"sort_by" form:"sort_by"`
	PriceMax        int64    `json:"price_max" form:"price_max"`
	PriceMin        int64    `json:"price_min" form:"price_min"`
	Creator         int64    `json:"creator" form:"creator"`
	Ids             []int64  `json:"ids"`
	ItemIds         []int64  `json:"item_ids"`
	Cols            []string `json:"cols"`
	Price           string   `json:"price" form:"price"`
	CreatedAt       string   `json:"created_at" form:"created_at"`
	TotalLike       string   `json:"total_like" form:"total_like"`
	ContractAddress string   `json:"contract_address"`
	Nft             string   `json:"nft"`
	IsDiscover      int      `json:"is_discover"`
	UserId          int64    `json:"user_id"`
	MarketContract  string   `json:"market_contract"`
	IsHot           string   `json:"is_hot" form:"is_hot"`
	IsTrending      int      `json:"is_trending" form:"is_trending"`
	Unscoped        bool     `json:"unscoped"`
	IsHide          int      `json:"is_hide"`
	SellType        int      `json:"sell_type" form:"sell_type"`
}

type UpdateMarketPlaceRequest struct {
	IsHot        int64 `json:"is_hot"`
	Id           int64 `json:"id"`
	DisplayOrder int64 `json:"display_order"`
}

type ItemEventFilter struct {
	IsShowAll     int    `json:"is_show_all" form:"is_show_all"`
	UserId        int64  `json:"user_id" from:"user_id"`
	CurrentUserId int64  `json:"current_user_id" from:"current_user_id"`
	CategoryId    int64  `json:"category_id" form:"category_id"`
	SortBy        string `json:"sort_by" form:"sort_by"`
	IsHide        int    `json:"is_hide"`
	Search        string `json:"search" form:"search"`
}

type StakingFilter struct {
	Cols          []string `json:"cols"`
	UserId        int64    `json:"user_id"`
	MarketPlaceId int64    `json:"market_place_id"`
	Distinct      []string `json:"distinct"`
}

type VoteFilter struct {
	Cols          []string `json:"cols"`
	UserId        int64    `json:"user_id"`
	MarketPlaceId int64    `json:"market_place_id"`
	Distinct      []string `json:"distinct"`
}

type CollectionLikeFilter struct {
	Cols          []string `json:"cols"`
	Id            int64    `json:"id"`
	Ids           []int64  `json:"ids"`
	UserId        int64    `json:"user_id"`
	IsDeleted     int64    `json:"is_deleted"`
	CollectionId  int64    `json:"collection_id"`
	CollectionIds []int64  `json:"collection_ids"`
}

type MarketPlaceSoldFilter struct {
	ItemIds []int64 `json:"item_ids"`
	MIds    []int64 `json:"m_ids"`
}

type BrandFilter struct {
	IsHot int   `json:"is_hot" form:"is_hot"`
	Id    int64 `json:"id" form:"id"`
}

type ExploreCollectionFilter struct {
	TotalItemMin  int64  `json:"total_item_min" form:"total_item_min"`
	TotalItemMax  int64  `json:"total_item_max" form:"total_item_max"`
	OrderBy       string `json:"order_by" form:"order_by"`
	FloorPriceMin int64  `json:"floor_price_min" form:"floor_price_min"`
	FloorPriceMax int64  `json:"floor_price_max" form:"floor_price_max"`
	Search        string `json:"search" form:"search"`
}

type ExploreNftFilter struct {
	Id           int64  `json:"id" form:"id"`
	UserId       int64  `json:"user_id" form:"user_id"`
	Status       int64  `json:"status" form:"status"`
	CollectionId int64  `json:"collection_id" form:"collection_id"`
	CategoryId   int64  `json:"category_id" form:"category_id"`
	PriceMin     int64  `json:"price_min" form:"price_min"`
	PriceMax     int64  `json:"price_max" form:"price_max"`
	Properties   string `json:"properties" form:"properties"`
	Search       string `json:"search" form:"search"`
	OrderBy      string `json:"order_by" form:"order_by"`
}

type GetAuctionFilter struct {
	Id             int64   `json:"id" form:"id"`
	MarketPlaceId  int64   `json:"market_place_id" form:"market_place_id"`
	MarketPlaceIds []int64 `json:"market_place_ids" form:"market_place_ids"`
	ItemId         int64   `json:"item_id" form:"item_id"`
	ItemIds        []int64 `json:"item_ids" form:"item_id"`
	IsGetUser      int     `json:"is_get_user"`
	OrderBy        string  `json:"order_by" form:"order_by"`
	GroupBy        string  `json:"group_by" form:"group_by"`
}

type NewPackInfo struct {
	IdOnMarket  []int64 `json:"id_on_market"`
	PackId      int64   `json:"pack_id"`
	Transaction string  `json:"transaction"`
	SellType    int     `json:"sell_type"`
}
