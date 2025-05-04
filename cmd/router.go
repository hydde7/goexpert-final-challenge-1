package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/hydde7/goexpert-final-challenge-1/internal/config"
	"github.com/hydde7/goexpert-final-challenge-1/internal/limiter"
)

func SetupRouter(redisClient *redis.Client) *gin.Engine {
	var store limiter.Store
	if config.App.UseRedis {
		store = limiter.NewRedisStore(redisClient)
	} else {
		store = limiter.NewInMemoryStore()
	}

	r := gin.Default()
	r.Use(RateLimiter(store, Config{
		IPLimit:            config.App.IPLimit,
		TokenLimit:         config.App.TokenLimit,
		BlockDurationIP:    config.App.BlockDurationIP,
		BlockDurationToken: config.App.BlockDurationToken,
		Window:             config.App.RateLimitWindow,
		IPBlockMode:        config.App.IPBlockMode,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return r
}
