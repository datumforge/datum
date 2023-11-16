// Package fga is a wrapper around openfga client
// credit to https://github.com/canonical/ofga/blob/main/tuples.go
// TODO: can we contribute this back once we have this in a working place
package fga

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/datumforge/datum/internal/echox"

	"github.com/openfga/go-sdk/client"
)

type TupleKey struct {
	Subject  Entity
	Object   Entity
	Relation Relation `json:"relation"`
}

func NewTupleKey() TupleKey { return TupleKey{} }

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
	Kind       Kind
	Identifier string
	Relation   Relation
}

// String returns a string representation of the entity/entity-set.
func (e *Entity) String() string {
	if e.Relation == "" {
		return e.Kind.String() + ":" + e.Identifier
	}

	return e.Kind.String() + ":" + e.Identifier + "#" + e.Relation.String()
}

// ParseEntity will parse a string representation into an Entity. It expects to
// find entities of the form:
//   - <entityType>:<Identifier>
//     eg. organization:datum
//   - <entityType>:<Identifier>#<relationship-set>
//     eg. organization:datum#member
func ParseEntity(s string) (Entity, error) {
	// entities should only contain a single colon
	c := strings.Count(s, ":")
	if c != 1 {
		return Entity{}, newInvalidEntityError(s)
	}

	match := entityRegex.FindStringSubmatch(s)
	if match == nil {
		return Entity{}, newInvalidEntityError(s)
	}

	// Extract and return the relevant information from the sub-matches.
	return Entity{
		Kind:       Kind(match[1]),
		Identifier: match[2],
		Relation:   Relation(match[4]),
	}, nil
}

// CreateCheckTupleWithUser gets the user id (currently the jwt sub, but that will change) and creates a Check Request for openFGA
func (c *Client) CreateCheckTupleWithUser(ctx context.Context, relation, object string) (*client.ClientCheckRequest, error) {
	if relation == "" {
		return nil, ErrMissingRelation
	}

	if object == "" {
		return nil, ErrMissingObject
	}

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
