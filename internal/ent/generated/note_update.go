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
	"github.com/datumforge/datum/internal/ent/generated/entity"
	"github.com/datumforge/datum/internal/ent/generated/note"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// NoteUpdate is the builder for updating Note entities.
type NoteUpdate struct {
	config
	hooks    []Hook
	mutation *NoteMutation
}

// Where appends a list predicates to the NoteUpdate builder.
func (nu *NoteUpdate) Where(ps ...predicate.Note) *NoteUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NoteUpdate) SetUpdatedAt(t time.Time) *NoteUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nu *NoteUpdate) ClearUpdatedAt() *NoteUpdate {
	nu.mutation.ClearUpdatedAt()
	return nu
}

// SetUpdatedBy sets the "updated_by" field.
func (nu *NoteUpdate) SetUpdatedBy(s string) *NoteUpdate {
	nu.mutation.SetUpdatedBy(s)
	return nu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableUpdatedBy(s *string) *NoteUpdate {
	if s != nil {
		nu.SetUpdatedBy(*s)
	}
	return nu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (nu *NoteUpdate) ClearUpdatedBy() *NoteUpdate {
	nu.mutation.ClearUpdatedBy()
	return nu
}

// SetDeletedAt sets the "deleted_at" field.
func (nu *NoteUpdate) SetDeletedAt(t time.Time) *NoteUpdate {
	nu.mutation.SetDeletedAt(t)
	return nu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableDeletedAt(t *time.Time) *NoteUpdate {
	if t != nil {
		nu.SetDeletedAt(*t)
	}
	return nu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nu *NoteUpdate) ClearDeletedAt() *NoteUpdate {
	nu.mutation.ClearDeletedAt()
	return nu
}

// SetDeletedBy sets the "deleted_by" field.
func (nu *NoteUpdate) SetDeletedBy(s string) *NoteUpdate {
	nu.mutation.SetDeletedBy(s)
	return nu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableDeletedBy(s *string) *NoteUpdate {
	if s != nil {
		nu.SetDeletedBy(*s)
	}
	return nu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (nu *NoteUpdate) ClearDeletedBy() *NoteUpdate {
	nu.mutation.ClearDeletedBy()
	return nu
}

// SetTags sets the "tags" field.
func (nu *NoteUpdate) SetTags(s []string) *NoteUpdate {
	nu.mutation.SetTags(s)
	return nu
}

// AppendTags appends s to the "tags" field.
func (nu *NoteUpdate) AppendTags(s []string) *NoteUpdate {
	nu.mutation.AppendTags(s)
	return nu
}

// ClearTags clears the value of the "tags" field.
func (nu *NoteUpdate) ClearTags() *NoteUpdate {
	nu.mutation.ClearTags()
	return nu
}

// SetOwnerID sets the "owner_id" field.
func (nu *NoteUpdate) SetOwnerID(s string) *NoteUpdate {
	nu.mutation.SetOwnerID(s)
	return nu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableOwnerID(s *string) *NoteUpdate {
	if s != nil {
		nu.SetOwnerID(*s)
	}
	return nu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (nu *NoteUpdate) ClearOwnerID() *NoteUpdate {
	nu.mutation.ClearOwnerID()
	return nu
}

// SetText sets the "text" field.
func (nu *NoteUpdate) SetText(s string) *NoteUpdate {
	nu.mutation.SetText(s)
	return nu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nu *NoteUpdate) SetNillableText(s *string) *NoteUpdate {
	if s != nil {
		nu.SetText(*s)
	}
	return nu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (nu *NoteUpdate) SetOwner(o *Organization) *NoteUpdate {
	return nu.SetOwnerID(o.ID)
}

// SetEntityID sets the "entity" edge to the Entity entity by ID.
func (nu *NoteUpdate) SetEntityID(id string) *NoteUpdate {
	nu.mutation.SetEntityID(id)
	return nu
}

// SetNillableEntityID sets the "entity" edge to the Entity entity by ID if the given value is not nil.
func (nu *NoteUpdate) SetNillableEntityID(id *string) *NoteUpdate {
	if id != nil {
		nu = nu.SetEntityID(*id)
	}
	return nu
}

// SetEntity sets the "entity" edge to the Entity entity.
func (nu *NoteUpdate) SetEntity(e *Entity) *NoteUpdate {
	return nu.SetEntityID(e.ID)
}

// Mutation returns the NoteMutation object of the builder.
func (nu *NoteUpdate) Mutation() *NoteMutation {
	return nu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (nu *NoteUpdate) ClearOwner() *NoteUpdate {
	nu.mutation.ClearOwner()
	return nu
}

// ClearEntity clears the "entity" edge to the Entity entity.
func (nu *NoteUpdate) ClearEntity() *NoteUpdate {
	nu.mutation.ClearEntity()
	return nu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NoteUpdate) Save(ctx context.Context) (int, error) {
	if err := nu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NoteUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NoteUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NoteUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NoteUpdate) defaults() error {
	if _, ok := nu.mutation.UpdatedAt(); !ok && !nu.mutation.UpdatedAtCleared() {
		if note.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized note.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := note.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nu *NoteUpdate) check() error {
	if v, ok := nu.mutation.OwnerID(); ok {
		if err := note.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Note.owner_id": %w`, err)}
		}
	}
	if v, ok := nu.mutation.Text(); ok {
		if err := note.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`generated: validator failed for field "Note.text": %w`, err)}
		}
	}
	return nil
}

func (nu *NoteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := nu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nu.mutation.CreatedAtCleared() {
		_spec.ClearField(note.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(note.FieldUpdatedAt, field.TypeTime, value)
	}
	if nu.mutation.UpdatedAtCleared() {
		_spec.ClearField(note.FieldUpdatedAt, field.TypeTime)
	}
	if nu.mutation.CreatedByCleared() {
		_spec.ClearField(note.FieldCreatedBy, field.TypeString)
	}
	if value, ok := nu.mutation.UpdatedBy(); ok {
		_spec.SetField(note.FieldUpdatedBy, field.TypeString, value)
	}
	if nu.mutation.UpdatedByCleared() {
		_spec.ClearField(note.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := nu.mutation.DeletedAt(); ok {
		_spec.SetField(note.FieldDeletedAt, field.TypeTime, value)
	}
	if nu.mutation.DeletedAtCleared() {
		_spec.ClearField(note.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nu.mutation.DeletedBy(); ok {
		_spec.SetField(note.FieldDeletedBy, field.TypeString, value)
	}
	if nu.mutation.DeletedByCleared() {
		_spec.ClearField(note.FieldDeletedBy, field.TypeString)
	}
	if value, ok := nu.mutation.Tags(); ok {
		_spec.SetField(note.FieldTags, field.TypeJSON, value)
	}
	if value, ok := nu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, note.FieldTags, value)
		})
	}
	if nu.mutation.TagsCleared() {
		_spec.ClearField(note.FieldTags, field.TypeJSON)
	}
	if value, ok := nu.mutation.Text(); ok {
		_spec.SetField(note.FieldText, field.TypeString, value)
	}
	if nu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.OwnerTable,
			Columns: []string{note.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = nu.schemaConfig.Note
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.OwnerTable,
			Columns: []string{note.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = nu.schemaConfig.Note
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nu.mutation.EntityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.EntityTable,
			Columns: []string{note.EntityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString),
			},
		}
		edge.Schema = nu.schemaConfig.Note
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.EntityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.EntityTable,
			Columns: []string{note.EntityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString),
			},
		}
		edge.Schema = nu.schemaConfig.Note
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = nu.schemaConfig.Note
	ctx = internal.NewSchemaConfigContext(ctx, nu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NoteUpdateOne is the builder for updating a single Note entity.
type NoteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NoteMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NoteUpdateOne) SetUpdatedAt(t time.Time) *NoteUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (nuo *NoteUpdateOne) ClearUpdatedAt() *NoteUpdateOne {
	nuo.mutation.ClearUpdatedAt()
	return nuo
}

// SetUpdatedBy sets the "updated_by" field.
func (nuo *NoteUpdateOne) SetUpdatedBy(s string) *NoteUpdateOne {
	nuo.mutation.SetUpdatedBy(s)
	return nuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableUpdatedBy(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetUpdatedBy(*s)
	}
	return nuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (nuo *NoteUpdateOne) ClearUpdatedBy() *NoteUpdateOne {
	nuo.mutation.ClearUpdatedBy()
	return nuo
}

// SetDeletedAt sets the "deleted_at" field.
func (nuo *NoteUpdateOne) SetDeletedAt(t time.Time) *NoteUpdateOne {
	nuo.mutation.SetDeletedAt(t)
	return nuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableDeletedAt(t *time.Time) *NoteUpdateOne {
	if t != nil {
		nuo.SetDeletedAt(*t)
	}
	return nuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (nuo *NoteUpdateOne) ClearDeletedAt() *NoteUpdateOne {
	nuo.mutation.ClearDeletedAt()
	return nuo
}

// SetDeletedBy sets the "deleted_by" field.
func (nuo *NoteUpdateOne) SetDeletedBy(s string) *NoteUpdateOne {
	nuo.mutation.SetDeletedBy(s)
	return nuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableDeletedBy(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetDeletedBy(*s)
	}
	return nuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (nuo *NoteUpdateOne) ClearDeletedBy() *NoteUpdateOne {
	nuo.mutation.ClearDeletedBy()
	return nuo
}

// SetTags sets the "tags" field.
func (nuo *NoteUpdateOne) SetTags(s []string) *NoteUpdateOne {
	nuo.mutation.SetTags(s)
	return nuo
}

// AppendTags appends s to the "tags" field.
func (nuo *NoteUpdateOne) AppendTags(s []string) *NoteUpdateOne {
	nuo.mutation.AppendTags(s)
	return nuo
}

// ClearTags clears the value of the "tags" field.
func (nuo *NoteUpdateOne) ClearTags() *NoteUpdateOne {
	nuo.mutation.ClearTags()
	return nuo
}

// SetOwnerID sets the "owner_id" field.
func (nuo *NoteUpdateOne) SetOwnerID(s string) *NoteUpdateOne {
	nuo.mutation.SetOwnerID(s)
	return nuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableOwnerID(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetOwnerID(*s)
	}
	return nuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (nuo *NoteUpdateOne) ClearOwnerID() *NoteUpdateOne {
	nuo.mutation.ClearOwnerID()
	return nuo
}

// SetText sets the "text" field.
func (nuo *NoteUpdateOne) SetText(s string) *NoteUpdateOne {
	nuo.mutation.SetText(s)
	return nuo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableText(s *string) *NoteUpdateOne {
	if s != nil {
		nuo.SetText(*s)
	}
	return nuo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (nuo *NoteUpdateOne) SetOwner(o *Organization) *NoteUpdateOne {
	return nuo.SetOwnerID(o.ID)
}

// SetEntityID sets the "entity" edge to the Entity entity by ID.
func (nuo *NoteUpdateOne) SetEntityID(id string) *NoteUpdateOne {
	nuo.mutation.SetEntityID(id)
	return nuo
}

// SetNillableEntityID sets the "entity" edge to the Entity entity by ID if the given value is not nil.
func (nuo *NoteUpdateOne) SetNillableEntityID(id *string) *NoteUpdateOne {
	if id != nil {
		nuo = nuo.SetEntityID(*id)
	}
	return nuo
}

// SetEntity sets the "entity" edge to the Entity entity.
func (nuo *NoteUpdateOne) SetEntity(e *Entity) *NoteUpdateOne {
	return nuo.SetEntityID(e.ID)
}

// Mutation returns the NoteMutation object of the builder.
func (nuo *NoteUpdateOne) Mutation() *NoteMutation {
	return nuo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (nuo *NoteUpdateOne) ClearOwner() *NoteUpdateOne {
	nuo.mutation.ClearOwner()
	return nuo
}

// ClearEntity clears the "entity" edge to the Entity entity.
func (nuo *NoteUpdateOne) ClearEntity() *NoteUpdateOne {
	nuo.mutation.ClearEntity()
	return nuo
}

// Where appends a list predicates to the NoteUpdate builder.
func (nuo *NoteUpdateOne) Where(ps ...predicate.Note) *NoteUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NoteUpdateOne) Select(field string, fields ...string) *NoteUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Note entity.
func (nuo *NoteUpdateOne) Save(ctx context.Context) (*Note, error) {
	if err := nuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NoteUpdateOne) SaveX(ctx context.Context) *Note {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NoteUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NoteUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NoteUpdateOne) defaults() error {
	if _, ok := nuo.mutation.UpdatedAt(); !ok && !nuo.mutation.UpdatedAtCleared() {
		if note.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized note.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := note.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (nuo *NoteUpdateOne) check() error {
	if v, ok := nuo.mutation.OwnerID(); ok {
		if err := note.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Note.owner_id": %w`, err)}
		}
	}
	if v, ok := nuo.mutation.Text(); ok {
		if err := note.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`generated: validator failed for field "Note.text": %w`, err)}
		}
	}
	return nil
}

func (nuo *NoteUpdateOne) sqlSave(ctx context.Context) (_node *Note, err error) {
	if err := nuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(note.Table, note.Columns, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Note.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, note.FieldID)
		for _, f := range fields {
			if !note.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != note.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nuo.mutation.CreatedAtCleared() {
		_spec.ClearField(note.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(note.FieldUpdatedAt, field.TypeTime, value)
	}
	if nuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(note.FieldUpdatedAt, field.TypeTime)
	}
	if nuo.mutation.CreatedByCleared() {
		_spec.ClearField(note.FieldCreatedBy, field.TypeString)
	}
	if value, ok := nuo.mutation.UpdatedBy(); ok {
		_spec.SetField(note.FieldUpdatedBy, field.TypeString, value)
	}
	if nuo.mutation.UpdatedByCleared() {
		_spec.ClearField(note.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := nuo.mutation.DeletedAt(); ok {
		_spec.SetField(note.FieldDeletedAt, field.TypeTime, value)
	}
	if nuo.mutation.DeletedAtCleared() {
		_spec.ClearField(note.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := nuo.mutation.DeletedBy(); ok {
		_spec.SetField(note.FieldDeletedBy, field.TypeString, value)
	}
	if nuo.mutation.DeletedByCleared() {
		_spec.ClearField(note.FieldDeletedBy, field.TypeString)
	}
	if value, ok := nuo.mutation.Tags(); ok {
		_spec.SetField(note.FieldTags, field.TypeJSON, value)
	}
	if value, ok := nuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, note.FieldTags, value)
		})
	}
	if nuo.mutation.TagsCleared() {
		_spec.ClearField(note.FieldTags, field.TypeJSON)
	}
	if value, ok := nuo.mutation.Text(); ok {
		_spec.SetField(note.FieldText, field.TypeString, value)
	}
	if nuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.OwnerTable,
			Columns: []string{note.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = nuo.schemaConfig.Note
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.OwnerTable,
			Columns: []string{note.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = nuo.schemaConfig.Note
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nuo.mutation.EntityCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.EntityTable,
			Columns: []string{note.EntityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString),
			},
		}
		edge.Schema = nuo.schemaConfig.Note
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.EntityIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.EntityTable,
			Columns: []string{note.EntityColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString),
			},
		}
		edge.Schema = nuo.schemaConfig.Note
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = nuo.schemaConfig.Note
	ctx = internal.NewSchemaConfigContext(ctx, nuo.schemaConfig)
	_node = &Note{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{note.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
