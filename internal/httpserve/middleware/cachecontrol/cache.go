package cachecontrol

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

var epoch = time.Unix(0, 0).Format(time.RFC1123)

// Config is the config values for the cache-control middleware
type Config struct {
	// Skipper defines a function to skip middleware
	Skipper middleware.Skipper

	// noCacheHeaders is the header <-> match map pair to match in http for entity headers to remove
	noCacheHeaders map[string]string

	// etagHeaders is the string of entity headers to remove
	etagHeaders []string
}

// DefaultConfig is the default configuration of the middleware
var DefaultConfig = Config{
	Skipper: middleware.DefaultSkipper,
	noCacheHeaders: map[string]string{
		"Expires":         epoch,
		"Cache-Control":   "no-cache, private, max-age=0",
		"Pragma":          "no-cache",
		"X-Accel-Expires": "0",
	},
	etagHeaders: []string{
		"ETag",
		"If-Modified-Since",
		"If-Match",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
	},
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
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			// Delete any ETag headers that may have been set
			for _, v := range DefaultConfig.etagHeaders {
				if req.Header.Get(v) != "" {
					req.Header.Del(v)
				}
			}

			// Set our NoCache headers
			res := c.Response()
			for k, v := range DefaultConfig.noCacheHeaders {
				res.Header().Set(k, v)
			}

			return next(c)
		}
	}
}
