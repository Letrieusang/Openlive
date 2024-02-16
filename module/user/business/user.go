package business

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	marketplacestorage "api-openlive/module/market_place/storage"
	usermodel "api-openlive/module/user/model"
	userstorage "api-openlive/module/user/storage"
	"api-openlive/utils"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	common_eth "github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/twinj/uuid"
	"gorm.io/gorm"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	Store              userstorage.Storage
	MarketPlaceStorage marketplacestorage.Storage
	Cache              *redis.Client
}

func (u UserController) CreateUser(rq usermodel.CreateUserRequest) (*usermodel.UserCreateResponse, error) {
	// clear otp code used
	exist, err := u.Store.CheckExistUser(&usermodel.CreateUserRequest{WalletAddress: rq.WalletAddress})
	if err != nil {
		return &usermodel.UserCreateResponse{}, err
	}

	if exist {
		return &usermodel.UserCreateResponse{}, common.ErrInUsed
	}
	user := &usermodel.User{
		WalletAddress: rq.WalletAddress,
	}
	if err = u.Store.CreateUser(user); err != nil {
		return &usermodel.UserCreateResponse{}, err
	}

	//refCode := utils.GetMD5Hash(fmt.Sprintf("%v.%v", user.WalletAddress, user.Id))
	//user.Ref = utils.BuildRefLink(refCode)
	user.Ref = user.WalletAddress
	err = u.Store.UpdateUser(user)
	if err != nil {
		return &usermodel.UserCreateResponse{}, err
	}
	response := &usermodel.UserCreateResponse{User: user}

	if len(rq.RefCode) > 0 {
		info, err := u.Store.GetUserByCondition(usermodel.UserFilter{Ref: rq.RefCode})
		if err != nil && err != gorm.ErrRecordNotFound {
			return &usermodel.UserCreateResponse{}, err
		}

		if err == gorm.ErrRecordNotFound {
			return &usermodel.UserCreateResponse{}, errors.New("not found ref code")
		}

		refUser := &usermodel.RefUser{
			UserId:    user.Id,
			RefCode:   rq.RefCode,
			Profit:    common.DEFAULT_PROFIT,
			UserRefId: info.Id,
		}

		err = u.Store.CreateRefUser(refUser)
		if err != nil {
			return nil, err
		}
		response.RefWalletAddress = info.WalletAddress
	}

	setting := &usermodel.NotificationSetting{
		UserId:              user.Id,
		ItemSold:            1,
		BiddingActivities:   1,
		AuctionExpiration:   1,
		OutBid:              1,
		NewLetter:           1,
		FollowingActivities: 1,
		LikeAndFollow:       1,
		AuctionWin:          1,
		YourActivities:      1,
	}
	if errSetting := u.Store.CreateNotificationSetting(setting); errSetting != nil {
		return nil, errSetting
	}

	return response, nil

}

func (u UserController) DetailUser(rq usermodel.UserFilter) (*usermodel.UserDetailResponse, error) {
	var userFilter usermodel.UserFilter
	if rq.Id > 0 {
		userFilter.Id = rq.Id
	}
	if len(rq.WalletAddress) > 0 {
		userFilter.WalletAddress = rq.WalletAddress
	}

	info, err := u.Store.GetUserByCondition(userFilter)
	if err != nil {
		return nil, err
	}
	if info.Id < 1 {
		return nil, errors.New("User not exist")
	}

	ref, err := u.Store.GetRefUserByUserId(rq.Id)
	if err != nil {
		return nil, err
	}

	follower, err := u.Store.CountFollowers(rq.Id)
	if err != nil {
		return nil, err
	}

	following, err := u.Store.CountFollowing(rq.Id)
	if err != nil {
		return nil, err
	}

	isFollowing, err := u.Store.CheckFollowing(rq.CurrentId, rq.Id)
	if err != nil {
		return nil, err
	}

	info.Ref = utils.BuildRefLink(info.Ref)
	info.Avatar = utils.RebuildUrlPath(info.Avatar)
	info.Cover = utils.RebuildUrlPath(info.Cover)
	userDetail := &usermodel.UserDetailResponse{
		User:        info,
		UsedRefCode: ref.RefCode,
		Follower:    follower,
		Following:   following,
		IsFollowing: isFollowing,
	}

	//opv, err := u.MarketPlaceStorage.FindStaking(&marketplacemodel.Staking{UserId: info.Id})
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return nil, err
	//}
	//
	//if opv != nil {
	//	userDetail.StakingOpv = opv.Amount
	//}

	return userDetail, nil

}

func (u UserController) CanAddRefProcess(rq usermodel.UserFilter) (bool, error) {
	info, err := u.Store.CheckCanAddRefUser(&usermodel.CreateUserRequest{UserId: rq.Id})
	if err != nil {
		return false, err
	}
	return !info, err
}

func (u UserController) LoginUser(rq usermodel.UserLoginRequest) (*usermodel.Credential, error) {
	nonce, err := u.Cache.Get(context.Background(), rq.WalletAddress).Result()
	if err != nil || nonce == "" {
		return nil, common.ErrNonceInvalid
	}
	u.Cache.Del(context.Background(), rq.WalletAddress)
	if os.Getenv("ENV") != "local" {
		if !utils.VerifySig(rq, nonce) {
			return nil, common.ErrVerify
		}
	}

	info, err := u.Store.GetUserByCondition(usermodel.UserFilter{WalletAddress: rq.WalletAddress})
	if err != nil {
		return nil, err
	}

	if info.Id == 0 {
		return nil, common.ErrNotFoundUser
	}

	res, errCreateToken := utils.CreateToken(info.Id, common.USER_TYPE)
	if errCreateToken != nil {
		return nil, errCreateToken
	}
	u.BuildUserDailyEvent(info.Id, common.DAILY_EVENT_TYPE_LOGIN)
	return res, nil
}

func (u UserController) RefreshToken(userId int64) (*usermodel.Credential, error) {
	if userId == 0 {
		return nil, common.ErrNotFoundUser
	}

	res, errCreateToken := utils.CreateToken(userId, common.USER_TYPE)
	if errCreateToken != nil {
		return nil, errCreateToken
	}

	return res, nil
}

func (u UserController) CheckExistUser(walletAddress string) (bool, error) {
	if len(walletAddress) == 0 {
		return false, common.ErrNotFoundUser
	}

	user, err := u.Store.GetUserByCondition(usermodel.UserFilter{WalletAddress: walletAddress})
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.Id == 0 {
		return false, nil
	}

	return true, nil
}

func (u UserController) CreateNounce(rq usermodel.NonceRequest) (*usermodel.NounceResponse, error) {
	userInfo, err := u.Store.GetUserByCondition(usermodel.UserFilter{
		WalletAddress: rq.WalletAddress,
	})
	if err != nil {
		return nil, err
	}

	nounce := uuid.NewV4().String()
	if stt := u.Cache.Set(context.Background(), rq.WalletAddress, nounce, 30*time.Minute); stt.Err() != nil {
		return nil, err
	}
	return &usermodel.NounceResponse{
		Nounce: nounce,
		UserId: userInfo.Id,
	}, nil

}

func (u UserController) MailOtp(rq usermodel.OTPRequest) error {
	nonce, err := u.Cache.Get(context.Background(), rq.WalletProvider).Result()
	if err != nil || nonce == "" || nonce != rq.Nonce {
		return common.ErrNonceInvalid
	}

	//gen OTP
	otp, err := u.setOTP("")
	if err != nil {
		return err
	}

	subject := "OPVLive OTP"
	if errSendMail := utils.SendMail(subject, otp, []string{rq.Email}); errSendMail != nil {
		return errSendMail
	}
	return nil
}

func (u UserController) MailOtpReset(rq usermodel.OTPRequestReset) error {
	user, err := u.Store.GetUserByCondition(usermodel.UserFilter{WalletAddress: rq.WalletProvider})
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return common.ErrNotFoundUser
	}

	//gen OTP
	otp, err := u.setOTP(user.Id)
	if err != nil {
		return err
	}

	subject := "OPVLive OTP Reset Password"
	if errSendMail := utils.SendMail(subject, otp, []string{user.Email}); errSendMail != nil {
		return errSendMail
	}
	return nil
}

