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

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPIHandler(router *Router) (err error) {
	path := "/api-docs"
	method := http.MethodGet

	route := echo.Route{
		Method: method,
		Path:   path,
		Handler: echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, router.OAS)
		}),
	}.ForGroup(V1Version, mw)

	if err := router.AddEchoOnlyRoute(path, method, route); err != nil {
		return err
	}

	return nil
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

	if err := router.AddEchoOnlyRoute(path, method, route); err != nil {
		return err
	}

	return nil
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

	if err := router.AddEchoOnlyRoute(path, method, route); err != nil {
		return err
	}

	return nil
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

	if err := router.AddEchoOnlyRoute(path, method, route); err != nil {
		return err
	}

	return nil
}
