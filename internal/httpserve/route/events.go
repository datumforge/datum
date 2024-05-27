package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerEventPublisher registers the event publisher endpoint
func registerEventPublisher(router *Router) (err error) {
	path := "/event/publish"
	method := http.MethodPost
	name := "EventPublisher"

	route := echo.Route{
		Name:   name,
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.EventPublisher(c)
		},
	}.ForGroup(V1Version, mw)

	eventOperation := router.Handler.BindEventPublisher()

	if err := router.AddRoute(path, method, eventOperation, route); err != nil {
		return err
	}

	return nil
}
