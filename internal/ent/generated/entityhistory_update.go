// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/entityhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntityHistoryUpdate is the builder for updating EntityHistory entities.
type EntityHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *EntityHistoryMutation
}

// Where appends a list predicates to the EntityHistoryUpdate builder.
func (ehu *EntityHistoryUpdate) Where(ps ...predicate.EntityHistory) *EntityHistoryUpdate {
	ehu.mutation.Where(ps...)
	return ehu
}

// SetUpdatedAt sets the "updated_at" field.
func (ehu *EntityHistoryUpdate) SetUpdatedAt(t time.Time) *EntityHistoryUpdate {
	ehu.mutation.SetUpdatedAt(t)
	return ehu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ehu *EntityHistoryUpdate) ClearUpdatedAt() *EntityHistoryUpdate {
	ehu.mutation.ClearUpdatedAt()
	return ehu
}

// SetUpdatedBy sets the "updated_by" field.
func (ehu *EntityHistoryUpdate) SetUpdatedBy(s string) *EntityHistoryUpdate {
	ehu.mutation.SetUpdatedBy(s)
	return ehu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableUpdatedBy(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetUpdatedBy(*s)
	}
	return ehu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehu *EntityHistoryUpdate) ClearUpdatedBy() *EntityHistoryUpdate {
	ehu.mutation.ClearUpdatedBy()
	return ehu
}

