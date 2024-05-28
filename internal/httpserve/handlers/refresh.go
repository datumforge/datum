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

// RefreshHandler allows users to refresh their access token using their refresh token.
func (h *Handler) RefreshHandler(ctx echo.Context) error {
	var in models.RefreshRequest
	if err := ctx.Bind(&in); err != nil {
		return h.BadRequest(ctx, err)
	}

	if in.RefreshToken == "" {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(rout.NewMissingRequiredFieldError("refresh_token")))
	}

	// verify the refresh token
	claims, err := h.TM.Verify(in.RefreshToken)
	if err != nil {
		h.Logger.Errorw("error verifying token", "error", err)

		return h.BadRequest(ctx, err)
	}

	// check user in the database, sub == claims subject and ensure only one record is returned
	user, err := h.getUserByMappingID(ctx.Request().Context(), claims.Subject)
	if err != nil {
		if ent.IsNotFound(err) {
			return ctx.JSON(http.StatusNotFound, ErrNoAuthUser)
		}

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	// ensure the user is still active
	if user.Edges.Setting.Status != "ACTIVE" {
		return ctx.JSON(http.StatusNotFound, ErrNoAuthUser)
	}

	// UserID is not on the refresh token, so we need to set it now
	claims.UserID = user.MappingID

	accessToken, refreshToken, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
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
	}

	return h.Success(ctx, out)
}

// BindRefreshHandler is used to bind the refresh endpoint to the OpenAPI schema
func (h *Handler) BindRefreshHandler() *openapi3.Operation {
	refresh := openapi3.NewOperation()
	refresh.Description = "The Refresh endpoint re-authenticates users and API keys using a refresh token rather than requiring a username and password or API key credentials a second time and returns a new access and refresh token pair with the current credentials of the user. This endpoint is intended to facilitate long-running connections to datum systems that last longer than the duration of an access token; e.g. long sessions on the Datum UI or (especially) long running publishers and subscribers (machine users) that need to stay authenticated semi-permanently."
	refresh.OperationID = "RefreshHandler"

	h.AddRequestBody("RefreshRequest", models.RefreshRequest{}, refresh)
	h.AddResponse("RefreshReply", "success", models.RefreshReply{}, refresh, http.StatusOK)
	refresh.AddResponse(http.StatusInternalServerError, internalServerError())
	refresh.AddResponse(http.StatusBadRequest, badRequest())
	refresh.AddResponse(http.StatusNotFound, notFound())

	return refresh
}
