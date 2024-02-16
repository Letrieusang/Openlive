package main

import (
	commonerr "api-openlive/common"
	"api-openlive/config"
	marketplacemodel "api-openlive/module/market_place/model"
	"fmt"
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

	var listItem []*marketplacemodel.Item
	errUser := db.Table(marketplacemodel.Item{}.TableName()).
		Where("status IN ?", []int64{commonerr.ITEM_READY_ON_SELL, commonerr.ITEM_STATUS_ON_MARKET}).
		Find(&listItem).Error
	if errUser != nil {
		log.Fatalln(errUser)
		return
	}

	for _, v := range listItem {
		v.IsApproval = commonerr.APPROVE
		errUpdate := db.Table(marketplacemodel.Item{}.TableName()).Where("id = ?", v.Id).Updates(v).Error
		if errUpdate != nil {
			log.Fatalln(errUpdate)
		}
		fmt.Println("update success")
	}

	log.Println("ok!!!")

}
