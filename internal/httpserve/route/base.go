package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/echox/middleware"
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
		Middlewares: []echo.MiddlewareFunc{middleware.Recover()},
	})

	return
}

func registerReadinessHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/ready",
		Handler: func(c echo.Context) error {
			return h.ReadyChecks.ReadyHandler(c)
		},
		Middlewares: []echo.MiddlewareFunc{middleware.Recover()},
	})

	return
}

func registerMetricsHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:      http.MethodGet,
		Path:        "/metrics",
		Handler:     echo.WrapHandler(promhttp.Handler()),
		Middlewares: []echo.MiddlewareFunc{middleware.Recover()},
	})

	return
}
