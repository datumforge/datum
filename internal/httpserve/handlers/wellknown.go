package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// JWKSWellKnownHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) JWKSWellKnownHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, h.JWTKeys)
}
