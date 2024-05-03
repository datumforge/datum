// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsetting"
)

// GroupSettingCreate is the builder for creating a GroupSetting entity.
type GroupSettingCreate struct {
	config
	mutation *GroupSettingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gsc *GroupSettingCreate) SetCreatedAt(t time.Time) *GroupSettingCreate {
	gsc.mutation.SetCreatedAt(t)
	return gsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableCreatedAt(t *time.Time) *GroupSettingCreate {
	if t != nil {
		gsc.SetCreatedAt(*t)
	}
	return gsc
}

// SetUpdatedAt sets the "updated_at" field.
func (gsc *GroupSettingCreate) SetUpdatedAt(t time.Time) *GroupSettingCreate {
	gsc.mutation.SetUpdatedAt(t)
	return gsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableUpdatedAt(t *time.Time) *GroupSettingCreate {
	if t != nil {
		gsc.SetUpdatedAt(*t)
	}
	return gsc
}

// SetCreatedBy sets the "created_by" field.
func (gsc *GroupSettingCreate) SetCreatedBy(s string) *GroupSettingCreate {
	gsc.mutation.SetCreatedBy(s)
	return gsc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableCreatedBy(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetCreatedBy(*s)
	}
	return gsc
}

// SetUpdatedBy sets the "updated_by" field.
func (gsc *GroupSettingCreate) SetUpdatedBy(s string) *GroupSettingCreate {
	gsc.mutation.SetUpdatedBy(s)
	return gsc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableUpdatedBy(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetUpdatedBy(*s)
	}
	return gsc
}

// SetMappingID sets the "mapping_id" field.
func (gsc *GroupSettingCreate) SetMappingID(s string) *GroupSettingCreate {
	gsc.mutation.SetMappingID(s)
	return gsc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableMappingID(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetMappingID(*s)
	}
	return gsc
}

// SetDeletedAt sets the "deleted_at" field.
func (gsc *GroupSettingCreate) SetDeletedAt(t time.Time) *GroupSettingCreate {
	gsc.mutation.SetDeletedAt(t)
	return gsc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableDeletedAt(t *time.Time) *GroupSettingCreate {
	if t != nil {
		gsc.SetDeletedAt(*t)
	}
	return gsc
}

// SetDeletedBy sets the "deleted_by" field.
func (gsc *GroupSettingCreate) SetDeletedBy(s string) *GroupSettingCreate {
	gsc.mutation.SetDeletedBy(s)
	return gsc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableDeletedBy(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetDeletedBy(*s)
	}
	return gsc
}

// SetVisibility sets the "visibility" field.
func (gsc *GroupSettingCreate) SetVisibility(e enums.Visibility) *GroupSettingCreate {
	gsc.mutation.SetVisibility(e)
	return gsc
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableVisibility(e *enums.Visibility) *GroupSettingCreate {
	if e != nil {
		gsc.SetVisibility(*e)
	}
	return gsc
}

// SetJoinPolicy sets the "join_policy" field.
func (gsc *GroupSettingCreate) SetJoinPolicy(ep enums.JoinPolicy) *GroupSettingCreate {
	gsc.mutation.SetJoinPolicy(ep)
	return gsc
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableJoinPolicy(ep *enums.JoinPolicy) *GroupSettingCreate {
	if ep != nil {
		gsc.SetJoinPolicy(*ep)
	}
	return gsc
}

// SetTags sets the "tags" field.
func (gsc *GroupSettingCreate) SetTags(s []string) *GroupSettingCreate {
	gsc.mutation.SetTags(s)
	return gsc
}

// SetSyncToSlack sets the "sync_to_slack" field.
func (gsc *GroupSettingCreate) SetSyncToSlack(b bool) *GroupSettingCreate {
	gsc.mutation.SetSyncToSlack(b)
	return gsc
}

// SetNillableSyncToSlack sets the "sync_to_slack" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableSyncToSlack(b *bool) *GroupSettingCreate {
	if b != nil {
		gsc.SetSyncToSlack(*b)
	}
	return gsc
}

// SetSyncToGithub sets the "sync_to_github" field.
func (gsc *GroupSettingCreate) SetSyncToGithub(b bool) *GroupSettingCreate {
	gsc.mutation.SetSyncToGithub(b)
	return gsc
}

// SetNillableSyncToGithub sets the "sync_to_github" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableSyncToGithub(b *bool) *GroupSettingCreate {
	if b != nil {
		gsc.SetSyncToGithub(*b)
	}
	return gsc
}

// SetGroupID sets the "group_id" field.
func (gsc *GroupSettingCreate) SetGroupID(s string) *GroupSettingCreate {
	gsc.mutation.SetGroupID(s)
	return gsc
}

// SetNillableGroupID sets the "group_id" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableGroupID(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetGroupID(*s)
	}
	return gsc
}

// SetID sets the "id" field.
func (gsc *GroupSettingCreate) SetID(s string) *GroupSettingCreate {
	gsc.mutation.SetID(s)
	return gsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gsc *GroupSettingCreate) SetNillableID(s *string) *GroupSettingCreate {
	if s != nil {
		gsc.SetID(*s)
	}
	return gsc
}

// SetGroup sets the "group" edge to the Group entity.
func (gsc *GroupSettingCreate) SetGroup(g *Group) *GroupSettingCreate {
	return gsc.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingMutation object of the builder.
func (gsc *GroupSettingCreate) Mutation() *GroupSettingMutation {
	return gsc.mutation
}

// Save creates the GroupSetting in the database.
func (gsc *GroupSettingCreate) Save(ctx context.Context) (*GroupSetting, error) {
	if err := gsc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gsc.sqlSave, gsc.mutation, gsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gsc *GroupSettingCreate) SaveX(ctx context.Context) *GroupSetting {
	v, err := gsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gsc *GroupSettingCreate) Exec(ctx context.Context) error {
	_, err := gsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsc *GroupSettingCreate) ExecX(ctx context.Context) {
	if err := gsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsc *GroupSettingCreate) defaults() error {
	if _, ok := gsc.mutation.CreatedAt(); !ok {
		if groupsetting.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := groupsetting.DefaultCreatedAt()
		gsc.mutation.SetCreatedAt(v)
	}
	if _, ok := gsc.mutation.UpdatedAt(); !ok {
		if groupsetting.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsetting.DefaultUpdatedAt()
		gsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gsc.mutation.MappingID(); !ok {
		if groupsetting.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := groupsetting.DefaultMappingID()
		gsc.mutation.SetMappingID(v)
	}
	if _, ok := gsc.mutation.Visibility(); !ok {
		v := groupsetting.DefaultVisibility
		gsc.mutation.SetVisibility(v)
	}
	if _, ok := gsc.mutation.JoinPolicy(); !ok {
		v := groupsetting.DefaultJoinPolicy
		gsc.mutation.SetJoinPolicy(v)
	}
	if _, ok := gsc.mutation.Tags(); !ok {
		v := groupsetting.DefaultTags
		gsc.mutation.SetTags(v)
	}
	if _, ok := gsc.mutation.SyncToSlack(); !ok {
		v := groupsetting.DefaultSyncToSlack
		gsc.mutation.SetSyncToSlack(v)
	}
	if _, ok := gsc.mutation.SyncToGithub(); !ok {
		v := groupsetting.DefaultSyncToGithub
		gsc.mutation.SetSyncToGithub(v)
	}
	if _, ok := gsc.mutation.ID(); !ok {
		if groupsetting.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized groupsetting.DefaultID (forgotten import generated/runtime?)")
		}
		v := groupsetting.DefaultID()
		gsc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsc *GroupSettingCreate) check() error {
	if _, ok := gsc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "GroupSetting.mapping_id"`)}
	}
	if _, ok := gsc.mutation.Visibility(); !ok {
		return &ValidationError{Name: "visibility", err: errors.New(`generated: missing required field "GroupSetting.visibility"`)}
	}
	if v, ok := gsc.mutation.Visibility(); ok {
		if err := groupsetting.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.visibility": %w`, err)}
		}
	}
	if _, ok := gsc.mutation.JoinPolicy(); !ok {
		return &ValidationError{Name: "join_policy", err: errors.New(`generated: missing required field "GroupSetting.join_policy"`)}
	}
	if v, ok := gsc.mutation.JoinPolicy(); ok {
		if err := groupsetting.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSetting.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gsc *GroupSettingCreate) sqlSave(ctx context.Context) (*GroupSetting, error) {
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
			return nil, fmt.Errorf("unexpected GroupSetting.ID type: %T", _spec.ID.Value)
		}
	}
	gsc.mutation.id = &_node.ID
	gsc.mutation.done = true
	return _node, nil
}

