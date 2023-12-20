// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetUpdatedAt sets the "updated_at" field.
func (gu *GroupUpdate) SetUpdatedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetUpdatedAt(t)
	return gu
}

// SetUpdatedBy sets the "updated_by" field.
func (gu *GroupUpdate) SetUpdatedBy(s string) *GroupUpdate {
	gu.mutation.SetUpdatedBy(s)
	return gu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableUpdatedBy(s *string) *GroupUpdate {
	if s != nil {
		gu.SetUpdatedBy(*s)
	}
	return gu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gu *GroupUpdate) ClearUpdatedBy() *GroupUpdate {
	gu.mutation.ClearUpdatedBy()
	return gu
}

// SetDeletedAt sets the "deleted_at" field.
func (gu *GroupUpdate) SetDeletedAt(t time.Time) *GroupUpdate {
	gu.mutation.SetDeletedAt(t)
	return gu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeletedAt(t *time.Time) *GroupUpdate {
	if t != nil {
		gu.SetDeletedAt(*t)
	}
	return gu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gu *GroupUpdate) ClearDeletedAt() *GroupUpdate {
	gu.mutation.ClearDeletedAt()
	return gu
}

// SetDeletedBy sets the "deleted_by" field.
func (gu *GroupUpdate) SetDeletedBy(s string) *GroupUpdate {
	gu.mutation.SetDeletedBy(s)
	return gu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDeletedBy(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDeletedBy(*s)
	}
	return gu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gu *GroupUpdate) ClearDeletedBy() *GroupUpdate {
	gu.mutation.ClearDeletedBy()
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableName(s *string) *GroupUpdate {
	if s != nil {
		gu.SetName(*s)
	}
	return gu
}

// SetDescription sets the "description" field.
func (gu *GroupUpdate) SetDescription(s string) *GroupUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDescription(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDescription(*s)
	}
	return gu
}

// ClearDescription clears the value of the "description" field.
func (gu *GroupUpdate) ClearDescription() *GroupUpdate {
	gu.mutation.ClearDescription()
	return gu
}

// SetGravatarLogoURL sets the "gravatar_logo_url" field.
func (gu *GroupUpdate) SetGravatarLogoURL(s string) *GroupUpdate {
	gu.mutation.SetGravatarLogoURL(s)
	return gu
}

// SetNillableGravatarLogoURL sets the "gravatar_logo_url" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableGravatarLogoURL(s *string) *GroupUpdate {
	if s != nil {
		gu.SetGravatarLogoURL(*s)
	}
	return gu
}

// ClearGravatarLogoURL clears the value of the "gravatar_logo_url" field.
func (gu *GroupUpdate) ClearGravatarLogoURL() *GroupUpdate {
	gu.mutation.ClearGravatarLogoURL()
	return gu
}

// SetLogoURL sets the "logo_url" field.
func (gu *GroupUpdate) SetLogoURL(s string) *GroupUpdate {
	gu.mutation.SetLogoURL(s)
	return gu
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableLogoURL(s *string) *GroupUpdate {
	if s != nil {
		gu.SetLogoURL(*s)
	}
	return gu
}

// ClearLogoURL clears the value of the "logo_url" field.
func (gu *GroupUpdate) ClearLogoURL() *GroupUpdate {
	gu.mutation.ClearLogoURL()
	return gu
}

// SetDisplayName sets the "display_name" field.
func (gu *GroupUpdate) SetDisplayName(s string) *GroupUpdate {
	gu.mutation.SetDisplayName(s)
	return gu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDisplayName(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDisplayName(*s)
	}
	return gu
}

// SetOrganizationID sets the "organization_id" field.
func (gu *GroupUpdate) SetOrganizationID(s string) *GroupUpdate {
	gu.mutation.SetOrganizationID(s)
	return gu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableOrganizationID(s *string) *GroupUpdate {
	if s != nil {
		gu.SetOrganizationID(*s)
	}
	return gu
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (gu *GroupUpdate) ClearOrganizationID() *GroupUpdate {
	gu.mutation.ClearOrganizationID()
	return gu
}

// SetSettingID sets the "setting" edge to the GroupSetting entity by ID.
func (gu *GroupUpdate) SetSettingID(id string) *GroupUpdate {
	gu.mutation.SetSettingID(id)
	return gu
}

// SetSetting sets the "setting" edge to the GroupSetting entity.
func (gu *GroupUpdate) SetSetting(g *GroupSetting) *GroupUpdate {
	return gu.SetSettingID(g.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gu *GroupUpdate) AddUserIDs(ids ...string) *GroupUpdate {
	gu.mutation.AddUserIDs(ids...)
	return gu
}

// AddUsers adds the "users" edges to the User entity.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (gu *GroupUpdate) SetOwnerID(id string) *GroupUpdate {
	gu.mutation.SetOwnerID(id)
	return gu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (gu *GroupUpdate) SetOwner(o *Organization) *GroupUpdate {
	return gu.SetOwnerID(o.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearSetting clears the "setting" edge to the GroupSetting entity.
func (gu *GroupUpdate) ClearSetting() *GroupUpdate {
	gu.mutation.ClearSetting()
	return gu
}

// ClearUsers clears all "users" edges to the User entity.
func (gu *GroupUpdate) ClearUsers() *GroupUpdate {
	gu.mutation.ClearUsers()
	return gu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (gu *GroupUpdate) RemoveUserIDs(ids ...string) *GroupUpdate {
	gu.mutation.RemoveUserIDs(ids...)
	return gu
}

// RemoveUsers removes "users" edges to User entities.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (gu *GroupUpdate) ClearOwner() *GroupUpdate {
	gu.mutation.ClearOwner()
	return gu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	if err := gu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gu *GroupUpdate) defaults() error {
	if _, ok := gu.mutation.UpdatedAt(); !ok {
		if group.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := group.UpdateDefaultUpdatedAt()
		gu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := gu.mutation.DisplayName(); ok {
		if err := group.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Group.display_name": %w`, err)}
		}
	}
	if _, ok := gu.mutation.SettingID(); gu.mutation.SettingCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.setting"`)
	}
	if _, ok := gu.mutation.OwnerID(); gu.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.owner"`)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if gu.mutation.CreatedByCleared() {
		_spec.ClearField(group.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gu.mutation.UpdatedBy(); ok {
		_spec.SetField(group.FieldUpdatedBy, field.TypeString, value)
	}
	if gu.mutation.UpdatedByCleared() {
		_spec.ClearField(group.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gu.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
	}
	if gu.mutation.DeletedAtCleared() {
		_spec.ClearField(group.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gu.mutation.DeletedBy(); ok {
		_spec.SetField(group.FieldDeletedBy, field.TypeString, value)
	}
	if gu.mutation.DeletedByCleared() {
		_spec.ClearField(group.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if gu.mutation.DescriptionCleared() {
		_spec.ClearField(group.FieldDescription, field.TypeString)
	}
	if value, ok := gu.mutation.GravatarLogoURL(); ok {
		_spec.SetField(group.FieldGravatarLogoURL, field.TypeString, value)
	}
	if gu.mutation.GravatarLogoURLCleared() {
		_spec.ClearField(group.FieldGravatarLogoURL, field.TypeString)
	}
	if value, ok := gu.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if gu.mutation.LogoURLCleared() {
		_spec.ClearField(group.FieldLogoURL, field.TypeString)
	}
	if value, ok := gu.mutation.DisplayName(); ok {
		_spec.SetField(group.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := gu.mutation.OrganizationID(); ok {
		_spec.SetField(group.FieldOrganizationID, field.TypeString, value)
	}
	if gu.mutation.OrganizationIDCleared() {
		_spec.ClearField(group.FieldOrganizationID, field.TypeString)
	}
	if gu.mutation.SettingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GroupSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.SettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GroupUsers
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !gu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.Group
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = gu.schemaConfig.Group
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = gu.schemaConfig.Group
	ctx = internal.NewSchemaConfigContext(ctx, gu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (guo *GroupUpdateOne) SetUpdatedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetUpdatedAt(t)
	return guo
}

// SetUpdatedBy sets the "updated_by" field.
func (guo *GroupUpdateOne) SetUpdatedBy(s string) *GroupUpdateOne {
	guo.mutation.SetUpdatedBy(s)
	return guo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableUpdatedBy(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetUpdatedBy(*s)
	}
	return guo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (guo *GroupUpdateOne) ClearUpdatedBy() *GroupUpdateOne {
	guo.mutation.ClearUpdatedBy()
	return guo
}

// SetDeletedAt sets the "deleted_at" field.
func (guo *GroupUpdateOne) SetDeletedAt(t time.Time) *GroupUpdateOne {
	guo.mutation.SetDeletedAt(t)
	return guo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupUpdateOne {
	if t != nil {
		guo.SetDeletedAt(*t)
	}
	return guo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (guo *GroupUpdateOne) ClearDeletedAt() *GroupUpdateOne {
	guo.mutation.ClearDeletedAt()
	return guo
}

// SetDeletedBy sets the "deleted_by" field.
func (guo *GroupUpdateOne) SetDeletedBy(s string) *GroupUpdateOne {
	guo.mutation.SetDeletedBy(s)
	return guo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDeletedBy(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDeletedBy(*s)
	}
	return guo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (guo *GroupUpdateOne) ClearDeletedBy() *GroupUpdateOne {
	guo.mutation.ClearDeletedBy()
	return guo
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableName(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetName(*s)
	}
	return guo
}

// SetDescription sets the "description" field.
func (guo *GroupUpdateOne) SetDescription(s string) *GroupUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDescription(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDescription(*s)
	}
	return guo
}

// ClearDescription clears the value of the "description" field.
func (guo *GroupUpdateOne) ClearDescription() *GroupUpdateOne {
	guo.mutation.ClearDescription()
	return guo
}

// SetGravatarLogoURL sets the "gravatar_logo_url" field.
func (guo *GroupUpdateOne) SetGravatarLogoURL(s string) *GroupUpdateOne {
	guo.mutation.SetGravatarLogoURL(s)
	return guo
}

// SetNillableGravatarLogoURL sets the "gravatar_logo_url" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableGravatarLogoURL(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetGravatarLogoURL(*s)
	}
	return guo
}

// ClearGravatarLogoURL clears the value of the "gravatar_logo_url" field.
func (guo *GroupUpdateOne) ClearGravatarLogoURL() *GroupUpdateOne {
	guo.mutation.ClearGravatarLogoURL()
	return guo
}

// SetLogoURL sets the "logo_url" field.
func (guo *GroupUpdateOne) SetLogoURL(s string) *GroupUpdateOne {
	guo.mutation.SetLogoURL(s)
	return guo
}

// SetNillableLogoURL sets the "logo_url" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableLogoURL(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetLogoURL(*s)
	}
	return guo
}

// ClearLogoURL clears the value of the "logo_url" field.
func (guo *GroupUpdateOne) ClearLogoURL() *GroupUpdateOne {
	guo.mutation.ClearLogoURL()
	return guo
}

// SetDisplayName sets the "display_name" field.
func (guo *GroupUpdateOne) SetDisplayName(s string) *GroupUpdateOne {
	guo.mutation.SetDisplayName(s)
	return guo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDisplayName(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDisplayName(*s)
	}
	return guo
}

// SetOrganizationID sets the "organization_id" field.
func (guo *GroupUpdateOne) SetOrganizationID(s string) *GroupUpdateOne {
	guo.mutation.SetOrganizationID(s)
	return guo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableOrganizationID(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetOrganizationID(*s)
	}
	return guo
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (guo *GroupUpdateOne) ClearOrganizationID() *GroupUpdateOne {
	guo.mutation.ClearOrganizationID()
	return guo
}

// SetSettingID sets the "setting" edge to the GroupSetting entity by ID.
func (guo *GroupUpdateOne) SetSettingID(id string) *GroupUpdateOne {
	guo.mutation.SetSettingID(id)
	return guo
}

// SetSetting sets the "setting" edge to the GroupSetting entity.
func (guo *GroupUpdateOne) SetSetting(g *GroupSetting) *GroupUpdateOne {
	return guo.SetSettingID(g.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (guo *GroupUpdateOne) AddUserIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.AddUserIDs(ids...)
	return guo
}

// AddUsers adds the "users" edges to the User entity.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (guo *GroupUpdateOne) SetOwnerID(id string) *GroupUpdateOne {
	guo.mutation.SetOwnerID(id)
	return guo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (guo *GroupUpdateOne) SetOwner(o *Organization) *GroupUpdateOne {
	return guo.SetOwnerID(o.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearSetting clears the "setting" edge to the GroupSetting entity.
func (guo *GroupUpdateOne) ClearSetting() *GroupUpdateOne {
	guo.mutation.ClearSetting()
	return guo
}

// ClearUsers clears all "users" edges to the User entity.
func (guo *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	guo.mutation.ClearUsers()
	return guo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...string) *GroupUpdateOne {
	guo.mutation.RemoveUserIDs(ids...)
	return guo
}

// RemoveUsers removes "users" edges to User entities.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (guo *GroupUpdateOne) ClearOwner() *GroupUpdateOne {
	guo.mutation.ClearOwner()
	return guo
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	if err := guo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (guo *GroupUpdateOne) defaults() error {
	if _, ok := guo.mutation.UpdatedAt(); !ok {
		if group.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized group.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := group.UpdateDefaultUpdatedAt()
		guo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := guo.mutation.DisplayName(); ok {
		if err := group.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Group.display_name": %w`, err)}
		}
	}
	if _, ok := guo.mutation.SettingID(); guo.mutation.SettingCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.setting"`)
	}
	if _, ok := guo.mutation.OwnerID(); guo.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.owner"`)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.UpdatedAt(); ok {
		_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
	}
	if guo.mutation.CreatedByCleared() {
		_spec.ClearField(group.FieldCreatedBy, field.TypeString)
	}
	if value, ok := guo.mutation.UpdatedBy(); ok {
		_spec.SetField(group.FieldUpdatedBy, field.TypeString, value)
	}
	if guo.mutation.UpdatedByCleared() {
		_spec.ClearField(group.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := guo.mutation.DeletedAt(); ok {
		_spec.SetField(group.FieldDeletedAt, field.TypeTime, value)
	}
	if guo.mutation.DeletedAtCleared() {
		_spec.ClearField(group.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := guo.mutation.DeletedBy(); ok {
		_spec.SetField(group.FieldDeletedBy, field.TypeString, value)
	}
	if guo.mutation.DeletedByCleared() {
		_spec.ClearField(group.FieldDeletedBy, field.TypeString)
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if guo.mutation.DescriptionCleared() {
		_spec.ClearField(group.FieldDescription, field.TypeString)
	}
	if value, ok := guo.mutation.GravatarLogoURL(); ok {
		_spec.SetField(group.FieldGravatarLogoURL, field.TypeString, value)
	}
	if guo.mutation.GravatarLogoURLCleared() {
		_spec.ClearField(group.FieldGravatarLogoURL, field.TypeString)
	}
	if value, ok := guo.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if guo.mutation.LogoURLCleared() {
		_spec.ClearField(group.FieldLogoURL, field.TypeString)
	}
	if value, ok := guo.mutation.DisplayName(); ok {
		_spec.SetField(group.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := guo.mutation.OrganizationID(); ok {
		_spec.SetField(group.FieldOrganizationID, field.TypeString, value)
	}
	if guo.mutation.OrganizationIDCleared() {
		_spec.ClearField(group.FieldOrganizationID, field.TypeString)
	}
	if guo.mutation.SettingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GroupSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.SettingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GroupUsers
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !guo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.Group
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   group.OwnerTable,
			Columns: []string{group.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = guo.schemaConfig.Group
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = guo.schemaConfig.Group
	ctx = internal.NewSchemaConfigContext(ctx, guo.schemaConfig)
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}
