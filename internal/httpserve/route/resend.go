package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// ResendEmail accepts an email address via a POST request and always returns a 204
// response, no matter the input or result of the processing. This is to ensure that
// no secure information is leaked from this unauthenticated endpoint. If the email
// address belongs to a user who has not been verified, another verification email is
// sent. If the post request contains an orgID and the user is invited to that
// organization but hasn't accepted the invite, then the invite is resent.
func registerResendEmailHandler(router *Router) (err error) {
	path := "/resend"
	method := http.MethodPost

	route := echo.Route{
		Name:   "ResendEmail",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.ResendEmail(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}
