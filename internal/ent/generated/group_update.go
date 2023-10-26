// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/google/uuid"
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

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
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

// SetLogoURL sets the "logo_url" field.
func (gu *GroupUpdate) SetLogoURL(s string) *GroupUpdate {
	gu.mutation.SetLogoURL(s)
	return gu
}

// SetSettingID sets the "setting" edge to the GroupSettings entity by ID.
func (gu *GroupUpdate) SetSettingID(id uuid.UUID) *GroupUpdate {
	gu.mutation.SetSettingID(id)
	return gu
}

// SetSetting sets the "setting" edge to the GroupSettings entity.
func (gu *GroupUpdate) SetSetting(g *GroupSettings) *GroupUpdate {
	return gu.SetSettingID(g.ID)
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (gu *GroupUpdate) AddMembershipIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.AddMembershipIDs(ids...)
	return gu
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (gu *GroupUpdate) AddMemberships(m ...*Membership) *GroupUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.AddMembershipIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearSetting clears the "setting" edge to the GroupSettings entity.
func (gu *GroupUpdate) ClearSetting() *GroupUpdate {
	gu.mutation.ClearSetting()
	return gu
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (gu *GroupUpdate) ClearMemberships() *GroupUpdate {
	gu.mutation.ClearMemberships()
	return gu
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (gu *GroupUpdate) RemoveMembershipIDs(ids ...uuid.UUID) *GroupUpdate {
	gu.mutation.RemoveMembershipIDs(ids...)
	return gu
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (gu *GroupUpdate) RemoveMemberships(m ...*Membership) *GroupUpdate {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gu.RemoveMembershipIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
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

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := gu.mutation.LogoURL(); ok {
		if err := group.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`generated: validator failed for field "Group.logo_url": %w`, err)}
		}
	}
	if _, ok := gu.mutation.SettingID(); gu.mutation.SettingCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.setting"`)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := gu.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if gu.mutation.SettingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID),
			},
		}
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
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !gu.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
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

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
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

// SetLogoURL sets the "logo_url" field.
func (guo *GroupUpdateOne) SetLogoURL(s string) *GroupUpdateOne {
	guo.mutation.SetLogoURL(s)
	return guo
}

// SetSettingID sets the "setting" edge to the GroupSettings entity by ID.
func (guo *GroupUpdateOne) SetSettingID(id uuid.UUID) *GroupUpdateOne {
	guo.mutation.SetSettingID(id)
	return guo
}

// SetSetting sets the "setting" edge to the GroupSettings entity.
func (guo *GroupUpdateOne) SetSetting(g *GroupSettings) *GroupUpdateOne {
	return guo.SetSettingID(g.ID)
}

// AddMembershipIDs adds the "memberships" edge to the Membership entity by IDs.
func (guo *GroupUpdateOne) AddMembershipIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.AddMembershipIDs(ids...)
	return guo
}

// AddMemberships adds the "memberships" edges to the Membership entity.
func (guo *GroupUpdateOne) AddMemberships(m ...*Membership) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.AddMembershipIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearSetting clears the "setting" edge to the GroupSettings entity.
func (guo *GroupUpdateOne) ClearSetting() *GroupUpdateOne {
	guo.mutation.ClearSetting()
	return guo
}

// ClearMemberships clears all "memberships" edges to the Membership entity.
func (guo *GroupUpdateOne) ClearMemberships() *GroupUpdateOne {
	guo.mutation.ClearMemberships()
	return guo
}

// RemoveMembershipIDs removes the "memberships" edge to Membership entities by IDs.
func (guo *GroupUpdateOne) RemoveMembershipIDs(ids ...uuid.UUID) *GroupUpdateOne {
	guo.mutation.RemoveMembershipIDs(ids...)
	return guo
}

// RemoveMemberships removes "memberships" edges to Membership entities.
func (guo *GroupUpdateOne) RemoveMemberships(m ...*Membership) *GroupUpdateOne {
	ids := make([]uuid.UUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return guo.RemoveMembershipIDs(ids...)
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

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := guo.mutation.LogoURL(); ok {
		if err := group.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`generated: validator failed for field "Group.logo_url": %w`, err)}
		}
	}
	if _, ok := guo.mutation.SettingID(); guo.mutation.SettingCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Group.setting"`)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID))
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
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := guo.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if guo.mutation.SettingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingTable,
			Columns: []string{group.SettingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID),
			},
		}
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
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMembershipsIDs(); len(nodes) > 0 && !guo.mutation.MembershipsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MembershipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembershipsTable,
			Columns: []string{group.MembershipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
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
