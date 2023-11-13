package fga

import (
	"context"

	"github.com/openfga/go-sdk/client"
)

// checkTuple checks the openFGA store for provided relationship tuple
func (c *Client) checkTuple(ctx context.Context, check client.ClientCheckRequest) (bool, error) {
	data, err := c.O.Check(ctx).Body(check).Execute()
	if err != nil {
		c.Logger.Infof("GetCheck error: [%s]", err.Error())

		return false, err
	}

	return *data.Allowed, nil
}

// CheckDirectUser checks the user:<uuid> tuple relation given the object
func (c *Client) CheckDirectUser(ctx context.Context, relation, object string) (bool, error) {
	tuple, err := c.CreateCheckTupleWithUser(ctx, relation, object)
	if err != nil {
		return false, err
	}

	c.Logger.Infow(
		"Checking permissions",
		"user",
		tuple.User,
		"object",
		tuple.Object,
		"relation",
		tuple.Relation,
	)

	return c.checkTuple(ctx, *tuple)
}
