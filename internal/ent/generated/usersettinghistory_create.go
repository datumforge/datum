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
	"github.com/datumforge/datum/internal/ent/generated/usersettinghistory"
	"github.com/datumforge/enthistory"
)

// UserSettingHistoryCreate is the builder for creating a UserSettingHistory entity.
type UserSettingHistoryCreate struct {
	config
	mutation *UserSettingHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ushc *UserSettingHistoryCreate) SetHistoryTime(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetHistoryTime(t)
	return ushc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableHistoryTime(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetHistoryTime(*t)
	}
	return ushc
}

// SetOperation sets the "operation" field.
func (ushc *UserSettingHistoryCreate) SetOperation(et enthistory.OpType) *UserSettingHistoryCreate {
	ushc.mutation.SetOperation(et)
	return ushc
}

// SetRef sets the "ref" field.
func (ushc *UserSettingHistoryCreate) SetRef(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetRef(s)
	return ushc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableRef(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetRef(*s)
	}
	return ushc
}

// SetCreatedAt sets the "created_at" field.
func (ushc *UserSettingHistoryCreate) SetCreatedAt(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetCreatedAt(t)
	return ushc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableCreatedAt(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetCreatedAt(*t)
	}
	return ushc
}

// SetUpdatedAt sets the "updated_at" field.
func (ushc *UserSettingHistoryCreate) SetUpdatedAt(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetUpdatedAt(t)
	return ushc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableUpdatedAt(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetUpdatedAt(*t)
	}
	return ushc
}

// SetCreatedBy sets the "created_by" field.
func (ushc *UserSettingHistoryCreate) SetCreatedBy(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetCreatedBy(s)
	return ushc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableCreatedBy(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetCreatedBy(*s)
	}
	return ushc
}

// SetUpdatedBy sets the "updated_by" field.
func (ushc *UserSettingHistoryCreate) SetUpdatedBy(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetUpdatedBy(s)
	return ushc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableUpdatedBy(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetUpdatedBy(*s)
	}
	return ushc
}

// SetDeletedAt sets the "deleted_at" field.
func (ushc *UserSettingHistoryCreate) SetDeletedAt(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetDeletedAt(t)
	return ushc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableDeletedAt(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetDeletedAt(*t)
	}
	return ushc
}

// SetDeletedBy sets the "deleted_by" field.
func (ushc *UserSettingHistoryCreate) SetDeletedBy(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetDeletedBy(s)
	return ushc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableDeletedBy(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetDeletedBy(*s)
	}
	return ushc
}

// SetLocked sets the "locked" field.
func (ushc *UserSettingHistoryCreate) SetLocked(b bool) *UserSettingHistoryCreate {
	ushc.mutation.SetLocked(b)
	return ushc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableLocked(b *bool) *UserSettingHistoryCreate {
	if b != nil {
		ushc.SetLocked(*b)
	}
	return ushc
}

// SetSilencedAt sets the "silenced_at" field.
func (ushc *UserSettingHistoryCreate) SetSilencedAt(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetSilencedAt(t)
	return ushc
}

// SetNillableSilencedAt sets the "silenced_at" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableSilencedAt(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetSilencedAt(*t)
	}
	return ushc
}

// SetSuspendedAt sets the "suspended_at" field.
func (ushc *UserSettingHistoryCreate) SetSuspendedAt(t time.Time) *UserSettingHistoryCreate {
	ushc.mutation.SetSuspendedAt(t)
	return ushc
}

// SetNillableSuspendedAt sets the "suspended_at" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableSuspendedAt(t *time.Time) *UserSettingHistoryCreate {
	if t != nil {
		ushc.SetSuspendedAt(*t)
	}
	return ushc
}

// SetRecoveryCode sets the "recovery_code" field.
func (ushc *UserSettingHistoryCreate) SetRecoveryCode(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetRecoveryCode(s)
	return ushc
}

// SetNillableRecoveryCode sets the "recovery_code" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableRecoveryCode(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetRecoveryCode(*s)
	}
	return ushc
}

