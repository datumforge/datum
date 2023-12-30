package handlers

import (
	"encoding/json"
	"net/http"

	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshHandler allows users to refresh their access token using their refresh token.
func (h *Handler) RefreshHandler(ctx echo.Context) error {
	var r RefreshRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&r); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	if r.RefreshToken == "" {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(newMissingRequiredFieldError("refresh_token")))
	}

	// verify the refresh token
	claims, err := h.TM.Verify(r.RefreshToken)
	if err != nil {
		h.Logger.Errorw("error verifying token", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// check user in the database, sub == claims subject and ensure only one record is returned
	user, err := h.getUserBySub(ctx.Request().Context(), claims.Subject)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, ErrNoAuthUser)
		}

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// ensure the user is still active
	if user.Edges.Setting.Status != "ACTIVE" {
		return ctx.JSON(http.StatusNotFound, ErrNoAuthUser)
	}

	// UserID is not on the refresh token, so we need to set it now
	claims.UserID = user.ID

	accessToken, refreshToken, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, accessToken, refreshToken, h.CookieDomain); err != nil {
		h.Logger.Errorw("error setting cookies", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	return ctx.JSON(http.StatusOK, Response{Message: "success"})
}
