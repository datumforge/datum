package hooks

import (
	"context"
	"fmt"

	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/fga"
)

func createTuple(ctx context.Context, c *fga.Client, relation, object string) ([]ofgaclient.ClientTupleKey, error) {
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

	tuples := []ofgaclient.ClientTupleKey{{
		User:     fmt.Sprintf("user:%s", actor),
		Relation: relation,
		Object:   object,
	}}

	return tuples, nil
}
