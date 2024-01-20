package sessions

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

// PersistentStore is defining an interface for session store
type PersistentStore interface {
	Exists(ctx context.Context, userID string) (int64, error)
	GetSession(ctx context.Context, userID string) (string, error)
	StoreSession(ctx context.Context, sessionID string, userID string) error
	DeleteSession(ctx context.Context, userID string) error
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
func (s *persistentStore) Exists(ctx context.Context, userID string) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Exists(ctx, userID).Result()
}

// GetSession checks to see if there is an existing session for the user
func (s *persistentStore) GetSession(ctx context.Context, userID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Get(ctx, userID).Result()
}

// StoreSession is used to store a session in the store
func (s *persistentStore) StoreSession(ctx context.Context, sessionID string, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Set(ctx, userID, sessionID, 0).Err()
}

// DeleteSession is used to delete a session from the store
func (s *persistentStore) DeleteSession(ctx context.Context, userID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.client.Del(ctx, userID).Err()
}
