package common

import "errors"

// marketplace error
var (
	ErrItemNotPricing        = errors.New("item not pricing")
	ErrNotFoundMarketPlaceId = errors.New("market place id not found")
	ErrUserNotStaking        = errors.New("can't vote because haven't staking")
	ErrUserLiked             = errors.New("liked")
	ErrUserDisLiked          = errors.New("disliked")
	ErrUserNotLiked          = errors.New("not like")
	ErrUpdateOwner           = errors.New("update owner of nft error")
	ErrUserFollowed          = errors.New("user followed")
	ErrUserNotFollowed       = errors.New("user not followed")
	ErrUserUnFollowed        = errors.New("user unfollowed")
	ErrNftHided              = errors.New("nft hided")
	ErrNftUnHided            = errors.New("nft unhided")
	ErrNumberOfItem          = errors.New("Invalid number of item")
)

// items error
var (
	ErrNotFoundItem = errors.New("not found item")
)

// user error
var (
	ErrInUsed                  = errors.New("login name or wallet or email has been used")
	ErrNonceInvalid            = errors.New("nonce invalid")
	ErrOTPInvalid              = errors.New("OTP code wrong or invalid")
	ErrVerify                  = errors.New("verify sign error")
	ErrPasswordIncorrect       = errors.New("password is incorrect")
	ErrNotFoundUser            = errors.New("not found user")
	ErrHashPassword            = errors.New("error hash password")
	ErrStatusNotMatch          = errors.New("user status invalid")
	ErrNotFoundCate            = errors.New("not found category")
	ErrNotFoundCol             = errors.New("not found collection")
	ErrNotPermissionUpdate     = errors.New("not permission update collection")
	ErrNotPermissionUpdateItem = errors.New("not permission update item")
	ErrFindWalletAddress       = errors.New("find wallet address error")
	ErrLogin                   = errors.New("error login fail")
	ErrAccountHasRegister      = errors.New("login name has been created")
	ErrCantNotUpdateNft        = errors.New("item on market can not update collection")
	ErrCantUpdateItemStatus    = errors.New("item is approval")
	ErrNFTCodeRequired         = errors.New("code is required to lock nft")
	ErrNFTPropertiesRequired   = errors.New("properties is required")
	ErrNotPermissionHideItem   = errors.New("not permission hide/unhide item")
)
