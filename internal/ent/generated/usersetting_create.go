// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
)

// UserSettingCreate is the builder for creating a UserSetting entity.
type UserSettingCreate struct {
	config
	mutation *UserSettingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSettingCreate) SetCreatedAt(t time.Time) *UserSettingCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableCreatedAt(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserSettingCreate) SetUpdatedAt(t time.Time) *UserSettingCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableUpdatedAt(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetCreatedBy sets the "created_by" field.
func (usc *UserSettingCreate) SetCreatedBy(s string) *UserSettingCreate {
	usc.mutation.SetCreatedBy(s)
	return usc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableCreatedBy(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetCreatedBy(*s)
	}
	return usc
}

// SetUpdatedBy sets the "updated_by" field.
func (usc *UserSettingCreate) SetUpdatedBy(s string) *UserSettingCreate {
	usc.mutation.SetUpdatedBy(s)
	return usc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableUpdatedBy(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetUpdatedBy(*s)
	}
	return usc
}

// SetDeletedAt sets the "deleted_at" field.
func (usc *UserSettingCreate) SetDeletedAt(t time.Time) *UserSettingCreate {
	usc.mutation.SetDeletedAt(t)
	return usc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableDeletedAt(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetDeletedAt(*t)
	}
	return usc
}

// SetDeletedBy sets the "deleted_by" field.
func (usc *UserSettingCreate) SetDeletedBy(s string) *UserSettingCreate {
	usc.mutation.SetDeletedBy(s)
	return usc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableDeletedBy(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetDeletedBy(*s)
	}
	return usc
}

// SetUserID sets the "user_id" field.
func (usc *UserSettingCreate) SetUserID(s string) *UserSettingCreate {
	usc.mutation.SetUserID(s)
	return usc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableUserID(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetUserID(*s)
	}
	return usc
}

// SetLocked sets the "locked" field.
func (usc *UserSettingCreate) SetLocked(b bool) *UserSettingCreate {
	usc.mutation.SetLocked(b)
	return usc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableLocked(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetLocked(*b)
	}
	return usc
}

// SetSilencedAt sets the "silenced_at" field.
func (usc *UserSettingCreate) SetSilencedAt(t time.Time) *UserSettingCreate {
	usc.mutation.SetSilencedAt(t)
	return usc
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableSilencedAt(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetSilencedAt(*t)
	}
	return usc
}

// SetSuspendedAt sets the "suspended_at" field.
func (usc *UserSettingCreate) SetSuspendedAt(t time.Time) *UserSettingCreate {
	usc.mutation.SetSuspendedAt(t)
	return usc
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableSuspendedAt(t *time.Time) *UserSettingCreate {
	if t != nil {
		usc.SetSuspendedAt(*t)
	}
	return usc
}

// SetStatus sets the "status" field.
func (usc *UserSettingCreate) SetStatus(es enums.UserStatus) *UserSettingCreate {
	usc.mutation.SetStatus(es)
	return usc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableStatus(es *enums.UserStatus) *UserSettingCreate {
	if es != nil {
		usc.SetStatus(*es)
	}
	return usc
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (usc *UserSettingCreate) SetEmailConfirmed(b bool) *UserSettingCreate {
	usc.mutation.SetEmailConfirmed(b)
	return usc
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableEmailConfirmed(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetEmailConfirmed(*b)
	}
	return usc
}

// SetTags sets the "tags" field.
func (usc *UserSettingCreate) SetTags(s []string) *UserSettingCreate {
	usc.mutation.SetTags(s)
	return usc
}

// SetTfaSecret sets the "tfa_secret" field.
func (usc *UserSettingCreate) SetTfaSecret(s string) *UserSettingCreate {
	usc.mutation.SetTfaSecret(s)
	return usc
}

// SetNillableTfaSecret sets the "tfa_secret" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableTfaSecret(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetTfaSecret(*s)
	}
	return usc
}

// SetIsPhoneOtpAllowed sets the "is_phone_otp_allowed" field.
func (usc *UserSettingCreate) SetIsPhoneOtpAllowed(b bool) *UserSettingCreate {
	usc.mutation.SetIsPhoneOtpAllowed(b)
	return usc
}

// SetNillableIsPhoneOtpAllowed sets the "is_phone_otp_allowed" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableIsPhoneOtpAllowed(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetIsPhoneOtpAllowed(*b)
	}
	return usc
}

// SetIsEmailOtpAllowed sets the "is_email_otp_allowed" field.
func (usc *UserSettingCreate) SetIsEmailOtpAllowed(b bool) *UserSettingCreate {
	usc.mutation.SetIsEmailOtpAllowed(b)
	return usc
}

// SetNillableIsEmailOtpAllowed sets the "is_email_otp_allowed" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableIsEmailOtpAllowed(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetIsEmailOtpAllowed(*b)
	}
	return usc
}

