package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
	_ "github.com/datumforge/datum/internal/ent/generated/runtime"
	"github.com/datumforge/datum/pkg/rout"
)

// CreateIntegration is the resolver for the createIntegration field.
func (r *mutationResolver) CreateIntegration(ctx context.Context, input generated.CreateIntegrationInput) (*IntegrationCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	i, err := withTransactionalMutation(ctx).Integration.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create integration", "error", err)
		return nil, ErrInternalServerError
	}

	return &IntegrationCreatePayload{Integration: i}, nil
}

// CreateBulkIntegration is the resolver for the createBulkIntegration field.
func (r *mutationResolver) CreateBulkIntegration(ctx context.Context, input []*generated.CreateIntegrationInput) (*IntegrationBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkIntegration - createBulkIntegration"))
}

// CreateBulkCSVIntegration is the resolver for the createBulkCSVIntegration field.
func (r *mutationResolver) CreateBulkCSVIntegration(ctx context.Context, input graphql.Upload) (*IntegrationBulkCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateBulkCSVIntegration - createBulkCSVIntegration"))
}

// UpdateIntegration is the resolver for the updateIntegration field.
func (r *mutationResolver) UpdateIntegration(ctx context.Context, id string, input generated.UpdateIntegrationInput) (*IntegrationUpdatePayload, error) {
	i, err := withTransactionalMutation(ctx).Integration.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get integration", "error", err)
		return nil, ErrInternalServerError
	}

	if err := setOrganizationInAuthContext(ctx, &i.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, ErrPermissionDenied
	}

	i, err = i.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update integration", "error", err)
		return nil, ErrInternalServerError
	}

	return &IntegrationUpdatePayload{Integration: i}, nil
}

// DeleteIntegration is the resolver for the deleteIntegration field.
func (r *mutationResolver) DeleteIntegration(ctx context.Context, id string) (*IntegrationDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Integration.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete integration", "error", err)
		return nil, err
	}

	if err := generated.IntegrationEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &IntegrationDeletePayload{DeletedID: id}, nil
}

// Integration is the resolver for the integration field.
func (r *queryResolver) Integration(ctx context.Context, id string) (*generated.Integration, error) {
	i, err := withTransactionalMutation(ctx).Integration.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get integration", "error", err)
		return nil, ErrInternalServerError
	}

	return i, nil
}
