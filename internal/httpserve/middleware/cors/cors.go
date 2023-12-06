package cors

import (
	"fmt"
	"strings"
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// Config holds the cors configuration settings
type Config struct {
	// Skipper defines a function to skip middleware.
	Skipper  middleware.Skipper
	Prefixes map[string][]string
}

// DefaultConfig creates a default config
var DefaultConfig = Config{
	Skipper:  middleware.DefaultSkipper,
	Prefixes: nil,
}

// New creates a new middleware function with the default config
func New() echo.MiddlewareFunc {
	mw, _ := NewWithConfig(DefaultConfig)

	return mw
}

// NewWithConfig creates a new middleware function with the provided config
func NewWithConfig(config Config) (echo.MiddlewareFunc, error) {
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}

	prefixes := make(map[string]echo.MiddlewareFunc)

	for prefix, origins := range config.Prefixes {
		if err := Validate(origins); err != nil {
			return nil, fmt.Errorf("CORS config for prefix %s is invalid: %w", prefix, err)
		}

		conf := middleware.CORSConfig{
			AllowOrigins:     origins,
			AllowMethods:     []string{"GET", "HEAD", "PUT", "POST", "DELETE", "PATCH"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           int((24 * time.Hour).Seconds()), //nolint:gomnd
		}

		prefixes[prefix] = middleware.CORSWithConfig(conf)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if origin := c.Request().Header.Get("Origin"); len(origin) == 0 {
				c.Request().Header.Set("Origin", "*")
			}

			path := c.Request().URL.Path

			var (
				middlewareFunc echo.MiddlewareFunc
				maxPrefixLen   int
			)

			for prefix, h := range prefixes {
				if strings.HasPrefix(path, prefix) {
					if len(prefix) > maxPrefixLen {
						maxPrefixLen = len(prefix)
						middlewareFunc = h
					}
				}
			}

			if middlewareFunc != nil {
				handler := middlewareFunc(next)
				return handler(c)
			}

			return next(c)
		}
	}, nil
}

// DefaultSchemas is a list of default allowed schemas for CORS origins
var DefaultSchemas = []string{
	"http://",
	"https://",
}

// Validate checks a list of origins to see if they comply with the allowed origins
func Validate(origins []string) error {
	for _, origin := range origins {
		if !strings.Contains(origin, "*") && !validateAllowedSchemas(origin) {
			allowed := fmt.Sprintf(" origins must contain '*' or include %s", strings.Join(getAllowedSchemas(), ", or "))

			return newValidationError("bad origin", allowed)
		}
	}

	return nil
}

func validateAllowedSchemas(origin string) bool {
	allowedSchemas := getAllowedSchemas()

	for _, schema := range allowedSchemas {
		if strings.HasPrefix(origin, schema) {
			return true
		}
	}

	return false
}

func getAllowedSchemas() []string {
	allowedSchemas := DefaultSchemas

	return allowedSchemas
}
