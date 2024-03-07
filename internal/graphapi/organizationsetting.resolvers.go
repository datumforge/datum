package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// UpdateOrganizationSetting is the resolver for the updateOrganizationSetting field.
func (r *mutationResolver) UpdateOrganizationSetting(ctx context.Context, id string, input generated.UpdateOrganizationSettingInput) (*OrganizationSettingUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateOrganizationSetting - updateOrganizationSetting"))
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateOrganizationSetting(ctx context.Context, input generated.CreateOrganizationSettingInput) (*OrganizationSettingCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateOrganizationSetting - createOrganizationSetting"))
}
func (r *mutationResolver) DeleteOrganizationSetting(ctx context.Context, id string) (*OrganizationSettingDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteOrganizationSetting - deleteOrganizationSetting"))
}
