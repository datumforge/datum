package route

import (
	"net/http"

	"github.com/datumforge/datum/internal/httpserve/config"
	"github.com/datumforge/echox"
)

func registerHandlers(router *echox.Echo) error {
	router.AddRoute(echox.Route{
		Method: http.MethodGet,
		Path:   "/api/hello",
		Handler: func(c echox.Context) error {
			obj := map[string]interface{}{"message": "Hello world! - New"}
			return c.JSON(http.StatusOK, obj)
		},
	})

	return nil
}

func validateAuth(c echox.Context, cfgProvider *config.ConfigProviderWithRefresh) error {
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
