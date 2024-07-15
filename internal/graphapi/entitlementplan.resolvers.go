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

// CreateEntitlementPlan is the resolver for the createEntitlementPlan field.
func (r *mutationResolver) CreateEntitlementPlan(ctx context.Context, input generated.CreateEntitlementPlanInput) (*EntitlementPlanCreatePayload, error) {
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, input.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, rout.NewMissingRequiredFieldError("owner_id")
	}

	res, err := withTransactionalMutation(ctx).EntitlementPlan.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "entitlementplan"}, r.logger)
	}

	return &EntitlementPlanCreatePayload{
		EntitlementPlan: res,
	}, nil
}

// CreateBulkEntitlementPlan is the resolver for the createBulkEntitlementPlan field.
func (r *mutationResolver) CreateBulkEntitlementPlan(ctx context.Context, input []*generated.CreateEntitlementPlanInput) (*EntitlementPlanBulkCreatePayload, error) {
	return r.bulkCreateEntitlementPlan(ctx, input)
}

// CreateBulkCSVEntitlementPlan is the resolver for the createBulkCSVEntitlementPlan field.
func (r *mutationResolver) CreateBulkCSVEntitlementPlan(ctx context.Context, input graphql.Upload) (*EntitlementPlanBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateEntitlementPlanInput](input)
	if err != nil {
		r.logger.Errorw("failed to unmarshal bulk data", "error", err)

		return nil, err
	}

	return r.bulkCreateEntitlementPlan(ctx, data)
}

// UpdateEntitlementPlan is the resolver for the updateEntitlementPlan field.
func (r *mutationResolver) UpdateEntitlementPlan(ctx context.Context, id string, input generated.UpdateEntitlementPlanInput) (*EntitlementPlanUpdatePayload, error) {
	res, err := withTransactionalMutation(ctx).EntitlementPlan.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitlementplan"}, r.logger)
	}
	// set the organization in the auth context if its not done for us
	if err := setOrganizationInAuthContext(ctx, &res.OwnerID); err != nil {
		r.logger.Errorw("failed to set organization in auth context", "error", err)

		return nil, ErrPermissionDenied
	}

	res, err = res.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "entitlementplan"}, r.logger)
	}

	return &EntitlementPlanUpdatePayload{
		EntitlementPlan: res,
	}, nil
}

// DeleteEntitlementPlan is the resolver for the deleteEntitlementPlan field.
func (r *mutationResolver) DeleteEntitlementPlan(ctx context.Context, id string) (*EntitlementPlanDeletePayload, error) {
	if err := withTransactionalMutation(ctx).EntitlementPlan.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "entitlementplan"}, r.logger)
	}

	if err := generated.EntitlementPlanEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &EntitlementPlanDeletePayload{
		DeletedID: id,
	}, nil
}

// EntitlementPlan is the resolver for the entitlementPlan field.
func (r *queryResolver) EntitlementPlan(ctx context.Context, id string) (*generated.EntitlementPlan, error) {
	res, err := withTransactionalMutation(ctx).EntitlementPlan.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "entitlementplan"}, r.logger)
	}

	return res, nil
}
