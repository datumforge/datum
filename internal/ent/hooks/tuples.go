package hooks

import (
	"context"
	"fmt"

	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/fga"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

func createTuple(ctx context.Context, c *fga.Client, relation, object string) ([]ofgaclient.ClientTupleKey, error) {
	actor, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		c.Logger.Errorw("unable to get user ID from context", "error", err)

		return nil, err
	}

	tuples := []ofgaclient.ClientTupleKey{{
		User:     fmt.Sprintf("user:%s", actor),
		Relation: relation,
		Object:   object,
	}}

	return tuples, nil
}

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
