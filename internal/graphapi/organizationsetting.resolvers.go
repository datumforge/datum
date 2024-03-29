package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/rout"
)

// UpdateOrganizationSetting is the resolver for the updateOrganizationSetting field.
func (r *mutationResolver) UpdateOrganizationSetting(ctx context.Context, id string, input generated.UpdateOrganizationSettingInput) (*OrganizationSettingUpdatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	organizationSetting, err := withTransactionalMutation(ctx).OrganizationSetting.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied
		}

		r.logger.Errorw("failed to get user setting", "error", err)
		return nil, ErrInternalServerError
	}

	organizationSetting, err = organizationSetting.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			ve := err.(*generated.ValidationError)

			return nil, rout.InvalidField(ve.Name)
		}

		r.logger.Errorw("failed to update user setting", "error", err)
		return nil, err
	}

	return &OrganizationSettingUpdatePayload{OrganizationSetting: organizationSetting}, nil
}

// OrganizationSetting is the resolver for the organizationSetting field.
func (r *queryResolver) OrganizationSetting(ctx context.Context, id string) (*generated.OrganizationSetting, error) {
	// setup view context
	v := viewer.UserViewer{
		OrgID: id,
	}

	ctx = viewer.NewContext(ctx, v)

	org, err := withTransactionalMutation(ctx).OrganizationSetting.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get organization settings", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "organization")
		}

		return nil, ErrInternalServerError
	}

	return org, nil
}
