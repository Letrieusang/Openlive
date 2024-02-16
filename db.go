package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func ConnectDb(sqlPath string) (*gorm.DB, error) {
	logMod := logger.Info
	if *Environment == "prod" {
		logMod = logger.Error
	}
	db, err := gorm.Open(mysql.Open(sqlPath), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logMod),
	})

	if err != nil {
		log.Fatalln("Error mariaDB:", err)
		return nil, err
	}

	tick := time.NewTicker(15 * time.Minute)
	go func(engine *gorm.DB) {
		for {
			select {
			case <-tick.C:
				sqlDB, err := engine.DB()
				if err != nil {
					log.Print("sql select DB")
				}
				err = sqlDB.Ping()
				if err != nil {
					log.Println("sql can not ping")
				}
			}
		}
	}(db)
	log.Print("Connected to: ", sqlPath)

	return db, nil

}
