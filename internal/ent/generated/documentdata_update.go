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
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/documentdata"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/template"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// DocumentDataUpdate is the builder for updating DocumentData entities.
type DocumentDataUpdate struct {
	config
	hooks    []Hook
	mutation *DocumentDataMutation
}

// Where appends a list predicates to the DocumentDataUpdate builder.
func (ddu *DocumentDataUpdate) Where(ps ...predicate.DocumentData) *DocumentDataUpdate {
	ddu.mutation.Where(ps...)
	return ddu
}

// SetUpdatedAt sets the "updated_at" field.
func (ddu *DocumentDataUpdate) SetUpdatedAt(t time.Time) *DocumentDataUpdate {
	ddu.mutation.SetUpdatedAt(t)
	return ddu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ddu *DocumentDataUpdate) ClearUpdatedAt() *DocumentDataUpdate {
	ddu.mutation.ClearUpdatedAt()
	return ddu
}

// SetUpdatedBy sets the "updated_by" field.
func (ddu *DocumentDataUpdate) SetUpdatedBy(s string) *DocumentDataUpdate {
	ddu.mutation.SetUpdatedBy(s)
	return ddu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ddu *DocumentDataUpdate) SetNillableUpdatedBy(s *string) *DocumentDataUpdate {
	if s != nil {
		ddu.SetUpdatedBy(*s)
	}
	return ddu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ddu *DocumentDataUpdate) ClearUpdatedBy() *DocumentDataUpdate {
	ddu.mutation.ClearUpdatedBy()
	return ddu
}

// SetDeletedAt sets the "deleted_at" field.
func (ddu *DocumentDataUpdate) SetDeletedAt(t time.Time) *DocumentDataUpdate {
	ddu.mutation.SetDeletedAt(t)
	return ddu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ddu *DocumentDataUpdate) SetNillableDeletedAt(t *time.Time) *DocumentDataUpdate {
	if t != nil {
		ddu.SetDeletedAt(*t)
	}
	return ddu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ddu *DocumentDataUpdate) ClearDeletedAt() *DocumentDataUpdate {
	ddu.mutation.ClearDeletedAt()
	return ddu
}

// SetDeletedBy sets the "deleted_by" field.
func (ddu *DocumentDataUpdate) SetDeletedBy(s string) *DocumentDataUpdate {
	ddu.mutation.SetDeletedBy(s)
	return ddu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ddu *DocumentDataUpdate) SetNillableDeletedBy(s *string) *DocumentDataUpdate {
	if s != nil {
		ddu.SetDeletedBy(*s)
	}
	return ddu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ddu *DocumentDataUpdate) ClearDeletedBy() *DocumentDataUpdate {
	ddu.mutation.ClearDeletedBy()
	return ddu
}

