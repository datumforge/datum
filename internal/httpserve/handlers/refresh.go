package handlers

import (
	"encoding/json"
	"net/http"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"

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
		auth.Unauthorized(ctx) //nolint:errcheck
		return ErrBadRequest
	}

	if r.RefreshToken == "" {
		auth.Unauthorized(ctx) //nolint:errcheck
		return ErrBadRequest
	}

	// verify the refresh token
	claims, err := h.TM.Verify(r.RefreshToken)
	if err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return err
	}

	// check user in the database, sub == claims subject and ensure only one record is returned
	user, err := h.DBClient.User.Query().WithSetting().Where(func(s *sql.Selector) {
		s.Where(sql.EQ("sub", claims.Subject))
	}).Only(ctx.Request().Context())
	if err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return auth.ErrNoAuthUser
	}

	// ensure the user is still active
	if user.Edges.Setting.Status != "ACTIVE" {
		auth.Unauthorized(ctx) //nolint:errcheck
		return auth.ErrNoAuthUser
	}

	// UserID is not on the refresh token, so we need to set it now
	claims.UserID = user.ID

	accessToken, refreshToken, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return err
	}

	// set cookies on request with the access and refresh token
	// when cookie domain is localhost, this is dropped but expected
	if err := auth.SetAuthCookies(ctx, accessToken, refreshToken, h.CookieDomain); err != nil {
		return auth.ErrorResponse(err)
	}

	return ctx.JSON(http.StatusOK, Response{Message: "success"})
}
