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
	"github.com/datumforge/datum/internal/ent/generated/integration"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// IntegrationUpdate is the builder for updating Integration entities.
type IntegrationUpdate struct {
	config
	hooks    []Hook
	mutation *IntegrationMutation
}

// Where appends a list predicates to the IntegrationUpdate builder.
func (iu *IntegrationUpdate) Where(ps ...predicate.Integration) *IntegrationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IntegrationUpdate) SetUpdatedAt(t time.Time) *IntegrationUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iu *IntegrationUpdate) ClearUpdatedAt() *IntegrationUpdate {
	iu.mutation.ClearUpdatedAt()
	return iu
}

// SetUpdatedBy sets the "updated_by" field.
func (iu *IntegrationUpdate) SetUpdatedBy(s string) *IntegrationUpdate {
	iu.mutation.SetUpdatedBy(s)
	return iu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableUpdatedBy(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetUpdatedBy(*s)
	}
	return iu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iu *IntegrationUpdate) ClearUpdatedBy() *IntegrationUpdate {
	iu.mutation.ClearUpdatedBy()
	return iu
}

// SetDeletedAt sets the "deleted_at" field.
func (iu *IntegrationUpdate) SetDeletedAt(t time.Time) *IntegrationUpdate {
	iu.mutation.SetDeletedAt(t)
	return iu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDeletedAt(t *time.Time) *IntegrationUpdate {
	if t != nil {
		iu.SetDeletedAt(*t)
	}
	return iu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iu *IntegrationUpdate) ClearDeletedAt() *IntegrationUpdate {
	iu.mutation.ClearDeletedAt()
	return iu
}

// SetDeletedBy sets the "deleted_by" field.
func (iu *IntegrationUpdate) SetDeletedBy(s string) *IntegrationUpdate {
	iu.mutation.SetDeletedBy(s)
	return iu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDeletedBy(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetDeletedBy(*s)
	}
	return iu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iu *IntegrationUpdate) ClearDeletedBy() *IntegrationUpdate {
	iu.mutation.ClearDeletedBy()
	return iu
}

// SetName sets the "name" field.
func (iu *IntegrationUpdate) SetName(s string) *IntegrationUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableName(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetName(*s)
	}
	return iu
}

// SetDescription sets the "description" field.
func (iu *IntegrationUpdate) SetDescription(s string) *IntegrationUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableDescription(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetDescription(*s)
	}
	return iu
}

// ClearDescription clears the value of the "description" field.
func (iu *IntegrationUpdate) ClearDescription() *IntegrationUpdate {
	iu.mutation.ClearDescription()
	return iu
}

