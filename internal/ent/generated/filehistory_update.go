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
	"github.com/datumforge/datum/internal/ent/generated/filehistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// FileHistoryUpdate is the builder for updating FileHistory entities.
type FileHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *FileHistoryMutation
}

// Where appends a list predicates to the FileHistoryUpdate builder.
func (fhu *FileHistoryUpdate) Where(ps ...predicate.FileHistory) *FileHistoryUpdate {
	fhu.mutation.Where(ps...)
	return fhu
}

// SetUpdatedAt sets the "updated_at" field.
func (fhu *FileHistoryUpdate) SetUpdatedAt(t time.Time) *FileHistoryUpdate {
	fhu.mutation.SetUpdatedAt(t)
	return fhu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *FileHistoryUpdate {
	if t != nil {
		fhu.SetUpdatedAt(*t)
	}
	return fhu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fhu *FileHistoryUpdate) ClearUpdatedAt() *FileHistoryUpdate {
	fhu.mutation.ClearUpdatedAt()
	return fhu
}

// SetUpdatedBy sets the "updated_by" field.
func (fhu *FileHistoryUpdate) SetUpdatedBy(s string) *FileHistoryUpdate {
	fhu.mutation.SetUpdatedBy(s)
	return fhu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableUpdatedBy(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetUpdatedBy(*s)
	}
	return fhu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fhu *FileHistoryUpdate) ClearUpdatedBy() *FileHistoryUpdate {
	fhu.mutation.ClearUpdatedBy()
	return fhu
}

// SetDeletedAt sets the "deleted_at" field.
func (fhu *FileHistoryUpdate) SetDeletedAt(t time.Time) *FileHistoryUpdate {
	fhu.mutation.SetDeletedAt(t)
	return fhu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableDeletedAt(t *time.Time) *FileHistoryUpdate {
	if t != nil {
		fhu.SetDeletedAt(*t)
	}
	return fhu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fhu *FileHistoryUpdate) ClearDeletedAt() *FileHistoryUpdate {
	fhu.mutation.ClearDeletedAt()
	return fhu
}

// SetDeletedBy sets the "deleted_by" field.
func (fhu *FileHistoryUpdate) SetDeletedBy(s string) *FileHistoryUpdate {
	fhu.mutation.SetDeletedBy(s)
	return fhu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableDeletedBy(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetDeletedBy(*s)
	}
	return fhu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (fhu *FileHistoryUpdate) ClearDeletedBy() *FileHistoryUpdate {
	fhu.mutation.ClearDeletedBy()
	return fhu
}

// SetFileName sets the "file_name" field.
func (fhu *FileHistoryUpdate) SetFileName(s string) *FileHistoryUpdate {
	fhu.mutation.SetFileName(s)
	return fhu
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableFileName(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetFileName(*s)
	}
	return fhu
}

// SetFileExtension sets the "file_extension" field.
func (fhu *FileHistoryUpdate) SetFileExtension(s string) *FileHistoryUpdate {
	fhu.mutation.SetFileExtension(s)
	return fhu
}

// SetNillableFileExtension sets the "file_extension" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableFileExtension(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetFileExtension(*s)
	}
	return fhu
}

// SetFileSize sets the "file_size" field.
func (fhu *FileHistoryUpdate) SetFileSize(i int) *FileHistoryUpdate {
	fhu.mutation.ResetFileSize()
	fhu.mutation.SetFileSize(i)
	return fhu
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableFileSize(i *int) *FileHistoryUpdate {
	if i != nil {
		fhu.SetFileSize(*i)
	}
	return fhu
}

// AddFileSize adds i to the "file_size" field.
func (fhu *FileHistoryUpdate) AddFileSize(i int) *FileHistoryUpdate {
	fhu.mutation.AddFileSize(i)
	return fhu
}

// ClearFileSize clears the value of the "file_size" field.
func (fhu *FileHistoryUpdate) ClearFileSize() *FileHistoryUpdate {
	fhu.mutation.ClearFileSize()
	return fhu
}

// SetContentType sets the "content_type" field.
func (fhu *FileHistoryUpdate) SetContentType(s string) *FileHistoryUpdate {
	fhu.mutation.SetContentType(s)
	return fhu
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableContentType(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetContentType(*s)
	}
	return fhu
}

// SetStoreKey sets the "store_key" field.
func (fhu *FileHistoryUpdate) SetStoreKey(s string) *FileHistoryUpdate {
	fhu.mutation.SetStoreKey(s)
	return fhu
}

// SetNillableStoreKey sets the "store_key" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableStoreKey(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetStoreKey(*s)
	}
	return fhu
}

