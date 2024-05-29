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

// CreateOrgMembership is the resolver for the createOrgMembership field.
func (r *mutationResolver) CreateOrgMembership(ctx context.Context, input generated.CreateOrgMembershipInput) (*OrgMembershipCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &input.OrganizationID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	om, err := withTransactionalMutation(ctx).OrgMembership.Create().SetInput(input).Save(ctx)
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
			return nil, newPermissionDeniedError(ActionCreate, "org member")
		}

		r.logger.Errorw("failed to create org member", "error", err)

		return nil, err
	}

	return &OrgMembershipCreatePayload{OrgMembership: om}, nil
}

// CreateBulkOrgMembership is the resolver for the createBulkOrgMembership field.
func (r *mutationResolver) CreateBulkOrgMembership(ctx context.Context, input []*generated.CreateOrgMembershipInput) (*OrgMembershipBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkOrgMembership - createBulkOrgMembership"))
}

// CreateBulkCSVOrgMembership is the resolver for the createBulkCSVOrgMembership field.
func (r *mutationResolver) CreateBulkCSVOrgMembership(ctx context.Context, input graphql.Upload) (*OrgMembershipBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkCSVOrgMembership - createBulkCSVOrgMembership"))
}

// UpdateOrgMembership is the resolver for the updateOrgMembership field.
func (r *mutationResolver) UpdateOrgMembership(ctx context.Context, id string, input generated.UpdateOrgMembershipInput) (*OrgMembershipUpdatePayload, error) {
	orgMember, err := withTransactionalMutation(ctx).OrgMembership.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get org member on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "org member")
		}

		r.logger.Errorw("failed to get org member", "error", err)
		return nil, ErrInternalServerError
	}

	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &orgMember.OrganizationID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	orgMember, err = orgMember.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update org member", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "org member")
		}

		r.logger.Errorw("failed to update org member role", "error", err)
		return nil, ErrInternalServerError
	}

	return &OrgMembershipUpdatePayload{OrgMembership: orgMember}, nil
}

// DeleteOrgMembership is the resolver for the deleteOrgMembership field.
func (r *mutationResolver) DeleteOrgMembership(ctx context.Context, id string) (*OrgMembershipDeletePayload, error) {
	if err := withTransactionalMutation(ctx).OrgMembership.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "org member")
		}

		r.logger.Errorw("failed to delete org member", "error", err)
		return nil, err
	}

	if err := generated.OrgMembershipEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &OrgMembershipDeletePayload{DeletedID: id}, nil
}

// OrgMembership is the resolver for the orgMembership field.
func (r *queryResolver) OrgMembership(ctx context.Context, id string) (*generated.OrgMembership, error) {
	org, err := withTransactionalMutation(ctx).OrgMembership.Get(ctx, id)
	if err != nil {
		r.logger.Errorw("failed to get members of organization", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionGet, "org members")
		}

		return nil, ErrInternalServerError
	}

	return org, nil
}