func (u UserController) setOTP(value interface{}) (string, error) {
	otp := utils.GenRandomOTP()
	var e error
	for e != nil {
		_, e = u.Cache.Get(context.Background(), otp).Result()
		otp = utils.GenRandomOTP()
	}
	if value == "" {
		value = otp
	}
	_, err := u.Cache.Set(context.Background(), otp, value, common.USER_OTP_DURATION*time.Minute).Result()
	return otp, err
}

func (u UserController) ChangeProfile(rq *usermodel.UpdateUserRequest, options map[string]interface{}) (*usermodel.UserUpdateResponse, error) {
	user := &usermodel.User{
		Id:          rq.Id,
		Phone:       rq.Phone,
		Description: rq.Description,
		DisplayName: rq.DisplayName,
		LinkInLink:  rq.LinkInLink,
		FbLink:      rq.FbLink,
		TwitterLink: rq.TwitterLink,
		InsLink:     rq.InsLink,
	}

	cover := utils.DESTINATION + utils.GetFileName(rq.Cover)
	avatar := utils.DESTINATION + utils.GetFileName(rq.Avatar)

	if len(rq.Cover) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Cover, cast.ToString(options["host"]), ""), cover); err != nil {
			return nil, err
		}
	}

	if len(rq.Avatar) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Avatar, cast.ToString(options["host"]), ""), avatar); err != nil {
			return nil, err
		}
	}

	if cover != utils.DESTINATION {
		user.Cover = cast.ToString(options["host"]) + cover
	}

	if avatar != utils.DESTINATION {
		user.Avatar = cast.ToString(options["host"]) + avatar
	}

	if err := u.Store.UpdateUser(user); err != nil {
		return nil, err
	}

	infoUpdated, err := u.Store.GetUserDetailById(rq.Id)
	if err != nil {
		return nil, err
	}

	response := &usermodel.UserUpdateResponse{UserDetail: &infoUpdated}

	if len(rq.RefCode) > 0 {
		ref, err := u.Store.GetRefUserByUserId(rq.Id)
		if err != nil {
			return nil, err
		}

		if ref.Id > 0 {
			return nil, errors.New("has ref, cant update ref user")
		}

		info, err := u.Store.GetUserByCondition(usermodel.UserFilter{Ref: rq.RefCode})
		if err != nil {
			return nil, err
		}

		if info.Id == 0 {
			return nil, errors.New("ref code invalid ,not found user from ref code")
		}

		refUser := &usermodel.RefUser{
			UserId:    rq.Id,
			RefCode:   rq.RefCode,
			Profit:    common.DEFAULT_PROFIT,
			UserRefId: info.Id,
		}

		err = u.Store.CreateRefUser(refUser)
		if err != nil {
			return nil, err
		}

		response.RefWalletAddress = info.WalletAddress
	}

	return response, nil

}

func (u UserController) TopSeller(paging *common.PagingWithTime) ([]*usermodel.TopSellerResponse, error) {
	userIdsTopSeller, err := u.MarketPlaceStorage.ListUserTopSeller(paging)
	if err != nil {
		return nil, err
	}

	logs, err := u.MarketPlaceStorage.TopSeller(userIdsTopSeller, paging)
	if err != nil {
		return nil, err
	}

	mod := big.NewFloat(1000000000000000000)
	mapTotalAmount := make(map[int64]float64, len(logs))
	mapLastTotalAmount := make(map[int64]float64, len(logs))
	mapCount := make(map[int64]int64, len(logs))
	mapLastCount := make(map[int64]int64, len(logs))
	mapVolumePerDay := make(map[int64]map[string]float64, len(logs))
	for _, v := range logs {
		distance, _ := new(big.Float).SetString(v.Price)
		div := new(big.Float).Quo(distance, mod)
		price, _ := div.Float64()

		if utils.InTimeSpan(paging.CustomStart, paging.CustomEnd, v.CreatedAt) {
			mapLastCount[v.SellerId] += 1
			mapLastTotalAmount[v.SellerId] += price
		}

		if utils.InTimeSpan(paging.Start, paging.End, v.CreatedAt) {
			mapCount[v.SellerId] += 1
			mapTotalAmount[v.SellerId] += price
			if mapVolumePerDay[v.SellerId] == nil {
				mapVolumePerDay[v.SellerId] = make(map[string]float64)
			}
			mapVolumePerDay[v.SellerId][v.CreatedAt.Format(common.ISO8601)] += price
		}
	}

	result, err := u.Store.GetListUserById(userIdsTopSeller)
	if err != nil {
		return nil, err
	}

	mapUser := make(map[int64]usermodel.UserDetail, len(result))
	for _, v := range result {
		v.Avatar = utils.RebuildUrlPath(v.Avatar)
		v.Cover = utils.RebuildUrlPath(v.Cover)
		mapUser[v.Id] = v
	}

	response := make([]*usermodel.TopSellerResponse, 0, len(logs))
	for _, v := range result {
		seller := &usermodel.TopSellerResponse{
			UserDetail:          mapUser[v.Id],
			TotalAmount:         cast.ToString(mapTotalAmount[v.Id]),
			TotalNFTSold:        mapCount[v.Id],
			PercentVolume:       utils.PercentChangeFloat64(mapLastTotalAmount[v.Id], mapTotalAmount[v.Id]),
			PercentTotalNftSold: utils.PercentChangeInt64(mapLastCount[v.Id], mapCount[v.Id]),
			USDTVolume:          utils.GetConvertedVolume(mapVolumePerDay[v.Id]),
		}
		response = append(response, seller)
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].USDTVolume > response[j].USDTVolume
	})

	return response, nil

}

func (u UserController) TopCollection(paging *common.PagingWithTime) ([]*marketplacemodel.TopCollectionResponse, error) {
	collectionIds, err := u.MarketPlaceStorage.TopVolumeCollectionId(paging)
	if err != nil {
		return nil, err
	}

	res, err := u.MarketPlaceStorage.TopVolume(collectionIds, paging)
	if err != nil {
		return nil, err
	}

	mod := float64(1000000000000000000)
	mapTotalPrice := make(map[int64]float64)
	mapLastTotalPrice := make(map[int64]float64)
	mapFloorPrice := make(map[int64]float64)
	mapVolumePerDay := make(map[int64]map[string]float64, len(res))
	for _, v := range res {
		price := v.SoldPrice / mod
		if utils.InTimeSpan(paging.CustomStart, paging.CustomEnd, v.TransCreated) {
			mapLastTotalPrice[v.Id] += price
		}

		if utils.InTimeSpan(paging.Start, paging.End, v.TransCreated) {
			mapTotalPrice[v.Id] += price
			if mapVolumePerDay[v.Id] == nil {
				mapVolumePerDay[v.Id] = make(map[string]float64)
			}
			mapVolumePerDay[v.Id][v.TransCreated.Format(common.ISO8601)] += price
		}

	}

	floorPrice, err := u.MarketPlaceStorage.MapFloorPrice(collectionIds)
	if err != nil {
		return nil, err
	}

	for _, v := range floorPrice {
		mapFloorPrice[v.Id] = v.FloorPrice / mod
	}

	mapExist := make(map[int64]bool)

	result := make([]*marketplacemodel.TopCollectionResponse, 0, len(res))
	for _, v := range res {
		if mapExist[v.Id] {
			continue
		}

		v.Collection.Logo = utils.RebuildUrlPath(v.Collection.Logo)
		v.Collection.FeaturedImage = utils.RebuildUrlPath(v.Collection.FeaturedImage)
		v.Collection.Banner = utils.RebuildUrlPath(v.Collection.Banner)

		result = append(result, &marketplacemodel.TopCollectionResponse{
			Collection:     v.Collection,
			FloorPrice:     mapFloorPrice[v.Id],
			Volume:         mapTotalPrice[v.Id],
			USDTFloorPrice: utils.GetConvertedFloorPrice(mapFloorPrice[v.Id]),
			USDTVolume:     utils.GetConvertedVolume(mapVolumePerDay[v.Id]),
			PercentVolume:  utils.PercentChangeFloat64(mapLastTotalPrice[v.Id], mapTotalPrice[v.Id]),
		})
		mapExist[v.Id] = true
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].USDTVolume > result[j].USDTVolume
	})

	return result, nil
}

