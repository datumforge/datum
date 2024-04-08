// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/template"
	"github.com/datumforge/datum/internal/ent/schematype"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TemplateUpdate is the builder for updating Template entities.
type TemplateUpdate struct {
	config
	hooks    []Hook
	mutation *TemplateMutation
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tu *TemplateUpdate) Where(ps ...predicate.Template) *TemplateUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TemplateUpdate) SetUpdatedAt(t time.Time) *TemplateUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TemplateUpdate) ClearUpdatedAt() *TemplateUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TemplateUpdate) SetUpdatedBy(s string) *TemplateUpdate {
	tu.mutation.SetUpdatedBy(s)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableUpdatedBy(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetUpdatedBy(*s)
	}
	return tu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tu *TemplateUpdate) ClearUpdatedBy() *TemplateUpdate {
	tu.mutation.ClearUpdatedBy()
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TemplateUpdate) SetDeletedAt(t time.Time) *TemplateUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableDeletedAt(t *time.Time) *TemplateUpdate {
	if t != nil {
		tu.SetDeletedAt(*t)
	}
	return tu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tu *TemplateUpdate) ClearDeletedAt() *TemplateUpdate {
	tu.mutation.ClearDeletedAt()
	return tu
}

// SetDeletedBy sets the "deleted_by" field.
func (tu *TemplateUpdate) SetDeletedBy(s string) *TemplateUpdate {
	tu.mutation.SetDeletedBy(s)
	return tu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableDeletedBy(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetDeletedBy(*s)
	}
	return tu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tu *TemplateUpdate) ClearDeletedBy() *TemplateUpdate {
	tu.mutation.ClearDeletedBy()
	return tu
}

// SetOwnerID sets the "owner_id" field.
func (tu *TemplateUpdate) SetOwnerID(s string) *TemplateUpdate {
	tu.mutation.SetOwnerID(s)
	return tu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableOwnerID(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetOwnerID(*s)
	}
	return tu
}

// SetName sets the "name" field.
func (tu *TemplateUpdate) SetName(s string) *TemplateUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableName(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TemplateUpdate) SetDescription(s string) *TemplateUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillableDescription(s *string) *TemplateUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TemplateUpdate) ClearDescription() *TemplateUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetJsonconfig sets the "jsonconfig" field.
func (tu *TemplateUpdate) SetJsonconfig(co customtypes.JSONObject) *TemplateUpdate {
	tu.mutation.SetJsonconfig(co)
	return tu
}

// ClearJsonconfig clears the value of the "jsonconfig" field.
func (tu *TemplateUpdate) ClearJsonconfig() *TemplateUpdate {
	tu.mutation.ClearJsonconfig()
	return tu
}

// SetOtherconfig sets the "otherconfig" field.
func (tu *TemplateUpdate) SetOtherconfig(sc *schematype.TemplateConfig) *TemplateUpdate {
	tu.mutation.SetOtherconfig(sc)
	return tu
}

// ClearOtherconfig clears the value of the "otherconfig" field.
func (tu *TemplateUpdate) ClearOtherconfig() *TemplateUpdate {
	tu.mutation.ClearOtherconfig()
	return tu
}

// SetPair sets the "pair" field.
func (tu *TemplateUpdate) SetPair(c customtypes.Pair) *TemplateUpdate {
	tu.mutation.SetPair(c)
	return tu
}

// SetNillablePair sets the "pair" field if the given value is not nil.
func (tu *TemplateUpdate) SetNillablePair(c *customtypes.Pair) *TemplateUpdate {
	if c != nil {
		tu.SetPair(*c)
	}
	return tu
}

// SetURL sets the "url" field.
func (tu *TemplateUpdate) SetURL(u *url.URL) *TemplateUpdate {
	tu.mutation.SetURL(u)
	return tu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (tu *TemplateUpdate) SetOwner(o *Organization) *TemplateUpdate {
	return tu.SetOwnerID(o.ID)
}

// Mutation returns the TemplateMutation object of the builder.
func (tu *TemplateUpdate) Mutation() *TemplateMutation {
	return tu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (tu *TemplateUpdate) ClearOwner() *TemplateUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TemplateUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TemplateUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TemplateUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TemplateUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		if template.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized template.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := template.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TemplateUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Template.name": %w`, err)}
		}
	}
	if _, ok := tu.mutation.OwnerID(); tu.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Template.owner"`)
	}
	return nil
}

