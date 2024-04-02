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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/tier"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TierUpdate is the builder for updating Tier entities.
type TierUpdate struct {
	config
	hooks    []Hook
	mutation *TierMutation
}

// Where appends a list predicates to the TierUpdate builder.
func (tu *TierUpdate) Where(ps ...predicate.Tier) *TierUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TierUpdate) SetUpdatedAt(t time.Time) *TierUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tu *TierUpdate) ClearUpdatedAt() *TierUpdate {
	tu.mutation.ClearUpdatedAt()
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TierUpdate) SetUpdatedBy(s string) *TierUpdate {
	tu.mutation.SetUpdatedBy(s)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TierUpdate) SetNillableUpdatedBy(s *string) *TierUpdate {
	if s != nil {
		tu.SetUpdatedBy(*s)
	}
	return tu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tu *TierUpdate) ClearUpdatedBy() *TierUpdate {
	tu.mutation.ClearUpdatedBy()
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TierUpdate) SetDeletedAt(t time.Time) *TierUpdate {
	tu.mutation.SetDeletedAt(t)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TierUpdate) SetNillableDeletedAt(t *time.Time) *TierUpdate {
	if t != nil {
		tu.SetDeletedAt(*t)
	}
	return tu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tu *TierUpdate) ClearDeletedAt() *TierUpdate {
	tu.mutation.ClearDeletedAt()
	return tu
}

// SetDeletedBy sets the "deleted_by" field.
func (tu *TierUpdate) SetDeletedBy(s string) *TierUpdate {
	tu.mutation.SetDeletedBy(s)
	return tu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tu *TierUpdate) SetNillableDeletedBy(s *string) *TierUpdate {
	if s != nil {
		tu.SetDeletedBy(*s)
	}
	return tu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tu *TierUpdate) ClearDeletedBy() *TierUpdate {
	tu.mutation.ClearDeletedBy()
	return tu
}

