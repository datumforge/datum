package route

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/ratelimit"
	"github.com/datumforge/datum/pkg/middleware/transaction"
	oas "github.com/datumforge/datum/pkg/oasrouter"
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

type OASRouter = oas.Router[echo.HandlerFunc, echo.RouteInfo]

// RegisterRoutes with the echo routers
func RegisterRoutes(router *echo.Echo, h *handlers.Handler, oasrouter *oas.Router[echo.HandlerFunc, echo.RouteInfo]) error {
	// add transaction middleware0
	transactionConfig := transaction.Client{
		EntDBClient: h.DBClient,
		Logger:      h.Logger,
	}

	mw = append(mw, transactionConfig.Middleware)

	// Middleware for restricted endpoints
	restrictedEndpointsMW = append(restrictedEndpointsMW, mw...)
	restrictedEndpointsMW = append(restrictedEndpointsMW, ratelimit.RateLimiterWithConfig(restrictedRateLimit)) // add restricted ratelimit middleware

	routeOASHandlers := []interface{}{
		registerOASHandler,
	}

	for _, route := range routeOASHandlers {
		if err := route.(func(*echo.Echo, *OASRouter) error)(router, oasrouter); err != nil {
			return err
		}
	}

	// routeHandlers that take the router and handler as input
	routeHandlers := []interface{}{
		registerReadinessHandler,
		registerLoginHandler,
		registerForgotPasswordHandler,
		registerVerifyHandler,
		registerResetPasswordHandler,
		registerResendEmailHandler,
		registerRegisterHandler,
		registerSubscribeHandler,
		registerVerifySubscribeHandler,
		registerUnsubscribeHandler,
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
		registerSwaggerStatic,
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
