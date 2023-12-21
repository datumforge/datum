package redirect

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// Config contains the types used in executing redirects via the redirect middleware
type Config struct {
	// Skipper defines a function to skip middleware.
	Skipper   middleware.Skipper
	Redirects map[string]string
}

// DefaultConfig is the default configuration of the redirect middleware
var DefaultConfig = Config{
	Skipper:   middleware.DefaultSkipper,
	Redirects: map[string]string{},
}

// New creates a new middleware function with the default config
func New() echo.MiddlewareFunc {
	return NewWithConfig(DefaultConfig)
}

// NewWithConfig returns a new router middleware handler
func NewWithConfig(config Config) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()

			if target, ok := config.Redirects[req.URL.Path]; ok {
				return c.Redirect(http.StatusMovedPermanently, target)
			}

			return next(c)
		}
	}
}
