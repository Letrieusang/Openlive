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

type LikeCount struct {
	Id        int64 `json:"id"`
	TotalLike int64 `json:"total_like"`
}

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

	var listItem []LikeCount
	errUser := db.Table(marketplacemodel.Like{}.TableName()).
		Select("COUNT(id) as total_like, item_id as id").
		Where("is_deleted = ?", commonerr.ACTIVE).
		Group("item_id").
		Find(&listItem).Error
	if errUser != nil {
		log.Fatalln(errUser)
		return
	}

	for _, v := range listItem {
		errUpdate := db.Table(marketplacemodel.Item{}.TableName()).Where("id = ?", v.Id).Updates(v).Error
		if errUpdate != nil {
			log.Fatalln(errUpdate)
		}
		fmt.Println("update success")
	}

	log.Println("ok!!!")
}
