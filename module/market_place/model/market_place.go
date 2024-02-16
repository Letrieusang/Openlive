package marketplacemodel

import (
	"api-openlive/common"
	"api-openlive/config"
	usermodel "api-openlive/module/user/model"
	"encoding/json"
	"fmt"
	"time"
)

type MarketPlace struct {
	Id               int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	TransactionId    string         `gorm:"type:VARCHAR(255) NOT NULL" json:"transaction_id" validate:"required"`
	ContractAddress  string         `gorm:"type:VARCHAR(255) NOT NULL" json:"contract_address" validate:"required"`
	IdOnMarket       int64          `gorm:"type:INT(11) NOT NULL" json:"id_on_market" validate:"required"`
	ItemId           int64          `gorm:"type:INT(11) NOT NULL" json:"item_id" validate:"required"`
	Item             Item           `gorm:"migration" json:"item,omitempty"`
	SellType         int            `gorm:"type:TINYINT(4) NOT NULL default 1" json:"sell_type"`
	SellerId         int64          `gorm:"type:INT(11) NOT NULL" json:"seller_id" validate:"required"`
	Seller           usermodel.User `gorm:"migration" json:"seller,omitempty"`
	Price            string         `gorm:"type:VARCHAR(60)" json:"price" validate:"required"`
	StartPrice       string         `gorm:"type:VARCHAR(60) default 0" json:"start_price"`
	AuctionStartTime string         `gorm:"type:VARCHAR(60)" json:"auction_start_time"`
	AuctionEndTime   string         `gorm:"type:VARCHAR(60)" json:"auction_end_time"`
	Status           int64          `gorm:"type:TINYINT(2)" json:"status" validate:"required"`
	TotalVote        int64          `gorm:"type:INT(11) default 0" json:"total_vote"`
	TotalLike        int64          `gorm:"type:INT(11) default 0" json:"total_like"`
	Description      string         `gorm:"type:TEXT" json:"description"`
	IsHot            int            `gorm:"type:TINYINT(2) default 0" json:"is_hot" `
	DisplayOrder     int            `gorm:"type:INT(11)" json:"display_order"`
	PackId           int64          `gorm:"type:INT(11) default 0" json:"pack_id"`
	common.SQLModel  `swaggerignore:"true"`
}

func (MarketPlace) TableName() string { return "market_places" }

type MarketPlaceSold struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	TransactionId   string `gorm:"type:VARCHAR(255) NOT NULL" json:"transaction_id" validate:"required"`
	ContractAddress string `gorm:"type:VARCHAR(255) NOT NULL" json:"contract_address" validate:"required"`
	MarketPlaceId   int64  `gorm:"type:INT(11) NOT NULL" json:"market_place_id" validate:"required"`
	BuyerId         int64  `gorm:"type:INT(11) NOT NULL" json:"buyer_id" validate:"required"`
	Price           string `gorm:"type:VARCHAR(60)" json:"image" validate:"required"`
	Status          int64  `gorm:"type:TINYINT(2)" json:"status" validate:"required"`
	Description     string `gorm:"type:TEXT" json:"description"`
	common.SQLModel `swaggerignore:"true"`
}

func (MarketPlaceSold) TableName() string { return "market_place_solds" }

type ShortInfoMarketPlace struct {
	Id                       int64                `json:"id,omitempty"`
	MarketPlaceSoldId        int64                `json:"market_place_sold_id"`
	SoldTransactionId        string               `json:"sold_transaction_id"`
	ItemId                   int64                `json:"item_id"`
	Item                     Item                 `json:"item,omitempty"`
	SellerId                 int64                `json:"seller_id"`
	Seller                   usermodel.UserDetail `json:"seller,omitempty"`
	BuyerId                  int64                `json:"buyer_id"`
	Buyer                    usermodel.UserDetail `json:"buyer,omitempty"`
	Price                    string               `json:"price"`
	MarketPlaceSoldUpdatedAt string               `json:"market_place_sold_updated_at"`
}

