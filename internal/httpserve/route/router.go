package route

import (
	"time"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/middleware/ratelimit"
	"github.com/datumforge/datum/pkg/middleware/transaction"
)

var (
	mw     = []echo.MiddlewareFunc{middleware.Recover()}
	authMW = []echo.MiddlewareFunc{}

	restrictedRateLimit   = &ratelimit.Config{RateLimit: 10, BurstLimit: 10, ExpiresIn: 15 * time.Minute} //nolint:mnd
	restrictedEndpointsMW = []echo.MiddlewareFunc{}
)

// Router is a struct that holds the echo router, the OpenAPI schema, and the handler - it's a way to group these components together
type Router struct {
	Echo    *echo.Echo
	OAS     *openapi3.T
	Handler *handlers.Handler
}

// AddRoute is used to add a route to the echo router and OpenAPI schema at the same time ensuring consistency between the spec and the server
func (r *Router) AddRoute(pattern, method string, op *openapi3.Operation, route echo.Routable) error {
	_, err := r.Echo.AddRoute(route)
	if err != nil {
		return err
	}

	r.OAS.AddOperation(pattern, method, op)

	return nil
}

// AddRoute is used to add a route to the echo router and OpenAPI schema at the same time ensuring consistency between the spec and the server
func (r *Router) Addv1Route(pattern, method string, op *openapi3.Operation, route echo.Routable) error {
	grp := r.VersionOne()

	_, err := grp.AddRoute(route)
	if err != nil {
		return err
	}

	r.OAS.AddOperation(pattern, method, op)

	return nil
}

// AddRoute is used to add a route to the echo router and OpenAPI schema at the same time ensuring consistency between the spec and the server
func (r *Router) AddUnversionedRoute(pattern, method string, op *openapi3.Operation, route echo.Routable) error {
	grp := r.Base()

	_, err := grp.AddRoute(route)
	if err != nil {
		return err
	}

	r.OAS.AddOperation(pattern, method, op)

	return nil
}

// AddEchoOnlyRoute is used to add a route to the echo router without adding it to the OpenAPI schema
func (r *Router) AddEchoOnlyRoute(pattern, method string, route echo.Routable) error {
	grp := r.Base()

	_, err := grp.AddRoute(route)
	if err != nil {
		return err
	}

	return nil
}

// VersionOne returns a new echo group for version 1 of the API
func (r *Router) VersionOne() *echo.Group {
	return r.Echo.Group("v1")
}

// VersionTwo returns a new echo group for version 2 of the API - lets anticipate the future
func (r *Router) VersionTwo() *echo.Group {
	return r.Echo.Group("v2")
}

// Base returns the base echo group - no "version" prefix for the router group
func (r *Router) Base() *echo.Group {
	return r.Echo.Group("")
}

// RegisterRoutes with the echo routers - Router is defined within openapi.go
func RegisterRoutes(router *Router) error {
	// add transaction middleware
	transactionConfig := transaction.Client{
		EntDBClient: router.Handler.DBClient,
		Logger:      router.Handler.Logger,
	}

	mw = append(mw, transactionConfig.Middleware)

	// Middleware for restricted endpoints
	restrictedEndpointsMW = append(restrictedEndpointsMW, mw...)
	restrictedEndpointsMW = append(restrictedEndpointsMW, ratelimit.RateLimiterWithConfig(restrictedRateLimit)) // add restricted ratelimit middleware

	// Middleware for authenticated endpoints
	authMW = append(authMW, mw...)
	authMW = append(authMW, router.Handler.AuthMiddleware...)

	// routeHandlers that take the router and handler as input
	routeHandlers := []interface{}{
		registerReadinessHandler,
		registerForgotPasswordHandler,
		registerVerifyHandler,
		registerResetPasswordHandler,
		registerResendEmailHandler,
		registerRegisterHandler,
		registerVerifySubscribeHandler,
		registerRefreshHandler,
		registerJwksWellKnownHandler,
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
		registerLivenessHandler,
		registerMetricsHandler,
		registerSecurityTxtHandler,
		registerRobotsHandler,
		registerFaviconHandler,
		registerOpenAPIHandler,
		registerLoginHandler,
		registerAccountAccessHandler,
	}

	for _, route := range routeHandlers {
		if err := route.(func(*Router) error)(router); err != nil {
			return err
		}
	}

	return nil
}