// SetOwnerID sets the "owner_id" field.
func (tu *TierUpdate) SetOwnerID(s string) *TierUpdate {
	tu.mutation.SetOwnerID(s)
	return tu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tu *TierUpdate) SetNillableOwnerID(s *string) *TierUpdate {
	if s != nil {
		tu.SetOwnerID(*s)
	}
	return tu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (tu *TierUpdate) ClearOwnerID() *TierUpdate {
	tu.mutation.ClearOwnerID()
	return tu
}

// SetName sets the "name" field.
func (tu *TierUpdate) SetName(s string) *TierUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TierUpdate) SetNillableName(s *string) *TierUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TierUpdate) SetDescription(s string) *TierUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TierUpdate) SetNillableDescription(s *string) *TierUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of the "description" field.
func (tu *TierUpdate) ClearDescription() *TierUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetOrganizationID sets the "organization_id" field.
func (tu *TierUpdate) SetOrganizationID(s string) *TierUpdate {
	tu.mutation.SetOrganizationID(s)
	return tu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (tu *TierUpdate) SetNillableOrganizationID(s *string) *TierUpdate {
	if s != nil {
		tu.SetOrganizationID(*s)
	}
	return tu
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (tu *TierUpdate) ClearOrganizationID() *TierUpdate {
	tu.mutation.ClearOrganizationID()
	return tu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (tu *TierUpdate) SetOwner(o *Organization) *TierUpdate {
	return tu.SetOwnerID(o.ID)
}

// Mutation returns the TierMutation object of the builder.
func (tu *TierUpdate) Mutation() *TierMutation {
	return tu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (tu *TierUpdate) ClearOwner() *TierUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TierUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TierUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TierUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TierUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TierUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok && !tu.mutation.UpdatedAtCleared() {
		if tier.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized tier.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := tier.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TierUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tier.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Tier.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TierUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tier.Table, tier.Columns, sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.CreatedAtCleared() {
		_spec.ClearField(tier.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tier.FieldUpdatedAt, field.TypeTime, value)
	}
	if tu.mutation.UpdatedAtCleared() {
		_spec.ClearField(tier.FieldUpdatedAt, field.TypeTime)
	}
	if tu.mutation.CreatedByCleared() {
		_spec.ClearField(tier.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(tier.FieldUpdatedBy, field.TypeString, value)
	}
	if tu.mutation.UpdatedByCleared() {
		_spec.ClearField(tier.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(tier.FieldDeletedAt, field.TypeTime, value)
	}
	if tu.mutation.DeletedAtCleared() {
		_spec.ClearField(tier.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.DeletedBy(); ok {
		_spec.SetField(tier.FieldDeletedBy, field.TypeString, value)
	}
	if tu.mutation.DeletedByCleared() {
		_spec.ClearField(tier.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tier.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tier.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.DescriptionCleared() {
		_spec.ClearField(tier.FieldDescription, field.TypeString)
	}
	if value, ok := tu.mutation.OrganizationID(); ok {
		_spec.SetField(tier.FieldOrganizationID, field.TypeString, value)
	}
	if tu.mutation.OrganizationIDCleared() {
		_spec.ClearField(tier.FieldOrganizationID, field.TypeString)
	}
	if tu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tier.OwnerTable,
			Columns: []string{tier.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Tier
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tier.OwnerTable,
			Columns: []string{tier.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Tier
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TierUpdateOne is the builder for updating a single Tier entity.
type TierUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TierMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TierUpdateOne) SetUpdatedAt(t time.Time) *TierUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (tuo *TierUpdateOne) ClearUpdatedAt() *TierUpdateOne {
	tuo.mutation.ClearUpdatedAt()
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TierUpdateOne) SetUpdatedBy(s string) *TierUpdateOne {
	tuo.mutation.SetUpdatedBy(s)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableUpdatedBy(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetUpdatedBy(*s)
	}
	return tuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (tuo *TierUpdateOne) ClearUpdatedBy() *TierUpdateOne {
	tuo.mutation.ClearUpdatedBy()
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TierUpdateOne) SetDeletedAt(t time.Time) *TierUpdateOne {
	tuo.mutation.SetDeletedAt(t)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableDeletedAt(t *time.Time) *TierUpdateOne {
	if t != nil {
		tuo.SetDeletedAt(*t)
	}
	return tuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (tuo *TierUpdateOne) ClearDeletedAt() *TierUpdateOne {
	tuo.mutation.ClearDeletedAt()
	return tuo
}

// SetDeletedBy sets the "deleted_by" field.
func (tuo *TierUpdateOne) SetDeletedBy(s string) *TierUpdateOne {
	tuo.mutation.SetDeletedBy(s)
	return tuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableDeletedBy(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetDeletedBy(*s)
	}
	return tuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (tuo *TierUpdateOne) ClearDeletedBy() *TierUpdateOne {
	tuo.mutation.ClearDeletedBy()
	return tuo
}

// SetOwnerID sets the "owner_id" field.
func (tuo *TierUpdateOne) SetOwnerID(s string) *TierUpdateOne {
	tuo.mutation.SetOwnerID(s)
	return tuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableOwnerID(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetOwnerID(*s)
	}
	return tuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (tuo *TierUpdateOne) ClearOwnerID() *TierUpdateOne {
	tuo.mutation.ClearOwnerID()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TierUpdateOne) SetName(s string) *TierUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableName(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TierUpdateOne) SetDescription(s string) *TierUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableDescription(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of the "description" field.
func (tuo *TierUpdateOne) ClearDescription() *TierUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetOrganizationID sets the "organization_id" field.
func (tuo *TierUpdateOne) SetOrganizationID(s string) *TierUpdateOne {
	tuo.mutation.SetOrganizationID(s)
	return tuo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (tuo *TierUpdateOne) SetNillableOrganizationID(s *string) *TierUpdateOne {
	if s != nil {
		tuo.SetOrganizationID(*s)
	}
	return tuo
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (tuo *TierUpdateOne) ClearOrganizationID() *TierUpdateOne {
	tuo.mutation.ClearOrganizationID()
	return tuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (tuo *TierUpdateOne) SetOwner(o *Organization) *TierUpdateOne {
	return tuo.SetOwnerID(o.ID)
}

// Mutation returns the TierMutation object of the builder.
func (tuo *TierUpdateOne) Mutation() *TierMutation {
	return tuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (tuo *TierUpdateOne) ClearOwner() *TierUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// Where appends a list predicates to the TierUpdate builder.
func (tuo *TierUpdateOne) Where(ps ...predicate.Tier) *TierUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TierUpdateOne) Select(field string, fields ...string) *TierUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tier entity.
func (tuo *TierUpdateOne) Save(ctx context.Context) (*Tier, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TierUpdateOne) SaveX(ctx context.Context) *Tier {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TierUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TierUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TierUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok && !tuo.mutation.UpdatedAtCleared() {
		if tier.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized tier.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := tier.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TierUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tier.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Tier.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TierUpdateOne) sqlSave(ctx context.Context) (_node *Tier, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tier.Table, tier.Columns, sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Tier.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tier.FieldID)
		for _, f := range fields {
			if !tier.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != tier.FieldID {
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
		_spec.ClearField(tier.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tier.FieldUpdatedAt, field.TypeTime, value)
	}
	if tuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(tier.FieldUpdatedAt, field.TypeTime)
	}
	if tuo.mutation.CreatedByCleared() {
		_spec.ClearField(tier.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(tier.FieldUpdatedBy, field.TypeString, value)
	}
	if tuo.mutation.UpdatedByCleared() {
		_spec.ClearField(tier.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(tier.FieldDeletedAt, field.TypeTime, value)
	}
	if tuo.mutation.DeletedAtCleared() {
		_spec.ClearField(tier.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.DeletedBy(); ok {
		_spec.SetField(tier.FieldDeletedBy, field.TypeString, value)
	}
	if tuo.mutation.DeletedByCleared() {
		_spec.ClearField(tier.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tier.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tier.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.ClearField(tier.FieldDescription, field.TypeString)
	}
	if value, ok := tuo.mutation.OrganizationID(); ok {
		_spec.SetField(tier.FieldOrganizationID, field.TypeString, value)
	}
	if tuo.mutation.OrganizationIDCleared() {
		_spec.ClearField(tier.FieldOrganizationID, field.TypeString)
	}
	if tuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tier.OwnerTable,
			Columns: []string{tier.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Tier
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tier.OwnerTable,
			Columns: []string{tier.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Tier
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Tier
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_node = &Tier{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tier.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
