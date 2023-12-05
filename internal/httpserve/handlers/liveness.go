package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// livenessCheckHandler ensures that the server is up and responding
func (h *Handler) livenessCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status": "UP",
	})
}
