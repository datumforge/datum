package sessions

import (
	"context"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultExpiration = 10 * time.Minute
)

// PersistentStore is defining an interface for session store
type PersistentStore interface {
	Exists(ctx context.Context, key string) (int64, error)
	GetSession(ctx context.Context, key string) (string, error)
	StoreSession(ctx context.Context, key, value string) error
	DeleteSession(ctx context.Context, key string) error
}

var _ PersistentStore = &persistentStore{}

// persistentStore stores Sessions in a persisent data store (redis)
type persistentStore struct {
	client *redis.Client
	mu     sync.Mutex
}

// NewStore returns a new Store that stores to a persistent backend (redis)
func NewStore(client *redis.Client) PersistentStore {
	return &persistentStore{
		client: client,
	}
}

// Exists checks to see if there is an existing session for the user
func (s *persistentStore) Exists(ctx context.Context, key string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Exists(ctx, key).Result()
}

// GetSession checks to see if there is an existing session for the user
func (s *persistentStore) GetSession(ctx context.Context, key string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Get(ctx, key).Result()
}

// StoreSession is used to store a session in the store
func (s *persistentStore) StoreSession(ctx context.Context, key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Set(ctx, key, value, defaultExpiration).Err()
}

// DeleteSession is used to delete a session from the store
func (s *persistentStore) DeleteSession(ctx context.Context, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Del(ctx, userID).Err()
}
