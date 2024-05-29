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
)

// CreateOhAuthTooToken is the resolver for the createOhAuthTooToken field.
func (r *mutationResolver) CreateOhAuthTooToken(ctx context.Context, input generated.CreateOhAuthTooTokenInput) (*OhAuthTooTokenCreatePayload, error) {
	res, err := withTransactionalMutation(ctx).OhAuthTooToken.Create().SetInput(input).Save(ctx)
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
			return nil, newPermissionDeniedError(ActionCreate, "ohauthtootoken")
		}

		r.logger.Errorw("failed to create ohauthtootoken", "error", err)

		return nil, ErrInternalServerError
	}

	return &OhAuthTooTokenCreatePayload{
		OhAuthTooToken: res,
	}, nil
}

// CreateBulkOhAuthTooToken is the resolver for the createBulkOhAuthTooToken field.
func (r *mutationResolver) CreateBulkOhAuthTooToken(ctx context.Context, input []*generated.CreateOhAuthTooTokenInput) (*OhAuthTooTokenBulkCreatePayload, error) {
	return r.bulkCreateOhAuthTooToken(ctx, input)
}

// CreateBulkCSVOhAuthTooToken is the resolver for the createBulkCSVOhAuthTooToken field.
func (r *mutationResolver) CreateBulkCSVOhAuthTooToken(ctx context.Context, input graphql.Upload) (*OhAuthTooTokenBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateOhAuthTooTokenInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateOhAuthTooToken(ctx, data)
}

// UpdateOhAuthTooToken is the resolver for the updateOhAuthTooToken field.
func (r *mutationResolver) UpdateOhAuthTooToken(ctx context.Context, id string, input generated.UpdateOhAuthTooTokenInput) (*OhAuthTooTokenUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).OhAuthTooToken.Get(ctx, id)
	if err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to get ohauthtootoken on update", "error", err)

			return nil, newPermissionDeniedError(ActionGet, "ohauthtootoken")
		}

		r.logger.Errorw("failed to get ohauthtootoken", "error", err)
		return nil, ErrInternalServerError
	}

	res, err = res.Update().SetInput(input).Save(ctx)
	if err != nil {
		if generated.IsValidationError(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			r.logger.Errorw("failed to update ohauthtootoken", "error", err)

			return nil, newPermissionDeniedError(ActionUpdate, "ohauthtootoken")
		}

		r.logger.Errorw("failed to update ohauthtootoken", "error", err)
		return nil, ErrInternalServerError
	}

	return &OhAuthTooTokenUpdatePayload{
		OhAuthTooToken: res,
	}, nil
}

// DeleteOhAuthTooToken is the resolver for the deleteOhAuthTooToken field.
func (r *mutationResolver) DeleteOhAuthTooToken(ctx context.Context, id string) (*OhAuthTooTokenDeletePayload, error) {
	if err := withTransactionalMutation(ctx).OhAuthTooToken.DeleteOneID(id).Exec(ctx); err != nil {
		if generated.IsNotFound(err) {
			return nil, err
		}

		if errors.Is(err, privacy.Deny) {
			return nil, newPermissionDeniedError(ActionDelete, "ohauthtootoken")
		}

		r.logger.Errorw("failed to delete ohauthtootoken", "error", err)
		return nil, err
	}

	if err := generated.OhAuthTooTokenEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &OhAuthTooTokenDeletePayload{
		DeletedID: id,
	}, nil
}

// OhAuthTooToken is the resolver for the ohAuthTooToken field.
func (r *queryResolver) OhAuthTooToken(ctx context.Context, id string) (*generated.OhAuthTooToken, error) {
	panic(fmt.Errorf("not implemented: OhAuthTooToken - ohAuthTooToken"))
}
