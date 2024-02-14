package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

// @Summary Liveness Handler
// @Description Handles liveness check
// @ID liveness-handler
// @Produce json
// @Router /livez [get]

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

// @Summary Readiness Handler
// @Description Handles readiness check
// @ID readiness-handler
// @Produce json
// @Router /ready [get]
func registerReadinessHandler(router *echo.Echo, h *handlers.Handler) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/ready",
		Handler: func(c echo.Context) error {
			return h.ReadyChecks.ReadyHandler(c)
		},
	}.ForGroup(unversioned, mw))

	return
}

// @Summary Metrics Handler
// @Description Handles metrics request
// @ID metrics-handler
// @Produce json
// @Router /metrics [get]
func registerMetricsHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/metrics",
		Handler: echo.WrapHandler(promhttp.Handler()),
	}.ForGroup(unversioned, mw))

	return
}
