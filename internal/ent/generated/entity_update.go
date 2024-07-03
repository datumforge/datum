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
	"github.com/datumforge/datum/internal/ent/generated/contact"
	"github.com/datumforge/datum/internal/ent/generated/documentdata"
	"github.com/datumforge/datum/internal/ent/generated/entity"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/pkg/enums"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntityUpdate is the builder for updating Entity entities.
type EntityUpdate struct {
	config
	hooks    []Hook
	mutation *EntityMutation
}

// Where appends a list predicates to the EntityUpdate builder.
func (eu *EntityUpdate) Where(ps ...predicate.Entity) *EntityUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EntityUpdate) SetUpdatedAt(t time.Time) *EntityUpdate {
	eu.mutation.SetUpdatedAt(t)
	return eu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (eu *EntityUpdate) ClearUpdatedAt() *EntityUpdate {
	eu.mutation.ClearUpdatedAt()
	return eu
}

// SetUpdatedBy sets the "updated_by" field.
func (eu *EntityUpdate) SetUpdatedBy(s string) *EntityUpdate {
	eu.mutation.SetUpdatedBy(s)
	return eu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableUpdatedBy(s *string) *EntityUpdate {
	if s != nil {
		eu.SetUpdatedBy(*s)
	}
	return eu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (eu *EntityUpdate) ClearUpdatedBy() *EntityUpdate {
	eu.mutation.ClearUpdatedBy()
	return eu
}

// SetDeletedAt sets the "deleted_at" field.
func (eu *EntityUpdate) SetDeletedAt(t time.Time) *EntityUpdate {
	eu.mutation.SetDeletedAt(t)
	return eu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableDeletedAt(t *time.Time) *EntityUpdate {
	if t != nil {
		eu.SetDeletedAt(*t)
	}
	return eu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (eu *EntityUpdate) ClearDeletedAt() *EntityUpdate {
	eu.mutation.ClearDeletedAt()
	return eu
}

// SetDeletedBy sets the "deleted_by" field.
func (eu *EntityUpdate) SetDeletedBy(s string) *EntityUpdate {
	eu.mutation.SetDeletedBy(s)
	return eu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableDeletedBy(s *string) *EntityUpdate {
	if s != nil {
		eu.SetDeletedBy(*s)
	}
	return eu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (eu *EntityUpdate) ClearDeletedBy() *EntityUpdate {
	eu.mutation.ClearDeletedBy()
	return eu
}

// SetTags sets the "tags" field.
func (eu *EntityUpdate) SetTags(s []string) *EntityUpdate {
	eu.mutation.SetTags(s)
	return eu
}

// AppendTags appends s to the "tags" field.
func (eu *EntityUpdate) AppendTags(s []string) *EntityUpdate {
	eu.mutation.AppendTags(s)
	return eu
}

// ClearTags clears the value of the "tags" field.
func (eu *EntityUpdate) ClearTags() *EntityUpdate {
	eu.mutation.ClearTags()
	return eu
}

// SetOwnerID sets the "owner_id" field.
func (eu *EntityUpdate) SetOwnerID(s string) *EntityUpdate {
	eu.mutation.SetOwnerID(s)
	return eu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableOwnerID(s *string) *EntityUpdate {
	if s != nil {
		eu.SetOwnerID(*s)
	}
	return eu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (eu *EntityUpdate) ClearOwnerID() *EntityUpdate {
	eu.mutation.ClearOwnerID()
	return eu
}

// SetName sets the "name" field.
func (eu *EntityUpdate) SetName(s string) *EntityUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableName(s *string) *EntityUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetDisplayName sets the "display_name" field.
func (eu *EntityUpdate) SetDisplayName(s string) *EntityUpdate {
	eu.mutation.SetDisplayName(s)
	return eu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableDisplayName(s *string) *EntityUpdate {
	if s != nil {
		eu.SetDisplayName(*s)
	}
	return eu
}

// SetDescription sets the "description" field.
func (eu *EntityUpdate) SetDescription(s string) *EntityUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableDescription(s *string) *EntityUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// ClearDescription clears the value of the "description" field.
func (eu *EntityUpdate) ClearDescription() *EntityUpdate {
	eu.mutation.ClearDescription()
	return eu
}

// SetEntityType sets the "entity_type" field.
func (eu *EntityUpdate) SetEntityType(et enums.EntityType) *EntityUpdate {
	eu.mutation.SetEntityType(et)
	return eu
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableEntityType(et *enums.EntityType) *EntityUpdate {
	if et != nil {
		eu.SetEntityType(*et)
	}
	return eu
}

// SetOwner sets the "owner" edge to the Organization entity.
func (eu *EntityUpdate) SetOwner(o *Organization) *EntityUpdate {
	return eu.SetOwnerID(o.ID)
}

// AddContactIDs adds the "contacts" edge to the Contact entity by IDs.
func (eu *EntityUpdate) AddContactIDs(ids ...string) *EntityUpdate {
	eu.mutation.AddContactIDs(ids...)
	return eu
}

// AddContacts adds the "contacts" edges to the Contact entity.
func (eu *EntityUpdate) AddContacts(c ...*Contact) *EntityUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.AddContactIDs(ids...)
}

// AddDocumentIDs adds the "documents" edge to the DocumentData entity by IDs.
func (eu *EntityUpdate) AddDocumentIDs(ids ...string) *EntityUpdate {
	eu.mutation.AddDocumentIDs(ids...)
	return eu
}

// AddDocuments adds the "documents" edges to the DocumentData entity.
func (eu *EntityUpdate) AddDocuments(d ...*DocumentData) *EntityUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.AddDocumentIDs(ids...)
}

// Mutation returns the EntityMutation object of the builder.
func (eu *EntityUpdate) Mutation() *EntityMutation {
	return eu.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (eu *EntityUpdate) ClearOwner() *EntityUpdate {
	eu.mutation.ClearOwner()
	return eu
}

// ClearContacts clears all "contacts" edges to the Contact entity.
func (eu *EntityUpdate) ClearContacts() *EntityUpdate {
	eu.mutation.ClearContacts()
	return eu
}

// RemoveContactIDs removes the "contacts" edge to Contact entities by IDs.
func (eu *EntityUpdate) RemoveContactIDs(ids ...string) *EntityUpdate {
	eu.mutation.RemoveContactIDs(ids...)
	return eu
}

// RemoveContacts removes "contacts" edges to Contact entities.
func (eu *EntityUpdate) RemoveContacts(c ...*Contact) *EntityUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.RemoveContactIDs(ids...)
}

// ClearDocuments clears all "documents" edges to the DocumentData entity.
func (eu *EntityUpdate) ClearDocuments() *EntityUpdate {
	eu.mutation.ClearDocuments()
	return eu
}

// RemoveDocumentIDs removes the "documents" edge to DocumentData entities by IDs.
func (eu *EntityUpdate) RemoveDocumentIDs(ids ...string) *EntityUpdate {
	eu.mutation.RemoveDocumentIDs(ids...)
	return eu
}

// RemoveDocuments removes "documents" edges to DocumentData entities.
func (eu *EntityUpdate) RemoveDocuments(d ...*DocumentData) *EntityUpdate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return eu.RemoveDocumentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntityUpdate) Save(ctx context.Context) (int, error) {
	if err := eu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntityUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntityUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntityUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EntityUpdate) defaults() error {
	if _, ok := eu.mutation.UpdatedAt(); !ok && !eu.mutation.UpdatedAtCleared() {
		if entity.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entity.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entity.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eu *EntityUpdate) check() error {
	if v, ok := eu.mutation.OwnerID(); ok {
		if err := entity.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Entity.owner_id": %w`, err)}
		}
	}
	if v, ok := eu.mutation.Name(); ok {
		if err := entity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Entity.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.DisplayName(); ok {
		if err := entity.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Entity.display_name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.EntityType(); ok {
		if err := entity.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`generated: validator failed for field "Entity.entity_type": %w`, err)}
		}
	}
	return nil
}

func (eu *EntityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(entity.Table, entity.Columns, sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eu.mutation.CreatedAtCleared() {
		_spec.ClearField(entity.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(entity.FieldUpdatedAt, field.TypeTime, value)
	}
	if eu.mutation.UpdatedAtCleared() {
		_spec.ClearField(entity.FieldUpdatedAt, field.TypeTime)
	}
	if eu.mutation.CreatedByCleared() {
		_spec.ClearField(entity.FieldCreatedBy, field.TypeString)
	}
	if value, ok := eu.mutation.UpdatedBy(); ok {
		_spec.SetField(entity.FieldUpdatedBy, field.TypeString, value)
	}
	if eu.mutation.UpdatedByCleared() {
		_spec.ClearField(entity.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := eu.mutation.DeletedAt(); ok {
		_spec.SetField(entity.FieldDeletedAt, field.TypeTime, value)
	}
	if eu.mutation.DeletedAtCleared() {
		_spec.ClearField(entity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := eu.mutation.DeletedBy(); ok {
		_spec.SetField(entity.FieldDeletedBy, field.TypeString, value)
	}
	if eu.mutation.DeletedByCleared() {
		_spec.ClearField(entity.FieldDeletedBy, field.TypeString)
	}
	if value, ok := eu.mutation.Tags(); ok {
		_spec.SetField(entity.FieldTags, field.TypeJSON, value)
	}
	if value, ok := eu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entity.FieldTags, value)
		})
	}
	if eu.mutation.TagsCleared() {
		_spec.ClearField(entity.FieldTags, field.TypeJSON)
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(entity.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.DisplayName(); ok {
		_spec.SetField(entity.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.SetField(entity.FieldDescription, field.TypeString, value)
	}
	if eu.mutation.DescriptionCleared() {
		_spec.ClearField(entity.FieldDescription, field.TypeString)
	}
	if value, ok := eu.mutation.EntityType(); ok {
		_spec.SetField(entity.FieldEntityType, field.TypeEnum, value)
	}
	if eu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.OwnerTable,
			Columns: []string{entity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.Entity
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.OwnerTable,
			Columns: []string{entity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.Entity
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityContacts
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedContactsIDs(); len(nodes) > 0 && !eu.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityContacts
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityContacts
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityDocuments
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !eu.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityDocuments
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = eu.schemaConfig.EntityDocuments
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = eu.schemaConfig.Entity
	ctx = internal.NewSchemaConfigContext(ctx, eu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EntityUpdateOne is the builder for updating a single Entity entity.
type EntityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EntityUpdateOne) SetUpdatedAt(t time.Time) *EntityUpdateOne {
	euo.mutation.SetUpdatedAt(t)
	return euo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (euo *EntityUpdateOne) ClearUpdatedAt() *EntityUpdateOne {
	euo.mutation.ClearUpdatedAt()
	return euo
}

// SetUpdatedBy sets the "updated_by" field.
func (euo *EntityUpdateOne) SetUpdatedBy(s string) *EntityUpdateOne {
	euo.mutation.SetUpdatedBy(s)
	return euo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableUpdatedBy(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetUpdatedBy(*s)
	}
	return euo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (euo *EntityUpdateOne) ClearUpdatedBy() *EntityUpdateOne {
	euo.mutation.ClearUpdatedBy()
	return euo
}

// SetDeletedAt sets the "deleted_at" field.
func (euo *EntityUpdateOne) SetDeletedAt(t time.Time) *EntityUpdateOne {
	euo.mutation.SetDeletedAt(t)
	return euo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableDeletedAt(t *time.Time) *EntityUpdateOne {
	if t != nil {
		euo.SetDeletedAt(*t)
	}
	return euo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (euo *EntityUpdateOne) ClearDeletedAt() *EntityUpdateOne {
	euo.mutation.ClearDeletedAt()
	return euo
}

// SetDeletedBy sets the "deleted_by" field.
func (euo *EntityUpdateOne) SetDeletedBy(s string) *EntityUpdateOne {
	euo.mutation.SetDeletedBy(s)
	return euo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableDeletedBy(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetDeletedBy(*s)
	}
	return euo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (euo *EntityUpdateOne) ClearDeletedBy() *EntityUpdateOne {
	euo.mutation.ClearDeletedBy()
	return euo
}

// SetTags sets the "tags" field.
func (euo *EntityUpdateOne) SetTags(s []string) *EntityUpdateOne {
	euo.mutation.SetTags(s)
	return euo
}

// AppendTags appends s to the "tags" field.
func (euo *EntityUpdateOne) AppendTags(s []string) *EntityUpdateOne {
	euo.mutation.AppendTags(s)
	return euo
}

// ClearTags clears the value of the "tags" field.
func (euo *EntityUpdateOne) ClearTags() *EntityUpdateOne {
	euo.mutation.ClearTags()
	return euo
}

// SetOwnerID sets the "owner_id" field.
func (euo *EntityUpdateOne) SetOwnerID(s string) *EntityUpdateOne {
	euo.mutation.SetOwnerID(s)
	return euo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableOwnerID(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetOwnerID(*s)
	}
	return euo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (euo *EntityUpdateOne) ClearOwnerID() *EntityUpdateOne {
	euo.mutation.ClearOwnerID()
	return euo
}

// SetName sets the "name" field.
func (euo *EntityUpdateOne) SetName(s string) *EntityUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableName(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetDisplayName sets the "display_name" field.
func (euo *EntityUpdateOne) SetDisplayName(s string) *EntityUpdateOne {
	euo.mutation.SetDisplayName(s)
	return euo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableDisplayName(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetDisplayName(*s)
	}
	return euo
}

// SetDescription sets the "description" field.
func (euo *EntityUpdateOne) SetDescription(s string) *EntityUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableDescription(s *string) *EntityUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// ClearDescription clears the value of the "description" field.
func (euo *EntityUpdateOne) ClearDescription() *EntityUpdateOne {
	euo.mutation.ClearDescription()
	return euo
}

// SetEntityType sets the "entity_type" field.
func (euo *EntityUpdateOne) SetEntityType(et enums.EntityType) *EntityUpdateOne {
	euo.mutation.SetEntityType(et)
	return euo
}

// SetNillableEntityType sets the "entity_type" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableEntityType(et *enums.EntityType) *EntityUpdateOne {
	if et != nil {
		euo.SetEntityType(*et)
	}
	return euo
}

// SetOwner sets the "owner" edge to the Organization entity.
func (euo *EntityUpdateOne) SetOwner(o *Organization) *EntityUpdateOne {
	return euo.SetOwnerID(o.ID)
}

// AddContactIDs adds the "contacts" edge to the Contact entity by IDs.
func (euo *EntityUpdateOne) AddContactIDs(ids ...string) *EntityUpdateOne {
	euo.mutation.AddContactIDs(ids...)
	return euo
}

// AddContacts adds the "contacts" edges to the Contact entity.
func (euo *EntityUpdateOne) AddContacts(c ...*Contact) *EntityUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.AddContactIDs(ids...)
}

// AddDocumentIDs adds the "documents" edge to the DocumentData entity by IDs.
func (euo *EntityUpdateOne) AddDocumentIDs(ids ...string) *EntityUpdateOne {
	euo.mutation.AddDocumentIDs(ids...)
	return euo
}

// AddDocuments adds the "documents" edges to the DocumentData entity.
func (euo *EntityUpdateOne) AddDocuments(d ...*DocumentData) *EntityUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.AddDocumentIDs(ids...)
}

// Mutation returns the EntityMutation object of the builder.
func (euo *EntityUpdateOne) Mutation() *EntityMutation {
	return euo.mutation
}

// ClearOwner clears the "owner" edge to the Organization entity.
func (euo *EntityUpdateOne) ClearOwner() *EntityUpdateOne {
	euo.mutation.ClearOwner()
	return euo
}

// ClearContacts clears all "contacts" edges to the Contact entity.
func (euo *EntityUpdateOne) ClearContacts() *EntityUpdateOne {
	euo.mutation.ClearContacts()
	return euo
}

// RemoveContactIDs removes the "contacts" edge to Contact entities by IDs.
func (euo *EntityUpdateOne) RemoveContactIDs(ids ...string) *EntityUpdateOne {
	euo.mutation.RemoveContactIDs(ids...)
	return euo
}

// RemoveContacts removes "contacts" edges to Contact entities.
func (euo *EntityUpdateOne) RemoveContacts(c ...*Contact) *EntityUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.RemoveContactIDs(ids...)
}

// ClearDocuments clears all "documents" edges to the DocumentData entity.
func (euo *EntityUpdateOne) ClearDocuments() *EntityUpdateOne {
	euo.mutation.ClearDocuments()
	return euo
}

// RemoveDocumentIDs removes the "documents" edge to DocumentData entities by IDs.
func (euo *EntityUpdateOne) RemoveDocumentIDs(ids ...string) *EntityUpdateOne {
	euo.mutation.RemoveDocumentIDs(ids...)
	return euo
}

// RemoveDocuments removes "documents" edges to DocumentData entities.
func (euo *EntityUpdateOne) RemoveDocuments(d ...*DocumentData) *EntityUpdateOne {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return euo.RemoveDocumentIDs(ids...)
}

// Where appends a list predicates to the EntityUpdate builder.
func (euo *EntityUpdateOne) Where(ps ...predicate.Entity) *EntityUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntityUpdateOne) Select(field string, fields ...string) *EntityUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entity entity.
func (euo *EntityUpdateOne) Save(ctx context.Context) (*Entity, error) {
	if err := euo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntityUpdateOne) SaveX(ctx context.Context) *Entity {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntityUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntityUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EntityUpdateOne) defaults() error {
	if _, ok := euo.mutation.UpdatedAt(); !ok && !euo.mutation.UpdatedAtCleared() {
		if entity.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized entity.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := entity.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (euo *EntityUpdateOne) check() error {
	if v, ok := euo.mutation.OwnerID(); ok {
		if err := entity.OwnerIDValidator(v); err != nil {
			return &ValidationError{Name: "owner_id", err: fmt.Errorf(`generated: validator failed for field "Entity.owner_id": %w`, err)}
		}
	}
	if v, ok := euo.mutation.Name(); ok {
		if err := entity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Entity.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.DisplayName(); ok {
		if err := entity.DisplayNameValidator(v); err != nil {
			return &ValidationError{Name: "display_name", err: fmt.Errorf(`generated: validator failed for field "Entity.display_name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.EntityType(); ok {
		if err := entity.EntityTypeValidator(v); err != nil {
			return &ValidationError{Name: "entity_type", err: fmt.Errorf(`generated: validator failed for field "Entity.entity_type": %w`, err)}
		}
	}
	return nil
}

func (euo *EntityUpdateOne) sqlSave(ctx context.Context) (_node *Entity, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(entity.Table, entity.Columns, sqlgraph.NewFieldSpec(entity.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Entity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entity.FieldID)
		for _, f := range fields {
			if !entity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != entity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if euo.mutation.CreatedAtCleared() {
		_spec.ClearField(entity.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(entity.FieldUpdatedAt, field.TypeTime, value)
	}
	if euo.mutation.UpdatedAtCleared() {
		_spec.ClearField(entity.FieldUpdatedAt, field.TypeTime)
	}
	if euo.mutation.CreatedByCleared() {
		_spec.ClearField(entity.FieldCreatedBy, field.TypeString)
	}
	if value, ok := euo.mutation.UpdatedBy(); ok {
		_spec.SetField(entity.FieldUpdatedBy, field.TypeString, value)
	}
	if euo.mutation.UpdatedByCleared() {
		_spec.ClearField(entity.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := euo.mutation.DeletedAt(); ok {
		_spec.SetField(entity.FieldDeletedAt, field.TypeTime, value)
	}
	if euo.mutation.DeletedAtCleared() {
		_spec.ClearField(entity.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := euo.mutation.DeletedBy(); ok {
		_spec.SetField(entity.FieldDeletedBy, field.TypeString, value)
	}
	if euo.mutation.DeletedByCleared() {
		_spec.ClearField(entity.FieldDeletedBy, field.TypeString)
	}
	if value, ok := euo.mutation.Tags(); ok {
		_spec.SetField(entity.FieldTags, field.TypeJSON, value)
	}
	if value, ok := euo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, entity.FieldTags, value)
		})
	}
	if euo.mutation.TagsCleared() {
		_spec.ClearField(entity.FieldTags, field.TypeJSON)
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(entity.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.DisplayName(); ok {
		_spec.SetField(entity.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.SetField(entity.FieldDescription, field.TypeString, value)
	}
	if euo.mutation.DescriptionCleared() {
		_spec.ClearField(entity.FieldDescription, field.TypeString)
	}
	if value, ok := euo.mutation.EntityType(); ok {
		_spec.SetField(entity.FieldEntityType, field.TypeEnum, value)
	}
	if euo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.OwnerTable,
			Columns: []string{entity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.Entity
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entity.OwnerTable,
			Columns: []string{entity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.Entity
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityContacts
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedContactsIDs(); len(nodes) > 0 && !euo.mutation.ContactsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityContacts
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ContactsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.ContactsTable,
			Columns: entity.ContactsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityContacts
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityDocuments
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedDocumentsIDs(); len(nodes) > 0 && !euo.mutation.DocumentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityDocuments
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.DocumentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   entity.DocumentsTable,
			Columns: entity.DocumentsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString),
			},
		}
		edge.Schema = euo.schemaConfig.EntityDocuments
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = euo.schemaConfig.Entity
	ctx = internal.NewSchemaConfigContext(ctx, euo.schemaConfig)
	_node = &Entity{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