// SetCategory sets the "category" field.
func (fhu *FileHistoryUpdate) SetCategory(s string) *FileHistoryUpdate {
	fhu.mutation.SetCategory(s)
	return fhu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableCategory(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetCategory(*s)
	}
	return fhu
}

// ClearCategory clears the value of the "category" field.
func (fhu *FileHistoryUpdate) ClearCategory() *FileHistoryUpdate {
	fhu.mutation.ClearCategory()
	return fhu
}

// SetAnnotation sets the "annotation" field.
func (fhu *FileHistoryUpdate) SetAnnotation(s string) *FileHistoryUpdate {
	fhu.mutation.SetAnnotation(s)
	return fhu
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (fhu *FileHistoryUpdate) SetNillableAnnotation(s *string) *FileHistoryUpdate {
	if s != nil {
		fhu.SetAnnotation(*s)
	}
	return fhu
}

// ClearAnnotation clears the value of the "annotation" field.
func (fhu *FileHistoryUpdate) ClearAnnotation() *FileHistoryUpdate {
	fhu.mutation.ClearAnnotation()
	return fhu
}

// Mutation returns the FileHistoryMutation object of the builder.
func (fhu *FileHistoryUpdate) Mutation() *FileHistoryMutation {
	return fhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fhu *FileHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fhu.sqlSave, fhu.mutation, fhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fhu *FileHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := fhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fhu *FileHistoryUpdate) Exec(ctx context.Context) error {
	_, err := fhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhu *FileHistoryUpdate) ExecX(ctx context.Context) {
	if err := fhu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fhu *FileHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(filehistory.Table, filehistory.Columns, sqlgraph.NewFieldSpec(filehistory.FieldID, field.TypeString))
	if ps := fhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fhu.mutation.RefCleared() {
		_spec.ClearField(filehistory.FieldRef, field.TypeString)
	}
	if fhu.mutation.CreatedAtCleared() {
		_spec.ClearField(filehistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := fhu.mutation.UpdatedAt(); ok {
		_spec.SetField(filehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if fhu.mutation.UpdatedAtCleared() {
		_spec.ClearField(filehistory.FieldUpdatedAt, field.TypeTime)
	}
	if fhu.mutation.CreatedByCleared() {
		_spec.ClearField(filehistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := fhu.mutation.UpdatedBy(); ok {
		_spec.SetField(filehistory.FieldUpdatedBy, field.TypeString, value)
	}
	if fhu.mutation.UpdatedByCleared() {
		_spec.ClearField(filehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := fhu.mutation.DeletedAt(); ok {
		_spec.SetField(filehistory.FieldDeletedAt, field.TypeTime, value)
	}
	if fhu.mutation.DeletedAtCleared() {
		_spec.ClearField(filehistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fhu.mutation.DeletedBy(); ok {
		_spec.SetField(filehistory.FieldDeletedBy, field.TypeString, value)
	}
	if fhu.mutation.DeletedByCleared() {
		_spec.ClearField(filehistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := fhu.mutation.FileName(); ok {
		_spec.SetField(filehistory.FieldFileName, field.TypeString, value)
	}
	if value, ok := fhu.mutation.FileExtension(); ok {
		_spec.SetField(filehistory.FieldFileExtension, field.TypeString, value)
	}
	if value, ok := fhu.mutation.FileSize(); ok {
		_spec.SetField(filehistory.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := fhu.mutation.AddedFileSize(); ok {
		_spec.AddField(filehistory.FieldFileSize, field.TypeInt, value)
	}
	if fhu.mutation.FileSizeCleared() {
		_spec.ClearField(filehistory.FieldFileSize, field.TypeInt)
	}
	if value, ok := fhu.mutation.ContentType(); ok {
		_spec.SetField(filehistory.FieldContentType, field.TypeString, value)
	}
	if value, ok := fhu.mutation.StoreKey(); ok {
		_spec.SetField(filehistory.FieldStoreKey, field.TypeString, value)
	}
	if value, ok := fhu.mutation.Category(); ok {
		_spec.SetField(filehistory.FieldCategory, field.TypeString, value)
	}
	if fhu.mutation.CategoryCleared() {
		_spec.ClearField(filehistory.FieldCategory, field.TypeString)
	}
	if value, ok := fhu.mutation.Annotation(); ok {
		_spec.SetField(filehistory.FieldAnnotation, field.TypeString, value)
	}
	if fhu.mutation.AnnotationCleared() {
		_spec.ClearField(filehistory.FieldAnnotation, field.TypeString)
	}
	_spec.Node.Schema = fhu.schemaConfig.FileHistory
	ctx = internal.NewSchemaConfigContext(ctx, fhu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, fhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fhu.mutation.done = true
	return n, nil
}

// FileHistoryUpdateOne is the builder for updating a single FileHistory entity.
type FileHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fhuo *FileHistoryUpdateOne) SetUpdatedAt(t time.Time) *FileHistoryUpdateOne {
	fhuo.mutation.SetUpdatedAt(t)
	return fhuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *FileHistoryUpdateOne {
	if t != nil {
		fhuo.SetUpdatedAt(*t)
	}
	return fhuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fhuo *FileHistoryUpdateOne) ClearUpdatedAt() *FileHistoryUpdateOne {
	fhuo.mutation.ClearUpdatedAt()
	return fhuo
}

// SetUpdatedBy sets the "updated_by" field.
func (fhuo *FileHistoryUpdateOne) SetUpdatedBy(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetUpdatedBy(s)
	return fhuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableUpdatedBy(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetUpdatedBy(*s)
	}
	return fhuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fhuo *FileHistoryUpdateOne) ClearUpdatedBy() *FileHistoryUpdateOne {
	fhuo.mutation.ClearUpdatedBy()
	return fhuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fhuo *FileHistoryUpdateOne) SetDeletedAt(t time.Time) *FileHistoryUpdateOne {
	fhuo.mutation.SetDeletedAt(t)
	return fhuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *FileHistoryUpdateOne {
	if t != nil {
		fhuo.SetDeletedAt(*t)
	}
	return fhuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fhuo *FileHistoryUpdateOne) ClearDeletedAt() *FileHistoryUpdateOne {
	fhuo.mutation.ClearDeletedAt()
	return fhuo
}

// SetDeletedBy sets the "deleted_by" field.
func (fhuo *FileHistoryUpdateOne) SetDeletedBy(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetDeletedBy(s)
	return fhuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableDeletedBy(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetDeletedBy(*s)
	}
	return fhuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (fhuo *FileHistoryUpdateOne) ClearDeletedBy() *FileHistoryUpdateOne {
	fhuo.mutation.ClearDeletedBy()
	return fhuo
}

// SetFileName sets the "file_name" field.
func (fhuo *FileHistoryUpdateOne) SetFileName(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetFileName(s)
	return fhuo
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableFileName(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetFileName(*s)
	}
	return fhuo
}

// SetFileExtension sets the "file_extension" field.
func (fhuo *FileHistoryUpdateOne) SetFileExtension(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetFileExtension(s)
	return fhuo
}

// SetNillableFileExtension sets the "file_extension" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableFileExtension(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetFileExtension(*s)
	}
	return fhuo
}

// SetFileSize sets the "file_size" field.
func (fhuo *FileHistoryUpdateOne) SetFileSize(i int) *FileHistoryUpdateOne {
	fhuo.mutation.ResetFileSize()
	fhuo.mutation.SetFileSize(i)
	return fhuo
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableFileSize(i *int) *FileHistoryUpdateOne {
	if i != nil {
		fhuo.SetFileSize(*i)
	}
	return fhuo
}

// AddFileSize adds i to the "file_size" field.
func (fhuo *FileHistoryUpdateOne) AddFileSize(i int) *FileHistoryUpdateOne {
	fhuo.mutation.AddFileSize(i)
	return fhuo
}

// ClearFileSize clears the value of the "file_size" field.
func (fhuo *FileHistoryUpdateOne) ClearFileSize() *FileHistoryUpdateOne {
	fhuo.mutation.ClearFileSize()
	return fhuo
}

// SetContentType sets the "content_type" field.
func (fhuo *FileHistoryUpdateOne) SetContentType(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetContentType(s)
	return fhuo
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableContentType(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetContentType(*s)
	}
	return fhuo
}

// SetStoreKey sets the "store_key" field.
func (fhuo *FileHistoryUpdateOne) SetStoreKey(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetStoreKey(s)
	return fhuo
}

// SetNillableStoreKey sets the "store_key" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableStoreKey(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetStoreKey(*s)
	}
	return fhuo
}

// SetCategory sets the "category" field.
func (fhuo *FileHistoryUpdateOne) SetCategory(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetCategory(s)
	return fhuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableCategory(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetCategory(*s)
	}
	return fhuo
}

// ClearCategory clears the value of the "category" field.
func (fhuo *FileHistoryUpdateOne) ClearCategory() *FileHistoryUpdateOne {
	fhuo.mutation.ClearCategory()
	return fhuo
}

// SetAnnotation sets the "annotation" field.
func (fhuo *FileHistoryUpdateOne) SetAnnotation(s string) *FileHistoryUpdateOne {
	fhuo.mutation.SetAnnotation(s)
	return fhuo
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (fhuo *FileHistoryUpdateOne) SetNillableAnnotation(s *string) *FileHistoryUpdateOne {
	if s != nil {
		fhuo.SetAnnotation(*s)
	}
	return fhuo
}

// ClearAnnotation clears the value of the "annotation" field.
func (fhuo *FileHistoryUpdateOne) ClearAnnotation() *FileHistoryUpdateOne {
	fhuo.mutation.ClearAnnotation()
	return fhuo
}

// Mutation returns the FileHistoryMutation object of the builder.
func (fhuo *FileHistoryUpdateOne) Mutation() *FileHistoryMutation {
	return fhuo.mutation
}

// Where appends a list predicates to the FileHistoryUpdate builder.
func (fhuo *FileHistoryUpdateOne) Where(ps ...predicate.FileHistory) *FileHistoryUpdateOne {
	fhuo.mutation.Where(ps...)
	return fhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fhuo *FileHistoryUpdateOne) Select(field string, fields ...string) *FileHistoryUpdateOne {
	fhuo.fields = append([]string{field}, fields...)
	return fhuo
}

// Save executes the query and returns the updated FileHistory entity.
func (fhuo *FileHistoryUpdateOne) Save(ctx context.Context) (*FileHistory, error) {
	return withHooks(ctx, fhuo.sqlSave, fhuo.mutation, fhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fhuo *FileHistoryUpdateOne) SaveX(ctx context.Context) *FileHistory {
	node, err := fhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fhuo *FileHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := fhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fhuo *FileHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := fhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fhuo *FileHistoryUpdateOne) sqlSave(ctx context.Context) (_node *FileHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(filehistory.Table, filehistory.Columns, sqlgraph.NewFieldSpec(filehistory.FieldID, field.TypeString))
	id, ok := fhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "FileHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filehistory.FieldID)
		for _, f := range fields {
			if !filehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != filehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fhuo.mutation.RefCleared() {
		_spec.ClearField(filehistory.FieldRef, field.TypeString)
	}
	if fhuo.mutation.CreatedAtCleared() {
		_spec.ClearField(filehistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := fhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(filehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if fhuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(filehistory.FieldUpdatedAt, field.TypeTime)
	}
	if fhuo.mutation.CreatedByCleared() {
		_spec.ClearField(filehistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := fhuo.mutation.UpdatedBy(); ok {
		_spec.SetField(filehistory.FieldUpdatedBy, field.TypeString, value)
	}
	if fhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(filehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := fhuo.mutation.DeletedAt(); ok {
		_spec.SetField(filehistory.FieldDeletedAt, field.TypeTime, value)
	}
	if fhuo.mutation.DeletedAtCleared() {
		_spec.ClearField(filehistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fhuo.mutation.DeletedBy(); ok {
		_spec.SetField(filehistory.FieldDeletedBy, field.TypeString, value)
	}
	if fhuo.mutation.DeletedByCleared() {
		_spec.ClearField(filehistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := fhuo.mutation.FileName(); ok {
		_spec.SetField(filehistory.FieldFileName, field.TypeString, value)
	}
	if value, ok := fhuo.mutation.FileExtension(); ok {
		_spec.SetField(filehistory.FieldFileExtension, field.TypeString, value)
	}
	if value, ok := fhuo.mutation.FileSize(); ok {
		_spec.SetField(filehistory.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := fhuo.mutation.AddedFileSize(); ok {
		_spec.AddField(filehistory.FieldFileSize, field.TypeInt, value)
	}
	if fhuo.mutation.FileSizeCleared() {
		_spec.ClearField(filehistory.FieldFileSize, field.TypeInt)
	}
	if value, ok := fhuo.mutation.ContentType(); ok {
		_spec.SetField(filehistory.FieldContentType, field.TypeString, value)
	}
	if value, ok := fhuo.mutation.StoreKey(); ok {
		_spec.SetField(filehistory.FieldStoreKey, field.TypeString, value)
	}
	if value, ok := fhuo.mutation.Category(); ok {
		_spec.SetField(filehistory.FieldCategory, field.TypeString, value)
	}
	if fhuo.mutation.CategoryCleared() {
		_spec.ClearField(filehistory.FieldCategory, field.TypeString)
	}
	if value, ok := fhuo.mutation.Annotation(); ok {
		_spec.SetField(filehistory.FieldAnnotation, field.TypeString, value)
	}
	if fhuo.mutation.AnnotationCleared() {
		_spec.ClearField(filehistory.FieldAnnotation, field.TypeString)
	}
	_spec.Node.Schema = fhuo.schemaConfig.FileHistory
	ctx = internal.NewSchemaConfigContext(ctx, fhuo.schemaConfig)
	_node = &FileHistory{config: fhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fhuo.mutation.done = true
	return _node, nil
}