// SetTemplateID sets the "template_id" field.
func (ddu *DocumentDataUpdate) SetTemplateID(s string) *DocumentDataUpdate {
	ddu.mutation.SetTemplateID(s)
	return ddu
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (ddu *DocumentDataUpdate) SetNillableTemplateID(s *string) *DocumentDataUpdate {
	if s != nil {
		ddu.SetTemplateID(*s)
	}
	return ddu
}

// SetData sets the "data" field.
func (ddu *DocumentDataUpdate) SetData(co customtypes.JSONObject) *DocumentDataUpdate {
	ddu.mutation.SetData(co)
	return ddu
}

// SetTemplate sets the "template" edge to the Template entity.
func (ddu *DocumentDataUpdate) SetTemplate(t *Template) *DocumentDataUpdate {
	return ddu.SetTemplateID(t.ID)
}

// Mutation returns the DocumentDataMutation object of the builder.
func (ddu *DocumentDataUpdate) Mutation() *DocumentDataMutation {
	return ddu.mutation
}

// ClearTemplate clears the "template" edge to the Template entity.
func (ddu *DocumentDataUpdate) ClearTemplate() *DocumentDataUpdate {
	ddu.mutation.ClearTemplate()
	return ddu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddu *DocumentDataUpdate) Save(ctx context.Context) (int, error) {
	if err := ddu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ddu.sqlSave, ddu.mutation, ddu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ddu *DocumentDataUpdate) SaveX(ctx context.Context) int {
	affected, err := ddu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddu *DocumentDataUpdate) Exec(ctx context.Context) error {
	_, err := ddu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddu *DocumentDataUpdate) ExecX(ctx context.Context) {
	if err := ddu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddu *DocumentDataUpdate) defaults() error {
	if _, ok := ddu.mutation.UpdatedAt(); !ok && !ddu.mutation.UpdatedAtCleared() {
		if documentdata.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized documentdata.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := documentdata.UpdateDefaultUpdatedAt()
		ddu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ddu *DocumentDataUpdate) check() error {
	if _, ok := ddu.mutation.TemplateID(); ddu.mutation.TemplateCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "DocumentData.template"`)
	}
	return nil
}

func (ddu *DocumentDataUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ddu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(documentdata.Table, documentdata.Columns, sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString))
	if ps := ddu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ddu.mutation.CreatedAtCleared() {
		_spec.ClearField(documentdata.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ddu.mutation.UpdatedAt(); ok {
		_spec.SetField(documentdata.FieldUpdatedAt, field.TypeTime, value)
	}
	if ddu.mutation.UpdatedAtCleared() {
		_spec.ClearField(documentdata.FieldUpdatedAt, field.TypeTime)
	}
	if ddu.mutation.CreatedByCleared() {
		_spec.ClearField(documentdata.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ddu.mutation.UpdatedBy(); ok {
		_spec.SetField(documentdata.FieldUpdatedBy, field.TypeString, value)
	}
	if ddu.mutation.UpdatedByCleared() {
		_spec.ClearField(documentdata.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ddu.mutation.DeletedAt(); ok {
		_spec.SetField(documentdata.FieldDeletedAt, field.TypeTime, value)
	}
	if ddu.mutation.DeletedAtCleared() {
		_spec.ClearField(documentdata.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ddu.mutation.DeletedBy(); ok {
		_spec.SetField(documentdata.FieldDeletedBy, field.TypeString, value)
	}
	if ddu.mutation.DeletedByCleared() {
		_spec.ClearField(documentdata.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ddu.mutation.Data(); ok {
		_spec.SetField(documentdata.FieldData, field.TypeJSON, value)
	}
	if ddu.mutation.TemplateCleared() {
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
		edge.Schema = ddu.schemaConfig.DocumentData
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ddu.mutation.TemplateIDs(); len(nodes) > 0 {
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
		edge.Schema = ddu.schemaConfig.DocumentData
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = ddu.schemaConfig.DocumentData
	ctx = internal.NewSchemaConfigContext(ctx, ddu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ddu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{documentdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ddu.mutation.done = true
	return n, nil
}

// DocumentDataUpdateOne is the builder for updating a single DocumentData entity.
type DocumentDataUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DocumentDataMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (dduo *DocumentDataUpdateOne) SetUpdatedAt(t time.Time) *DocumentDataUpdateOne {
	dduo.mutation.SetUpdatedAt(t)
	return dduo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (dduo *DocumentDataUpdateOne) ClearUpdatedAt() *DocumentDataUpdateOne {
	dduo.mutation.ClearUpdatedAt()
	return dduo
}

// SetUpdatedBy sets the "updated_by" field.
func (dduo *DocumentDataUpdateOne) SetUpdatedBy(s string) *DocumentDataUpdateOne {
	dduo.mutation.SetUpdatedBy(s)
	return dduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (dduo *DocumentDataUpdateOne) SetNillableUpdatedBy(s *string) *DocumentDataUpdateOne {
	if s != nil {
		dduo.SetUpdatedBy(*s)
	}
	return dduo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (dduo *DocumentDataUpdateOne) ClearUpdatedBy() *DocumentDataUpdateOne {
	dduo.mutation.ClearUpdatedBy()
	return dduo
}

// SetDeletedAt sets the "deleted_at" field.
func (dduo *DocumentDataUpdateOne) SetDeletedAt(t time.Time) *DocumentDataUpdateOne {
	dduo.mutation.SetDeletedAt(t)
	return dduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dduo *DocumentDataUpdateOne) SetNillableDeletedAt(t *time.Time) *DocumentDataUpdateOne {
	if t != nil {
		dduo.SetDeletedAt(*t)
	}
	return dduo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (dduo *DocumentDataUpdateOne) ClearDeletedAt() *DocumentDataUpdateOne {
	dduo.mutation.ClearDeletedAt()
	return dduo
}

// SetDeletedBy sets the "deleted_by" field.
func (dduo *DocumentDataUpdateOne) SetDeletedBy(s string) *DocumentDataUpdateOne {
	dduo.mutation.SetDeletedBy(s)
	return dduo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (dduo *DocumentDataUpdateOne) SetNillableDeletedBy(s *string) *DocumentDataUpdateOne {
	if s != nil {
		dduo.SetDeletedBy(*s)
	}
	return dduo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (dduo *DocumentDataUpdateOne) ClearDeletedBy() *DocumentDataUpdateOne {
	dduo.mutation.ClearDeletedBy()
	return dduo
}

// SetTemplateID sets the "template_id" field.
func (dduo *DocumentDataUpdateOne) SetTemplateID(s string) *DocumentDataUpdateOne {
	dduo.mutation.SetTemplateID(s)
	return dduo
}

// SetNillableTemplateID sets the "template_id" field if the given value is not nil.
func (dduo *DocumentDataUpdateOne) SetNillableTemplateID(s *string) *DocumentDataUpdateOne {
	if s != nil {
		dduo.SetTemplateID(*s)
	}
	return dduo
}

// SetData sets the "data" field.
func (dduo *DocumentDataUpdateOne) SetData(co customtypes.JSONObject) *DocumentDataUpdateOne {
	dduo.mutation.SetData(co)
	return dduo
}

// SetTemplate sets the "template" edge to the Template entity.
func (dduo *DocumentDataUpdateOne) SetTemplate(t *Template) *DocumentDataUpdateOne {
	return dduo.SetTemplateID(t.ID)
}

// Mutation returns the DocumentDataMutation object of the builder.
func (dduo *DocumentDataUpdateOne) Mutation() *DocumentDataMutation {
	return dduo.mutation
}

// ClearTemplate clears the "template" edge to the Template entity.
func (dduo *DocumentDataUpdateOne) ClearTemplate() *DocumentDataUpdateOne {
	dduo.mutation.ClearTemplate()
	return dduo
}

// Where appends a list predicates to the DocumentDataUpdate builder.
func (dduo *DocumentDataUpdateOne) Where(ps ...predicate.DocumentData) *DocumentDataUpdateOne {
	dduo.mutation.Where(ps...)
	return dduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dduo *DocumentDataUpdateOne) Select(field string, fields ...string) *DocumentDataUpdateOne {
	dduo.fields = append([]string{field}, fields...)
	return dduo
}

// Save executes the query and returns the updated DocumentData entity.
func (dduo *DocumentDataUpdateOne) Save(ctx context.Context) (*DocumentData, error) {
	if err := dduo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, dduo.sqlSave, dduo.mutation, dduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dduo *DocumentDataUpdateOne) SaveX(ctx context.Context) *DocumentData {
	node, err := dduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dduo *DocumentDataUpdateOne) Exec(ctx context.Context) error {
	_, err := dduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dduo *DocumentDataUpdateOne) ExecX(ctx context.Context) {
	if err := dduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dduo *DocumentDataUpdateOne) defaults() error {
	if _, ok := dduo.mutation.UpdatedAt(); !ok && !dduo.mutation.UpdatedAtCleared() {
		if documentdata.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized documentdata.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := documentdata.UpdateDefaultUpdatedAt()
		dduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dduo *DocumentDataUpdateOne) check() error {
	if _, ok := dduo.mutation.TemplateID(); dduo.mutation.TemplateCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "DocumentData.template"`)
	}
	return nil
}

