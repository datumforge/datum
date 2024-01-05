package route

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/ratelimit"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
)

const (
	V1Version   = "v1"
	unversioned = ""
)

var (
	mw = []echo.MiddlewareFunc{middleware.Recover()}

	restrictedRateLimit = &ratelimit.Config{
		RateLimit:  1,
		BurstLimit: 1,
		ExpiresIn:  15 * time.Minute, //nolint:gomnd
	}
	restrictedEndpointsMW = []echo.MiddlewareFunc{}
)

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc

	Name string
}

// RegisterRoutes with the echo routers
func RegisterRoutes(router *echo.Echo, h *handlers.Handler) error {
	// add transaction middleware
	transactionConfig := transaction.Client{
		EntDBClient: h.DBClient,
		Logger:      h.Logger,
	}

	mw = append(mw, transactionConfig.Middleware)

	// Middleware for restricted endpoints
	restrictedEndpointsMW = append(restrictedEndpointsMW, mw...)
	restrictedEndpointsMW = append(restrictedEndpointsMW, ratelimit.RateLimiterWithConfig(restrictedRateLimit)) // add restricted ratelimit middleware

	// register handlers
	if err := registerLivenessHandler(router); err != nil {
		return err
	}

	if err := registerReadinessHandler(router, h); err != nil {
		return err
	}

	if err := registerMetricsHandler(router); err != nil {
		return err
	}

	if err := registerLoginHandler(router, h); err != nil {
		return err
	}

	if err := registerForgotPasswordHandler(router, h); err != nil {
		return err
	}

	if err := registerVerifyHandler(router, h); err != nil {
		return err
	}

	if err := registerResetPasswordHandler(router); err != nil {
		return err
	}

	if err := registerResendEmailHandler(router, h); err != nil {
		return err
	}

	if err := registerRegisterHandler(router, h); err != nil {
		return err
	}

	if err := registerRefreshHandler(router, h); err != nil {
		return err
	}

	if err := registerAuthenticateHandler(router); err != nil {
		return err
	}

	if err := registerJwksWellKnownHandler(router, h); err != nil {
		return err
	}

	return nil
}

// RegisterRoute with the echo server given a method, path, and handler definition
func (r *Route) RegisterRoute(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:      r.Method,
		Path:        r.Path,
		Handler:     r.Handler,
		Middlewares: r.Middlewares,

		Name: r.Name,
	})

	return
}
