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
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementUpdate is the builder for updating Entitlement entities.
type EntitlementUpdate struct {
	config
	hooks    []Hook
	mutation *EntitlementMutation
}

// Where appends a list predicates to the EntitlementUpdate builder.
func (eu *EntitlementUpdate) Where(ps ...predicate.Entitlement) *EntitlementUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EntitlementUpdate) SetUpdatedAt(t time.Time) *EntitlementUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// SetCreatedBy sets the "created_by" field.
func (eu *EntitlementUpdate) SetCreatedBy(s string) *EntitlementUpdate {
	eu.mutation.SetCreatedBy(s)
	return eu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableCreatedBy(s *string) *EntitlementUpdate {
	if s != nil {
		eu.SetCreatedBy(*s)
	}
	return eu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (eu *EntitlementUpdate) ClearCreatedBy() *EntitlementUpdate {
	eu.mutation.ClearCreatedBy()
	return eu
}

// SetUpdatedBy sets the "updated_by" field.
func (eu *EntitlementUpdate) SetUpdatedBy(s string) *EntitlementUpdate {
	eu.mutation.SetUpdatedBy(s)
	return eu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableUpdatedBy(s *string) *EntitlementUpdate {
	if s != nil {
		eu.SetUpdatedBy(*s)
	}
	return eu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (eu *EntitlementUpdate) ClearUpdatedBy() *EntitlementUpdate {
	eu.mutation.ClearUpdatedBy()
	return eu
}

// SetTier sets the "tier" field.
func (eu *EntitlementUpdate) SetTier(e entitlement.Tier) *EntitlementUpdate {
	eu.mutation.SetTier(e)
	return eu
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableTier(e *entitlement.Tier) *EntitlementUpdate {
	if e != nil {
		eu.SetTier(*e)
	}
	return eu
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (eu *EntitlementUpdate) SetStripeCustomerID(s string) *EntitlementUpdate {
	eu.mutation.SetStripeCustomerID(s)
	return eu
}

// SetNillableStripeCustomerID sets the "stripe_customer_id" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableStripeCustomerID(s *string) *EntitlementUpdate {
	if s != nil {
		eu.SetStripeCustomerID(*s)
	}
	return eu
}

// ClearStripeCustomerID clears the value of the "stripe_customer_id" field.
func (eu *EntitlementUpdate) ClearStripeCustomerID() *EntitlementUpdate {
	eu.mutation.ClearStripeCustomerID()
	return eu
}

// SetStripeSubscriptionID sets the "stripe_subscription_id" field.
func (eu *EntitlementUpdate) SetStripeSubscriptionID(s string) *EntitlementUpdate {
	eu.mutation.SetStripeSubscriptionID(s)
	return eu
}

// SetNillableStripeSubscriptionID sets the "stripe_subscription_id" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableStripeSubscriptionID(s *string) *EntitlementUpdate {
	if s != nil {
		eu.SetStripeSubscriptionID(*s)
	}
	return eu
}

// ClearStripeSubscriptionID clears the value of the "stripe_subscription_id" field.
func (eu *EntitlementUpdate) ClearStripeSubscriptionID() *EntitlementUpdate {
	eu.mutation.ClearStripeSubscriptionID()
	return eu
}

// SetExpiresAt sets the "expires_at" field.
func (eu *EntitlementUpdate) SetExpiresAt(t time.Time) *EntitlementUpdate {
	eu.mutation.SetExpiresAt(t)
	return eu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableExpiresAt(t *time.Time) *EntitlementUpdate {
	if t != nil {
		eu.SetExpiresAt(*t)
	}
	return eu
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (eu *EntitlementUpdate) ClearExpiresAt() *EntitlementUpdate {
	eu.mutation.ClearExpiresAt()
	return eu
}

// SetCancelled sets the "cancelled" field.
func (eu *EntitlementUpdate) SetCancelled(b bool) *EntitlementUpdate {
	eu.mutation.SetCancelled(b)
	return eu
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (eu *EntitlementUpdate) SetNillableCancelled(b *bool) *EntitlementUpdate {
	if b != nil {
		eu.SetCancelled(*b)
	}
	return eu
}

// Mutation returns the EntitlementMutation object of the builder.
func (eu *EntitlementUpdate) Mutation() *EntitlementMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntitlementUpdate) Save(ctx context.Context) (int, error) {
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntitlementUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntitlementUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntitlementUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EntitlementUpdate) defaults() error {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		if entitlement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlement.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entitlement.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eu *EntitlementUpdate) check() error {
	if v, ok := eu.mutation.Tier(); ok {
		if err := entitlement.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "Entitlement.tier": %w`, err)}
		}
	}
	return nil
}

func (eu *EntitlementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlement.Table, entitlement.Columns, sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := eu.mutation.CreatedBy(); ok {
		_spec.SetField(entitlement.FieldCreatedBy, field.TypeString, value)
	}
	if eu.mutation.CreatedByCleared() {
		_spec.ClearField(entitlement.FieldCreatedBy, field.TypeString)
	}
	if value, ok := eu.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlement.FieldUpdatedBy, field.TypeString, value)
	}
	if eu.mutation.UpdatedByCleared() {
		_spec.ClearField(entitlement.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := eu.mutation.Tier(); ok {
		_spec.SetField(entitlement.FieldTier, field.TypeEnum, value)
	}
	if value, ok := eu.mutation.StripeCustomerID(); ok {
		_spec.SetField(entitlement.FieldStripeCustomerID, field.TypeString, value)
	}
	if eu.mutation.StripeCustomerIDCleared() {
		_spec.ClearField(entitlement.FieldStripeCustomerID, field.TypeString)
	}
	if value, ok := eu.mutation.StripeSubscriptionID(); ok {
		_spec.SetField(entitlement.FieldStripeSubscriptionID, field.TypeString, value)
	}
	if eu.mutation.StripeSubscriptionIDCleared() {
		_spec.ClearField(entitlement.FieldStripeSubscriptionID, field.TypeString)
	}
	if value, ok := eu.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlement.FieldExpiresAt, field.TypeTime, value)
	}
	if eu.mutation.ExpiresAtCleared() {
		_spec.ClearField(entitlement.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := eu.mutation.Cancelled(); ok {
		_spec.SetField(entitlement.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = eu.schemaConfig.Entitlement
	ctx = internal.NewSchemaConfigContext(ctx, eu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EntitlementUpdateOne is the builder for updating a single Entitlement entity.
type EntitlementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntitlementMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EntitlementUpdateOne) SetUpdatedAt(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// SetCreatedBy sets the "created_by" field.
func (euo *EntitlementUpdateOne) SetCreatedBy(s string) *EntitlementUpdateOne {
	euo.mutation.SetCreatedBy(s)
	return euo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableCreatedBy(s *string) *EntitlementUpdateOne {
	if s != nil {
		euo.SetCreatedBy(*s)
	}
	return euo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (euo *EntitlementUpdateOne) ClearCreatedBy() *EntitlementUpdateOne {
	euo.mutation.ClearCreatedBy()
	return euo
}

// SetUpdatedBy sets the "updated_by" field.
func (euo *EntitlementUpdateOne) SetUpdatedBy(s string) *EntitlementUpdateOne {
	euo.mutation.SetUpdatedBy(s)
	return euo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableUpdatedBy(s *string) *EntitlementUpdateOne {
	if s != nil {
		euo.SetUpdatedBy(*s)
	}
	return euo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (euo *EntitlementUpdateOne) ClearUpdatedBy() *EntitlementUpdateOne {
	euo.mutation.ClearUpdatedBy()
	return euo
}

// SetTier sets the "tier" field.
func (euo *EntitlementUpdateOne) SetTier(e entitlement.Tier) *EntitlementUpdateOne {
	euo.mutation.SetTier(e)
	return euo
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableTier(e *entitlement.Tier) *EntitlementUpdateOne {
	if e != nil {
		euo.SetTier(*e)
	}
	return euo
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (euo *EntitlementUpdateOne) SetStripeCustomerID(s string) *EntitlementUpdateOne {
	euo.mutation.SetStripeCustomerID(s)
	return euo
}

// SetNillableStripeCustomerID sets the "stripe_customer_id" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableStripeCustomerID(s *string) *EntitlementUpdateOne {
	if s != nil {
		euo.SetStripeCustomerID(*s)
	}
	return euo
}

// ClearStripeCustomerID clears the value of the "stripe_customer_id" field.
func (euo *EntitlementUpdateOne) ClearStripeCustomerID() *EntitlementUpdateOne {
	euo.mutation.ClearStripeCustomerID()
	return euo
}

// SetStripeSubscriptionID sets the "stripe_subscription_id" field.
func (euo *EntitlementUpdateOne) SetStripeSubscriptionID(s string) *EntitlementUpdateOne {
	euo.mutation.SetStripeSubscriptionID(s)
	return euo
}

// SetNillableStripeSubscriptionID sets the "stripe_subscription_id" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableStripeSubscriptionID(s *string) *EntitlementUpdateOne {
	if s != nil {
		euo.SetStripeSubscriptionID(*s)
	}
	return euo
}

// ClearStripeSubscriptionID clears the value of the "stripe_subscription_id" field.
func (euo *EntitlementUpdateOne) ClearStripeSubscriptionID() *EntitlementUpdateOne {
	euo.mutation.ClearStripeSubscriptionID()
	return euo
}

// SetExpiresAt sets the "expires_at" field.
func (euo *EntitlementUpdateOne) SetExpiresAt(t time.Time) *EntitlementUpdateOne {
	euo.mutation.SetExpiresAt(t)
	return euo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableExpiresAt(t *time.Time) *EntitlementUpdateOne {
	if t != nil {
		euo.SetExpiresAt(*t)
	}
	return euo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (euo *EntitlementUpdateOne) ClearExpiresAt() *EntitlementUpdateOne {
	euo.mutation.ClearExpiresAt()
	return euo
}

// SetCancelled sets the "cancelled" field.
func (euo *EntitlementUpdateOne) SetCancelled(b bool) *EntitlementUpdateOne {
	euo.mutation.SetCancelled(b)
	return euo
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (euo *EntitlementUpdateOne) SetNillableCancelled(b *bool) *EntitlementUpdateOne {
	if b != nil {
		euo.SetCancelled(*b)
	}
	return euo
}

// Mutation returns the EntitlementMutation object of the builder.
func (euo *EntitlementUpdateOne) Mutation() *EntitlementMutation {
	return euo.mutation
}

// Where appends a list predicates to the EntitlementUpdate builder.
func (euo *EntitlementUpdateOne) Where(ps ...predicate.Entitlement) *EntitlementUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntitlementUpdateOne) Select(field string, fields ...string) *EntitlementUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entitlement entity.
func (euo *EntitlementUpdateOne) Save(ctx context.Context) (*Entitlement, error) {
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntitlementUpdateOne) SaveX(ctx context.Context) *Entitlement {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntitlementUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntitlementUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EntitlementUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		if entitlement.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entitlement.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entitlement.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (euo *EntitlementUpdateOne) check() error {
	if v, ok := euo.mutation.Tier(); ok {
		if err := entitlement.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "Entitlement.tier": %w`, err)}
		}
	}
	return nil
}

func (euo *EntitlementUpdateOne) sqlSave(ctx context.Context) (_node *Entitlement, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entitlement.Table, entitlement.Columns, sqlgraph.NewFieldSpec(entitlement.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Entitlement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entitlement.FieldID)
		for _, f := range fields {
			if !entitlement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != entitlement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlement.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := euo.mutation.CreatedBy(); ok {
		_spec.SetField(entitlement.FieldCreatedBy, field.TypeString, value)
	}
	if euo.mutation.CreatedByCleared() {
		_spec.ClearField(entitlement.FieldCreatedBy, field.TypeString)
	}
	if value, ok := euo.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlement.FieldUpdatedBy, field.TypeString, value)
	}
	if euo.mutation.UpdatedByCleared() {
		_spec.ClearField(entitlement.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := euo.mutation.Tier(); ok {
		_spec.SetField(entitlement.FieldTier, field.TypeEnum, value)
	}
	if value, ok := euo.mutation.StripeCustomerID(); ok {
		_spec.SetField(entitlement.FieldStripeCustomerID, field.TypeString, value)
	}
	if euo.mutation.StripeCustomerIDCleared() {
		_spec.ClearField(entitlement.FieldStripeCustomerID, field.TypeString)
	}
	if value, ok := euo.mutation.StripeSubscriptionID(); ok {
		_spec.SetField(entitlement.FieldStripeSubscriptionID, field.TypeString, value)
	}
	if euo.mutation.StripeSubscriptionIDCleared() {
		_spec.ClearField(entitlement.FieldStripeSubscriptionID, field.TypeString)
	}
	if value, ok := euo.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlement.FieldExpiresAt, field.TypeTime, value)
	}
	if euo.mutation.ExpiresAtCleared() {
		_spec.ClearField(entitlement.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := euo.mutation.Cancelled(); ok {
		_spec.SetField(entitlement.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = euo.schemaConfig.Entitlement
	ctx = internal.NewSchemaConfigContext(ctx, euo.schemaConfig)
	_node = &Entitlement{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entitlement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
