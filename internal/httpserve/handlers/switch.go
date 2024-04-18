package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/golang-jwt/jwt/v5"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

// SwitchOrganizationRequest contains the target organization ID being switched to
type SwitchOrganizationRequest struct {
	TargetOrganizationID string `json:"target_organization_id"`
}

// SwitchOrganizationReply holds the new authentication and session information for the user for the new organization
type SwitchOrganizationReply struct {
	rout.Reply
	AccessToken            string `json:"access_token"`
	RefreshToken           string `json:"refresh_token"`
	Session                string `json:"session"`
	NewOrganizationContext string `json:"new_organization_context"`
}

// SwitchHandler is responsible for handling requests to the `/switch` endpoint, and changing the user's logged in organization context
func (h *Handler) SwitchHandler(ctx echo.Context) error {
	var req SwitchOrganizationRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// get the current user claims
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// ensure the user is not already in the target organization
	if claims.OrgID == req.TargetOrganizationID {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse("already switched to organization"))
	}

	newUserCtx := viewer.NewContext(ctx.Request().Context(), viewer.NewUserViewerFromID(claims.UserID, true))

	// get the user from the database by subject
	user, err := h.getUserBySub(newUserCtx, claims.Subject)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, ErrNoAuthUser)
		}

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	// ensure user is already a member of the destination organization
	if err := h.confirmOrgMembership(newUserCtx, claims.UserID, req.TargetOrganizationID); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// create new claims for the user
	newClaims := switchClaims(user, req.TargetOrganizationID)

	newClaims.UserID = user.ID
	// create a new token pair for the user
	access, refresh, err := h.TM.CreateTokenPair(newClaims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, user.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// return the session value for the UI to use
	session, err := sessions.SessionToken(newUserCtx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// last sanity check ensuring that newly issued claims match the requested organization ID, which does not match what the user starts with, and the target organization ID which the user is already a member
	if newClaims.OrgID != req.TargetOrganizationID {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse("currently issued claims do not match target organization id"))
	}

	// track the organization switch event
	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("target_organization_id", newClaims.OrgID).
		Set("auth_provider", user.AuthProvider).
		Set("previous_organization_id", claims.OrgID)

	h.AnalyticsClient.Event("organization_switched", props)

	// set the out attributes we send back to the client only on success
	out := &SwitchOrganizationReply{
		Reply:                  rout.Reply{Success: true},
		AccessToken:            access,
		RefreshToken:           refresh,
		Session:                session,
		NewOrganizationContext: newClaims.OrgID,
	}

	return ctx.JSON(http.StatusOK, out)
}

// switchClaims creates a new set of claims for the user based on the target organization and returns them
func switchClaims(u *generated.User, targetOrg string) *tokens.Claims {
	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.ID,
		},
		UserID:      u.ID,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		AvatarURL:   *u.AvatarRemoteURL,
		OrgID:       targetOrg,
	}
}
