package session

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

type SessionConfig struct {
	Skipper        middleware.Skipper
	SessionManager *scs.SessionManager
}

var DefaultSessionConfig = SessionConfig{
	Skipper: middleware.DefaultSkipper,
}

// LoadAndSave is a middleware function that loads and saves session data using a
// provided session manager. It takes a `SessionManager` as input and returns a middleware function
// that can be used with an Echo framework application
func LoadAndSave(sessionManager *scs.SessionManager) echo.MiddlewareFunc {
	c := DefaultSessionConfig
	c.SessionManager = sessionManager

	return LoadAndSaveWithConfig(c)
}

// LoadAndSaveWithConfig is a middleware that loads and saves session data
// using a provided session manager configuration. It takes a `SessionConfig` struct as input, which
// contains the skipper function and the session manager
func LoadAndSaveWithConfig(config SessionConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultSessionConfig.Skipper
	}

	if config.SessionManager == nil {
		panic("Session middleware requires a session manager")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			ctx := c.Request().Context()

			var token string

			cookie, err := c.Cookie(config.SessionManager.Cookie.Name)

			if err == nil {
				token = cookie.Value
			}

			ctx, err = config.SessionManager.Load(ctx, token)
			if err != nil {
				return err
			}

			c.SetRequest(c.Request().WithContext(ctx))

			c.Response().Before(func() {
				if config.SessionManager.Status(ctx) != scs.Unmodified {
					responseCookie := &http.Cookie{
						Name:     config.SessionManager.Cookie.Name,
						Path:     config.SessionManager.Cookie.Path,
						Domain:   config.SessionManager.Cookie.Domain,
						Secure:   config.SessionManager.Cookie.Secure,
						HttpOnly: config.SessionManager.Cookie.HttpOnly,
						SameSite: config.SessionManager.Cookie.SameSite,
					}

					switch config.SessionManager.Status(ctx) {
					case scs.Modified:
						token, _, err := config.SessionManager.Commit(ctx)
						if err != nil {
							panic(err)
						}

						responseCookie.Value = token

					case scs.Destroyed:
						responseCookie.Expires = time.Unix(1, 0)
						responseCookie.MaxAge = -1
					}

					c.SetCookie(responseCookie)
					addHeaderIfMissing(c.Response(), "Cache-Control", `no-cache="Set-Cookie"`)
					addHeaderIfMissing(c.Response(), "Vary", "Cookie")
				}
			})

			return next(c)
		}
	}
}

// addHeaderIfMissing function is used to add a header to the HTTP response if it is not already
// present. It takes in the response writer (`http.ResponseWriter`), the header key, and the header
// value as parameters
func addHeaderIfMissing(w http.ResponseWriter, key, value string) {
	for _, h := range w.Header()[key] {
		if h == value {
			return
		}
	}

	w.Header().Add(key, value)
}
