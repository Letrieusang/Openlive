package main

import (
	"api-openlive/blockchain"
	"api-openlive/config"
	"api-openlive/cronjob"
	markettransport "api-openlive/module/market_place"
	marketbusiness "api-openlive/module/market_place/business"
	marketstorage "api-openlive/module/market_place/storage"
	usertransport "api-openlive/module/user"
	bussiness "api-openlive/module/user/business"
	userstorage "api-openlive/module/user/storage"
	"api-openlive/utils"
	"fmt"
	"log"
	"os"
	sync2 "sync"
	"time"
)

type server struct {
	UserClient        *usertransport.UserHandler
	MarketPlaceClient *markettransport.MarketPlaceHandler
}

func Init() {
	cfg := config.GetConfig()
	db, err := NewModule(cfg.GetString("server.sql_path"))
	rdb := NewRedisClient(cfg.GetString("server.rds_addr"), cfg.GetString("server.rds_pwd"))
	if err != nil {
		panic(err)
	}
	userSvc := &bussiness.UserController{
		Store:              userstorage.NewUserStore(db),
		MarketPlaceStorage: marketstorage.NewMarketPlaceStore(db),
		Cache:              rdb,
	}

	mutex := &sync2.RWMutex{}
	convertRate, err := utils.GetExchangeRates("tether", 90)
	if err != nil {
		log.Println("error when get convert rate nft", err)
	}
	log.Println("get convert rate success", err)
	utils.SetRate(convertRate, mutex)

	marketSvc := &marketbusiness.MarketPlaceController{
		Store:     marketstorage.NewMarketPlaceStore(db),
		UserStore: userstorage.NewUserStore(db),
	}
	s := server{
		UserClient: &usertransport.UserHandler{
			Client: userSvc,
		},
		MarketPlaceClient: &markettransport.MarketPlaceHandler{
			Client: marketSvc,
			Blc:    blockchain.NewClientHttp(),
			OptimisticLock: &utils.OptimisticLock{
				Client:   rdb,
				Prefix:   "market_place_sv",
				Duration: 1 * time.Minute,
			},
		},
	}
	sync := blockchain.NewSynchronization(config.GetBlockchainConfig(), marketSvc, userSvc, rdb)
	go sync.StartSync(0)
	cron := cronjob.NewCronJob(sync, mutex)
	go cron.StartCronJob()
	r := s.NewRouter()
	os.Setenv("ENV", *Environment)
	r.Run(fmt.Sprintf(":%v", cfg.GetString("server.port")))
}
