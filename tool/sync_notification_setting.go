package main

import (
	"api-openlive/config"
	usermodel "api-openlive/module/user/model"
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

	var users []*usermodel.User
	errUser := db.Table(usermodel.User{}.TableName()).Select("id").
		Find(&users).Error
	if errUser != nil {
		log.Fatalln(errUser)
		return
	}

	setting := make([]usermodel.NotificationSetting, 0, len(users))
	for _, v := range users {
		setting = append(setting, usermodel.NotificationSetting{
			UserId:              v.Id,
			ItemSold:            1,
			BiddingActivities:   1,
			AuctionExpiration:   1,
			OutBid:              1,
			NewLetter:           1,
			FollowingActivities: 1,
			LikeAndFollow:       1,
			AuctionWin:          1,
			YourActivities:      1,
		})
	}

	errCreate := db.Table(usermodel.NotificationSetting{}.TableName()).CreateInBatches(&setting, 100).Error
	if errCreate != nil {
		log.Fatalln(errCreate)
	}
	fmt.Println("update success")
	log.Println("ok!!!")

}
