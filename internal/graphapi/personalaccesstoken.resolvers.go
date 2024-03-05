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

// CreatePersonalAccessToken is the resolver for the createPersonalAccessToken field.
func (r *mutationResolver) CreatePersonalAccessToken(ctx context.Context, input generated.CreatePersonalAccessTokenInput) (*PersonalAccessTokenCreatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	pat, err := withTransactionalMutation(ctx).PersonalAccessToken.Create().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if generated.IsConstraintError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to create personal access token", "error", err)
		return nil, ErrInternalServerError
	}

	return &PersonalAccessTokenCreatePayload{PersonalAccessToken: pat}, err
}

// UpdatePersonalAccessToken is the resolver for the updatePersonalAccessToken field.
func (r *mutationResolver) UpdatePersonalAccessToken(ctx context.Context, id string, input generated.UpdatePersonalAccessTokenInput) (*PersonalAccessTokenUpdatePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	pat, err := withTransactionalMutation(ctx).PersonalAccessToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, ErrPermissionDenied

		}

		r.logger.Errorw("failed to get personal access token", "error", err)
		return nil, ErrInternalServerError
	}

	pat, err = pat.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if generated.IsConstraintError(err) {
			return nil, err
		}

		r.logger.Errorw("failed to update personal access token", "error", err)

		return nil, ErrInternalServerError
	}

	return &PersonalAccessTokenUpdatePayload{PersonalAccessToken: pat}, err
}

// DeletePersonalAccessToken is the resolver for the deletePersonalAccessToken field.
func (r *mutationResolver) DeletePersonalAccessToken(ctx context.Context, id string) (*PersonalAccessTokenDeletePayload, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	if err := withTransactionalMutation(ctx).PersonalAccessToken.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete personal access token", "error", err)
		return nil, err
	}

	if err := generated.PersonalAccessTokenEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &PersonalAccessTokenDeletePayload{DeletedID: id}, nil
}

// PersonalAccessToken is the resolver for the PersonalAccessToken field.
func (r *queryResolver) PersonalAccessToken(ctx context.Context, id string) (*generated.PersonalAccessToken, error) {
	ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))

	pat, err := withTransactionalMutation(ctx).PersonalAccessToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get token", "error", err)

		return nil, ErrInternalServerError
	}

	return pat, nil
}