type Vote struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	TransactionId   string         `gorm:"type:VARCHAR(255) NOT NULL" json:"transaction_id" validate:"required"`
	ContractAddress string         `gorm:"type:VARCHAR(255) NOT NULL" json:"contract_address" validate:"required"`
	MarketPlaceId   int64          `gorm:"INT(11) NOT NULL" json:"market_place_id" validate:"required"`
	MarketPlace     MarketPlace    `gorm:"-" json:"market_place,omitempty"`
	UserId          int64          `gorm:"INT(11) NOT NULL" json:"user_id" validate:"required"`
	User            usermodel.User `gorm:"-" json:"user,omitempty"`
	common.SQLModel `swaggerignore:"true"`
}

func (Vote) TableName() string {
	return "votes"
}

type Staking struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	TransactionId   string         `gorm:"type:VARCHAR(255) NOT NULL" json:"transaction_id" validate:"required"`
	ContractAddress string         `gorm:"type:VARCHAR(255) NOT NULL" json:"contract_address" validate:"required"`
	UserId          int64          `gorm:"INT(11) NOT NULL" json:"user_id" validate:"required"`
	User            usermodel.User `gorm:"-" json:"user,omitempty"`
	Status          int64          `gorm:"INT(11) NOT NULL" json:"status" validate:"required"`
	Amount          float64        `gorm:"DOUBLE NOT NULL" json:"amount" validate:"required"`
	InterestRate    float64        `gorm:"DOUBLE NOT NULL" json:"interest_rate" validate:"required"`
	TimeDeposit     float64        `gorm:"DOUBLE NOT NULL" json:"time_deposit" validate:"required"`
	common.SQLModel `swaggerignore:"true"`
}

func (Staking) TableName() string {
	return "stakings"
}

type Like struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	ItemId          int64          `gorm:"INT(11) NOT NULL" json:"item_id" validate:"required"`
	Item            Item           `gorm:"-" json:"item,omitempty"`
	UserId          int64          `gorm:"INT(11) NOT NULL" json:"user_id" validate:"required"`
	User            usermodel.User `gorm:"-" json:"user,omitempty"`
	IsDeleted       int64          `gorm:"type:INT(2);column:is_deleted"`
	common.SQLModel `swaggerignore:"true"`
}

func (Like) TableName() string {
	return "likes"
}

