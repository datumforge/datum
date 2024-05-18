package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// registerOAuthRegisterHandler registers the oauth register handler used by the UI to register
// users logging in with an oauth provider
func registerOAuthRegisterHandler(router *Router) (err error) {
	path := "/oauth/register"
	method := http.MethodPost

	route := echo.Route{
		Name:   "OAuthRegister",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.OauthRegister(c)
		},
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerUserInfoHandler registers the userinfo handler
func registerUserInfoHandler(router *Router) (err error) {
	authMW := mw
	authMW = append(authMW, router.Handler.AuthMiddleware...)

	path := "/oauth/userinfo"
	method := http.MethodGet

	route := echo.Route{
		Name:   "UserInfo",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

			return router.Handler.UserInfo(c)
		},
	}.ForGroup(unversioned, authMW)

	router.AddRoute(path, method, nil, route)

	return
}

// registerGithubLoginHandler registers the github login handler
func registerGithubLoginHandler(router *Router) (err error) {
	path := "/github/login"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GitHubLogin",
		Method:  method,
		Path:    path,
		Handler: githubLogin(router.Handler),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerGithubCallbackHandler registers the github callback handler
func registerGithubCallbackHandler(router *Router) (err error) {
	path := "/github/callback"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GitHubCallback",
		Method:  method,
		Path:    path,
		Handler: githubCallback(router.Handler),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerGoogleLoginHandler registers the google login handler
func registerGoogleLoginHandler(router *Router) (err error) {
	path := "/google/login"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GoogleLogin",
		Method:  method,
		Path:    path,
		Handler: googleLogin(router.Handler),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}

// registerGoogleCallbackHandler registers the google callback handler
func registerGoogleCallbackHandler(router *Router) (err error) {
	path := "/google/callback"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GoogleCallback",
		Method:  method,
		Path:    path,
		Handler: googleCallback(router.Handler),
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

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
