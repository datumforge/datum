package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateHush is the resolver for the createHush field
func (r *mutationResolver) CreateHush(ctx context.Context, input generated.CreateHushInput) (*HushCreatePayload, error) {
	t, err := withTransactionalMutation(ctx).Hush.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionCreate, object: "hush"}, r.logger)
	}

	return &HushCreatePayload{Hush: t}, nil
}

// CreateBulkHush is the resolver for the createBulkHush field.
func (r *mutationResolver) CreateBulkHush(ctx context.Context, input []*generated.CreateHushInput) (*HushBulkCreatePayload, error) {
	return r.bulkCreateHush(ctx, input)
}

// CreateBulkCSVHush is the resolver for the createBulkCSVHush field.
func (r *mutationResolver) CreateBulkCSVHush(ctx context.Context, input graphql.Upload) (*HushBulkCreatePayload, error) {
	data, err := unmarshalBulkData[generated.CreateHushInput](input)
	if err != nil {
		return nil, err
	}

	return r.bulkCreateHush(ctx, data)
}

// UpdateHush is the resolver for the updateHush field
func (r *mutationResolver) UpdateHush(ctx context.Context, id string, input generated.UpdateHushInput) (*HushUpdatePayload, error) {
	hush, err := withTransactionalMutation(ctx).Hush.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "hush"}, r.logger)
	}

	hush, err = hush.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionUpdate, object: "hush"}, r.logger)
	}

	return &HushUpdatePayload{Hush: hush}, nil
}

// DeleteHush is the resolver for the deleteHush field
func (r *mutationResolver) DeleteHush(ctx context.Context, id string) (*HushDeletePayload, error) {
	if err := withTransactionalMutation(ctx).Hush.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, parseRequestError(err, action{action: ActionDelete, object: "hush"}, r.logger)
	}

	if err := generated.HushEdgeCleanup(ctx, id); err != nil {
		return nil, newCascadeDeleteError(err)
	}

	return &HushDeletePayload{DeletedID: id}, nil
}

// Hush is the resolver for the hush field
func (r *queryResolver) Hush(ctx context.Context, id string) (*generated.Hush, error) {
	hush, err := withTransactionalMutation(ctx).Hush.Get(ctx, id)
	if err != nil {
		return nil, parseRequestError(err, action{action: ActionGet, object: "hush"}, r.logger)
	}

	return hush, nil
}
