// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/eventhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EventHistoryUpdate is the builder for updating EventHistory entities.
type EventHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *EventHistoryMutation
}

// Where appends a list predicates to the EventHistoryUpdate builder.
func (ehu *EventHistoryUpdate) Where(ps ...predicate.EventHistory) *EventHistoryUpdate {
	ehu.mutation.Where(ps...)
	return ehu
}

// SetUpdatedAt sets the "updated_at" field.
func (ehu *EventHistoryUpdate) SetUpdatedAt(t time.Time) *EventHistoryUpdate {
	ehu.mutation.SetUpdatedAt(t)
	return ehu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ehu *EventHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *EventHistoryUpdate {
	if t != nil {
		ehu.SetUpdatedAt(*t)
	}
	return ehu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ehu *EventHistoryUpdate) ClearUpdatedAt() *EventHistoryUpdate {
	ehu.mutation.ClearUpdatedAt()
	return ehu
}

// SetUpdatedBy sets the "updated_by" field.
func (ehu *EventHistoryUpdate) SetUpdatedBy(s string) *EventHistoryUpdate {
	ehu.mutation.SetUpdatedBy(s)
	return ehu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehu *EventHistoryUpdate) SetNillableUpdatedBy(s *string) *EventHistoryUpdate {
	if s != nil {
		ehu.SetUpdatedBy(*s)
	}
	return ehu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehu *EventHistoryUpdate) ClearUpdatedBy() *EventHistoryUpdate {
	ehu.mutation.ClearUpdatedBy()
	return ehu
}

// SetEventID sets the "event_id" field.
func (ehu *EventHistoryUpdate) SetEventID(s string) *EventHistoryUpdate {
	ehu.mutation.SetEventID(s)
	return ehu
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (ehu *EventHistoryUpdate) SetNillableEventID(s *string) *EventHistoryUpdate {
	if s != nil {
		ehu.SetEventID(*s)
	}
	return ehu
}

// ClearEventID clears the value of the "event_id" field.
func (ehu *EventHistoryUpdate) ClearEventID() *EventHistoryUpdate {
	ehu.mutation.ClearEventID()
	return ehu
}

// SetCorrelationID sets the "correlation_id" field.
func (ehu *EventHistoryUpdate) SetCorrelationID(s string) *EventHistoryUpdate {
	ehu.mutation.SetCorrelationID(s)
	return ehu
}

// SetNillableCorrelationID sets the "correlation_id" field if the given value is not nil.
func (ehu *EventHistoryUpdate) SetNillableCorrelationID(s *string) *EventHistoryUpdate {
	if s != nil {
		ehu.SetCorrelationID(*s)
	}
	return ehu
}

// ClearCorrelationID clears the value of the "correlation_id" field.
func (ehu *EventHistoryUpdate) ClearCorrelationID() *EventHistoryUpdate {
	ehu.mutation.ClearCorrelationID()
	return ehu
}

// SetEventType sets the "event_type" field.
func (ehu *EventHistoryUpdate) SetEventType(s string) *EventHistoryUpdate {
	ehu.mutation.SetEventType(s)
	return ehu
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (ehu *EventHistoryUpdate) SetNillableEventType(s *string) *EventHistoryUpdate {
	if s != nil {
		ehu.SetEventType(*s)
	}
	return ehu
}

// SetMetadata sets the "metadata" field.
func (ehu *EventHistoryUpdate) SetMetadata(m map[string]interface{}) *EventHistoryUpdate {
	ehu.mutation.SetMetadata(m)
	return ehu
}

// ClearMetadata clears the value of the "metadata" field.
func (ehu *EventHistoryUpdate) ClearMetadata() *EventHistoryUpdate {
	ehu.mutation.ClearMetadata()
	return ehu
}

// Mutation returns the EventHistoryMutation object of the builder.
func (ehu *EventHistoryUpdate) Mutation() *EventHistoryMutation {
	return ehu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ehu *EventHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ehu.sqlSave, ehu.mutation, ehu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehu *EventHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ehu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ehu *EventHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ehu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehu *EventHistoryUpdate) ExecX(ctx context.Context) {
	if err := ehu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ehu *EventHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(eventhistory.Table, eventhistory.Columns, sqlgraph.NewFieldSpec(eventhistory.FieldID, field.TypeString))
	if ps := ehu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ehu.mutation.RefCleared() {
		_spec.ClearField(eventhistory.FieldRef, field.TypeString)
	}
	if ehu.mutation.CreatedAtCleared() {
		_spec.ClearField(eventhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ehu.mutation.UpdatedAt(); ok {
		_spec.SetField(eventhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehu.mutation.UpdatedAtCleared() {
		_spec.ClearField(eventhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ehu.mutation.CreatedByCleared() {
		_spec.ClearField(eventhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.UpdatedBy(); ok {
		_spec.SetField(eventhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehu.mutation.UpdatedByCleared() {
		_spec.ClearField(eventhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.EventID(); ok {
		_spec.SetField(eventhistory.FieldEventID, field.TypeString, value)
	}
	if ehu.mutation.EventIDCleared() {
		_spec.ClearField(eventhistory.FieldEventID, field.TypeString)
	}
	if value, ok := ehu.mutation.CorrelationID(); ok {
		_spec.SetField(eventhistory.FieldCorrelationID, field.TypeString, value)
	}
	if ehu.mutation.CorrelationIDCleared() {
		_spec.ClearField(eventhistory.FieldCorrelationID, field.TypeString)
	}
	if value, ok := ehu.mutation.EventType(); ok {
		_spec.SetField(eventhistory.FieldEventType, field.TypeString, value)
	}
	if value, ok := ehu.mutation.Metadata(); ok {
		_spec.SetField(eventhistory.FieldMetadata, field.TypeJSON, value)
	}
	if ehu.mutation.MetadataCleared() {
		_spec.ClearField(eventhistory.FieldMetadata, field.TypeJSON)
	}
	_spec.Node.Schema = ehu.schemaConfig.EventHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ehu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ehu.mutation.done = true
	return n, nil
}

// EventHistoryUpdateOne is the builder for updating a single EventHistory entity.
type EventHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ehuo *EventHistoryUpdateOne) SetUpdatedAt(t time.Time) *EventHistoryUpdateOne {
	ehuo.mutation.SetUpdatedAt(t)
	return ehuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ehuo *EventHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *EventHistoryUpdateOne {
	if t != nil {
		ehuo.SetUpdatedAt(*t)
	}
	return ehuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ehuo *EventHistoryUpdateOne) ClearUpdatedAt() *EventHistoryUpdateOne {
	ehuo.mutation.ClearUpdatedAt()
	return ehuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ehuo *EventHistoryUpdateOne) SetUpdatedBy(s string) *EventHistoryUpdateOne {
	ehuo.mutation.SetUpdatedBy(s)
	return ehuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehuo *EventHistoryUpdateOne) SetNillableUpdatedBy(s *string) *EventHistoryUpdateOne {
	if s != nil {
		ehuo.SetUpdatedBy(*s)
	}
	return ehuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehuo *EventHistoryUpdateOne) ClearUpdatedBy() *EventHistoryUpdateOne {
	ehuo.mutation.ClearUpdatedBy()
	return ehuo
}

// SetEventID sets the "event_id" field.
func (ehuo *EventHistoryUpdateOne) SetEventID(s string) *EventHistoryUpdateOne {
	ehuo.mutation.SetEventID(s)
	return ehuo
}

// SetNillableEventID sets the "event_id" field if the given value is not nil.
func (ehuo *EventHistoryUpdateOne) SetNillableEventID(s *string) *EventHistoryUpdateOne {
	if s != nil {
		ehuo.SetEventID(*s)
	}
	return ehuo
}

// ClearEventID clears the value of the "event_id" field.
func (ehuo *EventHistoryUpdateOne) ClearEventID() *EventHistoryUpdateOne {
	ehuo.mutation.ClearEventID()
	return ehuo
}

// SetCorrelationID sets the "correlation_id" field.
func (ehuo *EventHistoryUpdateOne) SetCorrelationID(s string) *EventHistoryUpdateOne {
	ehuo.mutation.SetCorrelationID(s)
	return ehuo
}

// SetNillableCorrelationID sets the "correlation_id" field if the given value is not nil.
func (ehuo *EventHistoryUpdateOne) SetNillableCorrelationID(s *string) *EventHistoryUpdateOne {
	if s != nil {
		ehuo.SetCorrelationID(*s)
	}
	return ehuo
}

// ClearCorrelationID clears the value of the "correlation_id" field.
func (ehuo *EventHistoryUpdateOne) ClearCorrelationID() *EventHistoryUpdateOne {
	ehuo.mutation.ClearCorrelationID()
	return ehuo
}

// SetEventType sets the "event_type" field.
func (ehuo *EventHistoryUpdateOne) SetEventType(s string) *EventHistoryUpdateOne {
	ehuo.mutation.SetEventType(s)
	return ehuo
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (ehuo *EventHistoryUpdateOne) SetNillableEventType(s *string) *EventHistoryUpdateOne {
	if s != nil {
		ehuo.SetEventType(*s)
	}
	return ehuo
}

// SetMetadata sets the "metadata" field.
func (ehuo *EventHistoryUpdateOne) SetMetadata(m map[string]interface{}) *EventHistoryUpdateOne {
	ehuo.mutation.SetMetadata(m)
	return ehuo
}

// ClearMetadata clears the value of the "metadata" field.
func (ehuo *EventHistoryUpdateOne) ClearMetadata() *EventHistoryUpdateOne {
	ehuo.mutation.ClearMetadata()
	return ehuo
}

// Mutation returns the EventHistoryMutation object of the builder.
func (ehuo *EventHistoryUpdateOne) Mutation() *EventHistoryMutation {
	return ehuo.mutation
}

// Where appends a list predicates to the EventHistoryUpdate builder.
func (ehuo *EventHistoryUpdateOne) Where(ps ...predicate.EventHistory) *EventHistoryUpdateOne {
	ehuo.mutation.Where(ps...)
	return ehuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ehuo *EventHistoryUpdateOne) Select(field string, fields ...string) *EventHistoryUpdateOne {
	ehuo.fields = append([]string{field}, fields...)
	return ehuo
}

// Save executes the query and returns the updated EventHistory entity.
func (ehuo *EventHistoryUpdateOne) Save(ctx context.Context) (*EventHistory, error) {
	return withHooks(ctx, ehuo.sqlSave, ehuo.mutation, ehuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehuo *EventHistoryUpdateOne) SaveX(ctx context.Context) *EventHistory {
	node, err := ehuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ehuo *EventHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ehuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehuo *EventHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ehuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ehuo *EventHistoryUpdateOne) sqlSave(ctx context.Context) (_node *EventHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(eventhistory.Table, eventhistory.Columns, sqlgraph.NewFieldSpec(eventhistory.FieldID, field.TypeString))
	id, ok := ehuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "EventHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ehuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventhistory.FieldID)
		for _, f := range fields {
			if !eventhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != eventhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ehuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ehuo.mutation.RefCleared() {
		_spec.ClearField(eventhistory.FieldRef, field.TypeString)
	}
	if ehuo.mutation.CreatedAtCleared() {
		_spec.ClearField(eventhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ehuo.mutation.UpdatedAt(); ok {
		_spec.SetField(eventhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(eventhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ehuo.mutation.CreatedByCleared() {
		_spec.ClearField(eventhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.UpdatedBy(); ok {
		_spec.SetField(eventhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehuo.mutation.UpdatedByCleared() {
		_spec.ClearField(eventhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.EventID(); ok {
		_spec.SetField(eventhistory.FieldEventID, field.TypeString, value)
	}
	if ehuo.mutation.EventIDCleared() {
		_spec.ClearField(eventhistory.FieldEventID, field.TypeString)
	}
	if value, ok := ehuo.mutation.CorrelationID(); ok {
		_spec.SetField(eventhistory.FieldCorrelationID, field.TypeString, value)
	}
	if ehuo.mutation.CorrelationIDCleared() {
		_spec.ClearField(eventhistory.FieldCorrelationID, field.TypeString)
	}
	if value, ok := ehuo.mutation.EventType(); ok {
		_spec.SetField(eventhistory.FieldEventType, field.TypeString, value)
	}
	if value, ok := ehuo.mutation.Metadata(); ok {
		_spec.SetField(eventhistory.FieldMetadata, field.TypeJSON, value)
	}
	if ehuo.mutation.MetadataCleared() {
		_spec.ClearField(eventhistory.FieldMetadata, field.TypeJSON)
	}
	_spec.Node.Schema = ehuo.schemaConfig.EventHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehuo.schemaConfig)
	_node = &EventHistory{config: ehuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ehuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{eventhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ehuo.mutation.done = true
	return _node, nil
}