// SetIsTotpAllowed sets the "is_totp_allowed" field.
func (usc *UserSettingCreate) SetIsTotpAllowed(b bool) *UserSettingCreate {
	usc.mutation.SetIsTotpAllowed(b)
	return usc
}

// SetNillableIsTotpAllowed sets the "is_totp_allowed" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableIsTotpAllowed(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetIsTotpAllowed(*b)
	}
	return usc
}

// SetIsWebauthnAllowed sets the "is_webauthn_allowed" field.
func (usc *UserSettingCreate) SetIsWebauthnAllowed(b bool) *UserSettingCreate {
	usc.mutation.SetIsWebauthnAllowed(b)
	return usc
}

// SetNillableIsWebauthnAllowed sets the "is_webauthn_allowed" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableIsWebauthnAllowed(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetIsWebauthnAllowed(*b)
	}
	return usc
}

// SetIsTfaEnabled sets the "is_tfa_enabled" field.
func (usc *UserSettingCreate) SetIsTfaEnabled(b bool) *UserSettingCreate {
	usc.mutation.SetIsTfaEnabled(b)
	return usc
}

// SetNillableIsTfaEnabled sets the "is_tfa_enabled" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableIsTfaEnabled(b *bool) *UserSettingCreate {
	if b != nil {
		usc.SetIsTfaEnabled(*b)
	}
	return usc
}

// SetID sets the "id" field.
func (usc *UserSettingCreate) SetID(s string) *UserSettingCreate {
	usc.mutation.SetID(s)
	return usc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (usc *UserSettingCreate) SetNillableID(s *string) *UserSettingCreate {
	if s != nil {
		usc.SetID(*s)
	}
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSettingCreate) SetUser(u *User) *UserSettingCreate {
	return usc.SetUserID(u.ID)
}

// SetDefaultOrgID sets the "default_org" edge to the Organization entity by ID.
func (usc *UserSettingCreate) SetDefaultOrgID(id string) *UserSettingCreate {
	usc.mutation.SetDefaultOrgID(id)
	return usc
}

// SetNillableDefaultOrgID sets the "default_org" edge to the Organization entity by ID if the given value is not nil.
func (usc *UserSettingCreate) SetNillableDefaultOrgID(id *string) *UserSettingCreate {
	if id != nil {
		usc = usc.SetDefaultOrgID(*id)
	}
	return usc
}

// SetDefaultOrg sets the "default_org" edge to the Organization entity.
func (usc *UserSettingCreate) SetDefaultOrg(o *Organization) *UserSettingCreate {
	return usc.SetDefaultOrgID(o.ID)
}

// Mutation returns the UserSettingMutation object of the builder.
func (usc *UserSettingCreate) Mutation() *UserSettingMutation {
	return usc.mutation
}

// Save creates the UserSetting in the database.
func (usc *UserSettingCreate) Save(ctx context.Context) (*UserSetting, error) {
	if err := usc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSettingCreate) SaveX(ctx context.Context) *UserSetting {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSettingCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSettingCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSettingCreate) defaults() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		if usersetting.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized usersetting.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := usersetting.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		if usersetting.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized usersetting.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := usersetting.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
	if _, ok := usc.mutation.Locked(); !ok {
		v := usersetting.DefaultLocked
		usc.mutation.SetLocked(v)
	}
	if _, ok := usc.mutation.Status(); !ok {
		v := usersetting.DefaultStatus
		usc.mutation.SetStatus(v)
	}
	if _, ok := usc.mutation.EmailConfirmed(); !ok {
		v := usersetting.DefaultEmailConfirmed
		usc.mutation.SetEmailConfirmed(v)
	}
	if _, ok := usc.mutation.Tags(); !ok {
		v := usersetting.DefaultTags
		usc.mutation.SetTags(v)
	}
	if _, ok := usc.mutation.IsPhoneOtpAllowed(); !ok {
		v := usersetting.DefaultIsPhoneOtpAllowed
		usc.mutation.SetIsPhoneOtpAllowed(v)
	}
	if _, ok := usc.mutation.IsEmailOtpAllowed(); !ok {
		v := usersetting.DefaultIsEmailOtpAllowed
		usc.mutation.SetIsEmailOtpAllowed(v)
	}
	if _, ok := usc.mutation.IsTotpAllowed(); !ok {
		v := usersetting.DefaultIsTotpAllowed
		usc.mutation.SetIsTotpAllowed(v)
	}
	if _, ok := usc.mutation.IsWebauthnAllowed(); !ok {
		v := usersetting.DefaultIsWebauthnAllowed
		usc.mutation.SetIsWebauthnAllowed(v)
	}
	if _, ok := usc.mutation.IsTfaEnabled(); !ok {
		v := usersetting.DefaultIsTfaEnabled
		usc.mutation.SetIsTfaEnabled(v)
	}
	if _, ok := usc.mutation.ID(); !ok {
		if usersetting.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized usersetting.DefaultID (forgotten import generated/runtime?)")
		}
		v := usersetting.DefaultID()
		usc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSettingCreate) check() error {
	if _, ok := usc.mutation.Locked(); !ok {
		return &ValidationError{Name: "locked", err: errors.New(`generated: missing required field "UserSetting.locked"`)}
	}
	if _, ok := usc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "UserSetting.status"`)}
	}
	if v, ok := usc.mutation.Status(); ok {
		if err := usersetting.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSetting.status": %w`, err)}
		}
	}
	if _, ok := usc.mutation.EmailConfirmed(); !ok {
		return &ValidationError{Name: "email_confirmed", err: errors.New(`generated: missing required field "UserSetting.email_confirmed"`)}
	}
	if _, ok := usc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`generated: missing required field "UserSetting.tags"`)}
	}
	return nil
}

