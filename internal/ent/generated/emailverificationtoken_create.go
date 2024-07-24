// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// EmailVerificationTokenCreate is the builder for creating a EmailVerificationToken entity.
type EmailVerificationTokenCreate struct {
	config
	mutation *EmailVerificationTokenMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (evtc *EmailVerificationTokenCreate) SetCreatedAt(t time.Time) *EmailVerificationTokenCreate {
	evtc.mutation.SetCreatedAt(t)
	return evtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableCreatedAt(t *time.Time) *EmailVerificationTokenCreate {
	if t != nil {
		evtc.SetCreatedAt(*t)
	}
	return evtc
}

// SetUpdatedAt sets the "updated_at" field.
func (evtc *EmailVerificationTokenCreate) SetUpdatedAt(t time.Time) *EmailVerificationTokenCreate {
	evtc.mutation.SetUpdatedAt(t)
	return evtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableUpdatedAt(t *time.Time) *EmailVerificationTokenCreate {
	if t != nil {
		evtc.SetUpdatedAt(*t)
	}
	return evtc
}

// SetCreatedBy sets the "created_by" field.
func (evtc *EmailVerificationTokenCreate) SetCreatedBy(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetCreatedBy(s)
	return evtc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableCreatedBy(s *string) *EmailVerificationTokenCreate {
	if s != nil {
		evtc.SetCreatedBy(*s)
	}
	return evtc
}

// SetUpdatedBy sets the "updated_by" field.
func (evtc *EmailVerificationTokenCreate) SetUpdatedBy(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetUpdatedBy(s)
	return evtc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableUpdatedBy(s *string) *EmailVerificationTokenCreate {
	if s != nil {
		evtc.SetUpdatedBy(*s)
	}
	return evtc
}

// SetMappingID sets the "mapping_id" field.
func (evtc *EmailVerificationTokenCreate) SetMappingID(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetMappingID(s)
	return evtc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableMappingID(s *string) *EmailVerificationTokenCreate {
	if s != nil {
		evtc.SetMappingID(*s)
	}
	return evtc
}

// SetDeletedAt sets the "deleted_at" field.
func (evtc *EmailVerificationTokenCreate) SetDeletedAt(t time.Time) *EmailVerificationTokenCreate {
	evtc.mutation.SetDeletedAt(t)
	return evtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableDeletedAt(t *time.Time) *EmailVerificationTokenCreate {
	if t != nil {
		evtc.SetDeletedAt(*t)
	}
	return evtc
}

// SetDeletedBy sets the "deleted_by" field.
func (evtc *EmailVerificationTokenCreate) SetDeletedBy(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetDeletedBy(s)
	return evtc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableDeletedBy(s *string) *EmailVerificationTokenCreate {
	if s != nil {
		evtc.SetDeletedBy(*s)
	}
	return evtc
}

// SetOwnerID sets the "owner_id" field.
func (evtc *EmailVerificationTokenCreate) SetOwnerID(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetOwnerID(s)
	return evtc
}

// SetToken sets the "token" field.
func (evtc *EmailVerificationTokenCreate) SetToken(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetToken(s)
	return evtc
}

// SetTTL sets the "ttl" field.
func (evtc *EmailVerificationTokenCreate) SetTTL(t time.Time) *EmailVerificationTokenCreate {
	evtc.mutation.SetTTL(t)
	return evtc
}

// SetEmail sets the "email" field.
func (evtc *EmailVerificationTokenCreate) SetEmail(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetEmail(s)
	return evtc
}

// SetSecret sets the "secret" field.
func (evtc *EmailVerificationTokenCreate) SetSecret(b []byte) *EmailVerificationTokenCreate {
	evtc.mutation.SetSecret(b)
	return evtc
}

// SetID sets the "id" field.
func (evtc *EmailVerificationTokenCreate) SetID(s string) *EmailVerificationTokenCreate {
	evtc.mutation.SetID(s)
	return evtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (evtc *EmailVerificationTokenCreate) SetNillableID(s *string) *EmailVerificationTokenCreate {
	if s != nil {
		evtc.SetID(*s)
	}
	return evtc
}

// SetOwner sets the "owner" edge to the User entity.
func (evtc *EmailVerificationTokenCreate) SetOwner(u *User) *EmailVerificationTokenCreate {
	return evtc.SetOwnerID(u.ID)
}

// Mutation returns the EmailVerificationTokenMutation object of the builder.
func (evtc *EmailVerificationTokenCreate) Mutation() *EmailVerificationTokenMutation {
	return evtc.mutation
}

// Save creates the EmailVerificationToken in the database.
func (evtc *EmailVerificationTokenCreate) Save(ctx context.Context) (*EmailVerificationToken, error) {
	if err := evtc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, evtc.sqlSave, evtc.mutation, evtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (evtc *EmailVerificationTokenCreate) SaveX(ctx context.Context) *EmailVerificationToken {
	v, err := evtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (evtc *EmailVerificationTokenCreate) Exec(ctx context.Context) error {
	_, err := evtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evtc *EmailVerificationTokenCreate) ExecX(ctx context.Context) {
	if err := evtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (evtc *EmailVerificationTokenCreate) defaults() error {
	if _, ok := evtc.mutation.CreatedAt(); !ok {
		if emailverificationtoken.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized emailverificationtoken.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := emailverificationtoken.DefaultCreatedAt()
		evtc.mutation.SetCreatedAt(v)
	}
	if _, ok := evtc.mutation.UpdatedAt(); !ok {
		if emailverificationtoken.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized emailverificationtoken.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := emailverificationtoken.DefaultUpdatedAt()
		evtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := evtc.mutation.MappingID(); !ok {
		if emailverificationtoken.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized emailverificationtoken.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := emailverificationtoken.DefaultMappingID()
		evtc.mutation.SetMappingID(v)
	}
	if _, ok := evtc.mutation.ID(); !ok {
		if emailverificationtoken.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized emailverificationtoken.DefaultID (forgotten import generated/runtime?)")
		}
		v := emailverificationtoken.DefaultID()
		evtc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (evtc *EmailVerificationTokenCreate) check() error {
	if _, ok := evtc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "EmailVerificationToken.mapping_id"`)}
	}
	if _, ok := evtc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "EmailVerificationToken.owner_id"`)}
	}
	if _, ok := evtc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`generated: missing required field "EmailVerificationToken.token"`)}
	}
	if v, ok := evtc.mutation.Token(); ok {
		if err := emailverificationtoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`generated: validator failed for field "EmailVerificationToken.token": %w`, err)}
		}
	}
	if _, ok := evtc.mutation.TTL(); !ok {
		return &ValidationError{Name: "ttl", err: errors.New(`generated: missing required field "EmailVerificationToken.ttl"`)}
	}
	if _, ok := evtc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`generated: missing required field "EmailVerificationToken.email"`)}
	}
	if v, ok := evtc.mutation.Email(); ok {
		if err := emailverificationtoken.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "EmailVerificationToken.email": %w`, err)}
		}
	}
	if _, ok := evtc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`generated: missing required field "EmailVerificationToken.secret"`)}
	}
	if v, ok := evtc.mutation.Secret(); ok {
		if err := emailverificationtoken.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`generated: validator failed for field "EmailVerificationToken.secret": %w`, err)}
		}
	}
	if len(evtc.mutation.OwnerIDs()) == 0 {
		return &ValidationError{Name: "owner", err: errors.New(`generated: missing required edge "EmailVerificationToken.owner"`)}
	}
	return nil
}

func (evtc *EmailVerificationTokenCreate) sqlSave(ctx context.Context) (*EmailVerificationToken, error) {
	if err := evtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := evtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, evtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EmailVerificationToken.ID type: %T", _spec.ID.Value)
		}
	}
	evtc.mutation.id = &_node.ID
	evtc.mutation.done = true
	return _node, nil
}

func (evtc *EmailVerificationTokenCreate) createSpec() (*EmailVerificationToken, *sqlgraph.CreateSpec) {
	var (
		_node = &EmailVerificationToken{config: evtc.config}
		_spec = sqlgraph.NewCreateSpec(emailverificationtoken.Table, sqlgraph.NewFieldSpec(emailverificationtoken.FieldID, field.TypeString))
	)
	_spec.Schema = evtc.schemaConfig.EmailVerificationToken
	if id, ok := evtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := evtc.mutation.CreatedAt(); ok {
		_spec.SetField(emailverificationtoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := evtc.mutation.UpdatedAt(); ok {
		_spec.SetField(emailverificationtoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := evtc.mutation.CreatedBy(); ok {
		_spec.SetField(emailverificationtoken.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := evtc.mutation.UpdatedBy(); ok {
		_spec.SetField(emailverificationtoken.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := evtc.mutation.MappingID(); ok {
		_spec.SetField(emailverificationtoken.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := evtc.mutation.DeletedAt(); ok {
		_spec.SetField(emailverificationtoken.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := evtc.mutation.DeletedBy(); ok {
		_spec.SetField(emailverificationtoken.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := evtc.mutation.Token(); ok {
		_spec.SetField(emailverificationtoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := evtc.mutation.TTL(); ok {
		_spec.SetField(emailverificationtoken.FieldTTL, field.TypeTime, value)
		_node.TTL = &value
	}
	if value, ok := evtc.mutation.Email(); ok {
		_spec.SetField(emailverificationtoken.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := evtc.mutation.Secret(); ok {
		_spec.SetField(emailverificationtoken.FieldSecret, field.TypeBytes, value)
		_node.Secret = &value
	}
	if nodes := evtc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailverificationtoken.OwnerTable,
			Columns: []string{emailverificationtoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = evtc.schemaConfig.EmailVerificationToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EmailVerificationTokenCreateBulk is the builder for creating many EmailVerificationToken entities in bulk.
type EmailVerificationTokenCreateBulk struct {
	config
	err      error
	builders []*EmailVerificationTokenCreate
}

// Save creates the EmailVerificationToken entities in the database.
func (evtcb *EmailVerificationTokenCreateBulk) Save(ctx context.Context) ([]*EmailVerificationToken, error) {
	if evtcb.err != nil {
		return nil, evtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(evtcb.builders))
	nodes := make([]*EmailVerificationToken, len(evtcb.builders))
	mutators := make([]Mutator, len(evtcb.builders))
	for i := range evtcb.builders {
		func(i int, root context.Context) {
			builder := evtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EmailVerificationTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, evtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, evtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, evtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (evtcb *EmailVerificationTokenCreateBulk) SaveX(ctx context.Context) []*EmailVerificationToken {
	v, err := evtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (evtcb *EmailVerificationTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := evtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evtcb *EmailVerificationTokenCreateBulk) ExecX(ctx context.Context) {
	if err := evtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
