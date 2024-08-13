package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string) (*GlobalSearchResultConnection, error) {
	var (
		orgResults []*generated.Organization
		orgErr     error

		groupResults []*generated.Group
		groupErr     error

		userResults []*generated.User
		userErr     error

		subscriberResults []*generated.Subscriber
		subscriberErr     error

		entityResults []*generated.Entity
		entityErr     error
	)

	r.withPool().SubmitMultipleAndWait([]func(){
		func() {
			orgResults, orgErr = searchOrganizations(ctx, query)
		},
		func() {
			groupResults, groupErr = searchGroups(ctx, query)
		},
		func() {
			userResults, userErr = searchUsers(ctx, query)
		},
		func() {
			subscriberResults, subscriberErr = searchSubscribers(ctx, query)
		},
		func() {
			entityResults, entityErr = searchEntities(ctx, query)
		},
	})

	// Check all errors and return a single error if any of the searches failed
	if orgErr != nil || groupErr != nil || userErr != nil || subscriberErr != nil || entityErr != nil {
		r.logger.Errorw("search failed", "error",
			"org", orgErr,
			"group", groupErr,
			"user", userErr,
			"subscriber", subscriberErr,
			"entity", entityErr,
		)

		return nil, ErrSearchFailed
	}

	// return the results
	return &GlobalSearchResultConnection{
		Nodes: []GlobalSearchResult{
			OrganizationSearchResult{
				Organizations: orgResults,
			},
			GroupSearchResult{
				Groups: groupResults,
			},
			UserSearchResult{
				Users: userResults,
			},
			SubscriberSearchResult{
				Subscribers: subscriberResults,
			},
			EntitySearchResult{
				Entities: entityResults,
			},
		},
	}, nil
}