func (u UserController) TopCreator(paging *common.PagingWithTime) ([]*usermodel.TopCreatorResponse, error) {
	userIds, err := u.MarketPlaceStorage.TopCreatorUserId(paging)
	if err != nil {
		return nil, err
	}

	res, err := u.MarketPlaceStorage.TopCreator(userIds, paging)
	if err != nil {
		return nil, err
	}

	mapTotalCreated := make(map[int64]int64, len(res))
	mapLastTotalCreated := make(map[int64]int64, len(res))
	for _, v := range res {
		if utils.InTimeSpan(paging.CustomStart, paging.CustomEnd, v.CreatedAt) {
			mapLastTotalCreated[v.UserId] += 1
		}

		if utils.InTimeSpan(paging.Start, paging.End, v.CreatedAt) {
			mapTotalCreated[v.UserId] += 1
		}
	}

	result, err := u.Store.GetListUserById(userIds)
	if err != nil {
		return nil, err
	}

	mapUser := make(map[int64]usermodel.UserDetail, len(result))
	for _, v := range result {
		v.Avatar = utils.RebuildUrlPath(v.Avatar)
		v.Cover = utils.RebuildUrlPath(v.Cover)
		mapUser[v.Id] = v
	}

	//get collection
	collections, err := u.MarketPlaceStorage.ListCollectionsByCondition(&marketplacemodel.CollectionFilter{
		UserIds: userIds,
		Cols:    []string{"id", "user_id"},
	}, nil)
	if err != nil {
		return nil, err
	}

	mapCollection := make(map[int64]int64, len(collections))
	for _, v := range collections {
		mapCollection[v.UserId] += 1
	}

	response := make([]*usermodel.TopCreatorResponse, 0, len(res))
	for _, v := range userIds {
		seller := &usermodel.TopCreatorResponse{
			UserDetail:          mapUser[v],
			TotalCreated:        mapTotalCreated[v],
			TotalCollection:     mapCollection[v],
			PercentTotalCreated: utils.PercentChangeInt64(mapLastTotalCreated[v], mapTotalCreated[v]),
		}
		response = append(response, seller)
	}

	sort.Slice(response, func(i, j int) bool {
		return response[i].TotalCreated > response[j].TotalCreated
	})

	return response, nil

}

func (u UserController) ChangeStatus(rq *usermodel.UpdateUserStatus) (*usermodel.User, error) {
	user := &usermodel.User{
		Id:     rq.Id,
		Status: rq.Status,
	}
	if err := u.Store.UpdateUser(user); err != nil {
		return nil, err
	}
	return user, nil

}

func (u UserController) ChangeNotificationStatus(rq *usermodel.UpdateNotificationStatus) error {
	noti := &usermodel.Notification{
		Status: common.READ,
	}
	if err := u.Store.UpdateNotification(noti, &usermodel.Notification{Id: rq.Id, UserId: rq.UserId, Status: common.ACTIVE}); err != nil {
		return err
	}
	return nil
}

func (u UserController) ReadAll(rq *usermodel.UpdateNotificationStatus) error {
	noti := &usermodel.Notification{
		Status: common.READ,
	}
	if err := u.Store.UpdateNotification(noti, &usermodel.Notification{UserId: rq.UserId, Status: common.ACTIVE}); err != nil {
		return err
	}
	return nil
}

func (u UserController) UpdateNotiSetting(rq *usermodel.NotificationSetting) error {
	if err := u.Store.UpdateNotificationSetting(rq, &usermodel.NotificationSetting{UserId: rq.UserId}); err != nil {
		return err
	}
	return nil
}

func (u UserController) NotificationSetting(rq *usermodel.NotificationSetting) (*usermodel.NotificationSetting, error) {
	result, err := u.Store.FindNotificationSetting(rq)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u UserController) ListNotification(rq *usermodel.NotificationFilter, paging *common.Paging) ([]usermodel.NotificationDTO, error) {
	settings, err := u.Store.FindNotificationSetting(&usermodel.NotificationSetting{UserId: rq.UserId})
	if err != nil {
		return nil, err
	}
	settingMap := utils.ConvertStructToMap(settings)
	notiType := make([]int, 0, 20)
	for k, v := range settingMap {
		if v == 1 {
			notiType = append(notiType, common.MapSetting[k]...)
		}
	}
	rq.CustomType = notiType

	result, err := u.Store.ListNotification(rq, paging)
	if err != nil {
		return nil, err
	}

	response := make([]usermodel.NotificationDTO, 0, len(result))

	for _, v := range result {
		dto := v.ConvertNotificationDTO()
		response = append(response, dto)
	}

	return response, nil
}

func (u UserController) CMSUserLogin(rq usermodel.CMSLoginRequest) (*usermodel.Credential, error) {

	info, errInfo := u.Store.GetLoginInformation(rq.LoginName)
	if errInfo != nil {
		return nil, common.ErrLogin
	}
	if !utils.CheckPasswordHash(rq.Password, info.Password) {
		return nil, common.ErrPasswordIncorrect
	}
	res, errCreateToken := utils.CreateToken(int64(info.CMSUserId), common.CMS_USER_TYPE)
	if errCreateToken != nil {
		return nil, common.ErrLogin
	}

	return res, nil
}

func (u UserController) CMSUserRegister(rq usermodel.CMSUser) (*usermodel.CMSUser, error) {

	hasLoginName, err := u.Store.HasLoginName(rq.LoginName)
	if err != nil || hasLoginName {
		return nil, common.ErrAccountHasRegister
	}

	rq.Password, _ = utils.HashPassword(rq.Password)
	if err := u.Store.CreateUserCMS(&rq); err != nil {
		return nil, err
	}

	return &rq, nil
}

func (u UserController) GetUsersByLogTransfer(receiver common_eth.Address, sender common_eth.Address) ([]usermodel.User, error) {
	walletAddresses := []string{receiver.Hex(), sender.Hex()}
	return u.Store.GetUserWithDistinctWalletAddresses(walletAddresses)
}

func (u UserController) GetUsersByLogCreateEvent(walletAddress common_eth.Address) (*usermodel.User, error) {
	return u.Store.GetUserWithWalletAddresses(walletAddress.Hex())
}

func (u UserController) GetUsersByLogCreateNft(walletAddress common_eth.Address) (*usermodel.User, error) {
	return u.Store.GetUserWithWalletAddresses(walletAddress.Hex())
}

func (i UserController) FollowUser(rq usermodel.FollowingRequest) error {
	user, err := i.Store.GetUserDetailById(rq.FollowerId)
	if err != nil {
		return err
	}

	if user.Id == 0 {
		return errors.New("Invalid user")
	}

	followRq := &usermodel.UserFollow{
		FollowerId:      rq.FollowerId,
		UserId:          rq.UserId,
		IsDeleted:       common.ACTIVE,
		TimeStartFollow: time.Now().Unix(),
	}
	err = i.Store.FollowUser(followRq)
	if err != nil {
		return err
	}

	notiObj, _ := json.Marshal(map[string]string{
		"user_name":      user.DisplayName,
		"user_avatar":    user.Avatar,
		"wallet_address": user.WalletAddress,
	})

	exist, err := i.Store.IsExistNotification(&usermodel.Notification{
		UserId:             rq.FollowerId,
		NotificationType:   common.FOLLOWED,
		NotificationObject: base58.Encode(notiObj),
	})

	if !exist && rq.UserId != user.Id {
		notification := &usermodel.Notification{
			UserId:             rq.FollowerId,
			LastTimeRead:       0,
			NotificationType:   common.FOLLOWED,
			Status:             common.UNREAD,
			NotificationObject: base58.Encode(notiObj),
		}
		err = i.Store.InsertNotification(notification)
		if err != nil {
			return err
		}
	}

	err = i.Store.InsertActivities(&usermodel.Activities{
		UserId:       rq.UserId,
		FollowId:     rq.FollowerId,
		MetaData:     "",
		ActivityType: common.FOLLOW_USER,
	})
	if err != nil {
		return err
	}

	return nil
}

