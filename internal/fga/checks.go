package fga

import (
	"context"

	ofgaclient "github.com/openfga/go-sdk/client"
)

// CheckTuple checks the openFGA store for provided relationship tuple
func (c *Client) CheckTuple(ctx context.Context, check ofgaclient.ClientCheckRequest) (bool, error) {
	data, err := c.Ofga.Check(ctx).Body(check).Execute()
	if err != nil {
		c.Logger.Errorw("error checking tuple", "tuple", check, "error", err.Error())

		return false, err
	}

	return *data.Allowed, nil
}

func (c *Client) CheckOrgAccess(ctx context.Context, userID, orgID, relation string) (bool, error) {
	sub := Entity{
		Kind:       "user",
		Identifier: userID,
	}

	obj := Entity{
		Kind:       "organization",
		Identifier: orgID,
	}

	c.Logger.Infow("checking relationship tuples", "relation", relation, "object", obj.String())

	checkReq := ofgaclient.ClientCheckRequest{
		User:     sub.String(),
		Relation: relation,
		Object:   obj.String(),
	}

	return c.CheckTuple(ctx, checkReq)
}

func (c *Client) CheckGroupAccess(ctx context.Context, userID, groupID, relation string) (bool, error) {
	sub := Entity{
		Kind:       "user",
		Identifier: userID,
	}

	obj := Entity{
		Kind:       "group",
		Identifier: groupID,
	}

	c.Logger.Infow("checking relationship tuples", "relation", relation, "object", obj.String())

	checkReq := ofgaclient.ClientCheckRequest{
		User:     sub.String(),
		Relation: relation,
		Object:   obj.String(),
	}

	return c.CheckTuple(ctx, checkReq)
}
