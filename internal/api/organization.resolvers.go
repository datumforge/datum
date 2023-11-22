package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/fga"
)

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input generated.CreateOrganizationInput) (*OrganizationCreatePayload, error) {
	// TODO - add permissions checks
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

		r.logger.Errorw("failed to create organization", "error", err)
		return nil, ErrInternalServerError
	}

	return &OrganizationCreatePayload{Organization: org}, nil
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, id string, input generated.UpdateOrganizationInput) (*OrganizationUpdatePayload, error) {
	userID, err := echox.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	sub := fga.Entity{
		Kind:       "user",
		Identifier: userID,
	}

	obj := fga.Entity{
		Kind:       "organization",
		Identifier: id,
	}

	v := viewer.UserViewer{
		UserID: userID,
		Authz:  r.fgaClient,
		Key: fga.TupleKey{
			Subject:  sub,
			Relation: fga.CanEdit,
			Object:   obj,
		},
	}

	vc := viewer.NewContext(ctx, v)

	org, err := r.client.Organization.Get(vc, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get organization", "error", err)
		return nil, ErrInternalServerError
	}

	org, err = org.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update organization", "error", err)
		return nil, ErrInternalServerError
	}

	return &OrganizationUpdatePayload{Organization: org}, nil
}

// DeleteOrganization is the resolver for the deleteOrganization field.
func (r *mutationResolver) DeleteOrganization(ctx context.Context, id string) (*OrganizationDeletePayload, error) {
	userID, err := echox.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	sub := fga.Entity{
		Kind:       "user",
		Identifier: userID,
	}

	obj := fga.Entity{
		Kind:       "organization",
		Identifier: id,
	}

	v := viewer.UserViewer{
		UserID: userID,
		Authz:  r.fgaClient,
		Key: fga.TupleKey{
			Subject:  sub,
			Relation: fga.CanDelete,
			Object:   obj,
		},
	}

	vc := viewer.NewContext(ctx, v)

	if err := r.client.Organization.DeleteOneID(id).Exec(vc); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete organization", "error", err)
		return nil, err
	}

	return &OrganizationDeletePayload{DeletedID: id}, nil
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id string) (*generated.Organization, error) {
	userID, err := echox.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	sub := fga.Entity{
		Kind:       "user",
		Identifier: userID,
	}

	obj := fga.Entity{
		Kind:       "organization",
		Identifier: id,
	}

	v := viewer.UserViewer{
		UserID: userID,
		Authz:  r.fgaClient,
		Key: fga.TupleKey{
			Subject:  sub,
			Relation: fga.CanView,
			Object:   obj,
		},
	}

	vc := viewer.NewContext(ctx, v)

	org, err := r.client.Organization.Get(vc, id)
	if err != nil {
		r.logger.Errorw("failed to get organization", "error", err)

		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied
		}

		return nil, ErrInternalServerError
	}

	return org, nil
}
