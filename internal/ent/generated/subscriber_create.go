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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/subscriber"
)

// SubscriberCreate is the builder for creating a Subscriber entity.
type SubscriberCreate struct {
	config
	mutation *SubscriberMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubscriberCreate) SetCreatedAt(t time.Time) *SubscriberCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableCreatedAt(t *time.Time) *SubscriberCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubscriberCreate) SetUpdatedAt(t time.Time) *SubscriberCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableUpdatedAt(t *time.Time) *SubscriberCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetCreatedBy sets the "created_by" field.
func (sc *SubscriberCreate) SetCreatedBy(s string) *SubscriberCreate {
	sc.mutation.SetCreatedBy(s)
	return sc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableCreatedBy(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetCreatedBy(*s)
	}
	return sc
}

// SetUpdatedBy sets the "updated_by" field.
func (sc *SubscriberCreate) SetUpdatedBy(s string) *SubscriberCreate {
	sc.mutation.SetUpdatedBy(s)
	return sc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableUpdatedBy(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetUpdatedBy(*s)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SubscriberCreate) SetDeletedAt(t time.Time) *SubscriberCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableDeletedAt(t *time.Time) *SubscriberCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// SetDeletedBy sets the "deleted_by" field.
func (sc *SubscriberCreate) SetDeletedBy(s string) *SubscriberCreate {
	sc.mutation.SetDeletedBy(s)
	return sc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableDeletedBy(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetDeletedBy(*s)
	}
	return sc
}

// SetOwnerID sets the "owner_id" field.
func (sc *SubscriberCreate) SetOwnerID(s string) *SubscriberCreate {
	sc.mutation.SetOwnerID(s)
	return sc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableOwnerID(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetOwnerID(*s)
	}
	return sc
}

// SetEmail sets the "email" field.
func (sc *SubscriberCreate) SetEmail(s string) *SubscriberCreate {
	sc.mutation.SetEmail(s)
	return sc
}

// SetPhoneNumber sets the "phone_number" field.
func (sc *SubscriberCreate) SetPhoneNumber(s string) *SubscriberCreate {
	sc.mutation.SetPhoneNumber(s)
	return sc
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillablePhoneNumber(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetPhoneNumber(*s)
	}
	return sc
}

// SetVerifiedEmail sets the "verified_email" field.
func (sc *SubscriberCreate) SetVerifiedEmail(b bool) *SubscriberCreate {
	sc.mutation.SetVerifiedEmail(b)
	return sc
}

// SetNillableVerifiedEmail sets the "verified_email" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableVerifiedEmail(b *bool) *SubscriberCreate {
	if b != nil {
		sc.SetVerifiedEmail(*b)
	}
	return sc
}

// SetVerifiedPhone sets the "verified_phone" field.
func (sc *SubscriberCreate) SetVerifiedPhone(b bool) *SubscriberCreate {
	sc.mutation.SetVerifiedPhone(b)
	return sc
}

// SetNillableVerifiedPhone sets the "verified_phone" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableVerifiedPhone(b *bool) *SubscriberCreate {
	if b != nil {
		sc.SetVerifiedPhone(*b)
	}
	return sc
}