func (i UserController) UnfollowUser(rq usermodel.FollowingRequest) error {
	user, err := i.Store.GetUserDetailById(rq.FollowerId)
	if err != nil {
		return err
	}

	if user.Id == 0 {
		return errors.New("Invalid user")
	}

	followRq := &usermodel.UserFollow{
		FollowerId: rq.FollowerId,
		UserId:     rq.UserId,
		IsDeleted:  common.ACTIVE,
	}
	err = i.Store.UnFollowUser(followRq)
	if err != nil {
		return err
	}

	err = i.Store.InsertActivities(&usermodel.Activities{
		UserId:       rq.UserId,
		FollowId:     rq.FollowerId,
		MetaData:     "",
		ActivityType: common.UNFOLLOW_USER,
	})
	if err != nil {
		return err
	}

	return nil
}

func (u UserController) ProfileStatistic(userId int64) (map[string]interface{}, error) {
	user, err := u.Store.GetUserDetailById(userId)
	if err != nil {
		return nil, err
	}

	onSale, err := u.MarketPlaceStorage.CountItemOnSaleByUserid(user)
	if err != nil {
		return nil, err
	}

	collection, err := u.MarketPlaceStorage.CountCollectionByUserid(user)
	if err != nil {
		return nil, err
	}

	created, err := u.MarketPlaceStorage.CountNFTCreatedByUserid(user)
	if err != nil {
		return nil, err
	}

	nftOwned, err := u.MarketPlaceStorage.CountOwnerNFTByUserId(user, []int{
		common.ITEM_NEW,
		common.ITEM_READY_ON_SELL,
		common.ITEM_ON_SALE,
	})
	if err != nil {
		return nil, err
	}

	liked, err := u.MarketPlaceStorage.CountLikedNFTByUserid(user)
	if err != nil {
		return nil, err
	}

	sold, err := u.MarketPlaceStorage.CountNFTSoldByUserid(user)
	if err != nil {
		return nil, err
	}

	event, err := u.MarketPlaceStorage.CountNFTEventByUserid(user)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"on_sale":    onSale,
		"collection": collection,
		"created":    created,
		"favorite":   liked,
		"owned":      nftOwned,
		"sold":       sold,
		"event":      event,
		"activities": 0,
	}, nil

}

