package sessions

import (
	"net/http"

	"github.com/gorilla/securecookie"

	"github.com/datumforge/datum/internal/cookies"
)

// A CookieStore manages creating, accessing, writing, and expiring Sessions.
type CookieStore interface {
	// New returns a new named Session
	New(name string) *Session
	// Get a named Session from the request
	Get(req *http.Request, name string) (*Session, error)
	// Save writes a Session to the ResponseWriter
	Save(w http.ResponseWriter, session *Session) error
	// Destroy removes (expires) a named Session
	Destroy(w http.ResponseWriter, name string)
	// GetUserFromSession with the provided cookie name
	GetUserFromSession(req *http.Request, name string) (string, error)
}

var _ CookieStore = &cookieStore{}

// cookieStore stores Sessions in secure cookies (i.e. client-side)
type cookieStore struct {
	config *cookies.CookieConfig
	// encodes and decodes signed and optionally encrypted cookie values
	codecs []securecookie.Codec
}

// NewCookieStore returns a new Store that signs and optionally encrypts
// session state in http cookies.
func NewCookieStore(config *cookies.CookieConfig, keyPairs ...[]byte) CookieStore {
	if config == nil {
		config = &cookies.DefaultCookieConfig
	}

	return &cookieStore{
		config: config,
		codecs: securecookie.CodecsFromPairs(keyPairs...),
	}
}

// New returns a new named Session.
func (s *cookieStore) New(name string) *Session {
	return NewWithName(s, name)
}

// NewSessionCookie creates a cookie from a session id
func NewSessionCookie(session string) *http.Cookie {
	config := cookies.DefaultCookieConfig
	config.Name = DefaultSessionName

	return cookies.NewCookie(config, session)
}

// NewDebugSessionCookie creates a debug cookie from a session id
func NewDebugSessionCookie(session string) *http.Cookie {
	config := cookies.DebugOnlyCookieConfig
	config.Name = DefaultSessionName

	return cookies.NewCookie(config, session)
}

// Get returns the named Session from the Request. Returns an error if the
// session cookie cannot be found, the cookie verification fails, or an error
// occurs decoding the cookie value.
func (s *cookieStore) Get(req *http.Request, name string) (session *Session, err error) {
	cookie, err := req.Cookie(name)
	if err == nil {
		session = s.New(name)
		err = securecookie.DecodeMulti(name, cookie.Value, &session.values, s.codecs...)
	}

	return session, err
}

// GetUserFromSession gets the cookies from the http.Request and gets the key (user ID) from the values
func (s *cookieStore) GetUserFromSession(req *http.Request, name string) (string, error) {
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
func (s *cookieStore) Save(w http.ResponseWriter, session *Session) error {
	cookieValue, err := securecookie.EncodeMulti(session.Name(), &session.values, s.codecs...)
	if err != nil {
		return err
	}

	cookies.SetCookie(w, cookieValue, session.Name(), *s.config)

	return nil
}

// Destroy deletes the Session with the given name by issuing an expired
// session cookie with the same name.
func (s *cookieStore) Destroy(w http.ResponseWriter, name string) {
	cookies.RemoveCookie(w, name, *s.config)
}
