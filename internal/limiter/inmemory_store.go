package limiter

import (
	"sync"
	"time"
)

type memoryEntry struct {
	count     int64
	expiresAt time.Time
}

type InMemoryStore struct {
	mu      sync.Mutex
	counts  map[string]*memoryEntry
	blocked map[string]time.Time
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		counts:  make(map[string]*memoryEntry),
		blocked: make(map[string]time.Time),
	}
}

func (m *InMemoryStore) Allow(key string, limit int64, window time.Duration, blockKey string, blockTTL time.Duration) (bool, error) {
	now := time.Now()
	m.mu.Lock()
	defer m.mu.Unlock()

	if t, ok := m.blocked[blockKey]; ok {
		if now.Before(t) {
			return false, nil
		}
		delete(m.blocked, blockKey)
	}

	entry, ok := m.counts[key]
	if !ok || now.After(entry.expiresAt) {
		entry = &memoryEntry{count: 0, expiresAt: now.Add(window)}
		m.counts[key] = entry
	}
	entry.count++
	if entry.count > limit {
		m.blocked[blockKey] = now.Add(blockTTL)
		return false, nil
	}
	return true, nil
}
