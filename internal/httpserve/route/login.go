package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// Login is oriented towards human users who use their email and password for
// authentication - see the handlers/login.go for more information
func registerLoginHandler(router *Router) (err error) {
	path := "/login"
	method := http.MethodPost

	route := echo.Route{
		Name:   "Login",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.LoginHandler(c)
		},
	}.ForGroup(V1Version, mw)

	loginOperation := router.Handler.BindLoginHandler()

	if err := router.AddRoute(path, method, loginOperation, route); err != nil {
		return err
	}

	return nil
}
