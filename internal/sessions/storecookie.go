package sessions

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

type Store[T any] interface {
	// New returns a new named Session
	New(name string) *Session[T]
	// Get a named Session from the request
	Get(req *http.Request, name string) (*Session[T], error)
	// Save writes a Session to the ResponseWriter
	Save(w http.ResponseWriter, session *Session[T]) error
	// Destroy removes (expires) a named Session
	Destroy(w http.ResponseWriter, name string)
}

var _ Store[any] = &cookieStore[any]{}

// cookieStore stores Sessions in secure cookies (i.e. client-side)
type cookieStore[T any] struct {
	config *CookieConfig
	// encodes and decodes signed and optionally encrypted cookie values
	codecs []securecookie.Codec
}

// NewCookieStore returns a new Store that signs and optionally encrypts
// session state in http cookies.
func NewCookieStore[T any](config *CookieConfig, keyPairs ...[]byte) Store[T] {
	if config == nil {
		config = DefaultCookieConfig
	}

	return &cookieStore[T]{
		config: config,
		codecs: securecookie.CodecsFromPairs(keyPairs...),
	}
}

// New returns a new named Session
func (s *cookieStore[T]) New(name string) *Session[T] {
	return NewSession[T](s, name)
}

// Get returns the named Session from the Request. Returns an error if the
// session cookie cannot be found, the cookie verification fails, or an error
// occurs decoding the cookie value.
func (s *cookieStore[T]) Get(req *http.Request, name string) (session *Session[T], err error) {
	cookie, err := req.Cookie(name)
	if err == nil {
		session = s.New(name)
		err = securecookie.DecodeMulti(name, cookie.Value, &session.values, s.codecs...)
	}

	return session, err
}

// GetUserFromSession gets the cookies from the http.Request and gets the key (user ID) from the values
func (s *cookieStore[T]) GetUserFromSession(req *http.Request, name string) (string, error) {
	cookie, err := req.Cookie(name)
	if err != nil {
		return "", err
	}

	session := s.New(name)
	if err = securecookie.DecodeMulti(name, cookie.Value, &session.values, s.codecs...); err != nil {
		return "", err
	}

	for k := range session.values {
		return k, nil
	}

	return "", err
}

// Save adds or updates the Session on the response via a signed and optionally
// encrypted session cookie. Session Values are encoded into the cookie value
// and the session Config sets cookie properties.
func (s *cookieStore[T]) Save(w http.ResponseWriter, session *Session[T]) error {
	cookieValue, err := securecookie.EncodeMulti(session.Name(), &session.values, s.codecs...)
	if err != nil {
		return err
	}

	http.SetCookie(w, NewCookie(session.Name(), cookieValue, s.config))

	return nil
}

// Destroy deletes the Session with the given name by issuing an expired
// session cookie with the same name.
func (s *cookieStore[T]) Destroy(w http.ResponseWriter, name string) {
	http.SetCookie(w, NewCookie(name, "", &CookieConfig{MaxAge: -1, Path: s.config.Path}))
}
