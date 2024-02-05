package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

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

	meow := echo.WrapHandler(login)

	return meow
}

// googleLogin wraps getloginhandlers
func googleLogin(h *handlers.Handler) echo.HandlerFunc {
	login, _ := h.GetGoogleLoginHandlers()

	meow := echo.WrapHandler(login)

	return meow
}

// githubCallback wraps getloginhandlers
func githubCallback(h *handlers.Handler) echo.HandlerFunc {
	_, callback := h.GetGitHubLoginHandlers()

	meow := echo.WrapHandler(callback)

	return meow
}

// googleCallback wraps getloginhandlers
func googleCallback(h *handlers.Handler) echo.HandlerFunc {
	_, callback := h.GetGoogleLoginHandlers()

	meow := echo.WrapHandler(callback)

	return meow
}
