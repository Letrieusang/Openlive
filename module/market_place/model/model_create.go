package marketplacemodel

type CreateMarketPlaceRequest struct {
	WalletAddress    string `json:"wallet_address" validate:"required"`
	Nft              string `json:"nft" validate:"required"`
	NftContract      string `json:"nft_contract"`
	IdOnMarket       int64  `json:"id_on_market" validate:"required"`
	TransactionId    string `json:"transaction_id" validate:"required"`
	ContractAddress  string `json:"contract_address" validate:"required"`
	Price            string `json:"price"`
	StartPrice       string `json:"start_price"`
	Description      string `json:"description"`
	SellType         int    `json:"sell_type"`
	AuctionStartTime string `json:"auction_start_time"`
	AuctionEndTime   string `json:"auction_end_time"`
}

type CancelMarketplace struct {
	IdOnMarket      int64 `json:"id_on_market"`
	Nft             string
	ContractAddress string `json:"contract_address"`
	SellType        int    `json:"sell_type"`
	Status          int64  `json:"status" form:"status"`
}

type CreateMarketPlaceSoldRequest struct {
	WalletAddress   string `json:"wallet_address" validate:"required"`
	IdOnMarket      int64  `json:"id_on_market" validate:"required"`
	TransactionId   string `json:"transaction_id" validate:"required"`
	ContractAddress string `json:"contract_address" validate:"required"`
	Price           string `json:"price"`
	Description     string `json:"description"`
	MarketContract  string `json:"market_contract"`
	SellType        int    `json:"sell_type"`
}

type VoteItemRequest struct {
	TransactionId   string `json:"transaction_id" validate:"required"`
	ContractAddress string `json:"contract_address" validate:"required"`
	UserId          int64  `json:"user_id" validate:"required"`
	MarketPlaceId   int64  `json:"market_place_id" validate:"required"`
}

type StakingRequest struct {
	TransactionId   string  `json:"transaction_id" validate:"required"`
	ContractAddress string  `json:"contract_address" validate:"required"`
	WalletAddress   string  `json:"wallet_address" validate:"required"`
	Amount          float64 `json:"amount" validate:"required gt=0"`
	InterestRate    float64 `json:"interest_rate" validate:"required gt=0"`
	TimeDeposit     float64 `json:"time_deposit" validate:"required gt=0"`
}

type LikeItemRequest struct {
	Id     int64 `json:"id" validate:"required"`
	UserId int64 `json:"user_id" validate:"required"`
}

type LikeCollectionRequest struct {
	Id     int64 `json:"id" validate:"required"`
	UserId int64 `json:"user_id" validate:"required"`
}

type CreateCategoryRequest struct {
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type CreateCollectionRequest struct {
	Description     string            `json:"description" validate:"required,min=0,max=1000"`
	Name            string            `json:"name" validate:"required"`
	UserId          int64             `json:"user_id"`
	IsHold          bool              `json:"is_hold"`
	ContractAddress string            `json:"contract_address"`
	TimePublicSale  int64             `json:"time_public_sale"`
	TimeStartHold   int64             `json:"time_start_hold"`
	TimeEndHold     int64             `json:"time_end_hold"`
	Banner          string            `json:"banner" validate:"required"`
	FeaturedImage   string            `json:"featured_image" validate:"required"`
	Logo            string            `json:"logo" validate:"required"`
	MinOpvHolding   string            `json:"min_opv_holding"`
	ItemIds         []int64           `json:"item_ids" validate:"required"`
	Links           map[string]string `json:"links" validate:"required"`
}

type CanCreateCollectionRequest struct {
	Description     string `json:"description"`
	Name            string `json:"name" validate:"required"`
	UserId          int64  `json:"user_id"`
	IsHold          bool   `json:"is_hold"`
	ContractAddress string `json:"contract_address"`
	TimePublicSale  int64  `json:"time_public_sale"`
	TimeStartHold   int64  `json:"time_start_hold"`
	TimeEndHold     int64  `json:"time_end_hold"`
	Banner          string `json:"banner"`
	FeaturedImage   string `json:"featured_image"`
	Logo            string `json:"logo"`
	MinOpvHolding   string `json:"min_opv_holding"`
}

type CreateItemRequest struct {
	Nft             string     `json:"nft"`
	TransactionId   string     `json:"transaction_id"`
	ContractAddress string     `json:"contract_address"`
	Description     string     `json:"description" validate:"required"`
	Title           string     `gorm:"type:VARCHAR(255)" json:"title" validate:"required"`
	Image           string     `gorm:"type:VARCHAR(60)" json:"image" validate:"required"`
	Cover           string     `gorm:"type:VARCHAR(60)" json:"cover"`
	CategoryId      int64      `gorm:"type:INT(11)" json:"category_id" validate:"required"`
	CollectionId    int64      `gorm:"type:INT(11)" json:"collection_id"`
	IsAirDrop       int        `json:"is_air_drop"`
	Property        []Property `json:"property"`
	NumOfCopies     int64      `json:"num_of_copies"`
	IsLocking       bool       `json:"is_locking"`
	Code            string     `json:"code"`
}

type Property struct {
	Name  string      `json:"name" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}

type AuctionBidRequest struct {
	Id            int64  `json:"id,omitempty"`
	MarketPlaceId int64  `json:"market_place_id" validate:"required"`
	UserId        int64  `json:"user_id" validate:"required"`
	Price         string `json:"image" validate:"required"`
}
