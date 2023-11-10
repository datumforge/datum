// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
)

// GroupSettingsCreate is the builder for creating a GroupSettings entity.
type GroupSettingsCreate struct {
	config
	mutation *GroupSettingsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gsc *GroupSettingsCreate) SetCreatedAt(t time.Time) *GroupSettingsCreate {
	gsc.mutation.SetCreatedAt(t)
	return gsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableCreatedAt(t *time.Time) *GroupSettingsCreate {
	if t != nil {
		gsc.SetCreatedAt(*t)
	}
	return gsc
}

// SetUpdatedAt sets the "updated_at" field.
func (gsc *GroupSettingsCreate) SetUpdatedAt(t time.Time) *GroupSettingsCreate {
	gsc.mutation.SetUpdatedAt(t)
	return gsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableUpdatedAt(t *time.Time) *GroupSettingsCreate {
	if t != nil {
		gsc.SetUpdatedAt(*t)
	}
	return gsc
}

// SetCreatedBy sets the "created_by" field.
func (gsc *GroupSettingsCreate) SetCreatedBy(s string) *GroupSettingsCreate {
	gsc.mutation.SetCreatedBy(s)
	return gsc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableCreatedBy(s *string) *GroupSettingsCreate {
	if s != nil {
		gsc.SetCreatedBy(*s)
	}
	return gsc
}

// SetUpdatedBy sets the "updated_by" field.
func (gsc *GroupSettingsCreate) SetUpdatedBy(s string) *GroupSettingsCreate {
	gsc.mutation.SetUpdatedBy(s)
	return gsc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableUpdatedBy(s *string) *GroupSettingsCreate {
	if s != nil {
		gsc.SetUpdatedBy(*s)
	}
	return gsc
}

// SetVisibility sets the "visibility" field.
func (gsc *GroupSettingsCreate) SetVisibility(gr groupsettings.Visibility) *GroupSettingsCreate {
	gsc.mutation.SetVisibility(gr)
	return gsc
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableVisibility(gr *groupsettings.Visibility) *GroupSettingsCreate {
	if gr != nil {
		gsc.SetVisibility(*gr)
	}
	return gsc
}

// SetJoinPolicy sets the "join_policy" field.
func (gsc *GroupSettingsCreate) SetJoinPolicy(gp groupsettings.JoinPolicy) *GroupSettingsCreate {
	gsc.mutation.SetJoinPolicy(gp)
	return gsc
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableJoinPolicy(gp *groupsettings.JoinPolicy) *GroupSettingsCreate {
	if gp != nil {
		gsc.SetJoinPolicy(*gp)
	}
	return gsc
}

// SetTags sets the "tags" field.
func (gsc *GroupSettingsCreate) SetTags(s []string) *GroupSettingsCreate {
	gsc.mutation.SetTags(s)
	return gsc
}

// SetID sets the "id" field.
func (gsc *GroupSettingsCreate) SetID(s string) *GroupSettingsCreate {
	gsc.mutation.SetID(s)
	return gsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableID(s *string) *GroupSettingsCreate {
	if s != nil {
		gsc.SetID(*s)
	}
	return gsc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gsc *GroupSettingsCreate) SetGroupID(id string) *GroupSettingsCreate {
	gsc.mutation.SetGroupID(id)
	return gsc
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gsc *GroupSettingsCreate) SetNillableGroupID(id *string) *GroupSettingsCreate {
	if id != nil {
		gsc = gsc.SetGroupID(*id)
	}
	return gsc
}

// SetGroup sets the "group" edge to the Group entity.
func (gsc *GroupSettingsCreate) SetGroup(g *Group) *GroupSettingsCreate {
	return gsc.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingsMutation object of the builder.
func (gsc *GroupSettingsCreate) Mutation() *GroupSettingsMutation {
	return gsc.mutation
}

// Save creates the GroupSettings in the database.
func (gsc *GroupSettingsCreate) Save(ctx context.Context) (*GroupSettings, error) {
	if err := gsc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gsc.sqlSave, gsc.mutation, gsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gsc *GroupSettingsCreate) SaveX(ctx context.Context) *GroupSettings {
	v, err := gsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gsc *GroupSettingsCreate) Exec(ctx context.Context) error {
	_, err := gsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsc *GroupSettingsCreate) ExecX(ctx context.Context) {
	if err := gsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsc *GroupSettingsCreate) defaults() error {
	if _, ok := gsc.mutation.CreatedAt(); !ok {
		if groupsettings.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsettings.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := groupsettings.DefaultCreatedAt()
		gsc.mutation.SetCreatedAt(v)
	}
	if _, ok := gsc.mutation.UpdatedAt(); !ok {
		if groupsettings.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsettings.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsettings.DefaultUpdatedAt()
		gsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gsc.mutation.Visibility(); !ok {
		v := groupsettings.DefaultVisibility
		gsc.mutation.SetVisibility(v)
	}
	if _, ok := gsc.mutation.JoinPolicy(); !ok {
		v := groupsettings.DefaultJoinPolicy
		gsc.mutation.SetJoinPolicy(v)
	}
	if _, ok := gsc.mutation.Tags(); !ok {
		v := groupsettings.DefaultTags
		gsc.mutation.SetTags(v)
	}
	if _, ok := gsc.mutation.ID(); !ok {
		if groupsettings.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized groupsettings.DefaultID (forgotten import generated/runtime?)")
		}
		v := groupsettings.DefaultID()
		gsc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsc *GroupSettingsCreate) check() error {
	if _, ok := gsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "GroupSettings.created_at"`)}
	}
	if _, ok := gsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "GroupSettings.updated_at"`)}
	}
	if _, ok := gsc.mutation.Visibility(); !ok {
		return &ValidationError{Name: "visibility", err: errors.New(`generated: missing required field "GroupSettings.visibility"`)}
	}
	if v, ok := gsc.mutation.Visibility(); ok {
		if err := groupsettings.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.visibility": %w`, err)}
		}
	}
	if _, ok := gsc.mutation.JoinPolicy(); !ok {
		return &ValidationError{Name: "join_policy", err: errors.New(`generated: missing required field "GroupSettings.join_policy"`)}
	}
	if v, ok := gsc.mutation.JoinPolicy(); ok {
		if err := groupsettings.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.join_policy": %w`, err)}
		}
	}
	if _, ok := gsc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`generated: missing required field "GroupSettings.tags"`)}
	}
	return nil
}

func (gsc *GroupSettingsCreate) sqlSave(ctx context.Context) (*GroupSettings, error) {
	if err := gsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected GroupSettings.ID type: %T", _spec.ID.Value)
		}
	}
	gsc.mutation.id = &_node.ID
	gsc.mutation.done = true
	return _node, nil
}

