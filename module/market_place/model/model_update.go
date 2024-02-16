package marketplacemodel

type ChangeMPStatusRequest struct {
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}

type UpdateStakingStatusRequest struct {
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}

type UpdateCategoryRequest struct {
	Id          int64  `json:"id,omitempty"`
	Description string `json:"description" validate:"required"`
	Name        string `json:"name" validate:"required"`
}

type UpdateItemRequest struct {
	Id           int64 `json:"id"`
	CollectionId int64 `json:"collection_id"`
	UserId       int64 `json:"user_id"`
	Status       int64 `json:"status"`
	IsApproved   int64 `json:"is_approved"`
}

type UpdateItemStatusRequest struct {
	Id     []int64 `json:"id" validate:"required"`
	Status int64   `json:"status" validate:"required"`
}

type UpdateItemCmsRequest struct {
	Id           int64 `json:"id" validate:"required"`
	Status       int64 `json:"status"`
	IsTrending   int   `json:"is_trending"`
	DisplayOrder int64 `json:"display_order"`
	UserId       int64 `json:"user_id"`
}

type UpdateCollectionRequest struct {
	Id            int64             `json:"id,omitempty"`
	Description   string            `json:"description"`
	Name          string            `json:"name"`
	Logo          string            `json:"Logo"`
	FeaturedImage string            `json:"featured_image"`
	Banner        string            `json:"banner"`
	UserId        int64             `json:"user_id"`
	MinOpvHolding string            `json:"min_opv_holding"`
	ItemIds       []int64           `json:"item_ids"`
	Links         map[string]string `json:"links"`
}

type UpdateCollectionCmsRequest struct {
	Id           int64 `json:"id,omitempty"`
	IsHot        int   `json:"is_hot"`
	IsTrending   int   `json:"is_trending"`
	DisplayOrder int64 `json:"display_order"`
}

type HideNftRequest struct {
	IsAdmin int   `json:"is_admin"`
	Id      int64 `json:"id,omitempty"`
	UserId  int64 `json:"user_id"`
}
