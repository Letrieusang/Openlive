package usermodel

type NonceRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required"`
}

type CheckExistUserRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required"`
}

type OTPRequest struct {
	Email          string `json:"email" validate:"required"`
	Nonce          string `json:"nonce" validate:"required"`
	WalletProvider string `json:"wallet_provider" validate:"required"`
}

type OTPRequestReset struct {
	WalletProvider string `json:"wallet_provider" validate:"required"`
}

type CreateUserRequest struct {
	WalletAddress string `json:"wallet_address" validate:"required"`
	RefCode       string `json:"ref_code"`
	UserId        int64  `json:"user_id"`
}

type UserLoginRequest struct {
	WalletAddress string `json:"walletAddress" validate:"required"`
	HashData      string `json:"hashData" validate:"required"`
}

type CMSLoginRequest struct {
	LoginName string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type FollowingRequest struct {
	UserId     int64 `json:"user_id" validate:"required"`
	FollowerId int64 `json:"follower_id" validate:"required"`
}
