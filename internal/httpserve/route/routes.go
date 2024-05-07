package route

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/ratelimit"
	"github.com/datumforge/datum/pkg/middleware/transaction"
)

const (
	V1Version   = "v1"
	unversioned = ""
)

var (
	mw = []echo.MiddlewareFunc{middleware.Recover()}

	restrictedRateLimit   = &ratelimit.Config{RateLimit: 10, BurstLimit: 10, ExpiresIn: 15 * time.Minute} //nolint:gomnd
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

	// routeHandlers that take the router and handler as input
	routeHandlers := []interface{}{
		registerReadinessHandler,
		registerLoginHandler,
		registerForgotPasswordHandler,
		registerVerifyHandler,
		registerResetPasswordHandler,
		registerResendEmailHandler,
		registerRegisterHandler,
		registerVerifySubscribeHandler,
		registerRefreshHandler,
		registerAuthenticateHandler,
		registerJwksWellKnownHandler,
		registerOIDCHandler,
		registerInviteHandler,
		registerGithubLoginHandler,
		registerGithubCallbackHandler,
		registerGoogleLoginHandler,
		registerGoogleCallbackHandler,
		registerWebauthnRegistrationHandler,
		registerWebauthnVerificationsHandler,
		registerWebauthnAuthenticationHandler,
		registerWebauthnAuthVerificationHandler,
		registerUserInfoHandler,
		registerOAuthRegisterHandler,
		registerSwitchRoute,
		registerEventPublisher,
	}

	for _, route := range routeHandlers {
		if err := route.(func(*echo.Echo, *handlers.Handler) error)(router, h); err != nil {
			return err
		}
	}

	// register additional handlers that only require router input
	additionalHandlers := []interface{}{
		registerLivenessHandler,
		registerOpenAPISpecHandler,
		registerMetricsHandler,
		registerSecurityTxtHandler,
		registerRobotsHandler,
	}

	for _, route := range additionalHandlers {
		if err := route.(func(*echo.Echo) error)(router); err != nil {
			return err
		}
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