type BlockChainBlock struct {
	Id              uint64 `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	BlockNumber     uint64 `gorm:"INT(11) NOT NULL" json:"block_number" validate:"required"`
	Status          int    `gorm:"TINYINT(4) default 1" json:"status" `
	IsBlockProcess  int    `gorm:"TINYINT(4) default 1" json:"is_block_process" `
	common.SQLModel `swaggerignore:"true"`
}

func (BlockChainBlock) TableName() string {
	return "block_chain_blocks"
}

type Item struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	UserId          int64  `gorm:"type:INT(11)" json:"user_id" validate:"required"`
	Status          int64  `gorm:"type:INT(11)" json:"status" validate:"required"`
	TransactionId   string `gorm:"type:VARCHAR(255)" json:"transaction_id" validate:"required"`
	ContractAddress string `gorm:"type:VARCHAR(255)" json:"contract_address" validate:"required"`
	Title           string `gorm:"type:VARCHAR(255)" json:"title" validate:"required"`
	Nft             string `gorm:"type:VARCHAR(255)" json:"nft" validate:"required"`
	Image           string `gorm:"type:TEXT" json:"image" validate:"required"`
	Cover           string `gorm:"type:TEXT" json:"cover"`
	Description     string `gorm:"type:TEXT" json:"description"`
	CategoryId      int64  `gorm:"type:INT(11)" json:"category_id" validate:"required"`
	CollectionId    int64  `gorm:"type:INT(11)" json:"collection_id" validate:"required"`
	View            int64  `gorm:"type:INT(11) default 0" json:"view" validate:"required"`
	IsApproval      int64  `gorm:"type:INT(2)" json:"is_approval" validate:"required"`
	IsJoinEvent     uint8  `gorm:"type:TINYINT(2) default 0" json:"is_join_event" validate:"required"`
	Properties      string `gorm:"type:TEXT" json:"properties" validate:"required"`
	NumOfCopies     int64  `gorm:"type:INT(11)" json:"num_of_copies" validate:"required"`
	TimeJoinEvent   int64  `gorm:"type:INT(11)" json:"time_join_event" validate:"required"`
	IsLocking       int64  `gorm:"type:INT(11)" json:"is_locking" validate:"required"`
	Code            string `gorm:"type:VARCHAR(255)" json:"code" validate:"required"`
	IsTrending      int    `gorm:"type:TINYINT(2) default 0" json:"is_trending"`
	IsHide          int    `gorm:"type:TINYINT(2) default 0" json:"is_hide"`
	DisplayOrder    int64  `gorm:"tye:INT(10) default 0" json:"display_order"`
	MainItemId      int64  `gorm:"type:INT(11)" json:"main_item_id"`
	TotalLike       int64  `gorm:"type:INT(11) default 0" json:"total_like"`
	common.SQLModel `swaggerignore:"true"`
}

func DefaultItem(userId int64, nft string, contractAddress string) Item {
	commonToken := fmt.Sprintf("#%v", nft)
	return Item{
		UserId:          userId,
		Status:          common.ITEM_NEW,
		ContractAddress: contractAddress,
		Title:           commonToken,
		Nft:             nft,
		Image:           config.GetConfig().GetString("blockchain.default_image"),
		Description:     commonToken,
	}
}

type ItemDetail struct {
	Id              int64       `json:"id,omitempty"`
	UserId          int64       `json:"user_id"`
	Status          int64       `json:"status"`
	TransactionId   string      `json:"transaction_id"`
	ContractAddress string      `json:"contract_address"`
	Title           string      `json:"title"`
	Nft             string      `json:"nft"`
	Image           string      `json:"image"`
	Description     string      `json:"description"`
	CategoryId      int64       `json:"category_id"`
	CollectionId    int64       `json:"collection_id"`
	Collection      *Collection `json:"collection"`
	IsApproval      int64       `json:"is_approval"`
	IsJoinEvent     uint8       `json:"is_join_event"`
	IsLike          int64       `json:"is_like"`
	IsLocking       int64       `json:"is_locking"`
	View            int64       `json:"view"`
	Code            string      `json:"code"`
	IsTrending      int         `json:"is_trending"`
	DisplayOrder    int64       `json:"display_order"`
	Properties      []Property  `json:"properties"`
	NumOfCopies     int64       `json:"num_of_copies"`
	ItemCreatorId   int64       `json:"item_creator_id"`
	IsHide          int         `json:"is_hide"`
	TotalLike       int64       `json:"total_like"`
	MarketPlaceId   int64       `json:"market_place_id"`
	IdOnMarket      string      `json:"id_on_market"`
	common.SQLModel `swaggerignore:"true"`
}

func (i Item) ConvertToDetail() ItemDetail {
	property := make([]Property, 0, 10)
	if len(i.Properties) > 0 {
		_ = json.Unmarshal([]byte(i.Properties), &property)
	}
	return ItemDetail{
		Id:              i.Id,
		UserId:          i.UserId,
		Status:          i.Status,
		TransactionId:   i.TransactionId,
		ContractAddress: i.ContractAddress,
		Title:           i.Title,
		Nft:             i.Nft,
		Image:           i.Image,
		Description:     i.Description,
		CategoryId:      i.CategoryId,
		CollectionId:    i.CollectionId,
		IsApproval:      i.IsApproval,
		IsJoinEvent:     i.IsJoinEvent,
		IsTrending:      i.IsTrending,
		DisplayOrder:    i.DisplayOrder,
		Properties:      property,
		NumOfCopies:     i.NumOfCopies,
		IsLocking:       i.IsLocking,
		View:            i.View,
		IsHide:          i.IsHide,
		TotalLike:       i.TotalLike,
	}
}

func (Item) TableName() string { return "items" }

type Setting struct {
	Id              int64 `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	IsVerifySellNft int64 `gorm:"type:TINYINT(2) default 0" json:"is_verify_sell_nft" `
	common.SQLModel `swaggerignore:"true"`
}

