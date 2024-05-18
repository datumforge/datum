package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func registerLivenessHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/livez",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"status": "UP",
			})
		},
	}.ForGroup(unversioned, mw))

	return
}

func registerReadinessHandler(router *Router) (err error) {
	path := "/ready"
	method := http.MethodGet

	route := echo.Route{
		Name:   "Ready",
		Method: method,
		Path:   path,
		Handler: func(c echo.Context) error {
			return router.Handler.ReadyChecks.ReadyHandler(c)
		},
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}

func registerMetricsHandler(router *Router) (err error) {
	path := "/metrics"
	method := http.MethodGet

	route := echo.Route{
		Name:    "Metrics",
		Method:  method,
		Path:    path,
		Handler: echo.WrapHandler(promhttp.Handler()),
	}.ForGroup(unversioned, mw)

	router.AddRoute(path, method, nil, route)

	return
}
