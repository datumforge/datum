package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// registerEventPublisher registers the event publisher endpoint
func registerEventPublisher(router *Router) (err error) {
	path := "/event/publish"
	method := http.MethodPost

	route := echo.Route{
		Name:   "EventPublisher",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.EventPublisher(c)
		},
	}.ForGroup(V1Version, mw)

	router.AddRoute(path, method, nil, route)

	return
}
