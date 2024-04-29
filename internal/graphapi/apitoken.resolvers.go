package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateAPIToken is the resolver for the createAPIToken field.
func (r *mutationResolver) CreateAPIToken(ctx context.Context, input generated.CreateAPITokenInput) (*APITokenCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateAPIToken - createAPIToken"))
}

// UpdateAPIToken is the resolver for the updateAPIToken field.
func (r *mutationResolver) UpdateAPIToken(ctx context.Context, id string, input generated.UpdateAPITokenInput) (*APITokenUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateAPIToken - updateAPIToken"))
}

// DeleteAPIToken is the resolver for the deleteAPIToken field.
func (r *mutationResolver) DeleteAPIToken(ctx context.Context, id string) (*APITokenDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteAPIToken - deleteAPIToken"))
}

// APIToken is the resolver for the apiToken field.
func (r *queryResolver) APIToken(ctx context.Context, id string) (*generated.APIToken, error) {
	panic(fmt.Errorf("not implemented: APIToken - apiToken"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
