package marketplacemodel

type UpdateItemEvent struct {
	ItemId          int64
	ItemName        string
	UserAvatar      string
	WalletAddress   string
	UserDisplay     string
	UserId          int64
	Nft             string
	Revenue         string
	ContractAddress string
}
