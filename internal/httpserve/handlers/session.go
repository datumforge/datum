package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
)

// SessionPutHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) SessionPutHandler(ctx echo.Context) error {
	// test storing a new key and value into the session data
	h.SM.Put(ctx.Request().Context(), "message", "matt is the best!")

	return ctx.JSON(http.StatusOK, http.StatusOK)
}

// SessionGetHandler provides the JWK used to verify all Datum-issued JWTs
func (h *Handler) SessionGetHandler(ctx echo.Context) error {
	// test storing a new key and value into the session data
	msg := h.SM.GetString(ctx.Request().Context(), "message")

	return ctx.JSON(http.StatusOK, msg)
}
