package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func NewRedisClient(rdsAddr string, rdsPwd string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdsAddr,
		Password: rdsPwd, // no password set
		DB:       0,      // use default DB
	})

	log.Println("RDS client ready")
	tick := time.NewTicker(10 * time.Minute)
	ctx := context.Background()
	go func(client *redis.Client) {
		for {
			<-tick.C
			if err := client.Ping(ctx).Err(); err != nil {
				panic(err)
			}
		}
	}(rdb)

	return rdb
}
