package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/datumforge/datum/internal/httpserve/handlers"
	"github.com/datumforge/datum/pkg/oas"
)

func registerLivenessHandler(router *echo.Echo, h *handlers.Handler, oasRouter *OASRouter) (err error) {
	e := echo.Route{
		Method: http.MethodGet,
		Path:   "/livez",
		Handler: func(c echo.Context) error {
			return h.LivenessCheck(c)
		},
	}.ForGroup(unversioned, mw)

	_, err = oasRouter.AddRoute(e.ToRoute().Method, e.ToRoute().Path, e.ToRoute().Handler, oas.Definitions{})
	if err != nil {
		return err
	}

	return
}

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

func registerMetricsHandler(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/metrics",
		Handler: echo.WrapHandler(promhttp.Handler()),
	}.ForGroup(unversioned, mw))

	return
}
