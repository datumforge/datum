package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerInviteHandler registers the invite handler
func registerInviteHandler(router *Router) (err error) {
	path := "/invite"
	method := http.MethodPost

	authMW := mw
	authMW = append(authMW, router.Handler.AuthMiddleware...)

	route := echo.Route{
		Name:   "OrganizationInviteAccept",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.OrganizationInviteAccept(c)
		},
	}.ForGroup(V1Version, authMW)

	inviteOperation := router.Handler.BindOrganizationInviteAccept()

	if err := router.AddRoute(path, method, inviteOperation, route); err != nil {
		return err
	}

	return nil
}
