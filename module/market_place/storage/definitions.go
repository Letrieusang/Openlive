package marketplacestorage

import (
	"api-openlive/common"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
)

type Storage interface {
	InsertMarketPlace(rq *marketplacemodel.MarketPlace) error
	InsertMarketPlaceSold(rq *marketplacemodel.MarketPlaceSold) error
	InsertTransactionLog(logs *marketplacemodel.TransactionLog) error
	GetMarketPlaceByItemId(marketPlace *marketplacemodel.MarketPlaceFilter) (*marketplacemodel.MarketPlace, error)
	GetMarketPlaceByIdOnMarket(marketPlace *marketplacemodel.MarketPlaceFilter) (*marketplacemodel.MarketPlace, error)
	ListMarketPlace(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.MarketPlace, error)
	ListMarketPlaceRebuild(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]*marketplacemodel.MarketPlaceWithCollectionName, error)
	ListMostLikeMarketPlace(limit int) ([]*marketplacemodel.MarketPlace, error)
	UpdateMarketPlaceStatus(market *marketplacemodel.MarketPlace) error
	UpdateMarketPlacePack(idOnMarket []int64, transaction string, sellType int, packId int64) error
	CountMarketPlacePack(idOnMarket []int64, transaction string, sellType int) (int64, error)
	InsertVoteItem(in *marketplacemodel.Vote) error
	InsertStaking(rq *marketplacemodel.Staking) (*marketplacemodel.Staking, error)
	FindStaking(rq *marketplacemodel.Staking) (*marketplacemodel.Staking, error)
	UpdateStaking(*marketplacemodel.Staking) error
	IsStakingByUser(rq *marketplacemodel.Staking) (bool, error)
	ListVoting(rq *marketplacemodel.VoteFilter) ([]*marketplacemodel.Vote, error)
	IsExistMarketPlace(rq *marketplacemodel.MarketPlace) (bool, error)
	DisLikeItem(in *marketplacemodel.Like) error
	InsertLikeItem(in *marketplacemodel.Like) error
	CancelSellItem(in *marketplacemodel.CancelMarketplace) error
	DisLikeCollection(in *marketplacemodel.CollectionLike) error
	InsertLikeCollection(in *marketplacemodel.CollectionLike) error
	CountCollectionLike(id []int64) ([]*marketplacemodel.CollectionLikeResponse, error)
	ListCollectionLike(filter marketplacemodel.CollectionLikeFilter) ([]*marketplacemodel.CollectionLike, error)
	UpdateViewItem(itemId int64) error
	UpdateViewCollection(id int64) error
	CountMarketPlace(rq marketplacemodel.MarketPlaceFilter) (int64, error)
	ListMarketPlaceSold(rq marketplacemodel.MarketPlaceSoldFilter) ([]*marketplacemodel.MarketPlaceSold, error)
	ListTransactionLogs(marketPlace marketplacemodel.TransactionLogFilter) ([]*marketplacemodel.TransactionLog, error)
	ListTransactionLogsWithTime(marketPlace marketplacemodel.TransactionLogFilter, paging *common.PagingWithTime) ([]*marketplacemodel.TransactionLog, error)
	ListMarketPlaceWithTime(rq marketplacemodel.MarketPlaceFilter, paging *common.PagingWithTime) ([]*marketplacemodel.MarketPlace, error)

	// items
	GetItemWithCondition(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.Item, error)
	GetItemWithConditionNew(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.ItemDetail, error)
	GetItemWithConditionNewWithMarket(rq marketplacemodel.ItemFilter, paging *common.Paging) ([]marketplacemodel.ItemDetail, error)
	GetItemByID(rq marketplacemodel.ItemFilter) (*marketplacemodel.Item, error)
	UpdateItem(updator, selector *marketplacemodel.Item) error
	UpdateItemCms(id int64, update map[string]interface{}) error
	FindItem(selector *marketplacemodel.Item) (*marketplacemodel.Item, error)
	GetListItemById(listUserId []int64) ([]marketplacemodel.Item, error)
	ListItems(rq marketplacemodel.ItemFilter) ([]marketplacemodel.Item, error)
	GetListItemByEvent(rq marketplacemodel.ItemEventFilter, paging *common.Paging) ([]*marketplacemodel.ListItemEventResult, error)
	CountListItemByEvent() (int64, error)
	UpdateItemStatus(id []int64, status int) error
	UpdateItemSync(items []marketplacemodel.Item) error
	UpdateItemUserId(id, userId, collectionId int64, status int) error
	InsertCollection(rq *marketplacemodel.Collection) error
	InsertCategory(rq *marketplacemodel.Category) error
	UpdateCategory(rq *marketplacemodel.Category) error
	DeleteCategory(rq *marketplacemodel.Category) error
	ListCategories(rq *marketplacemodel.CategoryFilter) ([]*marketplacemodel.Category, error)
	ListCollections(rq *marketplacemodel.CollectionFilter, paging *common.Paging) ([]*marketplacemodel.Collection, error)
	ListCollectionsByCondition(rq *marketplacemodel.CollectionFilter, paging *common.Paging) ([]*marketplacemodel.Collection, error)
	ItemCreators(rq []marketplacemodel.ItemCreator) error
	ItemCreator(rq *marketplacemodel.ItemCreator) error
	InsertItem(rq marketplacemodel.Item) (int64, error)
	InsertItems(item marketplacemodel.Item, rq []marketplacemodel.Item) (int64, error)
	BatchInsertItem(items []marketplacemodel.Item) (int64, error)
	IsExistCategory(rq *marketplacemodel.Category) (bool, error)
	IsExistCollection(rq *marketplacemodel.Collection) (bool, error)
	IsLikedItem(itemId int64, userId int64) (bool, error)
	GetByUserIdAndNftIds(userId int64, tokenId []string, contractAddress string) (map[string]bool, error)
	UpdateItemStatusByUserIdAndNft(status int, userId int64, nft string) error
	BuyItem(sellerId int64, ownerId int64, nft string) error
	ListLikedItem(itemId []int64, userId int64) ([]marketplacemodel.Like, error)
	ListLikes(itemId []int64) ([]marketplacemodel.Like, error)
	BatchUpdateCollectionItem(ids []int64, updater map[string]interface{}) error
	CountLikeById(itemId int64) (int64, error)
	IsLikeItem(currentUserId, itemId int64) (int64, error)
	UpdateOwnerOfItem(request marketplacemodel.UpdateItemOwnerRequest) error
	UpdateItemJoinEvent(request marketplacemodel.UpdateItemEvent) error
	UpdateItemCancelEvent(request marketplacemodel.UpdateItemEvent) error
	CountItemEventByCate(cateId int64) (int64, error)
	CheckItemCanJoinEvent(walletAddress string, nft string, userId int) (bool, error)
	CheckItemJoinEvent(nft string) (bool, error)
	ListCollectionByIds(ids []int64) ([]*marketplacemodel.Collection, error)
	CheckNameCollection(rq *marketplacemodel.Collection) (bool, error)
	CountAvailableItemCopied(filter marketplacemodel.CountItemFilter) (int64, error)
	ExploreNft(filter marketplacemodel.ExploreNftFilter, paging *common.Paging) ([]*marketplacemodel.ExploreNftQueryResponse, error)

	//colection
	DetailCollection(collectionId int64) (*marketplacemodel.Collection, error)
	GetCollectionByCondition(rq *marketplacemodel.CanCreateCollectionRequest) (marketplacemodel.Collection, error)
	UpdateCollection(updator, selector *marketplacemodel.Collection) error
	UpdateCollectionCms(id int64, update map[string]interface{}) error
	FindCollection(selector *marketplacemodel.Collection) (*marketplacemodel.Collection, error)
	GetCollectionItems(listCollectionId []int64) ([]marketplacemodel.Item, error)
	TopSeller(in []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TransactionLogTopSeller, error)
	TopCreator(ids []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TopCreatorResponse, error)
	TopVolume(ids []int64, paging *common.PagingWithTime) ([]*marketplacemodel.TopVolume, error)
	MapFloorPrice(collectionIds []int64) ([]*marketplacemodel.MapFloorPrice, error)
	CountItemSell(userId []int64) ([]*marketplacemodel.TransactionLogTopSeller, error)
	CollectionItemProperties(collectionId int64) ([]*marketplacemodel.Item, error)
	ListUserTopSeller(paging *common.PagingWithTime) ([]int64, error)
	TopVolumeCollectionId(paging *common.PagingWithTime) ([]int64, error)
	TopCreatorUserId(paging *common.PagingWithTime) ([]int64, error)

	//transaction
	GetListTransaction(rq marketplacemodel.MarketPlaceFilter, paging *common.Paging) ([]marketplacemodel.ShortInfoMarketPlace, error)

	//setting
	GetSetting() (*marketplacemodel.Setting, error)

	//block
	InsertBlock(block []marketplacemodel.BlockChainBlock) error
	AddWsBlock(block marketplacemodel.BlockChainBlock) error
	UpdateWsBlockSuccess(block []marketplacemodel.BlockChainBlock, listIds []uint64) error
	GetLatestBlock() (marketplacemodel.BlockChainBlock, error)
	GetWsBlock() ([]marketplacemodel.BlockChainBlock, error)

	//statistic
	CountItemOnSaleByUserid(user usermodel.UserDetail) (int64, error)
	CountCollectionByUserid(user usermodel.UserDetail) (int64, error)
	CountNFTByUserid(userId int64, status []int) (int64, error)
	CountOwnerNFTByUserId(user usermodel.UserDetail, status []int) (int64, error)
	CountLikedNFTByUserid(user usermodel.UserDetail) (int64, error)
	CountNFTEventByUserid(user usermodel.UserDetail) (int64, error)
	CountNFTSoldByUserid(user usermodel.UserDetail) (int64, error)
	CountNFTCreatedByUserid(user usermodel.UserDetail) (int64, error)

	CollectionStats(itemIds []int64, lastTime int64) ([]marketplacemodel.MarketPlace, error)
	ExploreStatistic(total *int64) ([]*marketplacemodel.CategoryStatistic, error)

	//ref
	GetItemJoinEventHistory(filter usermodel.RefNetworkFilter) ([]marketplacemodel.ListRefInfoResult, error)
	GetRefMarketPlace(filter usermodel.RefNetworkFilter) ([]marketplacemodel.ListRefInfoResult, error)
	CheckOldItemEvent(itemId int64) (bool, error)

	//brands
	GetBrandImage(ListBrandsImage []int64) ([]marketplacemodel.BrandImage, error)
	ListBrands(rq *marketplacemodel.BrandFilter, paging *common.Paging) ([]*marketplacemodel.Brand, error)
	UpdateBrandCms(id int64, update map[string]interface{}) error
	UpdateBrandImageCms(id int64, update map[string]interface{}) error
	CreateBrandCms(brand marketplacemodel.Brand, brandImages []marketplacemodel.BrandImage) error
	CreateBrandImageCms(create *marketplacemodel.BrandImageDetail) error
	DeleteBrandImage(id int64) error
	DeleteBrand(id int64) error

	//reports
	ListReport(paging *common.Paging) ([]*marketplacemodel.ReportTableDetail, error)
	ChangeReport(rq *marketplacemodel.ChangeReportRequest) error
	RequestReport(rq *marketplacemodel.ReportTableDetail) error

	//banner
	AddBannerCms(update map[string]interface{}) error
	ListBannerCms() ([]marketplacemodel.AddBannerCmsRequest, error)

	HideNftCms(id int64) error
	UnHideNftCms(id int64) error
	NFTStatistic() (int64, error)
	CountTotalNft() (int64, error)

	CollectionExplore(filter marketplacemodel.ExploreCollectionFilter, paging *common.Paging) ([]*marketplacemodel.CollectionExploreResponse, error)

	UpdateMarketPlaceCms(id int64, update map[string]interface{}) error

	SaveAuctionBid(bid marketplacemodel.AuctionBid) error
	GetAuctionBid(filter marketplacemodel.GetAuctionFilter, paging *common.Paging) ([]marketplacemodel.AuctionBid, error)
}
