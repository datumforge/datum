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
	"github.com/datumforge/datum/internal/ent/generated/personalaccesstoken"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// PersonalAccessTokenUpdate is the builder for updating PersonalAccessToken entities.
type PersonalAccessTokenUpdate struct {
	config
	hooks    []Hook
	mutation *PersonalAccessTokenMutation
}

// Where appends a list predicates to the PersonalAccessTokenUpdate builder.
func (patu *PersonalAccessTokenUpdate) Where(ps ...predicate.PersonalAccessToken) *PersonalAccessTokenUpdate {
	patu.mutation.Where(ps...)
	return patu
}

// SetUpdatedAt sets the "updated_at" field.
func (patu *PersonalAccessTokenUpdate) SetUpdatedAt(t time.Time) *PersonalAccessTokenUpdate {
	patu.mutation.SetUpdatedAt(t)
	return patu
}

// SetUpdatedBy sets the "updated_by" field.
func (patu *PersonalAccessTokenUpdate) SetUpdatedBy(s string) *PersonalAccessTokenUpdate {
	patu.mutation.SetUpdatedBy(s)
	return patu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableUpdatedBy(s *string) *PersonalAccessTokenUpdate {
	if s != nil {
		patu.SetUpdatedBy(*s)
	}
	return patu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (patu *PersonalAccessTokenUpdate) ClearUpdatedBy() *PersonalAccessTokenUpdate {
	patu.mutation.ClearUpdatedBy()
	return patu
}

// SetDeletedAt sets the "deleted_at" field.
func (patu *PersonalAccessTokenUpdate) SetDeletedAt(t time.Time) *PersonalAccessTokenUpdate {
	patu.mutation.SetDeletedAt(t)
	return patu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableDeletedAt(t *time.Time) *PersonalAccessTokenUpdate {
	if t != nil {
		patu.SetDeletedAt(*t)
	}
	return patu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (patu *PersonalAccessTokenUpdate) ClearDeletedAt() *PersonalAccessTokenUpdate {
	patu.mutation.ClearDeletedAt()
	return patu
}

// SetDeletedBy sets the "deleted_by" field.
func (patu *PersonalAccessTokenUpdate) SetDeletedBy(s string) *PersonalAccessTokenUpdate {
	patu.mutation.SetDeletedBy(s)
	return patu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableDeletedBy(s *string) *PersonalAccessTokenUpdate {
	if s != nil {
		patu.SetDeletedBy(*s)
	}
	return patu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (patu *PersonalAccessTokenUpdate) ClearDeletedBy() *PersonalAccessTokenUpdate {
	patu.mutation.ClearDeletedBy()
	return patu
}

// SetName sets the "name" field.
func (patu *PersonalAccessTokenUpdate) SetName(s string) *PersonalAccessTokenUpdate {
	patu.mutation.SetName(s)
	return patu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableName(s *string) *PersonalAccessTokenUpdate {
	if s != nil {
		patu.SetName(*s)
	}
	return patu
}

// SetAbilities sets the "abilities" field.
func (patu *PersonalAccessTokenUpdate) SetAbilities(s []string) *PersonalAccessTokenUpdate {
	patu.mutation.SetAbilities(s)
	return patu
}

// AppendAbilities appends s to the "abilities" field.
func (patu *PersonalAccessTokenUpdate) AppendAbilities(s []string) *PersonalAccessTokenUpdate {
	patu.mutation.AppendAbilities(s)
	return patu
}

// ClearAbilities clears the value of the "abilities" field.
func (patu *PersonalAccessTokenUpdate) ClearAbilities() *PersonalAccessTokenUpdate {
	patu.mutation.ClearAbilities()
	return patu
}

// SetExpiresAt sets the "expires_at" field.
func (patu *PersonalAccessTokenUpdate) SetExpiresAt(t time.Time) *PersonalAccessTokenUpdate {
	patu.mutation.SetExpiresAt(t)
	return patu
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableExpiresAt(t *time.Time) *PersonalAccessTokenUpdate {
	if t != nil {
		patu.SetExpiresAt(*t)
	}
	return patu
}

// SetDescription sets the "description" field.
func (patu *PersonalAccessTokenUpdate) SetDescription(s string) *PersonalAccessTokenUpdate {
	patu.mutation.SetDescription(s)
	return patu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (patu *PersonalAccessTokenUpdate) SetNillableDescription(s *string) *PersonalAccessTokenUpdate {
	if s != nil {
		patu.SetDescription(*s)
	}
	return patu
}

// ClearDescription clears the value of the "description" field.
func (patu *PersonalAccessTokenUpdate) ClearDescription() *PersonalAccessTokenUpdate {
	patu.mutation.ClearDescription()
	return patu
}

// SetLastUsedAt sets the "last_used_at" field.
func (patu *PersonalAccessTokenUpdate) SetLastUsedAt(t time.Time) *PersonalAccessTokenUpdate {
	patu.mutation.SetLastUsedAt(t)
	return patu
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (patu *PersonalAccessTokenUpdate) ClearLastUsedAt() *PersonalAccessTokenUpdate {
	patu.mutation.ClearLastUsedAt()
	return patu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (patu *PersonalAccessTokenUpdate) SetOwnerID(id string) *PersonalAccessTokenUpdate {
	patu.mutation.SetOwnerID(id)
	return patu
}

// SetOwner sets the "owner" edge to the User entity.
func (patu *PersonalAccessTokenUpdate) SetOwner(u *User) *PersonalAccessTokenUpdate {
	return patu.SetOwnerID(u.ID)
}

// Mutation returns the PersonalAccessTokenMutation object of the builder.
func (patu *PersonalAccessTokenUpdate) Mutation() *PersonalAccessTokenMutation {
	return patu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (patu *PersonalAccessTokenUpdate) ClearOwner() *PersonalAccessTokenUpdate {
	patu.mutation.ClearOwner()
	return patu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (patu *PersonalAccessTokenUpdate) Save(ctx context.Context) (int, error) {
	if err := patu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, patu.sqlSave, patu.mutation, patu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (patu *PersonalAccessTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := patu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (patu *PersonalAccessTokenUpdate) Exec(ctx context.Context) error {
	_, err := patu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patu *PersonalAccessTokenUpdate) ExecX(ctx context.Context) {
	if err := patu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (patu *PersonalAccessTokenUpdate) defaults() error {
	if _, ok := patu.mutation.UpdatedAt(); !ok {
		if personalaccesstoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized personalaccesstoken.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := personalaccesstoken.UpdateDefaultUpdatedAt()
		patu.mutation.SetUpdatedAt(v)
	}
	if _, ok := patu.mutation.LastUsedAt(); !ok && !patu.mutation.LastUsedAtCleared() {
		if personalaccesstoken.UpdateDefaultLastUsedAt == nil {
			return fmt.Errorf("generated: uninitialized personalaccesstoken.UpdateDefaultLastUsedAt (forgotten import generated/runtime?)")
		}
		v := personalaccesstoken.UpdateDefaultLastUsedAt()
		patu.mutation.SetLastUsedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (patu *PersonalAccessTokenUpdate) check() error {
	if _, ok := patu.mutation.OwnerID(); patu.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "PersonalAccessToken.owner"`)
	}
	return nil
}

func (patu *PersonalAccessTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := patu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccesstoken.Table, personalaccesstoken.Columns, sqlgraph.NewFieldSpec(personalaccesstoken.FieldID, field.TypeString))
	if ps := patu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := patu.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if patu.mutation.CreatedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldCreatedBy, field.TypeString)
	}
	if value, ok := patu.mutation.UpdatedBy(); ok {
		_spec.SetField(personalaccesstoken.FieldUpdatedBy, field.TypeString, value)
	}
	if patu.mutation.UpdatedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := patu.mutation.DeletedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldDeletedAt, field.TypeTime, value)
	}
	if patu.mutation.DeletedAtCleared() {
		_spec.ClearField(personalaccesstoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := patu.mutation.DeletedBy(); ok {
		_spec.SetField(personalaccesstoken.FieldDeletedBy, field.TypeString, value)
	}
	if patu.mutation.DeletedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldDeletedBy, field.TypeString)
	}
	if value, ok := patu.mutation.Name(); ok {
		_spec.SetField(personalaccesstoken.FieldName, field.TypeString, value)
	}
	if value, ok := patu.mutation.Abilities(); ok {
		_spec.SetField(personalaccesstoken.FieldAbilities, field.TypeJSON, value)
	}
	if value, ok := patu.mutation.AppendedAbilities(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, personalaccesstoken.FieldAbilities, value)
		})
	}
	if patu.mutation.AbilitiesCleared() {
		_spec.ClearField(personalaccesstoken.FieldAbilities, field.TypeJSON)
	}
	if value, ok := patu.mutation.ExpiresAt(); ok {
		_spec.SetField(personalaccesstoken.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := patu.mutation.Description(); ok {
		_spec.SetField(personalaccesstoken.FieldDescription, field.TypeString, value)
	}
	if patu.mutation.DescriptionCleared() {
		_spec.ClearField(personalaccesstoken.FieldDescription, field.TypeString)
	}
	if value, ok := patu.mutation.LastUsedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if patu.mutation.LastUsedAtCleared() {
		_spec.ClearField(personalaccesstoken.FieldLastUsedAt, field.TypeTime)
	}
	if patu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccesstoken.OwnerTable,
			Columns: []string{personalaccesstoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = patu.schemaConfig.PersonalAccessToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := patu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccesstoken.OwnerTable,
			Columns: []string{personalaccesstoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = patu.schemaConfig.PersonalAccessToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = patu.schemaConfig.PersonalAccessToken
	ctx = internal.NewSchemaConfigContext(ctx, patu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, patu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	patu.mutation.done = true
	return n, nil
}

// PersonalAccessTokenUpdateOne is the builder for updating a single PersonalAccessToken entity.
type PersonalAccessTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PersonalAccessTokenMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (patuo *PersonalAccessTokenUpdateOne) SetUpdatedAt(t time.Time) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetUpdatedAt(t)
	return patuo
}

// SetUpdatedBy sets the "updated_by" field.
func (patuo *PersonalAccessTokenUpdateOne) SetUpdatedBy(s string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetUpdatedBy(s)
	return patuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableUpdatedBy(s *string) *PersonalAccessTokenUpdateOne {
	if s != nil {
		patuo.SetUpdatedBy(*s)
	}
	return patuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearUpdatedBy() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearUpdatedBy()
	return patuo
}

// SetDeletedAt sets the "deleted_at" field.
func (patuo *PersonalAccessTokenUpdateOne) SetDeletedAt(t time.Time) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetDeletedAt(t)
	return patuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableDeletedAt(t *time.Time) *PersonalAccessTokenUpdateOne {
	if t != nil {
		patuo.SetDeletedAt(*t)
	}
	return patuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearDeletedAt() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearDeletedAt()
	return patuo
}

// SetDeletedBy sets the "deleted_by" field.
func (patuo *PersonalAccessTokenUpdateOne) SetDeletedBy(s string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetDeletedBy(s)
	return patuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableDeletedBy(s *string) *PersonalAccessTokenUpdateOne {
	if s != nil {
		patuo.SetDeletedBy(*s)
	}
	return patuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearDeletedBy() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearDeletedBy()
	return patuo
}

// SetName sets the "name" field.
func (patuo *PersonalAccessTokenUpdateOne) SetName(s string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetName(s)
	return patuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableName(s *string) *PersonalAccessTokenUpdateOne {
	if s != nil {
		patuo.SetName(*s)
	}
	return patuo
}

// SetAbilities sets the "abilities" field.
func (patuo *PersonalAccessTokenUpdateOne) SetAbilities(s []string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetAbilities(s)
	return patuo
}

// AppendAbilities appends s to the "abilities" field.
func (patuo *PersonalAccessTokenUpdateOne) AppendAbilities(s []string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.AppendAbilities(s)
	return patuo
}

// ClearAbilities clears the value of the "abilities" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearAbilities() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearAbilities()
	return patuo
}

// SetExpiresAt sets the "expires_at" field.
func (patuo *PersonalAccessTokenUpdateOne) SetExpiresAt(t time.Time) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetExpiresAt(t)
	return patuo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableExpiresAt(t *time.Time) *PersonalAccessTokenUpdateOne {
	if t != nil {
		patuo.SetExpiresAt(*t)
	}
	return patuo
}

// SetDescription sets the "description" field.
func (patuo *PersonalAccessTokenUpdateOne) SetDescription(s string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetDescription(s)
	return patuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (patuo *PersonalAccessTokenUpdateOne) SetNillableDescription(s *string) *PersonalAccessTokenUpdateOne {
	if s != nil {
		patuo.SetDescription(*s)
	}
	return patuo
}

// ClearDescription clears the value of the "description" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearDescription() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearDescription()
	return patuo
}

// SetLastUsedAt sets the "last_used_at" field.
func (patuo *PersonalAccessTokenUpdateOne) SetLastUsedAt(t time.Time) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetLastUsedAt(t)
	return patuo
}

// ClearLastUsedAt clears the value of the "last_used_at" field.
func (patuo *PersonalAccessTokenUpdateOne) ClearLastUsedAt() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearLastUsedAt()
	return patuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (patuo *PersonalAccessTokenUpdateOne) SetOwnerID(id string) *PersonalAccessTokenUpdateOne {
	patuo.mutation.SetOwnerID(id)
	return patuo
}

// SetOwner sets the "owner" edge to the User entity.
func (patuo *PersonalAccessTokenUpdateOne) SetOwner(u *User) *PersonalAccessTokenUpdateOne {
	return patuo.SetOwnerID(u.ID)
}

// Mutation returns the PersonalAccessTokenMutation object of the builder.
func (patuo *PersonalAccessTokenUpdateOne) Mutation() *PersonalAccessTokenMutation {
	return patuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (patuo *PersonalAccessTokenUpdateOne) ClearOwner() *PersonalAccessTokenUpdateOne {
	patuo.mutation.ClearOwner()
	return patuo
}

// Where appends a list predicates to the PersonalAccessTokenUpdate builder.
func (patuo *PersonalAccessTokenUpdateOne) Where(ps ...predicate.PersonalAccessToken) *PersonalAccessTokenUpdateOne {
	patuo.mutation.Where(ps...)
	return patuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (patuo *PersonalAccessTokenUpdateOne) Select(field string, fields ...string) *PersonalAccessTokenUpdateOne {
	patuo.fields = append([]string{field}, fields...)
	return patuo
}

// Save executes the query and returns the updated PersonalAccessToken entity.
func (patuo *PersonalAccessTokenUpdateOne) Save(ctx context.Context) (*PersonalAccessToken, error) {
	if err := patuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, patuo.sqlSave, patuo.mutation, patuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (patuo *PersonalAccessTokenUpdateOne) SaveX(ctx context.Context) *PersonalAccessToken {
	node, err := patuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (patuo *PersonalAccessTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := patuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (patuo *PersonalAccessTokenUpdateOne) ExecX(ctx context.Context) {
	if err := patuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (patuo *PersonalAccessTokenUpdateOne) defaults() error {
	if _, ok := patuo.mutation.UpdatedAt(); !ok {
		if personalaccesstoken.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized personalaccesstoken.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := personalaccesstoken.UpdateDefaultUpdatedAt()
		patuo.mutation.SetUpdatedAt(v)
	}
	if _, ok := patuo.mutation.LastUsedAt(); !ok && !patuo.mutation.LastUsedAtCleared() {
		if personalaccesstoken.UpdateDefaultLastUsedAt == nil {
			return fmt.Errorf("generated: uninitialized personalaccesstoken.UpdateDefaultLastUsedAt (forgotten import generated/runtime?)")
		}
		v := personalaccesstoken.UpdateDefaultLastUsedAt()
		patuo.mutation.SetLastUsedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (patuo *PersonalAccessTokenUpdateOne) check() error {
	if _, ok := patuo.mutation.OwnerID(); patuo.mutation.OwnerCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "PersonalAccessToken.owner"`)
	}
	return nil
}

func (patuo *PersonalAccessTokenUpdateOne) sqlSave(ctx context.Context) (_node *PersonalAccessToken, err error) {
	if err := patuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(personalaccesstoken.Table, personalaccesstoken.Columns, sqlgraph.NewFieldSpec(personalaccesstoken.FieldID, field.TypeString))
	id, ok := patuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "PersonalAccessToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := patuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, personalaccesstoken.FieldID)
		for _, f := range fields {
			if !personalaccesstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != personalaccesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := patuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := patuo.mutation.UpdatedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if patuo.mutation.CreatedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldCreatedBy, field.TypeString)
	}
	if value, ok := patuo.mutation.UpdatedBy(); ok {
		_spec.SetField(personalaccesstoken.FieldUpdatedBy, field.TypeString, value)
	}
	if patuo.mutation.UpdatedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := patuo.mutation.DeletedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldDeletedAt, field.TypeTime, value)
	}
	if patuo.mutation.DeletedAtCleared() {
		_spec.ClearField(personalaccesstoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := patuo.mutation.DeletedBy(); ok {
		_spec.SetField(personalaccesstoken.FieldDeletedBy, field.TypeString, value)
	}
	if patuo.mutation.DeletedByCleared() {
		_spec.ClearField(personalaccesstoken.FieldDeletedBy, field.TypeString)
	}
	if value, ok := patuo.mutation.Name(); ok {
		_spec.SetField(personalaccesstoken.FieldName, field.TypeString, value)
	}
	if value, ok := patuo.mutation.Abilities(); ok {
		_spec.SetField(personalaccesstoken.FieldAbilities, field.TypeJSON, value)
	}
	if value, ok := patuo.mutation.AppendedAbilities(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, personalaccesstoken.FieldAbilities, value)
		})
	}
	if patuo.mutation.AbilitiesCleared() {
		_spec.ClearField(personalaccesstoken.FieldAbilities, field.TypeJSON)
	}
	if value, ok := patuo.mutation.ExpiresAt(); ok {
		_spec.SetField(personalaccesstoken.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := patuo.mutation.Description(); ok {
		_spec.SetField(personalaccesstoken.FieldDescription, field.TypeString, value)
	}
	if patuo.mutation.DescriptionCleared() {
		_spec.ClearField(personalaccesstoken.FieldDescription, field.TypeString)
	}
	if value, ok := patuo.mutation.LastUsedAt(); ok {
		_spec.SetField(personalaccesstoken.FieldLastUsedAt, field.TypeTime, value)
	}
	if patuo.mutation.LastUsedAtCleared() {
		_spec.ClearField(personalaccesstoken.FieldLastUsedAt, field.TypeTime)
	}
	if patuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccesstoken.OwnerTable,
			Columns: []string{personalaccesstoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = patuo.schemaConfig.PersonalAccessToken
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := patuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   personalaccesstoken.OwnerTable,
			Columns: []string{personalaccesstoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = patuo.schemaConfig.PersonalAccessToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = patuo.schemaConfig.PersonalAccessToken
	ctx = internal.NewSchemaConfigContext(ctx, patuo.schemaConfig)
	_node = &PersonalAccessToken{config: patuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, patuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{personalaccesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	patuo.mutation.done = true
	return _node, nil
}
