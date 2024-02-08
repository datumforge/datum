package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerUserInfoHandler registers the userinfo handler
func registerUserInfoHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	authMW := mw
	authMW = append(authMW, h.AuthMiddleware...)
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/oauth/userinfo",
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return h.UserInfo(c)
		},
	}.ForGroup(unversioned, authMW))

	return
}

// registerGithubLoginHandler registers the github login handler
func registerGithubLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/github/login",
		Handler: githubLogin(h),
	}.ForGroup(V1Version, mw))

	return
}

// registerGithubCallbackHandler registers the github callback handler
func registerGithubCallbackHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/github/callback",
		Handler: githubCallback(h),
	}.ForGroup(V1Version, mw))

	return
}

// registerGoogleLoginHandler registers the google login handler
func registerGoogleLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/google/login",
		Handler: googleLogin(h),
	}.ForGroup(V1Version, mw))

	return
}

// registerGoogleCallbackHandler registers the google callback handler
func registerGoogleCallbackHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/google/callback",
		Handler: googleCallback(h),
	}.ForGroup(V1Version, mw))

	return
}

// githubLogin wraps getloginhandlers
func githubLogin(h *handlers.Handler) echo.HandlerFunc {
	login, _ := h.GetGitHubLoginHandlers()

	return echo.WrapHandler(login)
}

// googleLogin wraps getloginhandlers
func googleLogin(h *handlers.Handler) echo.HandlerFunc {
	login, _ := h.GetGoogleLoginHandlers()

	return echo.WrapHandler(login)
}

// githubCallback wraps getloginhandlers
func githubCallback(h *handlers.Handler) echo.HandlerFunc {
	_, callback := h.GetGitHubLoginHandlers()

	return echo.WrapHandler(callback)
}

// googleCallback wraps getloginhandlers
func googleCallback(h *handlers.Handler) echo.HandlerFunc {
	_, callback := h.GetGoogleLoginHandlers()

	return echo.WrapHandler(callback)
}
