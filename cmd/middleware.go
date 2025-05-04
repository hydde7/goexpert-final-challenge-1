package cmd

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hydde7/goexpert-final-challenge-1/internal/limiter"
	"github.com/sirupsen/logrus"
)

type Config struct {
	IPLimit            int64
	TokenLimit         int64
	BlockDurationIP    time.Duration
	BlockDurationToken time.Duration
	Window             time.Duration
	IPBlockMode        bool
}

func RateLimiter(store limiter.Store, cfg Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			key      string
			blockKey string
			limit    int64
			blockTTL time.Duration
		)

		if cfg.IPBlockMode {
			ip := c.ClientIP()
			key = "ip:" + ip
			blockKey = "block:ip:" + ip
			limit = cfg.IPLimit
			blockTTL = cfg.BlockDurationIP
		} else {
			token := c.GetHeader("API_KEY")
			if token != "" {
				key = "token:" + token
				blockKey = "block:token:" + token
				limit = cfg.TokenLimit
				blockTTL = cfg.BlockDurationToken
			} else {
				c.AbortWithStatusJSON(401, gin.H{"message": "API_KEY header is required"})
				return
			}
		}

		allowed, err := store.Allow(key, limit, cfg.Window, blockKey, blockTTL)
		if err != nil {
			logrus.Errorf("error checking rate limit: %v", err)
			c.AbortWithStatusJSON(500, gin.H{"error": "internal server error"})
			return
		}
		if !allowed {
			c.AbortWithStatusJSON(429, gin.H{"message": "you have reached the maximum number of requests or actions allowed within a certain time frame"})
			return
		}
		c.Next()
	}
}
