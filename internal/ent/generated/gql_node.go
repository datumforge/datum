// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupmembership"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/invite"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
	"github.com/datumforge/datum/internal/ent/generated/tfasettings"
	"github.com/datumforge/datum/internal/ent/generated/tier"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/hashicorp/go-multierror"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Entitlement) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Group) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *GroupMembership) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *GroupSetting) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Integration) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Invite) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *OauthProvider) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *OhAuthTooToken) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *OrgMembership) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Organization) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *OrganizationSetting) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *PersonalAccessToken) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Subscriber) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *TFASettings) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Tier) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *User) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *UserSetting) IsNode() {}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, string) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, string) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, string) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id string) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id string, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id string) (Noder, error) {
	switch table {
	case entitlement.Table:
		query := c.Entitlement.Query().
			Where(entitlement.ID(id))
		query, err := query.CollectFields(ctx, "Entitlement")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		query := c.Group.Query().
			Where(group.ID(id))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupmembership.Table:
		query := c.GroupMembership.Query().
			Where(groupmembership.ID(id))
		query, err := query.CollectFields(ctx, "GroupMembership")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupsetting.Table:
		query := c.GroupSetting.Query().
			Where(groupsetting.ID(id))
		query, err := query.CollectFields(ctx, "GroupSetting")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case integration.Table:
		query := c.Integration.Query().
			Where(integration.ID(id))
		query, err := query.CollectFields(ctx, "Integration")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case invite.Table:
		query := c.Invite.Query().
			Where(invite.ID(id))
		query, err := query.CollectFields(ctx, "Invite")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case oauthprovider.Table:
		query := c.OauthProvider.Query().
			Where(oauthprovider.ID(id))
		query, err := query.CollectFields(ctx, "OauthProvider")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case ohauthtootoken.Table:
		query := c.OhAuthTooToken.Query().
			Where(ohauthtootoken.ID(id))
		query, err := query.CollectFields(ctx, "OhAuthTooToken")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case orgmembership.Table:
		query := c.OrgMembership.Query().
			Where(orgmembership.ID(id))
		query, err := query.CollectFields(ctx, "OrgMembership")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case organization.Table:
		query := c.Organization.Query().
			Where(organization.ID(id))
		query, err := query.CollectFields(ctx, "Organization")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case organizationsetting.Table:
		query := c.OrganizationSetting.Query().
			Where(organizationsetting.ID(id))
		query, err := query.CollectFields(ctx, "OrganizationSetting")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case personalaccesstoken.Table:
		query := c.PersonalAccessToken.Query().
			Where(personalaccesstoken.ID(id))
		query, err := query.CollectFields(ctx, "PersonalAccessToken")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case subscriber.Table:
		query := c.Subscriber.Query().
			Where(subscriber.ID(id))
		query, err := query.CollectFields(ctx, "Subscriber")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tfasettings.Table:
		query := c.TFASettings.Query().
			Where(tfasettings.ID(id))
		query, err := query.CollectFields(ctx, "TFASettings")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case tier.Table:
		query := c.Tier.Query().
			Where(tier.ID(id))
		query, err := query.CollectFields(ctx, "Tier")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		query := c.User.Query().
			Where(user.ID(id))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case usersetting.Table:
		query := c.UserSetting.Query().
			Where(usersetting.ID(id))
		query, err := query.CollectFields(ctx, "UserSetting")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []string, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]string)
	id2idx := make(map[string][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []string) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[string][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case entitlement.Table:
		query := c.Entitlement.Query().
			Where(entitlement.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Entitlement")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		query := c.Group.Query().
			Where(group.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupmembership.Table:
		query := c.GroupMembership.Query().
			Where(groupmembership.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GroupMembership")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupsetting.Table:
		query := c.GroupSetting.Query().
			Where(groupsetting.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GroupSetting")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case integration.Table:
		query := c.Integration.Query().
			Where(integration.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Integration")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case invite.Table:
		query := c.Invite.Query().
			Where(invite.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Invite")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case oauthprovider.Table:
		query := c.OauthProvider.Query().
			Where(oauthprovider.IDIn(ids...))
		query, err := query.CollectFields(ctx, "OauthProvider")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case ohauthtootoken.Table:
		query := c.OhAuthTooToken.Query().
			Where(ohauthtootoken.IDIn(ids...))
		query, err := query.CollectFields(ctx, "OhAuthTooToken")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case orgmembership.Table:
		query := c.OrgMembership.Query().
			Where(orgmembership.IDIn(ids...))
		query, err := query.CollectFields(ctx, "OrgMembership")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case organization.Table:
		query := c.Organization.Query().
			Where(organization.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Organization")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case organizationsetting.Table:
		query := c.OrganizationSetting.Query().
			Where(organizationsetting.IDIn(ids...))
		query, err := query.CollectFields(ctx, "OrganizationSetting")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case personalaccesstoken.Table:
		query := c.PersonalAccessToken.Query().
			Where(personalaccesstoken.IDIn(ids...))
		query, err := query.CollectFields(ctx, "PersonalAccessToken")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case subscriber.Table:
		query := c.Subscriber.Query().
			Where(subscriber.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Subscriber")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tfasettings.Table:
		query := c.TFASettings.Query().
			Where(tfasettings.IDIn(ids...))
		query, err := query.CollectFields(ctx, "TFASettings")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case tier.Table:
		query := c.Tier.Query().
			Where(tier.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Tier")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case usersetting.Table:
		query := c.UserSetting.Query().
			Where(usersetting.IDIn(ids...))
		query, err := query.CollectFields(ctx, "UserSetting")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}
