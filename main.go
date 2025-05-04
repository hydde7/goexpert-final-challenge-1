package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hydde7/goexpert-final-challenge-1/cmd"
	"github.com/hydde7/goexpert-final-challenge-1/internal/config"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.RedisAddr,
		Password: config.Redis.RedisPassword,
		DB:       config.Redis.RedisDB,
	})

	r := cmd.SetupRouter(redisClient)
	r.Run(fmt.Sprintf(":%s", config.App.Port))
}
