package ratelimit

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
)

// RateLimiter returns a middleware function for rate limiting requests, see https://echo.labstack.com/docs/middleware/rate-limiter
// TODO: https://github.com/datumforge/datum/issues/287
func RateLimiter() echo.MiddlewareFunc {
	rateLimitConfig := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      10,              // nolint: gomnd
				Burst:     30,              // nolint: gomnd
				ExpiresIn: 1 * time.Minute, // nolint: gomnd
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
	// TODO: make this configurable with inputs
	return middleware.RateLimiterWithConfig(rateLimitConfig)
}
