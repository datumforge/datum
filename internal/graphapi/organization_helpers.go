package api

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
)

func (r *mutationResolver) createOrg(ctx context.Context, input generated.CreateOrganizationInput) (*OrganizationCreatePayload, error) {
	org, err := r.client.Organization.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			validationError := err.(*generated.ValidationError)

			r.logger.Debugw("validation error", "field", validationError.Name, "error", validationError.Error())

			return nil, validationError
		}

		if generated.IsConstraintError(err) {
			constraintError := err.(*generated.ConstraintError)

			r.logger.Debugw("constraint error", "error", constraintError.Error())

			return nil, constraintError
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionCreate, "organization")
		}

		r.logger.Errorw("failed to create organization", "error", err)

		return nil, ErrInternalServerError
	}

	return &OrganizationCreatePayload{Organization: org}, nil
}
