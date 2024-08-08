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

// DefaultAllRelations is the default list of relations to check
// these come from the fga/model/datum.fga file relations
var DefaultAllRelations = []string{
	"can_view",
	"can_edit",
	"can_delete",
	"audit_log_viewer",
	"can_invite_admins",
	"can_invite_members",
}

// AccountAccessHandler checks if a subject has access to an object
func (h *Handler) AccountListRolesHandler(ctx echo.Context) error {
	var in models.AccountListRolesRequest
	if err := ctx.Bind(&in); err != nil {
		return h.InvalidInput(ctx, err)
	}

	if err := in.Validate(); err != nil {
		return h.BadRequest(ctx, err)
	}

	req := fgax.ListAccess{
		SubjectType: in.SubjectType,
		ObjectID:    in.ObjectID,
		ObjectType:  fgax.Kind(in.ObjectType),
		Relations:   in.Relations,
	}

	// if no relations are provided, default to all relations
	if len(req.Relations) == 0 {
		req.Relations = DefaultAllRelations
	}

	subjectID, err := auth.GetUserIDFromContext(ctx.Request().Context())
	if err != nil {
		h.Logger.Error("error getting user id from context", "error", err)

		return h.InternalServerError(ctx, err)
	}

	req.SubjectID = subjectID

	roles, err := h.DBClient.Authz.ListRelations(ctx.Request().Context(), req)
	if err != nil {
		h.Logger.Error("error checking access", "error", err)

		return h.InternalServerError(ctx, err)
	}

	return h.Success(ctx, models.AccountListRolesReply{
		Reply: rout.Reply{Success: true},
		Roles: roles,
	})
}

// BindAccountListRoles returns the OpenAPI3 operation for accepting an account list roles request
func (h *Handler) BindAccountListRoles() *openapi3.Operation {
	listRoles := openapi3.NewOperation()
	listRoles.Description = "List roles a subject has in relation to an object"
	listRoles.OperationID = "AccountListRoles"
	listRoles.Security = &openapi3.SecurityRequirements{
		openapi3.SecurityRequirement{
			"bearerAuth": []string{},
		},
	}

	h.AddRequestBody("AccountListRolesRequest", models.ExampleAccountListRolesRequest, listRoles)
	h.AddResponse("AccountListRolesReply", "success", models.ExampleAccountListRolesReply, listRoles, http.StatusOK)
	listRoles.AddResponse(http.StatusInternalServerError, internalServerError())
	listRoles.AddResponse(http.StatusBadRequest, badRequest())
	listRoles.AddResponse(http.StatusUnauthorized, unauthorized())

	return listRoles
}