func (gsc *GroupSettingCreate) createSpec() (*GroupSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupSetting{config: gsc.config}
		_spec = sqlgraph.NewCreateSpec(groupsetting.Table, sqlgraph.NewFieldSpec(groupsetting.FieldID, field.TypeString))
	)
	_spec.Schema = gsc.schemaConfig.GroupSetting
	if id, ok := gsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gsc.mutation.CreatedAt(); ok {
		_spec.SetField(groupsetting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gsc.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsetting.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gsc.mutation.CreatedBy(); ok {
		_spec.SetField(groupsetting.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gsc.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsetting.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gsc.mutation.MappingID(); ok {
		_spec.SetField(groupsetting.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := gsc.mutation.DeletedAt(); ok {
		_spec.SetField(groupsetting.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gsc.mutation.DeletedBy(); ok {
		_spec.SetField(groupsetting.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := gsc.mutation.Visibility(); ok {
		_spec.SetField(groupsetting.FieldVisibility, field.TypeEnum, value)
		_node.Visibility = value
	}
	if value, ok := gsc.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsetting.FieldJoinPolicy, field.TypeEnum, value)
		_node.JoinPolicy = value
	}
	if value, ok := gsc.mutation.Tags(); ok {
		_spec.SetField(groupsetting.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := gsc.mutation.SyncToSlack(); ok {
		_spec.SetField(groupsetting.FieldSyncToSlack, field.TypeBool, value)
		_node.SyncToSlack = value
	}
	if value, ok := gsc.mutation.SyncToGithub(); ok {
		_spec.SetField(groupsetting.FieldSyncToGithub, field.TypeBool, value)
		_node.SyncToGithub = value
	}
	if nodes := gsc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsetting.GroupTable,
			Columns: []string{groupsetting.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gsc.schemaConfig.GroupSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupSettingCreateBulk is the builder for creating many GroupSetting entities in bulk.
type GroupSettingCreateBulk struct {
	config
	err      error
	builders []*GroupSettingCreate
}

// Save creates the GroupSetting entities in the database.
func (gscb *GroupSettingCreateBulk) Save(ctx context.Context) ([]*GroupSetting, error) {
	if gscb.err != nil {
		return nil, gscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gscb.builders))
	nodes := make([]*GroupSetting, len(gscb.builders))
	mutators := make([]Mutator, len(gscb.builders))
	for i := range gscb.builders {
		func(i int, root context.Context) {
			builder := gscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupSettingMutation)
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
func (gscb *GroupSettingCreateBulk) SaveX(ctx context.Context) []*GroupSetting {
	v, err := gscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gscb *GroupSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := gscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gscb *GroupSettingCreateBulk) ExecX(ctx context.Context) {
	if err := gscb.Exec(ctx); err != nil {
		panic(err)
	}
}
