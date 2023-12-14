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

// TODO: implement resendemail handler
func registerResendEmailHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/resend",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
