package limiter

import "time"

type Store interface {
	Allow(key string, limit int64, window time.Duration, blockKey string, blockTTL time.Duration) (bool, error)
}
