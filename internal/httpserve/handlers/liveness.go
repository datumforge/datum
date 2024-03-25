package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

func (h *Handler) LivenessCheck(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": "UP",
	})
}
