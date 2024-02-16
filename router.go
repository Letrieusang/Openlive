package main

import (
	"api-openlive/common"
	"api-openlive/docs"
	"api-openlive/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func (s *server) NewRouter() *gin.Engine {
	router := gin.New()
	if *Environment != "prod" {
		router.Use(gin.Logger())
	}
	//router.Use(gin.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a OPVLive server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "apipi.openlivenft.io"
	docs.SwaggerInfo.Schemes = []string{"https"}
	if *Environment == "local" {
		// development
		docs.SwaggerInfo.Host = "localhost:8081"
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	cms := router.Group("/cms")
	{
		cms.GET("/all/list_reward_department", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.GetAllListReWardDepartmentCms())
		cms.GET("/list_reward_department", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.GetListReWardDepartmentCms())
		cms.GET("/list_reward", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.GetListReWardCms())
		cms.POST("/hide", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.HideNftCMS())
		cms.POST("/unhide", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.UnHideNftCMS())
		cms.GET("/list/report", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ListReportCms())
		cms.PUT("update/report", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ChangeReportCms())
		cms.PATCH("/market-place/update/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.UpdateMarketPlaceCMS())
		cms.GET("/market-place/list-item", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ListItemCMS())
		cms.DELETE("/delete/brandimage/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.DeleteBrandImageCms())
		cms.GET("getbanner", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.GetBannerCms())
		cms.DELETE("/delete/brand/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.DeleteBrandCms())
		cms.POST("/updatebanner", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.AddBannerCms())
		cms.GET("/listbrands", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ListBrandCms())
		cms.PATCH("/change/brands/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ChangeBrandCms())
		cms.PATCH("change/brandsimage/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ChangeBrandImageCms())
		cms.POST("create/brands", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.CreateBrandCms())
		cms.POST("create/brandsimage", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.CreateBrandImageCms())
		cms.POST("/user/update-status", s.UserClient.ChangeStatus())
		cms.GET("/item/list", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ListItemsCMS())
		cms.GET("/item/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.DetailItemCMS())
		cms.POST("/register", s.UserClient.RegisterCMSUser())
		cms.POST("/login", s.UserClient.LoginCMSUser())
		cms.POST("/item/update-status", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.UpdateItemStatusCMS())
		cms.PATCH("/item/update/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.UpdateItemCMS())
		cms.GET("/collection", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.ListCollectionsCms())
		cms.PATCH("/collection/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.MarketPlaceClient.UpdateCollectionCms())
		cms.POST("/setting/item", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.SaveSettingItem())
		cms.DELETE("/setting/item/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.DeleteSettingItem())
		cms.GET("/setting", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.GetSetting())
		cms.POST("/daily-event", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.SaveDailyEvent())
		cms.PUT("/daily-event/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.UpdateDailyEvent())
		cms.DELETE("/daily-event/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.DeleteDailyEvent())
		cms.GET("/daily-event/:id", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.DetailDailyEvent())
		cms.GET("/daily-event", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.ListDailyEvent())
		cms.POST("/reward-department", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.SaveRewardDepartment())
		cms.POST("/reward", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.SaveReward())
	}

	commonRoute := router.Group("/common")
	{
		commonRoute.GET("/exchange", s.UserClient.ExchangeOPV())
		commonRoute.POST("/upload/:type", middleware.Authenticate(common.USER_TYPE), s.UserClient.Upload())
		commonRoute.POST("/item/upload", middleware.Authenticate(common.USER_TYPE), s.UserClient.UploadImage())
		commonRoute.POST("/cms/upload/:type", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.Upload())
		commonRoute.POST("/cms/item/upload", middleware.Authenticate(common.CMS_USER_TYPE), s.UserClient.UploadImage())
		commonRoute.POST("/search", s.UserClient.Search())
		commonRoute.GET("/cms/search/collection", s.UserClient.SearchCollectionCms())
		commonRoute.GET("/activities", s.UserClient.Activities())
	}

	notification := router.Group("/notification")
	{
		notification.POST("/change-status", middleware.Authenticate(common.USER_TYPE), s.UserClient.UpdateStatusNotification())
		notification.POST("/read-all", middleware.Authenticate(common.USER_TYPE), s.UserClient.ReadAll())
		notification.GET("/list", middleware.Authenticate(common.USER_TYPE), s.UserClient.ListNotification())
		notification.POST("/update-setting", middleware.Authenticate(common.USER_TYPE), s.UserClient.UpdateSetting())
		notification.GET("/setting", middleware.Authenticate(common.USER_TYPE), s.UserClient.NotificationSetting())
	}

	user := router.Group("/v1/user")
	{
		user.POST("/verify-email", middleware.Authenticate(common.USER_TYPE), s.UserClient.VerifyEmail())
		user.POST("/send-verify-email", middleware.Authenticate(common.USER_TYPE), s.UserClient.SendVerifyEmail())
		user.POST("/register", s.UserClient.Register())
		user.POST("/login", s.UserClient.LoginUser())
		user.POST("/check-exist", s.UserClient.CheckExistUser())
		user.POST("/create-nonce", s.UserClient.CreateNounce())
		user.POST("/send-otp", s.UserClient.SendMailOTP())
		user.POST("/refresh-token", middleware.Authenticate(common.USER_TYPE), s.UserClient.RefreshToken())
		user.GET("/detail", middleware.AuthenticateFalse(common.USER_TYPE), s.UserClient.DetailUser())
		user.GET("/:id", s.UserClient.DetailOtherUser())
		user.PATCH("/change-profile", middleware.Authenticate(common.USER_TYPE), s.UserClient.ChangeProfile())
		user.GET("/top-seller", s.UserClient.TopSeller())
		user.GET("/top-creator", s.UserClient.TopCreator())
		user.GET("/top-collection", s.UserClient.TopCollection())
		user.POST("/follow", middleware.Authenticate(common.USER_TYPE), s.UserClient.FollowingUser())
		user.POST("/unfollow", middleware.Authenticate(common.USER_TYPE), s.UserClient.UnfollowUser())
		user.GET("/profile-statistic", middleware.AuthenticateFalse(common.USER_TYPE), s.UserClient.ProfileStatistic())
		user.GET("/search-history", middleware.Authenticate(common.USER_TYPE), s.UserClient.SearchHistory())
		user.GET("/watch-list", middleware.Authenticate(common.USER_TYPE), s.UserClient.WatchList())
		user.GET("/:id/followers", middleware.Authenticate(common.USER_TYPE), s.UserClient.ListFollowers())
		user.GET("/:id/following", middleware.Authenticate(common.USER_TYPE), s.UserClient.ListFollowing())
		user.GET("/activities", middleware.AuthenticateFalse(common.USER_TYPE), s.UserClient.UserActivities())
		user.GET("/can-add-ref", middleware.Authenticate(common.USER_TYPE), s.UserClient.CanAddRef())
		user.GET("/get-user/:wallet_address", s.UserClient.DetailUserTransfer())
		user.GET("/explore-user", s.UserClient.ExploreUser())
		user.GET("/:id/followers-view", s.UserClient.ListFollowersView())
		user.GET("/:id/following-view", s.UserClient.ListFollowingView())
		user.GET("/daily-event", middleware.Authenticate(common.USER_TYPE), s.UserClient.GetUserDailyEvent())
		user.POST("/daily-event/check/:is_share", middleware.Authenticate(common.USER_TYPE), s.UserClient.UserDailyEventCheck())
		user.GET("/list_reward", middleware.Authenticate(common.USER_TYPE), s.UserClient.GetListReWard())
		user.GET("/reward/:id", middleware.Authenticate(common.USER_TYPE), s.UserClient.GetReWard())
		user.POST("/claim-reward", middleware.Authenticate(common.USER_TYPE), s.UserClient.ClaimReward())
		user.GET("/user_claim_reward", middleware.Authenticate(common.USER_TYPE), s.UserClient.GetUserReward())
	}

	item := router.Group("/item")
	{
		item.POST("/hide", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.CMSHideNft())
		item.POST("/unhide", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.CMSUnHideNft())
		item.POST("/sync", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.SyncItems())
		item.GET("/:smart_contract/:nft", middleware.AuthenticateFalse(common.USER_TYPE), s.MarketPlaceClient.DetailItem())
		item.GET("/event/:smart_contract/:nft", s.MarketPlaceClient.DetailItemEvent())
		item.GET("/event/statistic", s.MarketPlaceClient.ItemEventStatistic())
		item.GET("/list", middleware.AuthenticateFalse(common.USER_TYPE), s.MarketPlaceClient.ListItems())
		item.GET("/explore", middleware.AuthenticateFalse(common.USER_TYPE), s.MarketPlaceClient.ExploreNft())
		item.GET("/collection/item/:id", s.MarketPlaceClient.ListCollectionItems())
		item.GET("/list-category", s.MarketPlaceClient.ListCategory())
		item.GET("/list-collection", s.MarketPlaceClient.ListCollections())
		item.GET("/user-collection", middleware.AuthenticateFalse(common.USER_TYPE), s.MarketPlaceClient.ListSelfCollections())
		item.POST("/category", s.MarketPlaceClient.CreateCategory())
		item.POST("/collection", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.CreateCollection())
		item.GET("/collection/:id", s.MarketPlaceClient.CollectionDetail())
		item.GET("/collection/:id/stats", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.Stats())
		item.POST("/is-can-create-collection", s.MarketPlaceClient.IsCanCreateCollection())
		item.POST("/create-item", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.ItemCreator())
		item.PUT("/category/:id", s.MarketPlaceClient.UpdateCategory())
		item.PATCH("/:id", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.UpdateItem())
		item.PATCH("collection/:id", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.UpdateCollection())
		item.POST("/collection/like", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.LikeCollection())
		item.POST("/collection/dislike", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.DisLikeCollection())
		item.GET("/can-join-event/:smart_contract/:nft", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.CanJoinEvent())
		item.GET("/:smart_contract/:nft/history", s.MarketPlaceClient.History())
		item.GET("/collection/:id/history", s.MarketPlaceClient.CollectionHistory())
		item.GET("/collection/explore", s.MarketPlaceClient.ExploreCollection())
		item.GET("/collection-properties/:id", s.MarketPlaceClient.CollectionProperties())
		item.GET("/bid/:id", s.MarketPlaceClient.ItemBidDetail())
	}

	itemNft := router.Group("/item-nft")
	{
		itemNft.GET("/:smart_contract/:nft", s.MarketPlaceClient.DetailItemShow())
	}

	market := router.Group("/market-place")
	{
		market.POST("/request/report", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.RequestReport())
		market.GET("/explore-nft", middleware.AuthenticateFalse(common.USER_TYPE), s.UserClient.ExploreNft())
		market.POST("/sell-item", s.MarketPlaceClient.SellItem())
		market.GET("/brand/:id", s.MarketPlaceClient.Brand())
		market.GET("/list-brand", s.MarketPlaceClient.ListBrand())
		market.POST("/buy-item", s.MarketPlaceClient.BuyItem())
		market.POST("/staking-item", s.MarketPlaceClient.StakingItem())
		market.POST("/staking-complete", s.MarketPlaceClient.UpdateStakingStatus())
		market.POST("/vote", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.Vote())
		market.POST("/like", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.Like())
		market.POST("/dislike", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.DisLike())
		market.PUT("/:id/change-status", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.ChangeStatus())
		market.GET("/list-item", s.MarketPlaceClient.ListItem())
		market.GET("/list-item-event", s.MarketPlaceClient.ListItemEvent())
		market.GET("/most-liked/list", s.MarketPlaceClient.ListMostLikedItem())
		market.GET("/list-item-voted", s.MarketPlaceClient.ListVotedItem())
		market.GET("/item/:contract_address/:nft", s.MarketPlaceClient.DetailSellItem())
		market.GET("/transaction", s.MarketPlaceClient.ListTransaction())
		market.GET("/explore-statistic", s.MarketPlaceClient.ExploreStatistic())
		market.GET("/ref-network", middleware.Authenticate(common.USER_TYPE), s.UserClient.RefNetwork())
		market.GET("/setting/default", s.UserClient.DefaultSetting())
		market.POST("/get-list-item-token", middleware.Authenticate(common.USER_TYPE), s.MarketPlaceClient.GetListItemToken())
	}

	payment := router.Group("/payment")
	{
		payment.GET("/payment", s.UserClient.DefaultSetting())
	}

	router.GET("/storage/*filepath", s.UserClient.GetImageStatic)
	router.GET("/tmp/*filepath", s.UserClient.GetTmpImageStatic)
	if *Environment != "prod" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return router
}

func NewModule(sqlPath string) (*gorm.DB, error) {
	db, err := ConnectDb(sqlPath)
	if err != nil {
		return nil, err
	}
	return db, nil
}
