package route

import (
	"net/http"

	"github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/httpserve/config"
)

func registerHandlers(router *echox.Echo) error { //nolint:unused
	// register handlers, example
	_, err := router.AddRoute(echox.Route{
		Method: http.MethodGet,
		Path:   "/api/hello",
		Handler: func(c echox.Context) error {
			obj := map[string]interface{}{"message": "Hello world! - New"}
			return c.JSON(http.StatusOK, obj)
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func validateAuth(c echox.Context, cfgProvider *config.ConfigProviderWithRefresh) error { //nolint:unused
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
