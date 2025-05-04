package config

import (
	"time"

	"github.com/hydde7/goexpert-final-challenge-1/internal/utils"
)

var App = &appConfig{
	Port:               utils.GetStringEnv("APP_PORT", "8080"),
	UseRedis:           utils.GetBoolEnv("APP_USE_REDIS", false),
	IPLimit:            utils.GetInt64Env("APP_IP_LIMIT", 10),
	TokenLimit:         utils.GetInt64Env("APP_TOKEN_LIMIT", 10),
	RateLimitWindow:    utils.GetDurationEnv("APP_RATE_LIMIT_WINDOW", time.Minute),
	BlockDurationIP:    utils.GetDurationEnv("APP_BLOCK_DURATION_IP", time.Minute),
	BlockDurationToken: utils.GetDurationEnv("APP_BLOCK_DURATION_TOKEN", time.Minute),
	IPBlockMode:        utils.GetBoolEnv("APP_IP_BLOCK_MODE", true),
}

type appConfig struct {
	Port               string        `json:"APP_PORT"`
	UseRedis           bool          `json:"APP_USE_REDIS"`
	IPLimit            int64         `json:"APP_IP_LIMIT"`
	TokenLimit         int64         `json:"APP_TOKEN_LIMIT"`
	RateLimitWindow    time.Duration `json:"APP_RATE_LIMIT_WINDOW"`
	BlockDurationIP    time.Duration `json:"APP_BLOCK_DURATION_IP"`
	BlockDurationToken time.Duration `json:"APP_BLOCK_DURATION_TOKEN"`
	IPBlockMode        bool          `json:"APP_IP_BLOCK_MODE"`
}
