package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// CreatePersonalAccessToken is the resolver for the createPersonalAccessToken field.
func (r *mutationResolver) CreatePersonalAccessToken(ctx context.Context, input generated.CreatePersonalAccessTokenInput) (*PersonalAccessTokenCreatePayload, error) {
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
	panic(fmt.Errorf("not implemented: CreateOrganizationSetting - createOrganizationSetting"))
}

// DeletePersonalAccessToken is the resolver for the deletePersonalAccessToken field.
func (r *mutationResolver) DeletePersonalAccessToken(ctx context.Context, id string) (*PersonalAccessTokenDeletePayload, error) {
	if err := withTransactionalMutation(ctx).PersonalAccessToken.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to delete personal access token", "error", err)
		return nil, err
	}

	return &PersonalAccessTokenDeletePayload{DeletedID: id}, nil
}

// PersonalAccessToken is the resolver for the PersonalAccessToken field.
func (r *queryResolver) PersonalAccessToken(ctx context.Context, id string) (*generated.PersonalAccessToken, error) {
	if r.authDisabled {
		ctx = privacy.DecisionContext(ctx, privacy.Allow)
	} else {
		// setup view context
		ctx = viewer.NewContext(ctx, viewer.NewUserViewerFromSubject(ctx))
	}

	pat, err := r.client.PersonalAccessToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		r.logger.Errorw("failed to get token", "error", err)
		return nil, ErrInternalServerError
	}

	return pat, nil
}
