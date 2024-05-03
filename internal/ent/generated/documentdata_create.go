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
	"github.com/datumforge/datum/internal/ent/generated/documentdata"
	"github.com/datumforge/datum/internal/ent/generated/template"
)

// DocumentDataCreate is the builder for creating a DocumentData entity.
type DocumentDataCreate struct {
	config
	mutation *DocumentDataMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ddc *DocumentDataCreate) SetCreatedAt(t time.Time) *DocumentDataCreate {
	ddc.mutation.SetCreatedAt(t)
	return ddc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableCreatedAt(t *time.Time) *DocumentDataCreate {
	if t != nil {
		ddc.SetCreatedAt(*t)
	}
	return ddc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddc *DocumentDataCreate) SetUpdatedAt(t time.Time) *DocumentDataCreate {
	ddc.mutation.SetUpdatedAt(t)
	return ddc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableUpdatedAt(t *time.Time) *DocumentDataCreate {
	if t != nil {
		ddc.SetUpdatedAt(*t)
	}
	return ddc
}

// SetCreatedBy sets the "created_by" field.
func (ddc *DocumentDataCreate) SetCreatedBy(s string) *DocumentDataCreate {
	ddc.mutation.SetCreatedBy(s)
	return ddc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableCreatedBy(s *string) *DocumentDataCreate {
	if s != nil {
		ddc.SetCreatedBy(*s)
	}
	return ddc
}

// SetUpdatedBy sets the "updated_by" field.
func (ddc *DocumentDataCreate) SetUpdatedBy(s string) *DocumentDataCreate {
	ddc.mutation.SetUpdatedBy(s)
	return ddc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableUpdatedBy(s *string) *DocumentDataCreate {
	if s != nil {
		ddc.SetUpdatedBy(*s)
	}
	return ddc
}

// SetMappingID sets the "mapping_id" field.
func (ddc *DocumentDataCreate) SetMappingID(s string) *DocumentDataCreate {
	ddc.mutation.SetMappingID(s)
	return ddc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableMappingID(s *string) *DocumentDataCreate {
	if s != nil {
		ddc.SetMappingID(*s)
	}
	return ddc
}

// SetDeletedAt sets the "deleted_at" field.
func (ddc *DocumentDataCreate) SetDeletedAt(t time.Time) *DocumentDataCreate {
	ddc.mutation.SetDeletedAt(t)
	return ddc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableDeletedAt(t *time.Time) *DocumentDataCreate {
	if t != nil {
		ddc.SetDeletedAt(*t)
	}
	return ddc
}

// SetDeletedBy sets the "deleted_by" field.
func (ddc *DocumentDataCreate) SetDeletedBy(s string) *DocumentDataCreate {
	ddc.mutation.SetDeletedBy(s)
	return ddc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableDeletedBy(s *string) *DocumentDataCreate {
	if s != nil {
		ddc.SetDeletedBy(*s)
	}
	return ddc
}

// SetTemplateID sets the "template_id" field.
func (ddc *DocumentDataCreate) SetTemplateID(s string) *DocumentDataCreate {
	ddc.mutation.SetTemplateID(s)
	return ddc
}

// SetData sets the "data" field.
func (ddc *DocumentDataCreate) SetData(co customtypes.JSONObject) *DocumentDataCreate {
	ddc.mutation.SetData(co)
	return ddc
}

// SetID sets the "id" field.
func (ddc *DocumentDataCreate) SetID(s string) *DocumentDataCreate {
	ddc.mutation.SetID(s)
	return ddc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ddc *DocumentDataCreate) SetNillableID(s *string) *DocumentDataCreate {
	if s != nil {
		ddc.SetID(*s)
	}
	return ddc
}

// SetTemplate sets the "template" edge to the Template entity.
func (ddc *DocumentDataCreate) SetTemplate(t *Template) *DocumentDataCreate {
	return ddc.SetTemplateID(t.ID)
}

// Mutation returns the DocumentDataMutation object of the builder.
func (ddc *DocumentDataCreate) Mutation() *DocumentDataMutation {
	return ddc.mutation
}

// Save creates the DocumentData in the database.
func (ddc *DocumentDataCreate) Save(ctx context.Context) (*DocumentData, error) {
	if err := ddc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ddc.sqlSave, ddc.mutation, ddc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddc *DocumentDataCreate) SaveX(ctx context.Context) *DocumentData {
	v, err := ddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddc *DocumentDataCreate) Exec(ctx context.Context) error {
	_, err := ddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddc *DocumentDataCreate) ExecX(ctx context.Context) {
	if err := ddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddc *DocumentDataCreate) defaults() error {
	if _, ok := ddc.mutation.CreatedAt(); !ok {
		if documentdata.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized documentdata.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := documentdata.DefaultCreatedAt()
		ddc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddc.mutation.UpdatedAt(); !ok {
		if documentdata.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized documentdata.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := documentdata.DefaultUpdatedAt()
		ddc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ddc.mutation.MappingID(); !ok {
		if documentdata.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized documentdata.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := documentdata.DefaultMappingID()
		ddc.mutation.SetMappingID(v)
	}
	if _, ok := ddc.mutation.ID(); !ok {
		if documentdata.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized documentdata.DefaultID (forgotten import generated/runtime?)")
		}
		v := documentdata.DefaultID()
		ddc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ddc *DocumentDataCreate) check() error {
	if _, ok := ddc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "DocumentData.mapping_id"`)}
	}
	if _, ok := ddc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template_id", err: errors.New(`generated: missing required field "DocumentData.template_id"`)}
	}
	if _, ok := ddc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`generated: missing required field "DocumentData.data"`)}
	}
	if _, ok := ddc.mutation.TemplateID(); !ok {
		return &ValidationError{Name: "template", err: errors.New(`generated: missing required edge "DocumentData.template"`)}
	}
	return nil
}

func (ddc *DocumentDataCreate) sqlSave(ctx context.Context) (*DocumentData, error) {
	if err := ddc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected DocumentData.ID type: %T", _spec.ID.Value)
		}
	}
	ddc.mutation.id = &_node.ID
	ddc.mutation.done = true
	return _node, nil
}

func (ddc *DocumentDataCreate) createSpec() (*DocumentData, *sqlgraph.CreateSpec) {
	var (
		_node = &DocumentData{config: ddc.config}
		_spec = sqlgraph.NewCreateSpec(documentdata.Table, sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString))
	)
	_spec.Schema = ddc.schemaConfig.DocumentData
	if id, ok := ddc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ddc.mutation.CreatedAt(); ok {
		_spec.SetField(documentdata.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ddc.mutation.UpdatedAt(); ok {
		_spec.SetField(documentdata.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ddc.mutation.CreatedBy(); ok {
		_spec.SetField(documentdata.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ddc.mutation.UpdatedBy(); ok {
		_spec.SetField(documentdata.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ddc.mutation.MappingID(); ok {
		_spec.SetField(documentdata.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := ddc.mutation.DeletedAt(); ok {
		_spec.SetField(documentdata.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ddc.mutation.DeletedBy(); ok {
		_spec.SetField(documentdata.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ddc.mutation.Data(); ok {
		_spec.SetField(documentdata.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if nodes := ddc.mutation.TemplateIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   documentdata.TemplateTable,
			Columns: []string{documentdata.TemplateColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(template.FieldID, field.TypeString),
			},
		}
		edge.Schema = ddc.schemaConfig.DocumentData
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TemplateID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DocumentDataCreateBulk is the builder for creating many DocumentData entities in bulk.
type DocumentDataCreateBulk struct {
	config
	err      error
	builders []*DocumentDataCreate
}

// Save creates the DocumentData entities in the database.
func (ddcb *DocumentDataCreateBulk) Save(ctx context.Context) ([]*DocumentData, error) {
	if ddcb.err != nil {
		return nil, ddcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ddcb.builders))
	nodes := make([]*DocumentData, len(ddcb.builders))
	mutators := make([]Mutator, len(ddcb.builders))
	for i := range ddcb.builders {
		func(i int, root context.Context) {
			builder := ddcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DocumentDataMutation)
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
					_, err = mutators[i+1].Mutate(root, ddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddcb *DocumentDataCreateBulk) SaveX(ctx context.Context) []*DocumentData {
	v, err := ddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddcb *DocumentDataCreateBulk) Exec(ctx context.Context) error {
	_, err := ddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddcb *DocumentDataCreateBulk) ExecX(ctx context.Context) {
	if err := ddcb.Exec(ctx); err != nil {
		panic(err)
	}
}
