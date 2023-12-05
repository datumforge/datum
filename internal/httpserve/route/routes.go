package route

import (
	echo "github.com/datumforge/echox"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

// RegisterBaseRoutes with the echo routers
func RegisterBaseRoutes(router *echo.Echo) error {
	// register handlers
	if err := registerLivenessHandler(router); err != nil {
		return err
	}

	if err := registerReadinessHandler(router); err != nil {
		return err
	}

	if err := registerMetricsHandler(router); err != nil {
		return err
	}

	return nil
}

func (r *Route) RegisterRoute(router *echo.Echo) (err error) {
	_, err = router.AddRoute(echo.Route{
		Method:  r.Method,
		Path:    r.Path,
		Handler: r.Handler,
	})

	return
}
