package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateEntitlement is the resolver for the createEntitlement field.
func (r *mutationResolver) CreateEntitlement(ctx context.Context, input generated.CreateEntitlementInput) (*EntitlementCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateEntitlement - createEntitlement"))
}

// UpdateEntitlement is the resolver for the updateEntitlement field.
func (r *mutationResolver) UpdateEntitlement(ctx context.Context, id string, input generated.UpdateEntitlementInput) (*EntitlementUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateEntitlement - updateEntitlement"))
}

// DeleteEntitlement is the resolver for the deleteEntitlement field.
func (r *mutationResolver) DeleteEntitlement(ctx context.Context, id string) (*EntitlementDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteEntitlement - deleteEntitlement"))
}

// Entitlement is the resolver for the entitlement field.
func (r *queryResolver) Entitlement(ctx context.Context, id string) (*generated.Entitlement, error) {
	panic(fmt.Errorf("not implemented: Entitlement - entitlement"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
