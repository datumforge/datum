// Package fga is a wrapper around openfga client
// credit to https://github.com/canonical/ofga/blob/main/tuples.go
// TODO: can we contribute this back once we have this in a working place
package fga

import (
	"context"
	"fmt"
	"regexp"

	"github.com/datumforge/datum/internal/echox"

	"github.com/openfga/go-sdk/client"
)

// entityRegex is used to validate that a string represents an Entity/EntitySet
// and helps to convert from a string representation into an Entity struct.
var entityRegex = regexp.MustCompile(`([A-za-z0-9_][A-za-z0-9_-]*):([A-za-z0-9_][A-za-z0-9_@.+-]*)(#([A-za-z0-9_][A-za-z0-9_-]*))?`)

// Kind represents the type of the entity in OpenFGA.
type Kind string

// String implements the Stringer interface.
func (k Kind) String() string {
	return string(k)
}

// Relation represents the type of relation between entities in OpenFGA.
type Relation string

// String implements the Stringer interface.
func (r Relation) String() string {
	return string(r)
}

// Entity represents an entity/entity-set in OpenFGA.
// Example: `user:<user-id>`, `org:<org-id>#member`
type Entity struct {
	Kind     Kind
	ID       string
	Relation Relation
}

// String returns a string representation of the entity/entity-set.
func (e *Entity) String() string {
	if e.Relation == "" {
		return e.Kind.String() + ":" + e.ID
	}

	return e.Kind.String() + ":" + e.ID + "#" + e.Relation.String()
}

// ParseEntity will parse a string representation into an Entity. It expects to
// find entities of the form:
//   - <entityType>:<ID>
//     eg. organization:canonical
//   - <entityType>:<ID>#<relationship-set>
//     eg. organization:canonical#member
func ParseEntity(s string) (Entity, error) {
	match := entityRegex.FindStringSubmatch(s)
	if match == nil {
		return Entity{}, newInvalidEntityError(s)
	}

	// Extract and return the relevant information from the sub-matches.
	return Entity{
		Kind:     Kind(match[1]),
		ID:       match[2],
		Relation: Relation(match[4]),
	}, nil
}

// CreateCheckTupleWithUser gets the user id (currently the jwt sub, but that will change) and creates a Check Request for openFGA
func (c *Client) CreateCheckTupleWithUser(ctx context.Context, relation, object string) (*client.ClientCheckRequest, error) {
	ec, err := echox.EchoContextFromContext(ctx)
	if err != nil {
		c.Logger.Errorw("unable to get echo context", "error", err)

		return nil, err
	}

	actor, err := echox.GetActorSubject(*ec)
	if err != nil {
		return nil, err
	}

	// TODO: convert jwt sub --> uuid

	return &client.ClientCheckRequest{
		User:     fmt.Sprintf("user:%s", actor),
		Relation: relation,
		Object:   object,
	}, nil
}

// CreateRelationshipTupleWithUser gets the user id (currently the jwt sub, but that will change) and creates a relationship tuple
// with the given relation and object reference
func (c *Client) CreateRelationshipTupleWithUser(ctx context.Context, relation, object string) error {
	ec, err := echox.EchoContextFromContext(ctx)
	if err != nil {
		c.Logger.Errorw("unable to get echo context", "error", err)

		return err
	}

	actor, err := echox.GetActorSubject(*ec)
	if err != nil {
		return err
	}

	// TODO: convert jwt sub --> uuid

	tuples := []client.ClientTupleKey{{
		User:     fmt.Sprintf("user:%s", actor),
		Relation: relation,
		Object:   object,
	}}

	return c.createRelationshipTuple(ctx, tuples)
}

// CreateRelationshipTuple creates a relationship tuple in the openFGA store
func (c *Client) createRelationshipTuple(ctx context.Context, tuples []client.ClientTupleKey) error {
	if _, err := c.O.WriteTuples(ctx).Body(tuples).Execute(); err != nil {
		c.Logger.Infof("CreateRelationshipTuple error: [%s][%v]", err.Error())

		return err
	}

	return nil
}
