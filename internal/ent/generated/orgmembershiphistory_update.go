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
	"github.com/datumforge/datum/internal/ent/generated/orgmembershiphistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/pkg/enums"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrgMembershipHistoryUpdate is the builder for updating OrgMembershipHistory entities.
type OrgMembershipHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrgMembershipHistoryMutation
}

// Where appends a list predicates to the OrgMembershipHistoryUpdate builder.
func (omhu *OrgMembershipHistoryUpdate) Where(ps ...predicate.OrgMembershipHistory) *OrgMembershipHistoryUpdate {
	omhu.mutation.Where(ps...)
	return omhu
}

// SetUpdatedAt sets the "updated_at" field.
func (omhu *OrgMembershipHistoryUpdate) SetUpdatedAt(t time.Time) *OrgMembershipHistoryUpdate {
	omhu.mutation.SetUpdatedAt(t)
	return omhu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (omhu *OrgMembershipHistoryUpdate) ClearUpdatedAt() *OrgMembershipHistoryUpdate {
	omhu.mutation.ClearUpdatedAt()
	return omhu
}

// SetUpdatedBy sets the "updated_by" field.
func (omhu *OrgMembershipHistoryUpdate) SetUpdatedBy(s string) *OrgMembershipHistoryUpdate {
	omhu.mutation.SetUpdatedBy(s)
	return omhu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (omhu *OrgMembershipHistoryUpdate) SetNillableUpdatedBy(s *string) *OrgMembershipHistoryUpdate {
	if s != nil {
		omhu.SetUpdatedBy(*s)
	}
	return omhu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (omhu *OrgMembershipHistoryUpdate) ClearUpdatedBy() *OrgMembershipHistoryUpdate {
	omhu.mutation.ClearUpdatedBy()
	return omhu
}

// SetDeletedAt sets the "deleted_at" field.
func (omhu *OrgMembershipHistoryUpdate) SetDeletedAt(t time.Time) *OrgMembershipHistoryUpdate {
	omhu.mutation.SetDeletedAt(t)
	return omhu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (omhu *OrgMembershipHistoryUpdate) SetNillableDeletedAt(t *time.Time) *OrgMembershipHistoryUpdate {
	if t != nil {
		omhu.SetDeletedAt(*t)
	}
	return omhu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (omhu *OrgMembershipHistoryUpdate) ClearDeletedAt() *OrgMembershipHistoryUpdate {
	omhu.mutation.ClearDeletedAt()
	return omhu
}

// SetDeletedBy sets the "deleted_by" field.
func (omhu *OrgMembershipHistoryUpdate) SetDeletedBy(s string) *OrgMembershipHistoryUpdate {
	omhu.mutation.SetDeletedBy(s)
	return omhu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (omhu *OrgMembershipHistoryUpdate) SetNillableDeletedBy(s *string) *OrgMembershipHistoryUpdate {
	if s != nil {
		omhu.SetDeletedBy(*s)
	}
	return omhu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (omhu *OrgMembershipHistoryUpdate) ClearDeletedBy() *OrgMembershipHistoryUpdate {
	omhu.mutation.ClearDeletedBy()
	return omhu
}

// SetRole sets the "role" field.
func (omhu *OrgMembershipHistoryUpdate) SetRole(e enums.Role) *OrgMembershipHistoryUpdate {
	omhu.mutation.SetRole(e)
	return omhu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (omhu *OrgMembershipHistoryUpdate) SetNillableRole(e *enums.Role) *OrgMembershipHistoryUpdate {
	if e != nil {
		omhu.SetRole(*e)
	}
	return omhu
}

// Mutation returns the OrgMembershipHistoryMutation object of the builder.
func (omhu *OrgMembershipHistoryUpdate) Mutation() *OrgMembershipHistoryMutation {
	return omhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (omhu *OrgMembershipHistoryUpdate) Save(ctx context.Context) (int, error) {
	if err := omhu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, omhu.sqlSave, omhu.mutation, omhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (omhu *OrgMembershipHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := omhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (omhu *OrgMembershipHistoryUpdate) Exec(ctx context.Context) error {
	_, err := omhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omhu *OrgMembershipHistoryUpdate) ExecX(ctx context.Context) {
	if err := omhu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (omhu *OrgMembershipHistoryUpdate) defaults() error {
	if _, ok := omhu.mutation.UpdatedAt(); !ok && !omhu.mutation.UpdatedAtCleared() {
		if orgmembershiphistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.UpdateDefaultUpdatedAt()
		omhu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (omhu *OrgMembershipHistoryUpdate) check() error {
	if v, ok := omhu.mutation.Role(); ok {
		if err := orgmembershiphistory.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "OrgMembershipHistory.role": %w`, err)}
		}
	}
	return nil
}

func (omhu *OrgMembershipHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := omhu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgmembershiphistory.Table, orgmembershiphistory.Columns, sqlgraph.NewFieldSpec(orgmembershiphistory.FieldID, field.TypeString))
	if ps := omhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if omhu.mutation.RefCleared() {
		_spec.ClearField(orgmembershiphistory.FieldRef, field.TypeString)
	}
	if omhu.mutation.CreatedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := omhu.mutation.UpdatedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if omhu.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldUpdatedAt, field.TypeTime)
	}
	if omhu.mutation.CreatedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := omhu.mutation.UpdatedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedBy, field.TypeString, value)
	}
	if omhu.mutation.UpdatedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := omhu.mutation.DeletedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedAt, field.TypeTime, value)
	}
	if omhu.mutation.DeletedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := omhu.mutation.DeletedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedBy, field.TypeString, value)
	}
	if omhu.mutation.DeletedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := omhu.mutation.Role(); ok {
		_spec.SetField(orgmembershiphistory.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = omhu.schemaConfig.OrgMembershipHistory
	ctx = internal.NewSchemaConfigContext(ctx, omhu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, omhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgmembershiphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	omhu.mutation.done = true
	return n, nil
}

// OrgMembershipHistoryUpdateOne is the builder for updating a single OrgMembershipHistory entity.
type OrgMembershipHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrgMembershipHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (omhuo *OrgMembershipHistoryUpdateOne) SetUpdatedAt(t time.Time) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.SetUpdatedAt(t)
	return omhuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (omhuo *OrgMembershipHistoryUpdateOne) ClearUpdatedAt() *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.ClearUpdatedAt()
	return omhuo
}

// SetUpdatedBy sets the "updated_by" field.
func (omhuo *OrgMembershipHistoryUpdateOne) SetUpdatedBy(s string) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.SetUpdatedBy(s)
	return omhuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (omhuo *OrgMembershipHistoryUpdateOne) SetNillableUpdatedBy(s *string) *OrgMembershipHistoryUpdateOne {
	if s != nil {
		omhuo.SetUpdatedBy(*s)
	}
	return omhuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (omhuo *OrgMembershipHistoryUpdateOne) ClearUpdatedBy() *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.ClearUpdatedBy()
	return omhuo
}

// SetDeletedAt sets the "deleted_at" field.
func (omhuo *OrgMembershipHistoryUpdateOne) SetDeletedAt(t time.Time) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.SetDeletedAt(t)
	return omhuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (omhuo *OrgMembershipHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *OrgMembershipHistoryUpdateOne {
	if t != nil {
		omhuo.SetDeletedAt(*t)
	}
	return omhuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (omhuo *OrgMembershipHistoryUpdateOne) ClearDeletedAt() *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.ClearDeletedAt()
	return omhuo
}

// SetDeletedBy sets the "deleted_by" field.
func (omhuo *OrgMembershipHistoryUpdateOne) SetDeletedBy(s string) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.SetDeletedBy(s)
	return omhuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (omhuo *OrgMembershipHistoryUpdateOne) SetNillableDeletedBy(s *string) *OrgMembershipHistoryUpdateOne {
	if s != nil {
		omhuo.SetDeletedBy(*s)
	}
	return omhuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (omhuo *OrgMembershipHistoryUpdateOne) ClearDeletedBy() *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.ClearDeletedBy()
	return omhuo
}

// SetRole sets the "role" field.
func (omhuo *OrgMembershipHistoryUpdateOne) SetRole(e enums.Role) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.SetRole(e)
	return omhuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (omhuo *OrgMembershipHistoryUpdateOne) SetNillableRole(e *enums.Role) *OrgMembershipHistoryUpdateOne {
	if e != nil {
		omhuo.SetRole(*e)
	}
	return omhuo
}

// Mutation returns the OrgMembershipHistoryMutation object of the builder.
func (omhuo *OrgMembershipHistoryUpdateOne) Mutation() *OrgMembershipHistoryMutation {
	return omhuo.mutation
}

// Where appends a list predicates to the OrgMembershipHistoryUpdate builder.
func (omhuo *OrgMembershipHistoryUpdateOne) Where(ps ...predicate.OrgMembershipHistory) *OrgMembershipHistoryUpdateOne {
	omhuo.mutation.Where(ps...)
	return omhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (omhuo *OrgMembershipHistoryUpdateOne) Select(field string, fields ...string) *OrgMembershipHistoryUpdateOne {
	omhuo.fields = append([]string{field}, fields...)
	return omhuo
}

// Save executes the query and returns the updated OrgMembershipHistory entity.
func (omhuo *OrgMembershipHistoryUpdateOne) Save(ctx context.Context) (*OrgMembershipHistory, error) {
	if err := omhuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, omhuo.sqlSave, omhuo.mutation, omhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (omhuo *OrgMembershipHistoryUpdateOne) SaveX(ctx context.Context) *OrgMembershipHistory {
	node, err := omhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (omhuo *OrgMembershipHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := omhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (omhuo *OrgMembershipHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := omhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (omhuo *OrgMembershipHistoryUpdateOne) defaults() error {
	if _, ok := omhuo.mutation.UpdatedAt(); !ok && !omhuo.mutation.UpdatedAtCleared() {
		if orgmembershiphistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized orgmembershiphistory.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := orgmembershiphistory.UpdateDefaultUpdatedAt()
		omhuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (omhuo *OrgMembershipHistoryUpdateOne) check() error {
	if v, ok := omhuo.mutation.Role(); ok {
		if err := orgmembershiphistory.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "OrgMembershipHistory.role": %w`, err)}
		}
	}
	return nil
}

func (omhuo *OrgMembershipHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OrgMembershipHistory, err error) {
	if err := omhuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(orgmembershiphistory.Table, orgmembershiphistory.Columns, sqlgraph.NewFieldSpec(orgmembershiphistory.FieldID, field.TypeString))
	id, ok := omhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrgMembershipHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := omhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orgmembershiphistory.FieldID)
		for _, f := range fields {
			if !orgmembershiphistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != orgmembershiphistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := omhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if omhuo.mutation.RefCleared() {
		_spec.ClearField(orgmembershiphistory.FieldRef, field.TypeString)
	}
	if omhuo.mutation.CreatedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := omhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if omhuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldUpdatedAt, field.TypeTime)
	}
	if omhuo.mutation.CreatedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := omhuo.mutation.UpdatedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldUpdatedBy, field.TypeString, value)
	}
	if omhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := omhuo.mutation.DeletedAt(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedAt, field.TypeTime, value)
	}
	if omhuo.mutation.DeletedAtCleared() {
		_spec.ClearField(orgmembershiphistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := omhuo.mutation.DeletedBy(); ok {
		_spec.SetField(orgmembershiphistory.FieldDeletedBy, field.TypeString, value)
	}
	if omhuo.mutation.DeletedByCleared() {
		_spec.ClearField(orgmembershiphistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := omhuo.mutation.Role(); ok {
		_spec.SetField(orgmembershiphistory.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = omhuo.schemaConfig.OrgMembershipHistory
	ctx = internal.NewSchemaConfigContext(ctx, omhuo.schemaConfig)
	_node = &OrgMembershipHistory{config: omhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, omhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orgmembershiphistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	omhuo.mutation.done = true
	return _node, nil
}
