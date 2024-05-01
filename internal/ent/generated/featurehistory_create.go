// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/featurehistory"
	"github.com/datumforge/enthistory"
)

// FeatureHistoryCreate is the builder for creating a FeatureHistory entity.
type FeatureHistoryCreate struct {
	config
	mutation *FeatureHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (fhc *FeatureHistoryCreate) SetHistoryTime(t time.Time) *FeatureHistoryCreate {
	fhc.mutation.SetHistoryTime(t)
	return fhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableHistoryTime(t *time.Time) *FeatureHistoryCreate {
	if t != nil {
		fhc.SetHistoryTime(*t)
	}
	return fhc
}

// SetOperation sets the "operation" field.
func (fhc *FeatureHistoryCreate) SetOperation(et enthistory.OpType) *FeatureHistoryCreate {
	fhc.mutation.SetOperation(et)
	return fhc
}

// SetRef sets the "ref" field.
func (fhc *FeatureHistoryCreate) SetRef(s string) *FeatureHistoryCreate {
	fhc.mutation.SetRef(s)
	return fhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableRef(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetRef(*s)
	}
	return fhc
}

// SetCreatedAt sets the "created_at" field.
func (fhc *FeatureHistoryCreate) SetCreatedAt(t time.Time) *FeatureHistoryCreate {
	fhc.mutation.SetCreatedAt(t)
	return fhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableCreatedAt(t *time.Time) *FeatureHistoryCreate {
	if t != nil {
		fhc.SetCreatedAt(*t)
	}
	return fhc
}

// SetUpdatedAt sets the "updated_at" field.
func (fhc *FeatureHistoryCreate) SetUpdatedAt(t time.Time) *FeatureHistoryCreate {
	fhc.mutation.SetUpdatedAt(t)
	return fhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableUpdatedAt(t *time.Time) *FeatureHistoryCreate {
	if t != nil {
		fhc.SetUpdatedAt(*t)
	}
	return fhc
}

// SetCreatedBy sets the "created_by" field.
func (fhc *FeatureHistoryCreate) SetCreatedBy(s string) *FeatureHistoryCreate {
	fhc.mutation.SetCreatedBy(s)
	return fhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableCreatedBy(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetCreatedBy(*s)
	}
	return fhc
}

// SetUpdatedBy sets the "updated_by" field.
func (fhc *FeatureHistoryCreate) SetUpdatedBy(s string) *FeatureHistoryCreate {
	fhc.mutation.SetUpdatedBy(s)
	return fhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableUpdatedBy(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetUpdatedBy(*s)
	}
	return fhc
}

// SetDeletedAt sets the "deleted_at" field.
func (fhc *FeatureHistoryCreate) SetDeletedAt(t time.Time) *FeatureHistoryCreate {
	fhc.mutation.SetDeletedAt(t)
	return fhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableDeletedAt(t *time.Time) *FeatureHistoryCreate {
	if t != nil {
		fhc.SetDeletedAt(*t)
	}
	return fhc
}

// SetDeletedBy sets the "deleted_by" field.
func (fhc *FeatureHistoryCreate) SetDeletedBy(s string) *FeatureHistoryCreate {
	fhc.mutation.SetDeletedBy(s)
	return fhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableDeletedBy(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetDeletedBy(*s)
	}
	return fhc
}

// SetName sets the "name" field.
func (fhc *FeatureHistoryCreate) SetName(s string) *FeatureHistoryCreate {
	fhc.mutation.SetName(s)
	return fhc
}

// SetGlobal sets the "global" field.
func (fhc *FeatureHistoryCreate) SetGlobal(b bool) *FeatureHistoryCreate {
	fhc.mutation.SetGlobal(b)
	return fhc
}

// SetNillableGlobal sets the "global" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableGlobal(b *bool) *FeatureHistoryCreate {
	if b != nil {
		fhc.SetGlobal(*b)
	}
	return fhc
}

// SetEnabled sets the "enabled" field.
func (fhc *FeatureHistoryCreate) SetEnabled(b bool) *FeatureHistoryCreate {
	fhc.mutation.SetEnabled(b)
	return fhc
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableEnabled(b *bool) *FeatureHistoryCreate {
	if b != nil {
		fhc.SetEnabled(*b)
	}
	return fhc
}

// SetDescription sets the "description" field.
func (fhc *FeatureHistoryCreate) SetDescription(s string) *FeatureHistoryCreate {
	fhc.mutation.SetDescription(s)
	return fhc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableDescription(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetDescription(*s)
	}
	return fhc
}

// SetID sets the "id" field.
func (fhc *FeatureHistoryCreate) SetID(s string) *FeatureHistoryCreate {
	fhc.mutation.SetID(s)
	return fhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fhc *FeatureHistoryCreate) SetNillableID(s *string) *FeatureHistoryCreate {
	if s != nil {
		fhc.SetID(*s)
	}
	return fhc
}

// Mutation returns the FeatureHistoryMutation object of the builder.
func (fhc *FeatureHistoryCreate) Mutation() *FeatureHistoryMutation {
	return fhc.mutation
}

// Save creates the FeatureHistory in the database.
func (fhc *FeatureHistoryCreate) Save(ctx context.Context) (*FeatureHistory, error) {
	fhc.defaults()
	return withHooks(ctx, fhc.sqlSave, fhc.mutation, fhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fhc *FeatureHistoryCreate) SaveX(ctx context.Context) *FeatureHistory {
	v, err := fhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fhc *FeatureHistoryCreate) Exec(ctx context.Context) error {
	_, err := fhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhc *FeatureHistoryCreate) ExecX(ctx context.Context) {
	if err := fhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fhc *FeatureHistoryCreate) defaults() {
	if _, ok := fhc.mutation.HistoryTime(); !ok {
		v := featurehistory.DefaultHistoryTime()
		fhc.mutation.SetHistoryTime(v)
	}
	if _, ok := fhc.mutation.CreatedAt(); !ok {
		v := featurehistory.DefaultCreatedAt()
		fhc.mutation.SetCreatedAt(v)
	}
	if _, ok := fhc.mutation.UpdatedAt(); !ok {
		v := featurehistory.DefaultUpdatedAt()
		fhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fhc.mutation.Global(); !ok {
		v := featurehistory.DefaultGlobal
		fhc.mutation.SetGlobal(v)
	}
	if _, ok := fhc.mutation.Enabled(); !ok {
		v := featurehistory.DefaultEnabled
		fhc.mutation.SetEnabled(v)
	}
	if _, ok := fhc.mutation.ID(); !ok {
		v := featurehistory.DefaultID()
		fhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fhc *FeatureHistoryCreate) check() error {
	if _, ok := fhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "FeatureHistory.history_time"`)}
	}
	if _, ok := fhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "FeatureHistory.operation"`)}
	}
	if v, ok := fhc.mutation.Operation(); ok {
		if err := featurehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "FeatureHistory.operation": %w`, err)}
		}
	}
	if _, ok := fhc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "FeatureHistory.name"`)}
	}
	if _, ok := fhc.mutation.Global(); !ok {
		return &ValidationError{Name: "global", err: errors.New(`generated: missing required field "FeatureHistory.global"`)}
	}
	if _, ok := fhc.mutation.Enabled(); !ok {
		return &ValidationError{Name: "enabled", err: errors.New(`generated: missing required field "FeatureHistory.enabled"`)}
	}
	return nil
}

func (fhc *FeatureHistoryCreate) sqlSave(ctx context.Context) (*FeatureHistory, error) {
	if err := fhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected FeatureHistory.ID type: %T", _spec.ID.Value)
		}
	}
	fhc.mutation.id = &_node.ID
	fhc.mutation.done = true
	return _node, nil
}

func (fhc *FeatureHistoryCreate) createSpec() (*FeatureHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &FeatureHistory{config: fhc.config}
		_spec = sqlgraph.NewCreateSpec(featurehistory.Table, sqlgraph.NewFieldSpec(featurehistory.FieldID, field.TypeString))
	)
	_spec.Schema = fhc.schemaConfig.FeatureHistory
	if id, ok := fhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fhc.mutation.HistoryTime(); ok {
		_spec.SetField(featurehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := fhc.mutation.Operation(); ok {
		_spec.SetField(featurehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := fhc.mutation.Ref(); ok {
		_spec.SetField(featurehistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := fhc.mutation.CreatedAt(); ok {
		_spec.SetField(featurehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fhc.mutation.UpdatedAt(); ok {
		_spec.SetField(featurehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fhc.mutation.CreatedBy(); ok {
		_spec.SetField(featurehistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := fhc.mutation.UpdatedBy(); ok {
		_spec.SetField(featurehistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := fhc.mutation.DeletedAt(); ok {
		_spec.SetField(featurehistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := fhc.mutation.DeletedBy(); ok {
		_spec.SetField(featurehistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := fhc.mutation.Name(); ok {
		_spec.SetField(featurehistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fhc.mutation.Global(); ok {
		_spec.SetField(featurehistory.FieldGlobal, field.TypeBool, value)
		_node.Global = value
	}
	if value, ok := fhc.mutation.Enabled(); ok {
		_spec.SetField(featurehistory.FieldEnabled, field.TypeBool, value)
		_node.Enabled = value
	}
	if value, ok := fhc.mutation.Description(); ok {
		_spec.SetField(featurehistory.FieldDescription, field.TypeString, value)
		_node.Description = &value
	}
	return _node, _spec
}

// FeatureHistoryCreateBulk is the builder for creating many FeatureHistory entities in bulk.
type FeatureHistoryCreateBulk struct {
	config
	err      error
	builders []*FeatureHistoryCreate
}

// Save creates the FeatureHistory entities in the database.
func (fhcb *FeatureHistoryCreateBulk) Save(ctx context.Context) ([]*FeatureHistory, error) {
	if fhcb.err != nil {
		return nil, fhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fhcb.builders))
	nodes := make([]*FeatureHistory, len(fhcb.builders))
	mutators := make([]Mutator, len(fhcb.builders))
	for i := range fhcb.builders {
		func(i int, root context.Context) {
			builder := fhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeatureHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, fhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fhcb *FeatureHistoryCreateBulk) SaveX(ctx context.Context) []*FeatureHistory {
	v, err := fhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fhcb *FeatureHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := fhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhcb *FeatureHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := fhcb.Exec(ctx); err != nil {
		panic(err)
	}
}