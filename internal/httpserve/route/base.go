package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/datumforge/datum/internal/httpserve/handlers"
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
	})

	return
}

func registerReadinessHandler(router *echo.Echo, readinessChecks *handlers.Checks) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/ready",
		Handler: func(c echo.Context) error {
			return readinessChecks.ReadyHandler(c)
		},
	})

	return
}

func registerMetricsHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/metrics",
		Handler: echo.WrapHandler(promhttp.Handler()),
	})

	return
}
