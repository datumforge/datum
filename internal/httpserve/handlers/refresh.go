package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
)

// RefreshHandler allows users to refresh their access token using their refresh token

func (h *Handler) RefreshHandler(ctx echo.Context) error {
	var in models.RefreshRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return h.InvalidInput(ctx, err)
	}

	// verify the refresh token
	claims, err := h.TokenManager.Verify(in.RefreshToken)
	if err != nil {
		h.Logger.Errorw("error verifying token", "error", err)

		return h.BadRequest(ctx, err)
	}

	// check user in the database, sub == claims subject and ensure only one record is returned
	user, err := h.getUserDetailsByID(ctx.Request().Context(), claims.Subject)
	if err != nil {
		if ent.IsNotFound(err) {
			return h.NotFound(ctx, ErrNoAuthUser)
		}

		return h.InternalServerError(ctx, ErrProcessingRequest)
	}

	// ensure the user is still active
	if user.Edges.Setting.Status != "ACTIVE" {
		return h.NotFound(ctx, ErrNoAuthUser)
	}

	// UserID is not on the refresh token, so we need to set it now
	claims.UserID = user.ID

	accessToken, refreshToken, err := h.TokenManager.CreateTokenPair(claims)
	if err != nil {
		h.Logger.Errorw("error creating token pair", "error", err)

		return h.InternalServerError(ctx, ErrProcessingRequest)
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, accessToken, refreshToken)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, user.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return err
	}

	out := &models.RefreshReply{
		Reply:   rout.Reply{Success: true},
		Message: "success",
		AuthData: models.AuthData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return h.Success(ctx, out)
}

// BindRefreshHandler is used to bind the refresh endpoint to the OpenAPI schema
func (h *Handler) BindRefreshHandler() *openapi3.Operation {
	refresh := openapi3.NewOperation()
	refresh.Description = "The Refresh endpoint re-authenticates users and API keys using a refresh token rather than requiring a username and password or API key credentials a second time and returns a new access and refresh token pair with the current credentials of the user. This endpoint is intended to facilitate long-running connections to datum systems that last longer than the duration of an access token; e.g. long sessions on the Datum UI or (especially) long running publishers and subscribers (machine users) that need to stay authenticated semi-permanently."
	refresh.OperationID = "RefreshHandler"
	refresh.Security = &openapi3.SecurityRequirements{
		openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		},
		openapi3.SecurityRequirement{
			"basicAuth": []string{},
		},
	}

	h.AddRequestBody("RefreshRequest", models.ExampleRefreshRequest, refresh)
	h.AddResponse("RefreshReply", "success", models.ExampleRefreshSuccessResponse, refresh, http.StatusOK)
	refresh.AddResponse(http.StatusInternalServerError, internalServerError())
	refresh.AddResponse(http.StatusBadRequest, badRequest())
	refresh.AddResponse(http.StatusNotFound, notFound())

	return refresh
}
