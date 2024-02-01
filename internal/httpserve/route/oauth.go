package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerGithubLoginHandler
func registerGithubLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/github/login",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.GithubLoginHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerGithubCallbackHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/github/callback",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.GithubCallbackHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubLoginHandler
func registerGoogleLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/google/login",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.GoogleLoginHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerGoogleCallbackHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/google/callback",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.GoogleCallbackHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerDeviceAuthorizationHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/device/authorization",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.DeviceAuthorizationHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerOAuth2AuthorizeHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/oauth2/authorize",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.OAuth2AuthorizeHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerOAuth2TokenHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/oauth2/token",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.OAuth2TokenHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// registerGithubCallbackHandler
func registerOAuth2RedirectHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/oauth2/redirect",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.OAuth2RedirectHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}
