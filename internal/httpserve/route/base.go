package route

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/datumforge/datum/internal/httpserve/config"
)

// RegisterHandlers with the echo routers
func RegisterHandlers(router *echo.Echo) error { //nolint:unused
	// register handlers
	_, err := router.AddRoute(echo.Route{
		Method: http.MethodGet,
		Path:   "/livez",
		Handler: func(c echo.Context) error {
			return c.JSON(http.StatusOK, echo.Map{
				"status": "UP",
			})
		},
	})
	if err != nil {
		return err
	}

	// TODO: add readiness handlers

	_, err = router.AddRoute(echo.Route{
		Method:  http.MethodGet,
		Path:    "/metrics",
		Handler: echo.WrapHandler(promhttp.Handler()),
	})
	if err != nil {
		return err
	}

	return nil
}

func validateAuth(c echo.Context, cfgProvider *config.ConfigProviderWithRefresh) error { //nolint:unused
	cfg, err := cfgProvider.GetConfig()
	if err != nil {
		return err
	}

	isEnabled := cfg.Auth.Enabled
	if !isEnabled {
		return nil
	}

	//	sess, _ := session.Get("auth", c)
	//	if sess == nil {
	//		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	//	}
	//
	//	token := sess.Values["access-token"]
	//	if token == nil {
	//		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	//	}

	return nil
}
