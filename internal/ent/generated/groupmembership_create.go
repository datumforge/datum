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
	"github.com/datumforge/datum/internal/ent/generated/groupmembership"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// GroupMembershipCreate is the builder for creating a GroupMembership entity.
type GroupMembershipCreate struct {
	config
	mutation *GroupMembershipMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (gmc *GroupMembershipCreate) SetCreatedAt(t time.Time) *GroupMembershipCreate {
	gmc.mutation.SetCreatedAt(t)
	return gmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableCreatedAt(t *time.Time) *GroupMembershipCreate {
	if t != nil {
		gmc.SetCreatedAt(*t)
	}
	return gmc
}

// SetUpdatedAt sets the "updated_at" field.
func (gmc *GroupMembershipCreate) SetUpdatedAt(t time.Time) *GroupMembershipCreate {
	gmc.mutation.SetUpdatedAt(t)
	return gmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableUpdatedAt(t *time.Time) *GroupMembershipCreate {
	if t != nil {
		gmc.SetUpdatedAt(*t)
	}
	return gmc
}

// SetCreatedBy sets the "created_by" field.
func (gmc *GroupMembershipCreate) SetCreatedBy(s string) *GroupMembershipCreate {
	gmc.mutation.SetCreatedBy(s)
	return gmc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableCreatedBy(s *string) *GroupMembershipCreate {
	if s != nil {
		gmc.SetCreatedBy(*s)
	}
	return gmc
}

// SetUpdatedBy sets the "updated_by" field.
func (gmc *GroupMembershipCreate) SetUpdatedBy(s string) *GroupMembershipCreate {
	gmc.mutation.SetUpdatedBy(s)
	return gmc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableUpdatedBy(s *string) *GroupMembershipCreate {
	if s != nil {
		gmc.SetUpdatedBy(*s)
	}
	return gmc
}

// SetDeletedAt sets the "deleted_at" field.
func (gmc *GroupMembershipCreate) SetDeletedAt(t time.Time) *GroupMembershipCreate {
	gmc.mutation.SetDeletedAt(t)
	return gmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableDeletedAt(t *time.Time) *GroupMembershipCreate {
	if t != nil {
		gmc.SetDeletedAt(*t)
	}
	return gmc
}

// SetDeletedBy sets the "deleted_by" field.
func (gmc *GroupMembershipCreate) SetDeletedBy(s string) *GroupMembershipCreate {
	gmc.mutation.SetDeletedBy(s)
	return gmc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableDeletedBy(s *string) *GroupMembershipCreate {
	if s != nil {
		gmc.SetDeletedBy(*s)
	}
	return gmc
}

// SetRole sets the "role" field.
func (gmc *GroupMembershipCreate) SetRole(e enums.Role) *GroupMembershipCreate {
	gmc.mutation.SetRole(e)
	return gmc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableRole(e *enums.Role) *GroupMembershipCreate {
	if e != nil {
		gmc.SetRole(*e)
	}
	return gmc
}

// SetGroupID sets the "group_id" field.
func (gmc *GroupMembershipCreate) SetGroupID(s string) *GroupMembershipCreate {
	gmc.mutation.SetGroupID(s)
	return gmc
}

// SetUserID sets the "user_id" field.
func (gmc *GroupMembershipCreate) SetUserID(s string) *GroupMembershipCreate {
	gmc.mutation.SetUserID(s)
	return gmc
}

// SetID sets the "id" field.
func (gmc *GroupMembershipCreate) SetID(s string) *GroupMembershipCreate {
	gmc.mutation.SetID(s)
	return gmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableID(s *string) *GroupMembershipCreate {
	if s != nil {
		gmc.SetID(*s)
	}
	return gmc
}

// SetGroup sets the "group" edge to the Group entity.
func (gmc *GroupMembershipCreate) SetGroup(g *Group) *GroupMembershipCreate {
	return gmc.SetGroupID(g.ID)
}

// SetUser sets the "user" edge to the User entity.
func (gmc *GroupMembershipCreate) SetUser(u *User) *GroupMembershipCreate {
	return gmc.SetUserID(u.ID)
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmc *GroupMembershipCreate) Mutation() *GroupMembershipMutation {
	return gmc.mutation
}

// Save creates the GroupMembership in the database.
func (gmc *GroupMembershipCreate) Save(ctx context.Context) (*GroupMembership, error) {
	if err := gmc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gmc.sqlSave, gmc.mutation, gmc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMembershipCreate) SaveX(ctx context.Context) *GroupMembership {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMembershipCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMembershipCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmc *GroupMembershipCreate) defaults() error {
	if _, ok := gmc.mutation.CreatedAt(); !ok {
		if groupmembership.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupmembership.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := groupmembership.DefaultCreatedAt()
		gmc.mutation.SetCreatedAt(v)
	}
	if _, ok := gmc.mutation.UpdatedAt(); !ok {
		if groupmembership.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupmembership.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupmembership.DefaultUpdatedAt()
		gmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gmc.mutation.Role(); !ok {
		v := groupmembership.DefaultRole
		gmc.mutation.SetRole(v)
	}
	if _, ok := gmc.mutation.ID(); !ok {
		if groupmembership.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized groupmembership.DefaultID (forgotten import generated/runtime?)")
		}
		v := groupmembership.DefaultID()
		gmc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMembershipCreate) check() error {
	if _, ok := gmc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`generated: missing required field "GroupMembership.role"`)}
	}
	if v, ok := gmc.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`generated: missing required field "GroupMembership.group_id"`)}
	}
	if _, ok := gmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`generated: missing required field "GroupMembership.user_id"`)}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`generated: missing required edge "GroupMembership.group"`)}
	}
	if _, ok := gmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`generated: missing required edge "GroupMembership.user"`)}
	}
	return nil
}

func (gmc *GroupMembershipCreate) sqlSave(ctx context.Context) (*GroupMembership, error) {
	if err := gmc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected GroupMembership.ID type: %T", _spec.ID.Value)
		}
	}
	gmc.mutation.id = &_node.ID
	gmc.mutation.done = true
	return _node, nil
}

func (gmc *GroupMembershipCreate) createSpec() (*GroupMembership, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMembership{config: gmc.config}
		_spec = sqlgraph.NewCreateSpec(groupmembership.Table, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	)
	_spec.Schema = gmc.schemaConfig.GroupMembership
	if id, ok := gmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gmc.mutation.CreatedAt(); ok {
		_spec.SetField(groupmembership.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := gmc.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmembership.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := gmc.mutation.CreatedBy(); ok {
		_spec.SetField(groupmembership.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := gmc.mutation.UpdatedBy(); ok {
		_spec.SetField(groupmembership.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := gmc.mutation.DeletedAt(); ok {
		_spec.SetField(groupmembership.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := gmc.mutation.DeletedBy(); ok {
		_spec.SetField(groupmembership.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := gmc.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := gmc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = gmc.schemaConfig.GroupMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = gmc.schemaConfig.GroupMembership
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupMembershipCreateBulk is the builder for creating many GroupMembership entities in bulk.
type GroupMembershipCreateBulk struct {
	config
	err      error
	builders []*GroupMembershipCreate
}

// Save creates the GroupMembership entities in the database.
func (gmcb *GroupMembershipCreateBulk) Save(ctx context.Context) ([]*GroupMembership, error) {
	if gmcb.err != nil {
		return nil, gmcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMembership, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMembershipMutation)
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
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) SaveX(ctx context.Context) []*GroupMembership {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}
