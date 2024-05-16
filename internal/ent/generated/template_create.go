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
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/documentdata"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/template"
)

// TemplateCreate is the builder for creating a Template entity.
type TemplateCreate struct {
	config
	mutation *TemplateMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (tc *TemplateCreate) SetCreatedAt(t time.Time) *TemplateCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreatedAt(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TemplateCreate) SetUpdatedAt(t time.Time) *TemplateCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableUpdatedAt(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetCreatedBy sets the "created_by" field.
func (tc *TemplateCreate) SetCreatedBy(s string) *TemplateCreate {
	tc.mutation.SetCreatedBy(s)
	return tc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableCreatedBy(s *string) *TemplateCreate {
	if s != nil {
		tc.SetCreatedBy(*s)
	}
	return tc
}

// SetUpdatedBy sets the "updated_by" field.
func (tc *TemplateCreate) SetUpdatedBy(s string) *TemplateCreate {
	tc.mutation.SetUpdatedBy(s)
	return tc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableUpdatedBy(s *string) *TemplateCreate {
	if s != nil {
		tc.SetUpdatedBy(*s)
	}
	return tc
}

// SetDeletedAt sets the "deleted_at" field.
func (tc *TemplateCreate) SetDeletedAt(t time.Time) *TemplateCreate {
	tc.mutation.SetDeletedAt(t)
	return tc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableDeletedAt(t *time.Time) *TemplateCreate {
	if t != nil {
		tc.SetDeletedAt(*t)
	}
	return tc
}

// SetDeletedBy sets the "deleted_by" field.
func (tc *TemplateCreate) SetDeletedBy(s string) *TemplateCreate {
	tc.mutation.SetDeletedBy(s)
	return tc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableDeletedBy(s *string) *TemplateCreate {
	if s != nil {
		tc.SetDeletedBy(*s)
	}
	return tc
}

// SetMappingID sets the "mapping_id" field.
func (tc *TemplateCreate) SetMappingID(s string) *TemplateCreate {
	tc.mutation.SetMappingID(s)
	return tc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableMappingID(s *string) *TemplateCreate {
	if s != nil {
		tc.SetMappingID(*s)
	}
	return tc
}

// SetTags sets the "tags" field.
func (tc *TemplateCreate) SetTags(s []string) *TemplateCreate {
	tc.mutation.SetTags(s)
	return tc
}

// SetOwnerID sets the "owner_id" field.
func (tc *TemplateCreate) SetOwnerID(s string) *TemplateCreate {
	tc.mutation.SetOwnerID(s)
	return tc
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableOwnerID(s *string) *TemplateCreate {
	if s != nil {
		tc.SetOwnerID(*s)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TemplateCreate) SetName(s string) *TemplateCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetTemplateType sets the "template_type" field.
func (tc *TemplateCreate) SetTemplateType(et enums.DocumentType) *TemplateCreate {
	tc.mutation.SetTemplateType(et)
	return tc
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableTemplateType(et *enums.DocumentType) *TemplateCreate {
	if et != nil {
		tc.SetTemplateType(*et)
	}
	return tc
}

// SetDescription sets the "description" field.
func (tc *TemplateCreate) SetDescription(s string) *TemplateCreate {
	tc.mutation.SetDescription(s)
	return tc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableDescription(s *string) *TemplateCreate {
	if s != nil {
		tc.SetDescription(*s)
	}
	return tc
}

// SetJsonconfig sets the "jsonconfig" field.
func (tc *TemplateCreate) SetJsonconfig(co customtypes.JSONObject) *TemplateCreate {
	tc.mutation.SetJsonconfig(co)
	return tc
}

// SetUischema sets the "uischema" field.
func (tc *TemplateCreate) SetUischema(co customtypes.JSONObject) *TemplateCreate {
	tc.mutation.SetUischema(co)
	return tc
}

// SetID sets the "id" field.
func (tc *TemplateCreate) SetID(s string) *TemplateCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tc *TemplateCreate) SetNillableID(s *string) *TemplateCreate {
	if s != nil {
		tc.SetID(*s)
	}
	return tc
}

// SetOwner sets the "owner" edge to the Organization entity.
func (tc *TemplateCreate) SetOwner(o *Organization) *TemplateCreate {
	return tc.SetOwnerID(o.ID)
}

// AddDocumentIDs adds the "documents" edge to the DocumentData entity by IDs.
func (tc *TemplateCreate) AddDocumentIDs(ids ...string) *TemplateCreate {
	tc.mutation.AddDocumentIDs(ids...)
	return tc
}

// AddDocuments adds the "documents" edges to the DocumentData entity.
func (tc *TemplateCreate) AddDocuments(d ...*DocumentData) *TemplateCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return tc.AddDocumentIDs(ids...)
}

// Mutation returns the TemplateMutation object of the builder.
func (tc *TemplateCreate) Mutation() *TemplateMutation {
	return tc.mutation
}

// Save creates the Template in the database.
func (tc *TemplateCreate) Save(ctx context.Context) (*Template, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TemplateCreate) SaveX(ctx context.Context) *Template {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TemplateCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TemplateCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TemplateCreate) defaults() error {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		if template.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized template.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := template.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		if template.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized template.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := template.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tc.mutation.MappingID(); !ok {
		if template.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized template.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := template.DefaultMappingID()
		tc.mutation.SetMappingID(v)
	}
	if _, ok := tc.mutation.Tags(); !ok {
		v := template.DefaultTags
		tc.mutation.SetTags(v)
	}
	if _, ok := tc.mutation.TemplateType(); !ok {
		v := template.DefaultTemplateType
		tc.mutation.SetTemplateType(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		if template.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized template.DefaultID (forgotten import generated/runtime?)")
		}
		v := template.DefaultID()
		tc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TemplateCreate) check() error {
	if _, ok := tc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "Template.mapping_id"`)}
	}
	if v, ok := tc.mutation.OwnerID(); ok {
		if err := template.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Template.owner_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "Template.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Template.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.TemplateType(); !ok {
		return &ValidationError{Name: "template_type", err: errors.New(`generated: missing required field "Template.template_type"`)}
	}
	if v, ok := tc.mutation.TemplateType(); ok {
		if err := template.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`generated: validator failed for field "Template.template_type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Jsonconfig(); !ok {
		return &ValidationError{Name: "jsonconfig", err: errors.New(`generated: missing required field "Template.jsonconfig"`)}
	}
	return nil
}

func (tc *TemplateCreate) sqlSave(ctx context.Context) (*Template, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Template.ID type: %T", _spec.ID.Value)
		}
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TemplateCreate) createSpec() (*Template, *sqlgraph.CreateSpec) {
	var (
		_node = &Template{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(template.Table, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	)
	_spec.Schema = tc.schemaConfig.Template
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(template.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(template.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tc.mutation.CreatedBy(); ok {
		_spec.SetField(template.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := tc.mutation.UpdatedBy(); ok {
		_spec.SetField(template.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := tc.mutation.DeletedAt(); ok {
		_spec.SetField(template.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := tc.mutation.DeletedBy(); ok {
		_spec.SetField(template.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := tc.mutation.MappingID(); ok {
		_spec.SetField(template.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := tc.mutation.Tags(); ok {
		_spec.SetField(template.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(template.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.TemplateType(); ok {
		_spec.SetField(template.FieldTemplateType, field.TypeEnum, value)
		_node.TemplateType = value
	}
	if value, ok := tc.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := tc.mutation.Jsonconfig(); ok {
		_spec.SetField(template.FieldJsonconfig, field.TypeJSON, value)
		_node.Jsonconfig = value
	}
	if value, ok := tc.mutation.Uischema(); ok {
		_spec.SetField(template.FieldUischema, field.TypeJSON, value)
		_node.Uischema = value
	}
	if nodes := tc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   template.OwnerTable,
			Columns: []string{template.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = tc.schemaConfig.Template
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   template.DocumentsTable,
			Columns: []string{template.DocumentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = tc.schemaConfig.DocumentData
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TemplateCreateBulk is the builder for creating many Template entities in bulk.
type TemplateCreateBulk struct {
	config
	err      error
	builders []*TemplateCreate
}

// Save creates the Template entities in the database.
func (tcb *TemplateCreateBulk) Save(ctx context.Context) ([]*Template, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Template, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TemplateMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TemplateCreateBulk) SaveX(ctx context.Context) []*Template {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TemplateCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
