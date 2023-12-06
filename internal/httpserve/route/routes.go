package route

import (
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// RegisterRoutes with the echo routers
func RegisterRoutes(router *echo.Echo, checks *handlers.Checks) error {
	// register handlers
	if err := registerLivenessHandler(router); err != nil {
		return err
	}

	if err := registerReadinessHandler(router, checks); err != nil {
		return err
	}

	if err := registerMetricsHandler(router); err != nil {
		return err
	}

	if err := registerLoginHandler(router); err != nil {
		return err
	}

	return nil
}

// RegisterRoute with the echo server given a method, path, and handler definition
func (r *Route) RegisterRoute(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  r.Method,
		Path:    r.Path,
		Handler: r.Handler,
	})

	return
}