func (u UserController) SearchCollection(rq *usermodel.SearchRequest) (map[string]interface{}, error) {
	searchUser, err := u.Store.SearchUser(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchUser {
		searchUser[index].Avatar = utils.RebuildUrlPath(searchUser[index].Avatar)
	}

	searchUserWd, err := u.Store.SearchUserWd(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchUserWd {
		searchUserWd[index].Avatar = utils.RebuildUrlPath(searchUserWd[index].Avatar)
	}

	searchUser = append(searchUser, searchUserWd...)

	searchCollection, err := u.Store.SearchCollection(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchCollection {
		searchCollection[index].FeaturedImage = utils.RebuildUrlPath(searchCollection[index].FeaturedImage)
		searchCollection[index].Logo = utils.RebuildUrlPath(searchCollection[index].Logo)
	}

	res := map[string]interface{}{
		"users":       searchUser,
		"collections": searchCollection,
	}
	return res, nil
}

func (u UserController) Search(rq *usermodel.SearchRequest, paging *common.Paging) (map[string]interface{}, error) {
	market, err := u.Store.SearchMarketPlaceNft([]int64{}, rq, nil)
	if err != nil {
		return nil, err
	}

	var listMarketItemIds []int64
	var listMarketItemMainIds = make(map[int64][]int64)
	mapItemMarket := make(map[int64]usermodel.NFTSearchResponse)
	for _, v := range market {
		mapItemMarket[v.Id] = *v
		if v.PackId > 0 {
			listMarketItemMainIds[v.PackId] = append(listMarketItemMainIds[v.PackId], v.Id)
		}
	}
	for _, itemIds := range listMarketItemMainIds {
		for i := 0; i < len(itemIds); i++ {
			if i == 0 {
				continue
			}
			listMarketItemIds = append(listMarketItemIds, itemIds[i])
		}
	}
	rq.OrNotListItemIds = listMarketItemIds
	searchNFT, err := u.Store.SearchNFT(rq, paging)
	if err != nil {
		return nil, err
	}

	listItemId := make([]int64, 0, len(searchNFT))
	userIds := make([]int64, 0, len(searchNFT))
	itemIds := make([]int64, 0, len(searchNFT))
	for _, v := range searchNFT {
		listItemId = append(listItemId, v.Id)
		userIds = append(userIds, v.UserId)
		itemIds = append(itemIds, v.Id)
	}

	for _, v := range searchNFT {
		if v.NumOfCopies > 1 {
			var countItemCopyFilter marketplacemodel.CountItemFilter
			countItemCopyFilter.Status = common.ITEM_NEW
			if v.MainItemId > 0 {
				countItemCopyFilter.MainItemId = v.MainItemId
			} else {
				countItemCopyFilter.MainItemId = v.Id
			}
			v.AvailableCopy, _ = u.Store.CountAvailableItemCopied(countItemCopyFilter)
			if v.MainItemId == 0 {
				v.AvailableCopy++
			}
		}
	}

	users, err := u.Store.GetListUserById(userIds)
	if err != nil {
		if err != nil {
			return nil, err
		}
	}
	mapLiked := make(map[int64]bool, 100)
	if rq.UserId > 0 {
		// list liked item
		listLiked, err := u.MarketPlaceStorage.ListLikedItem(listItemId, rq.UserId)
		if err != nil {
			return nil, err
		}
		for _, v := range listLiked {
			mapLiked[v.ItemId] = true
		}
	}
	mapUser := make(map[int64]usermodel.UserDetail)
	for _, v := range users {
		mapUser[v.Id] = v
	}
	for rIndex, valueResult := range searchNFT {
		searchNFT[rIndex].IsLiked = mapLiked[valueResult.Id]
	}
	var resultNFT, noPriceNFT []*usermodel.NFTSearchResponse
	if rq.PriceMax > 0 || rq.PriceMin > 0 || rq.SellType > 0 || rq.OrderBy == 2 || rq.OrderBy == 3 {
		for index, v := range searchNFT {
			searchNFT[index].Image = utils.RebuildUrlPath(searchNFT[index].Image)
			v.Owner = mapUser[v.UserId]
			v.Price = mapItemMarket[v.Id].Price
			v.AuctionEndTime = mapItemMarket[v.Id].AuctionEndTime
			v.AuctionStartTime = mapItemMarket[v.Id].AuctionStartTime
			if len(v.Price) == 0 {
				noPriceNFT = append(noPriceNFT, v)
				searchNFT[index] = nil
			} else {
				resultNFT = append(resultNFT, v)
			}
		}
	} else {
		for index, v := range searchNFT {
			searchNFT[index].Image = utils.RebuildUrlPath(searchNFT[index].Image)
			v.Owner = mapUser[v.UserId]
			v.Price = mapItemMarket[v.Id].Price
			v.AuctionEndTime = mapItemMarket[v.Id].AuctionEndTime
			v.AuctionStartTime = mapItemMarket[v.Id].AuctionStartTime
		}
		resultNFT = searchNFT
	}
	if rq.UserId != 0 {
		_ = u.Store.InsertSearchHistory(&usermodel.UserSearch{
			UserId: rq.UserId,
			Search: rq.Value,
		})
	}

	var res map[string]interface{}
	if rq.IsJustGetNft == 1 {
		if rq.OrderBy == 2 && rq.PriceMin == 0 && rq.SellType == 0 {
			resultNFT = append(resultNFT, noPriceNFT...)
			res = map[string]interface{}{
				"items": resultNFT,
			}
		} else if rq.OrderBy == 3 && rq.PriceMin == 0 && rq.SellType == 0 {
			noPriceNFT = append(noPriceNFT, resultNFT...)
			res = map[string]interface{}{
				"items": noPriceNFT,
			}
		} else {
			res = map[string]interface{}{
				"items": resultNFT,
			}
		}
		return res, nil
	}
	//user
	searchUser, err := u.Store.SearchUser(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchUser {
		searchUser[index].Avatar = utils.RebuildUrlPath(searchUser[index].Avatar)
	}

	searchUserWd, err := u.Store.SearchUserWd(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchUserWd {
		searchUserWd[index].Avatar = utils.RebuildUrlPath(searchUserWd[index].Avatar)
	}

	searchUser = append(searchUser, searchUserWd...)

	//collection
	searchCollection, err := u.Store.SearchCollection(rq)
	if err != nil {
		return nil, err
	}
	for index, _ := range searchCollection {
		searchCollection[index].FeaturedImage = utils.RebuildUrlPath(searchCollection[index].FeaturedImage)
		searchCollection[index].Logo = utils.RebuildUrlPath(searchCollection[index].Logo)
	}

	res = map[string]interface{}{
		"users":       searchUser,
		"collections": searchCollection,
		"items":       resultNFT,
	}
	return res, nil
}

func (u UserController) SearchHistory(userId int64) ([]usermodel.SearchResponse, error) {
	search, err := u.Store.GetListSearchHistory(userId)
	if err != nil {
		return nil, err
	}

	return usermodel.SearchResponseDto(search), nil
}

func (u UserController) WatchList(userId int64, paging *common.PagingWithTime) ([]marketplacemodel.CollectionWatchList, error) {
	response := make([]marketplacemodel.CollectionWatchList, 0, 100)
	collectionLikes, err := u.MarketPlaceStorage.ListCollectionLike(marketplacemodel.CollectionLikeFilter{
		UserId:    userId,
		IsDeleted: common.ACTIVE,
		Cols:      []string{"id", "collection_id"},
	})
	if err != nil {
		return nil, err
	}

	colIds := make([]int64, 0, len(collectionLikes))
	for _, v := range collectionLikes {
		if v.CollectionId != 0 {
			colIds = append(colIds, v.CollectionId)
		}
	}

	if len(colIds) == 0 {
		return response, nil
	}

	collections, err := u.MarketPlaceStorage.ListCollectionByIds(colIds)
	if err != nil {
		return nil, err
	}
	mapCol := make(map[int64]marketplacemodel.Collection, len(collections))
	for _, v := range collections {
		mapCol[v.Id] = *v
	}

	items, err := u.MarketPlaceStorage.ListItems(
		marketplacemodel.ItemFilter{
			CollectionIds: colIds,
			IsApproval:    1,
			IsTrending:    -1,
			Stts:          []int64{common.ITEM_ON_SALE, common.ITEM_NEW, common.ITEM_READY_ON_SELL},
		})
	if err != nil {
		return nil, err
	}
	mapItem := make(map[int64]int64, len(colIds))
	mapItemCol := make(map[int64]int64, len(colIds))
	itemIds := make([]int64, 0, len(items))
	mapOwner := make(map[int64]int64, len(colIds))
	mapCheckExist := make(map[string]bool, len(colIds))
	for _, v := range items {
		mapItem[v.CollectionId] += 1
		itemIds = append(itemIds, v.Id)
		keyCheck := fmt.Sprintf("%v.%v", v.CollectionId, v.UserId)
		if !mapCheckExist[keyCheck] {
			mapOwner[v.CollectionId] += 1
			mapCheckExist[keyCheck] = true
		}
		mapItemCol[v.Id] = v.CollectionId
	}

	marketPlace, err := u.MarketPlaceStorage.ListMarketPlaceWithTime(marketplacemodel.MarketPlaceFilter{
		ItemIds: itemIds,
		Status:  common.MARKET_PLACE_STATUS_SELLING,
	}, paging)
	if err != nil {
		return nil, err
	}

	logs, err := u.MarketPlaceStorage.ListTransactionLogsWithTime(marketplacemodel.TransactionLogFilter{
		ItemIds: itemIds,
	}, paging)
	if err != nil {
		return nil, err
	}

	mapMarket := make(map[int64]*marketplacemodel.MarketPlace, len(marketPlace))
	mapFloorPrice := make(map[int64]float64, len(marketPlace))
	mod := big.NewFloat(1000000000000000000)
	mIds := make([]int64, 0, len(marketPlace))

	for _, v := range marketPlace {
		mapMarket[v.ItemId] = v
		mIds = append(mIds, v.Id)

		distance, _ := new(big.Float).SetString(v.Price)
		price, _ := new(big.Float).Quo(distance, mod).Float64()
		if mapFloorPrice[mapItemCol[v.ItemId]] == 0 && v.Status == common.MARKET_PLACE_STATUS_SELLING {
			mapFloorPrice[mapItemCol[v.ItemId]] = price
		}

		if mapFloorPrice[mapItemCol[v.ItemId]] > price && v.Status == common.MARKET_PLACE_STATUS_SELLING {
			mapFloorPrice[mapItemCol[v.ItemId]] = price
		}
	}

	mapTotalPrice := make(map[int64]float64, len(logs))
	mapLastTotalPrice := make(map[int64]float64, len(logs))
	mapVolumePerDay := make(map[int64]map[string]float64, len(logs))
	for _, v := range logs {
		distance, _ := new(big.Float).SetString(v.Price)
		price, _ := new(big.Float).Quo(distance, mod).Float64()
		if utils.InTimeSpan(paging.CustomStart, paging.CustomEnd, v.CreatedAt) {
			mapLastTotalPrice[mapItemCol[v.ItemId]] += price
		}

		if utils.InTimeSpan(paging.Start, paging.End, v.CreatedAt) {
			if mapVolumePerDay[mapItemCol[v.ItemId]] == nil {
				mapVolumePerDay[mapItemCol[v.ItemId]] = make(map[string]float64)
			}
			mapVolumePerDay[mapItemCol[v.ItemId]][v.CreatedAt.Format(common.ISO8601)] += price
			mapTotalPrice[mapItemCol[v.ItemId]] += price
		}
	}

	for _, v := range collectionLikes {
		res := marketplacemodel.CollectionWatchList{
			Id:             v.CollectionId,
			Logo:           utils.RebuildUrlPath(mapCol[v.CollectionId].Logo),
			Banner:         utils.RebuildUrlPath(mapCol[v.CollectionId].Banner),
			FeaturedImage:  utils.RebuildUrlPath(mapCol[v.CollectionId].FeaturedImage),
			CollectionName: mapCol[v.CollectionId].Name,
			FloorPrice:     mapFloorPrice[v.CollectionId],
			USDTFloorPrice: utils.GetConvertedFloorPrice(mapFloorPrice[v.CollectionId]),
			TotalVolume:    mapTotalPrice[v.CollectionId],
			TotalOwner:     mapOwner[v.CollectionId],
			TotalItem:      mapItem[v.CollectionId],
			USDTVolume:     utils.GetConvertedVolume(mapVolumePerDay[v.CollectionId]),
			PercentVolume:  utils.PercentChangeFloat64(mapLastTotalPrice[v.CollectionId], mapTotalPrice[v.CollectionId]),
		}
		response = append(response, res)
	}

	return response, nil

}

func (u UserController) ListFollowing(userId int64, currentId int64) ([]*usermodel.ListFollowResponse, error) {
	listUserFollow, err := u.Store.ListFollowingUser(&usermodel.UserFollow{UserId: userId})
	if err != nil {
		return nil, err
	}

	followerIds := make([]int64, 0, len(listUserFollow))
	for _, v := range listUserFollow {
		followerIds = append(followerIds, v.FollowerId)
	}

	followers, err := u.Store.ListFollow(followerIds, currentId)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0, len(followers))
	mapFollow := make(map[int64]*usermodel.ListFollowQueryResponse)
	for _, v := range followers {
		userIds = append(userIds, v.FollowerId)
		mapFollow[v.FollowerId] = v
	}

	users, err := u.Store.GetListUserById(followerIds)
	if err != nil {
		return nil, err
	}

	result := make([]*usermodel.ListFollowResponse, 0, len(userIds))
	for _, v := range users {
		if mapFollow[v.Id] == nil {
			mapFollow[v.Id] = &usermodel.ListFollowQueryResponse{}
		}
		e := &usermodel.ListFollowResponse{
			Id:             v.Id,
			DisplayName:    v.DisplayName,
			WalletAddress:  v.WalletAddress,
			Avatar:         utils.RebuildUrlPath(v.Avatar),
			IsFollowing:    mapFollow[v.Id].IsFollowing,
			TotalFollowers: mapFollow[v.Id].TotalFollowers,
		}
		result = append(result, e)
	}

	return result, nil
}

func (u UserController) ListFollowers(userId int64, currentId int64) ([]*usermodel.ListFollowResponse, error) {
	listUserFollowing, err := u.Store.ListFollowingUser(&usermodel.UserFollow{FollowerId: userId})
	if err != nil {
		return nil, err
	}

	followingIds := make([]int64, 0, len(listUserFollowing))
	for _, v := range listUserFollowing {
		followingIds = append(followingIds, v.UserId)
	}

	followers, err := u.Store.ListFollow(followingIds, currentId)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0, len(followers))
	mapFollow := make(map[int64]*usermodel.ListFollowQueryResponse)
	for _, v := range followers {
		userIds = append(userIds, v.FollowerId)
		mapFollow[v.FollowerId] = v
	}

	users, err := u.Store.GetListUserById(followingIds)
	if err != nil {
		return nil, err
	}

	result := make([]*usermodel.ListFollowResponse, 0, len(userIds))
	for _, v := range users {
		if mapFollow[v.Id] == nil {
			mapFollow[v.Id] = &usermodel.ListFollowQueryResponse{}
		}
		e := &usermodel.ListFollowResponse{
			Id:             v.Id,
			DisplayName:    v.DisplayName,
			WalletAddress:  v.WalletAddress,
			Avatar:         utils.RebuildUrlPath(v.Avatar),
			IsFollowing:    mapFollow[v.Id].IsFollowing,
			TotalFollowers: mapFollow[v.Id].TotalFollowers,
		}
		result = append(result, e)
	}

	return result, nil
}

func (u UserController) Activities(filter usermodel.ActivitiesFilter) ([]marketplacemodel.ActivitiesUserResponse, error) {
	activities, err := u.Store.ListActivities(filter)
	if err != nil {
		return nil, err
	}

	itemIds := make([]int64, 0, len(activities))
	userIds := make([]int64, 0, len(activities))
	collectionIds := make([]int64, 0, len(activities))
	for _, v := range activities {
		if v.ItemId > 0 {
			itemIds = append(itemIds, v.ItemId)
		}
		if v.FollowId > 0 {
			userIds = append(userIds, v.FollowId)
		}
		if v.ActivityUserId > 0 {
			userIds = append(userIds, v.ActivityUserId)
		}
		if v.CollectionId > 0 {
			collectionIds = append(collectionIds, v.CollectionId)
		}
		if v.UserId > 0 {
			userIds = append(userIds, v.UserId)
		}
	}

	mapItem := make(map[int64]marketplacemodel.Item)
	if len(itemIds) > 0 {
		items, err := u.MarketPlaceStorage.ListItems(marketplacemodel.ItemFilter{Ids: itemIds})
		if err != nil {
			return nil, err
		}
		for _, v := range items {
			mapItem[v.Id] = v
		}
	}

	mapUser := make(map[int64]usermodel.UserDetail)
	if len(userIds) > 0 {
		users, err := u.Store.GetListUserById(userIds)
		if err != nil {
			return nil, err
		}
		for _, v := range users {
			mapUser[v.Id] = v
		}
	}

	mapCollection := make(map[int64]marketplacemodel.Collection)
	if len(collectionIds) > 0 {
		collections, err := u.MarketPlaceStorage.ListCollectionByIds(collectionIds)
		if err != nil {
			return nil, err
		}
		for _, v := range collections {
			mapCollection[v.Id] = *v
		}
	}

	result := make([]marketplacemodel.ActivitiesUserResponse, 0, len(activities))
	for _, v := range activities {
		data := make(map[string]string)
		_ = json.Unmarshal([]byte(v.MetaData), &data)
		result = append(result, marketplacemodel.ActivitiesUserResponse{
			Id:           v.Id,
			UserId:       v.UserId,
			ItemId:       v.ItemId,
			FollowId:     v.FollowId,
			MetaData:     data,
			CollectionId: v.CollectionId,
			ActivityType: v.ActivityType,
			User:         mapUser[v.UserId],
			Follow:       mapUser[v.FollowId],
			ActivityUser: mapUser[v.ActivityUserId],
			Item:         mapItem[v.ItemId],
			Collection:   mapCollection[v.CollectionId],
			CreatedAt:    v.CreatedAt,
		})
	}

	return result, nil

}

func (u UserController) RefNetworkProcess(filter usermodel.RefNetworkFilter) (usermodel.RefNetworkResult, error) {
	var result usermodel.RefNetworkResult
	//get list user
	listUsers, err := u.Store.GetListUserRef(filter)
	if err != nil {
		return result, err
	}
	userF0, errF0 := u.Store.GetUserF0Ref(filter.UserId)
	if errF0 != nil {
		return result, errF0
	}
	result.InviterId = userF0.WalletAddress
	var listUserId []int64
	for _, user := range listUsers {
		if user.Id == filter.UserId {
			result.UserDetailAddTimeJoinRef = user
			result.RefLink = utils.BuildRefLink(user.Ref)
			continue
		}
		result.TotalMember++
		listUserId = append(listUserId, user.Id)
	}
	if len(listUserId) < 2 {
		return result, nil
	}
	filter.UserIds = listUserId
	//get list info
	if filter.TimeFilter > 0 {
		filter.StartTimeStamp = utils.BuildDateFilterByDay(filter.TimeFilter)
	}
	listItemJoinEvent, errJoinEvent := u.MarketPlaceStorage.GetItemJoinEventHistory(filter)
	if errJoinEvent != nil {
		return result, errJoinEvent
	}

	listMarket, errMarket := u.MarketPlaceStorage.GetRefMarketPlace(filter)
	if errMarket != nil {
		return result, errMarket
	}

	mod := big.NewInt(1000000000000000000)
	modFloat := big.NewFloat(1000000000000000000)
	var tmpListUserRef []usermodel.ListRefUserResult
	var tmpTotalEarning float64 = 0
	for _, user := range listUsers {
		if user.Id == filter.UserId {
			result.UserDetailAddTimeJoinRef = user
			result.RefLink = utils.BuildRefLink(user.Ref)
			continue
		}
		var tmpUserRef usermodel.ListRefUserResult
		tmpUserRef.Id = user.Id
		tmpUserRef.Avatar = utils.RebuildUrlPath(user.Avatar)
		tmpUserRef.DisplayName = user.DisplayName
		tmpUserRef.WalletAddress = user.WalletAddress
		tmpEventVolume := big.NewFloat(0)
		for _, item := range listItemJoinEvent {
			if user.Id == item.UserId && item.CreatedAt.Sub(user.TimeJoinRef) >= 0 {
				tmpDistance, _ := strconv.ParseFloat(item.Price, 64)
				distance := big.NewFloat(tmpDistance)
				tmpEventVolume = tmpEventVolume.Add(tmpEventVolume, distance)
				//price := new(big.Int).Div(distance, mod).Int64()
				//eventVolume += price
				tmpUserRef.EventTotalNFT++
			}
		}

		var marketTotalVolume int64 = 0
		for _, market := range listMarket {
			if user.Id == market.UserId && market.CreatedAt.Sub(user.TimeJoinRef) >= 0 {
				distance, _ := new(big.Int).SetString(market.Price, 10)
				price := new(big.Int).Div(distance, mod).Int64()
				marketTotalVolume += price
			}
		}

		eventVolume := new(big.Float).Quo(tmpEventVolume, modFloat)
		tmpEventVolumeString := fmt.Sprint(eventVolume)
		tmpEventVolumeIndex := strings.Index(tmpEventVolumeString, ".")
		tmpEventVolumeString = tmpEventVolumeString[:tmpEventVolumeIndex+2]
		tmpEventVolumeFloat, _ := strconv.ParseFloat(tmpEventVolumeString, 64)
		tmpUserRef.EventVolume = tmpEventVolumeString
		tmpTotalEarning += tmpEventVolumeFloat
		tmpMarketTotalVolume := float64(marketTotalVolume) * 0.001
		tmpTotalEarning += tmpMarketTotalVolume
		tmpUserRef.MarketTotalVolume = fmt.Sprint(marketTotalVolume)
		tmpUserRef.MarketVolume = fmt.Sprint(tmpMarketTotalVolume)
		tmpListUserRef = append(tmpListUserRef, tmpUserRef)
	}
	result.TotalEarning = fmt.Sprint(tmpTotalEarning)
	//default
	sort.Slice(tmpListUserRef, func(i, j int) bool {
		return tmpListUserRef[i].EventVolume < tmpListUserRef[j].EventVolume
	})

	if filter.SortUserName == 1 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].DisplayName < tmpListUserRef[j].DisplayName
		})
	} else if filter.SortUserName == 2 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].DisplayName > tmpListUserRef[j].DisplayName
		})
	}

	if filter.SortEventVolume == 1 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].EventVolume < tmpListUserRef[j].EventVolume
		})
	} else if filter.SortEventVolume == 2 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].EventVolume > tmpListUserRef[j].EventVolume
		})
	}

	if filter.SortMarketVolume == 1 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].MarketVolume < tmpListUserRef[j].MarketVolume
		})
	} else if filter.SortMarketVolume == 2 {
		sort.Slice(tmpListUserRef, func(i, j int) bool {
			return tmpListUserRef[i].MarketVolume > tmpListUserRef[j].MarketVolume
		})
	}

	result.ListUserRefs = tmpListUserRef
	return result, nil
}

