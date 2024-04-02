// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/invite"
	"github.com/datumforge/datum/internal/ent/generated/oauthprovider"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
	"github.com/datumforge/datum/internal/ent/generated/tier"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oc *OrganizationCreate) SetCreatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetCreatedAt(t)
	return oc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetCreatedAt(*t)
	}
	return oc
}

// SetUpdatedAt sets the "updated_at" field.
func (oc *OrganizationCreate) SetUpdatedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetUpdatedAt(t)
	return oc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetUpdatedAt(*t)
	}
	return oc
}

// SetCreatedBy sets the "created_by" field.
func (oc *OrganizationCreate) SetCreatedBy(s string) *OrganizationCreate {
	oc.mutation.SetCreatedBy(s)
	return oc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableCreatedBy(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetCreatedBy(*s)
	}
	return oc
}

// SetUpdatedBy sets the "updated_by" field.
func (oc *OrganizationCreate) SetUpdatedBy(s string) *OrganizationCreate {
	oc.mutation.SetUpdatedBy(s)
	return oc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableUpdatedBy(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetUpdatedBy(*s)
	}
	return oc
}

// SetDeletedAt sets the "deleted_at" field.
func (oc *OrganizationCreate) SetDeletedAt(t time.Time) *OrganizationCreate {
	oc.mutation.SetDeletedAt(t)
	return oc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDeletedAt(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetDeletedAt(*t)
	}
	return oc
}

// SetDeletedBy sets the "deleted_by" field.
func (oc *OrganizationCreate) SetDeletedBy(s string) *OrganizationCreate {
	oc.mutation.SetDeletedBy(s)
	return oc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDeletedBy(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDeletedBy(*s)
	}
	return oc
}

// SetName sets the "name" field.
func (oc *OrganizationCreate) SetName(s string) *OrganizationCreate {
	oc.mutation.SetName(s)
	return oc
}

// SetDisplayName sets the "display_name" field.
func (oc *OrganizationCreate) SetDisplayName(s string) *OrganizationCreate {
	oc.mutation.SetDisplayName(s)
	return oc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDisplayName(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDisplayName(*s)
	}
	return oc
}

// SetDescription sets the "description" field.
func (oc *OrganizationCreate) SetDescription(s string) *OrganizationCreate {
	oc.mutation.SetDescription(s)
	return oc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDescription(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetDescription(*s)
	}
	return oc
}

// SetParentOrganizationID sets the "parent_organization_id" field.
func (oc *OrganizationCreate) SetParentOrganizationID(s string) *OrganizationCreate {
	oc.mutation.SetParentOrganizationID(s)
	return oc
}

// SetNillableParentOrganizationID sets the "parent_organization_id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableParentOrganizationID(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetParentOrganizationID(*s)
	}
	return oc
}

// SetPersonalOrg sets the "personal_org" field.
func (oc *OrganizationCreate) SetPersonalOrg(b bool) *OrganizationCreate {
	oc.mutation.SetPersonalOrg(b)
	return oc
}

// SetNillablePersonalOrg sets the "personal_org" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillablePersonalOrg(b *bool) *OrganizationCreate {
	if b != nil {
		oc.SetPersonalOrg(*b)
	}
	return oc
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (oc *OrganizationCreate) SetAvatarRemoteURL(s string) *OrganizationCreate {
	oc.mutation.SetAvatarRemoteURL(s)
	return oc
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableAvatarRemoteURL(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetAvatarRemoteURL(*s)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OrganizationCreate) SetID(s string) *OrganizationCreate {
	oc.mutation.SetID(s)
	return oc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableID(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetID(*s)
	}
	return oc
}

// SetParentID sets the "parent" edge to the Organization entity by ID.
func (oc *OrganizationCreate) SetParentID(id string) *OrganizationCreate {
	oc.mutation.SetParentID(id)
	return oc
}

// SetNillableParentID sets the "parent" edge to the Organization entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableParentID(id *string) *OrganizationCreate {
	if id != nil {
		oc = oc.SetParentID(*id)
	}
	return oc
}

// SetParent sets the "parent" edge to the Organization entity.
func (oc *OrganizationCreate) SetParent(o *Organization) *OrganizationCreate {
	return oc.SetParentID(o.ID)
}

// AddChildIDs adds the "children" edge to the Organization entity by IDs.
func (oc *OrganizationCreate) AddChildIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddChildIDs(ids...)
	return oc
}

// AddChildren adds the "children" edges to the Organization entity.
func (oc *OrganizationCreate) AddChildren(o ...*Organization) *OrganizationCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddChildIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (oc *OrganizationCreate) AddGroupIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddGroupIDs(ids...)
	return oc
}