func (dduo *DocumentDataUpdateOne) sqlSave(ctx context.Context) (_node *DocumentData, err error) {
	if err := dduo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(documentdata.Table, documentdata.Columns, sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString))
	id, ok := dduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "DocumentData.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, documentdata.FieldID)
		for _, f := range fields {
			if !documentdata.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != documentdata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if dduo.mutation.CreatedAtCleared() {
		_spec.ClearField(documentdata.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := dduo.mutation.UpdatedAt(); ok {
		_spec.SetField(documentdata.FieldUpdatedAt, field.TypeTime, value)
	}
	if dduo.mutation.UpdatedAtCleared() {
		_spec.ClearField(documentdata.FieldUpdatedAt, field.TypeTime)
	}
	if dduo.mutation.CreatedByCleared() {
		_spec.ClearField(documentdata.FieldCreatedBy, field.TypeString)
	}
	if value, ok := dduo.mutation.UpdatedBy(); ok {
		_spec.SetField(documentdata.FieldUpdatedBy, field.TypeString, value)
	}
	if dduo.mutation.UpdatedByCleared() {
		_spec.ClearField(documentdata.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := dduo.mutation.DeletedAt(); ok {
		_spec.SetField(documentdata.FieldDeletedAt, field.TypeTime, value)
	}
	if dduo.mutation.DeletedAtCleared() {
		_spec.ClearField(documentdata.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := dduo.mutation.DeletedBy(); ok {
		_spec.SetField(documentdata.FieldDeletedBy, field.TypeString, value)
	}
	if dduo.mutation.DeletedByCleared() {
		_spec.ClearField(documentdata.FieldDeletedBy, field.TypeString)
	}
	if value, ok := dduo.mutation.Data(); ok {
		_spec.SetField(documentdata.FieldData, field.TypeJSON, value)
	}
	if dduo.mutation.TemplateCleared() {
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
		edge.Schema = dduo.schemaConfig.DocumentData
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dduo.mutation.TemplateIDs(); len(nodes) > 0 {
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
		edge.Schema = dduo.schemaConfig.DocumentData
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = dduo.schemaConfig.DocumentData
	ctx = internal.NewSchemaConfigContext(ctx, dduo.schemaConfig)
	_node = &DocumentData{config: dduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{documentdata.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dduo.mutation.done = true
	return _node, nil
}
