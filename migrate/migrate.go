package main

import (
	"api-openlive/config"
	marketplacemodel "api-openlive/module/market_place/model"
	usermodel "api-openlive/module/user/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func init() {
	env := "dev"
	if os.Getenv("ENV") == "local" {
		env = os.Getenv("ENV")
	}
	config.Init(env)
}

func main() {
	svConfig := config.GetConfig()
	db, err := gorm.Open(mysql.Open(svConfig.GetString("server.sql_path")), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}

	//migrate here
	db.AutoMigrate(&usermodel.User{})
	db.AutoMigrate(&marketplacemodel.Item{})
	db.AutoMigrate(&marketplacemodel.Category{})
	db.AutoMigrate(&marketplacemodel.MarketPlace{})
	db.AutoMigrate(&marketplacemodel.MarketPlaceSold{})
	db.AutoMigrate(&marketplacemodel.Vote{})
	db.AutoMigrate(&marketplacemodel.Staking{})
	db.AutoMigrate(&marketplacemodel.Like{})
	db.AutoMigrate(&marketplacemodel.ItemCreator{})
	db.AutoMigrate(&marketplacemodel.Collection{})
	db.AutoMigrate(&marketplacemodel.TransactionLog{})
	db.AutoMigrate(&usermodel.Notification{})
	db.AutoMigrate(&usermodel.CMSUser{})
	db.AutoMigrate(&marketplacemodel.CollectionLike{})
	db.AutoMigrate(&usermodel.RefUser{})
	db.AutoMigrate(&marketplacemodel.Setting{})
	db.AutoMigrate(&marketplacemodel.BlockChainBlock{})
	db.AutoMigrate(&usermodel.UserFollow{})
	db.AutoMigrate(&usermodel.UserSearch{})
	db.AutoMigrate(&usermodel.Activities{})
	db.AutoMigrate(&usermodel.NotificationSetting{})
	db.AutoMigrate(&marketplacemodel.ItemJoinEventHistory{})
	db.AutoMigrate(&usermodel.SystemSetting{})
	db.AutoMigrate(&usermodel.BankIncome{})
	db.AutoMigrate(&marketplacemodel.BrandImage{})
	db.AutoMigrate(&marketplacemodel.Brand{})
	db.AutoMigrate(&marketplacemodel.AddBannerCmsRequest{})
	db.AutoMigrate(&marketplacemodel.AuctionBid{})
	db.AutoMigrate(&usermodel.DailyEvent{})
	db.AutoMigrate(&usermodel.UserDailyEvent{})
	db.AutoMigrate(&usermodel.Reward{})
	db.AutoMigrate(&usermodel.RewardDepartment{})
	db.AutoMigrate(&usermodel.UserReward{})
	log.Println("migrate success")
}
