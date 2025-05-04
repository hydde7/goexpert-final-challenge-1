package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{client: client, ctx: context.Background()}
}

func (r *RedisStore) Allow(key string, limit int64, window time.Duration, blockKey string, blockTTL time.Duration) (bool, error) {
	if b, err := r.client.Exists(r.ctx, blockKey).Result(); err != nil {
		return false, err
	} else if b > 0 {
		return false, nil
	}

	count, err := r.client.Incr(r.ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		if err := r.client.Expire(r.ctx, key, window).Err(); err != nil {
			return false, err
		}
	}

	if count > limit {
		if err := r.client.Set(r.ctx, blockKey, "1", blockTTL).Err(); err != nil {
			return false, err
		}
		return false, nil
	}
	return true, nil
}
