package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/rout"
)

// CreateOrganizationSetting is the resolver for the createOrganizationSetting field.
func (r *mutationResolver) CreateOrganizationSetting(ctx context.Context, input generated.CreateOrganizationSettingInput) (*OrganizationSettingCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateOrganizationSetting - createOrganizationSetting"))
}

// CreateBulkOrganizationSetting is the resolver for the createBulkOrganizationSetting field.
func (r *mutationResolver) CreateBulkOrganizationSetting(ctx context.Context, input []*generated.CreateOrganizationSettingInput) (*OrganizationSettingBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkOrganizationSetting - createBulkOrganizationSetting"))
}

// CreateBulkCSVOrganizationSetting is the resolver for the createBulkCSVOrganizationSetting field.
func (r *mutationResolver) CreateBulkCSVOrganizationSetting(ctx context.Context, input graphql.Upload) (*OrganizationSettingBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkCSVOrganizationSetting - createBulkCSVOrganizationSetting"))
}

// UpdateOrganizationSetting is the resolver for the updateOrganizationSetting field.
func (r *mutationResolver) UpdateOrganizationSetting(ctx context.Context, id string, input generated.UpdateOrganizationSettingInput) (*OrganizationSettingUpdatePayload, error) {
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

// DeleteOrganizationSetting is the resolver for the deleteOrganizationSetting field.
func (r *mutationResolver) DeleteOrganizationSetting(ctx context.Context, id string) (*OrganizationSettingDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteOrganizationSetting - deleteOrganizationSetting"))
}

// OrganizationSetting is the resolver for the organizationSetting field.
func (r *queryResolver) OrganizationSetting(ctx context.Context, id string) (*generated.OrganizationSetting, error) {
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
