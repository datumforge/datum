package graphapi

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// searchResult is a generic struct to hold the result of a search operation
type searchResult[T any] struct {
	result T
	err    error
}

// searchOrganizations searches for organizations based on the query string looking for matches in the name, description and display name
func searchOrganizations(ctx context.Context, query string, c chan<- searchResult[[]*generated.Organization]) {
	res, err := withTransactionalMutation(ctx).Organization.Query().Where(
		organization.Or(
			organization.NameContains(query),        // search by name
			organization.DescriptionContains(query), // search by description
			organization.DisplayNameContains(query), // search by display name
		),
	).All(ctx)

	c <- searchResult[[]*generated.Organization]{result: res, err: err}
}

// searchGroups searches for groups based on the query string looking for matches in the name, description and display name
func searchGroups(ctx context.Context, query string, c chan<- searchResult[[]*generated.Group]) {
	res, err := withTransactionalMutation(ctx).Group.Query().Where(
		group.Or(
			group.NameContains(query),        // search by name
			group.DescriptionContains(query), // search by description
			group.DisplayNameContains(query), // search by display name
		),
	).All(ctx)

	c <- searchResult[[]*generated.Group]{result: res, err: err}
}

// searchUsers searches for org members based on the query string looking for matches in the email, display name, first name and last name
func searchUsers(ctx context.Context, query string, c chan<- searchResult[[]*generated.User]) {
	members, err := withTransactionalMutation(ctx).OrgMembership.Query().Where(
		orgmembership.Or(
			orgmembership.HasUserWith(user.EmailContains(query)),       // search by email
			orgmembership.HasUserWith(user.DisplayNameContains(query)), // search by display name
			orgmembership.HasUserWith(user.FirstNameContains(query)),   // search by first name
			orgmembership.HasUserWith(user.LastNameContains(query)),    // search by last name
		),
	).WithUser().All(ctx)

	if members == nil || err != nil {
		c <- searchResult[[]*generated.User]{result: nil, err: err}
		return
	}

	users := make([]*generated.User, 0, len(members))
	for _, member := range members {
		users = append(users, member.Edges.User)
	}

	c <- searchResult[[]*generated.User]{result: users, err: err}
}

// searchSubscriber searches for subscribers based on the query string looking for matches in the email
func searchSubscriber(ctx context.Context, query string, c chan<- searchResult[[]*generated.Subscriber]) {
	res, err := withTransactionalMutation(ctx).Subscriber.Query().Where(
		subscriber.Or(
			subscriber.EmailContains(query), // search by email
		),
	).All(ctx)

	c <- searchResult[[]*generated.Subscriber]{result: res, err: err}
}