func (tu *TemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(template.Table, template.Columns, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.CreatedAtCleared() {
		_spec.ClearField(template.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(template.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(template.FieldUpdatedAt, field.TypeTime)
	}
	if tu.mutation.CreatedByCleared() {
		_spec.ClearField(template.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(template.FieldUpdatedBy, field.TypeString, value)
	}
	if tu.mutation.UpdatedByCleared() {
		_spec.ClearField(template.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(template.FieldDeletedAt, field.TypeTime, value)
	}
	if tu.mutation.DeletedAtCleared() {
		_spec.ClearField(template.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.DeletedBy(); ok {
		_spec.SetField(template.FieldDeletedBy, field.TypeString, value)
	}
	if tu.mutation.DeletedByCleared() {
		_spec.ClearField(template.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(template.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(template.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.Jsonconfig(); ok {
		_spec.SetField(template.FieldJsonconfig, field.TypeJSON, value)
	}
	if tu.mutation.JsonconfigCleared() {
		_spec.ClearField(template.FieldJsonconfig, field.TypeJSON)
	}
	if value, ok := tu.mutation.Otherconfig(); ok {
		_spec.SetField(template.FieldOtherconfig, field.TypeOther, value)
	}
	if tu.mutation.OtherconfigCleared() {
		_spec.ClearField(template.FieldOtherconfig, field.TypeOther)
	}
	if value, ok := tu.mutation.Pair(); ok {
		_spec.SetField(template.FieldPair, field.TypeBytes, value)
	}
	if value, ok := tu.mutation.URL(); ok {
		vv, err := template.ValueScanner.URL.Value(value)
		if err != nil {
			return 0, err
		}
		_spec.SetField(template.FieldURL, field.TypeString, vv)
	}
	if tu.mutation.OwnerCleared() {
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
		edge.Schema = tu.schemaConfig.Template
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		edge.Schema = tu.schemaConfig.Template
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Template
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TemplateUpdateOne is the builder for updating a single Template entity.
type TemplateUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TemplateMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TemplateUpdateOne) SetUpdatedAt(t time.Time) *TemplateUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TemplateUpdateOne) ClearUpdatedAt() *TemplateUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TemplateUpdateOne) SetUpdatedBy(s string) *TemplateUpdateOne {
	tuo.mutation.SetUpdatedBy(s)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableUpdatedBy(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetUpdatedBy(*s)
	}
	return tuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tuo *TemplateUpdateOne) ClearUpdatedBy() *TemplateUpdateOne {
	tuo.mutation.ClearUpdatedBy()
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TemplateUpdateOne) SetDeletedAt(t time.Time) *TemplateUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableDeletedAt(t *time.Time) *TemplateUpdateOne {
	if t != nil {
		tuo.SetDeletedAt(*t)
	}
	return tuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tuo *TemplateUpdateOne) ClearDeletedAt() *TemplateUpdateOne {
	tuo.mutation.ClearDeletedAt()
	return tuo
}

// SetDeletedBy sets the "deleted_by" field.
func (tuo *TemplateUpdateOne) SetDeletedBy(s string) *TemplateUpdateOne {
	tuo.mutation.SetDeletedBy(s)
	return tuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableDeletedBy(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetDeletedBy(*s)
	}
	return tuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tuo *TemplateUpdateOne) ClearDeletedBy() *TemplateUpdateOne {
	tuo.mutation.ClearDeletedBy()
	return tuo
}

// SetOwnerID sets the "owner_id" field.
func (tuo *TemplateUpdateOne) SetOwnerID(s string) *TemplateUpdateOne {
	tuo.mutation.SetOwnerID(s)
	return tuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableOwnerID(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetOwnerID(*s)
	}
	return tuo
}

// SetName sets the "name" field.
func (tuo *TemplateUpdateOne) SetName(s string) *TemplateUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableName(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TemplateUpdateOne) SetDescription(s string) *TemplateUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillableDescription(s *string) *TemplateUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TemplateUpdateOne) ClearDescription() *TemplateUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetJsonconfig sets the "jsonconfig" field.
func (tuo *TemplateUpdateOne) SetJsonconfig(co customtypes.JSONObject) *TemplateUpdateOne {
	tuo.mutation.SetJsonconfig(co)
	return tuo
}

// ClearJsonconfig clears the value of the "jsonconfig" field.
func (tuo *TemplateUpdateOne) ClearJsonconfig() *TemplateUpdateOne {
	tuo.mutation.ClearJsonconfig()
	return tuo
}

// SetOtherconfig sets the "otherconfig" field.
func (tuo *TemplateUpdateOne) SetOtherconfig(sc *schematype.TemplateConfig) *TemplateUpdateOne {
	tuo.mutation.SetOtherconfig(sc)
	return tuo
}

// ClearOtherconfig clears the value of the "otherconfig" field.
func (tuo *TemplateUpdateOne) ClearOtherconfig() *TemplateUpdateOne {
	tuo.mutation.ClearOtherconfig()
	return tuo
}

// SetPair sets the "pair" field.
func (tuo *TemplateUpdateOne) SetPair(c customtypes.Pair) *TemplateUpdateOne {
	tuo.mutation.SetPair(c)
	return tuo
}

// SetNillablePair sets the "pair" field if the given value is not nil.
func (tuo *TemplateUpdateOne) SetNillablePair(c *customtypes.Pair) *TemplateUpdateOne {
	if c != nil {
		tuo.SetPair(*c)
	}
	return tuo
}

// SetURL sets the "url" field.
func (tuo *TemplateUpdateOne) SetURL(u *url.URL) *TemplateUpdateOne {
	tuo.mutation.SetURL(u)
	return tuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (tuo *TemplateUpdateOne) SetOwner(o *Organization) *TemplateUpdateOne {
	return tuo.SetOwnerID(o.ID)
}

// Mutation returns the TemplateMutation object of the builder.
func (tuo *TemplateUpdateOne) Mutation() *TemplateMutation {
	return tuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (tuo *TemplateUpdateOne) ClearOwner() *TemplateUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// Where appends a list predicates to the TemplateUpdate builder.
func (tuo *TemplateUpdateOne) Where(ps ...predicate.Template) *TemplateUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TemplateUpdateOne) Select(field string, fields ...string) *TemplateUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Template entity.
func (tuo *TemplateUpdateOne) Save(ctx context.Context) (*Template, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TemplateUpdateOne) SaveX(ctx context.Context) *Template {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TemplateUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TemplateUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		if template.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized template.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := template.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TemplateUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := template.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Template.name": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.OwnerID(); tuo.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "Template.owner"`)
	}
	return nil
}

func (tuo *TemplateUpdateOne) sqlSave(ctx context.Context) (_node *Template, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(template.Table, template.Columns, sqlgraph.NewFieldSpec(template.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Template.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, template.FieldID)
		for _, f := range fields {
			if !template.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != template.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tuo.mutation.CreatedAtCleared() {
		_spec.ClearField(template.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(template.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(template.FieldUpdatedAt, field.TypeTime)
	}
	if tuo.mutation.CreatedByCleared() {
		_spec.ClearField(template.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(template.FieldUpdatedBy, field.TypeString, value)
	}
	if tuo.mutation.UpdatedByCleared() {
		_spec.ClearField(template.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(template.FieldDeletedAt, field.TypeTime, value)
	}
	if tuo.mutation.DeletedAtCleared() {
		_spec.ClearField(template.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.DeletedBy(); ok {
		_spec.SetField(template.FieldDeletedBy, field.TypeString, value)
	}
	if tuo.mutation.DeletedByCleared() {
		_spec.ClearField(template.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(template.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(template.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(template.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.Jsonconfig(); ok {
		_spec.SetField(template.FieldJsonconfig, field.TypeJSON, value)
	}
	if tuo.mutation.JsonconfigCleared() {
		_spec.ClearField(template.FieldJsonconfig, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Otherconfig(); ok {
		_spec.SetField(template.FieldOtherconfig, field.TypeOther, value)
	}
	if tuo.mutation.OtherconfigCleared() {
		_spec.ClearField(template.FieldOtherconfig, field.TypeOther)
	}
	if value, ok := tuo.mutation.Pair(); ok {
		_spec.SetField(template.FieldPair, field.TypeBytes, value)
	}
	if value, ok := tuo.mutation.URL(); ok {
		vv, err := template.ValueScanner.URL.Value(value)
		if err != nil {
			return nil, err
		}
		_spec.SetField(template.FieldURL, field.TypeString, vv)
	}
	if tuo.mutation.OwnerCleared() {
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
		edge.Schema = tuo.schemaConfig.Template
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		edge.Schema = tuo.schemaConfig.Template
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Template
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_node = &Template{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{template.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
