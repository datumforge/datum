package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerWebauthnRegistrationHandler registers the webauthn registration handler
func registerWebauthnRegistrationHandler(router *Router) (err error) {
	path := "/registration/options"
	method := http.MethodPost

	route := echo.Route{
		Name:   "WebauthnRegistration",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.BeginWebauthnRegistration(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerWebauthnVerificationsHandler registers the webauthn registration verification handler
func registerWebauthnVerificationsHandler(router *Router) (err error) {
	path := "/registration/verification"
	method := http.MethodPost

	route := echo.Route{
		Name:   "WebauthnRegistrationVerification",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.FinishWebauthnRegistration(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerWebauthnAuthenticationHandler registers the webauthn authentication handler
func registerWebauthnAuthenticationHandler(router *Router) (err error) {
	path := "/authentication/options"
	method := http.MethodPost

	route := echo.Route{
		Name:   "WebauthnAuthentication",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.BeginWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerWebauthnAuthVerificationHandler registers the webauthn authentication verification handler
func registerWebauthnAuthVerificationHandler(router *Router) (err error) {
	path := "/authentication/verification"
	method := http.MethodPost

	route := echo.Route{
		Name:   "WebauthnAuthenticationVerification",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.FinishWebauthnLogin(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}
