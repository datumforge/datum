package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/fgax"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang-jwt/jwt/v5"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

// SwitchHandler is responsible for handling requests to the `/switch` endpoint, and changing the user's logged in organization context
func (h *Handler) SwitchHandler(ctx echo.Context) error {
	var in models.SwitchOrganizationRequest

	if err := ctx.Bind(&in); err != nil {
		return h.BadRequest(ctx, err)
	}

	reqCtx := ctx.Request().Context()

	userID, err := auth.GetUserIDFromContext(reqCtx)
	if err != nil {
		h.Logger.Errorw("unable to get user id from context", "error", err)

		return h.BadRequest(ctx, err)
	}

	// get user from database by subject
	user, err := h.getUserDetailsByID(reqCtx, userID)
	if err != nil {
		h.Logger.Errorw("unable to get user by subject", "error", err)

		return h.BadRequest(ctx, err)
	}

	orgID, err := auth.GetOrganizationIDFromContext(reqCtx)
	if err != nil {
		h.Logger.Errorw("unable to get organization id from context", "error", err)

		return h.BadRequest(ctx, err)
	}

	// ensure the user is not already in the target organization
	if orgID == in.TargetOrganizationID {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse("already switched to organization"))
	}

	// ensure user is already a member of the destination organization
	if allow, err := h.DBClient.Authz.CheckOrgAccess(reqCtx, userID, auth.UserSubjectType, orgID, fgax.CanView); err != nil || !allow {
		h.Logger.Errorw("user not authorized to access organization", "error", err)

		return ctx.JSON(http.StatusUnauthorized, rout.ErrorResponse("unauthorized"))
	}

	// get the target organization
	orgGetCtx := privacy.DecisionContext(reqCtx, privacy.Allow)

	org, err := h.getOrgByID(orgGetCtx, in.TargetOrganizationID)
	if err != nil {
		h.Logger.Errorw("unable to get target organization by id", "error", err)
	}

	// create new claims for the user
	newClaims := switchClaims(user, org.MappingID)

	// create a new token pair for the user
	access, refresh, err := h.TM.CreateTokenPair(newClaims)
	if err != nil {
		return h.InternalServerError(ctx, err)
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, user.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return h.InternalServerError(ctx, err)
	}

	// return the session value for the UI to use
	session, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return h.InternalServerError(ctx, err)
	}

	// track the organization switch event
	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("target_organization_id", org.ID).
		Set("auth_provider", user.AuthProvider).
		Set("previous_organization_id", orgID)

	h.AnalyticsClient.Event("organization_switched", props)

	// set the out attributes we send back to the client only on success
	out := &models.SwitchOrganizationReply{
		Reply:        rout.Reply{Success: true},
		AccessToken:  access,
		RefreshToken: refresh,
		Session:      session,
	}

	return h.Success(ctx, out)
}

// switchClaims creates a new set of claims for the user based on the target organization and returns them
func switchClaims(u *generated.User, targetOrgMappingID string) *tokens.Claims {
	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.MappingID,
		},
		UserID: u.MappingID,
		OrgID:  targetOrgMappingID,
	}
}

// BindResetPassword binds the reset password handler to the OpenAPI schema
func (h *Handler) BindSwitchHandler() *openapi3.Operation {
	switchHandler := openapi3.NewOperation()
	switchHandler.Description = "Switch the user's organization context"
	switchHandler.OperationID = "OrganizationSwitch"

	h.AddRequestBody("SwitchOrganizationRequest", models.PublishRequest{}, switchHandler)
	h.AddResponse("SwitchOrganizationReply", "success", models.PublishReply{}, switchHandler, http.StatusOK)
	switchHandler.AddResponse(http.StatusInternalServerError, internalServerError())
	switchHandler.AddResponse(http.StatusBadRequest, badRequest())
	switchHandler.AddResponse(http.StatusUnauthorized, unauthorized())

	return switchHandler
}
