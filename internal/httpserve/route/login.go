package route

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/oas"
)

// Login is oriented towards human users who use their email and password for
// authentication (whereas authenticate is used for machine access using API keys).
// Login verifies the password submitted for the user is correct by looking up the user
// by email and using the argon2 derived key verification process to confirm the
// password matches. Upon authentication an access token and a refresh token with the
// authorized claims of the user are returned. The user can use the
// access token to authenticate to Datum systems. The access token has an expiration and
// the refresh token can be used with the refresh endpoint to get a new access token
// without the user having to log in again. The refresh token overlaps with the access
// token to provide a seamless authentication experience and the user can refresh their
// access token so long as the refresh token is valid.
func registerLoginHandler(router *echo.Echo, h *handlers.Handler, oasRouter *OASRouter) (err error) {
	e := echo.Route{
		Method: http.MethodPost,
		Path:   "/login",
		Handler: func(c echo.Context) error {
			return h.LoginHandler(c)
		},
	}.ForGroup(V1Version, mw)

	opr := oas.NewOperation()
	opr.Description = "User login and authentication via email and password. The user can optionally supply an OrgID to log into a specific organization that they belong to. An _access token_ and a _refresh token_ are returned that contain the authorization claims of the user. The access token should be added to the Authorize header as a Bearer token for all requests that require authentication. The refresh token can be used to obtain a new access token after it expires (extending the user's session) without requiring the user to supply their username and password again.The access token contains claims that help identify the user (e.g. name, email, picture) as well as the ID of the organization the user is currently logged into. A user can only be logged into one organization at a time. Additionally the claims contain the permissions the user has defined by the user's role. The subject of the claims is the ID of the user."

	_, err = oasRouter.AddRoute(e.ToRoute().Method, e.ToRoute().Path, e.ToRoute().Handler, oas.Definitions{
		RequestBody: &oas.ContentValue{
			Content: oas.Content{
				"application/json": {Value: handlers.LoginRequest{}},
			},
		},
		Responses: map[int]oas.ContentValue{
			200: {
				Content: oas.Content{
					"application/json": {Value: handlers.LoginReply{}},
				},
			},
		},
	})

	return
}
