package sessions

import (
	"context"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

const (
	DefaultSessionName = "__session_id"
)

// SessionContextKey is the context key for the session-context
var SessionContextKey = &ContextKey{"SessionContextKey"}

// ContextKey is the key name for the additional context
type ContextKey struct {
	name string
}

type Config struct {
	SigningKey    string `yaml:"signingKey" split_words:"true" default:"my-signing-secret"`   // $DATUM_SESSIONS_SIGNING_KEY
	EncryptionKey string `yaml:"encryptionKey" split_words:"true" default:"encryptionsecret"` // $DATUM_SESSIONS_ENCRYPTION_KEY
}

// Session represents state values maintained in a sessions Store.
type Session struct {
	name   string
	values map[string]string
	mu     sync.Mutex
	// convenience methods Save and Destroy use store
	store CookieStore
}

// New returns a new Session.
func New(store CookieStore) *Session {
	return &Session{
		name:   DefaultSessionName,
		values: make(map[string]string),
		store:  store,
	}
}

// NewWithName returns a new Session with the provided name
func NewWithName(store CookieStore, name string) *Session {
	if name == "" {
		name = DefaultSessionName
	}

	return &Session{
		name:   name,
		values: make(map[string]string),
		store:  store,
	}
}

// Name returns the name of the session.
func (s *Session) Name() string {
	return s.name
}

// Set sets a key/value pair in the session state.
func (s *Session) Set(key string, value string) {
	s.values[key] = value
}

// Get returns the state value for the given key.
func (s *Session) Get(key string) string {
	return s.values[key]
}

// GetOk returns the state value for the given key and whether they key exists.
func (s *Session) GetOk(key string) (string, bool) {
	value, ok := s.values[key]
	return value, ok
}

// Save adds or updates the session. Identical to calling
// store.Save(w, session).
func (s *Session) Save(w http.ResponseWriter) error {
	return s.store.Save(w, s)
}

// Destroy destroys the session. Identical to calling
// store.Destroy(w, session.name).
func (s *Session) Destroy(w http.ResponseWriter) {
	s.store.Destroy(w, s.name)
}

func GenerateSessionID() string {
	id, _ := uuid.NewRandom()

	return id.String()
}

// Token returns the session token from the context
func Token(ctx context.Context) map[string]string {
	sd := getSessionDataFromContext(ctx)

	sd.mu.Lock()
	defer sd.mu.Unlock()

	return sd.values
}

// addSessionDataToContext adds the session details to the context
func (s *Session) addSessionDataToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, SessionContextKey, s)
}

// getSessionDataFromContext gets the session information from the context
func getSessionDataFromContext(ctx context.Context) *Session {
	c, ok := ctx.Value(SessionContextKey).(*Session)
	if !ok {
		panic("no session data in context")
	}

	return c
}