func (gsc *GroupSettingsCreate) createSpec() (*GroupSettings, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupSettings{config: gsc.config}
		_spec = sqlgraph.NewCreateSpec(groupsettings.Table, sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString))
	)
	_spec.Schema = gsc.schemaConfig.GroupSettings
	if id, ok := gsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gsc.mutation.CreatedAt(); ok {
		_spec.SetField(groupsettings.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gsc.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsettings.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gsc.mutation.CreatedBy(); ok {
		_spec.SetField(groupsettings.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gsc.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsettings.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gsc.mutation.Visibility(); ok {
		_spec.SetField(groupsettings.FieldVisibility, field.TypeEnum, value)
		_node.Visibility = value
	}
	if value, ok := gsc.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsettings.FieldJoinPolicy, field.TypeEnum, value)
		_node.JoinPolicy = value
	}
	if value, ok := gsc.mutation.Tags(); ok {
		_spec.SetField(groupsettings.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if nodes := gsc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsettings.GroupTable,
			Columns: []string{groupsettings.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsc.schemaConfig.GroupSettings
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_setting = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupSettingsCreateBulk is the builder for creating many GroupSettings entities in bulk.
type GroupSettingsCreateBulk struct {
	config
	err      error
	builders []*GroupSettingsCreate
}

// Save creates the GroupSettings entities in the database.
func (gscb *GroupSettingsCreateBulk) Save(ctx context.Context) ([]*GroupSettings, error) {
	if gscb.err != nil {
		return nil, gscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gscb.builders))
	nodes := make([]*GroupSettings, len(gscb.builders))
	mutators := make([]Mutator, len(gscb.builders))
	for i := range gscb.builders {
		func(i int, root context.Context) {
			builder := gscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupSettingsMutation)
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
					_, err = mutators[i+1].Mutate(root, gscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gscb *GroupSettingsCreateBulk) SaveX(ctx context.Context) []*GroupSettings {
	v, err := gscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gscb *GroupSettingsCreateBulk) Exec(ctx context.Context) error {
	_, err := gscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gscb *GroupSettingsCreateBulk) ExecX(ctx context.Context) {
	if err := gscb.Exec(ctx); err != nil {
		panic(err)
	}
}
