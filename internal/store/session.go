package store

import (
	"encoding/base64"
	"net/http"

	"github.com/gorilla/sessions"
)

// SessionStore interface specifies a single method `Get` that takes an
// `http.Request` and a string as parameters and returns a `*sessions.Session` and an error. Any type that implements this `Get` method can be considered as implementing the `SessionStore` interface
type SessionStore interface {
	Get(r *http.Request, name string) (*sessions.Session, error)
}

// GetSession function retrieves a session from a session store based on a cookie name and a request
func GetSession(r *http.Request, cookieName string, sessionStore SessionStore) (*sessions.Session, error) {
	return sessionStore.Get(r, cookieName)
}

// SetSessionB64 function sets a base64-encoded session value in a cookie and returns the session ID
func SetSessionB64(r *http.Request, w http.ResponseWriter, body []byte, cookieName, valueName string, sessionStore SessionStore) (string, error) {
	cookieValue := base64.StdEncoding.EncodeToString(body)

	if err := SetSession(r, w, cookieValue, cookieName, valueName, sessionStore); err != nil {
		return "", err
	}

	return cookieValue, nil
}

// SetSession function sets a session value in a session store and also sets a corresponding cookie in the response
func SetSession(r *http.Request, w http.ResponseWriter, value, cookieName, valueName string, sessionStore SessionStore) error {
	// set the cookie
	session, err := sessionStore.Get(r, cookieName)
	if err != nil {
		return err
	}

	session.Values[valueName] = value

	return session.Save(r, w)
}

// RemoveSession function removes a session from the session store based on the provided cookie name
func RemoveSession(r *http.Request, w http.ResponseWriter, cookieName string, sessionStore SessionStore) error {
	session, err := sessionStore.Get(r, cookieName)
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1

	return session.Save(r, w)
}
