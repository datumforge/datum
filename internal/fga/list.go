package fga

import (
	"context"

	ofgaclient "github.com/openfga/go-sdk/client"
)

// listObjects checks the openFGA store for all objects associated with a user+relation
func (c *Client) listObjects(ctx context.Context, req ofgaclient.ClientListObjectsRequest) (*ofgaclient.ClientListObjectsResponse, error) {
	list, err := c.Ofga.ListObjects(ctx).Body(req).Execute()
	if err != nil {
		c.Logger.Errorw("error listing objects",
			"user", req.User,
			"relation", req.Relation,
			"type", req.Type,
			"error", err.Error())

		return nil, err
	}

	return list, nil
}

// ListObjectsRequest creates the ClientListObjectsRequest and queries the FGA store for all objects with the user+relation
func (c *Client) ListObjectsRequest(ctx context.Context, userID, objectType, relation string) (*ofgaclient.ClientListObjectsResponse, error) {
	sub := Entity{
		Kind:       "user",
		Identifier: userID,
	}

	listReq := ofgaclient.ClientListObjectsRequest{
		User:     sub.String(),
		Relation: relation,
		Type:     objectType,
		// TODO: Support contextual tuples
	}

	c.Logger.Infow("listing objects", "relation", "user", sub.String(), relation, "type", objectType)

	return c.listObjects(ctx, listReq)
}

// ListContains checks the results of an fga ListObjects and parses the entities
// to get the identifier to compare to another identifier based on entity type
func ListContains(entityType string, l []string, i string) bool {
	for _, o := range l {
		e, _ := ParseEntity(o)

		// make sure its the correct entity type
		if e.Kind.String() != entityType {
			continue
		}

		if i == e.Identifier {
			return true
		}
	}

	return false
}
