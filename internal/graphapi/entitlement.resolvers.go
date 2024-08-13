package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/rout"
)

// CreateEntitlement is the resolver for the createEntitlement field.
func (r *mutationResolver) CreateEntitlement(ctx context.Context, input generated.CreateEntitlementInput) (*EntitlementCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).Entitlement.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "entitlement"}, r.logger)
	}

	return &EntitlementCreatePayload{
		Entitlement: res,
	}, nil
}

// CreateBulkEntitlement is the resolver for the createBulkEntitlement field.
func (r *mutationResolver) CreateBulkEntitlement(ctx context.Context, input []*generated.CreateEntitlementInput) (*EntitlementBulkCreatePayload, error) {
	return r.bulkCreateEntitlement(ctx, input)
}

// CreateBulkCSVEntitlement is the resolver for the createBulkCSVEntitlement field.
func (r *mutationResolver) CreateBulkCSVEntitlement(ctx context.Context, input graphql.Upload) (*EntitlementBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateEntitlementInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateEntitlement(ctx, data)
}

// UpdateEntitlement is the resolver for the updateEntitlement field.
func (r *mutationResolver) UpdateEntitlement(ctx context.Context, id string, input generated.UpdateEntitlementInput) (*EntitlementUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).Entitlement.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitlement"}, r.logger)
	}
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, ErrPermissionDenied
	}

	// setup update request
	req := res.Update().SetInput(input).AppendTags(input.AppendTags)

	res, err = req.Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitlement"}, r.logger)
	}

	return &EntitlementUpdatePayload{
		Entitlement: res,
	}, nil
}

// DeleteEntitlement is the resolver for the deleteEntitlement field.
func (r *mutationResolver) DeleteEntitlement(ctx context.Context, id string) (*EntitlementDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Entitlement.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "entitlement"}, r.logger)
	}

	if err := generated.EntitlementEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &EntitlementDeletePayload{
		DeletedID: id,
	}, nil
}

// Entitlement is the resolver for the entitlement field.
func (r *queryResolver) Entitlement(ctx context.Context, id string) (*generated.Entitlement, error) {
	res, err := withTransactionalMutation(ctx).Entitlement.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "entitlement"}, r.logger)
	}

	return res, nil
}
