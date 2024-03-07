// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (e *Entitlement) Owner(ctx context.Context) (*Organization, error) {
	result, err := e.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryOwner().Only(ctx)
	}
	return result, err
}

func (gr *Group) Owner(ctx context.Context) (*Organization, error) {
	result, err := gr.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryOwner().Only(ctx)
	}
	return result, err
}

func (gr *Group) Setting(ctx context.Context) (*GroupSetting, error) {
	result, err := gr.Edges.SettingOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QuerySetting().Only(ctx)
	}
	return result, err
}

func (gr *Group) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryUsers().All(ctx)
	}
	return result, err
}

func (gr *Group) Members(ctx context.Context) (result []*GroupMembership, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedMembers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.MembersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryMembers().All(ctx)
	}
	return result, err
}

func (gm *GroupMembership) Group(ctx context.Context) (*Group, error) {
	result, err := gm.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = gm.QueryGroup().Only(ctx)
	}
	return result, err
}

func (gm *GroupMembership) User(ctx context.Context) (*User, error) {
	result, err := gm.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = gm.QueryUser().Only(ctx)
	}
	return result, err
}

func (gs *GroupSetting) Group(ctx context.Context) (*Group, error) {
	result, err := gs.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = gs.QueryGroup().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Integration) Owner(ctx context.Context) (*Organization, error) {
	result, err := i.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (i *Invite) Owner(ctx context.Context) (*Organization, error) {
	result, err := i.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = i.QueryOwner().Only(ctx)
	}
	return result, err
}

func (op *OauthProvider) Owner(ctx context.Context) (*Organization, error) {
	result, err := op.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = op.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (om *OrgMembership) Organization(ctx context.Context) (*Organization, error) {
	result, err := om.Edges.OrganizationOrErr()
	if IsNotLoaded(err) {
		result, err = om.QueryOrganization().Only(ctx)
	}
	return result, err
}

func (om *OrgMembership) User(ctx context.Context) (*User, error) {
	result, err := om.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = om.QueryUser().Only(ctx)
	}
	return result, err
}

func (o *Organization) Parent(ctx context.Context) (*Organization, error) {
	result, err := o.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = o.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (o *Organization) Children(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *OrganizationOrder, where *OrganizationWhereInput,
) (*OrganizationConnection, error) {
	opts := []OrganizationPaginateOption{
		WithOrganizationOrder(orderBy),
		WithOrganizationFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := o.Edges.totalCount[1][alias]
	if nodes, err := o.NamedChildren(alias); err == nil || hasTotalCount {
		pager, err := newOrganizationPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &OrganizationConnection{Edges: []*OrganizationEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return o.QueryChildren().Paginate(ctx, after, first, before, last, opts...)
}

func (o *Organization) Groups(ctx context.Context) (result []*Group, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedGroups(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.GroupsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryGroups().All(ctx)
	}
	return result, err
}

func (o *Organization) Integrations(ctx context.Context) (result []*Integration, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedIntegrations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.IntegrationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryIntegrations().All(ctx)
	}
	return result, err
}

func (o *Organization) Setting(ctx context.Context) (*OrganizationSetting, error) {
	result, err := o.Edges.SettingOrErr()
	if IsNotLoaded(err) {
		result, err = o.QuerySetting().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (o *Organization) Entitlements(ctx context.Context) (result []*Entitlement, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedEntitlements(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.EntitlementsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryEntitlements().All(ctx)
	}
	return result, err
}

func (o *Organization) PersonalAccessTokens(ctx context.Context) (result []*PersonalAccessToken, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedPersonalAccessTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.PersonalAccessTokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryPersonalAccessTokens().All(ctx)
	}
	return result, err
}

func (o *Organization) Oauthprovider(ctx context.Context) (result []*OauthProvider, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedOauthprovider(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.OauthproviderOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryOauthprovider().All(ctx)
	}
	return result, err
}

func (o *Organization) Users(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.UsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryUsers().All(ctx)
	}
	return result, err
}

func (o *Organization) Invites(ctx context.Context) (result []*Invite, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedInvites(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.InvitesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryInvites().All(ctx)
	}
	return result, err
}

func (o *Organization) Members(ctx context.Context) (result []*OrgMembership, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = o.NamedMembers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = o.Edges.MembersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = o.QueryMembers().All(ctx)
	}
	return result, err
}

func (os *OrganizationSetting) Organization(ctx context.Context) (*Organization, error) {
	result, err := os.Edges.OrganizationOrErr()
	if IsNotLoaded(err) {
		result, err = os.QueryOrganization().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pat *PersonalAccessToken) Owner(ctx context.Context) (*User, error) {
	result, err := pat.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = pat.QueryOwner().Only(ctx)
	}
	return result, err
}

func (pat *PersonalAccessToken) Organizations(ctx context.Context) (result []*Organization, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pat.NamedOrganizations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pat.Edges.OrganizationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pat.QueryOrganizations().All(ctx)
	}
	return result, err
}

func (ts *TFASettings) Owner(ctx context.Context) (*User, error) {
	result, err := ts.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = ts.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) PersonalAccessTokens(ctx context.Context) (result []*PersonalAccessToken, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPersonalAccessTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PersonalAccessTokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPersonalAccessTokens().All(ctx)
	}
	return result, err
}

func (u *User) TfaSettings(ctx context.Context) (*TFASettings, error) {
	result, err := u.Edges.TfaSettingsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryTfaSettings().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Setting(ctx context.Context) (*UserSetting, error) {
	result, err := u.Edges.SettingOrErr()
	if IsNotLoaded(err) {
		result, err = u.QuerySetting().Only(ctx)
	}
	return result, err
}

func (u *User) Groups(ctx context.Context) (result []*Group, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroups(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroups().All(ctx)
	}
	return result, err
}

func (u *User) Organizations(ctx context.Context) (result []*Organization, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedOrganizations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.OrganizationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryOrganizations().All(ctx)
	}
	return result, err
}

func (u *User) GroupMemberships(ctx context.Context) (result []*GroupMembership, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroupMemberships(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupMembershipsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroupMemberships().All(ctx)
	}
	return result, err
}

func (u *User) OrgMemberships(ctx context.Context) (result []*OrgMembership, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedOrgMemberships(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.OrgMembershipsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryOrgMemberships().All(ctx)
	}
	return result, err
}

func (us *UserSetting) User(ctx context.Context) (*User, error) {
	result, err := us.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = us.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (us *UserSetting) DefaultOrg(ctx context.Context) (*Organization, error) {
	result, err := us.Edges.DefaultOrgOrErr()
	if IsNotLoaded(err) {
		result, err = us.QueryDefaultOrg().Only(ctx)
	}
	return result, MaskNotFound(err)
}