func (u UserController) SaveSettingProcess(rq []usermodel.SystemSetting, options map[string]interface{}) ([]usermodel.SystemSetting, error) {
	for index, systemSetting := range rq {
		image := utils.DESTINATION + utils.GetFileName(systemSetting.Image)

		if len(systemSetting.Image) > 0 {
			if err := utils.MoveFile(strings.ReplaceAll(systemSetting.Image, cast.ToString(options["host"]), ""), image); err != nil {
				return nil, err
			}
		}
		if image != utils.DESTINATION {
			image = cast.ToString(options["host"]) + image
		}
		rq[index].Image = image
	}
	//save
	if err := u.Store.SystemSettingBatchSave(rq); err != nil {
		return nil, err
	}
	return rq, nil
}

func (u UserController) DeleteSettingProcess(id int64) error {
	if err := u.Store.DeleteSystemSetting(id); err != nil {
		return err
	}
	return nil
}

func (u UserController) DefaultSystemSettingProcess() (usermodel.DefaultSetting, error) {
	var response usermodel.DefaultSetting
	listSystemSettings, err := u.Store.GetAllSystemSetting()
	if err != nil {
		return response, err
	}
	if len(listSystemSettings) < 1 {
		return response, nil
	}

	for _, setting := range listSystemSettings {
		if setting.Type == common.SYSTEM_SETTING_TYPE_TOP_BANNER {
			response.TopBanner = append(response.TopBanner, setting)
		} else if setting.Type == common.SYSTEM_SETTING_TYPE_LEFT_BANNER {
			response.LeftBanner = append(response.LeftBanner, setting)
		} else if setting.Type == common.SYSTEM_SETTING_TYPE_RIGHT_BANNER {
			response.RightBanner = append(response.RightBanner, setting)
		} else if setting.Type == common.SYSTEM_SETTING_TYPE_TEXT_BANNER {
			response.TextBanner = append(response.TextBanner, setting)
		}
	}
	return response, nil
}