func (Setting) TableName() string { return "settings" }

type Category struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Description     string `gorm:"type:TEXT" json:"description" validate:"required"`
	Name            string `gorm:"type:VARCHAR(255)" json:"name" validate:"required"`
	common.SQLModel `swaggerignore:"true"`
}

func (Category) TableName() string { return "categories" }

type ItemWithContract struct {
	NftItems        []NftItem `json:"nft_items" binding:"required"`
	ContractAddress string    `json:"contract_address" binding:"required"`
}

type CreateItemResult struct {
	TotalItemCreated int64  `json:"total_item_created"`
	ContractAddress  string `json:"contract_address"`
}

type NftItem struct {
	Id string `json:"id" binding:"required"`
}

type ItemCreator struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	UserId          int64          `gorm:"type:INT(11)" json:"user_id"`
	User            usermodel.User `gorm:"-" json:"user"`
	Item            Item           `gorm:"-" json:"item"`
	ItemId          int64          `gorm:"type:INT(11)" json:"item_id"`
	common.SQLModel `swaggerignore:"true"`
}

func (ItemCreator) TableName() string { return "item_creators" }

type Collection struct {
	Id              int64           `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	Description     string          `gorm:"type:TEXT" json:"description"`
	Name            string          `gorm:"type:VARCHAR(255)" json:"name"`
	Logo            string          `gorm:"type:VARCHAR(255)" json:"Logo"`
	FeaturedImage   string          `gorm:"type:VARCHAR(255)" json:"featured_image"`
	Banner          string          `gorm:"type:VARCHAR(255)" json:"banner"`
	UserId          int64           `gorm:"type:INT(11)" json:"user_id"`
	Status          int64           `gorm:"type:INT(11)" json:"status"`
	ContractAddress string          `gorm:"type:VARCHAR(255)" json:"contract_address"`
	TimePublicSale  int64           `gorm:"type:INT(11)" json:"time_public_sale"`
	TimeStartHold   int64           `gorm:"type:INT(11)" json:"time_start_hold"`
	TimeEndHold     int64           `gorm:"type:INT(11)" json:"time_end_hold"`
	IsHold          bool            `gorm:"type:INT(1)" json:"is_hold"`
	IsTrending      int             `gorm:"type:TINYINT(2) default 0" json:"is_trending"`
	DisplayOrder    int             `gorm:"type:TINYINT(2) default 0" json:"display_order"`
	MinOpvHolding   string          `gorm:"type:VARCHAR(255)" json:"min_opv_holding"`
	IsHot           int             `gorm:"type:TINYINT(2) default 0" json:"is_hot"`
	Seller          *usermodel.User `gorm:"-" json:"seller"`
	Links           string          `json:"links"`
	View            int64           `gorm:"type:INT(11) default 0" json:"view" validate:"required"`
	common.SQLModel `swaggerignore:"true"`
}

func (Collection) TableName() string { return "collections" }

type CollectionDetail struct {
	Id              int64                  `json:"id,omitempty"`
	Description     string                 `json:"description"`
	Name            string                 `json:"name"`
	Logo            string                 `json:"Logo"`
	FeaturedImage   string                 `json:"featured_image"`
	Banner          string                 `json:"banner"`
	UserId          int64                  `json:"user_id"`
	Status          int64                  `json:"status"`
	ContractAddress string                 `json:"contract_address"`
	TimePublicSale  int64                  `json:"time_public_sale"`
	TimeStartHold   int64                  `json:"time_start_hold"`
	TimeEndHold     int64                  `json:"time_end_hold"`
	IsHold          bool                   `json:"is_hold"`
	IsHot           int                    `json:"is_hot"`
	Items           []Item                 `json:"items"`
	Seller          *usermodel.UserDetail  `json:"seller"`
	Owner           *usermodel.UserDetail  `json:"owner"`
	IsTrending      int                    `json:"is_trending"`
	DisplayOrder    int                    `json:"display_order"`
	TotalLike       int64                  `json:"total_like"`
	IsLiked         bool                   `json:"is_liked"`
	UserLikes       []usermodel.UserDetail `json:"user_likes"`
	TotalItem       int                    `json:"total_item"`
	View            int64                  `json:"view"`
	TotalOwner      int                    `json:"total_owner"`
	FloorPrice      string                 `json:"floor_price"`
	TotalVolume     string                 `json:"total_volume"`
	MinOPVHolding   string                 `json:"min_opv_holding"`
	Links           map[string]string      `json:"links"`
}

func (c Collection) ConvertToDetail() CollectionDetail {
	linkStr := make(map[string]string)
	_ = json.Unmarshal([]byte(c.Links), &linkStr)
	return CollectionDetail{
		Id:              c.Id,
		Description:     c.Description,
		Name:            c.Name,
		Logo:            c.Logo,
		FeaturedImage:   c.FeaturedImage,
		Banner:          c.Banner,
		UserId:          c.UserId,
		Status:          c.Status,
		ContractAddress: c.ContractAddress,
		TimeEndHold:     c.TimeEndHold,
		TimePublicSale:  c.TimePublicSale,
		TimeStartHold:   c.TimeStartHold,
		IsHold:          c.IsHold,
		IsHot:           c.IsHot,
		MinOPVHolding:   c.MinOpvHolding,
		IsTrending:      c.IsTrending,
		DisplayOrder:    c.DisplayOrder,
		Links:           linkStr,
		View:            c.View,
	}
}

type TransactionLog struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	SellerId        int64  `gorm:"type:INT(11)" json:"seller_id"`
	BuyerId         int64  `gorm:"type:INT(11)" json:"buyer_id"`
	ItemId          int64  `gorm:"type:INT(11)" json:"item_id"`
	Price           string `gorm:"type:VARCHAR(60)" json:"image" validate:"required"`
	TransactionId   string `gorm:"type:VARCHAR(255) NOT NULL" json:"transaction_id" validate:"required"`
	ContractAddress string `gorm:"type:VARCHAR(255) NOT NULL" json:"contract_address" validate:"required"`
	common.SQLModel `swaggerignore:"true"`
}

func (TransactionLog) TableName() string { return "transaction_logs" }

type TransactionLogTopSeller struct {
	SellerId  int64     `json:"seller_id"`
	ItemId    int64     `json:"item_id"`
	CreatedAt time.Time `json:"created_at"`
	Price     string    `json:"price"`
}

type CollectionLikeResponse struct {
	Id           int64 `json:"id"`
	CollectionId int64 `json:"collection_id"`
	TotalLike    int64 `json:"total_like"`
}

type TopCreatorResponse struct {
	UserId    int64     `json:"user_id"`
	Count     int64     `json:"count"`
	CreatedAt time.Time `json:"created_at"`
}

type CollectionLike struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	CollectionId    int64          `gorm:"INT(11) NOT NULL" json:"collection_id" validate:"required"`
	Collection      Collection     `gorm:"-" json:"collection,omitempty"`
	UserId          int64          `gorm:"INT(11) NOT NULL" json:"user_id" validate:"required"`
	User            usermodel.User `gorm:"-" json:"user,omitempty"`
	IsDeleted       int64          `gorm:"type:INT(2);column:is_deleted"`
	common.SQLModel `swaggerignore:"true"`
}

func (CollectionLike) TableName() string {
	return "collection_likes"
}

type MarketPlaceWithCollectionName struct {
	Id               int64          `json:"id,omitempty"`
	TransactionId    string         `json:"transaction_id"`
	ContractAddress  string         `json:"contract_address"`
	IdOnMarket       int64          `json:"id_on_market"`
	ItemId           int64          `json:"item_id"`
	Item             Item           `json:"item,omitempty"`
	SellerId         int64          `json:"seller_id"`
	Seller           usermodel.User `json:"seller,omitempty"`
	Price            string         `json:"price"`
	Status           int64          `json:"status"`
	TotalVote        int64          `json:"total_vote"`
	TotalLike        int64          `json:"total_like"`
	Description      string         `json:"description"`
	CollectionName   string         `json:"collection_name"`
	SellType         int            `json:"sell_type"`
	StartPrice       string         `json:"start_price"`
	IsHot            int            `json:"is_hot"`
	DisplayOrder     int            `json:"display_order"`
	AuctionStartTime string         `json:"auction_start_time"`
	AuctionEndTime   string         `json:"auction_end_time"`
	PackId           int64          `json:"pack_id"`
	common.SQLModel  `swaggerignore:"true"`
}

type ItemJoinEventHistory struct {
	Id              int     `gorm:"type:INT(11) AUTO_INCREMENT;primarykey;column:id" json:"id,omitempty" swaggerignore:"true"`
	UserJoinEventId int64   `gorm:"type:INT(11)" json:"user_join_event_id"`
	ItemId          int64   `gorm:"type:INT(11)" json:"item_id"`
	Status          int     `gorm:"type:TINYINT(2) default 1" json:"status"`
	Profit          float64 `gorm:"type:DOUBLE" json:"profit"`
	Revenue         string  `gorm:"type:VARCHAR(255) default 0" json:"revenue"`
	common.SQLModel `swaggerignore:"true"`
}

func (ItemJoinEventHistory) TableName() string { return "item_join_event_histories" }

type Brand struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	IsHot           int    `gorm:"type:TINYINT(2) default 0" json:"is_hot"`
	BrandTitle      string `gorm:"type:VARCHAR(255)" json:"brand_title"`
	Description     string `gorm:"type:TEXT" json:"description"`
	DisplayOrder    int    `gorm:"type:INT(4) default 0" json:"display_order"`
	BackgroundImage string `gorm:"type:TEXT" json:"background_image"`
	common.SQLModel `swaggerignore:"true"`
}

func (Brand) TableName() string { return "brands" }

type BrandImage struct {
	Id                  int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	BrandName           string `gorm:"type:VARCHAR(255)" json:"brand_name"`
	Link                string `gorm:"type:TEXT" json:"link"`
	BrandId             int64  `gorm:"type:INT(11)" json:"brand_id"`
	ItemBackgroundImage string `gorm:"type:TEXT" json:"item_background_image"`
	ItemAvatarImage     string `gorm:"type:TEXT" json:"item_avatar_image"`
	DisplayOrder        int    `gorm:"type:INT(4) default 0" json:"display_order"`
	common.SQLModel     `swaggerignore:"true"`
}

func (BrandImage) TableName() string { return "brand_image" }

type BrandDetail struct {
	Id              int64              `json:"id,omitempty"`
	IsHot           int                `json:"is_hot"`
	BrandTitle      string             `json:"brand_title"`
	Description     string             `json:"description"`
	BackgroundImage string             `json:"background_image"`
	DisplayOrder    int                `json:"display_order"`
	Images          []BrandImageDetail `json:"images"`
}

type BrandImageDetail struct {
	Id                  int64  `json:"id,omitempty"`
	BrandName           string `json:"brand_name"`
	Link                string `json:"link"`
	BrandId             int64  `json:"brand_id"`
	ItemBackgroundImage string `json:"item_background_image"`
	ItemAvatarImage     string `json:"item_avatar_image"`
	DisplayOrder        int    `json:"display_order"`
}

type CreateBrandCmsRequest struct {
	IsHot           int                          `json:"is_hot"`
	BrandTitle      string                       `json:"brand_title"`
	Description     string                       `json:"description"`
	BackgroundImage string                       `json:"background_image"`
	DisplayOrder    int                          `json:"display_order"`
	BrandImages     []CreateBrandImageCmsRequest `json:"brand_images"`
}

func (c CreateBrandCmsRequest) ConvertToBrand() Brand {
	return Brand{
		BrandTitle:      c.BrandTitle,
		IsHot:           c.IsHot,
		Description:     c.Description,
		BackgroundImage: c.BackgroundImage,
		DisplayOrder:    c.DisplayOrder,
	}
}

func (c CreateBrandCmsRequest) ConvertToBrandImage() []BrandImage {
	var result []BrandImage
	for _, value := range c.BrandImages {
		var tmp BrandImage
		tmp.BrandName = value.BrandName
		tmp.ItemAvatarImage = value.ItemAvatarImage
		tmp.Link = value.Link
		tmp.ItemBackgroundImage = value.ItemBackGroundImage
		tmp.DisplayOrder = value.DisplayOrder
		result = append(result, tmp)
	}
	return result
}

type CreateBrandImageCmsRequest struct {
	BrandName           string `json:"brand_name"`
	Link                string `json:"link"`
	BrandId             int64  `json:"brand_id"`
	ItemBackGroundImage string `json:"item_background_image"`
	ItemAvatarImage     string `json:"item_avatar_image"`
	DisplayOrder        int    `json:"display_order"`
}

type UpdateBrandCmsRequest struct {
	Ishot           int64  `json:"is_hot"`
	Id              int64  `json:"id,omitempty"`
	BrandTitle      string `json:"brand_title"`
	Description     string `json:"description"`
	BackgroundImage string `json:"background_image"`
	DisplayOrder    int64  `json:"display_order"`
}

type UpdateBrandImageCmsRequest struct {
	Id                  int64  `json:"id,omitempty"`
	BrandName           string `json:"brand_name"`
	Link                string `json:"link"`
	ItemBackgroundImage string `json:"item_background_image"`
	ItemAvatarImage     string `json:"item_avatar_image"`
	DisplayOrder        int    `json:"display_order"`
}

func (c Brand) ConvertToDetail() BrandDetail {
	return BrandDetail{
		Id:              c.Id,
		BrandTitle:      c.BrandTitle,
		IsHot:           c.IsHot,
		Description:     c.Description,
		BackgroundImage: c.BackgroundImage,
		DisplayOrder:    c.DisplayOrder,
	}
}

func (c BrandImage) ConvertToDetail() BrandImageDetail {
	return BrandImageDetail{
		Id:                  c.Id,
		BrandName:           c.BrandName,
		Link:                c.Link,
		BrandId:             c.BrandId,
		ItemBackgroundImage: c.ItemBackgroundImage,
		ItemAvatarImage:     c.ItemAvatarImage,
		DisplayOrder:        c.DisplayOrder,
	}
}

type AddBannerCmsRequest struct {
	Title           string `gorm:"type:VARCHAR(255)" json:"title"`
	BlueTitle       string `gorm:"type:VARCHAR(255)" json:"blue_title"`
	Line1           string `gorm:"type:VARCHAR(255)" json:"line1"`
	Line2           string `gorm:"type:VARCHAR(255)" json:"line2"`
	Line3           string `gorm:"type:VARCHAR(255)" json:"line3"`
	Line4           string `gorm:"type:VARCHAR(255)" json:"line4"`
	common.SQLModel `swaggerignore:"true"`
}

func (AddBannerCmsRequest) TableName() string { return "banner" }

type BannerDetail struct {
	Title     string `json:"title"`
	BlueTitle string `json:"blue_title"`
	Line1     string `json:"line1"`
	Line2     string `json:"line2"`
	Line3     string `json:"line3"`
	Line4     string `json:"line4"`
}

func (c AddBannerCmsRequest) ConvertToDetail() BannerDetail {
	return BannerDetail{
		Title:     c.Title,
		BlueTitle: c.BlueTitle,
		Line1:     c.Line1,
		Line2:     c.Line2,
		Line3:     c.Line3,
		Line4:     c.Line4,
	}
}

type AuctionBid struct {
	Id              int64          `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	MarketPlaceId   int64          `gorm:"type:INT(11) NOT NULL" json:"market_place_id" validate:"required"`
	UserId          int64          `gorm:"type:INT(11) NOT NULL" json:"user_id" validate:"required"`
	Price           string         `gorm:"type:VARCHAR(60)" json:"price" validate:"required"`
	UserInfo        usermodel.User `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:Id"`
	common.SQLModel `swaggerignore:"true"`
}

