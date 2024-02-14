//go:generate swagger generate spec
package route

import (
	"embed"
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerSecurityTxtHandler serves up the text output of datum's security.txt
func registerSecurityTxtHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/security.txt",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
			return h.SecurityHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerRobotsHandler serves up the robots.txt file via the RobotsHandler
func registerRobotsHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/robots.txt",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
			return h.RobotsHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerJwksWellKnownHandler supplies the JWKS endpoint.
// This endpoint will contain the JWK used to verify all Datum JWTs
func registerJwksWellKnownHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/.well-known/jwks.json",
		Handler: func(c echo.Context) error {
			return h.JWKSWellKnownHandler(c)
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
var embeddedFiles embed.FS

//go:embed doc.json
var openapifiles embed.FS

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPISpecHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/api-docs",
		Handler: echo.StaticFileHandler("openapi.json", embeddedFiles),
	}.ForGroup(V1Version, mw))

	return
}

// registerSwaggerStatic embeds our generated open api specs and serves it behind /doc.json
func registerSwaggerStatic(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/doc.json",
		Handler: echo.StaticFileHandler("doc.json", openapifiles),
	}.ForGroup(V1Version, mw))

	return
}
