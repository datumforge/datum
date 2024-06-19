// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/templatehistory"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/enthistory"
)

// TemplateHistoryCreate is the builder for creating a TemplateHistory entity.
type TemplateHistoryCreate struct {
	config
	mutation *TemplateHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (thc *TemplateHistoryCreate) SetHistoryTime(t time.Time) *TemplateHistoryCreate {
	thc.mutation.SetHistoryTime(t)
	return thc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableHistoryTime(t *time.Time) *TemplateHistoryCreate {
	if t != nil {
		thc.SetHistoryTime(*t)
	}
	return thc
}

// SetOperation sets the "operation" field.
func (thc *TemplateHistoryCreate) SetOperation(et enthistory.OpType) *TemplateHistoryCreate {
	thc.mutation.SetOperation(et)
	return thc
}

// SetRef sets the "ref" field.
func (thc *TemplateHistoryCreate) SetRef(s string) *TemplateHistoryCreate {
	thc.mutation.SetRef(s)
	return thc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableRef(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetRef(*s)
	}
	return thc
}

// SetCreatedAt sets the "created_at" field.
func (thc *TemplateHistoryCreate) SetCreatedAt(t time.Time) *TemplateHistoryCreate {
	thc.mutation.SetCreatedAt(t)
	return thc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableCreatedAt(t *time.Time) *TemplateHistoryCreate {
	if t != nil {
		thc.SetCreatedAt(*t)
	}
	return thc
}

// SetUpdatedAt sets the "updated_at" field.
func (thc *TemplateHistoryCreate) SetUpdatedAt(t time.Time) *TemplateHistoryCreate {
	thc.mutation.SetUpdatedAt(t)
	return thc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableUpdatedAt(t *time.Time) *TemplateHistoryCreate {
	if t != nil {
		thc.SetUpdatedAt(*t)
	}
	return thc
}

// SetCreatedBy sets the "created_by" field.
func (thc *TemplateHistoryCreate) SetCreatedBy(s string) *TemplateHistoryCreate {
	thc.mutation.SetCreatedBy(s)
	return thc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableCreatedBy(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetCreatedBy(*s)
	}
	return thc
}

// SetUpdatedBy sets the "updated_by" field.
func (thc *TemplateHistoryCreate) SetUpdatedBy(s string) *TemplateHistoryCreate {
	thc.mutation.SetUpdatedBy(s)
	return thc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableUpdatedBy(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetUpdatedBy(*s)
	}
	return thc
}

// SetDeletedAt sets the "deleted_at" field.
func (thc *TemplateHistoryCreate) SetDeletedAt(t time.Time) *TemplateHistoryCreate {
	thc.mutation.SetDeletedAt(t)
	return thc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableDeletedAt(t *time.Time) *TemplateHistoryCreate {
	if t != nil {
		thc.SetDeletedAt(*t)
	}
	return thc
}

// SetDeletedBy sets the "deleted_by" field.
func (thc *TemplateHistoryCreate) SetDeletedBy(s string) *TemplateHistoryCreate {
	thc.mutation.SetDeletedBy(s)
	return thc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableDeletedBy(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetDeletedBy(*s)
	}
	return thc
}

// SetMappingID sets the "mapping_id" field.
func (thc *TemplateHistoryCreate) SetMappingID(s string) *TemplateHistoryCreate {
	thc.mutation.SetMappingID(s)
	return thc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableMappingID(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetMappingID(*s)
	}
	return thc
}

// SetTags sets the "tags" field.
func (thc *TemplateHistoryCreate) SetTags(s []string) *TemplateHistoryCreate {
	thc.mutation.SetTags(s)
	return thc
}

// SetOwnerID sets the "owner_id" field.
func (thc *TemplateHistoryCreate) SetOwnerID(s string) *TemplateHistoryCreate {
	thc.mutation.SetOwnerID(s)
	return thc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableOwnerID(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetOwnerID(*s)
	}
	return thc
}

// SetName sets the "name" field.
func (thc *TemplateHistoryCreate) SetName(s string) *TemplateHistoryCreate {
	thc.mutation.SetName(s)
	return thc
}

// SetTemplateType sets the "template_type" field.
func (thc *TemplateHistoryCreate) SetTemplateType(et enums.DocumentType) *TemplateHistoryCreate {
	thc.mutation.SetTemplateType(et)
	return thc
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableTemplateType(et *enums.DocumentType) *TemplateHistoryCreate {
	if et != nil {
		thc.SetTemplateType(*et)
	}
	return thc
}

// SetDescription sets the "description" field.
func (thc *TemplateHistoryCreate) SetDescription(s string) *TemplateHistoryCreate {
	thc.mutation.SetDescription(s)
	return thc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableDescription(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetDescription(*s)
	}
	return thc
}

// SetJsonconfig sets the "jsonconfig" field.
func (thc *TemplateHistoryCreate) SetJsonconfig(co customtypes.JSONObject) *TemplateHistoryCreate {
	thc.mutation.SetJsonconfig(co)
	return thc
}

// SetUischema sets the "uischema" field.
func (thc *TemplateHistoryCreate) SetUischema(co customtypes.JSONObject) *TemplateHistoryCreate {
	thc.mutation.SetUischema(co)
	return thc
}

// SetID sets the "id" field.
func (thc *TemplateHistoryCreate) SetID(s string) *TemplateHistoryCreate {
	thc.mutation.SetID(s)
	return thc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (thc *TemplateHistoryCreate) SetNillableID(s *string) *TemplateHistoryCreate {
	if s != nil {
		thc.SetID(*s)
	}
	return thc
}

// Mutation returns the TemplateHistoryMutation object of the builder.
func (thc *TemplateHistoryCreate) Mutation() *TemplateHistoryMutation {
	return thc.mutation
}

// Save creates the TemplateHistory in the database.
func (thc *TemplateHistoryCreate) Save(ctx context.Context) (*TemplateHistory, error) {
	thc.defaults()
	return withHooks(ctx, thc.sqlSave, thc.mutation, thc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (thc *TemplateHistoryCreate) SaveX(ctx context.Context) *TemplateHistory {
	v, err := thc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thc *TemplateHistoryCreate) Exec(ctx context.Context) error {
	_, err := thc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thc *TemplateHistoryCreate) ExecX(ctx context.Context) {
	if err := thc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (thc *TemplateHistoryCreate) defaults() {
	if _, ok := thc.mutation.HistoryTime(); !ok {
		v := templatehistory.DefaultHistoryTime()
		thc.mutation.SetHistoryTime(v)
	}
	if _, ok := thc.mutation.CreatedAt(); !ok {
		v := templatehistory.DefaultCreatedAt()
		thc.mutation.SetCreatedAt(v)
	}
	if _, ok := thc.mutation.UpdatedAt(); !ok {
		v := templatehistory.DefaultUpdatedAt()
		thc.mutation.SetUpdatedAt(v)
	}
	if _, ok := thc.mutation.MappingID(); !ok {
		v := templatehistory.DefaultMappingID()
		thc.mutation.SetMappingID(v)
	}
	if _, ok := thc.mutation.Tags(); !ok {
		v := templatehistory.DefaultTags
		thc.mutation.SetTags(v)
	}
	if _, ok := thc.mutation.TemplateType(); !ok {
		v := templatehistory.DefaultTemplateType
		thc.mutation.SetTemplateType(v)
	}
	if _, ok := thc.mutation.ID(); !ok {
		v := templatehistory.DefaultID()
		thc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thc *TemplateHistoryCreate) check() error {
	if _, ok := thc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "TemplateHistory.history_time"`)}
	}
	if _, ok := thc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "TemplateHistory.operation"`)}
	}
	if v, ok := thc.mutation.Operation(); ok {
		if err := templatehistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "TemplateHistory.operation": %w`, err)}
		}
	}
	if _, ok := thc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "TemplateHistory.mapping_id"`)}
	}
	if _, ok := thc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "TemplateHistory.name"`)}
	}
	if _, ok := thc.mutation.TemplateType(); !ok {
		return &ValidationError{Name: "template_type", err: errors.New(`generated: missing required field "TemplateHistory.template_type"`)}
	}
	if v, ok := thc.mutation.TemplateType(); ok {
		if err := templatehistory.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`generated: validator failed for field "TemplateHistory.template_type": %w`, err)}
		}
	}
	if _, ok := thc.mutation.Jsonconfig(); !ok {
		return &ValidationError{Name: "jsonconfig", err: errors.New(`generated: missing required field "TemplateHistory.jsonconfig"`)}
	}
	return nil
}

func (thc *TemplateHistoryCreate) sqlSave(ctx context.Context) (*TemplateHistory, error) {
	if err := thc.check(); err != nil {
		return nil, err
	}
	_node, _spec := thc.createSpec()
	if err := sqlgraph.CreateNode(ctx, thc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TemplateHistory.ID type: %T", _spec.ID.Value)
		}
	}
	thc.mutation.id = &_node.ID
	thc.mutation.done = true
	return _node, nil
}

func (thc *TemplateHistoryCreate) createSpec() (*TemplateHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &TemplateHistory{config: thc.config}
		_spec = sqlgraph.NewCreateSpec(templatehistory.Table, sqlgraph.NewFieldSpec(templatehistory.FieldID, field.TypeString))
	)
	_spec.Schema = thc.schemaConfig.TemplateHistory
	if id, ok := thc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := thc.mutation.HistoryTime(); ok {
		_spec.SetField(templatehistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := thc.mutation.Operation(); ok {
		_spec.SetField(templatehistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := thc.mutation.Ref(); ok {
		_spec.SetField(templatehistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := thc.mutation.CreatedAt(); ok {
		_spec.SetField(templatehistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := thc.mutation.UpdatedAt(); ok {
		_spec.SetField(templatehistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := thc.mutation.CreatedBy(); ok {
		_spec.SetField(templatehistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := thc.mutation.UpdatedBy(); ok {
		_spec.SetField(templatehistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := thc.mutation.DeletedAt(); ok {
		_spec.SetField(templatehistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := thc.mutation.DeletedBy(); ok {
		_spec.SetField(templatehistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := thc.mutation.MappingID(); ok {
		_spec.SetField(templatehistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := thc.mutation.Tags(); ok {
		_spec.SetField(templatehistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := thc.mutation.OwnerID(); ok {
		_spec.SetField(templatehistory.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := thc.mutation.Name(); ok {
		_spec.SetField(templatehistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := thc.mutation.TemplateType(); ok {
		_spec.SetField(templatehistory.FieldTemplateType, field.TypeEnum, value)
		_node.TemplateType = value
	}
	if value, ok := thc.mutation.Description(); ok {
		_spec.SetField(templatehistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := thc.mutation.Jsonconfig(); ok {
		_spec.SetField(templatehistory.FieldJsonconfig, field.TypeJSON, value)
		_node.Jsonconfig = value
	}
	if value, ok := thc.mutation.Uischema(); ok {
		_spec.SetField(templatehistory.FieldUischema, field.TypeJSON, value)
		_node.Uischema = value
	}
	return _node, _spec
}

// TemplateHistoryCreateBulk is the builder for creating many TemplateHistory entities in bulk.
type TemplateHistoryCreateBulk struct {
	config
	err      error
	builders []*TemplateHistoryCreate
}

// Save creates the TemplateHistory entities in the database.
func (thcb *TemplateHistoryCreateBulk) Save(ctx context.Context) ([]*TemplateHistory, error) {
	if thcb.err != nil {
		return nil, thcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(thcb.builders))
	nodes := make([]*TemplateHistory, len(thcb.builders))
	mutators := make([]Mutator, len(thcb.builders))
	for i := range thcb.builders {
		func(i int, root context.Context) {
			builder := thcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, thcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, thcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, thcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (thcb *TemplateHistoryCreateBulk) SaveX(ctx context.Context) []*TemplateHistory {
	v, err := thcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (thcb *TemplateHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := thcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thcb *TemplateHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := thcb.Exec(ctx); err != nil {
		panic(err)
	}
}
