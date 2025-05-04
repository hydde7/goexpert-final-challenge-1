package config

import "github.com/hydde7/goexpert-final-challenge-1/internal/utils"

var Redis = &redisConfig{
	RedisAddr:     utils.GetStringEnv("REDIS_ADDR", "localhost:6379"),
	RedisPassword: utils.GetStringEnv("REDIS_PASSWORD", ""),
	RedisDB:       utils.GetIntEnv("REDIS_DB", 0),
}

type redisConfig struct {
	RedisAddr     string `json:"REDIS_ADDR"`
	RedisPassword string `json:"REDIS_PASSWORD"`
	RedisDB       int    `json:"REDIS_DB"`
}
