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
	"github.com/datumforge/datum/internal/ent/generated/documentdatahistory"
	"github.com/datumforge/enthistory"
)

// DocumentDataHistoryCreate is the builder for creating a DocumentDataHistory entity.
type DocumentDataHistoryCreate struct {
	config
	mutation *DocumentDataHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ddhc *DocumentDataHistoryCreate) SetHistoryTime(t time.Time) *DocumentDataHistoryCreate {
	ddhc.mutation.SetHistoryTime(t)
	return ddhc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableHistoryTime(t *time.Time) *DocumentDataHistoryCreate {
	if t != nil {
		ddhc.SetHistoryTime(*t)
	}
	return ddhc
}

// SetOperation sets the "operation" field.
func (ddhc *DocumentDataHistoryCreate) SetOperation(et enthistory.OpType) *DocumentDataHistoryCreate {
	ddhc.mutation.SetOperation(et)
	return ddhc
}

// SetRef sets the "ref" field.
func (ddhc *DocumentDataHistoryCreate) SetRef(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetRef(s)
	return ddhc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableRef(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetRef(*s)
	}
	return ddhc
}

// SetCreatedAt sets the "created_at" field.
func (ddhc *DocumentDataHistoryCreate) SetCreatedAt(t time.Time) *DocumentDataHistoryCreate {
	ddhc.mutation.SetCreatedAt(t)
	return ddhc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableCreatedAt(t *time.Time) *DocumentDataHistoryCreate {
	if t != nil {
		ddhc.SetCreatedAt(*t)
	}
	return ddhc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddhc *DocumentDataHistoryCreate) SetUpdatedAt(t time.Time) *DocumentDataHistoryCreate {
	ddhc.mutation.SetUpdatedAt(t)
	return ddhc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableUpdatedAt(t *time.Time) *DocumentDataHistoryCreate {
	if t != nil {
		ddhc.SetUpdatedAt(*t)
	}
	return ddhc
}

// SetCreatedBy sets the "created_by" field.
func (ddhc *DocumentDataHistoryCreate) SetCreatedBy(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetCreatedBy(s)
	return ddhc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableCreatedBy(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetCreatedBy(*s)
	}
	return ddhc
}

// SetUpdatedBy sets the "updated_by" field.
func (ddhc *DocumentDataHistoryCreate) SetUpdatedBy(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetUpdatedBy(s)
	return ddhc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableUpdatedBy(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetUpdatedBy(*s)
	}
	return ddhc
}

// SetMappingID sets the "mapping_id" field.
func (ddhc *DocumentDataHistoryCreate) SetMappingID(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetMappingID(s)
	return ddhc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableMappingID(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetMappingID(*s)
	}
	return ddhc
}

// SetDeletedAt sets the "deleted_at" field.
func (ddhc *DocumentDataHistoryCreate) SetDeletedAt(t time.Time) *DocumentDataHistoryCreate {
	ddhc.mutation.SetDeletedAt(t)
	return ddhc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableDeletedAt(t *time.Time) *DocumentDataHistoryCreate {
	if t != nil {
		ddhc.SetDeletedAt(*t)
	}
	return ddhc
}

// SetDeletedBy sets the "deleted_by" field.
func (ddhc *DocumentDataHistoryCreate) SetDeletedBy(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetDeletedBy(s)
	return ddhc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableDeletedBy(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetDeletedBy(*s)
	}
	return ddhc
}

// SetTemplateID sets the "template_id" field.
func (ddhc *DocumentDataHistoryCreate) SetTemplateID(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetTemplateID(s)
	return ddhc
}

// SetData sets the "data" field.
func (ddhc *DocumentDataHistoryCreate) SetData(co customtypes.JSONObject) *DocumentDataHistoryCreate {
	ddhc.mutation.SetData(co)
	return ddhc
}

// SetID sets the "id" field.
func (ddhc *DocumentDataHistoryCreate) SetID(s string) *DocumentDataHistoryCreate {
	ddhc.mutation.SetID(s)
	return ddhc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ddhc *DocumentDataHistoryCreate) SetNillableID(s *string) *DocumentDataHistoryCreate {
	if s != nil {
		ddhc.SetID(*s)
	}
	return ddhc
}

// Mutation returns the DocumentDataHistoryMutation object of the builder.
func (ddhc *DocumentDataHistoryCreate) Mutation() *DocumentDataHistoryMutation {
	return ddhc.mutation
}

// Save creates the DocumentDataHistory in the database.
func (ddhc *DocumentDataHistoryCreate) Save(ctx context.Context) (*DocumentDataHistory, error) {
	ddhc.defaults()
	return withHooks(ctx, ddhc.sqlSave, ddhc.mutation, ddhc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddhc *DocumentDataHistoryCreate) SaveX(ctx context.Context) *DocumentDataHistory {
	v, err := ddhc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddhc *DocumentDataHistoryCreate) Exec(ctx context.Context) error {
	_, err := ddhc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddhc *DocumentDataHistoryCreate) ExecX(ctx context.Context) {
	if err := ddhc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddhc *DocumentDataHistoryCreate) defaults() {
	if _, ok := ddhc.mutation.HistoryTime(); !ok {
		v := documentdatahistory.DefaultHistoryTime()
		ddhc.mutation.SetHistoryTime(v)
	}
	if _, ok := ddhc.mutation.CreatedAt(); !ok {
		v := documentdatahistory.DefaultCreatedAt()
		ddhc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddhc.mutation.UpdatedAt(); !ok {
		v := documentdatahistory.DefaultUpdatedAt()
		ddhc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ddhc.mutation.MappingID(); !ok {
		v := documentdatahistory.DefaultMappingID()
		ddhc.mutation.SetMappingID(v)
	}
	if _, ok := ddhc.mutation.ID(); !ok {
		v := documentdatahistory.DefaultID()
		ddhc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddhc *DocumentDataHistoryCreate) check() error {
	if _, ok := ddhc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "DocumentDataHistory.history_time"`)}
	}
	if _, ok := ddhc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "DocumentDataHistory.operation"`)}
	}
	if v, ok := ddhc.mutation.Operation(); ok {
		if err := documentdatahistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "DocumentDataHistory.operation": %w`, err)}
		}
	}
	if _, ok := ddhc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "DocumentDataHistory.mapping_id"`)}
	}
	if _, ok := ddhc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template_id", err: errors.New(`generated: missing required field "DocumentDataHistory.template_id"`)}
	}
	if _, ok := ddhc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`generated: missing required field "DocumentDataHistory.data"`)}
	}
	return nil
}

