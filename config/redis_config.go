package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func NewRedisInstance() *redis.Client {
	addr := os.Getenv("REDIS_URL")
	// db := os.Getenv("REDIS_DB")

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	if pong := rdb.Ping(context.Background()); pong.String() != "ping: PONG" {
		log.Println("-------------Error connection redis ----------:", pong)
		return nil
	} else {
		log.Println("-------CONNECTED REDIS ----------")
	}
	return rdb
}
