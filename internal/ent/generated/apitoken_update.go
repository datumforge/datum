// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/apitoken"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// APITokenUpdate is the builder for updating APIToken entities.
type APITokenUpdate struct {
	config
	hooks    []Hook
	mutation *APITokenMutation
}

// Where appends a list predicates to the APITokenUpdate builder.
func (atu *APITokenUpdate) Where(ps ...predicate.APIToken) *APITokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *APITokenUpdate) SetUpdatedAt(t time.Time) *APITokenUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atu *APITokenUpdate) ClearUpdatedAt() *APITokenUpdate {
	atu.mutation.ClearUpdatedAt()
	return atu
}

// SetUpdatedBy sets the "updated_by" field.
func (atu *APITokenUpdate) SetUpdatedBy(s string) *APITokenUpdate {
	atu.mutation.SetUpdatedBy(s)
	return atu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableUpdatedBy(s *string) *APITokenUpdate {
	if s != nil {
		atu.SetUpdatedBy(*s)
	}
	return atu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (atu *APITokenUpdate) ClearUpdatedBy() *APITokenUpdate {
	atu.mutation.ClearUpdatedBy()
	return atu
}

// SetDeletedAt sets the "deleted_at" field.
func (atu *APITokenUpdate) SetDeletedAt(t time.Time) *APITokenUpdate {
	atu.mutation.SetDeletedAt(t)
	return atu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableDeletedAt(t *time.Time) *APITokenUpdate {
	if t != nil {
		atu.SetDeletedAt(*t)
	}
	return atu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atu *APITokenUpdate) ClearDeletedAt() *APITokenUpdate {
	atu.mutation.ClearDeletedAt()
	return atu
}

// SetDeletedBy sets the "deleted_by" field.
func (atu *APITokenUpdate) SetDeletedBy(s string) *APITokenUpdate {
	atu.mutation.SetDeletedBy(s)
	return atu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableDeletedBy(s *string) *APITokenUpdate {
	if s != nil {
		atu.SetDeletedBy(*s)
	}
	return atu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (atu *APITokenUpdate) ClearDeletedBy() *APITokenUpdate {
	atu.mutation.ClearDeletedBy()
	return atu
}

// SetOwnerID sets the "owner_id" field.
func (atu *APITokenUpdate) SetOwnerID(s string) *APITokenUpdate {
	atu.mutation.SetOwnerID(s)
	return atu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableOwnerID(s *string) *APITokenUpdate {
	if s != nil {
		atu.SetOwnerID(*s)
	}
	return atu
}

// SetName sets the "name" field.
func (atu *APITokenUpdate) SetName(s string) *APITokenUpdate {
	atu.mutation.SetName(s)
	return atu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableName(s *string) *APITokenUpdate {
	if s != nil {
		atu.SetName(*s)
	}
	return atu
}

// SetExpiresAt sets the "expires_at" field.
func (atu *APITokenUpdate) SetExpiresAt(t time.Time) *APITokenUpdate {
	atu.mutation.SetExpiresAt(t)
	return atu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableExpiresAt(t *time.Time) *APITokenUpdate {
	if t != nil {
		atu.SetExpiresAt(*t)
	}
	return atu
}

// SetDescription sets the "description" field.
func (atu *APITokenUpdate) SetDescription(s string) *APITokenUpdate {
	atu.mutation.SetDescription(s)
	return atu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableDescription(s *string) *APITokenUpdate {
	if s != nil {
		atu.SetDescription(*s)
	}
	return atu
}

// ClearDescription clears the value of the "description" field.
func (atu *APITokenUpdate) ClearDescription() *APITokenUpdate {
	atu.mutation.ClearDescription()
	return atu
}

// SetScopes sets the "scopes" field.
func (atu *APITokenUpdate) SetScopes(s []string) *APITokenUpdate {
	atu.mutation.SetScopes(s)
	return atu
}

// AppendScopes appends s to the "scopes" field.
func (atu *APITokenUpdate) AppendScopes(s []string) *APITokenUpdate {
	atu.mutation.AppendScopes(s)
	return atu
}

// ClearScopes clears the value of the "scopes" field.
func (atu *APITokenUpdate) ClearScopes() *APITokenUpdate {
	atu.mutation.ClearScopes()
	return atu
}

// SetLastUsedAt sets the "last_used_at" field.
func (atu *APITokenUpdate) SetLastUsedAt(t time.Time) *APITokenUpdate {
	atu.mutation.SetLastUsedAt(t)
	return atu
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (atu *APITokenUpdate) SetNillableLastUsedAt(t *time.Time) *APITokenUpdate {
	if t != nil {
		atu.SetLastUsedAt(*t)
	}
	return atu
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (atu *APITokenUpdate) ClearLastUsedAt() *APITokenUpdate {
	atu.mutation.ClearLastUsedAt()
	return atu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (atu *APITokenUpdate) SetOwner(o *Organization) *APITokenUpdate {
	return atu.SetOwnerID(o.ID)
}

// Mutation returns the APITokenMutation object of the builder.
func (atu *APITokenUpdate) Mutation() *APITokenMutation {
	return atu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (atu *APITokenUpdate) ClearOwner() *APITokenUpdate {
	atu.mutation.ClearOwner()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *APITokenUpdate) Save(ctx context.Context) (int, error) {
	if err := atu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *APITokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *APITokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *APITokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *APITokenUpdate) defaults() error {
	if _, ok := atu.mutation.UpdatedAt(); !ok && !atu.mutation.UpdatedAtCleared() {
		if apitoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized apitoken.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := apitoken.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (atu *APITokenUpdate) check() error {
	if v, ok := atu.mutation.Name(); ok {
		if err := apitoken.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "APIToken.name": %w`, err)}
		}
	}
	if _, ok := atu.mutation.OwnerID(); atu.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "APIToken.owner"`)
	}
	return nil
}

func (atu *APITokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := atu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(apitoken.Table, apitoken.Columns, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atu.mutation.CreatedAtCleared() {
		_spec.ClearField(apitoken.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atu.mutation.UpdatedAtCleared() {
		_spec.ClearField(apitoken.FieldUpdatedAt, field.TypeTime)
	}
	if atu.mutation.CreatedByCleared() {
		_spec.ClearField(apitoken.FieldCreatedBy, field.TypeString)
	}
	if value, ok := atu.mutation.UpdatedBy(); ok {
		_spec.SetField(apitoken.FieldUpdatedBy, field.TypeString, value)
	}
	if atu.mutation.UpdatedByCleared() {
		_spec.ClearField(apitoken.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := atu.mutation.DeletedAt(); ok {
		_spec.SetField(apitoken.FieldDeletedAt, field.TypeTime, value)
	}
	if atu.mutation.DeletedAtCleared() {
		_spec.ClearField(apitoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.DeletedBy(); ok {
		_spec.SetField(apitoken.FieldDeletedBy, field.TypeString, value)
	}
	if atu.mutation.DeletedByCleared() {
		_spec.ClearField(apitoken.FieldDeletedBy, field.TypeString)
	}
	if value, ok := atu.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
	}
	if value, ok := atu.mutation.ExpiresAt(); ok {
		_spec.SetField(apitoken.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := atu.mutation.Description(); ok {
		_spec.SetField(apitoken.FieldDescription, field.TypeString, value)
	}
	if atu.mutation.DescriptionCleared() {
		_spec.ClearField(apitoken.FieldDescription, field.TypeString)
	}
	if value, ok := atu.mutation.Scopes(); ok {
		_spec.SetField(apitoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := atu.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, apitoken.FieldScopes, value)
		})
	}
	if atu.mutation.ScopesCleared() {
		_spec.ClearField(apitoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := atu.mutation.LastUsedAt(); ok {
		_spec.SetField(apitoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if atu.mutation.LastUsedAtCleared() {
		_spec.ClearField(apitoken.FieldLastUsedAt, field.TypeTime)
	}
	if atu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.OwnerTable,
			Columns: []string{apitoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = atu.schemaConfig.APIToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.OwnerTable,
			Columns: []string{apitoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = atu.schemaConfig.APIToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = atu.schemaConfig.APIToken
	ctx = internal.NewSchemaConfigContext(ctx, atu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// APITokenUpdateOne is the builder for updating a single APIToken entity.
type APITokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APITokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *APITokenUpdateOne) SetUpdatedAt(t time.Time) *APITokenUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atuo *APITokenUpdateOne) ClearUpdatedAt() *APITokenUpdateOne {
	atuo.mutation.ClearUpdatedAt()
	return atuo
}

// SetUpdatedBy sets the "updated_by" field.
func (atuo *APITokenUpdateOne) SetUpdatedBy(s string) *APITokenUpdateOne {
	atuo.mutation.SetUpdatedBy(s)
	return atuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableUpdatedBy(s *string) *APITokenUpdateOne {
	if s != nil {
		atuo.SetUpdatedBy(*s)
	}
	return atuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (atuo *APITokenUpdateOne) ClearUpdatedBy() *APITokenUpdateOne {
	atuo.mutation.ClearUpdatedBy()
	return atuo
}

// SetDeletedAt sets the "deleted_at" field.
func (atuo *APITokenUpdateOne) SetDeletedAt(t time.Time) *APITokenUpdateOne {
	atuo.mutation.SetDeletedAt(t)
	return atuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableDeletedAt(t *time.Time) *APITokenUpdateOne {
	if t != nil {
		atuo.SetDeletedAt(*t)
	}
	return atuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atuo *APITokenUpdateOne) ClearDeletedAt() *APITokenUpdateOne {
	atuo.mutation.ClearDeletedAt()
	return atuo
}

// SetDeletedBy sets the "deleted_by" field.
func (atuo *APITokenUpdateOne) SetDeletedBy(s string) *APITokenUpdateOne {
	atuo.mutation.SetDeletedBy(s)
	return atuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableDeletedBy(s *string) *APITokenUpdateOne {
	if s != nil {
		atuo.SetDeletedBy(*s)
	}
	return atuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (atuo *APITokenUpdateOne) ClearDeletedBy() *APITokenUpdateOne {
	atuo.mutation.ClearDeletedBy()
	return atuo
}

// SetOwnerID sets the "owner_id" field.
func (atuo *APITokenUpdateOne) SetOwnerID(s string) *APITokenUpdateOne {
	atuo.mutation.SetOwnerID(s)
	return atuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableOwnerID(s *string) *APITokenUpdateOne {
	if s != nil {
		atuo.SetOwnerID(*s)
	}
	return atuo
}

// SetName sets the "name" field.
func (atuo *APITokenUpdateOne) SetName(s string) *APITokenUpdateOne {
	atuo.mutation.SetName(s)
	return atuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableName(s *string) *APITokenUpdateOne {
	if s != nil {
		atuo.SetName(*s)
	}
	return atuo
}

// SetExpiresAt sets the "expires_at" field.
func (atuo *APITokenUpdateOne) SetExpiresAt(t time.Time) *APITokenUpdateOne {
	atuo.mutation.SetExpiresAt(t)
	return atuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableExpiresAt(t *time.Time) *APITokenUpdateOne {
	if t != nil {
		atuo.SetExpiresAt(*t)
	}
	return atuo
}

// SetDescription sets the "description" field.
func (atuo *APITokenUpdateOne) SetDescription(s string) *APITokenUpdateOne {
	atuo.mutation.SetDescription(s)
	return atuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableDescription(s *string) *APITokenUpdateOne {
	if s != nil {
		atuo.SetDescription(*s)
	}
	return atuo
}

// ClearDescription clears the value of the "description" field.
func (atuo *APITokenUpdateOne) ClearDescription() *APITokenUpdateOne {
	atuo.mutation.ClearDescription()
	return atuo
}

// SetScopes sets the "scopes" field.
func (atuo *APITokenUpdateOne) SetScopes(s []string) *APITokenUpdateOne {
	atuo.mutation.SetScopes(s)
	return atuo
}

// AppendScopes appends s to the "scopes" field.
func (atuo *APITokenUpdateOne) AppendScopes(s []string) *APITokenUpdateOne {
	atuo.mutation.AppendScopes(s)
	return atuo
}

// ClearScopes clears the value of the "scopes" field.
func (atuo *APITokenUpdateOne) ClearScopes() *APITokenUpdateOne {
	atuo.mutation.ClearScopes()
	return atuo
}

// SetLastUsedAt sets the "last_used_at" field.
func (atuo *APITokenUpdateOne) SetLastUsedAt(t time.Time) *APITokenUpdateOne {
	atuo.mutation.SetLastUsedAt(t)
	return atuo
}

// SetNillableLastUsedAt sets the "last_used_at" field if the given value is not nil.
func (atuo *APITokenUpdateOne) SetNillableLastUsedAt(t *time.Time) *APITokenUpdateOne {
	if t != nil {
		atuo.SetLastUsedAt(*t)
	}
	return atuo
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (atuo *APITokenUpdateOne) ClearLastUsedAt() *APITokenUpdateOne {
	atuo.mutation.ClearLastUsedAt()
	return atuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (atuo *APITokenUpdateOne) SetOwner(o *Organization) *APITokenUpdateOne {
	return atuo.SetOwnerID(o.ID)
}

// Mutation returns the APITokenMutation object of the builder.
func (atuo *APITokenUpdateOne) Mutation() *APITokenMutation {
	return atuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (atuo *APITokenUpdateOne) ClearOwner() *APITokenUpdateOne {
	atuo.mutation.ClearOwner()
	return atuo
}

// Where appends a list predicates to the APITokenUpdate builder.
func (atuo *APITokenUpdateOne) Where(ps ...predicate.APIToken) *APITokenUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *APITokenUpdateOne) Select(field string, fields ...string) *APITokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated APIToken entity.
func (atuo *APITokenUpdateOne) Save(ctx context.Context) (*APIToken, error) {
	if err := atuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *APITokenUpdateOne) SaveX(ctx context.Context) *APIToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *APITokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *APITokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *APITokenUpdateOne) defaults() error {
	if _, ok := atuo.mutation.UpdatedAt(); !ok && !atuo.mutation.UpdatedAtCleared() {
		if apitoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized apitoken.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := apitoken.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (atuo *APITokenUpdateOne) check() error {
	if v, ok := atuo.mutation.Name(); ok {
		if err := apitoken.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "APIToken.name": %w`, err)}
		}
	}
	if _, ok := atuo.mutation.OwnerID(); atuo.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "APIToken.owner"`)
	}
	return nil
}

func (atuo *APITokenUpdateOne) sqlSave(ctx context.Context) (_node *APIToken, err error) {
	if err := atuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(apitoken.Table, apitoken.Columns, sqlgraph.NewFieldSpec(apitoken.FieldID, field.TypeString))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "APIToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, apitoken.FieldID)
		for _, f := range fields {
			if !apitoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != apitoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if atuo.mutation.CreatedAtCleared() {
		_spec.ClearField(apitoken.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(apitoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(apitoken.FieldUpdatedAt, field.TypeTime)
	}
	if atuo.mutation.CreatedByCleared() {
		_spec.ClearField(apitoken.FieldCreatedBy, field.TypeString)
	}
	if value, ok := atuo.mutation.UpdatedBy(); ok {
		_spec.SetField(apitoken.FieldUpdatedBy, field.TypeString, value)
	}
	if atuo.mutation.UpdatedByCleared() {
		_spec.ClearField(apitoken.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := atuo.mutation.DeletedAt(); ok {
		_spec.SetField(apitoken.FieldDeletedAt, field.TypeTime, value)
	}
	if atuo.mutation.DeletedAtCleared() {
		_spec.ClearField(apitoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.DeletedBy(); ok {
		_spec.SetField(apitoken.FieldDeletedBy, field.TypeString, value)
	}
	if atuo.mutation.DeletedByCleared() {
		_spec.ClearField(apitoken.FieldDeletedBy, field.TypeString)
	}
	if value, ok := atuo.mutation.Name(); ok {
		_spec.SetField(apitoken.FieldName, field.TypeString, value)
	}
	if value, ok := atuo.mutation.ExpiresAt(); ok {
		_spec.SetField(apitoken.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := atuo.mutation.Description(); ok {
		_spec.SetField(apitoken.FieldDescription, field.TypeString, value)
	}
	if atuo.mutation.DescriptionCleared() {
		_spec.ClearField(apitoken.FieldDescription, field.TypeString)
	}
	if value, ok := atuo.mutation.Scopes(); ok {
		_spec.SetField(apitoken.FieldScopes, field.TypeJSON, value)
	}
	if value, ok := atuo.mutation.AppendedScopes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, apitoken.FieldScopes, value)
		})
	}
	if atuo.mutation.ScopesCleared() {
		_spec.ClearField(apitoken.FieldScopes, field.TypeJSON)
	}
	if value, ok := atuo.mutation.LastUsedAt(); ok {
		_spec.SetField(apitoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if atuo.mutation.LastUsedAtCleared() {
		_spec.ClearField(apitoken.FieldLastUsedAt, field.TypeTime)
	}
	if atuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.OwnerTable,
			Columns: []string{apitoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = atuo.schemaConfig.APIToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   apitoken.OwnerTable,
			Columns: []string{apitoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = atuo.schemaConfig.APIToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = atuo.schemaConfig.APIToken
	ctx = internal.NewSchemaConfigContext(ctx, atuo.schemaConfig)
	_node = &APIToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{apitoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
