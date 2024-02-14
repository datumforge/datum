//go:generate swagger generate spec
package route

import (
	"embed"
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/internal/httpserve/middleware/swagger"
	_ "github.com/datumforge/datum/openapi"
)

// registerJwksWellKnownHandler supplies the JWKS endpoint.
// This endpoint will contain the JWK used to verify all Datum JWTs
func registerJwksWellKnownHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/.well-known/jwks.json",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusOK, h.JWTKeys)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerOIDCHandler supplies the open-configuration endpoint
func registerOIDCHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/.well-known/openid-configuration",
		Handler: func(c echo.Context) error {
			return h.OpenIDConfiguration(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

//go:embed openapi.json
//go:embed robots.txt
//go:embed security.txt
//go:embed doc.json
var openapi embed.FS

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPISpecHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/api-docs",
		Handler: echo.StaticFileHandler("openapi.json", openapi),
	}.ForGroup(V1Version, mw))

	return
}

// registerSwaggerStatic embeds our generated open api specs and serves it behind /doc.json
func registerSwaggerStatic(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/restapi-docs",
		Handler: echo.StaticFileHandler("doc.json", openapi),
	}.ForGroup(V1Version, mw))

	return
}

// registerSecurityTxtHandler serves up the text output of datum's security.txt
func registerSecurityTxtHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/.well-known/security.txt",
		Handler: echo.StaticFileHandler("security.txt", openapi),
	}.ForGroup(unversioned, mw))

	return
}

// registerRobotsHandler serves up the robots.txt file via the RobotsHandler
func registerRobotsHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/robots.txt",
		Handler: echo.StaticFileHandler("robots.txt", openapi),
	}.ForGroup(unversioned, mw))

	return
}

func registerSwaggerMeowHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/swagger/*",
		Handler: swagger.WrapHandler,
	}.ForGroup(V1Version, mw))

	return
}
