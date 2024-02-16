package usermodel

import (
	"api-openlive/common"
	"encoding/json"
	"github.com/btcsuite/btcutil/base58"
	"log"
	"time"
)

type NounceResponse struct {
	Nounce string `json:"nounce"`
	UserId int64  `json:"user_id"`
}

type Credential struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type UserDetailResponse struct {
	User
	Follower    int64   `json:"follower"`
	Following   int64   `json:"following"`
	IsFollowing int64   `json:"is_following"`
	StakingOpv  float64 `json:"staking_opv"`
	UsedRefCode string  `json:"used_ref_code"`
}

type UserResponse struct {
	Id            int64  `json:"id,omitempty"`
	LoginName     string `json:"username,omitempty"`
	Email         string `json:"email,omitempty"`
	Phone         string `json:"phone,omitempty"`
	DisplayName   string `json:"display_name,omitempty"`
	Avatar        string `json:"avatar,omitempty"`
	WalletAddress string `json:"wallet_address,omitempty"`
	Ref           string `json:"ref,omitempty"`
}

func UserDto(user *User) UserResponse {
	dto := UserResponse{
		Id:            user.Id,
		Email:         user.Email,
		Phone:         user.Phone,
		Avatar:        user.Avatar,
		WalletAddress: user.WalletAddress,
		DisplayName:   user.DisplayName,
	}
	return dto
}

type NotificationDTO struct {
	Id               int64                  `json:"id"`
	Status           int64                  `json:"status"`
	Message          string                 `json:"message"`
	LastTimeRead     int64                  `json:"last_time_read"`
	Type             int64                  `json:"type"`
	NotificationType int64                  `json:"notification_type"`
	CreatedAt        time.Time              `json:"created_at"`
	UpdatedAt        time.Time              `json:"updated_at"`
	MetaData         map[string]interface{} `json:"meta_data"`
}

func convertMessage(data string) map[string]interface{} {
	var notiObj map[string]interface{}
	dataObject := base58.Decode(data)
	err := json.Unmarshal(dataObject, &notiObj)
	if err != nil {
		log.Println("err when unmarshal ", err.Error())
		return nil
	}

	return notiObj
}

func (n Notification) ConvertNotificationDTO() NotificationDTO {
	metaData := convertMessage(n.NotificationObject)
	notiType := 0
	switch n.NotificationType {
	case common.NOTIFICATION_LISTED, common.NOTIFICATION_UNLISTED, common.PURCHASED, common.CREATED_COLLECTION, common.MINTED, common.NOTIFICATION_AUCTION_OTHER_PURCHASED:
		notiType = common.FOLLOWING_USER_ACTIVITIES
	case common.FOLLOWED:
		notiType = common.FOLLOWER_USER_ACTIVITIES
	case common.AUCTION_WIN, common.AUCTION_BEATEN, common.AUCTION_ACCEPTED, common.NOTIFICATION_AUCTION_CLAIM_NFT:
		notiType = common.AUCTION_ACTIVITIES
	case common.LIKED:
		notiType = common.OTHER_USER_ACTIVITIES
	case common.ORTHER_USER_SELL:
		notiType = common.SELL_ACTIVITIES
	case common.SELL, common.CANCEL, common.BUY, common.BID, common.USER_CREATE_AUCTION, common.JOIN_EVENT, common.REMOVE_FROM_EVENT, common.NOTIFICATION_AUCTION_BID, common.NOTIFICATION_AUCTION_CANCEL, common.NOTIFICATION_AUCTION:
		notiType = common.USERS_ACTIVITIES
	}

	return NotificationDTO{
		Id:               n.Id,
		Status:           n.Status,
		LastTimeRead:     n.LastTimeRead,
		Type:             int64(notiType),
		CreatedAt:        n.CreatedAt,
		UpdatedAt:        n.UpdatedAt,
		MetaData:         metaData,
		NotificationType: n.NotificationType,
	}
}

type UserCreateResponse struct {
	*User
	RefWalletAddress string `json:"ref_wallet_address"`
}

type UserUpdateResponse struct {
	*UserDetail
	RefWalletAddress string `json:"ref_wallet_address"`
}

type UserSearchResponse struct {
	Id            int64  `json:"id"`
	Avatar        string `json:"avatar"`
	DisplayName   string `json:"display_name"`
	Type          int64  `json:"type"`
	Followers     int64  `json:"followers"`
	WalletAddress string `json:"wallet_address"`
}

type UserFollowSearchResponse struct {
	UserFollowId int64 `json:"user_follow_id"`
	UserId       int64 `json:"user_id"`
}

type CollectionSearchResponse struct {
	Id            int64  `json:"id"`
	View          int64  `json:"view"`
	Logo          string `json:"logo"`
	Name          string `json:"name"`
	FeaturedImage string `json:"featured_image"`
}

type NFTSearchResponse struct {
	Id               int64      `json:"id"`
	UserId           int64      `json:"user_id"`
	IsLiked          bool       `json:"is_like"`
	Owner            UserDetail `json:"owner" gorm:"-"`
	ContractAddress  string     `json:"contract_address"`
	Nft              string     `json:"nft"`
	Image            string     `json:"image"`
	CollectionName   string     `json:"collection_name"`
	Title            string     `json:"title"`
	Price            string     `json:"price"`
	TotalLike        int64      `json:"total_like"`
	Status           int64      `json:"status"`
	IsJoinEvent      int64      `json:"is_join_event"`
	View             int64      `json:"view"`
	TransactionId    string     `json:"transaction_id"`
	NumOfCopies      int64      `json:"num_of_copies" gorm:"num_of_copies"`
	AuctionStartTime string     `json:"auction_start_time" gorm:"auction_end_time"`
	AuctionEndTime   string     `json:"auction_end_time"  gorm:"auction_end_time"`
	AvailableCopy    int64      `json:"available_copy"`
	MainItemId       int64      `json:"main_item_id"`
	PackId           int64      `json:"pack_id"`
}

type SearchResponse struct {
	Id        int64  `json:"id"`
	UserId    int64  `json:"user_id"`
	Search    string `json:"search"`
	CreatedAt string `json:"created_at"`
}

func SearchResponseDto(s []*UserSearch) []SearchResponse {
	result := make([]SearchResponse, 0, len(s))
	for _, v := range s {
		result = append(result, SearchResponse{
			Id:        v.Id,
			UserId:    v.UserId,
			Search:    v.Search,
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return result
}

type ListFollowQueryResponse struct {
	UserId         int64 `json:"user_id"`
	FollowerId     int64 `json:"follower_id"`
	IsFollowing    bool  `json:"is_following"`
	TotalFollowers int64 `json:"total_followers"`
}

type ListFollowResponse struct {
	Id             int64  `json:"id"`
	DisplayName    string `json:"display_name"`
	WalletAddress  string `json:"wallet_address"`
	Avatar         string `json:"avatar"`
	IsFollowing    bool   `json:"is_following"`
	TotalFollowers int64  `json:"total_followers"`
}