func (ddhc *DocumentDataHistoryCreate) sqlSave(ctx context.Context) (*DocumentDataHistory, error) {
	if err := ddhc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddhc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddhc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DocumentDataHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ddhc.mutation.id = &_node.ID
	ddhc.mutation.done = true
	return _node, nil
}

func (ddhc *DocumentDataHistoryCreate) createSpec() (*DocumentDataHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &DocumentDataHistory{config: ddhc.config}
		_spec = sqlgraph.NewCreateSpec(documentdatahistory.Table, sqlgraph.NewFieldSpec(documentdatahistory.FieldID, field.TypeString))
	)
	_spec.Schema = ddhc.schemaConfig.DocumentDataHistory
	if id, ok := ddhc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ddhc.mutation.HistoryTime(); ok {
		_spec.SetField(documentdatahistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ddhc.mutation.Operation(); ok {
		_spec.SetField(documentdatahistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ddhc.mutation.Ref(); ok {
		_spec.SetField(documentdatahistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ddhc.mutation.CreatedAt(); ok {
		_spec.SetField(documentdatahistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ddhc.mutation.UpdatedAt(); ok {
		_spec.SetField(documentdatahistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ddhc.mutation.CreatedBy(); ok {
		_spec.SetField(documentdatahistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ddhc.mutation.UpdatedBy(); ok {
		_spec.SetField(documentdatahistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ddhc.mutation.MappingID(); ok {
		_spec.SetField(documentdatahistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := ddhc.mutation.DeletedAt(); ok {
		_spec.SetField(documentdatahistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ddhc.mutation.DeletedBy(); ok {
		_spec.SetField(documentdatahistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ddhc.mutation.TemplateID(); ok {
		_spec.SetField(documentdatahistory.FieldTemplateID, field.TypeString, value)
		_node.TemplateID = value
	}
	if value, ok := ddhc.mutation.Data(); ok {
		_spec.SetField(documentdatahistory.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	return _node, _spec
}

// DocumentDataHistoryCreateBulk is the builder for creating many DocumentDataHistory entities in bulk.
type DocumentDataHistoryCreateBulk struct {
	config
	err      error
	builders []*DocumentDataHistoryCreate
}

// Save creates the DocumentDataHistory entities in the database.
func (ddhcb *DocumentDataHistoryCreateBulk) Save(ctx context.Context) ([]*DocumentDataHistory, error) {
	if ddhcb.err != nil {
		return nil, ddhcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ddhcb.builders))
	nodes := make([]*DocumentDataHistory, len(ddhcb.builders))
	mutators := make([]Mutator, len(ddhcb.builders))
	for i := range ddhcb.builders {
		func(i int, root context.Context) {
			builder := ddhcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocumentDataHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ddhcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddhcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ddhcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddhcb *DocumentDataHistoryCreateBulk) SaveX(ctx context.Context) []*DocumentDataHistory {
	v, err := ddhcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddhcb *DocumentDataHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ddhcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddhcb *DocumentDataHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ddhcb.Exec(ctx); err != nil {
		panic(err)
	}
}
