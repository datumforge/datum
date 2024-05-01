// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/eventhistory"
	"github.com/datumforge/enthistory"
)

// EventHistoryCreate is the builder for creating a EventHistory entity.
type EventHistoryCreate struct {
	config
	mutation *EventHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ehc *EventHistoryCreate) SetHistoryTime(t time.Time) *EventHistoryCreate {
	ehc.mutation.SetHistoryTime(t)
	return ehc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableHistoryTime(t *time.Time) *EventHistoryCreate {
	if t != nil {
		ehc.SetHistoryTime(*t)
	}
	return ehc
}

// SetOperation sets the "operation" field.
func (ehc *EventHistoryCreate) SetOperation(et enthistory.OpType) *EventHistoryCreate {
	ehc.mutation.SetOperation(et)
	return ehc
}

// SetRef sets the "ref" field.
func (ehc *EventHistoryCreate) SetRef(s string) *EventHistoryCreate {
	ehc.mutation.SetRef(s)
	return ehc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableRef(s *string) *EventHistoryCreate {
	if s != nil {
		ehc.SetRef(*s)
	}
	return ehc
}

// SetCreatedAt sets the "created_at" field.
func (ehc *EventHistoryCreate) SetCreatedAt(t time.Time) *EventHistoryCreate {
	ehc.mutation.SetCreatedAt(t)
	return ehc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableCreatedAt(t *time.Time) *EventHistoryCreate {
	if t != nil {
		ehc.SetCreatedAt(*t)
	}
	return ehc
}

// SetUpdatedAt sets the "updated_at" field.
func (ehc *EventHistoryCreate) SetUpdatedAt(t time.Time) *EventHistoryCreate {
	ehc.mutation.SetUpdatedAt(t)
	return ehc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableUpdatedAt(t *time.Time) *EventHistoryCreate {
	if t != nil {
		ehc.SetUpdatedAt(*t)
	}
	return ehc
}

// SetCreatedBy sets the "created_by" field.
func (ehc *EventHistoryCreate) SetCreatedBy(s string) *EventHistoryCreate {
	ehc.mutation.SetCreatedBy(s)
	return ehc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableCreatedBy(s *string) *EventHistoryCreate {
	if s != nil {
		ehc.SetCreatedBy(*s)
	}
	return ehc
}

// SetUpdatedBy sets the "updated_by" field.
func (ehc *EventHistoryCreate) SetUpdatedBy(s string) *EventHistoryCreate {
	ehc.mutation.SetUpdatedBy(s)
	return ehc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableUpdatedBy(s *string) *EventHistoryCreate {
	if s != nil {
		ehc.SetUpdatedBy(*s)
	}
	return ehc
}

// SetEventID sets the "event_id" field.
func (ehc *EventHistoryCreate) SetEventID(s string) *EventHistoryCreate {
	ehc.mutation.SetEventID(s)
	return ehc
}

// SetCorrelationID sets the "correlation_id" field.
func (ehc *EventHistoryCreate) SetCorrelationID(s string) *EventHistoryCreate {
	ehc.mutation.SetCorrelationID(s)
	return ehc
}

// SetEventType sets the "event_type" field.
func (ehc *EventHistoryCreate) SetEventType(s string) *EventHistoryCreate {
	ehc.mutation.SetEventType(s)
	return ehc
}

// SetMetadata sets the "metadata" field.
func (ehc *EventHistoryCreate) SetMetadata(m map[string]interface{}) *EventHistoryCreate {
	ehc.mutation.SetMetadata(m)
	return ehc
}

// SetID sets the "id" field.
func (ehc *EventHistoryCreate) SetID(s string) *EventHistoryCreate {
	ehc.mutation.SetID(s)
	return ehc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ehc *EventHistoryCreate) SetNillableID(s *string) *EventHistoryCreate {
	if s != nil {
		ehc.SetID(*s)
	}
	return ehc
}

// Mutation returns the EventHistoryMutation object of the builder.
func (ehc *EventHistoryCreate) Mutation() *EventHistoryMutation {
	return ehc.mutation
}

// Save creates the EventHistory in the database.
func (ehc *EventHistoryCreate) Save(ctx context.Context) (*EventHistory, error) {
	ehc.defaults()
	return withHooks(ctx, ehc.sqlSave, ehc.mutation, ehc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ehc *EventHistoryCreate) SaveX(ctx context.Context) *EventHistory {
	v, err := ehc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehc *EventHistoryCreate) Exec(ctx context.Context) error {
	_, err := ehc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehc *EventHistoryCreate) ExecX(ctx context.Context) {
	if err := ehc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehc *EventHistoryCreate) defaults() {
	if _, ok := ehc.mutation.HistoryTime(); !ok {
		v := eventhistory.DefaultHistoryTime()
		ehc.mutation.SetHistoryTime(v)
	}
	if _, ok := ehc.mutation.CreatedAt(); !ok {
		v := eventhistory.DefaultCreatedAt()
		ehc.mutation.SetCreatedAt(v)
	}
	if _, ok := ehc.mutation.UpdatedAt(); !ok {
		v := eventhistory.DefaultUpdatedAt()
		ehc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ehc.mutation.ID(); !ok {
		v := eventhistory.DefaultID()
		ehc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ehc *EventHistoryCreate) check() error {
	if _, ok := ehc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "EventHistory.history_time"`)}
	}
	if _, ok := ehc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "EventHistory.operation"`)}
	}
	if v, ok := ehc.mutation.Operation(); ok {
		if err := eventhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "EventHistory.operation": %w`, err)}
		}
	}
	if _, ok := ehc.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`generated: missing required field "EventHistory.event_id"`)}
	}
	if _, ok := ehc.mutation.CorrelationID(); !ok {
		return &ValidationError{Name: "correlation_id", err: errors.New(`generated: missing required field "EventHistory.correlation_id"`)}
	}
	if _, ok := ehc.mutation.EventType(); !ok {
		return &ValidationError{Name: "event_type", err: errors.New(`generated: missing required field "EventHistory.event_type"`)}
	}
	return nil
}

func (ehc *EventHistoryCreate) sqlSave(ctx context.Context) (*EventHistory, error) {
	if err := ehc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ehc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ehc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EventHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ehc.mutation.id = &_node.ID
	ehc.mutation.done = true
	return _node, nil
}

func (ehc *EventHistoryCreate) createSpec() (*EventHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &EventHistory{config: ehc.config}
		_spec = sqlgraph.NewCreateSpec(eventhistory.Table, sqlgraph.NewFieldSpec(eventhistory.FieldID, field.TypeString))
	)
	_spec.Schema = ehc.schemaConfig.EventHistory
	if id, ok := ehc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ehc.mutation.HistoryTime(); ok {
		_spec.SetField(eventhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ehc.mutation.Operation(); ok {
		_spec.SetField(eventhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ehc.mutation.Ref(); ok {
		_spec.SetField(eventhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ehc.mutation.CreatedAt(); ok {
		_spec.SetField(eventhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ehc.mutation.UpdatedAt(); ok {
		_spec.SetField(eventhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ehc.mutation.CreatedBy(); ok {
		_spec.SetField(eventhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ehc.mutation.UpdatedBy(); ok {
		_spec.SetField(eventhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ehc.mutation.EventID(); ok {
		_spec.SetField(eventhistory.FieldEventID, field.TypeString, value)
		_node.EventID = value
	}
	if value, ok := ehc.mutation.CorrelationID(); ok {
		_spec.SetField(eventhistory.FieldCorrelationID, field.TypeString, value)
		_node.CorrelationID = value
	}
	if value, ok := ehc.mutation.EventType(); ok {
		_spec.SetField(eventhistory.FieldEventType, field.TypeString, value)
		_node.EventType = value
	}
	if value, ok := ehc.mutation.Metadata(); ok {
		_spec.SetField(eventhistory.FieldMetadata, field.TypeJSON, value)
		_node.Metadata = value
	}
	return _node, _spec
}

// EventHistoryCreateBulk is the builder for creating many EventHistory entities in bulk.
type EventHistoryCreateBulk struct {
	config
	err      error
	builders []*EventHistoryCreate
}

// Save creates the EventHistory entities in the database.
func (ehcb *EventHistoryCreateBulk) Save(ctx context.Context) ([]*EventHistory, error) {
	if ehcb.err != nil {
		return nil, ehcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ehcb.builders))
	nodes := make([]*EventHistory, len(ehcb.builders))
	mutators := make([]Mutator, len(ehcb.builders))
	for i := range ehcb.builders {
		func(i int, root context.Context) {
			builder := ehcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ehcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ehcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ehcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ehcb *EventHistoryCreateBulk) SaveX(ctx context.Context) []*EventHistory {
	v, err := ehcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehcb *EventHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ehcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehcb *EventHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ehcb.Exec(ctx); err != nil {
		panic(err)
	}
}