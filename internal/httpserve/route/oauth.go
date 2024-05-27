package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
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

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
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

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// registerGithubLoginHandler registers the github login handler
func registerGithubLoginHandler(router *Router) (err error) {
	path := "/github/login"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GitHubLogin",
		Method:  method,
		Path:    path,
		Handler: githubLogin(router),
	}.ForGroup(V1Version, mw)

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// registerGithubCallbackHandler registers the github callback handler
func registerGithubCallbackHandler(router *Router) (err error) {
	path := "/github/callback"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GitHubCallback",
		Method:  method,
		Path:    path,
		Handler: githubCallback(router),
	}.ForGroup(V1Version, mw)

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// registerGoogleLoginHandler registers the google login handler
func registerGoogleLoginHandler(router *Router) (err error) {
	path := "/google/login"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GoogleLogin",
		Method:  method,
		Path:    path,
		Handler: googleLogin(router),
	}.ForGroup(V1Version, mw)

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// registerGoogleCallbackHandler registers the google callback handler
func registerGoogleCallbackHandler(router *Router) (err error) {
	path := "/google/callback"
	method := http.MethodGet

	route := echo.Route{
		Name:    "GoogleCallback",
		Method:  method,
		Path:    path,
		Handler: googleCallback(router),
	}.ForGroup(V1Version, mw)

	if err := router.AddRoute(path, method, nil, route); err != nil {
		return err
	}

	return nil
}

// githubLogin wraps getloginhandlers
func githubLogin(h *Router) echo.HandlerFunc {
	login, _ := h.Handler.GetGitHubLoginHandlers()

	return echo.WrapHandler(login)
}

// googleLogin wraps getloginhandlers
func googleLogin(h *Router) echo.HandlerFunc {
	login, _ := h.Handler.GetGoogleLoginHandlers()

	return echo.WrapHandler(login)
}

// githubCallback wraps getloginhandlers
func githubCallback(h *Router) echo.HandlerFunc {
	_, callback := h.Handler.GetGitHubLoginHandlers()

	return echo.WrapHandler(callback)
}

// googleCallback wraps getloginhandlers
func googleCallback(h *Router) echo.HandlerFunc {
	_, callback := h.Handler.GetGoogleLoginHandlers()

	return echo.WrapHandler(callback)
}
