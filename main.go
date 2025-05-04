package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hydde7/goexpert-final-challenge-1/cmd"
	"github.com/hydde7/goexpert-final-challenge-1/internal/config"
	"github.com/sirupsen/logrus"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.RedisAddr,
		Password: config.Redis.RedisPassword,
		DB:       config.Redis.RedisDB,
	})
	status := redisClient.Ping(redisClient.Context())
	logrus.Info(status)
	r := cmd.SetupRouter(redisClient)
	r.Run(fmt.Sprintf(":%s", config.App.Port))
}
