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
	"github.com/datumforge/datum/internal/ent/generated/integrationhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationHistoryUpdate is the builder for updating IntegrationHistory entities.
type IntegrationHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *IntegrationHistoryMutation
}

// Where appends a list predicates to the IntegrationHistoryUpdate builder.
func (ihu *IntegrationHistoryUpdate) Where(ps ...predicate.IntegrationHistory) *IntegrationHistoryUpdate {
	ihu.mutation.Where(ps...)
	return ihu
}

// SetUpdatedAt sets the "updated_at" field.
func (ihu *IntegrationHistoryUpdate) SetUpdatedAt(t time.Time) *IntegrationHistoryUpdate {
	ihu.mutation.SetUpdatedAt(t)
	return ihu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *IntegrationHistoryUpdate {
	if t != nil {
		ihu.SetUpdatedAt(*t)
	}
	return ihu
}

// SetUpdatedBy sets the "updated_by" field.
func (ihu *IntegrationHistoryUpdate) SetUpdatedBy(s string) *IntegrationHistoryUpdate {
	ihu.mutation.SetUpdatedBy(s)
	return ihu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableUpdatedBy(s *string) *IntegrationHistoryUpdate {
	if s != nil {
		ihu.SetUpdatedBy(*s)
	}
	return ihu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ihu *IntegrationHistoryUpdate) ClearUpdatedBy() *IntegrationHistoryUpdate {
	ihu.mutation.ClearUpdatedBy()
	return ihu
}

// SetDeletedAt sets the "deleted_at" field.
func (ihu *IntegrationHistoryUpdate) SetDeletedAt(t time.Time) *IntegrationHistoryUpdate {
	ihu.mutation.SetDeletedAt(t)
	return ihu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableDeletedAt(t *time.Time) *IntegrationHistoryUpdate {
	if t != nil {
		ihu.SetDeletedAt(*t)
	}
	return ihu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ihu *IntegrationHistoryUpdate) ClearDeletedAt() *IntegrationHistoryUpdate {
	ihu.mutation.ClearDeletedAt()
	return ihu
}

// SetDeletedBy sets the "deleted_by" field.
func (ihu *IntegrationHistoryUpdate) SetDeletedBy(s string) *IntegrationHistoryUpdate {
	ihu.mutation.SetDeletedBy(s)
	return ihu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableDeletedBy(s *string) *IntegrationHistoryUpdate {
	if s != nil {
		ihu.SetDeletedBy(*s)
	}
	return ihu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ihu *IntegrationHistoryUpdate) ClearDeletedBy() *IntegrationHistoryUpdate {
	ihu.mutation.ClearDeletedBy()
	return ihu
}

// SetName sets the "name" field.
func (ihu *IntegrationHistoryUpdate) SetName(s string) *IntegrationHistoryUpdate {
	ihu.mutation.SetName(s)
	return ihu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableName(s *string) *IntegrationHistoryUpdate {
	if s != nil {
		ihu.SetName(*s)
	}
	return ihu
}

// SetDescription sets the "description" field.
func (ihu *IntegrationHistoryUpdate) SetDescription(s string) *IntegrationHistoryUpdate {
	ihu.mutation.SetDescription(s)
	return ihu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableDescription(s *string) *IntegrationHistoryUpdate {
	if s != nil {
		ihu.SetDescription(*s)
	}
	return ihu
}

// ClearDescription clears the value of the "description" field.
func (ihu *IntegrationHistoryUpdate) ClearDescription() *IntegrationHistoryUpdate {
	ihu.mutation.ClearDescription()
	return ihu
}

// SetKind sets the "kind" field.
func (ihu *IntegrationHistoryUpdate) SetKind(s string) *IntegrationHistoryUpdate {
	ihu.mutation.SetKind(s)
	return ihu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (ihu *IntegrationHistoryUpdate) SetNillableKind(s *string) *IntegrationHistoryUpdate {
	if s != nil {
		ihu.SetKind(*s)
	}
	return ihu
}

// ClearKind clears the value of the "kind" field.
func (ihu *IntegrationHistoryUpdate) ClearKind() *IntegrationHistoryUpdate {
	ihu.mutation.ClearKind()
	return ihu
}

// Mutation returns the IntegrationHistoryMutation object of the builder.
func (ihu *IntegrationHistoryUpdate) Mutation() *IntegrationHistoryMutation {
	return ihu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ihu *IntegrationHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ihu.sqlSave, ihu.mutation, ihu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ihu *IntegrationHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ihu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ihu *IntegrationHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ihu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihu *IntegrationHistoryUpdate) ExecX(ctx context.Context) {
	if err := ihu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ihu *IntegrationHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(integrationhistory.Table, integrationhistory.Columns, sqlgraph.NewFieldSpec(integrationhistory.FieldID, field.TypeString))
	if ps := ihu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ihu.mutation.RefCleared() {
		_spec.ClearField(integrationhistory.FieldRef, field.TypeString)
	}
	if value, ok := ihu.mutation.UpdatedAt(); ok {
		_spec.SetField(integrationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ihu.mutation.CreatedByCleared() {
		_spec.ClearField(integrationhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ihu.mutation.UpdatedBy(); ok {
		_spec.SetField(integrationhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ihu.mutation.UpdatedByCleared() {
		_spec.ClearField(integrationhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ihu.mutation.DeletedAt(); ok {
		_spec.SetField(integrationhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ihu.mutation.DeletedAtCleared() {
		_spec.ClearField(integrationhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ihu.mutation.DeletedBy(); ok {
		_spec.SetField(integrationhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ihu.mutation.DeletedByCleared() {
		_spec.ClearField(integrationhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ihu.mutation.Name(); ok {
		_spec.SetField(integrationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ihu.mutation.Description(); ok {
		_spec.SetField(integrationhistory.FieldDescription, field.TypeString, value)
	}
	if ihu.mutation.DescriptionCleared() {
		_spec.ClearField(integrationhistory.FieldDescription, field.TypeString)
	}
	if value, ok := ihu.mutation.Kind(); ok {
		_spec.SetField(integrationhistory.FieldKind, field.TypeString, value)
	}
	if ihu.mutation.KindCleared() {
		_spec.ClearField(integrationhistory.FieldKind, field.TypeString)
	}
	_spec.Node.Schema = ihu.schemaConfig.IntegrationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ihu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ihu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integrationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ihu.mutation.done = true
	return n, nil
}

// IntegrationHistoryUpdateOne is the builder for updating a single IntegrationHistory entity.
type IntegrationHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IntegrationHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ihuo *IntegrationHistoryUpdateOne) SetUpdatedAt(t time.Time) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetUpdatedAt(t)
	return ihuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *IntegrationHistoryUpdateOne {
	if t != nil {
		ihuo.SetUpdatedAt(*t)
	}
	return ihuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ihuo *IntegrationHistoryUpdateOne) SetUpdatedBy(s string) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetUpdatedBy(s)
	return ihuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableUpdatedBy(s *string) *IntegrationHistoryUpdateOne {
	if s != nil {
		ihuo.SetUpdatedBy(*s)
	}
	return ihuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ihuo *IntegrationHistoryUpdateOne) ClearUpdatedBy() *IntegrationHistoryUpdateOne {
	ihuo.mutation.ClearUpdatedBy()
	return ihuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ihuo *IntegrationHistoryUpdateOne) SetDeletedAt(t time.Time) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetDeletedAt(t)
	return ihuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *IntegrationHistoryUpdateOne {
	if t != nil {
		ihuo.SetDeletedAt(*t)
	}
	return ihuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ihuo *IntegrationHistoryUpdateOne) ClearDeletedAt() *IntegrationHistoryUpdateOne {
	ihuo.mutation.ClearDeletedAt()
	return ihuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ihuo *IntegrationHistoryUpdateOne) SetDeletedBy(s string) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetDeletedBy(s)
	return ihuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableDeletedBy(s *string) *IntegrationHistoryUpdateOne {
	if s != nil {
		ihuo.SetDeletedBy(*s)
	}
	return ihuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ihuo *IntegrationHistoryUpdateOne) ClearDeletedBy() *IntegrationHistoryUpdateOne {
	ihuo.mutation.ClearDeletedBy()
	return ihuo
}

// SetName sets the "name" field.
func (ihuo *IntegrationHistoryUpdateOne) SetName(s string) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetName(s)
	return ihuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableName(s *string) *IntegrationHistoryUpdateOne {
	if s != nil {
		ihuo.SetName(*s)
	}
	return ihuo
}

// SetDescription sets the "description" field.
func (ihuo *IntegrationHistoryUpdateOne) SetDescription(s string) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetDescription(s)
	return ihuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableDescription(s *string) *IntegrationHistoryUpdateOne {
	if s != nil {
		ihuo.SetDescription(*s)
	}
	return ihuo
}

// ClearDescription clears the value of the "description" field.
func (ihuo *IntegrationHistoryUpdateOne) ClearDescription() *IntegrationHistoryUpdateOne {
	ihuo.mutation.ClearDescription()
	return ihuo
}

// SetKind sets the "kind" field.
func (ihuo *IntegrationHistoryUpdateOne) SetKind(s string) *IntegrationHistoryUpdateOne {
	ihuo.mutation.SetKind(s)
	return ihuo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (ihuo *IntegrationHistoryUpdateOne) SetNillableKind(s *string) *IntegrationHistoryUpdateOne {
	if s != nil {
		ihuo.SetKind(*s)
	}
	return ihuo
}

// ClearKind clears the value of the "kind" field.
func (ihuo *IntegrationHistoryUpdateOne) ClearKind() *IntegrationHistoryUpdateOne {
	ihuo.mutation.ClearKind()
	return ihuo
}

// Mutation returns the IntegrationHistoryMutation object of the builder.
func (ihuo *IntegrationHistoryUpdateOne) Mutation() *IntegrationHistoryMutation {
	return ihuo.mutation
}

// Where appends a list predicates to the IntegrationHistoryUpdate builder.
func (ihuo *IntegrationHistoryUpdateOne) Where(ps ...predicate.IntegrationHistory) *IntegrationHistoryUpdateOne {
	ihuo.mutation.Where(ps...)
	return ihuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ihuo *IntegrationHistoryUpdateOne) Select(field string, fields ...string) *IntegrationHistoryUpdateOne {
	ihuo.fields = append([]string{field}, fields...)
	return ihuo
}

// Save executes the query and returns the updated IntegrationHistory entity.
func (ihuo *IntegrationHistoryUpdateOne) Save(ctx context.Context) (*IntegrationHistory, error) {
	return withHooks(ctx, ihuo.sqlSave, ihuo.mutation, ihuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ihuo *IntegrationHistoryUpdateOne) SaveX(ctx context.Context) *IntegrationHistory {
	node, err := ihuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ihuo *IntegrationHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ihuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ihuo *IntegrationHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ihuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ihuo *IntegrationHistoryUpdateOne) sqlSave(ctx context.Context) (_node *IntegrationHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(integrationhistory.Table, integrationhistory.Columns, sqlgraph.NewFieldSpec(integrationhistory.FieldID, field.TypeString))
	id, ok := ihuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "IntegrationHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ihuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integrationhistory.FieldID)
		for _, f := range fields {
			if !integrationhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != integrationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ihuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ihuo.mutation.RefCleared() {
		_spec.ClearField(integrationhistory.FieldRef, field.TypeString)
	}
	if value, ok := ihuo.mutation.UpdatedAt(); ok {
		_spec.SetField(integrationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ihuo.mutation.CreatedByCleared() {
		_spec.ClearField(integrationhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ihuo.mutation.UpdatedBy(); ok {
		_spec.SetField(integrationhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ihuo.mutation.UpdatedByCleared() {
		_spec.ClearField(integrationhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ihuo.mutation.DeletedAt(); ok {
		_spec.SetField(integrationhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ihuo.mutation.DeletedAtCleared() {
		_spec.ClearField(integrationhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ihuo.mutation.DeletedBy(); ok {
		_spec.SetField(integrationhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ihuo.mutation.DeletedByCleared() {
		_spec.ClearField(integrationhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ihuo.mutation.Name(); ok {
		_spec.SetField(integrationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ihuo.mutation.Description(); ok {
		_spec.SetField(integrationhistory.FieldDescription, field.TypeString, value)
	}
	if ihuo.mutation.DescriptionCleared() {
		_spec.ClearField(integrationhistory.FieldDescription, field.TypeString)
	}
	if value, ok := ihuo.mutation.Kind(); ok {
		_spec.SetField(integrationhistory.FieldKind, field.TypeString, value)
	}
	if ihuo.mutation.KindCleared() {
		_spec.ClearField(integrationhistory.FieldKind, field.TypeString)
	}
	_spec.Node.Schema = ihuo.schemaConfig.IntegrationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ihuo.schemaConfig)
	_node = &IntegrationHistory{config: ihuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ihuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integrationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ihuo.mutation.done = true
	return _node, nil
}