// SetKind sets the "kind" field.
func (iu *IntegrationUpdate) SetKind(s string) *IntegrationUpdate {
	iu.mutation.SetKind(s)
	return iu
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableKind(s *string) *IntegrationUpdate {
	if s != nil {
		iu.SetKind(*s)
	}
	return iu
}

// ClearKind clears the value of the "kind" field.
func (iu *IntegrationUpdate) ClearKind() *IntegrationUpdate {
	iu.mutation.ClearKind()
	return iu
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (iu *IntegrationUpdate) SetOwnerID(id string) *IntegrationUpdate {
	iu.mutation.SetOwnerID(id)
	return iu
}

// SetNillableOwnerID sets the "owner" edge to the Organization entity by ID if the given value is not nil.
func (iu *IntegrationUpdate) SetNillableOwnerID(id *string) *IntegrationUpdate {
	if id != nil {
		iu = iu.SetOwnerID(*id)
	}
	return iu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (iu *IntegrationUpdate) SetOwner(o *Organization) *IntegrationUpdate {
	return iu.SetOwnerID(o.ID)
}

// Mutation returns the IntegrationMutation object of the builder.
func (iu *IntegrationUpdate) Mutation() *IntegrationMutation {
	return iu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (iu *IntegrationUpdate) ClearOwner() *IntegrationUpdate {
	iu.mutation.ClearOwner()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IntegrationUpdate) Save(ctx context.Context) (int, error) {
	if err := iu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IntegrationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IntegrationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IntegrationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *IntegrationUpdate) defaults() error {
	if _, ok := iu.mutation.UpdatedAt(); !ok && !iu.mutation.UpdatedAtCleared() {
		if integration.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := integration.UpdateDefaultUpdatedAt()
		iu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (iu *IntegrationUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := integration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Integration.name": %w`, err)}
		}
	}
	return nil
}

func (iu *IntegrationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(integration.Table, integration.Columns, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iu.mutation.CreatedAtCleared() {
		_spec.ClearField(integration.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.SetField(integration.FieldUpdatedAt, field.TypeTime, value)
	}
	if iu.mutation.UpdatedAtCleared() {
		_spec.ClearField(integration.FieldUpdatedAt, field.TypeTime)
	}
	if iu.mutation.CreatedByCleared() {
		_spec.ClearField(integration.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.UpdatedBy(); ok {
		_spec.SetField(integration.FieldUpdatedBy, field.TypeString, value)
	}
	if iu.mutation.UpdatedByCleared() {
		_spec.ClearField(integration.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iu.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
	}
	if iu.mutation.DeletedAtCleared() {
		_spec.ClearField(integration.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iu.mutation.DeletedBy(); ok {
		_spec.SetField(integration.FieldDeletedBy, field.TypeString, value)
	}
	if iu.mutation.DeletedByCleared() {
		_spec.ClearField(integration.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
	}
	if iu.mutation.DescriptionCleared() {
		_spec.ClearField(integration.FieldDescription, field.TypeString)
	}
	if value, ok := iu.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
	}
	if iu.mutation.KindCleared() {
		_spec.ClearField(integration.FieldKind, field.TypeString)
	}
	if iu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.Integration
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iu.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = iu.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// IntegrationUpdateOne is the builder for updating a single Integration entity.
type IntegrationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IntegrationMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IntegrationUpdateOne) SetUpdatedAt(t time.Time) *IntegrationUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (iuo *IntegrationUpdateOne) ClearUpdatedAt() *IntegrationUpdateOne {
	iuo.mutation.ClearUpdatedAt()
	return iuo
}

// SetUpdatedBy sets the "updated_by" field.
func (iuo *IntegrationUpdateOne) SetUpdatedBy(s string) *IntegrationUpdateOne {
	iuo.mutation.SetUpdatedBy(s)
	return iuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableUpdatedBy(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetUpdatedBy(*s)
	}
	return iuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (iuo *IntegrationUpdateOne) ClearUpdatedBy() *IntegrationUpdateOne {
	iuo.mutation.ClearUpdatedBy()
	return iuo
}

// SetDeletedAt sets the "deleted_at" field.
func (iuo *IntegrationUpdateOne) SetDeletedAt(t time.Time) *IntegrationUpdateOne {
	iuo.mutation.SetDeletedAt(t)
	return iuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDeletedAt(t *time.Time) *IntegrationUpdateOne {
	if t != nil {
		iuo.SetDeletedAt(*t)
	}
	return iuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (iuo *IntegrationUpdateOne) ClearDeletedAt() *IntegrationUpdateOne {
	iuo.mutation.ClearDeletedAt()
	return iuo
}

// SetDeletedBy sets the "deleted_by" field.
func (iuo *IntegrationUpdateOne) SetDeletedBy(s string) *IntegrationUpdateOne {
	iuo.mutation.SetDeletedBy(s)
	return iuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDeletedBy(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetDeletedBy(*s)
	}
	return iuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (iuo *IntegrationUpdateOne) ClearDeletedBy() *IntegrationUpdateOne {
	iuo.mutation.ClearDeletedBy()
	return iuo
}

// SetName sets the "name" field.
func (iuo *IntegrationUpdateOne) SetName(s string) *IntegrationUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableName(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetName(*s)
	}
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IntegrationUpdateOne) SetDescription(s string) *IntegrationUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableDescription(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetDescription(*s)
	}
	return iuo
}

// ClearDescription clears the value of the "description" field.
func (iuo *IntegrationUpdateOne) ClearDescription() *IntegrationUpdateOne {
	iuo.mutation.ClearDescription()
	return iuo
}

// SetKind sets the "kind" field.
func (iuo *IntegrationUpdateOne) SetKind(s string) *IntegrationUpdateOne {
	iuo.mutation.SetKind(s)
	return iuo
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableKind(s *string) *IntegrationUpdateOne {
	if s != nil {
		iuo.SetKind(*s)
	}
	return iuo
}

// ClearKind clears the value of the "kind" field.
func (iuo *IntegrationUpdateOne) ClearKind() *IntegrationUpdateOne {
	iuo.mutation.ClearKind()
	return iuo
}

// SetOwnerID sets the "owner" edge to the Organization entity by ID.
func (iuo *IntegrationUpdateOne) SetOwnerID(id string) *IntegrationUpdateOne {
	iuo.mutation.SetOwnerID(id)
	return iuo
}

// SetNillableOwnerID sets the "owner" edge to the Organization entity by ID if the given value is not nil.
func (iuo *IntegrationUpdateOne) SetNillableOwnerID(id *string) *IntegrationUpdateOne {
	if id != nil {
		iuo = iuo.SetOwnerID(*id)
	}
	return iuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (iuo *IntegrationUpdateOne) SetOwner(o *Organization) *IntegrationUpdateOne {
	return iuo.SetOwnerID(o.ID)
}

// Mutation returns the IntegrationMutation object of the builder.
func (iuo *IntegrationUpdateOne) Mutation() *IntegrationMutation {
	return iuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (iuo *IntegrationUpdateOne) ClearOwner() *IntegrationUpdateOne {
	iuo.mutation.ClearOwner()
	return iuo
}

// Where appends a list predicates to the IntegrationUpdate builder.
func (iuo *IntegrationUpdateOne) Where(ps ...predicate.Integration) *IntegrationUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IntegrationUpdateOne) Select(field string, fields ...string) *IntegrationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Integration entity.
func (iuo *IntegrationUpdateOne) Save(ctx context.Context) (*Integration, error) {
	if err := iuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IntegrationUpdateOne) SaveX(ctx context.Context) *Integration {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IntegrationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IntegrationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *IntegrationUpdateOne) defaults() error {
	if _, ok := iuo.mutation.UpdatedAt(); !ok && !iuo.mutation.UpdatedAtCleared() {
		if integration.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized integration.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := integration.UpdateDefaultUpdatedAt()
		iuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IntegrationUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := integration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Integration.name": %w`, err)}
		}
	}
	return nil
}

func (iuo *IntegrationUpdateOne) sqlSave(ctx context.Context) (_node *Integration, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(integration.Table, integration.Columns, sqlgraph.NewFieldSpec(integration.FieldID, field.TypeString))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Integration.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, integration.FieldID)
		for _, f := range fields {
			if !integration.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != integration.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if iuo.mutation.CreatedAtCleared() {
		_spec.ClearField(integration.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.SetField(integration.FieldUpdatedAt, field.TypeTime, value)
	}
	if iuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(integration.FieldUpdatedAt, field.TypeTime)
	}
	if iuo.mutation.CreatedByCleared() {
		_spec.ClearField(integration.FieldCreatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.UpdatedBy(); ok {
		_spec.SetField(integration.FieldUpdatedBy, field.TypeString, value)
	}
	if iuo.mutation.UpdatedByCleared() {
		_spec.ClearField(integration.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.DeletedAt(); ok {
		_spec.SetField(integration.FieldDeletedAt, field.TypeTime, value)
	}
	if iuo.mutation.DeletedAtCleared() {
		_spec.ClearField(integration.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := iuo.mutation.DeletedBy(); ok {
		_spec.SetField(integration.FieldDeletedBy, field.TypeString, value)
	}
	if iuo.mutation.DeletedByCleared() {
		_spec.ClearField(integration.FieldDeletedBy, field.TypeString)
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(integration.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.SetField(integration.FieldDescription, field.TypeString, value)
	}
	if iuo.mutation.DescriptionCleared() {
		_spec.ClearField(integration.FieldDescription, field.TypeString)
	}
	if value, ok := iuo.mutation.Kind(); ok {
		_spec.SetField(integration.FieldKind, field.TypeString, value)
	}
	if iuo.mutation.KindCleared() {
		_spec.ClearField(integration.FieldKind, field.TypeString)
	}
	if iuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.Integration
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   integration.OwnerTable,
			Columns: []string{integration.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = iuo.schemaConfig.Integration
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = iuo.schemaConfig.Integration
	ctx = internal.NewSchemaConfigContext(ctx, iuo.schemaConfig)
	_node = &Integration{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{integration.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
