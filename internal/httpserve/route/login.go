package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// TODO: implement login handler
func registerLoginHandler(router *echo.Echo) (err error) { //nolint:unused
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/login",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusNotImplemented, echo.Map{
				"error": "Not implemented",
			})
		},
	})

	return
}
