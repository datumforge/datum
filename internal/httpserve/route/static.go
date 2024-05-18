package route

import (
	"embed"
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerJwksWellKnownHandler supplies the JWKS endpoint.
// This endpoint will contain the JWK used to verify all Datum JWTs
func registerJwksWellKnownHandler(router *Router) (err error) {
	path := "/.well-known/jwks.json"
	method := http.MethodGet

	route := echo.Route{
		Name:   "JWKS",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusOK, router.Handler.JWTKeys)
		},
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerOIDCHandler supplies the open-configuration endpoint
func registerOIDCHandler(router *Router) (err error) {
	path := "/.well-known/openid-configuration"
	method := http.MethodGet

	route := echo.Route{
		Name:   "OpenIDConfiguration",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.OpenIDConfiguration(c)
		},
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

//go:embed openapi.yaml
var openapi embed.FS

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPISpecHandler(router *Router) (err error) {
	path := "/api-docs"
	method := http.MethodGet

	route := echo.Route{
		Name:    "OpenAPISpec",
		Method:  method,
		Path:    path,
		Handler: echo.StaticFileHandler("openapi.yaml", openapi),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPIHandler(router *Router) (err error) {
	path := "/openapi"
	method := http.MethodGet

	route := echo.Route{
		Method: method,
		Path:   path,
		Handler: echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, router.OAS)
		}),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

//go:embed security.txt
var securityTxt embed.FS

// registerSecurityTxtHandler serves up the text output of datum's security.txt
func registerSecurityTxtHandler(router *Router) (err error) {
	path := "/.well-known/security.txt"
	method := http.MethodGet

	route := echo.Route{
		Name:    "SecurityTxt",
		Method:  method,
		Path:    path,
		Handler: echo.StaticFileHandler("security.txt", securityTxt),
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

//go:embed robots.txt
var robotsTxt embed.FS

// registerRobotsHandler serves up the robots.txt file via the RobotsHandler
func registerRobotsHandler(router *Router) (err error) {
	path := "/robots.txt"
	method := http.MethodGet

	route := echo.Route{
		Name:    "Robots",
		Method:  method,
		Path:    path,
		Handler: echo.StaticFileHandler("robots.txt", robotsTxt),
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

//go:embed assets/*
var assets embed.FS

// registerFaviconHandler serves up the favicon.ico
func registerFaviconHandler(router *Router) (err error) {
	path := "/favicon.ico"
	method := http.MethodGet

	route := echo.Route{
		Name:    "Favicon",
		Method:  method,
		Path:    path,
		Handler: echo.StaticFileHandler("assets/favicon.ico", assets),
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}