func (u UserController) GetSystemSettingProcess() ([]usermodel.SystemSetting, error) {
	var response []usermodel.SystemSetting
	listSystemSettings, err := u.Store.GetAllSystemSetting()
	if err != nil {
		return response, err
	}
	return listSystemSettings, nil
}

func (u UserController) ExploreUser(in usermodel.ExploreUserFilter) ([]usermodel.ExploreUserResponse, error) {
	result, err := u.Store.ExploreUser(in)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u UserController) PaymentNotificationProcess(rq usermodel.PaymentNotificationRequest) error {
	var filter usermodel.BankIncomeFilter
	filter.TransactionId = rq.OrderId
	filter.Status = common.BANK_INCOME_STATUS_WAITNG
	bankIncome, err := u.Store.GetBankIncome(filter)
	if err != nil {
		return err
	}
	if bankIncome.Id < 1 {
		return errors.New("record not found")
	}

	bankIncome.Status = common.BANK_INCOME_STATUS_SUCCESS
	bankIncome.Sign = rq.Sign
	bankIncome.Amount = rq.Amount
	if updateErr := u.Store.UpdateBankIncome(bankIncome.Id, bankIncome); updateErr != nil {
		return updateErr
	}
	return nil
}

func (u UserController) SaveDailyEventProcess(rq usermodel.DailyEvent, options map[string]interface{}) (usermodel.DailyEvent, error) {

	image := utils.DESTINATION + utils.GetFileName(rq.Thumbnail)
	if len(rq.Thumbnail) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Thumbnail, cast.ToString(options["host"]), ""), image); err != nil {
			return rq, err
		}
	}
	if image != utils.DESTINATION {
		image = cast.ToString(options["host"]) + image
	}
	rq.Thumbnail = image
	if rq.Id == 0 {
		if err := u.Store.SystemDailyEventCreate(rq); err != nil {
			return rq, err
		}
	} else {
		if err := u.Store.SystemDailyEventUpdate(rq.Id, rq); err != nil {
			return rq, err
		}
	}
	return rq, nil
}

func (u UserController) DeleteDailyEventProcess(id int64) error {
	return u.Store.SystemDailyEventDelete(id)
}

func (u UserController) DetailDailyEventProcess(req usermodel.DailyEventFilter) (usermodel.DailyEvent, error) {
	dailyEvent, err := u.Store.SystemDailyEventDetail(req)
	if err != nil {
		return dailyEvent, err
	}
	return dailyEvent, nil
}

func (u UserController) ListDailyEventProcess(req usermodel.DailyEventFilter, page *common.Paging) ([]usermodel.DailyEvent, error) {
	dailyEvent, err := u.Store.SystemDailyEventList(req, page)
	if err != nil {
		return dailyEvent, err
	}
	return dailyEvent, nil
}

func (u UserController) UserDailyEventProcess(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.UserDailyEventExInfo, error) {
	//intTime, _ := strconv.ParseInt(req.CurrentTime, 10, 64)
	timeCurrent := time.Now()
	tmpTimeCurrent := fmt.Sprint(timeCurrent.Year()) + "-" + fmt.Sprint(int(timeCurrent.Month())) + "-" + fmt.Sprint(timeCurrent.Day())
	timeCurrentStart, _ := time.Parse(common.ISO8601TIME, tmpTimeCurrent+" 00:00:00")
	timeCurrentEnd, _ := time.Parse(common.ISO8601TIME, tmpTimeCurrent+" 23:59:59")
	req.CurrentTimeStart = fmt.Sprint(timeCurrentStart.Unix())
	req.CurrentTimeEnd = fmt.Sprint(timeCurrentEnd.Unix())
	info, err := u.Store.SystemUserDailyEventList(req, nil)
	if err != nil {
		return nil, err
	}
	var result []usermodel.UserDailyEventExInfo
	var listUsed = make(map[int64]int64)
	for i := 0; i < len(info); i++ {
		var tmpResult usermodel.UserDailyEventExInfo
		if _, ok := listUsed[info[i].DailyEventId]; ok {
			continue
		}
		listUsed[info[i].DailyEventId] = info[i].DailyEventId
		tmpResult = info[i]
		tmpResult.NumberOfCurrentTask = 1
		for j := 1; j < len(info)-1; j++ {
			if info[i].DailyEventId == info[j].DailyEventId {
				tmpResult.NumberOfCurrentTask++
				tmpResult.RewardOp += info[j].RewardOp
			}
		}
		result = append(result, tmpResult)
	}
	return result, err
}

func (u UserController) UserDailyEventHistoryProcess(req usermodel.DailyEventFilter, paging *common.Paging) ([]usermodel.UserDailyEventExInfo, error) {
	//intTime, _ := strconv.ParseInt(req.CurrentTime, 10, 64)
	info, err := u.Store.SystemUserDailyEventList(req, paging)
	if err != nil {
		return nil, err
	}
	var result []usermodel.UserDailyEventExInfo
	for i := 0; i < len(info); i++ {
		var tmpResult usermodel.UserDailyEventExInfo
		tmpResult = info[i]
		tmpResult.NumberOfCurrentTask = 1
		result = append(result, tmpResult)
	}
	return result, err
}

