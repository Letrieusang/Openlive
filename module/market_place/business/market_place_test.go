package business

import (
	marketstorage "api-openlive/module/market_place/storage"
	userstorage "api-openlive/module/user/storage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"testing"
)

func intEnv() (*MarketPlaceController, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3308)/opv_market?parseTime=true"), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln("Error mariaDB:", err)
		return nil, err
	}

	marketSvc := &MarketPlaceController{
		Store:     marketstorage.NewMarketPlaceStore(db),
		UserStore: userstorage.NewUserStore(db),
	}

	return marketSvc, nil
}

func TestCollectionNftProperties(t *testing.T) {
	sv, err := intEnv()
	if err != nil {
		log.Fatalln(err)
		return
	}

	res, err := sv.CollectionItemProperties(198)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println(res)
}
