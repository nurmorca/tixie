package cache

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient(context context.Context, config RConfig) *redis.Client {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr, // use default Addr
		Password: "",   // no password set
		DB:       0,    // use default DB
	})

	_, err := rdb.Ping(context).Result()
	if err != nil {
		log.Error("unable to ping redis: %v\n", err)
		panic(err)
	}

	return rdb
}
