package main

import (
	"api-openlive/config"
	usermodel "api-openlive/module/user/model"
	"api-openlive/utils"
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

	var listUser []*usermodel.User
	errUser := db.Table(usermodel.User{}.TableName()).
		Find(&listUser).Error
	if errUser != nil {
		log.Fatalln(errUser)
		return
	}

	for _, v := range listUser {
		refCode := utils.GetMD5Hash(fmt.Sprintf("%v.%v", v.WalletAddress, v.Id))
		v.Ref = utils.BuildRefLink(refCode)
		errUpdate := db.Table(usermodel.User{}.TableName()).Where("id = ?", v.Id).Updates(v).Error
		if errUpdate != nil {
			log.Fatalln(errUpdate)
		}
		fmt.Println("update success")
	}

	log.Println("ok!!!")

}
