package main

import (
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"time"
)

func RemoveTmpFolder() {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -2).Format("2006-01-02")

	if _, errFound := os.Stat("tmp/" + yesterday); os.IsNotExist(errFound) {
		return
	}

	err := os.RemoveAll("tmp/" + yesterday)
	if err != nil {
		log.Println("Error remove temp file")
	}

	log.Println("Delete temp folder success")
	return
}

func CronRemoveTemp() {
	log.Println("----------------------")
	log.Println("start cron every ")
	c := cron.New()
	c.AddFunc("30 2 * * *", RemoveTmpFolder)
	c.Start()
}
