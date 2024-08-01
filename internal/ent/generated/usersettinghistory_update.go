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
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/usersettinghistory"
	"github.com/datumforge/datum/pkg/enums"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// UserSettingHistoryUpdate is the builder for updating UserSettingHistory entities.
type UserSettingHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *UserSettingHistoryMutation
}

// Where appends a list predicates to the UserSettingHistoryUpdate builder.
func (ushu *UserSettingHistoryUpdate) Where(ps ...predicate.UserSettingHistory) *UserSettingHistoryUpdate {
	ushu.mutation.Where(ps...)
	return ushu
}

// SetUpdatedAt sets the "updated_at" field.
func (ushu *UserSettingHistoryUpdate) SetUpdatedAt(t time.Time) *UserSettingHistoryUpdate {
	ushu.mutation.SetUpdatedAt(t)
	return ushu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ushu *UserSettingHistoryUpdate) ClearUpdatedAt() *UserSettingHistoryUpdate {
	ushu.mutation.ClearUpdatedAt()
	return ushu
}

// SetUpdatedBy sets the "updated_by" field.
func (ushu *UserSettingHistoryUpdate) SetUpdatedBy(s string) *UserSettingHistoryUpdate {
	ushu.mutation.SetUpdatedBy(s)
	return ushu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableUpdatedBy(s *string) *UserSettingHistoryUpdate {
	if s != nil {
		ushu.SetUpdatedBy(*s)
	}
	return ushu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ushu *UserSettingHistoryUpdate) ClearUpdatedBy() *UserSettingHistoryUpdate {
	ushu.mutation.ClearUpdatedBy()
	return ushu
}

// SetTags sets the "tags" field.
func (ushu *UserSettingHistoryUpdate) SetTags(s []string) *UserSettingHistoryUpdate {
	ushu.mutation.SetTags(s)
	return ushu
}

// AppendTags appends s to the "tags" field.
func (ushu *UserSettingHistoryUpdate) AppendTags(s []string) *UserSettingHistoryUpdate {
	ushu.mutation.AppendTags(s)
	return ushu
}

// ClearTags clears the value of the "tags" field.
func (ushu *UserSettingHistoryUpdate) ClearTags() *UserSettingHistoryUpdate {
	ushu.mutation.ClearTags()
	return ushu
}

// SetDeletedAt sets the "deleted_at" field.
func (ushu *UserSettingHistoryUpdate) SetDeletedAt(t time.Time) *UserSettingHistoryUpdate {
	ushu.mutation.SetDeletedAt(t)
	return ushu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableDeletedAt(t *time.Time) *UserSettingHistoryUpdate {
	if t != nil {
		ushu.SetDeletedAt(*t)
	}
	return ushu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ushu *UserSettingHistoryUpdate) ClearDeletedAt() *UserSettingHistoryUpdate {
	ushu.mutation.ClearDeletedAt()
	return ushu
}

// SetDeletedBy sets the "deleted_by" field.
func (ushu *UserSettingHistoryUpdate) SetDeletedBy(s string) *UserSettingHistoryUpdate {
	ushu.mutation.SetDeletedBy(s)
	return ushu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableDeletedBy(s *string) *UserSettingHistoryUpdate {
	if s != nil {
		ushu.SetDeletedBy(*s)
	}
	return ushu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ushu *UserSettingHistoryUpdate) ClearDeletedBy() *UserSettingHistoryUpdate {
	ushu.mutation.ClearDeletedBy()
	return ushu
}

// SetUserID sets the "user_id" field.
func (ushu *UserSettingHistoryUpdate) SetUserID(s string) *UserSettingHistoryUpdate {
	ushu.mutation.SetUserID(s)
	return ushu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableUserID(s *string) *UserSettingHistoryUpdate {
	if s != nil {
		ushu.SetUserID(*s)
	}
	return ushu
}

// ClearUserID clears the value of the "user_id" field.
func (ushu *UserSettingHistoryUpdate) ClearUserID() *UserSettingHistoryUpdate {
	ushu.mutation.ClearUserID()
	return ushu
}

// SetLocked sets the "locked" field.
func (ushu *UserSettingHistoryUpdate) SetLocked(b bool) *UserSettingHistoryUpdate {
	ushu.mutation.SetLocked(b)
	return ushu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableLocked(b *bool) *UserSettingHistoryUpdate {
	if b != nil {
		ushu.SetLocked(*b)
	}
	return ushu
}

// SetSilencedAt sets the "silenced_at" field.
func (ushu *UserSettingHistoryUpdate) SetSilencedAt(t time.Time) *UserSettingHistoryUpdate {
	ushu.mutation.SetSilencedAt(t)
	return ushu
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableSilencedAt(t *time.Time) *UserSettingHistoryUpdate {
	if t != nil {
		ushu.SetSilencedAt(*t)
	}
	return ushu
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (ushu *UserSettingHistoryUpdate) ClearSilencedAt() *UserSettingHistoryUpdate {
	ushu.mutation.ClearSilencedAt()
	return ushu
}

// SetSuspendedAt sets the "suspended_at" field.
func (ushu *UserSettingHistoryUpdate) SetSuspendedAt(t time.Time) *UserSettingHistoryUpdate {
	ushu.mutation.SetSuspendedAt(t)
	return ushu
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableSuspendedAt(t *time.Time) *UserSettingHistoryUpdate {
	if t != nil {
		ushu.SetSuspendedAt(*t)
	}
	return ushu
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (ushu *UserSettingHistoryUpdate) ClearSuspendedAt() *UserSettingHistoryUpdate {
	ushu.mutation.ClearSuspendedAt()
	return ushu
}

// SetStatus sets the "status" field.
func (ushu *UserSettingHistoryUpdate) SetStatus(es enums.UserStatus) *UserSettingHistoryUpdate {
	ushu.mutation.SetStatus(es)
	return ushu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableStatus(es *enums.UserStatus) *UserSettingHistoryUpdate {
	if es != nil {
		ushu.SetStatus(*es)
	}
	return ushu
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (ushu *UserSettingHistoryUpdate) SetEmailConfirmed(b bool) *UserSettingHistoryUpdate {
	ushu.mutation.SetEmailConfirmed(b)
	return ushu
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableEmailConfirmed(b *bool) *UserSettingHistoryUpdate {
	if b != nil {
		ushu.SetEmailConfirmed(*b)
	}
	return ushu
}

// SetIsWebauthnAllowed sets the "is_webauthn_allowed" field.
func (ushu *UserSettingHistoryUpdate) SetIsWebauthnAllowed(b bool) *UserSettingHistoryUpdate {
	ushu.mutation.SetIsWebauthnAllowed(b)
	return ushu
}

// SetNillableIsWebauthnAllowed sets the "is_webauthn_allowed" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableIsWebauthnAllowed(b *bool) *UserSettingHistoryUpdate {
	if b != nil {
		ushu.SetIsWebauthnAllowed(*b)
	}
	return ushu
}

// ClearIsWebauthnAllowed clears the value of the "is_webauthn_allowed" field.
func (ushu *UserSettingHistoryUpdate) ClearIsWebauthnAllowed() *UserSettingHistoryUpdate {
	ushu.mutation.ClearIsWebauthnAllowed()
	return ushu
}

// SetIsTfaEnabled sets the "is_tfa_enabled" field.
func (ushu *UserSettingHistoryUpdate) SetIsTfaEnabled(b bool) *UserSettingHistoryUpdate {
	ushu.mutation.SetIsTfaEnabled(b)
	return ushu
}

// SetNillableIsTfaEnabled sets the "is_tfa_enabled" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillableIsTfaEnabled(b *bool) *UserSettingHistoryUpdate {
	if b != nil {
		ushu.SetIsTfaEnabled(*b)
	}
	return ushu
}

// ClearIsTfaEnabled clears the value of the "is_tfa_enabled" field.
func (ushu *UserSettingHistoryUpdate) ClearIsTfaEnabled() *UserSettingHistoryUpdate {
	ushu.mutation.ClearIsTfaEnabled()
	return ushu
}

// SetPhoneNumber sets the "phone_number" field.
func (ushu *UserSettingHistoryUpdate) SetPhoneNumber(s string) *UserSettingHistoryUpdate {
	ushu.mutation.SetPhoneNumber(s)
	return ushu
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ushu *UserSettingHistoryUpdate) SetNillablePhoneNumber(s *string) *UserSettingHistoryUpdate {
	if s != nil {
		ushu.SetPhoneNumber(*s)
	}
	return ushu
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (ushu *UserSettingHistoryUpdate) ClearPhoneNumber() *UserSettingHistoryUpdate {
	ushu.mutation.ClearPhoneNumber()
	return ushu
}

// Mutation returns the UserSettingHistoryMutation object of the builder.
func (ushu *UserSettingHistoryUpdate) Mutation() *UserSettingHistoryMutation {
	return ushu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ushu *UserSettingHistoryUpdate) Save(ctx context.Context) (int, error) {
	ushu.defaults()
	return withHooks(ctx, ushu.sqlSave, ushu.mutation, ushu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ushu *UserSettingHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ushu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ushu *UserSettingHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ushu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ushu *UserSettingHistoryUpdate) ExecX(ctx context.Context) {
	if err := ushu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ushu *UserSettingHistoryUpdate) defaults() {
	if _, ok := ushu.mutation.UpdatedAt(); !ok && !ushu.mutation.UpdatedAtCleared() {
		v := usersettinghistory.UpdateDefaultUpdatedAt()
		ushu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ushu *UserSettingHistoryUpdate) check() error {
	if v, ok := ushu.mutation.Status(); ok {
		if err := usersettinghistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSettingHistory.status": %w`, err)}
		}
	}
	return nil
}

func (ushu *UserSettingHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ushu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersettinghistory.Table, usersettinghistory.Columns, sqlgraph.NewFieldSpec(usersettinghistory.FieldID, field.TypeString))
	if ps := ushu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ushu.mutation.RefCleared() {
		_spec.ClearField(usersettinghistory.FieldRef, field.TypeString)
	}
	if ushu.mutation.CreatedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ushu.mutation.UpdatedAt(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ushu.mutation.UpdatedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if ushu.mutation.CreatedByCleared() {
		_spec.ClearField(usersettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ushu.mutation.UpdatedBy(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ushu.mutation.UpdatedByCleared() {
		_spec.ClearField(usersettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ushu.mutation.Tags(); ok {
		_spec.SetField(usersettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ushu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, usersettinghistory.FieldTags, value)
		})
	}
	if ushu.mutation.TagsCleared() {
		_spec.ClearField(usersettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ushu.mutation.DeletedAt(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ushu.mutation.DeletedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ushu.mutation.DeletedBy(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if ushu.mutation.DeletedByCleared() {
		_spec.ClearField(usersettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ushu.mutation.UserID(); ok {
		_spec.SetField(usersettinghistory.FieldUserID, field.TypeString, value)
	}
	if ushu.mutation.UserIDCleared() {
		_spec.ClearField(usersettinghistory.FieldUserID, field.TypeString)
	}
	if value, ok := ushu.mutation.Locked(); ok {
		_spec.SetField(usersettinghistory.FieldLocked, field.TypeBool, value)
	}
	if value, ok := ushu.mutation.SilencedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSilencedAt, field.TypeTime, value)
	}
	if ushu.mutation.SilencedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := ushu.mutation.SuspendedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSuspendedAt, field.TypeTime, value)
	}
	if ushu.mutation.SuspendedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := ushu.mutation.Status(); ok {
		_spec.SetField(usersettinghistory.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ushu.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersettinghistory.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := ushu.mutation.IsWebauthnAllowed(); ok {
		_spec.SetField(usersettinghistory.FieldIsWebauthnAllowed, field.TypeBool, value)
	}
	if ushu.mutation.IsWebauthnAllowedCleared() {
		_spec.ClearField(usersettinghistory.FieldIsWebauthnAllowed, field.TypeBool)
	}
	if value, ok := ushu.mutation.IsTfaEnabled(); ok {
		_spec.SetField(usersettinghistory.FieldIsTfaEnabled, field.TypeBool, value)
	}
	if ushu.mutation.IsTfaEnabledCleared() {
		_spec.ClearField(usersettinghistory.FieldIsTfaEnabled, field.TypeBool)
	}
	if value, ok := ushu.mutation.PhoneNumber(); ok {
		_spec.SetField(usersettinghistory.FieldPhoneNumber, field.TypeString, value)
	}
	if ushu.mutation.PhoneNumberCleared() {
		_spec.ClearField(usersettinghistory.FieldPhoneNumber, field.TypeString)
	}
	_spec.Node.Schema = ushu.schemaConfig.UserSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, ushu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ushu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ushu.mutation.done = true
	return n, nil
}

// UserSettingHistoryUpdateOne is the builder for updating a single UserSettingHistory entity.
type UserSettingHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSettingHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ushuo *UserSettingHistoryUpdateOne) SetUpdatedAt(t time.Time) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetUpdatedAt(t)
	return ushuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearUpdatedAt() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearUpdatedAt()
	return ushuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ushuo *UserSettingHistoryUpdateOne) SetUpdatedBy(s string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetUpdatedBy(s)
	return ushuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableUpdatedBy(s *string) *UserSettingHistoryUpdateOne {
	if s != nil {
		ushuo.SetUpdatedBy(*s)
	}
	return ushuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearUpdatedBy() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearUpdatedBy()
	return ushuo
}

// SetTags sets the "tags" field.
func (ushuo *UserSettingHistoryUpdateOne) SetTags(s []string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetTags(s)
	return ushuo
}

// AppendTags appends s to the "tags" field.
func (ushuo *UserSettingHistoryUpdateOne) AppendTags(s []string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.AppendTags(s)
	return ushuo
}

// ClearTags clears the value of the "tags" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearTags() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearTags()
	return ushuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ushuo *UserSettingHistoryUpdateOne) SetDeletedAt(t time.Time) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetDeletedAt(t)
	return ushuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *UserSettingHistoryUpdateOne {
	if t != nil {
		ushuo.SetDeletedAt(*t)
	}
	return ushuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearDeletedAt() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearDeletedAt()
	return ushuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ushuo *UserSettingHistoryUpdateOne) SetDeletedBy(s string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetDeletedBy(s)
	return ushuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableDeletedBy(s *string) *UserSettingHistoryUpdateOne {
	if s != nil {
		ushuo.SetDeletedBy(*s)
	}
	return ushuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearDeletedBy() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearDeletedBy()
	return ushuo
}

// SetUserID sets the "user_id" field.
func (ushuo *UserSettingHistoryUpdateOne) SetUserID(s string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetUserID(s)
	return ushuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableUserID(s *string) *UserSettingHistoryUpdateOne {
	if s != nil {
		ushuo.SetUserID(*s)
	}
	return ushuo
}

// ClearUserID clears the value of the "user_id" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearUserID() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearUserID()
	return ushuo
}

// SetLocked sets the "locked" field.
func (ushuo *UserSettingHistoryUpdateOne) SetLocked(b bool) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetLocked(b)
	return ushuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableLocked(b *bool) *UserSettingHistoryUpdateOne {
	if b != nil {
		ushuo.SetLocked(*b)
	}
	return ushuo
}

// SetSilencedAt sets the "silenced_at" field.
func (ushuo *UserSettingHistoryUpdateOne) SetSilencedAt(t time.Time) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetSilencedAt(t)
	return ushuo
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableSilencedAt(t *time.Time) *UserSettingHistoryUpdateOne {
	if t != nil {
		ushuo.SetSilencedAt(*t)
	}
	return ushuo
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearSilencedAt() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearSilencedAt()
	return ushuo
}

// SetSuspendedAt sets the "suspended_at" field.
func (ushuo *UserSettingHistoryUpdateOne) SetSuspendedAt(t time.Time) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetSuspendedAt(t)
	return ushuo
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableSuspendedAt(t *time.Time) *UserSettingHistoryUpdateOne {
	if t != nil {
		ushuo.SetSuspendedAt(*t)
	}
	return ushuo
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearSuspendedAt() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearSuspendedAt()
	return ushuo
}

// SetStatus sets the "status" field.
func (ushuo *UserSettingHistoryUpdateOne) SetStatus(es enums.UserStatus) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetStatus(es)
	return ushuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableStatus(es *enums.UserStatus) *UserSettingHistoryUpdateOne {
	if es != nil {
		ushuo.SetStatus(*es)
	}
	return ushuo
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (ushuo *UserSettingHistoryUpdateOne) SetEmailConfirmed(b bool) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetEmailConfirmed(b)
	return ushuo
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableEmailConfirmed(b *bool) *UserSettingHistoryUpdateOne {
	if b != nil {
		ushuo.SetEmailConfirmed(*b)
	}
	return ushuo
}

// SetIsWebauthnAllowed sets the "is_webauthn_allowed" field.
func (ushuo *UserSettingHistoryUpdateOne) SetIsWebauthnAllowed(b bool) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetIsWebauthnAllowed(b)
	return ushuo
}

// SetNillableIsWebauthnAllowed sets the "is_webauthn_allowed" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableIsWebauthnAllowed(b *bool) *UserSettingHistoryUpdateOne {
	if b != nil {
		ushuo.SetIsWebauthnAllowed(*b)
	}
	return ushuo
}

// ClearIsWebauthnAllowed clears the value of the "is_webauthn_allowed" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearIsWebauthnAllowed() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearIsWebauthnAllowed()
	return ushuo
}

// SetIsTfaEnabled sets the "is_tfa_enabled" field.
func (ushuo *UserSettingHistoryUpdateOne) SetIsTfaEnabled(b bool) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetIsTfaEnabled(b)
	return ushuo
}

// SetNillableIsTfaEnabled sets the "is_tfa_enabled" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillableIsTfaEnabled(b *bool) *UserSettingHistoryUpdateOne {
	if b != nil {
		ushuo.SetIsTfaEnabled(*b)
	}
	return ushuo
}

// ClearIsTfaEnabled clears the value of the "is_tfa_enabled" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearIsTfaEnabled() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearIsTfaEnabled()
	return ushuo
}

// SetPhoneNumber sets the "phone_number" field.
func (ushuo *UserSettingHistoryUpdateOne) SetPhoneNumber(s string) *UserSettingHistoryUpdateOne {
	ushuo.mutation.SetPhoneNumber(s)
	return ushuo
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (ushuo *UserSettingHistoryUpdateOne) SetNillablePhoneNumber(s *string) *UserSettingHistoryUpdateOne {
	if s != nil {
		ushuo.SetPhoneNumber(*s)
	}
	return ushuo
}

// ClearPhoneNumber clears the value of the "phone_number" field.
func (ushuo *UserSettingHistoryUpdateOne) ClearPhoneNumber() *UserSettingHistoryUpdateOne {
	ushuo.mutation.ClearPhoneNumber()
	return ushuo
}

// Mutation returns the UserSettingHistoryMutation object of the builder.
func (ushuo *UserSettingHistoryUpdateOne) Mutation() *UserSettingHistoryMutation {
	return ushuo.mutation
}

// Where appends a list predicates to the UserSettingHistoryUpdate builder.
func (ushuo *UserSettingHistoryUpdateOne) Where(ps ...predicate.UserSettingHistory) *UserSettingHistoryUpdateOne {
	ushuo.mutation.Where(ps...)
	return ushuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ushuo *UserSettingHistoryUpdateOne) Select(field string, fields ...string) *UserSettingHistoryUpdateOne {
	ushuo.fields = append([]string{field}, fields...)
	return ushuo
}

// Save executes the query and returns the updated UserSettingHistory entity.
func (ushuo *UserSettingHistoryUpdateOne) Save(ctx context.Context) (*UserSettingHistory, error) {
	ushuo.defaults()
	return withHooks(ctx, ushuo.sqlSave, ushuo.mutation, ushuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ushuo *UserSettingHistoryUpdateOne) SaveX(ctx context.Context) *UserSettingHistory {
	node, err := ushuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ushuo *UserSettingHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ushuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ushuo *UserSettingHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ushuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ushuo *UserSettingHistoryUpdateOne) defaults() {
	if _, ok := ushuo.mutation.UpdatedAt(); !ok && !ushuo.mutation.UpdatedAtCleared() {
		v := usersettinghistory.UpdateDefaultUpdatedAt()
		ushuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ushuo *UserSettingHistoryUpdateOne) check() error {
	if v, ok := ushuo.mutation.Status(); ok {
		if err := usersettinghistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSettingHistory.status": %w`, err)}
		}
	}
	return nil
}

func (ushuo *UserSettingHistoryUpdateOne) sqlSave(ctx context.Context) (_node *UserSettingHistory, err error) {
	if err := ushuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersettinghistory.Table, usersettinghistory.Columns, sqlgraph.NewFieldSpec(usersettinghistory.FieldID, field.TypeString))
	id, ok := ushuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "UserSettingHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ushuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersettinghistory.FieldID)
		for _, f := range fields {
			if !usersettinghistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != usersettinghistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ushuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ushuo.mutation.RefCleared() {
		_spec.ClearField(usersettinghistory.FieldRef, field.TypeString)
	}
	if ushuo.mutation.CreatedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ushuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ushuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if ushuo.mutation.CreatedByCleared() {
		_spec.ClearField(usersettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ushuo.mutation.UpdatedBy(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ushuo.mutation.UpdatedByCleared() {
		_spec.ClearField(usersettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ushuo.mutation.Tags(); ok {
		_spec.SetField(usersettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ushuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, usersettinghistory.FieldTags, value)
		})
	}
	if ushuo.mutation.TagsCleared() {
		_spec.ClearField(usersettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ushuo.mutation.DeletedAt(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ushuo.mutation.DeletedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ushuo.mutation.DeletedBy(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if ushuo.mutation.DeletedByCleared() {
		_spec.ClearField(usersettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ushuo.mutation.UserID(); ok {
		_spec.SetField(usersettinghistory.FieldUserID, field.TypeString, value)
	}
	if ushuo.mutation.UserIDCleared() {
		_spec.ClearField(usersettinghistory.FieldUserID, field.TypeString)
	}
	if value, ok := ushuo.mutation.Locked(); ok {
		_spec.SetField(usersettinghistory.FieldLocked, field.TypeBool, value)
	}
	if value, ok := ushuo.mutation.SilencedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSilencedAt, field.TypeTime, value)
	}
	if ushuo.mutation.SilencedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := ushuo.mutation.SuspendedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSuspendedAt, field.TypeTime, value)
	}
	if ushuo.mutation.SuspendedAtCleared() {
		_spec.ClearField(usersettinghistory.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := ushuo.mutation.Status(); ok {
		_spec.SetField(usersettinghistory.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := ushuo.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersettinghistory.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := ushuo.mutation.IsWebauthnAllowed(); ok {
		_spec.SetField(usersettinghistory.FieldIsWebauthnAllowed, field.TypeBool, value)
	}
	if ushuo.mutation.IsWebauthnAllowedCleared() {
		_spec.ClearField(usersettinghistory.FieldIsWebauthnAllowed, field.TypeBool)
	}
	if value, ok := ushuo.mutation.IsTfaEnabled(); ok {
		_spec.SetField(usersettinghistory.FieldIsTfaEnabled, field.TypeBool, value)
	}
	if ushuo.mutation.IsTfaEnabledCleared() {
		_spec.ClearField(usersettinghistory.FieldIsTfaEnabled, field.TypeBool)
	}
	if value, ok := ushuo.mutation.PhoneNumber(); ok {
		_spec.SetField(usersettinghistory.FieldPhoneNumber, field.TypeString, value)
	}
	if ushuo.mutation.PhoneNumberCleared() {
		_spec.ClearField(usersettinghistory.FieldPhoneNumber, field.TypeString)
	}
	_spec.Node.Schema = ushuo.schemaConfig.UserSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, ushuo.schemaConfig)
	_node = &UserSettingHistory{config: ushuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ushuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ushuo.mutation.done = true
	return _node, nil
}
