package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/google/uuid"
)

// CreateOrganization is the resolver for the createOrganization field.
func (r *mutationResolver) CreateOrganization(ctx context.Context, input generated.CreateOrganizationInput) (*OrganizationCreatePayload, error) {
	// TODO - add permissions checks
	org, err := r.client.Organization.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create organization", "error", err)
		return nil, ErrInternalServerError
	}

	return &OrganizationCreatePayload{Organization: org}, nil
}

// UpdateOrganization is the resolver for the updateOrganization field.
func (r *mutationResolver) UpdateOrganization(ctx context.Context, id uuid.UUID, input generated.UpdateOrganizationInput) (*OrganizationUpdatePayload, error) {
	// TODO - add permissions checks

	org, err := r.client.Organization.Get(ctx, id)
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
func (r *mutationResolver) DeleteOrganization(ctx context.Context, id uuid.UUID) (*OrganizationDeletePayload, error) {
	// TODO - add permissions checks

	if err := r.client.Organization.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete organization", "error", err)
		return nil, err
	}

	return &OrganizationDeletePayload{DeletedID: id}, nil
}

// Organization is the resolver for the organization field.
func (r *queryResolver) Organization(ctx context.Context, id uuid.UUID) (*generated.Organization, error) {
	// TODO - add permissions checks

	org, err := r.client.Organization.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get organization", "error", err)
		return nil, ErrInternalServerError
	}

	return org, nil
}
