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
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// UserSettingUpdate is the builder for updating UserSetting entities.
type UserSettingUpdate struct {
	config
	hooks    []Hook
	mutation *UserSettingMutation
}

// Where appends a list predicates to the UserSettingUpdate builder.
func (usu *UserSettingUpdate) Where(ps ...predicate.UserSetting) *UserSettingUpdate {
	usu.mutation.Where(ps...)
	return usu
}

// SetUpdatedAt sets the "updated_at" field.
func (usu *UserSettingUpdate) SetUpdatedAt(t time.Time) *UserSettingUpdate {
	usu.mutation.SetUpdatedAt(t)
	return usu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (usu *UserSettingUpdate) ClearUpdatedAt() *UserSettingUpdate {
	usu.mutation.ClearUpdatedAt()
	return usu
}

// SetUpdatedBy sets the "updated_by" field.
func (usu *UserSettingUpdate) SetUpdatedBy(s string) *UserSettingUpdate {
	usu.mutation.SetUpdatedBy(s)
	return usu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableUpdatedBy(s *string) *UserSettingUpdate {
	if s != nil {
		usu.SetUpdatedBy(*s)
	}
	return usu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (usu *UserSettingUpdate) ClearUpdatedBy() *UserSettingUpdate {
	usu.mutation.ClearUpdatedBy()
	return usu
}

// SetDeletedAt sets the "deleted_at" field.
func (usu *UserSettingUpdate) SetDeletedAt(t time.Time) *UserSettingUpdate {
	usu.mutation.SetDeletedAt(t)
	return usu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableDeletedAt(t *time.Time) *UserSettingUpdate {
	if t != nil {
		usu.SetDeletedAt(*t)
	}
	return usu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (usu *UserSettingUpdate) ClearDeletedAt() *UserSettingUpdate {
	usu.mutation.ClearDeletedAt()
	return usu
}

// SetDeletedBy sets the "deleted_by" field.
func (usu *UserSettingUpdate) SetDeletedBy(s string) *UserSettingUpdate {
	usu.mutation.SetDeletedBy(s)
	return usu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableDeletedBy(s *string) *UserSettingUpdate {
	if s != nil {
		usu.SetDeletedBy(*s)
	}
	return usu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (usu *UserSettingUpdate) ClearDeletedBy() *UserSettingUpdate {
	usu.mutation.ClearDeletedBy()
	return usu
}

// SetLocked sets the "locked" field.
func (usu *UserSettingUpdate) SetLocked(b bool) *UserSettingUpdate {
	usu.mutation.SetLocked(b)
	return usu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableLocked(b *bool) *UserSettingUpdate {
	if b != nil {
		usu.SetLocked(*b)
	}
	return usu
}

// SetSilencedAt sets the "silenced_at" field.
func (usu *UserSettingUpdate) SetSilencedAt(t time.Time) *UserSettingUpdate {
	usu.mutation.SetSilencedAt(t)
	return usu
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableSilencedAt(t *time.Time) *UserSettingUpdate {
	if t != nil {
		usu.SetSilencedAt(*t)
	}
	return usu
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (usu *UserSettingUpdate) ClearSilencedAt() *UserSettingUpdate {
	usu.mutation.ClearSilencedAt()
	return usu
}

// SetSuspendedAt sets the "suspended_at" field.
func (usu *UserSettingUpdate) SetSuspendedAt(t time.Time) *UserSettingUpdate {
	usu.mutation.SetSuspendedAt(t)
	return usu
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableSuspendedAt(t *time.Time) *UserSettingUpdate {
	if t != nil {
		usu.SetSuspendedAt(*t)
	}
	return usu
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (usu *UserSettingUpdate) ClearSuspendedAt() *UserSettingUpdate {
	usu.mutation.ClearSuspendedAt()
	return usu
}

// SetRecoveryCode sets the "recovery_code" field.
func (usu *UserSettingUpdate) SetRecoveryCode(s string) *UserSettingUpdate {
	usu.mutation.SetRecoveryCode(s)
	return usu
}

// SetNillableRecoveryCode sets the "recovery_code" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableRecoveryCode(s *string) *UserSettingUpdate {
	if s != nil {
		usu.SetRecoveryCode(*s)
	}
	return usu
}

// ClearRecoveryCode clears the value of the "recovery_code" field.
func (usu *UserSettingUpdate) ClearRecoveryCode() *UserSettingUpdate {
	usu.mutation.ClearRecoveryCode()
	return usu
}

// SetStatus sets the "status" field.
func (usu *UserSettingUpdate) SetStatus(es enums.UserStatus) *UserSettingUpdate {
	usu.mutation.SetStatus(es)
	return usu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableStatus(es *enums.UserStatus) *UserSettingUpdate {
	if es != nil {
		usu.SetStatus(*es)
	}
	return usu
}

// SetDefaultOrg sets the "default_org" field.
func (usu *UserSettingUpdate) SetDefaultOrg(s string) *UserSettingUpdate {
	usu.mutation.SetDefaultOrg(s)
	return usu
}

// SetNillableDefaultOrg sets the "default_org" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableDefaultOrg(s *string) *UserSettingUpdate {
	if s != nil {
		usu.SetDefaultOrg(*s)
	}
	return usu
}

// ClearDefaultOrg clears the value of the "default_org" field.
func (usu *UserSettingUpdate) ClearDefaultOrg() *UserSettingUpdate {
	usu.mutation.ClearDefaultOrg()
	return usu
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (usu *UserSettingUpdate) SetEmailConfirmed(b bool) *UserSettingUpdate {
	usu.mutation.SetEmailConfirmed(b)
	return usu
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableEmailConfirmed(b *bool) *UserSettingUpdate {
	if b != nil {
		usu.SetEmailConfirmed(*b)
	}
	return usu
}

// SetTags sets the "tags" field.
func (usu *UserSettingUpdate) SetTags(s []string) *UserSettingUpdate {
	usu.mutation.SetTags(s)
	return usu
}

// AppendTags appends s to the "tags" field.
func (usu *UserSettingUpdate) AppendTags(s []string) *UserSettingUpdate {
	usu.mutation.AppendTags(s)
	return usu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usu *UserSettingUpdate) SetUserID(id string) *UserSettingUpdate {
	usu.mutation.SetUserID(id)
	return usu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usu *UserSettingUpdate) SetNillableUserID(id *string) *UserSettingUpdate {
	if id != nil {
		usu = usu.SetUserID(*id)
	}
	return usu
}

// SetUser sets the "user" edge to the User entity.
func (usu *UserSettingUpdate) SetUser(u *User) *UserSettingUpdate {
	return usu.SetUserID(u.ID)
}

// Mutation returns the UserSettingMutation object of the builder.
func (usu *UserSettingUpdate) Mutation() *UserSettingMutation {
	return usu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usu *UserSettingUpdate) ClearUser() *UserSettingUpdate {
	usu.mutation.ClearUser()
	return usu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (usu *UserSettingUpdate) Save(ctx context.Context) (int, error) {
	if err := usu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, usu.sqlSave, usu.mutation, usu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usu *UserSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := usu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (usu *UserSettingUpdate) Exec(ctx context.Context) error {
	_, err := usu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usu *UserSettingUpdate) ExecX(ctx context.Context) {
	if err := usu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usu *UserSettingUpdate) defaults() error {
	if _, ok := usu.mutation.UpdatedAt(); !ok && !usu.mutation.UpdatedAtCleared() {
		if usersetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized usersetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := usersetting.UpdateDefaultUpdatedAt()
		usu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usu *UserSettingUpdate) check() error {
	if v, ok := usu.mutation.Status(); ok {
		if err := usersetting.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSetting.status": %w`, err)}
		}
	}
	return nil
}

func (usu *UserSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := usu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersetting.Table, usersetting.Columns, sqlgraph.NewFieldSpec(usersetting.FieldID, field.TypeString))
	if ps := usu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if usu.mutation.CreatedAtCleared() {
		_spec.ClearField(usersetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := usu.mutation.UpdatedAt(); ok {
		_spec.SetField(usersetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if usu.mutation.UpdatedAtCleared() {
		_spec.ClearField(usersetting.FieldUpdatedAt, field.TypeTime)
	}
	if usu.mutation.CreatedByCleared() {
		_spec.ClearField(usersetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := usu.mutation.UpdatedBy(); ok {
		_spec.SetField(usersetting.FieldUpdatedBy, field.TypeString, value)
	}
	if usu.mutation.UpdatedByCleared() {
		_spec.ClearField(usersetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := usu.mutation.DeletedAt(); ok {
		_spec.SetField(usersetting.FieldDeletedAt, field.TypeTime, value)
	}
	if usu.mutation.DeletedAtCleared() {
		_spec.ClearField(usersetting.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := usu.mutation.DeletedBy(); ok {
		_spec.SetField(usersetting.FieldDeletedBy, field.TypeString, value)
	}
	if usu.mutation.DeletedByCleared() {
		_spec.ClearField(usersetting.FieldDeletedBy, field.TypeString)
	}
	if value, ok := usu.mutation.Locked(); ok {
		_spec.SetField(usersetting.FieldLocked, field.TypeBool, value)
	}
	if value, ok := usu.mutation.SilencedAt(); ok {
		_spec.SetField(usersetting.FieldSilencedAt, field.TypeTime, value)
	}
	if usu.mutation.SilencedAtCleared() {
		_spec.ClearField(usersetting.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := usu.mutation.SuspendedAt(); ok {
		_spec.SetField(usersetting.FieldSuspendedAt, field.TypeTime, value)
	}
	if usu.mutation.SuspendedAtCleared() {
		_spec.ClearField(usersetting.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := usu.mutation.RecoveryCode(); ok {
		_spec.SetField(usersetting.FieldRecoveryCode, field.TypeString, value)
	}
	if usu.mutation.RecoveryCodeCleared() {
		_spec.ClearField(usersetting.FieldRecoveryCode, field.TypeString)
	}
	if value, ok := usu.mutation.Status(); ok {
		_spec.SetField(usersetting.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := usu.mutation.DefaultOrg(); ok {
		_spec.SetField(usersetting.FieldDefaultOrg, field.TypeString, value)
	}
	if usu.mutation.DefaultOrgCleared() {
		_spec.ClearField(usersetting.FieldDefaultOrg, field.TypeString)
	}
	if value, ok := usu.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersetting.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := usu.mutation.Tags(); ok {
		_spec.SetField(usersetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := usu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, usersetting.FieldTags, value)
		})
	}
	if usu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersetting.UserTable,
			Columns: []string{usersetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = usu.schemaConfig.UserSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersetting.UserTable,
			Columns: []string{usersetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = usu.schemaConfig.UserSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = usu.schemaConfig.UserSetting
	ctx = internal.NewSchemaConfigContext(ctx, usu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, usu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	usu.mutation.done = true
	return n, nil
}

// UserSettingUpdateOne is the builder for updating a single UserSetting entity.
type UserSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserSettingMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (usuo *UserSettingUpdateOne) SetUpdatedAt(t time.Time) *UserSettingUpdateOne {
	usuo.mutation.SetUpdatedAt(t)
	return usuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (usuo *UserSettingUpdateOne) ClearUpdatedAt() *UserSettingUpdateOne {
	usuo.mutation.ClearUpdatedAt()
	return usuo
}

// SetUpdatedBy sets the "updated_by" field.
func (usuo *UserSettingUpdateOne) SetUpdatedBy(s string) *UserSettingUpdateOne {
	usuo.mutation.SetUpdatedBy(s)
	return usuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableUpdatedBy(s *string) *UserSettingUpdateOne {
	if s != nil {
		usuo.SetUpdatedBy(*s)
	}
	return usuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (usuo *UserSettingUpdateOne) ClearUpdatedBy() *UserSettingUpdateOne {
	usuo.mutation.ClearUpdatedBy()
	return usuo
}

// SetDeletedAt sets the "deleted_at" field.
func (usuo *UserSettingUpdateOne) SetDeletedAt(t time.Time) *UserSettingUpdateOne {
	usuo.mutation.SetDeletedAt(t)
	return usuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableDeletedAt(t *time.Time) *UserSettingUpdateOne {
	if t != nil {
		usuo.SetDeletedAt(*t)
	}
	return usuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (usuo *UserSettingUpdateOne) ClearDeletedAt() *UserSettingUpdateOne {
	usuo.mutation.ClearDeletedAt()
	return usuo
}

// SetDeletedBy sets the "deleted_by" field.
func (usuo *UserSettingUpdateOne) SetDeletedBy(s string) *UserSettingUpdateOne {
	usuo.mutation.SetDeletedBy(s)
	return usuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableDeletedBy(s *string) *UserSettingUpdateOne {
	if s != nil {
		usuo.SetDeletedBy(*s)
	}
	return usuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (usuo *UserSettingUpdateOne) ClearDeletedBy() *UserSettingUpdateOne {
	usuo.mutation.ClearDeletedBy()
	return usuo
}

// SetLocked sets the "locked" field.
func (usuo *UserSettingUpdateOne) SetLocked(b bool) *UserSettingUpdateOne {
	usuo.mutation.SetLocked(b)
	return usuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableLocked(b *bool) *UserSettingUpdateOne {
	if b != nil {
		usuo.SetLocked(*b)
	}
	return usuo
}

// SetSilencedAt sets the "silenced_at" field.
func (usuo *UserSettingUpdateOne) SetSilencedAt(t time.Time) *UserSettingUpdateOne {
	usuo.mutation.SetSilencedAt(t)
	return usuo
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableSilencedAt(t *time.Time) *UserSettingUpdateOne {
	if t != nil {
		usuo.SetSilencedAt(*t)
	}
	return usuo
}

// ClearSilencedAt clears the value of the "silenced_at" field.
func (usuo *UserSettingUpdateOne) ClearSilencedAt() *UserSettingUpdateOne {
	usuo.mutation.ClearSilencedAt()
	return usuo
}

// SetSuspendedAt sets the "suspended_at" field.
func (usuo *UserSettingUpdateOne) SetSuspendedAt(t time.Time) *UserSettingUpdateOne {
	usuo.mutation.SetSuspendedAt(t)
	return usuo
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableSuspendedAt(t *time.Time) *UserSettingUpdateOne {
	if t != nil {
		usuo.SetSuspendedAt(*t)
	}
	return usuo
}

// ClearSuspendedAt clears the value of the "suspended_at" field.
func (usuo *UserSettingUpdateOne) ClearSuspendedAt() *UserSettingUpdateOne {
	usuo.mutation.ClearSuspendedAt()
	return usuo
}

// SetRecoveryCode sets the "recovery_code" field.
func (usuo *UserSettingUpdateOne) SetRecoveryCode(s string) *UserSettingUpdateOne {
	usuo.mutation.SetRecoveryCode(s)
	return usuo
}

// SetNillableRecoveryCode sets the "recovery_code" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableRecoveryCode(s *string) *UserSettingUpdateOne {
	if s != nil {
		usuo.SetRecoveryCode(*s)
	}
	return usuo
}

// ClearRecoveryCode clears the value of the "recovery_code" field.
func (usuo *UserSettingUpdateOne) ClearRecoveryCode() *UserSettingUpdateOne {
	usuo.mutation.ClearRecoveryCode()
	return usuo
}

// SetStatus sets the "status" field.
func (usuo *UserSettingUpdateOne) SetStatus(es enums.UserStatus) *UserSettingUpdateOne {
	usuo.mutation.SetStatus(es)
	return usuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableStatus(es *enums.UserStatus) *UserSettingUpdateOne {
	if es != nil {
		usuo.SetStatus(*es)
	}
	return usuo
}

// SetDefaultOrg sets the "default_org" field.
func (usuo *UserSettingUpdateOne) SetDefaultOrg(s string) *UserSettingUpdateOne {
	usuo.mutation.SetDefaultOrg(s)
	return usuo
}

// SetNillableDefaultOrg sets the "default_org" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableDefaultOrg(s *string) *UserSettingUpdateOne {
	if s != nil {
		usuo.SetDefaultOrg(*s)
	}
	return usuo
}

// ClearDefaultOrg clears the value of the "default_org" field.
func (usuo *UserSettingUpdateOne) ClearDefaultOrg() *UserSettingUpdateOne {
	usuo.mutation.ClearDefaultOrg()
	return usuo
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (usuo *UserSettingUpdateOne) SetEmailConfirmed(b bool) *UserSettingUpdateOne {
	usuo.mutation.SetEmailConfirmed(b)
	return usuo
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableEmailConfirmed(b *bool) *UserSettingUpdateOne {
	if b != nil {
		usuo.SetEmailConfirmed(*b)
	}
	return usuo
}

// SetTags sets the "tags" field.
func (usuo *UserSettingUpdateOne) SetTags(s []string) *UserSettingUpdateOne {
	usuo.mutation.SetTags(s)
	return usuo
}

// AppendTags appends s to the "tags" field.
func (usuo *UserSettingUpdateOne) AppendTags(s []string) *UserSettingUpdateOne {
	usuo.mutation.AppendTags(s)
	return usuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (usuo *UserSettingUpdateOne) SetUserID(id string) *UserSettingUpdateOne {
	usuo.mutation.SetUserID(id)
	return usuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (usuo *UserSettingUpdateOne) SetNillableUserID(id *string) *UserSettingUpdateOne {
	if id != nil {
		usuo = usuo.SetUserID(*id)
	}
	return usuo
}

// SetUser sets the "user" edge to the User entity.
func (usuo *UserSettingUpdateOne) SetUser(u *User) *UserSettingUpdateOne {
	return usuo.SetUserID(u.ID)
}

// Mutation returns the UserSettingMutation object of the builder.
func (usuo *UserSettingUpdateOne) Mutation() *UserSettingMutation {
	return usuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (usuo *UserSettingUpdateOne) ClearUser() *UserSettingUpdateOne {
	usuo.mutation.ClearUser()
	return usuo
}

// Where appends a list predicates to the UserSettingUpdate builder.
func (usuo *UserSettingUpdateOne) Where(ps ...predicate.UserSetting) *UserSettingUpdateOne {
	usuo.mutation.Where(ps...)
	return usuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (usuo *UserSettingUpdateOne) Select(field string, fields ...string) *UserSettingUpdateOne {
	usuo.fields = append([]string{field}, fields...)
	return usuo
}

// Save executes the query and returns the updated UserSetting entity.
func (usuo *UserSettingUpdateOne) Save(ctx context.Context) (*UserSetting, error) {
	if err := usuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, usuo.sqlSave, usuo.mutation, usuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (usuo *UserSettingUpdateOne) SaveX(ctx context.Context) *UserSetting {
	node, err := usuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (usuo *UserSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := usuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usuo *UserSettingUpdateOne) ExecX(ctx context.Context) {
	if err := usuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usuo *UserSettingUpdateOne) defaults() error {
	if _, ok := usuo.mutation.UpdatedAt(); !ok && !usuo.mutation.UpdatedAtCleared() {
		if usersetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized usersetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := usersetting.UpdateDefaultUpdatedAt()
		usuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usuo *UserSettingUpdateOne) check() error {
	if v, ok := usuo.mutation.Status(); ok {
		if err := usersetting.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSetting.status": %w`, err)}
		}
	}
	return nil
}

func (usuo *UserSettingUpdateOne) sqlSave(ctx context.Context) (_node *UserSetting, err error) {
	if err := usuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usersetting.Table, usersetting.Columns, sqlgraph.NewFieldSpec(usersetting.FieldID, field.TypeString))
	id, ok := usuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "UserSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := usuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersetting.FieldID)
		for _, f := range fields {
			if !usersetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != usersetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := usuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if usuo.mutation.CreatedAtCleared() {
		_spec.ClearField(usersetting.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := usuo.mutation.UpdatedAt(); ok {
		_spec.SetField(usersetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if usuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(usersetting.FieldUpdatedAt, field.TypeTime)
	}
	if usuo.mutation.CreatedByCleared() {
		_spec.ClearField(usersetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := usuo.mutation.UpdatedBy(); ok {
		_spec.SetField(usersetting.FieldUpdatedBy, field.TypeString, value)
	}
	if usuo.mutation.UpdatedByCleared() {
		_spec.ClearField(usersetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := usuo.mutation.DeletedAt(); ok {
		_spec.SetField(usersetting.FieldDeletedAt, field.TypeTime, value)
	}
	if usuo.mutation.DeletedAtCleared() {
		_spec.ClearField(usersetting.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := usuo.mutation.DeletedBy(); ok {
		_spec.SetField(usersetting.FieldDeletedBy, field.TypeString, value)
	}
	if usuo.mutation.DeletedByCleared() {
		_spec.ClearField(usersetting.FieldDeletedBy, field.TypeString)
	}
	if value, ok := usuo.mutation.Locked(); ok {
		_spec.SetField(usersetting.FieldLocked, field.TypeBool, value)
	}
	if value, ok := usuo.mutation.SilencedAt(); ok {
		_spec.SetField(usersetting.FieldSilencedAt, field.TypeTime, value)
	}
	if usuo.mutation.SilencedAtCleared() {
		_spec.ClearField(usersetting.FieldSilencedAt, field.TypeTime)
	}
	if value, ok := usuo.mutation.SuspendedAt(); ok {
		_spec.SetField(usersetting.FieldSuspendedAt, field.TypeTime, value)
	}
	if usuo.mutation.SuspendedAtCleared() {
		_spec.ClearField(usersetting.FieldSuspendedAt, field.TypeTime)
	}
	if value, ok := usuo.mutation.RecoveryCode(); ok {
		_spec.SetField(usersetting.FieldRecoveryCode, field.TypeString, value)
	}
	if usuo.mutation.RecoveryCodeCleared() {
		_spec.ClearField(usersetting.FieldRecoveryCode, field.TypeString)
	}
	if value, ok := usuo.mutation.Status(); ok {
		_spec.SetField(usersetting.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := usuo.mutation.DefaultOrg(); ok {
		_spec.SetField(usersetting.FieldDefaultOrg, field.TypeString, value)
	}
	if usuo.mutation.DefaultOrgCleared() {
		_spec.ClearField(usersetting.FieldDefaultOrg, field.TypeString)
	}
	if value, ok := usuo.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersetting.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := usuo.mutation.Tags(); ok {
		_spec.SetField(usersetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := usuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, usersetting.FieldTags, value)
		})
	}
	if usuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersetting.UserTable,
			Columns: []string{usersetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = usuo.schemaConfig.UserSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := usuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usersetting.UserTable,
			Columns: []string{usersetting.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = usuo.schemaConfig.UserSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = usuo.schemaConfig.UserSetting
	ctx = internal.NewSchemaConfigContext(ctx, usuo.schemaConfig)
	_node = &UserSetting{config: usuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, usuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usersetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	usuo.mutation.done = true
	return _node, nil
}
