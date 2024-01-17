package cookies

import (
	"encoding/base64"
	"net/http"
	"time"
)

// CookieConfig configures http.Cookie creation
type CookieConfig struct {
	// Name is the desired cookie name
	Name string
	// Domain sets the cookie domain. Defaults to the host name of the responding
	// server when left zero valued
	Domain string
	// Path sets the cookie path. Defaults to the path of the URL responding to
	// the request when left zero valued
	Path string
	// MaxAge=0 means no 'Max-Age' attribute should be set
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	// Cookie 'Expires' will be set (or left unset) according to MaxAge
	MaxAge int
	// HTTPOnly indicates whether the browser should prohibit a cookie from
	// being accessible via Javascript. Recommended true.
	HTTPOnly bool
	// Secure flag indicating to the browser that the cookie should only be
	// transmitted over a TLS HTTPS connection. Recommended true in production
	Secure bool
	// SameSite attribute modes indicates that a browser not send a cookie in
	// cross-site requests
	SameSite http.SameSite
}

const (
	maxAge = 600 // 10 min
)

// DefaultCookieConfig configures short-lived temporary http.Cookie creation
var DefaultCookieConfig = CookieConfig{
	Name:     "datum-temporary-cookie",
	Path:     "/",
	Domain:   "",
	MaxAge:   maxAge,
	HTTPOnly: true,
	Secure:   true,
	SameSite: http.SameSiteStrictMode,
}

// DebugOnlyCookieConfig configures non-HTTPS cookies; for development only
var DebugOnlyCookieConfig = CookieConfig{
	Name:     "datum-debug-cookie",
	Path:     "/",
	MaxAge:   maxAge,
	HTTPOnly: true,
	Secure:   false,
	SameSite: http.SameSiteLaxMode,
}

// NewCookie returns a new http.Cookie with the given value and CookieConfig properties
func NewCookie(config CookieConfig, value string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     config.Name,
		Value:    value,
		Domain:   config.Domain,
		Path:     config.Path,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HTTPOnly,
		Secure:   config.Secure,
		SameSite: config.SameSite,
	}

	if expires, ok := expiresTime(config.MaxAge); ok {
		cookie.Expires = expires
	}

	return cookie
}

// expiresTime converts a maxAge time in seconds to a time.Time in the future
// ref http://golang.org/src/net/http/cookie.go?s=618:801#L23
func expiresTime(maxAge int) (time.Time, bool) {
	if maxAge > 0 {
		d := time.Duration(maxAge) * time.Second
		return time.Now().Add(d), true
	} else if maxAge < 0 {
		return time.Unix(1, 0), true
	}

	return time.Time{}, false
}

// GetCookie function retrieves a specific cookie from an HTTP request
func GetCookie(r *http.Request, cookieName string) (*http.Cookie, error) {
	return r.Cookie(cookieName)
}

// CookieExpired checks to see if a cookie is expired
func CookieExpired(cookie *http.Cookie) bool {
	if !cookie.Expires.IsZero() && cookie.Expires.Before(time.Now()) {
		return true
	}

	if cookie.MaxAge < 0 {
		return true
	}

	return false
}

// SetCookieB64 function sets a base64-encoded cookie with the given name and value in the HTTP response
func SetCookieB64(w http.ResponseWriter, body []byte, cookieName string, v CookieConfig) string {
	cookieValue := base64.StdEncoding.EncodeToString(body)
	// set the cookie
	SetCookie(w, cookieValue, cookieName, v)

	return cookieValue
}

// SetCookie function sets a cookie with the given value and name
func SetCookie(w http.ResponseWriter, value string, cookieName string, v CookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    value,
		Domain:   v.Domain,
		Path:     v.Path,
		MaxAge:   v.MaxAge,
		Secure:   v.Secure,
		SameSite: v.SameSite,
		HttpOnly: v.HTTPOnly,
	})
}

// RemoveCookie function removes a cookie from the HTTP response
func RemoveCookie(w http.ResponseWriter, cookieName string, v CookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Domain:   v.Domain,
		Path:     v.Path,
		MaxAge:   -1,
		Secure:   v.Secure,
		SameSite: v.SameSite,
		HttpOnly: v.HTTPOnly,
	})
}
