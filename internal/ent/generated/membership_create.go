// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/membership"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/google/uuid"
)

// MembershipCreate is the builder for creating a Membership entity.
type MembershipCreate struct {
	config
	mutation *MembershipMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MembershipCreate) SetCreatedAt(t time.Time) *MembershipCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableCreatedAt(t *time.Time) *MembershipCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MembershipCreate) SetUpdatedAt(t time.Time) *MembershipCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableUpdatedAt(t *time.Time) *MembershipCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetCreatedBy sets the "created_by" field.
func (mc *MembershipCreate) SetCreatedBy(i int) *MembershipCreate {
	mc.mutation.SetCreatedBy(i)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableCreatedBy(i *int) *MembershipCreate {
	if i != nil {
		mc.SetCreatedBy(*i)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *MembershipCreate) SetUpdatedBy(i int) *MembershipCreate {
	mc.mutation.SetUpdatedBy(i)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableUpdatedBy(i *int) *MembershipCreate {
	if i != nil {
		mc.SetUpdatedBy(*i)
	}
	return mc
}

// SetCurrent sets the "current" field.
func (mc *MembershipCreate) SetCurrent(b bool) *MembershipCreate {
	mc.mutation.SetCurrent(b)
	return mc
}

// SetNillableCurrent sets the "current" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableCurrent(b *bool) *MembershipCreate {
	if b != nil {
		mc.SetCurrent(*b)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MembershipCreate) SetID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableID(u *uuid.UUID) *MembershipCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (mc *MembershipCreate) SetOrganizationID(id uuid.UUID) *MembershipCreate {
	mc.mutation.SetOrganizationID(id)
	return mc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (mc *MembershipCreate) SetOrganization(o *Organization) *MembershipCreate {
	return mc.SetOrganizationID(o.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (mc *MembershipCreate) SetUserID(id uuid.UUID) *MembershipCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetUser sets the "user" edge to the User entity.
func (mc *MembershipCreate) SetUser(u *User) *MembershipCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mc *MembershipCreate) Mutation() *MembershipMutation {
	return mc.mutation
}

// Save creates the Membership in the database.
func (mc *MembershipCreate) Save(ctx context.Context) (*Membership, error) {
	if err := mc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MembershipCreate) SaveX(ctx context.Context) *Membership {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MembershipCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MembershipCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MembershipCreate) defaults() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		if membership.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized membership.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := membership.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		if membership.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized membership.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := membership.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Current(); !ok {
		v := membership.DefaultCurrent
		mc.mutation.SetCurrent(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		if membership.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized membership.DefaultID (forgotten import generated/runtime?)")
		}
		v := membership.DefaultID()
		mc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (mc *MembershipCreate) check() error {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "Membership.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "Membership.updated_at"`)}
	}
	if _, ok := mc.mutation.Current(); !ok {
		return &ValidationError{Name: "current", err: errors.New(`generated: missing required field "Membership.current"`)}
	}
	if _, ok := mc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New(`generated: missing required edge "Membership.organization"`)}
	}
	if _, ok := mc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`generated: missing required edge "Membership.user"`)}
	}
	return nil
}

func (mc *MembershipCreate) sqlSave(ctx context.Context) (*Membership, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MembershipCreate) createSpec() (*Membership, *sqlgraph.CreateSpec) {
	var (
		_node = &Membership{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(membership.Table, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(membership.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(membership.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(membership.FieldCreatedBy, field.TypeInt, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(membership.FieldUpdatedBy, field.TypeInt, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.Current(); ok {
		_spec.SetField(membership.FieldCurrent, field.TypeBool, value)
		_node.Current = value
	}
	if nodes := mc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.OrganizationTable,
			Columns: []string{membership.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.organization_memberships = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.UserTable,
			Columns: []string{membership.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_memberships = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MembershipCreateBulk is the builder for creating many Membership entities in bulk.
type MembershipCreateBulk struct {
	config
	err      error
	builders []*MembershipCreate
}

// Save creates the Membership entities in the database.
func (mcb *MembershipCreateBulk) Save(ctx context.Context) ([]*Membership, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Membership, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MembershipMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MembershipCreateBulk) SaveX(ctx context.Context) []*Membership {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MembershipCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