// SetStatus sets the "status" field.
func (ushc *UserSettingHistoryCreate) SetStatus(es enums.UserStatus) *UserSettingHistoryCreate {
	ushc.mutation.SetStatus(es)
	return ushc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableStatus(es *enums.UserStatus) *UserSettingHistoryCreate {
	if es != nil {
		ushc.SetStatus(*es)
	}
	return ushc
}

// SetDefaultOrg sets the "default_org" field.
func (ushc *UserSettingHistoryCreate) SetDefaultOrg(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetDefaultOrg(s)
	return ushc
}

// SetNillableDefaultOrg sets the "default_org" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableDefaultOrg(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetDefaultOrg(*s)
	}
	return ushc
}

// SetEmailConfirmed sets the "email_confirmed" field.
func (ushc *UserSettingHistoryCreate) SetEmailConfirmed(b bool) *UserSettingHistoryCreate {
	ushc.mutation.SetEmailConfirmed(b)
	return ushc
}

// SetNillableEmailConfirmed sets the "email_confirmed" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableEmailConfirmed(b *bool) *UserSettingHistoryCreate {
	if b != nil {
		ushc.SetEmailConfirmed(*b)
	}
	return ushc
}

// SetTags sets the "tags" field.
func (ushc *UserSettingHistoryCreate) SetTags(s []string) *UserSettingHistoryCreate {
	ushc.mutation.SetTags(s)
	return ushc
}

// SetID sets the "id" field.
func (ushc *UserSettingHistoryCreate) SetID(s string) *UserSettingHistoryCreate {
	ushc.mutation.SetID(s)
	return ushc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ushc *UserSettingHistoryCreate) SetNillableID(s *string) *UserSettingHistoryCreate {
	if s != nil {
		ushc.SetID(*s)
	}
	return ushc
}

// Mutation returns the UserSettingHistoryMutation object of the builder.
func (ushc *UserSettingHistoryCreate) Mutation() *UserSettingHistoryMutation {
	return ushc.mutation
}

