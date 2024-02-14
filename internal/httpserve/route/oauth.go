package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// @Summary OAuth Register
// @Description Register a new user using OAuth
// @Tags OAuth
// @Accept json
// @Produce json
// @Param oauth-register body OAuthRegisterRequest true "OAuth register request"
// @Success 200 {object} OAuthRegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /oauth/register [post]

func registerOAuthReigsterandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodPost,
		Path:   "/oauth/register",
		Handler: func(c echo.Context) error {
			return h.OauthRegister(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

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

// @Summary GitHub Login
// @Description Initiate GitHub login
// @Tags OAuth
// @Accept json
// @Produce json
// @Security OAuth2
// @Success 302 {string} string "Redirects to GitHub login page"
// @Failure 401 {object} ErrorResponse
// @Router /github/login [get]
func registerGithubLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/github/login",
		Handler: githubLogin(h),
	}.ForGroup(V1Version, mw))

	return
}

// @Summary GitHub Callback
// @Description Callback URL for GitHub OAuth flow
// @Tags OAuth
// @Accept json
// @Produce json
// @Security OAuth2
// @Success 302 {string} string "Redirects to the application after successful GitHub login"
// @Failure 401 {object} ErrorResponse
// @Router /github/callback [get]
func registerGithubCallbackHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/github/callback",
		Handler: githubCallback(h),
	}.ForGroup(V1Version, mw))

	return
}

// @Summary Google Login
// @Description Initiate Google login
// @Tags OAuth
// @Accept json
// @Produce json
// @Security OAuth2
// @Success 302 {string} string "Redirects to Google login page"
// @Failure 401 {object} ErrorResponse
// @Router /google/login [get]
func registerGoogleLoginHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/google/login",
		Handler: googleLogin(h),
	}.ForGroup(V1Version, mw))

	return
}

// @Summary Google Callback
// @Description Callback URL for Google OAuth flow
// @Tags OAuth
// @Accept json
// @Produce json
// @Security OAuth2
// @Success 302 {string} string "Redirects to the application after successful Google login"
// @Failure 401 {object} ErrorResponse
// @Router /google/callback [get]
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
