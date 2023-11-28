package store

import (
	"encoding/base64"
	"net/http"
)

// Config is used to store configuration settings for setting and removing cookies
type Config struct {
	// Domain for defines the host to which the cookie will be sent.
	Domain string `cfg:"domain"`
	// Path that must exist in the requested URL for the browser to send the Cookie header.
	Path string `cfg:"path"`
	// MaxAge the number of seconds until the cookie expires.
	MaxAge int `cfg:"max_age"`
	// Secure to cookie only sent over HTTPS.
	Secure bool `cfg:"secure"`
	// SameSite for Lax 2, Strict 3, None 4.
	SameSite http.SameSite `cfg:"same_site"`
	// HttpOnly for true for not accessible by JavaScript.
	HttpOnly bool `cfg:"http_only"` //nolint:stylecheck
}

// GetCookie function retrieves a specific cookie from an HTTP request
func GetCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	return r.Cookie(cookieName)
}

// SetCookieB64 function sets a base64-encoded cookie with the given name and value in the HTTP response
func SetCookieB64(w http.ResponseWriter, body []byte, cookieName string, v Config) string {
	cookieValue := base64.StdEncoding.EncodeToString(body)
	// set the cookie
	SetCookie(w, cookieValue, cookieName, v)

	return cookieValue
}

// SetCookie function sets a cookie with the given value and name
func SetCookie(w http.ResponseWriter, value string, cookieName string, v Config) {
	// set the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    value,
		Domain:   v.Domain,
		Path:     v.Path,
		MaxAge:   v.MaxAge,
		Secure:   v.Secure,
		SameSite: v.SameSite,
		HttpOnly: v.HttpOnly,
	})
}

// RemoveCookie function removes a cookie from the HTTP response
func RemoveCookie(w http.ResponseWriter, cookieName string, v Config) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Domain:   v.Domain,
		Path:     v.Path,
		MaxAge:   -1,
		Secure:   v.Secure,
		SameSite: v.SameSite,
		HttpOnly: v.HttpOnly,
	})
}
