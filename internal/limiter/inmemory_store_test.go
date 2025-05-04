package limiter_test

import (
	"testing"
	"time"

	"github.com/hydde7/goexpert-final-challenge-1/internal/limiter"
)

func TestInMemoryStore(t *testing.T) {
	store := limiter.NewInMemoryStore()
	limit := int64(2)
	window := 100 * time.Millisecond
	blockTTL := 100 * time.Millisecond

	key := "test"
	blockKey := "block:test"

	for i := 0; i < 2; i++ {
		ok, _ := store.Allow(key, limit, window, blockKey, blockTTL)
		if !ok {
			t.Errorf("esperava allowed para requisição %d", i+1)
		}
	}
	if ok, _ := store.Allow(key, limit, window, blockKey, blockTTL); ok {
		t.Error("esperava blocked na terceira requisição")
	}
	time.Sleep(150 * time.Millisecond)
	if ok, _ := store.Allow(key, limit, window, blockKey, blockTTL); !ok {
		t.Error("esperava allowed após tempo de bloqueio")
	}
}
