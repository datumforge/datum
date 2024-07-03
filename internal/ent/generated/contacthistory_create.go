// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/contacthistory"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/enthistory"
)

// ContactHistoryCreate is the builder for creating a ContactHistory entity.
type ContactHistoryCreate struct {
	config
	mutation *ContactHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (chc *ContactHistoryCreate) SetHistoryTime(t time.Time) *ContactHistoryCreate {
	chc.mutation.SetHistoryTime(t)
	return chc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableHistoryTime(t *time.Time) *ContactHistoryCreate {
	if t != nil {
		chc.SetHistoryTime(*t)
	}
	return chc
}

// SetOperation sets the "operation" field.
func (chc *ContactHistoryCreate) SetOperation(et enthistory.OpType) *ContactHistoryCreate {
	chc.mutation.SetOperation(et)
	return chc
}

// SetRef sets the "ref" field.
func (chc *ContactHistoryCreate) SetRef(s string) *ContactHistoryCreate {
	chc.mutation.SetRef(s)
	return chc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableRef(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetRef(*s)
	}
	return chc
}

// SetCreatedAt sets the "created_at" field.
func (chc *ContactHistoryCreate) SetCreatedAt(t time.Time) *ContactHistoryCreate {
	chc.mutation.SetCreatedAt(t)
	return chc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableCreatedAt(t *time.Time) *ContactHistoryCreate {
	if t != nil {
		chc.SetCreatedAt(*t)
	}
	return chc
}

// SetUpdatedAt sets the "updated_at" field.
func (chc *ContactHistoryCreate) SetUpdatedAt(t time.Time) *ContactHistoryCreate {
	chc.mutation.SetUpdatedAt(t)
	return chc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableUpdatedAt(t *time.Time) *ContactHistoryCreate {
	if t != nil {
		chc.SetUpdatedAt(*t)
	}
	return chc
}

// SetCreatedBy sets the "created_by" field.
func (chc *ContactHistoryCreate) SetCreatedBy(s string) *ContactHistoryCreate {
	chc.mutation.SetCreatedBy(s)
	return chc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableCreatedBy(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetCreatedBy(*s)
	}
	return chc
}

// SetUpdatedBy sets the "updated_by" field.
func (chc *ContactHistoryCreate) SetUpdatedBy(s string) *ContactHistoryCreate {
	chc.mutation.SetUpdatedBy(s)
	return chc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableUpdatedBy(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetUpdatedBy(*s)
	}
	return chc
}

// SetMappingID sets the "mapping_id" field.
func (chc *ContactHistoryCreate) SetMappingID(s string) *ContactHistoryCreate {
	chc.mutation.SetMappingID(s)
	return chc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableMappingID(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetMappingID(*s)
	}
	return chc
}

// SetDeletedAt sets the "deleted_at" field.
func (chc *ContactHistoryCreate) SetDeletedAt(t time.Time) *ContactHistoryCreate {
	chc.mutation.SetDeletedAt(t)
	return chc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableDeletedAt(t *time.Time) *ContactHistoryCreate {
	if t != nil {
		chc.SetDeletedAt(*t)
	}
	return chc
}

// SetDeletedBy sets the "deleted_by" field.
func (chc *ContactHistoryCreate) SetDeletedBy(s string) *ContactHistoryCreate {
	chc.mutation.SetDeletedBy(s)
	return chc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableDeletedBy(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetDeletedBy(*s)
	}
	return chc
}

// SetTags sets the "tags" field.
func (chc *ContactHistoryCreate) SetTags(s []string) *ContactHistoryCreate {
	chc.mutation.SetTags(s)
	return chc
}

// SetOwnerID sets the "owner_id" field.
func (chc *ContactHistoryCreate) SetOwnerID(s string) *ContactHistoryCreate {
	chc.mutation.SetOwnerID(s)
	return chc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableOwnerID(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetOwnerID(*s)
	}
	return chc
}

// SetFullName sets the "full_name" field.
func (chc *ContactHistoryCreate) SetFullName(s string) *ContactHistoryCreate {
	chc.mutation.SetFullName(s)
	return chc
}

// SetTitle sets the "title" field.
func (chc *ContactHistoryCreate) SetTitle(s string) *ContactHistoryCreate {
	chc.mutation.SetTitle(s)
	return chc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableTitle(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetTitle(*s)
	}
	return chc
}

// SetCompany sets the "company" field.
func (chc *ContactHistoryCreate) SetCompany(s string) *ContactHistoryCreate {
	chc.mutation.SetCompany(s)
	return chc
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableCompany(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetCompany(*s)
	}
	return chc
}

// SetEmail sets the "email" field.
func (chc *ContactHistoryCreate) SetEmail(s string) *ContactHistoryCreate {
	chc.mutation.SetEmail(s)
	return chc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableEmail(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetEmail(*s)
	}
	return chc
}

// SetPhoneNumber sets the "phone_number" field.
func (chc *ContactHistoryCreate) SetPhoneNumber(s string) *ContactHistoryCreate {
	chc.mutation.SetPhoneNumber(s)
	return chc
}

// SetNillablePhoneNumber sets the "phone_number" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillablePhoneNumber(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetPhoneNumber(*s)
	}
	return chc
}

// SetAddress sets the "address" field.
func (chc *ContactHistoryCreate) SetAddress(s string) *ContactHistoryCreate {
	chc.mutation.SetAddress(s)
	return chc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableAddress(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetAddress(*s)
	}
	return chc
}

// SetStatus sets the "status" field.
func (chc *ContactHistoryCreate) SetStatus(es enums.UserStatus) *ContactHistoryCreate {
	chc.mutation.SetStatus(es)
	return chc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableStatus(es *enums.UserStatus) *ContactHistoryCreate {
	if es != nil {
		chc.SetStatus(*es)
	}
	return chc
}

// SetID sets the "id" field.
func (chc *ContactHistoryCreate) SetID(s string) *ContactHistoryCreate {
	chc.mutation.SetID(s)
	return chc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (chc *ContactHistoryCreate) SetNillableID(s *string) *ContactHistoryCreate {
	if s != nil {
		chc.SetID(*s)
	}
	return chc
}

// Mutation returns the ContactHistoryMutation object of the builder.
func (chc *ContactHistoryCreate) Mutation() *ContactHistoryMutation {
	return chc.mutation
}

// Save creates the ContactHistory in the database.
func (chc *ContactHistoryCreate) Save(ctx context.Context) (*ContactHistory, error) {
	chc.defaults()
	return withHooks(ctx, chc.sqlSave, chc.mutation, chc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (chc *ContactHistoryCreate) SaveX(ctx context.Context) *ContactHistory {
	v, err := chc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chc *ContactHistoryCreate) Exec(ctx context.Context) error {
	_, err := chc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chc *ContactHistoryCreate) ExecX(ctx context.Context) {
	if err := chc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (chc *ContactHistoryCreate) defaults() {
	if _, ok := chc.mutation.HistoryTime(); !ok {
		v := contacthistory.DefaultHistoryTime()
		chc.mutation.SetHistoryTime(v)
	}
	if _, ok := chc.mutation.CreatedAt(); !ok {
		v := contacthistory.DefaultCreatedAt()
		chc.mutation.SetCreatedAt(v)
	}
	if _, ok := chc.mutation.UpdatedAt(); !ok {
		v := contacthistory.DefaultUpdatedAt()
		chc.mutation.SetUpdatedAt(v)
	}
	if _, ok := chc.mutation.MappingID(); !ok {
		v := contacthistory.DefaultMappingID()
		chc.mutation.SetMappingID(v)
	}
	if _, ok := chc.mutation.Tags(); !ok {
		v := contacthistory.DefaultTags
		chc.mutation.SetTags(v)
	}
	if _, ok := chc.mutation.Status(); !ok {
		v := contacthistory.DefaultStatus
		chc.mutation.SetStatus(v)
	}
	if _, ok := chc.mutation.ID(); !ok {
		v := contacthistory.DefaultID()
		chc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (chc *ContactHistoryCreate) check() error {
	if _, ok := chc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "ContactHistory.history_time"`)}
	}
	if _, ok := chc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "ContactHistory.operation"`)}
	}
	if v, ok := chc.mutation.Operation(); ok {
		if err := contacthistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "ContactHistory.operation": %w`, err)}
		}
	}
	if _, ok := chc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "ContactHistory.mapping_id"`)}
	}
	if _, ok := chc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`generated: missing required field "ContactHistory.full_name"`)}
	}
	if _, ok := chc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`generated: missing required field "ContactHistory.status"`)}
	}
	if v, ok := chc.mutation.Status(); ok {
		if err := contacthistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "ContactHistory.status": %w`, err)}
		}
	}
	return nil
}

func (chc *ContactHistoryCreate) sqlSave(ctx context.Context) (*ContactHistory, error) {
	if err := chc.check(); err != nil {
		return nil, err
	}
	_node, _spec := chc.createSpec()
	if err := sqlgraph.CreateNode(ctx, chc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ContactHistory.ID type: %T", _spec.ID.Value)
		}
	}
	chc.mutation.id = &_node.ID
	chc.mutation.done = true
	return _node, nil
}

func (chc *ContactHistoryCreate) createSpec() (*ContactHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &ContactHistory{config: chc.config}
		_spec = sqlgraph.NewCreateSpec(contacthistory.Table, sqlgraph.NewFieldSpec(contacthistory.FieldID, field.TypeString))
	)
	_spec.Schema = chc.schemaConfig.ContactHistory
	if id, ok := chc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := chc.mutation.HistoryTime(); ok {
		_spec.SetField(contacthistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := chc.mutation.Operation(); ok {
		_spec.SetField(contacthistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := chc.mutation.Ref(); ok {
		_spec.SetField(contacthistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := chc.mutation.CreatedAt(); ok {
		_spec.SetField(contacthistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := chc.mutation.UpdatedAt(); ok {
		_spec.SetField(contacthistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := chc.mutation.CreatedBy(); ok {
		_spec.SetField(contacthistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := chc.mutation.UpdatedBy(); ok {
		_spec.SetField(contacthistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := chc.mutation.MappingID(); ok {
		_spec.SetField(contacthistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := chc.mutation.DeletedAt(); ok {
		_spec.SetField(contacthistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := chc.mutation.DeletedBy(); ok {
		_spec.SetField(contacthistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := chc.mutation.Tags(); ok {
		_spec.SetField(contacthistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := chc.mutation.OwnerID(); ok {
		_spec.SetField(contacthistory.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := chc.mutation.FullName(); ok {
		_spec.SetField(contacthistory.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := chc.mutation.Title(); ok {
		_spec.SetField(contacthistory.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := chc.mutation.Company(); ok {
		_spec.SetField(contacthistory.FieldCompany, field.TypeString, value)
		_node.Company = value
	}
	if value, ok := chc.mutation.Email(); ok {
		_spec.SetField(contacthistory.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := chc.mutation.PhoneNumber(); ok {
		_spec.SetField(contacthistory.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := chc.mutation.Address(); ok {
		_spec.SetField(contacthistory.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := chc.mutation.Status(); ok {
		_spec.SetField(contacthistory.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	return _node, _spec
}

// ContactHistoryCreateBulk is the builder for creating many ContactHistory entities in bulk.
type ContactHistoryCreateBulk struct {
	config
	err      error
	builders []*ContactHistoryCreate
}

// Save creates the ContactHistory entities in the database.
func (chcb *ContactHistoryCreateBulk) Save(ctx context.Context) ([]*ContactHistory, error) {
	if chcb.err != nil {
		return nil, chcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(chcb.builders))
	nodes := make([]*ContactHistory, len(chcb.builders))
	mutators := make([]Mutator, len(chcb.builders))
	for i := range chcb.builders {
		func(i int, root context.Context) {
			builder := chcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContactHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, chcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, chcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, chcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (chcb *ContactHistoryCreateBulk) SaveX(ctx context.Context) []*ContactHistory {
	v, err := chcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (chcb *ContactHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := chcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (chcb *ContactHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := chcb.Exec(ctx); err != nil {
		panic(err)
	}
}
