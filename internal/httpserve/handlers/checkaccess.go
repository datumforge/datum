package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/datumforge/fgax"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
)

// CheckAccessHandler checks if a subject has access to an object
func (h *Handler) CheckAccessHandler(ctx echo.Context) error {
	var in models.CheckAccessRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return h.BadRequest(ctx, err)
	}

	req := fgax.AccessCheck{
		SubjectType: in.SubjectType,
		Relation:    in.Relation,
		ObjectID:    in.ObjectID,
		ObjectType:  fgax.Kind(in.ObjectType),
	}

	subjectID, err := auth.GetUserIDFromContext(ctx.Request().Context())
	if err != nil {
		h.Logger.Error("error getting user id from context", "error", err)

		return h.InternalServerError(ctx, err)
	}

	req.SubjectID = subjectID

	allow, err := h.DBClient.Authz.CheckAccess(ctx.Request().Context(), req)
	if err != nil {
		h.Logger.Error("error checking access", "error", err)

		return h.InternalServerError(ctx, err)
	}

	return h.Success(ctx, models.CheckAccessReply{
		Reply:   rout.Reply{Success: true},
		Allowed: allow,
	})
}

// BindCheckAccess returns the OpenAPI3 operation for accepting an check access request
func (h *Handler) BindCheckAccess() *openapi3.Operation {
	checkAccess := openapi3.NewOperation()
	checkAccess.Description = "Check Subject Access to Object"
	checkAccess.OperationID = "CheckAccess"
	checkAccess.Security = &openapi3.SecurityRequirements{}

	h.AddRequestBody("CheckAccessRequest", models.ExampleInviteRequest, checkAccess)
	h.AddResponse("CheckAccessReply", "success", models.ExampleInviteResponse, checkAccess, http.StatusOK)
	checkAccess.AddResponse(http.StatusInternalServerError, internalServerError())
	checkAccess.AddResponse(http.StatusBadRequest, badRequest())
	checkAccess.AddResponse(http.StatusUnauthorized, unauthorized())

	return checkAccess
}
