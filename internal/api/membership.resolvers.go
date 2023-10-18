package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.39

import (
	"context"
	"fmt"

	"github.com/datumforge/datum/internal/ent/generated"
)

// CreateMembership is the resolver for the createMembership field.
func (r *mutationResolver) CreateMembership(ctx context.Context, input generated.CreateMembershipInput) (*MembershipCreatePayload, error) {
	panic(fmt.Errorf("not implemented: CreateMembership - createMembership"))
}

// UpdateMembership is the resolver for the updateMembership field.
func (r *mutationResolver) UpdateMembership(ctx context.Context, id string, input generated.UpdateMembershipInput) (*MembershipUpdatePayload, error) {
	panic(fmt.Errorf("not implemented: UpdateMembership - updateMembership"))
}

// DeleteMembership is the resolver for the deleteMembership field.
func (r *mutationResolver) DeleteMembership(ctx context.Context, id string) (*MembershipDeletePayload, error) {
	panic(fmt.Errorf("not implemented: DeleteMembership - deleteMembership"))
}

// Membership is the resolver for the membership field.
func (r *queryResolver) Membership(ctx context.Context, id string) (*generated.Membership, error) {
	panic(fmt.Errorf("not implemented: Membership - membership"))
}