func (AuctionBid) TableName() string { return "auction_bids" }

type ReportTableDetail struct {
	Id              int64  `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	ItemId          int64  `gorm:"type:INT(11) NOT NULL" json:"item_id"`
	UserId          int64  `gorm:"type:INT(11) NOT NULL" json:"user_id"`
	ReportReason    string `gorm:"type:VARCHAR(255)" json:"report_reason"`
	ReportMessage   string `gorm:"type:VARCHAR(255)" json:"report_message"`
	IsHide          int    `gorm:"type:INT(4) default 0" json:"is_hide"`
	IsReject        int    `gorm:"type:INT(4) default 0" json:"is_reject"`
	HideMessage     string `gorm:"type:VARCHAR(255)" json:"hide_message"`
	RejectMessage   string `gorm:"type:VARCHAR(255)" json:"reject_message"`
	common.SQLModel `swaggerignore:"true"`
}

func (ReportTableDetail) TableName() string { return "table_reports" }

type ReportTableResponse struct {
	Id            int64                `json:"id"`
	ItemDetail    Item                 `json:"item_id"`
	UserDetail    usermodel.UserDetail `json:"user"`
	ReportReason  string               `json:"report_reason"`
	ReportMessage string               `json:"report_message"`
	IsHide        int                  `json:"is_hide"`
	IsReject      int                  `json:"is_reject"`
	HideMessage   string               `json:"hide_message"`
	RejectMessage string               `json:"reject_message"`
	CreatedAt     time.Time            `json:"created_at"`
}

func (c ReportTableDetail) ConvertToDetail() ReportTableResponse {
	return ReportTableResponse{
		Id:            c.Id,
		ReportMessage: c.ReportMessage,
		ReportReason:  c.ReportReason,
		IsHide:        c.IsHide,
		IsReject:      c.IsReject,
		HideMessage:   c.HideMessage,
		RejectMessage: c.RejectMessage,
		CreatedAt:     c.CreatedAt,
	}
}

type ChangeReportRequest struct {
	Id            int64  `json:"id" form:"id"`
	IsHide        int    `json:"is_hide" form:"is_hide"`
	IsReject      int    `json:"is_reject" form:"is_reject"`
	RejectMessage string `json:"reject_message" form:"reject_message"`
	HideMessage   string `json:"hide_message" form:"hide_message"`
}

type RequestReport struct {
	UserId        int64  `json:"user_id" form:"user_id"`
	ItemId        int64  `json:"item_id" form:"item_id"`
	ReportReason  string `json:"report_reason" gorm:"report_reason"`
	ReportMessage string `json:"report_message" gorm:"report_message"`
}
