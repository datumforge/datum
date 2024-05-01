// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/event"
	"github.com/datumforge/datum/internal/ent/generated/hush"
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// IntegrationCreate is the builder for creating a Integration entity.
type IntegrationCreate struct {
	config
	mutation *IntegrationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *IntegrationCreate) SetCreatedAt(t time.Time) *IntegrationCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableCreatedAt(t *time.Time) *IntegrationCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IntegrationCreate) SetUpdatedAt(t time.Time) *IntegrationCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableUpdatedAt(t *time.Time) *IntegrationCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetCreatedBy sets the "created_by" field.
func (ic *IntegrationCreate) SetCreatedBy(s string) *IntegrationCreate {
	ic.mutation.SetCreatedBy(s)
	return ic
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableCreatedBy(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetCreatedBy(*s)
	}
	return ic
}

// SetUpdatedBy sets the "updated_by" field.
func (ic *IntegrationCreate) SetUpdatedBy(s string) *IntegrationCreate {
	ic.mutation.SetUpdatedBy(s)
	return ic
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableUpdatedBy(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetUpdatedBy(*s)
	}
	return ic
}

// SetDeletedAt sets the "deleted_at" field.
func (ic *IntegrationCreate) SetDeletedAt(t time.Time) *IntegrationCreate {
	ic.mutation.SetDeletedAt(t)
	return ic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableDeletedAt(t *time.Time) *IntegrationCreate {
	if t != nil {
		ic.SetDeletedAt(*t)
	}
	return ic
}

// SetDeletedBy sets the "deleted_by" field.
func (ic *IntegrationCreate) SetDeletedBy(s string) *IntegrationCreate {
	ic.mutation.SetDeletedBy(s)
	return ic
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableDeletedBy(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetDeletedBy(*s)
	}
	return ic
}

// SetOwnerID sets the "owner_id" field.
func (ic *IntegrationCreate) SetOwnerID(s string) *IntegrationCreate {
	ic.mutation.SetOwnerID(s)
	return ic
}

// SetName sets the "name" field.
func (ic *IntegrationCreate) SetName(s string) *IntegrationCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetDescription sets the "description" field.
func (ic *IntegrationCreate) SetDescription(s string) *IntegrationCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableDescription(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetDescription(*s)
	}
	return ic
}

// SetKind sets the "kind" field.
func (ic *IntegrationCreate) SetKind(s string) *IntegrationCreate {
	ic.mutation.SetKind(s)
	return ic
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableKind(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetKind(*s)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IntegrationCreate) SetID(s string) *IntegrationCreate {
	ic.mutation.SetID(s)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IntegrationCreate) SetNillableID(s *string) *IntegrationCreate {
	if s != nil {
		ic.SetID(*s)
	}
	return ic
}

// SetOwner sets the "owner" edge to the Organization entity.
func (ic *IntegrationCreate) SetOwner(o *Organization) *IntegrationCreate {
	return ic.SetOwnerID(o.ID)
}

// AddSecretIDs adds the "secrets" edge to the Hush entity by IDs.
func (ic *IntegrationCreate) AddSecretIDs(ids ...string) *IntegrationCreate {
	ic.mutation.AddSecretIDs(ids...)
	return ic
}

// AddSecrets adds the "secrets" edges to the Hush entity.
func (ic *IntegrationCreate) AddSecrets(h ...*Hush) *IntegrationCreate {
	ids := make([]string, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return ic.AddSecretIDs(ids...)
}

// AddOauth2tokenIDs adds the "oauth2tokens" edge to the OhAuthTooToken entity by IDs.
func (ic *IntegrationCreate) AddOauth2tokenIDs(ids ...string) *IntegrationCreate {
	ic.mutation.AddOauth2tokenIDs(ids...)
	return ic
}

// AddOauth2tokens adds the "oauth2tokens" edges to the OhAuthTooToken entity.
func (ic *IntegrationCreate) AddOauth2tokens(o ...*OhAuthTooToken) *IntegrationCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return ic.AddOauth2tokenIDs(ids...)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (ic *IntegrationCreate) AddEventIDs(ids ...string) *IntegrationCreate {
	ic.mutation.AddEventIDs(ids...)
	return ic
}

// AddEvents adds the "events" edges to the Event entity.
func (ic *IntegrationCreate) AddEvents(e ...*Event) *IntegrationCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ic.AddEventIDs(ids...)
}

// Mutation returns the IntegrationMutation object of the builder.
func (ic *IntegrationCreate) Mutation() *IntegrationMutation {
	return ic.mutation
}

// Save creates the Integration in the database.
func (ic *IntegrationCreate) Save(ctx context.Context) (*Integration, error) {
	if err := ic.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IntegrationCreate) SaveX(ctx context.Context) *Integration {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IntegrationCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IntegrationCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IntegrationCreate) defaults() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		if integration.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := integration.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		if integration.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := integration.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		if integration.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized integration.DefaultID (forgotten import generated/runtime?)")
		}
		v := integration.DefaultID()
		ic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ic *IntegrationCreate) check() error {
	if _, ok := ic.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "Integration.owner_id"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Integration.name"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := integration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Integration.name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`generated: missing required edge "Integration.owner"`)}
	}
	return nil
}

func (ic *IntegrationCreate) sqlSave(ctx context.Context) (*Integration, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Integration.ID type: %T", _spec.ID.Value)
		}
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IntegrationCreate) createSpec() (*Integration, *sqlgraph.CreateSpec) {
	var (
		_node = &Integration{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(integration.Table, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	)
	_spec.Schema = ic.schemaConfig.Integration
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(integration.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(integration.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.CreatedBy(); ok {
		_spec.SetField(integration.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ic.mutation.UpdatedBy(); ok {
		_spec.SetField(integration.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ic.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ic.mutation.DeletedBy(); ok {
		_spec.SetField(integration.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if nodes := ic.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = ic.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.SecretsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.SecretsTable,
			Columns: integration.SecretsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hush.FieldID, field.TypeString),
			},
		}
		edge.Schema = ic.schemaConfig.IntegrationSecrets
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.Oauth2tokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.Oauth2tokensTable,
			Columns: integration.Oauth2tokensPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ohauthtootoken.FieldID, field.TypeString),
			},
		}
		edge.Schema = ic.schemaConfig.IntegrationOauth2tokens
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   integration.EventsTable,
			Columns: integration.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = ic.schemaConfig.IntegrationEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IntegrationCreateBulk is the builder for creating many Integration entities in bulk.
type IntegrationCreateBulk struct {
	config
	err      error
	builders []*IntegrationCreate
}

// Save creates the Integration entities in the database.
func (icb *IntegrationCreateBulk) Save(ctx context.Context) ([]*Integration, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Integration, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IntegrationMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IntegrationCreateBulk) SaveX(ctx context.Context) []*Integration {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IntegrationCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IntegrationCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
