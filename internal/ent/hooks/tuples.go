package hooks

import (
	"context"
	"fmt"

	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/fga"
)

// TODO: https://github.com/datumforge/datum/issues/262
// Not ideal to hard code `organization` in the function but will revisit as a part of another issue
func createOrgTuple(ctx context.Context, c *fga.Client, org, relation, object string) ([]ofgaclient.ClientTupleKey, error) {
	tuples := []ofgaclient.ClientTupleKey{{
		User:     fmt.Sprintf("organization:%s", org),
		Relation: relation,
		Object:   object,
	}}

	return tuples, nil
}

func getTupleKey(userID, objectID, objectType string, role enums.Role) (fga.TupleKey, error) {
	fgaRelation, err := roleToRelation(role)
	if err != nil {
		return fga.NewTupleKey(), err
	}

	sub := fga.Entity{
		Kind:       "user",
		Identifier: userID,
	}

	object := fga.Entity{
		Kind:       fga.Kind(objectType),
		Identifier: objectID,
	}

	return fga.TupleKey{
		Subject:  sub,
		Object:   object,
		Relation: fga.Relation(fgaRelation),
	}, nil
}