// AddGroups adds the "groups" edges to the Group entity.
func (oc *OrganizationCreate) AddGroups(g ...*Group) *OrganizationCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return oc.AddGroupIDs(ids...)
}

// AddIntegrationIDs adds the "integrations" edge to the Integration entity by IDs.
func (oc *OrganizationCreate) AddIntegrationIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddIntegrationIDs(ids...)
	return oc
}

// AddIntegrations adds the "integrations" edges to the Integration entity.
func (oc *OrganizationCreate) AddIntegrations(i ...*Integration) *OrganizationCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return oc.AddIntegrationIDs(ids...)
}

// SetSettingID sets the "setting" edge to the OrganizationSetting entity by ID.
func (oc *OrganizationCreate) SetSettingID(id string) *OrganizationCreate {
	oc.mutation.SetSettingID(id)
	return oc
}

// SetNillableSettingID sets the "setting" edge to the OrganizationSetting entity by ID if the given value is not nil.
func (oc *OrganizationCreate) SetNillableSettingID(id *string) *OrganizationCreate {
	if id != nil {
		oc = oc.SetSettingID(*id)
	}
	return oc
}

// SetSetting sets the "setting" edge to the OrganizationSetting entity.
func (oc *OrganizationCreate) SetSetting(o *OrganizationSetting) *OrganizationCreate {
	return oc.SetSettingID(o.ID)
}

// AddEntitlementIDs adds the "entitlements" edge to the Entitlement entity by IDs.
func (oc *OrganizationCreate) AddEntitlementIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddEntitlementIDs(ids...)
	return oc
}

// AddEntitlements adds the "entitlements" edges to the Entitlement entity.
func (oc *OrganizationCreate) AddEntitlements(e ...*Entitlement) *OrganizationCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return oc.AddEntitlementIDs(ids...)
}

// AddPersonalAccessTokenIDs adds the "personal_access_tokens" edge to the PersonalAccessToken entity by IDs.
func (oc *OrganizationCreate) AddPersonalAccessTokenIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddPersonalAccessTokenIDs(ids...)
	return oc
}

// AddPersonalAccessTokens adds the "personal_access_tokens" edges to the PersonalAccessToken entity.
func (oc *OrganizationCreate) AddPersonalAccessTokens(p ...*PersonalAccessToken) *OrganizationCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPersonalAccessTokenIDs(ids...)
}

// AddOauthproviderIDs adds the "oauthprovider" edge to the OauthProvider entity by IDs.
func (oc *OrganizationCreate) AddOauthproviderIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddOauthproviderIDs(ids...)
	return oc
}

// AddOauthprovider adds the "oauthprovider" edges to the OauthProvider entity.
func (oc *OrganizationCreate) AddOauthprovider(o ...*OauthProvider) *OrganizationCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOauthproviderIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (oc *OrganizationCreate) AddUserIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddUserIDs(ids...)
	return oc
}

// AddUsers adds the "users" edges to the User entity.
func (oc *OrganizationCreate) AddUsers(u ...*User) *OrganizationCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return oc.AddUserIDs(ids...)
}

// AddInviteIDs adds the "invites" edge to the Invite entity by IDs.
func (oc *OrganizationCreate) AddInviteIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddInviteIDs(ids...)
	return oc
}

// AddInvites adds the "invites" edges to the Invite entity.
func (oc *OrganizationCreate) AddInvites(i ...*Invite) *OrganizationCreate {
	ids := make([]string, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return oc.AddInviteIDs(ids...)
}

// AddSubscriberIDs adds the "subscribers" edge to the Subscriber entity by IDs.
func (oc *OrganizationCreate) AddSubscriberIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddSubscriberIDs(ids...)
	return oc
}

// AddSubscribers adds the "subscribers" edges to the Subscriber entity.
func (oc *OrganizationCreate) AddSubscribers(s ...*Subscriber) *OrganizationCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return oc.AddSubscriberIDs(ids...)
}

// AddTierIDs adds the "tiers" edge to the Tier entity by IDs.
func (oc *OrganizationCreate) AddTierIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddTierIDs(ids...)
	return oc
}

// AddTiers adds the "tiers" edges to the Tier entity.
func (oc *OrganizationCreate) AddTiers(t ...*Tier) *OrganizationCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return oc.AddTierIDs(ids...)
}

// AddMemberIDs adds the "members" edge to the OrgMembership entity by IDs.
func (oc *OrganizationCreate) AddMemberIDs(ids ...string) *OrganizationCreate {
	oc.mutation.AddMemberIDs(ids...)
	return oc
}

