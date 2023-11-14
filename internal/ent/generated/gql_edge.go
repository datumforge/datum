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
	return result, MaskNotFound(err)
}

func (gr *Group) Setting(ctx context.Context) (*GroupSettings, error) {
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

func (gr *Group) Owner(ctx context.Context) (*Organization, error) {
	result, err := gr.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QueryOwner().Only(ctx)
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

func (op *OauthProvider) User(ctx context.Context) (*User, error) {
	result, err := op.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = op.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
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

func (o *Organization) Setting(ctx context.Context) (*OrganizationSettings, error) {
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

func (pat *PersonalAccessToken) User(ctx context.Context) (*User, error) {
	result, err := pat.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pat.QueryUser().Only(ctx)
	}
	return result, err
}

func (rt *RefreshToken) User(ctx context.Context) (*User, error) {
	result, err := rt.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = rt.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Session) Users(ctx context.Context) (*User, error) {
	result, err := s.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryUsers().Only(ctx)
	}
	return result, MaskNotFound(err)
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

func (u *User) Sessions(ctx context.Context) (result []*Session, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedSessions(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.SessionsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QuerySessions().All(ctx)
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

func (u *User) Setting(ctx context.Context) (*UserSettings, error) {
	result, err := u.Edges.SettingOrErr()
	if IsNotLoaded(err) {
		result, err = u.QuerySetting().Only(ctx)
	}
	return result, err
}

func (u *User) Refreshtoken(ctx context.Context) (result []*RefreshToken, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedRefreshtoken(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.RefreshtokenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryRefreshtoken().All(ctx)
	}
	return result, err
}

func (u *User) Oauthprovider(ctx context.Context) (result []*OauthProvider, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedOauthprovider(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.OauthproviderOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryOauthprovider().All(ctx)
	}
	return result, err
}
