package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
	OrgID        string `json:"org_id,omitempty"`
	ParentOrgID  string `json:"parent_org_id,omitempty"`
}

// RefreshLogin allows users to refresh their access token using their refresh token.
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

	accessToken, refreshToken, err := h.refreshToken(claims, r)
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

func (h *Handler) refreshToken(claims *tokens.Claims, r RefreshRequest) (string, string, error) {
	orgID := claims.OrgID
	if r.OrgID != "" {
		orgID = r.OrgID
	}

	parentOrgID := claims.ParentOrgID
	if r.ParentOrgID != "" {
		orgID = r.ParentOrgID
	}

	// TODO: confirm user is part of that org

	// Create a new claims object using the user retrieved from the database
	refreshClaims := &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: claims.Subject,
		},
		UserID:      claims.UserID,
		OrgID:       orgID,
		ParentOrgID: parentOrgID,
	}

	return h.TM.CreateTokenPair(refreshClaims)
}
