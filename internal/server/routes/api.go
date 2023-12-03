package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/datumforge/datum/internal/server/config"
)

// SetAPIRoutes sets user routes
func SetAPIRoutes(e *echo.Echo, cfgProvider *config.ConfigProviderWithRefresh) error {
	api := e.Group("/user")
	api.GET("/v1/settings", getSettings(cfgProvider))
	api.Match([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}, "/*", temporalAPIHandler(cfgProvider))
	return nil
}

func temporalAPIHandler(cfgProvider *config.ConfigProviderWithRefresh) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := validateAuth(c, cfgProvider)
		if err != nil {
			return err
		}

		cfg, err := cfgProvider.GetConfig()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Sprintln(cfg)

		return nil
	}
}

func getCurrentUser(c echo.Context) error {
	sess, _ := session.Get("auth", c)
	email := sess.Values["email"]
	name := sess.Values["name"]
	picture := sess.Values["picture"]

	if email == nil {
		return c.JSON(http.StatusOK, nil)
	}

	user := struct {
		Email   string
		Name    string
		Picture string
	}{email.(string), name.(string), picture.(string)}

	return c.JSON(http.StatusOK, user)
}

// this is just a dummy function stub to show the function call from the api group
func getSettings(cfgProvier *config.ConfigProviderWithRefresh) func(echo.Context) error {
	return func(c echo.Context) error {
		cfg, err := cfgProvier.GetConfig()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Sprintln(cfg)

		return err
	}
}

func validateAuth(c echo.Context, cfgProvider *config.ConfigProviderWithRefresh) error {
	cfg, err := cfgProvider.GetConfig()
	if err != nil {
		return err
	}
	//TODO actually put stuff inside cfg.Auth.enabled
	fmt.Sprintln(cfg)

	isEnabled := true
	if !isEnabled {
		return nil
	}

	sess, _ := session.Get("auth", c)
	if sess == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	token := sess.Values["access-token"]
	if token == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	return nil
}
