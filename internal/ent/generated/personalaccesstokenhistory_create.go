// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstokenhistory"
	"github.com/flume/enthistory"
)

// PersonalAccessTokenHistoryCreate is the builder for creating a PersonalAccessTokenHistory entity.
type PersonalAccessTokenHistoryCreate struct {
	config
	mutation *PersonalAccessTokenHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetHistoryTime(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetHistoryTime(t)
	return pathc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableHistoryTime(t *time.Time) *PersonalAccessTokenHistoryCreate {
	if t != nil {
		pathc.SetHistoryTime(*t)
	}
	return pathc
}

// SetOperation sets the "operation" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetOperation(et enthistory.OpType) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetOperation(et)
	return pathc
}

// SetRef sets the "ref" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetRef(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetRef(s)
	return pathc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableRef(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetRef(*s)
	}
	return pathc
}

// SetCreatedAt sets the "created_at" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetCreatedAt(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetCreatedAt(t)
	return pathc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableCreatedAt(t *time.Time) *PersonalAccessTokenHistoryCreate {
	if t != nil {
		pathc.SetCreatedAt(*t)
	}
	return pathc
}

// SetUpdatedAt sets the "updated_at" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetUpdatedAt(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetUpdatedAt(t)
	return pathc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableUpdatedAt(t *time.Time) *PersonalAccessTokenHistoryCreate {
	if t != nil {
		pathc.SetUpdatedAt(*t)
	}
	return pathc
}

// SetCreatedBy sets the "created_by" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetCreatedBy(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetCreatedBy(s)
	return pathc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableCreatedBy(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetCreatedBy(*s)
	}
	return pathc
}

// SetUpdatedBy sets the "updated_by" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetUpdatedBy(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetUpdatedBy(s)
	return pathc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableUpdatedBy(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetUpdatedBy(*s)
	}
	return pathc
}

// SetDeletedAt sets the "deleted_at" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetDeletedAt(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetDeletedAt(t)
	return pathc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableDeletedAt(t *time.Time) *PersonalAccessTokenHistoryCreate {
	if t != nil {
		pathc.SetDeletedAt(*t)
	}
	return pathc
}

// SetDeletedBy sets the "deleted_by" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetDeletedBy(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetDeletedBy(s)
	return pathc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableDeletedBy(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetDeletedBy(*s)
	}
	return pathc
}

// SetName sets the "name" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetName(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetName(s)
	return pathc
}

// SetToken sets the "token" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetToken(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetToken(s)
	return pathc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableToken(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetToken(*s)
	}
	return pathc
}

// SetAbilities sets the "abilities" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetAbilities(s []string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetAbilities(s)
	return pathc
}

// SetExpiresAt sets the "expires_at" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetExpiresAt(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetExpiresAt(t)
	return pathc
}

// SetDescription sets the "description" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetDescription(s string) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetDescription(s)
	return pathc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableDescription(s *string) *PersonalAccessTokenHistoryCreate {
	if s != nil {
		pathc.SetDescription(*s)
	}
	return pathc
}

