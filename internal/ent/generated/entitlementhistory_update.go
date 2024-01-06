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
	"github.com/datumforge/datum/internal/ent/generated/entitlementhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementHistoryUpdate is the builder for updating EntitlementHistory entities.
type EntitlementHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *EntitlementHistoryMutation
}

// Where appends a list predicates to the EntitlementHistoryUpdate builder.
func (ehu *EntitlementHistoryUpdate) Where(ps ...predicate.EntitlementHistory) *EntitlementHistoryUpdate {
	ehu.mutation.Where(ps...)
	return ehu
}

// SetUpdatedAt sets the "updated_at" field.
func (ehu *EntitlementHistoryUpdate) SetUpdatedAt(t time.Time) *EntitlementHistoryUpdate {
	ehu.mutation.SetUpdatedAt(t)
	return ehu
}

// SetUpdatedBy sets the "updated_by" field.
func (ehu *EntitlementHistoryUpdate) SetUpdatedBy(s string) *EntitlementHistoryUpdate {
	ehu.mutation.SetUpdatedBy(s)
	return ehu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableUpdatedBy(s *string) *EntitlementHistoryUpdate {
	if s != nil {
		ehu.SetUpdatedBy(*s)
	}
	return ehu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehu *EntitlementHistoryUpdate) ClearUpdatedBy() *EntitlementHistoryUpdate {
	ehu.mutation.ClearUpdatedBy()
	return ehu
}

// SetDeletedAt sets the "deleted_at" field.
func (ehu *EntitlementHistoryUpdate) SetDeletedAt(t time.Time) *EntitlementHistoryUpdate {
	ehu.mutation.SetDeletedAt(t)
	return ehu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableDeletedAt(t *time.Time) *EntitlementHistoryUpdate {
	if t != nil {
		ehu.SetDeletedAt(*t)
	}
	return ehu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ehu *EntitlementHistoryUpdate) ClearDeletedAt() *EntitlementHistoryUpdate {
	ehu.mutation.ClearDeletedAt()
	return ehu
}

// SetDeletedBy sets the "deleted_by" field.
func (ehu *EntitlementHistoryUpdate) SetDeletedBy(s string) *EntitlementHistoryUpdate {
	ehu.mutation.SetDeletedBy(s)
	return ehu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableDeletedBy(s *string) *EntitlementHistoryUpdate {
	if s != nil {
		ehu.SetDeletedBy(*s)
	}
	return ehu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ehu *EntitlementHistoryUpdate) ClearDeletedBy() *EntitlementHistoryUpdate {
	ehu.mutation.ClearDeletedBy()
	return ehu
}

// SetTier sets the "tier" field.
func (ehu *EntitlementHistoryUpdate) SetTier(e entitlementhistory.Tier) *EntitlementHistoryUpdate {
	ehu.mutation.SetTier(e)
	return ehu
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableTier(e *entitlementhistory.Tier) *EntitlementHistoryUpdate {
	if e != nil {
		ehu.SetTier(*e)
	}
	return ehu
}

// SetExternalCustomerID sets the "external_customer_id" field.
func (ehu *EntitlementHistoryUpdate) SetExternalCustomerID(s string) *EntitlementHistoryUpdate {
	ehu.mutation.SetExternalCustomerID(s)
	return ehu
}

// SetNillableExternalCustomerID sets the "external_customer_id" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableExternalCustomerID(s *string) *EntitlementHistoryUpdate {
	if s != nil {
		ehu.SetExternalCustomerID(*s)
	}
	return ehu
}

// ClearExternalCustomerID clears the value of the "external_customer_id" field.
func (ehu *EntitlementHistoryUpdate) ClearExternalCustomerID() *EntitlementHistoryUpdate {
	ehu.mutation.ClearExternalCustomerID()
	return ehu
}

// SetExternalSubscriptionID sets the "external_subscription_id" field.
func (ehu *EntitlementHistoryUpdate) SetExternalSubscriptionID(s string) *EntitlementHistoryUpdate {
	ehu.mutation.SetExternalSubscriptionID(s)
	return ehu
}

// SetNillableExternalSubscriptionID sets the "external_subscription_id" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableExternalSubscriptionID(s *string) *EntitlementHistoryUpdate {
	if s != nil {
		ehu.SetExternalSubscriptionID(*s)
	}
	return ehu
}

// ClearExternalSubscriptionID clears the value of the "external_subscription_id" field.
func (ehu *EntitlementHistoryUpdate) ClearExternalSubscriptionID() *EntitlementHistoryUpdate {
	ehu.mutation.ClearExternalSubscriptionID()
	return ehu
}

// SetExpires sets the "expires" field.
func (ehu *EntitlementHistoryUpdate) SetExpires(b bool) *EntitlementHistoryUpdate {
	ehu.mutation.SetExpires(b)
	return ehu
}

// SetNillableExpires sets the "expires" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableExpires(b *bool) *EntitlementHistoryUpdate {
	if b != nil {
		ehu.SetExpires(*b)
	}
	return ehu
}

// SetExpiresAt sets the "expires_at" field.
func (ehu *EntitlementHistoryUpdate) SetExpiresAt(t time.Time) *EntitlementHistoryUpdate {
	ehu.mutation.SetExpiresAt(t)
	return ehu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableExpiresAt(t *time.Time) *EntitlementHistoryUpdate {
	if t != nil {
		ehu.SetExpiresAt(*t)
	}
	return ehu
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (ehu *EntitlementHistoryUpdate) ClearExpiresAt() *EntitlementHistoryUpdate {
	ehu.mutation.ClearExpiresAt()
	return ehu
}

// SetCancelled sets the "cancelled" field.
func (ehu *EntitlementHistoryUpdate) SetCancelled(b bool) *EntitlementHistoryUpdate {
	ehu.mutation.SetCancelled(b)
	return ehu
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ehu *EntitlementHistoryUpdate) SetNillableCancelled(b *bool) *EntitlementHistoryUpdate {
	if b != nil {
		ehu.SetCancelled(*b)
	}
	return ehu
}

// Mutation returns the EntitlementHistoryMutation object of the builder.
func (ehu *EntitlementHistoryUpdate) Mutation() *EntitlementHistoryMutation {
	return ehu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ehu *EntitlementHistoryUpdate) Save(ctx context.Context) (int, error) {
	ehu.defaults()
	return withHooks(ctx, ehu.sqlSave, ehu.mutation, ehu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehu *EntitlementHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ehu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ehu *EntitlementHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ehu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehu *EntitlementHistoryUpdate) ExecX(ctx context.Context) {
	if err := ehu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehu *EntitlementHistoryUpdate) defaults() {
	if _, ok := ehu.mutation.UpdatedAt(); !ok {
		v := entitlementhistory.UpdateDefaultUpdatedAt()
		ehu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ehu *EntitlementHistoryUpdate) check() error {
	if v, ok := ehu.mutation.Tier(); ok {
		if err := entitlementhistory.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "EntitlementHistory.tier": %w`, err)}
		}
	}
	return nil
}

func (ehu *EntitlementHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ehu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlementhistory.Table, entitlementhistory.Columns, sqlgraph.NewFieldSpec(entitlementhistory.FieldID, field.TypeString))
	if ps := ehu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ehu.mutation.RefCleared() {
		_spec.ClearField(entitlementhistory.FieldRef, field.TypeString)
	}
	if value, ok := ehu.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehu.mutation.CreatedByCleared() {
		_spec.ClearField(entitlementhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehu.mutation.UpdatedByCleared() {
		_spec.ClearField(entitlementhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.DeletedAt(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ehu.mutation.DeletedAtCleared() {
		_spec.ClearField(entitlementhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ehu.mutation.DeletedBy(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ehu.mutation.DeletedByCleared() {
		_spec.ClearField(entitlementhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ehu.mutation.Tier(); ok {
		_spec.SetField(entitlementhistory.FieldTier, field.TypeEnum, value)
	}
	if value, ok := ehu.mutation.ExternalCustomerID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalCustomerID, field.TypeString, value)
	}
	if ehu.mutation.ExternalCustomerIDCleared() {
		_spec.ClearField(entitlementhistory.FieldExternalCustomerID, field.TypeString)
	}
	if value, ok := ehu.mutation.ExternalSubscriptionID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalSubscriptionID, field.TypeString, value)
	}
	if ehu.mutation.ExternalSubscriptionIDCleared() {
		_spec.ClearField(entitlementhistory.FieldExternalSubscriptionID, field.TypeString)
	}
	if value, ok := ehu.mutation.Expires(); ok {
		_spec.SetField(entitlementhistory.FieldExpires, field.TypeBool, value)
	}
	if value, ok := ehu.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlementhistory.FieldExpiresAt, field.TypeTime, value)
	}
	if ehu.mutation.ExpiresAtCleared() {
		_spec.ClearField(entitlementhistory.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := ehu.mutation.Cancelled(); ok {
		_spec.SetField(entitlementhistory.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = ehu.schemaConfig.EntitlementHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ehu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlementhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ehu.mutation.done = true
	return n, nil
}

// EntitlementHistoryUpdateOne is the builder for updating a single EntitlementHistory entity.
type EntitlementHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntitlementHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ehuo *EntitlementHistoryUpdateOne) SetUpdatedAt(t time.Time) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetUpdatedAt(t)
	return ehuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ehuo *EntitlementHistoryUpdateOne) SetUpdatedBy(s string) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetUpdatedBy(s)
	return ehuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableUpdatedBy(s *string) *EntitlementHistoryUpdateOne {
	if s != nil {
		ehuo.SetUpdatedBy(*s)
	}
	return ehuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearUpdatedBy() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearUpdatedBy()
	return ehuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ehuo *EntitlementHistoryUpdateOne) SetDeletedAt(t time.Time) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetDeletedAt(t)
	return ehuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *EntitlementHistoryUpdateOne {
	if t != nil {
		ehuo.SetDeletedAt(*t)
	}
	return ehuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearDeletedAt() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearDeletedAt()
	return ehuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ehuo *EntitlementHistoryUpdateOne) SetDeletedBy(s string) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetDeletedBy(s)
	return ehuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableDeletedBy(s *string) *EntitlementHistoryUpdateOne {
	if s != nil {
		ehuo.SetDeletedBy(*s)
	}
	return ehuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearDeletedBy() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearDeletedBy()
	return ehuo
}

// SetTier sets the "tier" field.
func (ehuo *EntitlementHistoryUpdateOne) SetTier(e entitlementhistory.Tier) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetTier(e)
	return ehuo
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableTier(e *entitlementhistory.Tier) *EntitlementHistoryUpdateOne {
	if e != nil {
		ehuo.SetTier(*e)
	}
	return ehuo
}

// SetExternalCustomerID sets the "external_customer_id" field.
func (ehuo *EntitlementHistoryUpdateOne) SetExternalCustomerID(s string) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetExternalCustomerID(s)
	return ehuo
}

// SetNillableExternalCustomerID sets the "external_customer_id" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableExternalCustomerID(s *string) *EntitlementHistoryUpdateOne {
	if s != nil {
		ehuo.SetExternalCustomerID(*s)
	}
	return ehuo
}

// ClearExternalCustomerID clears the value of the "external_customer_id" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearExternalCustomerID() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearExternalCustomerID()
	return ehuo
}

// SetExternalSubscriptionID sets the "external_subscription_id" field.
func (ehuo *EntitlementHistoryUpdateOne) SetExternalSubscriptionID(s string) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetExternalSubscriptionID(s)
	return ehuo
}

// SetNillableExternalSubscriptionID sets the "external_subscription_id" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableExternalSubscriptionID(s *string) *EntitlementHistoryUpdateOne {
	if s != nil {
		ehuo.SetExternalSubscriptionID(*s)
	}
	return ehuo
}

// ClearExternalSubscriptionID clears the value of the "external_subscription_id" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearExternalSubscriptionID() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearExternalSubscriptionID()
	return ehuo
}

// SetExpires sets the "expires" field.
func (ehuo *EntitlementHistoryUpdateOne) SetExpires(b bool) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetExpires(b)
	return ehuo
}

// SetNillableExpires sets the "expires" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableExpires(b *bool) *EntitlementHistoryUpdateOne {
	if b != nil {
		ehuo.SetExpires(*b)
	}
	return ehuo
}

// SetExpiresAt sets the "expires_at" field.
func (ehuo *EntitlementHistoryUpdateOne) SetExpiresAt(t time.Time) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetExpiresAt(t)
	return ehuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableExpiresAt(t *time.Time) *EntitlementHistoryUpdateOne {
	if t != nil {
		ehuo.SetExpiresAt(*t)
	}
	return ehuo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (ehuo *EntitlementHistoryUpdateOne) ClearExpiresAt() *EntitlementHistoryUpdateOne {
	ehuo.mutation.ClearExpiresAt()
	return ehuo
}

// SetCancelled sets the "cancelled" field.
func (ehuo *EntitlementHistoryUpdateOne) SetCancelled(b bool) *EntitlementHistoryUpdateOne {
	ehuo.mutation.SetCancelled(b)
	return ehuo
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ehuo *EntitlementHistoryUpdateOne) SetNillableCancelled(b *bool) *EntitlementHistoryUpdateOne {
	if b != nil {
		ehuo.SetCancelled(*b)
	}
	return ehuo
}

// Mutation returns the EntitlementHistoryMutation object of the builder.
func (ehuo *EntitlementHistoryUpdateOne) Mutation() *EntitlementHistoryMutation {
	return ehuo.mutation
}

// Where appends a list predicates to the EntitlementHistoryUpdate builder.
func (ehuo *EntitlementHistoryUpdateOne) Where(ps ...predicate.EntitlementHistory) *EntitlementHistoryUpdateOne {
	ehuo.mutation.Where(ps...)
	return ehuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ehuo *EntitlementHistoryUpdateOne) Select(field string, fields ...string) *EntitlementHistoryUpdateOne {
	ehuo.fields = append([]string{field}, fields...)
	return ehuo
}

// Save executes the query and returns the updated EntitlementHistory entity.
func (ehuo *EntitlementHistoryUpdateOne) Save(ctx context.Context) (*EntitlementHistory, error) {
	ehuo.defaults()
	return withHooks(ctx, ehuo.sqlSave, ehuo.mutation, ehuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ehuo *EntitlementHistoryUpdateOne) SaveX(ctx context.Context) *EntitlementHistory {
	node, err := ehuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ehuo *EntitlementHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ehuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehuo *EntitlementHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ehuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehuo *EntitlementHistoryUpdateOne) defaults() {
	if _, ok := ehuo.mutation.UpdatedAt(); !ok {
		v := entitlementhistory.UpdateDefaultUpdatedAt()
		ehuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ehuo *EntitlementHistoryUpdateOne) check() error {
	if v, ok := ehuo.mutation.Tier(); ok {
		if err := entitlementhistory.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "EntitlementHistory.tier": %w`, err)}
		}
	}
	return nil
}

func (ehuo *EntitlementHistoryUpdateOne) sqlSave(ctx context.Context) (_node *EntitlementHistory, err error) {
	if err := ehuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlementhistory.Table, entitlementhistory.Columns, sqlgraph.NewFieldSpec(entitlementhistory.FieldID, field.TypeString))
	id, ok := ehuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "EntitlementHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ehuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlementhistory.FieldID)
		for _, f := range fields {
			if !entitlementhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != entitlementhistory.FieldID {
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
		_spec.ClearField(entitlementhistory.FieldRef, field.TypeString)
	}
	if value, ok := ehuo.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ehuo.mutation.CreatedByCleared() {
		_spec.ClearField(entitlementhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ehuo.mutation.UpdatedByCleared() {
		_spec.ClearField(entitlementhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.DeletedAt(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ehuo.mutation.DeletedAtCleared() {
		_spec.ClearField(entitlementhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ehuo.mutation.DeletedBy(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ehuo.mutation.DeletedByCleared() {
		_spec.ClearField(entitlementhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ehuo.mutation.Tier(); ok {
		_spec.SetField(entitlementhistory.FieldTier, field.TypeEnum, value)
	}
	if value, ok := ehuo.mutation.ExternalCustomerID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalCustomerID, field.TypeString, value)
	}
	if ehuo.mutation.ExternalCustomerIDCleared() {
		_spec.ClearField(entitlementhistory.FieldExternalCustomerID, field.TypeString)
	}
	if value, ok := ehuo.mutation.ExternalSubscriptionID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalSubscriptionID, field.TypeString, value)
	}
	if ehuo.mutation.ExternalSubscriptionIDCleared() {
		_spec.ClearField(entitlementhistory.FieldExternalSubscriptionID, field.TypeString)
	}
	if value, ok := ehuo.mutation.Expires(); ok {
		_spec.SetField(entitlementhistory.FieldExpires, field.TypeBool, value)
	}
	if value, ok := ehuo.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlementhistory.FieldExpiresAt, field.TypeTime, value)
	}
	if ehuo.mutation.ExpiresAtCleared() {
		_spec.ClearField(entitlementhistory.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := ehuo.mutation.Cancelled(); ok {
		_spec.SetField(entitlementhistory.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = ehuo.schemaConfig.EntitlementHistory
	ctx = internal.NewSchemaConfigContext(ctx, ehuo.schemaConfig)
	_node = &EntitlementHistory{config: ehuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ehuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlementhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ehuo.mutation.done = true
	return _node, nil
}