func (u UserController) BuildUserDailyEvent(userId int64, dailyEventType int) {
	var req usermodel.DailyEventFilter
	//get user
	user, errU := u.Store.GetUserDetailById(req.UserId)
	if errU != nil {
		return
	}
	req.EventType = dailyEventType
	dailyEvent, err := u.Store.SystemDailyEventDetail(req)
	if err != nil {
		return
	}
	timeCurrent := fmt.Sprint(time.Now().Year()) + "-" + fmt.Sprint(int(time.Now().Month())) + "-" + fmt.Sprint(time.Now().Day())
	timeCurrentStart, _ := time.Parse(common.ISO8601TIME, timeCurrent+" 00:00:00")
	timeCurrentEnd, _ := time.Parse(common.ISO8601TIME, timeCurrent+" 23:59:59")
	req.UserId = userId
	req.EventType = 0
	req.DailyEventId = dailyEvent.Id
	req.CurrentTimeStart = fmt.Sprint(timeCurrentStart.Unix())
	req.CurrentTimeEnd = fmt.Sprint(timeCurrentEnd.Unix())
	total, _ := u.Store.SystemUserDailyEventCount(req)
	var userDailyEvent usermodel.UserDailyEvent
	userDailyEvent.UserId = userId
	userDailyEvent.RewardOp = dailyEvent.RewardOp
	userDailyEvent.DailyEventId = dailyEvent.Id
	userDailyEvent.ActivityDate = fmt.Sprint(time.Now().Unix())
	if int(total) >= dailyEvent.NumberOfTask {
		return
	}
	u.Store.SystemUserDailyEventCreate(userDailyEvent)
	//update user
	errSave := u.Store.UpdateUserOp(userId, user.Op+dailyEvent.RewardOp, user.OpClaimed+dailyEvent.RewardOp, "")
	if errSave != nil {
		fmt.Println(errSave.Error())
	}
}

func (u UserController) ListRewardProcess(req usermodel.RewardFilter, paging *common.Paging) ([]usermodel.RewardDepartment, error) {
	result, err := u.Store.ListRewardDepartment(req, paging)
	return result, err
}

func (u UserController) DetailRewardProcess(req usermodel.RewardFilter) (usermodel.RewardDto, error) {
	tmpResult, err := u.Store.DetailReward(req)
	result := tmpResult.ConvertToDto()
	if tmpResult.Id < 1 {
		return result, errors.New("Record not found")
	}
	result.ClaimTime = fmt.Sprint(tmpResult.CreatedAt.Unix())
	return result, err
}

func (u UserController) SaveRewardDepartmentProcess(rq usermodel.RewardDepartment, options map[string]interface{}) (usermodel.RewardDepartment, error) {
	image := utils.DESTINATION + utils.GetFileName(rq.Thumbnail)
	if len(rq.Thumbnail) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Thumbnail, cast.ToString(options["host"]), ""), image); err != nil {
			return rq, err
		}
	}
	if image != utils.DESTINATION {
		image = cast.ToString(options["host"]) + image
	}
	rq.Thumbnail = image
	if rq.Id == 0 {
		if err := u.Store.SaveRewardDepartment(&rq); err != nil {
			return rq, err
		}
	} else {
		if err := u.Store.UpdateRewardDepartment(rq.Id, rq); err != nil {
			return rq, err
		}
	}
	return rq, nil
}

func (u UserController) SaveRewardProcess(rq usermodel.Reward, options map[string]interface{}) (usermodel.Reward, error) {
	image := utils.DESTINATION + utils.GetFileName(rq.Image)
	if len(rq.Image) > 0 {
		if err := utils.MoveFile(strings.ReplaceAll(rq.Image, cast.ToString(options["host"]), ""), image); err != nil {
			return rq, err
		}
	}
	if image != utils.DESTINATION {
		image = cast.ToString(options["host"]) + image
	}
	rq.Image = image
	if rq.Id == 0 {
		if err := u.Store.SaveReward(&rq); err != nil {
			return rq, err
		}
	} else {
		if err := u.Store.UpdateReward(rq.Id, rq); err != nil {
			return rq, err
		}
	}
	return rq, nil
}

func (u UserController) ClaimRewardProcess(rq usermodel.ClaimReward) (usermodel.UserReward, error) {
	var result usermodel.UserReward
	reward, errReward := u.Store.DetailReward(usermodel.RewardFilter{RewardId: rq.RewardId})
	if errReward != nil {
		return result, errReward
	}
	if reward.TotalReward < (reward.TotalRewarded + 1) {
		return result, errors.New("Not enough reward")
	}
	var userReward = usermodel.UserReward{
		UserId:          rq.UserId,
		RewardId:        rq.RewardId,
		Price:           reward.Price,
		Status:          common.USER_REWARD_STATUS_NEW,
		ExpiredDateTime: reward.ExpiredDateTime,
	}
	//get user
	user, errU := u.Store.GetUserDetailById(rq.UserId)
	if errU != nil {
		return result, errU
	}
	if user.Op < reward.Price {
		return result, errors.New("Not enough Op point")
	}

	if err := u.Store.SaveUserReward(&userReward, reward.TotalRewarded+1, user.OpSpend+reward.Price, user.Op-reward.Price); err != nil {
		return userReward, err
	}

	return userReward, nil
}

func (u UserController) GetUserRewardProcess(rq usermodel.RewardFilter) ([]usermodel.UserReward, error) {
	result, errReward := u.Store.ListUserReward(rq)
	if errReward != nil {
		return result, errReward
	}
	return result, nil
}

func (u UserController) ListAllRewardDepartment(paging *common.Paging) (map[int64]string, error) {
	data, err := u.Store.ListRewardDepart(paging)
	result := make(map[int64]string)

	for _, detail := range data {
		result[detail.Id] = detail.RewardDepartmentName
	}
	return result, err
}

func (u UserController) ListRewardDepart(paging *common.Paging) ([]usermodel.RewardDepartment, error) {
	result, err := u.Store.ListRewardDepart(paging)
	return result, err
}

func (u UserController) ListReward(paging *common.Paging) ([]usermodel.Reward, error) {
	result, err := u.Store.ListReward(paging)
	return result, err
}

func (u UserController) SendVerifyMail(rq usermodel.SendVerifyEmailRequest) error {
	userfilter := usermodel.UserFilter{
		Email:      rq.TmpEmail,
		IsVerified: common.EMAIL_VERIFIED,
	}
	CheckUsers, err := u.Store.GetUserByCondition(userfilter)
	if err != nil {
		return err
	}
	if CheckUsers.Id > 0 {
		return errors.New("email has exist")
	}
	log.Println(rq.TmpEmail)
	hash := sha256.New()
	hash.Write([]byte(string(rq.UserId)))
	verifyToken := hash.Sum(nil)
	token := hex.EncodeToString(verifyToken)
	ExpireDate := time.Now()
	ExpireDate = ExpireDate.AddDate(0, 0, 3)
	log.Println(token, ExpireDate, rq.TmpEmail, rq.DisplayName)
	user := &usermodel.User{
		Id:          rq.UserId,
		VerifyToken: token,
		ExpireMail:  ExpireDate,
		TmpEmail:    rq.TmpEmail,
		DisplayName: rq.DisplayName,
	}
	if err = u.Store.UpdateUser(user); err != nil {
		return err
	}
	if errSendMail := utils.SendVerifyMail(rq.DisplayName, token, rq.UserId, []string{rq.TmpEmail}); errSendMail != nil {
		return errSendMail
	}
	return nil
}

func (u UserController) VerifyMail(rq usermodel.VerifyEmailRequest) error {
	verify, errToken := u.Store.GetVerifyToken(rq.UserId)
	log.Println(verify.VerifyToken, verify.TmpEmail)
	if errToken != nil {
		return errToken
	}
	if verify.VerifyToken == rq.VerifyToken {
		user := &usermodel.User{
			Id:              rq.UserId,
			IsVerifiedEmail: common.EMAIL_VERIFIED,
			Email:           verify.TmpEmail,
		}
		if err := u.Store.UpdateUser(user); err != nil {
			return err
		}
	} else if verify.VerifyToken != rq.VerifyToken {
		return errors.New("Invalid Verify Token")
	}
	return nil
}

func (u UserController) ExchangeOPV(rq *usermodel.ExchangeOPV) (float64, error) {
	rq.OPV = utils.GetConvertedNow(rq.OPV)
	return rq.OPV, nil
}
