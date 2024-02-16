package marketplacemodel

import (
	usermodel "api-openlive/module/user/model"
	"time"
)

type UpdateItemOwnerRequest struct {
	SenderId        int64
	ReceiverId      int64
	Nft             string
	ContractAddress string
}

type c struct {
	SenderId        int64
	ReceiverId      int64
	Nft             string
	ContractAddress string
	Type            int64
}

type CollectionWatchList struct {
	Id             int64   `json:"id"`
	Logo           string  `json:"logo"`
	Banner         string  `json:"banner"`
	FeaturedImage  string  `json:"featured_image"`
	CollectionName string  `json:"collection_name"`
	FloorPrice     float64 `json:"floor_price"`
	TotalVolume    float64 `json:"total_volume"`
	TotalOwner     int64   `json:"total_owner"`
	TotalItem      int64   `json:"total_item"`
	USDTVolume     float64 `json:"usdt_volume"`
	PercentVolume  float64 `json:"percent_volume"`
	USDTFloorPrice float64 `json:"usdt_floor_price"`
}

type TopCollection struct {
	Collection
	SoldPrice  string `json:"sold_price"`
	FloorPrice string `json:"floor_price"`
	MStatus    int64  `json:"m_status"`
}

type TopVolume struct {
	Collection
	SoldPrice    float64   `json:"sold_price"`
	TransCreated time.Time `json:"trans_created"`
}

type MapFloorPrice struct {
	Collection
	FloorPrice float64 `json:"floor_price"`
}

type TopCollectionResponse struct {
	Collection
	Volume         float64 `json:"volume"`
	FloorPrice     float64 `json:"floor_price"`
	USDTVolume     float64 `json:"usdt_volume"`
	PercentVolume  float64 `json:"percent_volume"`
	USDTFloorPrice float64 `json:"usdt_floor_price"`
}

type CollectionStatsResponse struct {
	AvgPrice  float64       `json:"avg_price"`
	AvgVolume float64       `json:"avg_volume"`
	Detail    []StatsDetail `json:"detail"`
}

type StatsDetail struct {
	TotalNft    int64   `json:"total_nft"`
	TotalVolume float64 `json:"total_volume"`
	AvgPrice    float64 `json:"avg_price"`
	Time        string  `json:"time"`
}

type ExploreNftQueryResponse struct {
	Price float64 `json:"price"`
	Item
	IsLike int64 `json:"is_like"`
}

type ExploreNftResponse struct {
	Price float64 `json:"price"`
	ItemDetail
}

type BidItemDetailResponse struct {
	Id            int64                `gorm:"type:INT(11) AUTO_INCREMENT;primarykey" json:"id,omitempty"`
	MarketPlaceId int64                `gorm:"type:INT(11) NOT NULL" json:"market_place_id" validate:"required"`
	UserId        int64                `gorm:"type:INT(11) NOT NULL" json:"user_id" validate:"required"`
	Price         string               `gorm:"type:VARCHAR(60)" json:"price" validate:"required"`
	UserInfo      usermodel.UserDetail `json:"user_info,omitempty" gorm:"foreignKey:UserId;references:Id"`
	USDTPrice     float64              `json:"usdt_price"`
}
