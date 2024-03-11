package ratelimit

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/kelseyhightower/envconfig"
)

// Config defines the configuration settings for the default rate limiter
type Config struct {
	Enabled    bool          `json:"enabled" koanf:"enabled" default:"false"`
	RateLimit  float64       `json:"limit" koanf:"limit" default:"10"`
	BurstLimit int           `json:"burst" koanf:"burst" default:"30"`
	ExpiresIn  time.Duration `json:"expires" koanf:"expires" default:"10m"`
}

// DefaultRateLimiter returns a middleware function for rate limiting requests
func DefaultRateLimiter() echo.MiddlewareFunc {
	conf := &Config{}

	err := envconfig.Process("datum", conf)
	if err != nil {
		panic(err)
	}

	return RateLimiterWithConfig(conf)
}

// RateLimiterWithConfig returns a middleware function for rate limiting requests with a config supplied
func RateLimiterWithConfig(conf *Config) echo.MiddlewareFunc {
	rateLimitConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      conf.RateLimit,
				Burst:     conf.BurstLimit,
				ExpiresIn: conf.ExpiresIn,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return &echo.HTTPError{
				Code:     middleware.ErrExtractorError.Code,
				Message:  middleware.ErrExtractorError.Message,
				Internal: err,
			}
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return &echo.HTTPError{
				Code:     middleware.ErrRateLimitExceeded.Code,
				Message:  "Too many requests!",
				Internal: err,
			}
		},
	}

	return middleware.RateLimiterWithConfig(rateLimitConfig)
}
