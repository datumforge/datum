package sessions

import (
	"net/http"
	"time"
)

const (
	defaultMaxAge = 3600 * 24 * 7 // 1 week
)

// DefaultCookieConfig configures http.Cookie creation for production.
var DefaultCookieConfig = &CookieConfig{
	Path:     "/",
	MaxAge:   defaultMaxAge,
	HTTPOnly: true,
	Secure:   true,
	SameSite: http.SameSiteLaxMode,
}

// DebugCookieConfig configures http.Cookie creation for debugging. It
// does NOT require cookies be sent over HTTPS so it should only be used
// in development. Prefer DefaultCookieConfig.
var DebugCookieConfig = &CookieConfig{
	Path:     "/",
	MaxAge:   defaultMaxAge,
	HTTPOnly: true,
	Secure:   false,
	SameSite: http.SameSiteLaxMode,
}

// CookieConfig configures http.Cookie creation.
type CookieConfig struct {
	Name string
	// Cookie domain/path scope (leave zeroed for requested resource scope)
	// Defaults to the domain name of the responding server when unset
	Domain string
	// Defaults to the path of the responding URL when unset
	Path string
	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'.
	// MaxAge>0 means Max-Age attribute present and given in seconds.
	MaxAge int
	// cookie may only be transferred over HTTPS. Recommend true.
	Secure bool
	// browser should prohibit non-HTTP (i.e. javascript) cookie access. Recommend true
	HTTPOnly bool
	// prohibit sending in cross-site requests with SameSiteLaxMode or SameSiteStrictMode
	SameSite http.SameSite
}

// NewCookie returns a new http.Cookie with the given name, value, and
// properties from config.
func NewCookie(name, value string, config *CookieConfig) *http.Cookie {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     config.Path,
		Domain:   config.Domain,
		MaxAge:   config.MaxAge,
		HttpOnly: config.HTTPOnly,
		Secure:   config.Secure,
		SameSite: config.SameSite,
	}
	// IE <9 does not understand MaxAge, set Expires based on MaxAge
	if expires, present := cookieExpires(config.MaxAge); present {
		cookie.Expires = expires
	}

	return cookie
}

// cookieExpires takes the MaxAge number of seconds a Cookie should be valid
// and returns the Expires time.Time and whether the attribtue should be set.
// http://golang.org/src/net/http/cookie.go?s=618:801#L23
func cookieExpires(maxAge int) (time.Time, bool) {
	if maxAge > 0 {
		d := time.Duration(maxAge) * time.Second
		return time.Now().Add(d), true
	} else if maxAge < 0 {
		return time.Unix(1, 0), true // first second of the epoch
	}

	return time.Time{}, false
}