// Save creates the UserSettingHistory in the database.
func (ushc *UserSettingHistoryCreate) Save(ctx context.Context) (*UserSettingHistory, error) {
	ushc.defaults()
	return withHooks(ctx, ushc.sqlSave, ushc.mutation, ushc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ushc *UserSettingHistoryCreate) SaveX(ctx context.Context) *UserSettingHistory {
	v, err := ushc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ushc *UserSettingHistoryCreate) Exec(ctx context.Context) error {
	_, err := ushc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ushc *UserSettingHistoryCreate) ExecX(ctx context.Context) {
	if err := ushc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ushc *UserSettingHistoryCreate) defaults() {
	if _, ok := ushc.mutation.HistoryTime(); !ok {
		v := usersettinghistory.DefaultHistoryTime()
		ushc.mutation.SetHistoryTime(v)
	}
	if _, ok := ushc.mutation.CreatedAt(); !ok {
		v := usersettinghistory.DefaultCreatedAt()
		ushc.mutation.SetCreatedAt(v)
	}
	if _, ok := ushc.mutation.UpdatedAt(); !ok {
		v := usersettinghistory.DefaultUpdatedAt()
		ushc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ushc.mutation.Locked(); !ok {
		v := usersettinghistory.DefaultLocked
		ushc.mutation.SetLocked(v)
	}
	if _, ok := ushc.mutation.Status(); !ok {
		v := usersettinghistory.DefaultStatus
		ushc.mutation.SetStatus(v)
	}
	if _, ok := ushc.mutation.EmailConfirmed(); !ok {
		v := usersettinghistory.DefaultEmailConfirmed
		ushc.mutation.SetEmailConfirmed(v)
	}
	if _, ok := ushc.mutation.Tags(); !ok {
		v := usersettinghistory.DefaultTags
		ushc.mutation.SetTags(v)
	}
	if _, ok := ushc.mutation.ID(); !ok {
		v := usersettinghistory.DefaultID()
		ushc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ushc *UserSettingHistoryCreate) check() error {
	if _, ok := ushc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "UserSettingHistory.history_time"`)}
	}
	if _, ok := ushc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "UserSettingHistory.operation"`)}
	}
	if v, ok := ushc.mutation.Operation(); ok {
		if err := usersettinghistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "UserSettingHistory.operation": %w`, err)}
		}
	}
	if _, ok := ushc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "UserSettingHistory.created_at"`)}
	}
	if _, ok := ushc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "UserSettingHistory.updated_at"`)}
	}
	if _, ok := ushc.mutation.Locked(); !ok {
		return &ValidationError{Name: "locked", err: errors.New(`generated: missing required field "UserSettingHistory.locked"`)}
	}
	if _, ok := ushc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "UserSettingHistory.status"`)}
	}
	if v, ok := ushc.mutation.Status(); ok {
		if err := usersettinghistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "UserSettingHistory.status": %w`, err)}
		}
	}
	if _, ok := ushc.mutation.EmailConfirmed(); !ok {
		return &ValidationError{Name: "email_confirmed", err: errors.New(`generated: missing required field "UserSettingHistory.email_confirmed"`)}
	}
	if _, ok := ushc.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`generated: missing required field "UserSettingHistory.tags"`)}
	}
	return nil
}

func (ushc *UserSettingHistoryCreate) sqlSave(ctx context.Context) (*UserSettingHistory, error) {
	if err := ushc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ushc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ushc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserSettingHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ushc.mutation.id = &_node.ID
	ushc.mutation.done = true
	return _node, nil
}

func (ushc *UserSettingHistoryCreate) createSpec() (*UserSettingHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSettingHistory{config: ushc.config}
		_spec = sqlgraph.NewCreateSpec(usersettinghistory.Table, sqlgraph.NewFieldSpec(usersettinghistory.FieldID, field.TypeString))
	)
	_spec.Schema = ushc.schemaConfig.UserSettingHistory
	if id, ok := ushc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ushc.mutation.HistoryTime(); ok {
		_spec.SetField(usersettinghistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ushc.mutation.Operation(); ok {
		_spec.SetField(usersettinghistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ushc.mutation.Ref(); ok {
		_spec.SetField(usersettinghistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ushc.mutation.CreatedAt(); ok {
		_spec.SetField(usersettinghistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ushc.mutation.UpdatedAt(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ushc.mutation.CreatedBy(); ok {
		_spec.SetField(usersettinghistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ushc.mutation.UpdatedBy(); ok {
		_spec.SetField(usersettinghistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ushc.mutation.DeletedAt(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ushc.mutation.DeletedBy(); ok {
		_spec.SetField(usersettinghistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ushc.mutation.Locked(); ok {
		_spec.SetField(usersettinghistory.FieldLocked, field.TypeBool, value)
		_node.Locked = value
	}
	if value, ok := ushc.mutation.SilencedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSilencedAt, field.TypeTime, value)
		_node.SilencedAt = &value
	}
	if value, ok := ushc.mutation.SuspendedAt(); ok {
		_spec.SetField(usersettinghistory.FieldSuspendedAt, field.TypeTime, value)
		_node.SuspendedAt = &value
	}
	if value, ok := ushc.mutation.RecoveryCode(); ok {
		_spec.SetField(usersettinghistory.FieldRecoveryCode, field.TypeString, value)
		_node.RecoveryCode = &value
	}
	if value, ok := ushc.mutation.Status(); ok {
		_spec.SetField(usersettinghistory.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := ushc.mutation.DefaultOrg(); ok {
		_spec.SetField(usersettinghistory.FieldDefaultOrg, field.TypeString, value)
		_node.DefaultOrg = value
	}
	if value, ok := ushc.mutation.EmailConfirmed(); ok {
		_spec.SetField(usersettinghistory.FieldEmailConfirmed, field.TypeBool, value)
		_node.EmailConfirmed = value
	}
	if value, ok := ushc.mutation.Tags(); ok {
		_spec.SetField(usersettinghistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	return _node, _spec
}

// UserSettingHistoryCreateBulk is the builder for creating many UserSettingHistory entities in bulk.
type UserSettingHistoryCreateBulk struct {
	config
	err      error
	builders []*UserSettingHistoryCreate
}

// Save creates the UserSettingHistory entities in the database.
func (ushcb *UserSettingHistoryCreateBulk) Save(ctx context.Context) ([]*UserSettingHistory, error) {
	if ushcb.err != nil {
		return nil, ushcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ushcb.builders))
	nodes := make([]*UserSettingHistory, len(ushcb.builders))
	mutators := make([]Mutator, len(ushcb.builders))
	for i := range ushcb.builders {
		func(i int, root context.Context) {
			builder := ushcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSettingHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ushcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ushcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ushcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ushcb *UserSettingHistoryCreateBulk) SaveX(ctx context.Context) []*UserSettingHistory {
	v, err := ushcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ushcb *UserSettingHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ushcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ushcb *UserSettingHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ushcb.Exec(ctx); err != nil {
		panic(err)
	}
}
