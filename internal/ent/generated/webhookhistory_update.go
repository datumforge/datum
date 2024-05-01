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
	"github.com/datumforge/datum/internal/ent/generated/webhookhistory"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// WebhookHistoryUpdate is the builder for updating WebhookHistory entities.
type WebhookHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *WebhookHistoryMutation
}

// Where appends a list predicates to the WebhookHistoryUpdate builder.
func (whu *WebhookHistoryUpdate) Where(ps ...predicate.WebhookHistory) *WebhookHistoryUpdate {
	whu.mutation.Where(ps...)
	return whu
}

// SetUpdatedAt sets the "updated_at" field.
func (whu *WebhookHistoryUpdate) SetUpdatedAt(t time.Time) *WebhookHistoryUpdate {
	whu.mutation.SetUpdatedAt(t)
	return whu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *WebhookHistoryUpdate {
	if t != nil {
		whu.SetUpdatedAt(*t)
	}
	return whu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (whu *WebhookHistoryUpdate) ClearUpdatedAt() *WebhookHistoryUpdate {
	whu.mutation.ClearUpdatedAt()
	return whu
}

// SetUpdatedBy sets the "updated_by" field.
func (whu *WebhookHistoryUpdate) SetUpdatedBy(s string) *WebhookHistoryUpdate {
	whu.mutation.SetUpdatedBy(s)
	return whu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableUpdatedBy(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetUpdatedBy(*s)
	}
	return whu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (whu *WebhookHistoryUpdate) ClearUpdatedBy() *WebhookHistoryUpdate {
	whu.mutation.ClearUpdatedBy()
	return whu
}

// SetDeletedAt sets the "deleted_at" field.
func (whu *WebhookHistoryUpdate) SetDeletedAt(t time.Time) *WebhookHistoryUpdate {
	whu.mutation.SetDeletedAt(t)
	return whu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableDeletedAt(t *time.Time) *WebhookHistoryUpdate {
	if t != nil {
		whu.SetDeletedAt(*t)
	}
	return whu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (whu *WebhookHistoryUpdate) ClearDeletedAt() *WebhookHistoryUpdate {
	whu.mutation.ClearDeletedAt()
	return whu
}

// SetDeletedBy sets the "deleted_by" field.
func (whu *WebhookHistoryUpdate) SetDeletedBy(s string) *WebhookHistoryUpdate {
	whu.mutation.SetDeletedBy(s)
	return whu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableDeletedBy(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetDeletedBy(*s)
	}
	return whu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (whu *WebhookHistoryUpdate) ClearDeletedBy() *WebhookHistoryUpdate {
	whu.mutation.ClearDeletedBy()
	return whu
}

// SetOwnerID sets the "owner_id" field.
func (whu *WebhookHistoryUpdate) SetOwnerID(s string) *WebhookHistoryUpdate {
	whu.mutation.SetOwnerID(s)
	return whu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableOwnerID(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetOwnerID(*s)
	}
	return whu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (whu *WebhookHistoryUpdate) ClearOwnerID() *WebhookHistoryUpdate {
	whu.mutation.ClearOwnerID()
	return whu
}

// SetName sets the "name" field.
func (whu *WebhookHistoryUpdate) SetName(s string) *WebhookHistoryUpdate {
	whu.mutation.SetName(s)
	return whu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableName(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetName(*s)
	}
	return whu
}

// SetDescription sets the "description" field.
func (whu *WebhookHistoryUpdate) SetDescription(s string) *WebhookHistoryUpdate {
	whu.mutation.SetDescription(s)
	return whu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableDescription(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetDescription(*s)
	}
	return whu
}

// ClearDescription clears the value of the "description" field.
func (whu *WebhookHistoryUpdate) ClearDescription() *WebhookHistoryUpdate {
	whu.mutation.ClearDescription()
	return whu
}

// SetDestinationURL sets the "destination_url" field.
func (whu *WebhookHistoryUpdate) SetDestinationURL(s string) *WebhookHistoryUpdate {
	whu.mutation.SetDestinationURL(s)
	return whu
}

// SetNillableDestinationURL sets the "destination_url" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableDestinationURL(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetDestinationURL(*s)
	}
	return whu
}

// SetEnabled sets the "enabled" field.
func (whu *WebhookHistoryUpdate) SetEnabled(b bool) *WebhookHistoryUpdate {
	whu.mutation.SetEnabled(b)
	return whu
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableEnabled(b *bool) *WebhookHistoryUpdate {
	if b != nil {
		whu.SetEnabled(*b)
	}
	return whu
}

// SetCallback sets the "callback" field.
func (whu *WebhookHistoryUpdate) SetCallback(s string) *WebhookHistoryUpdate {
	whu.mutation.SetCallback(s)
	return whu
}

// SetNillableCallback sets the "callback" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableCallback(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetCallback(*s)
	}
	return whu
}

// ClearCallback clears the value of the "callback" field.
func (whu *WebhookHistoryUpdate) ClearCallback() *WebhookHistoryUpdate {
	whu.mutation.ClearCallback()
	return whu
}

// SetExpiresAt sets the "expires_at" field.
func (whu *WebhookHistoryUpdate) SetExpiresAt(t time.Time) *WebhookHistoryUpdate {
	whu.mutation.SetExpiresAt(t)
	return whu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableExpiresAt(t *time.Time) *WebhookHistoryUpdate {
	if t != nil {
		whu.SetExpiresAt(*t)
	}
	return whu
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (whu *WebhookHistoryUpdate) ClearExpiresAt() *WebhookHistoryUpdate {
	whu.mutation.ClearExpiresAt()
	return whu
}

// SetSecret sets the "secret" field.
func (whu *WebhookHistoryUpdate) SetSecret(b []byte) *WebhookHistoryUpdate {
	whu.mutation.SetSecret(b)
	return whu
}

// ClearSecret clears the value of the "secret" field.
func (whu *WebhookHistoryUpdate) ClearSecret() *WebhookHistoryUpdate {
	whu.mutation.ClearSecret()
	return whu
}

// SetFailures sets the "failures" field.
func (whu *WebhookHistoryUpdate) SetFailures(i int) *WebhookHistoryUpdate {
	whu.mutation.ResetFailures()
	whu.mutation.SetFailures(i)
	return whu
}

// SetNillableFailures sets the "failures" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableFailures(i *int) *WebhookHistoryUpdate {
	if i != nil {
		whu.SetFailures(*i)
	}
	return whu
}

// AddFailures adds i to the "failures" field.
func (whu *WebhookHistoryUpdate) AddFailures(i int) *WebhookHistoryUpdate {
	whu.mutation.AddFailures(i)
	return whu
}

// ClearFailures clears the value of the "failures" field.
func (whu *WebhookHistoryUpdate) ClearFailures() *WebhookHistoryUpdate {
	whu.mutation.ClearFailures()
	return whu
}

// SetLastError sets the "last_error" field.
func (whu *WebhookHistoryUpdate) SetLastError(s string) *WebhookHistoryUpdate {
	whu.mutation.SetLastError(s)
	return whu
}

// SetNillableLastError sets the "last_error" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableLastError(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetLastError(*s)
	}
	return whu
}

// ClearLastError clears the value of the "last_error" field.
func (whu *WebhookHistoryUpdate) ClearLastError() *WebhookHistoryUpdate {
	whu.mutation.ClearLastError()
	return whu
}

// SetLastResponse sets the "last_response" field.
func (whu *WebhookHistoryUpdate) SetLastResponse(s string) *WebhookHistoryUpdate {
	whu.mutation.SetLastResponse(s)
	return whu
}

// SetNillableLastResponse sets the "last_response" field if the given value is not nil.
func (whu *WebhookHistoryUpdate) SetNillableLastResponse(s *string) *WebhookHistoryUpdate {
	if s != nil {
		whu.SetLastResponse(*s)
	}
	return whu
}

// ClearLastResponse clears the value of the "last_response" field.
func (whu *WebhookHistoryUpdate) ClearLastResponse() *WebhookHistoryUpdate {
	whu.mutation.ClearLastResponse()
	return whu
}

// Mutation returns the WebhookHistoryMutation object of the builder.
func (whu *WebhookHistoryUpdate) Mutation() *WebhookHistoryMutation {
	return whu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (whu *WebhookHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, whu.sqlSave, whu.mutation, whu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (whu *WebhookHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := whu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (whu *WebhookHistoryUpdate) Exec(ctx context.Context) error {
	_, err := whu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (whu *WebhookHistoryUpdate) ExecX(ctx context.Context) {
	if err := whu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (whu *WebhookHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(webhookhistory.Table, webhookhistory.Columns, sqlgraph.NewFieldSpec(webhookhistory.FieldID, field.TypeString))
	if ps := whu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if whu.mutation.RefCleared() {
		_spec.ClearField(webhookhistory.FieldRef, field.TypeString)
	}
	if whu.mutation.CreatedAtCleared() {
		_spec.ClearField(webhookhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := whu.mutation.UpdatedAt(); ok {
		_spec.SetField(webhookhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if whu.mutation.UpdatedAtCleared() {
		_spec.ClearField(webhookhistory.FieldUpdatedAt, field.TypeTime)
	}
	if whu.mutation.CreatedByCleared() {
		_spec.ClearField(webhookhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := whu.mutation.UpdatedBy(); ok {
		_spec.SetField(webhookhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if whu.mutation.UpdatedByCleared() {
		_spec.ClearField(webhookhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := whu.mutation.DeletedAt(); ok {
		_spec.SetField(webhookhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if whu.mutation.DeletedAtCleared() {
		_spec.ClearField(webhookhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := whu.mutation.DeletedBy(); ok {
		_spec.SetField(webhookhistory.FieldDeletedBy, field.TypeString, value)
	}
	if whu.mutation.DeletedByCleared() {
		_spec.ClearField(webhookhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := whu.mutation.OwnerID(); ok {
		_spec.SetField(webhookhistory.FieldOwnerID, field.TypeString, value)
	}
	if whu.mutation.OwnerIDCleared() {
		_spec.ClearField(webhookhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := whu.mutation.Name(); ok {
		_spec.SetField(webhookhistory.FieldName, field.TypeString, value)
	}
	if value, ok := whu.mutation.Description(); ok {
		_spec.SetField(webhookhistory.FieldDescription, field.TypeString, value)
	}
	if whu.mutation.DescriptionCleared() {
		_spec.ClearField(webhookhistory.FieldDescription, field.TypeString)
	}
	if value, ok := whu.mutation.DestinationURL(); ok {
		_spec.SetField(webhookhistory.FieldDestinationURL, field.TypeString, value)
	}
	if value, ok := whu.mutation.Enabled(); ok {
		_spec.SetField(webhookhistory.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := whu.mutation.Callback(); ok {
		_spec.SetField(webhookhistory.FieldCallback, field.TypeString, value)
	}
	if whu.mutation.CallbackCleared() {
		_spec.ClearField(webhookhistory.FieldCallback, field.TypeString)
	}
	if value, ok := whu.mutation.ExpiresAt(); ok {
		_spec.SetField(webhookhistory.FieldExpiresAt, field.TypeTime, value)
	}
	if whu.mutation.ExpiresAtCleared() {
		_spec.ClearField(webhookhistory.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := whu.mutation.Secret(); ok {
		_spec.SetField(webhookhistory.FieldSecret, field.TypeBytes, value)
	}
	if whu.mutation.SecretCleared() {
		_spec.ClearField(webhookhistory.FieldSecret, field.TypeBytes)
	}
	if value, ok := whu.mutation.Failures(); ok {
		_spec.SetField(webhookhistory.FieldFailures, field.TypeInt, value)
	}
	if value, ok := whu.mutation.AddedFailures(); ok {
		_spec.AddField(webhookhistory.FieldFailures, field.TypeInt, value)
	}
	if whu.mutation.FailuresCleared() {
		_spec.ClearField(webhookhistory.FieldFailures, field.TypeInt)
	}
	if value, ok := whu.mutation.LastError(); ok {
		_spec.SetField(webhookhistory.FieldLastError, field.TypeString, value)
	}
	if whu.mutation.LastErrorCleared() {
		_spec.ClearField(webhookhistory.FieldLastError, field.TypeString)
	}
	if value, ok := whu.mutation.LastResponse(); ok {
		_spec.SetField(webhookhistory.FieldLastResponse, field.TypeString, value)
	}
	if whu.mutation.LastResponseCleared() {
		_spec.ClearField(webhookhistory.FieldLastResponse, field.TypeString)
	}
	_spec.Node.Schema = whu.schemaConfig.WebhookHistory
	ctx = internal.NewSchemaConfigContext(ctx, whu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, whu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webhookhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	whu.mutation.done = true
	return n, nil
}

// WebhookHistoryUpdateOne is the builder for updating a single WebhookHistory entity.
type WebhookHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebhookHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (whuo *WebhookHistoryUpdateOne) SetUpdatedAt(t time.Time) *WebhookHistoryUpdateOne {
	whuo.mutation.SetUpdatedAt(t)
	return whuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *WebhookHistoryUpdateOne {
	if t != nil {
		whuo.SetUpdatedAt(*t)
	}
	return whuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (whuo *WebhookHistoryUpdateOne) ClearUpdatedAt() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearUpdatedAt()
	return whuo
}

// SetUpdatedBy sets the "updated_by" field.
func (whuo *WebhookHistoryUpdateOne) SetUpdatedBy(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetUpdatedBy(s)
	return whuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableUpdatedBy(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetUpdatedBy(*s)
	}
	return whuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (whuo *WebhookHistoryUpdateOne) ClearUpdatedBy() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearUpdatedBy()
	return whuo
}

// SetDeletedAt sets the "deleted_at" field.
func (whuo *WebhookHistoryUpdateOne) SetDeletedAt(t time.Time) *WebhookHistoryUpdateOne {
	whuo.mutation.SetDeletedAt(t)
	return whuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *WebhookHistoryUpdateOne {
	if t != nil {
		whuo.SetDeletedAt(*t)
	}
	return whuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (whuo *WebhookHistoryUpdateOne) ClearDeletedAt() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearDeletedAt()
	return whuo
}

// SetDeletedBy sets the "deleted_by" field.
func (whuo *WebhookHistoryUpdateOne) SetDeletedBy(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetDeletedBy(s)
	return whuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableDeletedBy(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetDeletedBy(*s)
	}
	return whuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (whuo *WebhookHistoryUpdateOne) ClearDeletedBy() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearDeletedBy()
	return whuo
}

// SetOwnerID sets the "owner_id" field.
func (whuo *WebhookHistoryUpdateOne) SetOwnerID(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetOwnerID(s)
	return whuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableOwnerID(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetOwnerID(*s)
	}
	return whuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (whuo *WebhookHistoryUpdateOne) ClearOwnerID() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearOwnerID()
	return whuo
}

// SetName sets the "name" field.
func (whuo *WebhookHistoryUpdateOne) SetName(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetName(s)
	return whuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableName(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetName(*s)
	}
	return whuo
}

// SetDescription sets the "description" field.
func (whuo *WebhookHistoryUpdateOne) SetDescription(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetDescription(s)
	return whuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableDescription(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetDescription(*s)
	}
	return whuo
}

// ClearDescription clears the value of the "description" field.
func (whuo *WebhookHistoryUpdateOne) ClearDescription() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearDescription()
	return whuo
}

// SetDestinationURL sets the "destination_url" field.
func (whuo *WebhookHistoryUpdateOne) SetDestinationURL(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetDestinationURL(s)
	return whuo
}

// SetNillableDestinationURL sets the "destination_url" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableDestinationURL(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetDestinationURL(*s)
	}
	return whuo
}

// SetEnabled sets the "enabled" field.
func (whuo *WebhookHistoryUpdateOne) SetEnabled(b bool) *WebhookHistoryUpdateOne {
	whuo.mutation.SetEnabled(b)
	return whuo
}

// SetNillableEnabled sets the "enabled" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableEnabled(b *bool) *WebhookHistoryUpdateOne {
	if b != nil {
		whuo.SetEnabled(*b)
	}
	return whuo
}

// SetCallback sets the "callback" field.
func (whuo *WebhookHistoryUpdateOne) SetCallback(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetCallback(s)
	return whuo
}

// SetNillableCallback sets the "callback" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableCallback(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetCallback(*s)
	}
	return whuo
}

// ClearCallback clears the value of the "callback" field.
func (whuo *WebhookHistoryUpdateOne) ClearCallback() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearCallback()
	return whuo
}

// SetExpiresAt sets the "expires_at" field.
func (whuo *WebhookHistoryUpdateOne) SetExpiresAt(t time.Time) *WebhookHistoryUpdateOne {
	whuo.mutation.SetExpiresAt(t)
	return whuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableExpiresAt(t *time.Time) *WebhookHistoryUpdateOne {
	if t != nil {
		whuo.SetExpiresAt(*t)
	}
	return whuo
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (whuo *WebhookHistoryUpdateOne) ClearExpiresAt() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearExpiresAt()
	return whuo
}

// SetSecret sets the "secret" field.
func (whuo *WebhookHistoryUpdateOne) SetSecret(b []byte) *WebhookHistoryUpdateOne {
	whuo.mutation.SetSecret(b)
	return whuo
}

// ClearSecret clears the value of the "secret" field.
func (whuo *WebhookHistoryUpdateOne) ClearSecret() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearSecret()
	return whuo
}

// SetFailures sets the "failures" field.
func (whuo *WebhookHistoryUpdateOne) SetFailures(i int) *WebhookHistoryUpdateOne {
	whuo.mutation.ResetFailures()
	whuo.mutation.SetFailures(i)
	return whuo
}

// SetNillableFailures sets the "failures" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableFailures(i *int) *WebhookHistoryUpdateOne {
	if i != nil {
		whuo.SetFailures(*i)
	}
	return whuo
}

// AddFailures adds i to the "failures" field.
func (whuo *WebhookHistoryUpdateOne) AddFailures(i int) *WebhookHistoryUpdateOne {
	whuo.mutation.AddFailures(i)
	return whuo
}

// ClearFailures clears the value of the "failures" field.
func (whuo *WebhookHistoryUpdateOne) ClearFailures() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearFailures()
	return whuo
}

// SetLastError sets the "last_error" field.
func (whuo *WebhookHistoryUpdateOne) SetLastError(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetLastError(s)
	return whuo
}

// SetNillableLastError sets the "last_error" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableLastError(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetLastError(*s)
	}
	return whuo
}

// ClearLastError clears the value of the "last_error" field.
func (whuo *WebhookHistoryUpdateOne) ClearLastError() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearLastError()
	return whuo
}

// SetLastResponse sets the "last_response" field.
func (whuo *WebhookHistoryUpdateOne) SetLastResponse(s string) *WebhookHistoryUpdateOne {
	whuo.mutation.SetLastResponse(s)
	return whuo
}

// SetNillableLastResponse sets the "last_response" field if the given value is not nil.
func (whuo *WebhookHistoryUpdateOne) SetNillableLastResponse(s *string) *WebhookHistoryUpdateOne {
	if s != nil {
		whuo.SetLastResponse(*s)
	}
	return whuo
}

// ClearLastResponse clears the value of the "last_response" field.
func (whuo *WebhookHistoryUpdateOne) ClearLastResponse() *WebhookHistoryUpdateOne {
	whuo.mutation.ClearLastResponse()
	return whuo
}

// Mutation returns the WebhookHistoryMutation object of the builder.
func (whuo *WebhookHistoryUpdateOne) Mutation() *WebhookHistoryMutation {
	return whuo.mutation
}

// Where appends a list predicates to the WebhookHistoryUpdate builder.
func (whuo *WebhookHistoryUpdateOne) Where(ps ...predicate.WebhookHistory) *WebhookHistoryUpdateOne {
	whuo.mutation.Where(ps...)
	return whuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (whuo *WebhookHistoryUpdateOne) Select(field string, fields ...string) *WebhookHistoryUpdateOne {
	whuo.fields = append([]string{field}, fields...)
	return whuo
}

// Save executes the query and returns the updated WebhookHistory entity.
func (whuo *WebhookHistoryUpdateOne) Save(ctx context.Context) (*WebhookHistory, error) {
	return withHooks(ctx, whuo.sqlSave, whuo.mutation, whuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (whuo *WebhookHistoryUpdateOne) SaveX(ctx context.Context) *WebhookHistory {
	node, err := whuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (whuo *WebhookHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := whuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (whuo *WebhookHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := whuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (whuo *WebhookHistoryUpdateOne) sqlSave(ctx context.Context) (_node *WebhookHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(webhookhistory.Table, webhookhistory.Columns, sqlgraph.NewFieldSpec(webhookhistory.FieldID, field.TypeString))
	id, ok := whuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "WebhookHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := whuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webhookhistory.FieldID)
		for _, f := range fields {
			if !webhookhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != webhookhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := whuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if whuo.mutation.RefCleared() {
		_spec.ClearField(webhookhistory.FieldRef, field.TypeString)
	}
	if whuo.mutation.CreatedAtCleared() {
		_spec.ClearField(webhookhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := whuo.mutation.UpdatedAt(); ok {
		_spec.SetField(webhookhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if whuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(webhookhistory.FieldUpdatedAt, field.TypeTime)
	}
	if whuo.mutation.CreatedByCleared() {
		_spec.ClearField(webhookhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := whuo.mutation.UpdatedBy(); ok {
		_spec.SetField(webhookhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if whuo.mutation.UpdatedByCleared() {
		_spec.ClearField(webhookhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := whuo.mutation.DeletedAt(); ok {
		_spec.SetField(webhookhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if whuo.mutation.DeletedAtCleared() {
		_spec.ClearField(webhookhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := whuo.mutation.DeletedBy(); ok {
		_spec.SetField(webhookhistory.FieldDeletedBy, field.TypeString, value)
	}
	if whuo.mutation.DeletedByCleared() {
		_spec.ClearField(webhookhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := whuo.mutation.OwnerID(); ok {
		_spec.SetField(webhookhistory.FieldOwnerID, field.TypeString, value)
	}
	if whuo.mutation.OwnerIDCleared() {
		_spec.ClearField(webhookhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := whuo.mutation.Name(); ok {
		_spec.SetField(webhookhistory.FieldName, field.TypeString, value)
	}
	if value, ok := whuo.mutation.Description(); ok {
		_spec.SetField(webhookhistory.FieldDescription, field.TypeString, value)
	}
	if whuo.mutation.DescriptionCleared() {
		_spec.ClearField(webhookhistory.FieldDescription, field.TypeString)
	}
	if value, ok := whuo.mutation.DestinationURL(); ok {
		_spec.SetField(webhookhistory.FieldDestinationURL, field.TypeString, value)
	}
	if value, ok := whuo.mutation.Enabled(); ok {
		_spec.SetField(webhookhistory.FieldEnabled, field.TypeBool, value)
	}
	if value, ok := whuo.mutation.Callback(); ok {
		_spec.SetField(webhookhistory.FieldCallback, field.TypeString, value)
	}
	if whuo.mutation.CallbackCleared() {
		_spec.ClearField(webhookhistory.FieldCallback, field.TypeString)
	}
	if value, ok := whuo.mutation.ExpiresAt(); ok {
		_spec.SetField(webhookhistory.FieldExpiresAt, field.TypeTime, value)
	}
	if whuo.mutation.ExpiresAtCleared() {
		_spec.ClearField(webhookhistory.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := whuo.mutation.Secret(); ok {
		_spec.SetField(webhookhistory.FieldSecret, field.TypeBytes, value)
	}
	if whuo.mutation.SecretCleared() {
		_spec.ClearField(webhookhistory.FieldSecret, field.TypeBytes)
	}
	if value, ok := whuo.mutation.Failures(); ok {
		_spec.SetField(webhookhistory.FieldFailures, field.TypeInt, value)
	}
	if value, ok := whuo.mutation.AddedFailures(); ok {
		_spec.AddField(webhookhistory.FieldFailures, field.TypeInt, value)
	}
	if whuo.mutation.FailuresCleared() {
		_spec.ClearField(webhookhistory.FieldFailures, field.TypeInt)
	}
	if value, ok := whuo.mutation.LastError(); ok {
		_spec.SetField(webhookhistory.FieldLastError, field.TypeString, value)
	}
	if whuo.mutation.LastErrorCleared() {
		_spec.ClearField(webhookhistory.FieldLastError, field.TypeString)
	}
	if value, ok := whuo.mutation.LastResponse(); ok {
		_spec.SetField(webhookhistory.FieldLastResponse, field.TypeString, value)
	}
	if whuo.mutation.LastResponseCleared() {
		_spec.ClearField(webhookhistory.FieldLastResponse, field.TypeString)
	}
	_spec.Node.Schema = whuo.schemaConfig.WebhookHistory
	ctx = internal.NewSchemaConfigContext(ctx, whuo.schemaConfig)
	_node = &WebhookHistory{config: whuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, whuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{webhookhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	whuo.mutation.done = true
	return _node, nil
}
