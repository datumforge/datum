package route

import (
	"embed"
	"io"
	"net/http"
	"text/template"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
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

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

//go:embed openapi.yaml
//go:embed robots.txt
//go:embed security.txt
var openapi embed.FS

// registerOpenAPISpecHandler embeds our generated open api specs and serves it behind /api-docs
func registerOpenAPISpecHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/api-docs",
		Handler: echo.StaticFileHandler("openapi.yaml", openapi),
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

//go:embed assets/*
var assets embed.FS

// registerFaviconHandler serves up the favicon.ico
func registerFaviconHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/favicon.ico",
		Handler: echo.StaticFileHandler("assets/favicon.ico", assets),
	}.ForGroup(unversioned, mw))

	return
}
