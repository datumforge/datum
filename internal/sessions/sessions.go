package sessions

import (
	"net/http"
	"sync"

	"github.com/datumforge/datum/internal/utils/ulids"
)

type Config struct {
	// SigningKey must be a 16, 32, or 64 character string used to encode the cookie
	SigningKey string `yaml:"signingKey" split_words:"true" default:"my-signing-secret"` // $DATUM_SESSIONS_SIGNING_KEY
	// EncryptionKey must be a 16, 32, or 64 character string used to encode the cookie
	EncryptionKey string `yaml:"encryptionKey" split_words:"true" default:"encryptionsecret"` // $DATUM_SESSIONS_ENCRYPTION_KEY
}

// Session represents state values maintained in a sessions Store
type Session[T any] struct {
	name   string
	values map[string]T
	mu     sync.Mutex
	store  Store[T]
}

// NewSession returns a new Session.
func NewSession[T any](store Store[T], name string) *Session[T] {
	return &Session[T]{
		name:   name,
		values: make(map[string]T),
		store:  store,
	}
}

func (s *Session[T]) SetName(name string) {
	s.name = name
}

// Name returns the name of the session.
func (s *Session[T]) Name() string {
	return s.name
}

// Set sets a key/value pair in the session state.
func (s *Session[T]) Set(key string, value T) {
	s.values[key] = value
}

// Get returns the state value for the given key.
func (s *Session[T]) Get(key string) T {
	return s.values[key]
}

// GetOk returns the state value for the given key and whether they key exists.
func (s *Session[T]) GetOk(key string) (T, bool) {
	value, ok := s.values[key]
	return value, ok
}

// Save adds or updates the session. Identical to calling
// store.Save(w, session).
func (s *Session[T]) Save(w http.ResponseWriter) error {
	return s.store.Save(w, s)
}

// Destroy destroys the session. Identical to calling
// store.Destroy(w, session.name).
func (s *Session[T]) Destroy(w http.ResponseWriter) {
	s.store.Destroy(w, s.name)
}

// GenerateSessionID returns a random ulid
func GenerateSessionID() string {
	return ulids.New().String()
}