// SetActive sets the "active" field.
func (sc *SubscriberCreate) SetActive(b bool) *SubscriberCreate {
	sc.mutation.SetActive(b)
	return sc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableActive(b *bool) *SubscriberCreate {
	if b != nil {
		sc.SetActive(*b)
	}
	return sc
}

// SetToken sets the "token" field.
func (sc *SubscriberCreate) SetToken(s string) *SubscriberCreate {
	sc.mutation.SetToken(s)
	return sc
}

// SetTTL sets the "ttl" field.
func (sc *SubscriberCreate) SetTTL(t time.Time) *SubscriberCreate {
	sc.mutation.SetTTL(t)
	return sc
}

// SetSecret sets the "secret" field.
func (sc *SubscriberCreate) SetSecret(b []byte) *SubscriberCreate {
	sc.mutation.SetSecret(b)
	return sc
}

// SetID sets the "id" field.
func (sc *SubscriberCreate) SetID(s string) *SubscriberCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableID(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (sc *SubscriberCreate) SetOwner(o *Organization) *SubscriberCreate {
	return sc.SetOwnerID(o.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (sc *SubscriberCreate) AddEventIDs(ids ...string) *SubscriberCreate {
	sc.mutation.AddEventIDs(ids...)
	return sc
}

// AddEvents adds the "events" edges to the Event entity.
func (sc *SubscriberCreate) AddEvents(e ...*Event) *SubscriberCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return sc.AddEventIDs(ids...)
}

// Mutation returns the SubscriberMutation object of the builder.
func (sc *SubscriberCreate) Mutation() *SubscriberMutation {
	return sc.mutation
}

// Save creates the Subscriber in the database.
func (sc *SubscriberCreate) Save(ctx context.Context) (*Subscriber, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriberCreate) SaveX(ctx context.Context) *Subscriber {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubscriberCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubscriberCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubscriberCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if subscriber.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized subscriber.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := subscriber.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if subscriber.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized subscriber.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := subscriber.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.VerifiedEmail(); !ok {
		v := subscriber.DefaultVerifiedEmail
		sc.mutation.SetVerifiedEmail(v)
	}
	if _, ok := sc.mutation.VerifiedPhone(); !ok {
		v := subscriber.DefaultVerifiedPhone
		sc.mutation.SetVerifiedPhone(v)
	}
	if _, ok := sc.mutation.Active(); !ok {
		v := subscriber.DefaultActive
		sc.mutation.SetActive(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if subscriber.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized subscriber.DefaultID (forgotten import generated/runtime?)")
		}
		v := subscriber.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubscriberCreate) check() error {
	if _, ok := sc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`generated: missing required field "Subscriber.email"`)}
	}
	if v, ok := sc.mutation.Email(); ok {
		if err := subscriber.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "Subscriber.email": %w`, err)}
		}
	}
	if v, ok := sc.mutation.PhoneNumber(); ok {
		if err := subscriber.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`generated: validator failed for field "Subscriber.phone_number": %w`, err)}
		}
	}
	if _, ok := sc.mutation.VerifiedEmail(); !ok {
		return &ValidationError{Name: "verified_email", err: errors.New(`generated: missing required field "Subscriber.verified_email"`)}
	}
	if _, ok := sc.mutation.VerifiedPhone(); !ok {
		return &ValidationError{Name: "verified_phone", err: errors.New(`generated: missing required field "Subscriber.verified_phone"`)}
	}
	if _, ok := sc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`generated: missing required field "Subscriber.active"`)}
	}
	if _, ok := sc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`generated: missing required field "Subscriber.token"`)}
	}
	if v, ok := sc.mutation.Token(); ok {
		if err := subscriber.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`generated: validator failed for field "Subscriber.token": %w`, err)}
		}
	}
	if _, ok := sc.mutation.TTL(); !ok {
		return &ValidationError{Name: "ttl", err: errors.New(`generated: missing required field "Subscriber.ttl"`)}
	}
	if _, ok := sc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`generated: missing required field "Subscriber.secret"`)}
	}
	if v, ok := sc.mutation.Secret(); ok {
		if err := subscriber.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`generated: validator failed for field "Subscriber.secret": %w`, err)}
		}
	}
	return nil
}

func (sc *SubscriberCreate) sqlSave(ctx context.Context) (*Subscriber, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Subscriber.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubscriberCreate) createSpec() (*Subscriber, *sqlgraph.CreateSpec) {
	var (
		_node = &Subscriber{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subscriber.Table, sqlgraph.NewFieldSpec(subscriber.FieldID, field.TypeString))
	)
	_spec.Schema = sc.schemaConfig.Subscriber
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subscriber.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(subscriber.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.CreatedBy(); ok {
		_spec.SetField(subscriber.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := sc.mutation.UpdatedBy(); ok {
		_spec.SetField(subscriber.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(subscriber.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.DeletedBy(); ok {
		_spec.SetField(subscriber.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := sc.mutation.Email(); ok {
		_spec.SetField(subscriber.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := sc.mutation.PhoneNumber(); ok {
		_spec.SetField(subscriber.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := sc.mutation.VerifiedEmail(); ok {
		_spec.SetField(subscriber.FieldVerifiedEmail, field.TypeBool, value)
		_node.VerifiedEmail = value
	}
	if value, ok := sc.mutation.VerifiedPhone(); ok {
		_spec.SetField(subscriber.FieldVerifiedPhone, field.TypeBool, value)
		_node.VerifiedPhone = value
	}
	if value, ok := sc.mutation.Active(); ok {
		_spec.SetField(subscriber.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := sc.mutation.Token(); ok {
		_spec.SetField(subscriber.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := sc.mutation.TTL(); ok {
		_spec.SetField(subscriber.FieldTTL, field.TypeTime, value)
		_node.TTL = &value
	}
	if value, ok := sc.mutation.Secret(); ok {
		_spec.SetField(subscriber.FieldSecret, field.TypeBytes, value)
		_node.Secret = &value
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subscriber.OwnerTable,
			Columns: []string{subscriber.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.Subscriber
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subscriber.EventsTable,
			Columns: subscriber.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		edge.Schema = sc.schemaConfig.SubscriberEvents
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubscriberCreateBulk is the builder for creating many Subscriber entities in bulk.
type SubscriberCreateBulk struct {
	config
	err      error
	builders []*SubscriberCreate
}

// Save creates the Subscriber entities in the database.
func (scb *SubscriberCreateBulk) Save(ctx context.Context) ([]*Subscriber, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subscriber, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriberMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubscriberCreateBulk) SaveX(ctx context.Context) []*Subscriber {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubscriberCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubscriberCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