func (usc *UserSettingCreate) sqlSave(ctx context.Context) (*UserSetting, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserSetting.ID type: %T", _spec.ID.Value)
		}
	}
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSettingCreate) createSpec() (*UserSetting, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSetting{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersetting.Table, sqlgraph.NewFieldSpec(usersetting.FieldID, field.TypeString))
	)
	_spec.Schema = usc.schemaConfig.UserSetting
	if id, ok := usc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(usersetting.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.SetField(usersetting.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := usc.mutation.CreatedBy(); ok {
		_spec.SetField(usersetting.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := usc.mutation.UpdatedBy(); ok {
		_spec.SetField(usersetting.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := usc.mutation.DeletedAt(); ok {
		_spec.SetField(usersetting.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := usc.mutation.DeletedBy(); ok {
		_spec.SetField(usersetting.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := usc.mutation.Locked(); ok {
		_spec.SetField(usersetting.FieldLocked, field.TypeBool, value)
		_node.Locked = value
	}
	if value, ok := usc.mutation.SilencedAt(); ok {
		_spec.SetField(usersetting.FieldSilencedAt, field.TypeTime, value)
		_node.SilencedAt = &value
	}
	if value, ok := usc.mutation.SuspendedAt(); ok {
		_spec.SetField(usersetting.FieldSuspendedAt, field.TypeTime, value)
		_node.SuspendedAt = &value
	}
	if value, ok := usc.mutation.Status(); ok {
		_spec.SetField(usersetting.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := usc.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersetting.FieldEmailConfirmed, field.TypeBool, value)
		_node.EmailConfirmed = value
	}
	if value, ok := usc.mutation.Tags(); ok {
		_spec.SetField(usersetting.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := usc.mutation.TfaSecret(); ok {
		_spec.SetField(usersetting.FieldTfaSecret, field.TypeString, value)
		_node.TfaSecret = &value
	}
	if value, ok := usc.mutation.IsPhoneOtpAllowed(); ok {
		_spec.SetField(usersetting.FieldIsPhoneOtpAllowed, field.TypeBool, value)
		_node.IsPhoneOtpAllowed = value
	}
	if value, ok := usc.mutation.IsEmailOtpAllowed(); ok {
		_spec.SetField(usersetting.FieldIsEmailOtpAllowed, field.TypeBool, value)
		_node.IsEmailOtpAllowed = value
	}
	if value, ok := usc.mutation.IsTotpAllowed(); ok {
		_spec.SetField(usersetting.FieldIsTotpAllowed, field.TypeBool, value)
		_node.IsTotpAllowed = value
	}
	if value, ok := usc.mutation.IsWebauthnAllowed(); ok {
		_spec.SetField(usersetting.FieldIsWebauthnAllowed, field.TypeBool, value)
		_node.IsWebauthnAllowed = value
	}
	if value, ok := usc.mutation.IsTfaEnabled(); ok {
		_spec.SetField(usersetting.FieldIsTfaEnabled, field.TypeBool, value)
		_node.IsTfaEnabled = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
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
		edge.Schema = usc.schemaConfig.UserSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := usc.mutation.DefaultOrgIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usersetting.DefaultOrgTable,
			Columns: []string{usersetting.DefaultOrgColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = usc.schemaConfig.UserSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_setting_default_org = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSettingCreateBulk is the builder for creating many UserSetting entities in bulk.
type UserSettingCreateBulk struct {
	config
	err      error
	builders []*UserSettingCreate
}

// Save creates the UserSetting entities in the database.
func (uscb *UserSettingCreateBulk) Save(ctx context.Context) ([]*UserSetting, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSetting, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSettingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) SaveX(ctx context.Context) []*UserSetting {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSettingCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSettingCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
