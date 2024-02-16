package usermodel

type ChangePasswordRequest struct {
	Id          int    `json:"id"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

type ResetPasswordRequest struct {
	NewPassword string `json:"newPassword" validate:"required"`
	Otp         string `json:"otp" validate:"required"`
}

type UpdateUserRequest struct {
	Id          int64  `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Avatar      string ` json:"avatar,omitempty"`
	Cover       string ` json:"cover,omitempty"`
	DisplayName string ` json:"display_name,omitempty"`
	Description string ` json:"description,omitempty"`
	FbLink      string `json:"fb_link,omitempty"`
	TwitterLink string `json:"twitter_link,omitempty"`
	InsLink     string `json:"ins_link,omitempty"`
	LinkInLink  string `json:"link_in_link,omitempty"`
	RefCode     string `json:"ref_code"`
}

type UpdateUserStatus struct {
	Status int64 `json:"status"`
	Id     int64 `json:"id"`
}

type UpdateNotificationStatus struct {
	Id     int64 `json:"id" validate:"required"`
	UserId int64 `json:"user_id"`
}

type SearchRequest struct {
	IsJustGetNft     int        `json:"is_just_get_nft"`
	Value            string     `json:"value" form:"value"`
	UserId           int64      `json:"user_id" form:"user_id"`
	CategoryIds      string     `json:"category_ids" form:"category_ids"`
	PriceMax         float64    `json:"price_max" form:"price_max"`
	PriceMin         float64    `json:"price_min" form:"price_min"`
	Status           string     `json:"status" form:"status"`
	IsJoinEvent      int64      `json:"is_join_event" form:"is_join_event"`
	OrderBy          int64      `json:"order_by" form:"order_by"`
	CollectionId     int64      `json:"collection_id" form:"collection_id"`
	SellType         int64      `json:"sell_type" form:"sell_type"`
	Properties       []Property `json:"properties" form:"properties"`
	OrNotListItemIds []int64    `json:"or_not_list_item_ids"`
}

type Property struct {
	Name  string      `json:"name" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}