// SetLastUsedAt sets the "last_used_at" field.
func (pathc *PersonalAccessTokenHistoryCreate) SetLastUsedAt(t time.Time) *PersonalAccessTokenHistoryCreate {
	pathc.mutation.SetLastUsedAt(t)
	return pathc
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (pathc *PersonalAccessTokenHistoryCreate) SetNillableLastUsedAt(t *time.Time) *PersonalAccessTokenHistoryCreate {
	if t != nil {
		pathc.SetLastUsedAt(*t)
	}
	return pathc
}

// Mutation returns the PersonalAccessTokenHistoryMutation object of the builder.
func (pathc *PersonalAccessTokenHistoryCreate) Mutation() *PersonalAccessTokenHistoryMutation {
	return pathc.mutation
}

// Save creates the PersonalAccessTokenHistory in the database.
func (pathc *PersonalAccessTokenHistoryCreate) Save(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	pathc.defaults()
	return withHooks(ctx, pathc.sqlSave, pathc.mutation, pathc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pathc *PersonalAccessTokenHistoryCreate) SaveX(ctx context.Context) *PersonalAccessTokenHistory {
	v, err := pathc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pathc *PersonalAccessTokenHistoryCreate) Exec(ctx context.Context) error {
	_, err := pathc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pathc *PersonalAccessTokenHistoryCreate) ExecX(ctx context.Context) {
	if err := pathc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pathc *PersonalAccessTokenHistoryCreate) defaults() {
	if _, ok := pathc.mutation.HistoryTime(); !ok {
		v := personalaccesstokenhistory.DefaultHistoryTime()
		pathc.mutation.SetHistoryTime(v)
	}
	if _, ok := pathc.mutation.CreatedAt(); !ok {
		v := personalaccesstokenhistory.DefaultCreatedAt()
		pathc.mutation.SetCreatedAt(v)
	}
	if _, ok := pathc.mutation.UpdatedAt(); !ok {
		v := personalaccesstokenhistory.DefaultUpdatedAt()
		pathc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pathc.mutation.Token(); !ok {
		v := personalaccesstokenhistory.DefaultToken()
		pathc.mutation.SetToken(v)
	}
	if _, ok := pathc.mutation.Description(); !ok {
		v := personalaccesstokenhistory.DefaultDescription
		pathc.mutation.SetDescription(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pathc *PersonalAccessTokenHistoryCreate) check() error {
	if _, ok := pathc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.history_time"`)}
	}
	if _, ok := pathc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.operation"`)}
	}
	if v, ok := pathc.mutation.Operation(); ok {
		if err := personalaccesstokenhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "PersonalAccessTokenHistory.operation": %w`, err)}
		}
	}
	if _, ok := pathc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.created_at"`)}
	}
	if _, ok := pathc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.updated_at"`)}
	}
	if _, ok := pathc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.name"`)}
	}
	if _, ok := pathc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.token"`)}
	}
	if _, ok := pathc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`generated: missing required field "PersonalAccessTokenHistory.expires_at"`)}
	}
	return nil
}

func (pathc *PersonalAccessTokenHistoryCreate) sqlSave(ctx context.Context) (*PersonalAccessTokenHistory, error) {
	if err := pathc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pathc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pathc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PersonalAccessTokenHistory.ID type: %T", _spec.ID.Value)
		}
	}
	pathc.mutation.id = &_node.ID
	pathc.mutation.done = true
	return _node, nil
}

func (pathc *PersonalAccessTokenHistoryCreate) createSpec() (*PersonalAccessTokenHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &PersonalAccessTokenHistory{config: pathc.config}
		_spec = sqlgraph.NewCreateSpec(personalaccesstokenhistory.Table, sqlgraph.NewFieldSpec(personalaccesstokenhistory.FieldID, field.TypeString))
	)
	_spec.Schema = pathc.schemaConfig.PersonalAccessTokenHistory
	if value, ok := pathc.mutation.HistoryTime(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := pathc.mutation.Operation(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := pathc.mutation.Ref(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := pathc.mutation.CreatedAt(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pathc.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pathc.mutation.CreatedBy(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := pathc.mutation.UpdatedBy(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := pathc.mutation.DeletedAt(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := pathc.mutation.DeletedBy(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := pathc.mutation.Name(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pathc.mutation.Token(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := pathc.mutation.Abilities(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldAbilities, field.TypeJSON, value)
		_node.Abilities = value
	}
	if value, ok := pathc.mutation.ExpiresAt(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = &value
	}
	if value, ok := pathc.mutation.Description(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pathc.mutation.LastUsedAt(); ok {
		_spec.SetField(personalaccesstokenhistory.FieldLastUsedAt, field.TypeTime, value)
		_node.LastUsedAt = &value
	}
	return _node, _spec
}

// PersonalAccessTokenHistoryCreateBulk is the builder for creating many PersonalAccessTokenHistory entities in bulk.
type PersonalAccessTokenHistoryCreateBulk struct {
	config
	err      error
	builders []*PersonalAccessTokenHistoryCreate
}

// Save creates the PersonalAccessTokenHistory entities in the database.
func (pathcb *PersonalAccessTokenHistoryCreateBulk) Save(ctx context.Context) ([]*PersonalAccessTokenHistory, error) {
	if pathcb.err != nil {
		return nil, pathcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pathcb.builders))
	nodes := make([]*PersonalAccessTokenHistory, len(pathcb.builders))
	mutators := make([]Mutator, len(pathcb.builders))
	for i := range pathcb.builders {
		func(i int, root context.Context) {
			builder := pathcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonalAccessTokenHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, pathcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pathcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pathcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pathcb *PersonalAccessTokenHistoryCreateBulk) SaveX(ctx context.Context) []*PersonalAccessTokenHistory {
	v, err := pathcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pathcb *PersonalAccessTokenHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := pathcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pathcb *PersonalAccessTokenHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := pathcb.Exec(ctx); err != nil {
		panic(err)
	}
}
