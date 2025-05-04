package limiter_test

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hydde7/goexpert-final-challenge-1/internal/config"
	"github.com/hydde7/goexpert-final-challenge-1/internal/limiter"
	"github.com/stretchr/testify/assert"
)

func setupRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: config.Redis.RedisAddr,
	})
	client.FlushDB(context.Background())
	return client
}

func TestRedisStore_Allow(t *testing.T) {
	client := setupRedisClient()
	store := limiter.NewRedisStore(client)

	t.Run("allow request within limit", func(t *testing.T) {
		key := "test_key"
		blockKey := "block_key"
		limit := int64(5)
		window := time.Second * 10
		blockTTL := time.Second * 30

		for i := int64(0); i < limit; i++ {
			allowed, err := store.Allow(key, limit, window, blockKey, blockTTL)
			assert.NoError(t, err)
			assert.True(t, allowed)
		}
	})

	t.Run("block request exceeding limit", func(t *testing.T) {
		key := "test_key_exceed"
		blockKey := "block_key_exceed"
		limit := int64(3)
		window := time.Second * 10
		blockTTL := time.Second * 30

		for i := int64(0); i < limit; i++ {
			allowed, err := store.Allow(key, limit, window, blockKey, blockTTL)
			assert.NoError(t, err)
			assert.True(t, allowed)
		}

		allowed, err := store.Allow(key, limit, window, blockKey, blockTTL)
		assert.NoError(t, err)
		assert.False(t, allowed)

		exists, err := client.Exists(context.Background(), blockKey).Result()
		assert.NoError(t, err)
		assert.Equal(t, int64(1), exists)
	})

	t.Run("block request when blockKey exists", func(t *testing.T) {
		key := "test_key_block"
		blockKey := "block_key_block"
		limit := int64(3)
		window := time.Second * 10
		blockTTL := time.Second * 30

		err := client.Set(context.Background(), blockKey, "1", blockTTL).Err()
		assert.NoError(t, err)

		allowed, err := store.Allow(key, limit, window, blockKey, blockTTL)
		assert.NoError(t, err)
		assert.False(t, allowed)
	})
}
