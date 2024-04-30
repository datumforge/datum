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
)

// CreateAPIToken is the resolver for the createAPIToken field.
func (r *mutationResolver) CreateAPIToken(ctx context.Context, input generated.CreateAPITokenInput) (*APITokenCreatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	apiToken, err := withTransactionalMutation(ctx).APIToken.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if generated.IsConstraintError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create api token", "error", err)
		return nil, ErrInternalServerError
	}

	return &APITokenCreatePayload{APIToken: apiToken}, err
}

// UpdateAPIToken is the resolver for the updateAPIToken field.
func (r *mutationResolver) UpdateAPIToken(ctx context.Context, id string, input generated.UpdateAPITokenInput) (*APITokenUpdatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	apiToken, err := withTransactionalMutation(ctx).APIToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied

		}

		r.logger.Errorw("failed to get api token", "error", err)
		return nil, ErrInternalServerError
	}

	apiToken, err = apiToken.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if generated.IsConstraintError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update api token", "error", err)

		return nil, ErrInternalServerError
	}

	return &APITokenUpdatePayload{APIToken: apiToken}, err
}

// DeleteAPIToken is the resolver for the deleteAPIToken field.
func (r *mutationResolver) DeleteAPIToken(ctx context.Context, id string) (*APITokenDeletePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	if err := withTransactionalMutation(ctx).APIToken.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete api token", "error", err)
		return nil, err
	}

	if err := generated.APITokenEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &APITokenDeletePayload{DeletedID: id}, nil
}

// APIToken is the resolver for the apiToken field.
func (r *queryResolver) APIToken(ctx context.Context, id string) (*generated.APIToken, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	apiToken, err := withTransactionalMutation(ctx).APIToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get api token", "error", err)

		return nil, ErrInternalServerError
	}

	return apiToken, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