// AddMembers adds the "members" edges to the OrgMembership entity.
func (oc *OrganizationCreate) AddMembers(o ...*OrgMembership) *OrganizationCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddMemberIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	if err := oc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() error {
	if _, ok := oc.mutation.CreatedAt(); !ok {
		if organization.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized organization.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := organization.DefaultCreatedAt()
		oc.mutation.SetCreatedAt(v)
	}
	if _, ok := oc.mutation.UpdatedAt(); !ok {
		if organization.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organization.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organization.DefaultUpdatedAt()
		oc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oc.mutation.DisplayName(); !ok {
		v := organization.DefaultDisplayName
		oc.mutation.SetDisplayName(v)
	}
	if _, ok := oc.mutation.PersonalOrg(); !ok {
		v := organization.DefaultPersonalOrg
		oc.mutation.SetPersonalOrg(v)
	}
	if _, ok := oc.mutation.ID(); !ok {
		if organization.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized organization.DefaultID (forgotten import generated/runtime?)")
		}
		v := organization.DefaultID()
		oc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Organization.name"`)}
	}
	if v, ok := oc.mutation.Name(); ok {
		if err := organization.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Organization.name": %w`, err)}
		}
	}
	if _, ok := oc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`generated: missing required field "Organization.display_name"`)}
	}
	if v, ok := oc.mutation.DisplayName(); ok {
		if err := organization.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Organization.display_name": %w`, err)}
		}
	}
	if v, ok := oc.mutation.AvatarRemoteURL(); ok {
		if err := organization.AvatarRemoteURLValidator(v); err != nil {
			return &ValidationError{Name: "avatar_remote_url", err: fmt.Errorf(`generated: validator failed for field "Organization.avatar_remote_url": %w`, err)}
		}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Organization.ID type: %T", _spec.ID.Value)
		}
	}
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString))
	)
	_spec.Schema = oc.schemaConfig.Organization
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.CreatedAt(); ok {
		_spec.SetField(organization.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oc.mutation.UpdatedAt(); ok {
		_spec.SetField(organization.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oc.mutation.CreatedBy(); ok {
		_spec.SetField(organization.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := oc.mutation.UpdatedBy(); ok {
		_spec.SetField(organization.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := oc.mutation.DeletedAt(); ok {
		_spec.SetField(organization.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oc.mutation.DeletedBy(); ok {
		_spec.SetField(organization.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := oc.mutation.Name(); ok {
		_spec.SetField(organization.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oc.mutation.DisplayName(); ok {
		_spec.SetField(organization.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := oc.mutation.Description(); ok {
		_spec.SetField(organization.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := oc.mutation.PersonalOrg(); ok {
		_spec.SetField(organization.FieldPersonalOrg, field.TypeBool, value)
		_node.PersonalOrg = value
	}
	if value, ok := oc.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(organization.FieldAvatarRemoteURL, field.TypeString, value)
		_node.AvatarRemoteURL = &value
	}
	if nodes := oc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   organization.ParentTable,
			Columns: []string{organization.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Organization
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ParentOrganizationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.ChildrenTable,
			Columns: []string{organization.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Organization
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.GroupsTable,
			Columns: []string{organization.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Group
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.IntegrationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.IntegrationsTable,
			Columns: []string{organization.IntegrationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.SettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   organization.SettingTable,
			Columns: []string{organization.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.OrganizationSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.EntitlementsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.EntitlementsTable,
			Columns: []string{organization.EntitlementsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Entitlement
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.PersonalAccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   organization.PersonalAccessTokensTable,
			Columns: organization.PersonalAccessTokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(personalaccesstoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.OrganizationPersonalAccessTokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OauthproviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.OauthproviderTable,
			Columns: []string{organization.OauthproviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(oauthprovider.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.OauthProvider
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   organization.UsersTable,
			Columns: organization.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.OrgMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrgMembershipCreate{config: oc.config, mutation: newOrgMembershipMutation(oc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.InvitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.InvitesTable,
			Columns: []string{organization.InvitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invite.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Invite
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.SubscribersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.SubscribersTable,
			Columns: []string{organization.SubscribersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subscriber.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Subscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.TiersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   organization.TiersTable,
			Columns: []string{organization.TiersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.Tier
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.MembersTable,
			Columns: []string{organization.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orgmembership.FieldID, field.TypeString),
			},
		}
		edge.Schema = oc.schemaConfig.OrgMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	err      error
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	if ocb.err != nil {
		return nil, ocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