// SetDeletedAt sets the "deleted_at" field.
func (ehu *EntityHistoryUpdate) SetDeletedAt(t time.Time) *EntityHistoryUpdate {
	ehu.mutation.SetDeletedAt(t)
	return ehu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableDeletedAt(t *time.Time) *EntityHistoryUpdate {
	if t != nil {
		ehu.SetDeletedAt(*t)
	}
	return ehu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ehu *EntityHistoryUpdate) ClearDeletedAt() *EntityHistoryUpdate {
	ehu.mutation.ClearDeletedAt()
	return ehu
}

// SetDeletedBy sets the "deleted_by" field.
func (ehu *EntityHistoryUpdate) SetDeletedBy(s string) *EntityHistoryUpdate {
	ehu.mutation.SetDeletedBy(s)
	return ehu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableDeletedBy(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetDeletedBy(*s)
	}
	return ehu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ehu *EntityHistoryUpdate) ClearDeletedBy() *EntityHistoryUpdate {
	ehu.mutation.ClearDeletedBy()
	return ehu
}

// SetTags sets the "tags" field.
func (ehu *EntityHistoryUpdate) SetTags(s []string) *EntityHistoryUpdate {
	ehu.mutation.SetTags(s)
	return ehu
}

// AppendTags appends s to the "tags" field.
func (ehu *EntityHistoryUpdate) AppendTags(s []string) *EntityHistoryUpdate {
	ehu.mutation.AppendTags(s)
	return ehu
}

// ClearTags clears the value of the "tags" field.
func (ehu *EntityHistoryUpdate) ClearTags() *EntityHistoryUpdate {
	ehu.mutation.ClearTags()
	return ehu
}

// SetOwnerID sets the "owner_id" field.
func (ehu *EntityHistoryUpdate) SetOwnerID(s string) *EntityHistoryUpdate {
	ehu.mutation.SetOwnerID(s)
	return ehu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableOwnerID(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetOwnerID(*s)
	}
	return ehu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (ehu *EntityHistoryUpdate) ClearOwnerID() *EntityHistoryUpdate {
	ehu.mutation.ClearOwnerID()
	return ehu
}

// SetName sets the "name" field.
func (ehu *EntityHistoryUpdate) SetName(s string) *EntityHistoryUpdate {
	ehu.mutation.SetName(s)
	return ehu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableName(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetName(*s)
	}
	return ehu
}

// SetDisplayName sets the "display_name" field.
func (ehu *EntityHistoryUpdate) SetDisplayName(s string) *EntityHistoryUpdate {
	ehu.mutation.SetDisplayName(s)
	return ehu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableDisplayName(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetDisplayName(*s)
	}
	return ehu
}

// SetDescription sets the "description" field.
func (ehu *EntityHistoryUpdate) SetDescription(s string) *EntityHistoryUpdate {
	ehu.mutation.SetDescription(s)
	return ehu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableDescription(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetDescription(*s)
	}
	return ehu
}

// ClearDescription clears the value of the "description" field.
func (ehu *EntityHistoryUpdate) ClearDescription() *EntityHistoryUpdate {
	ehu.mutation.ClearDescription()
	return ehu
}

// SetEntityTypeID sets the "entity_type_id" field.
func (ehu *EntityHistoryUpdate) SetEntityTypeID(s string) *EntityHistoryUpdate {
	ehu.mutation.SetEntityTypeID(s)
	return ehu
}

// SetNillableEntityTypeID sets the "entity_type_id" field if the given value is not nil.
func (ehu *EntityHistoryUpdate) SetNillableEntityTypeID(s *string) *EntityHistoryUpdate {
	if s != nil {
		ehu.SetEntityTypeID(*s)
	}
	return ehu
}

// ClearEntityTypeID clears the value of the "entity_type_id" field.
func (ehu *EntityHistoryUpdate) ClearEntityTypeID() *EntityHistoryUpdate {
	ehu.mutation.ClearEntityTypeID()
	return ehu
}

// Mutation returns the EntityHistoryMutation object of the builder.
func (ehu *EntityHistoryUpdate) Mutation() *EntityHistoryMutation {
	return ehu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ehu *EntityHistoryUpdate) Save(ctx context.Context) (int, error) {
	if err := ehu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ehu.sqlSave, ehu.mutation, ehu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehu *EntityHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ehu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ehu *EntityHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ehu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehu *EntityHistoryUpdate) ExecX(ctx context.Context) {
	if err := ehu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehu *EntityHistoryUpdate) defaults() error {
	if _, ok := ehu.mutation.UpdatedAt(); !ok && !ehu.mutation.UpdatedAtCleared() {
		if entityhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entityhistory.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entityhistory.UpdateDefaultUpdatedAt()
		ehu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ehu *EntityHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(entityhistory.Table, entityhistory.Columns, sqlgraph.NewFieldSpec(entityhistory.FieldID, field.TypeString))
	if ps := ehu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ehu.mutation.RefCleared() {
		_spec.ClearField(entityhistory.FieldRef, field.TypeString)
	}
	if ehu.mutation.CreatedAtCleared() {
		_spec.ClearField(entityhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ehu.mutation.UpdatedAt(); ok {
		_spec.SetField(entityhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehu.mutation.UpdatedAtCleared() {
		_spec.ClearField(entityhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ehu.mutation.CreatedByCleared() {
		_spec.ClearField(entityhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.UpdatedBy(); ok {
		_spec.SetField(entityhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehu.mutation.UpdatedByCleared() {
		_spec.ClearField(entityhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.DeletedAt(); ok {
		_spec.SetField(entityhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ehu.mutation.DeletedAtCleared() {
		_spec.ClearField(entityhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ehu.mutation.DeletedBy(); ok {
		_spec.SetField(entityhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ehu.mutation.DeletedByCleared() {
		_spec.ClearField(entityhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.Tags(); ok {
		_spec.SetField(entityhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ehu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entityhistory.FieldTags, value)
		})
	}
	if ehu.mutation.TagsCleared() {
		_spec.ClearField(entityhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ehu.mutation.OwnerID(); ok {
		_spec.SetField(entityhistory.FieldOwnerID, field.TypeString, value)
	}
	if ehu.mutation.OwnerIDCleared() {
		_spec.ClearField(entityhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := ehu.mutation.Name(); ok {
		_spec.SetField(entityhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ehu.mutation.DisplayName(); ok {
		_spec.SetField(entityhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ehu.mutation.Description(); ok {
		_spec.SetField(entityhistory.FieldDescription, field.TypeString, value)
	}
	if ehu.mutation.DescriptionCleared() {
		_spec.ClearField(entityhistory.FieldDescription, field.TypeString)
	}
	if value, ok := ehu.mutation.EntityTypeID(); ok {
		_spec.SetField(entityhistory.FieldEntityTypeID, field.TypeString, value)
	}
	if ehu.mutation.EntityTypeIDCleared() {
		_spec.ClearField(entityhistory.FieldEntityTypeID, field.TypeString)
	}
	_spec.Node.Schema = ehu.schemaConfig.EntityHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ehu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ehu.mutation.done = true
	return n, nil
}

// EntityHistoryUpdateOne is the builder for updating a single EntityHistory entity.
type EntityHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ehuo *EntityHistoryUpdateOne) SetUpdatedAt(t time.Time) *EntityHistoryUpdateOne {
	ehuo.mutation.SetUpdatedAt(t)
	return ehuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ehuo *EntityHistoryUpdateOne) ClearUpdatedAt() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearUpdatedAt()
	return ehuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ehuo *EntityHistoryUpdateOne) SetUpdatedBy(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetUpdatedBy(s)
	return ehuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableUpdatedBy(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetUpdatedBy(*s)
	}
	return ehuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehuo *EntityHistoryUpdateOne) ClearUpdatedBy() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearUpdatedBy()
	return ehuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ehuo *EntityHistoryUpdateOne) SetDeletedAt(t time.Time) *EntityHistoryUpdateOne {
	ehuo.mutation.SetDeletedAt(t)
	return ehuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *EntityHistoryUpdateOne {
	if t != nil {
		ehuo.SetDeletedAt(*t)
	}
	return ehuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ehuo *EntityHistoryUpdateOne) ClearDeletedAt() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearDeletedAt()
	return ehuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ehuo *EntityHistoryUpdateOne) SetDeletedBy(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetDeletedBy(s)
	return ehuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableDeletedBy(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetDeletedBy(*s)
	}
	return ehuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ehuo *EntityHistoryUpdateOne) ClearDeletedBy() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearDeletedBy()
	return ehuo
}

// SetTags sets the "tags" field.
func (ehuo *EntityHistoryUpdateOne) SetTags(s []string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetTags(s)
	return ehuo
}

// AppendTags appends s to the "tags" field.
func (ehuo *EntityHistoryUpdateOne) AppendTags(s []string) *EntityHistoryUpdateOne {
	ehuo.mutation.AppendTags(s)
	return ehuo
}

// ClearTags clears the value of the "tags" field.
func (ehuo *EntityHistoryUpdateOne) ClearTags() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearTags()
	return ehuo
}

// SetOwnerID sets the "owner_id" field.
func (ehuo *EntityHistoryUpdateOne) SetOwnerID(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetOwnerID(s)
	return ehuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableOwnerID(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetOwnerID(*s)
	}
	return ehuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (ehuo *EntityHistoryUpdateOne) ClearOwnerID() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearOwnerID()
	return ehuo
}

// SetName sets the "name" field.
func (ehuo *EntityHistoryUpdateOne) SetName(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetName(s)
	return ehuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableName(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetName(*s)
	}
	return ehuo
}

// SetDisplayName sets the "display_name" field.
func (ehuo *EntityHistoryUpdateOne) SetDisplayName(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetDisplayName(s)
	return ehuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableDisplayName(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetDisplayName(*s)
	}
	return ehuo
}

// SetDescription sets the "description" field.
func (ehuo *EntityHistoryUpdateOne) SetDescription(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetDescription(s)
	return ehuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableDescription(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetDescription(*s)
	}
	return ehuo
}

// ClearDescription clears the value of the "description" field.
func (ehuo *EntityHistoryUpdateOne) ClearDescription() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearDescription()
	return ehuo
}

// SetEntityTypeID sets the "entity_type_id" field.
func (ehuo *EntityHistoryUpdateOne) SetEntityTypeID(s string) *EntityHistoryUpdateOne {
	ehuo.mutation.SetEntityTypeID(s)
	return ehuo
}

// SetNillableEntityTypeID sets the "entity_type_id" field if the given value is not nil.
func (ehuo *EntityHistoryUpdateOne) SetNillableEntityTypeID(s *string) *EntityHistoryUpdateOne {
	if s != nil {
		ehuo.SetEntityTypeID(*s)
	}
	return ehuo
}

// ClearEntityTypeID clears the value of the "entity_type_id" field.
func (ehuo *EntityHistoryUpdateOne) ClearEntityTypeID() *EntityHistoryUpdateOne {
	ehuo.mutation.ClearEntityTypeID()
	return ehuo
}

// Mutation returns the EntityHistoryMutation object of the builder.
func (ehuo *EntityHistoryUpdateOne) Mutation() *EntityHistoryMutation {
	return ehuo.mutation
}

// Where appends a list predicates to the EntityHistoryUpdate builder.
func (ehuo *EntityHistoryUpdateOne) Where(ps ...predicate.EntityHistory) *EntityHistoryUpdateOne {
	ehuo.mutation.Where(ps...)
	return ehuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ehuo *EntityHistoryUpdateOne) Select(field string, fields ...string) *EntityHistoryUpdateOne {
	ehuo.fields = append([]string{field}, fields...)
	return ehuo
}

// Save executes the query and returns the updated EntityHistory entity.
func (ehuo *EntityHistoryUpdateOne) Save(ctx context.Context) (*EntityHistory, error) {
	if err := ehuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ehuo.sqlSave, ehuo.mutation, ehuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehuo *EntityHistoryUpdateOne) SaveX(ctx context.Context) *EntityHistory {
	node, err := ehuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ehuo *EntityHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ehuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehuo *EntityHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ehuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehuo *EntityHistoryUpdateOne) defaults() error {
	if _, ok := ehuo.mutation.UpdatedAt(); !ok && !ehuo.mutation.UpdatedAtCleared() {
		if entityhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entityhistory.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entityhistory.UpdateDefaultUpdatedAt()
		ehuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ehuo *EntityHistoryUpdateOne) sqlSave(ctx context.Context) (_node *EntityHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(entityhistory.Table, entityhistory.Columns, sqlgraph.NewFieldSpec(entityhistory.FieldID, field.TypeString))
	id, ok := ehuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "EntityHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ehuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entityhistory.FieldID)
		for _, f := range fields {
			if !entityhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != entityhistory.FieldID {
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
		_spec.ClearField(entityhistory.FieldRef, field.TypeString)
	}
	if ehuo.mutation.CreatedAtCleared() {
		_spec.ClearField(entityhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ehuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entityhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(entityhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ehuo.mutation.CreatedByCleared() {
		_spec.ClearField(entityhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.UpdatedBy(); ok {
		_spec.SetField(entityhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehuo.mutation.UpdatedByCleared() {
		_spec.ClearField(entityhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.DeletedAt(); ok {
		_spec.SetField(entityhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ehuo.mutation.DeletedAtCleared() {
		_spec.ClearField(entityhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ehuo.mutation.DeletedBy(); ok {
		_spec.SetField(entityhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ehuo.mutation.DeletedByCleared() {
		_spec.ClearField(entityhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.Tags(); ok {
		_spec.SetField(entityhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ehuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entityhistory.FieldTags, value)
		})
	}
	if ehuo.mutation.TagsCleared() {
		_spec.ClearField(entityhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ehuo.mutation.OwnerID(); ok {
		_spec.SetField(entityhistory.FieldOwnerID, field.TypeString, value)
	}
	if ehuo.mutation.OwnerIDCleared() {
		_spec.ClearField(entityhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := ehuo.mutation.Name(); ok {
		_spec.SetField(entityhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ehuo.mutation.DisplayName(); ok {
		_spec.SetField(entityhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ehuo.mutation.Description(); ok {
		_spec.SetField(entityhistory.FieldDescription, field.TypeString, value)
	}
	if ehuo.mutation.DescriptionCleared() {
		_spec.ClearField(entityhistory.FieldDescription, field.TypeString)
	}
	if value, ok := ehuo.mutation.EntityTypeID(); ok {
		_spec.SetField(entityhistory.FieldEntityTypeID, field.TypeString, value)
	}
	if ehuo.mutation.EntityTypeIDCleared() {
		_spec.ClearField(entityhistory.FieldEntityTypeID, field.TypeString)
	}
	_spec.Node.Schema = ehuo.schemaConfig.EntityHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehuo.schemaConfig)
	_node = &EntityHistory{config: ehuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ehuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entityhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ehuo.mutation.done = true
	return _node, nil
}
