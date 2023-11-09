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
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/taco"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TacoUpdate is the builder for updating Taco entities.
type TacoUpdate struct {
	config
	hooks    []Hook
	mutation *TacoMutation
}

// Where appends a list predicates to the TacoUpdate builder.
func (tu *TacoUpdate) Where(ps ...predicate.Taco) *TacoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TacoUpdate) SetUpdatedAt(t time.Time) *TacoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetCreatedBy sets the "created_by" field.
func (tu *TacoUpdate) SetCreatedBy(s string) *TacoUpdate {
	tu.mutation.SetCreatedBy(s)
	return tu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableCreatedBy(s *string) *TacoUpdate {
	if s != nil {
		tu.SetCreatedBy(*s)
	}
	return tu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (tu *TacoUpdate) ClearCreatedBy() *TacoUpdate {
	tu.mutation.ClearCreatedBy()
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TacoUpdate) SetUpdatedBy(s string) *TacoUpdate {
	tu.mutation.SetUpdatedBy(s)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableUpdatedBy(s *string) *TacoUpdate {
	if s != nil {
		tu.SetUpdatedBy(*s)
	}
	return tu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tu *TacoUpdate) ClearUpdatedBy() *TacoUpdate {
	tu.mutation.ClearUpdatedBy()
	return tu
}

// SetTier sets the "tier" field.
func (tu *TacoUpdate) SetTier(t taco.Tier) *TacoUpdate {
	tu.mutation.SetTier(t)
	return tu
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableTier(t *taco.Tier) *TacoUpdate {
	if t != nil {
		tu.SetTier(*t)
	}
	return tu
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (tu *TacoUpdate) SetStripeCustomerID(s string) *TacoUpdate {
	tu.mutation.SetStripeCustomerID(s)
	return tu
}

// SetNillableStripeCustomerID sets the "stripe_customer_id" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableStripeCustomerID(s *string) *TacoUpdate {
	if s != nil {
		tu.SetStripeCustomerID(*s)
	}
	return tu
}

// ClearStripeCustomerID clears the value of the "stripe_customer_id" field.
func (tu *TacoUpdate) ClearStripeCustomerID() *TacoUpdate {
	tu.mutation.ClearStripeCustomerID()
	return tu
}

// SetStripeSubscriptionID sets the "stripe_subscription_id" field.
func (tu *TacoUpdate) SetStripeSubscriptionID(s string) *TacoUpdate {
	tu.mutation.SetStripeSubscriptionID(s)
	return tu
}

// SetNillableStripeSubscriptionID sets the "stripe_subscription_id" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableStripeSubscriptionID(s *string) *TacoUpdate {
	if s != nil {
		tu.SetStripeSubscriptionID(*s)
	}
	return tu
}

// ClearStripeSubscriptionID clears the value of the "stripe_subscription_id" field.
func (tu *TacoUpdate) ClearStripeSubscriptionID() *TacoUpdate {
	tu.mutation.ClearStripeSubscriptionID()
	return tu
}

// SetExpiresAt sets the "expires_at" field.
func (tu *TacoUpdate) SetExpiresAt(t time.Time) *TacoUpdate {
	tu.mutation.SetExpiresAt(t)
	return tu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableExpiresAt(t *time.Time) *TacoUpdate {
	if t != nil {
		tu.SetExpiresAt(*t)
	}
	return tu
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (tu *TacoUpdate) ClearExpiresAt() *TacoUpdate {
	tu.mutation.ClearExpiresAt()
	return tu
}

// SetCancelled sets the "cancelled" field.
func (tu *TacoUpdate) SetCancelled(b bool) *TacoUpdate {
	tu.mutation.SetCancelled(b)
	return tu
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (tu *TacoUpdate) SetNillableCancelled(b *bool) *TacoUpdate {
	if b != nil {
		tu.SetCancelled(*b)
	}
	return tu
}

// Mutation returns the TacoMutation object of the builder.
func (tu *TacoUpdate) Mutation() *TacoMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TacoUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TacoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TacoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TacoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TacoUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if taco.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized taco.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := taco.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TacoUpdate) check() error {
	if v, ok := tu.mutation.Tier(); ok {
		if err := taco.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "Taco.tier": %w`, err)}
		}
	}
	return nil
}

func (tu *TacoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(taco.Table, taco.Columns, sqlgraph.NewFieldSpec(taco.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(taco.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CreatedBy(); ok {
		_spec.SetField(taco.FieldCreatedBy, field.TypeString, value)
	}
	if tu.mutation.CreatedByCleared() {
		_spec.ClearField(taco.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(taco.FieldUpdatedBy, field.TypeString, value)
	}
	if tu.mutation.UpdatedByCleared() {
		_spec.ClearField(taco.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.Tier(); ok {
		_spec.SetField(taco.FieldTier, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.StripeCustomerID(); ok {
		_spec.SetField(taco.FieldStripeCustomerID, field.TypeString, value)
	}
	if tu.mutation.StripeCustomerIDCleared() {
		_spec.ClearField(taco.FieldStripeCustomerID, field.TypeString)
	}
	if value, ok := tu.mutation.StripeSubscriptionID(); ok {
		_spec.SetField(taco.FieldStripeSubscriptionID, field.TypeString, value)
	}
	if tu.mutation.StripeSubscriptionIDCleared() {
		_spec.ClearField(taco.FieldStripeSubscriptionID, field.TypeString)
	}
	if value, ok := tu.mutation.ExpiresAt(); ok {
		_spec.SetField(taco.FieldExpiresAt, field.TypeTime, value)
	}
	if tu.mutation.ExpiresAtCleared() {
		_spec.ClearField(taco.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Cancelled(); ok {
		_spec.SetField(taco.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = tu.schemaConfig.Taco
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taco.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TacoUpdateOne is the builder for updating a single Taco entity.
type TacoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TacoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TacoUpdateOne) SetUpdatedAt(t time.Time) *TacoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetCreatedBy sets the "created_by" field.
func (tuo *TacoUpdateOne) SetCreatedBy(s string) *TacoUpdateOne {
	tuo.mutation.SetCreatedBy(s)
	return tuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableCreatedBy(s *string) *TacoUpdateOne {
	if s != nil {
		tuo.SetCreatedBy(*s)
	}
	return tuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (tuo *TacoUpdateOne) ClearCreatedBy() *TacoUpdateOne {
	tuo.mutation.ClearCreatedBy()
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TacoUpdateOne) SetUpdatedBy(s string) *TacoUpdateOne {
	tuo.mutation.SetUpdatedBy(s)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableUpdatedBy(s *string) *TacoUpdateOne {
	if s != nil {
		tuo.SetUpdatedBy(*s)
	}
	return tuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tuo *TacoUpdateOne) ClearUpdatedBy() *TacoUpdateOne {
	tuo.mutation.ClearUpdatedBy()
	return tuo
}

// SetTier sets the "tier" field.
func (tuo *TacoUpdateOne) SetTier(t taco.Tier) *TacoUpdateOne {
	tuo.mutation.SetTier(t)
	return tuo
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableTier(t *taco.Tier) *TacoUpdateOne {
	if t != nil {
		tuo.SetTier(*t)
	}
	return tuo
}

// SetStripeCustomerID sets the "stripe_customer_id" field.
func (tuo *TacoUpdateOne) SetStripeCustomerID(s string) *TacoUpdateOne {
	tuo.mutation.SetStripeCustomerID(s)
	return tuo
}

// SetNillableStripeCustomerID sets the "stripe_customer_id" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableStripeCustomerID(s *string) *TacoUpdateOne {
	if s != nil {
		tuo.SetStripeCustomerID(*s)
	}
	return tuo
}

// ClearStripeCustomerID clears the value of the "stripe_customer_id" field.
func (tuo *TacoUpdateOne) ClearStripeCustomerID() *TacoUpdateOne {
	tuo.mutation.ClearStripeCustomerID()
	return tuo
}

// SetStripeSubscriptionID sets the "stripe_subscription_id" field.
func (tuo *TacoUpdateOne) SetStripeSubscriptionID(s string) *TacoUpdateOne {
	tuo.mutation.SetStripeSubscriptionID(s)
	return tuo
}

// SetNillableStripeSubscriptionID sets the "stripe_subscription_id" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableStripeSubscriptionID(s *string) *TacoUpdateOne {
	if s != nil {
		tuo.SetStripeSubscriptionID(*s)
	}
	return tuo
}

// ClearStripeSubscriptionID clears the value of the "stripe_subscription_id" field.
func (tuo *TacoUpdateOne) ClearStripeSubscriptionID() *TacoUpdateOne {
	tuo.mutation.ClearStripeSubscriptionID()
	return tuo
}

// SetExpiresAt sets the "expires_at" field.
func (tuo *TacoUpdateOne) SetExpiresAt(t time.Time) *TacoUpdateOne {
	tuo.mutation.SetExpiresAt(t)
	return tuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableExpiresAt(t *time.Time) *TacoUpdateOne {
	if t != nil {
		tuo.SetExpiresAt(*t)
	}
	return tuo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (tuo *TacoUpdateOne) ClearExpiresAt() *TacoUpdateOne {
	tuo.mutation.ClearExpiresAt()
	return tuo
}

// SetCancelled sets the "cancelled" field.
func (tuo *TacoUpdateOne) SetCancelled(b bool) *TacoUpdateOne {
	tuo.mutation.SetCancelled(b)
	return tuo
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (tuo *TacoUpdateOne) SetNillableCancelled(b *bool) *TacoUpdateOne {
	if b != nil {
		tuo.SetCancelled(*b)
	}
	return tuo
}

// Mutation returns the TacoMutation object of the builder.
func (tuo *TacoUpdateOne) Mutation() *TacoMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TacoUpdate builder.
func (tuo *TacoUpdateOne) Where(ps ...predicate.Taco) *TacoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TacoUpdateOne) Select(field string, fields ...string) *TacoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Taco entity.
func (tuo *TacoUpdateOne) Save(ctx context.Context) (*Taco, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TacoUpdateOne) SaveX(ctx context.Context) *Taco {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TacoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TacoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TacoUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if taco.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized taco.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := taco.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TacoUpdateOne) check() error {
	if v, ok := tuo.mutation.Tier(); ok {
		if err := taco.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "Taco.tier": %w`, err)}
		}
	}
	return nil
}

func (tuo *TacoUpdateOne) sqlSave(ctx context.Context) (_node *Taco, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(taco.Table, taco.Columns, sqlgraph.NewFieldSpec(taco.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Taco.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, taco.FieldID)
		for _, f := range fields {
			if !taco.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != taco.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(taco.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CreatedBy(); ok {
		_spec.SetField(taco.FieldCreatedBy, field.TypeString, value)
	}
	if tuo.mutation.CreatedByCleared() {
		_spec.ClearField(taco.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(taco.FieldUpdatedBy, field.TypeString, value)
	}
	if tuo.mutation.UpdatedByCleared() {
		_spec.ClearField(taco.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.Tier(); ok {
		_spec.SetField(taco.FieldTier, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.StripeCustomerID(); ok {
		_spec.SetField(taco.FieldStripeCustomerID, field.TypeString, value)
	}
	if tuo.mutation.StripeCustomerIDCleared() {
		_spec.ClearField(taco.FieldStripeCustomerID, field.TypeString)
	}
	if value, ok := tuo.mutation.StripeSubscriptionID(); ok {
		_spec.SetField(taco.FieldStripeSubscriptionID, field.TypeString, value)
	}
	if tuo.mutation.StripeSubscriptionIDCleared() {
		_spec.ClearField(taco.FieldStripeSubscriptionID, field.TypeString)
	}
	if value, ok := tuo.mutation.ExpiresAt(); ok {
		_spec.SetField(taco.FieldExpiresAt, field.TypeTime, value)
	}
	if tuo.mutation.ExpiresAtCleared() {
		_spec.ClearField(taco.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Cancelled(); ok {
		_spec.SetField(taco.FieldCancelled, field.TypeBool, value)
	}
	_spec.Node.Schema = tuo.schemaConfig.Taco
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_node = &Taco{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{taco.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
