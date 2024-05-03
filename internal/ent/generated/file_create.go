// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/file"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FileCreate) SetCreatedAt(t time.Time) *FileCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FileCreate) SetUpdatedAt(t time.Time) *FileCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetCreatedBy sets the "created_by" field.
func (fc *FileCreate) SetCreatedBy(s string) *FileCreate {
	fc.mutation.SetCreatedBy(s)
	return fc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreatedBy(s *string) *FileCreate {
	if s != nil {
		fc.SetCreatedBy(*s)
	}
	return fc
}

// SetUpdatedBy sets the "updated_by" field.
func (fc *FileCreate) SetUpdatedBy(s string) *FileCreate {
	fc.mutation.SetUpdatedBy(s)
	return fc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdatedBy(s *string) *FileCreate {
	if s != nil {
		fc.SetUpdatedBy(*s)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FileCreate) SetDeletedAt(t time.Time) *FileCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedAt(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetDeletedBy sets the "deleted_by" field.
func (fc *FileCreate) SetDeletedBy(s string) *FileCreate {
	fc.mutation.SetDeletedBy(s)
	return fc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fc *FileCreate) SetNillableDeletedBy(s *string) *FileCreate {
	if s != nil {
		fc.SetDeletedBy(*s)
	}
	return fc
}

// SetMappingID sets the "mapping_id" field.
func (fc *FileCreate) SetMappingID(s string) *FileCreate {
	fc.mutation.SetMappingID(s)
	return fc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (fc *FileCreate) SetNillableMappingID(s *string) *FileCreate {
	if s != nil {
		fc.SetMappingID(*s)
	}
	return fc
}

// SetFileName sets the "file_name" field.
func (fc *FileCreate) SetFileName(s string) *FileCreate {
	fc.mutation.SetFileName(s)
	return fc
}

// SetFileExtension sets the "file_extension" field.
func (fc *FileCreate) SetFileExtension(s string) *FileCreate {
	fc.mutation.SetFileExtension(s)
	return fc
}

// SetFileSize sets the "file_size" field.
func (fc *FileCreate) SetFileSize(i int) *FileCreate {
	fc.mutation.SetFileSize(i)
	return fc
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fc *FileCreate) SetNillableFileSize(i *int) *FileCreate {
	if i != nil {
		fc.SetFileSize(*i)
	}
	return fc
}

// SetContentType sets the "content_type" field.
func (fc *FileCreate) SetContentType(s string) *FileCreate {
	fc.mutation.SetContentType(s)
	return fc
}

// SetStoreKey sets the "store_key" field.
func (fc *FileCreate) SetStoreKey(s string) *FileCreate {
	fc.mutation.SetStoreKey(s)
	return fc
}

// SetCategory sets the "category" field.
func (fc *FileCreate) SetCategory(s string) *FileCreate {
	fc.mutation.SetCategory(s)
	return fc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (fc *FileCreate) SetNillableCategory(s *string) *FileCreate {
	if s != nil {
		fc.SetCategory(*s)
	}
	return fc
}

// SetAnnotation sets the "annotation" field.
func (fc *FileCreate) SetAnnotation(s string) *FileCreate {
	fc.mutation.SetAnnotation(s)
	return fc
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (fc *FileCreate) SetNillableAnnotation(s *string) *FileCreate {
	if s != nil {
		fc.SetAnnotation(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(s string) *FileCreate {
	fc.mutation.SetID(s)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FileCreate) SetNillableID(s *string) *FileCreate {
	if s != nil {
		fc.SetID(*s)
	}
	return fc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fc *FileCreate) SetUserID(id string) *FileCreate {
	fc.mutation.SetUserID(id)
	return fc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableUserID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetUserID(*id)
	}
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FileCreate) SetUser(u *User) *FileCreate {
	return fc.SetUserID(u.ID)
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (fc *FileCreate) AddOrganizationIDs(ids ...string) *FileCreate {
	fc.mutation.AddOrganizationIDs(ids...)
	return fc
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (fc *FileCreate) AddOrganization(o ...*Organization) *FileCreate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return fc.AddOrganizationIDs(ids...)
}

// AddGroupIDs adds the "group" edge to the Group entity by IDs.
func (fc *FileCreate) AddGroupIDs(ids ...string) *FileCreate {
	fc.mutation.AddGroupIDs(ids...)
	return fc
}

// AddGroup adds the "group" edges to the Group entity.
func (fc *FileCreate) AddGroup(g ...*Group) *FileCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fc.AddGroupIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	if err := fc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		if file.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized file.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := file.DefaultCreatedAt()
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		if file.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized file.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := file.DefaultUpdatedAt()
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.MappingID(); !ok {
		if file.DefaultMappingID == nil {
			return fmt.Errorf("generated: uninitialized file.DefaultMappingID (forgotten import generated/runtime?)")
		}
		v := file.DefaultMappingID()
		fc.mutation.SetMappingID(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		if file.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized file.DefaultID (forgotten import generated/runtime?)")
		}
		v := file.DefaultID()
		fc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "File.mapping_id"`)}
	}
	if _, ok := fc.mutation.FileName(); !ok {
		return &ValidationError{Name: "file_name", err: errors.New(`generated: missing required field "File.file_name"`)}
	}
	if _, ok := fc.mutation.FileExtension(); !ok {
		return &ValidationError{Name: "file_extension", err: errors.New(`generated: missing required field "File.file_extension"`)}
	}
	if v, ok := fc.mutation.FileSize(); ok {
		if err := file.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`generated: validator failed for field "File.file_size": %w`, err)}
		}
	}
	if _, ok := fc.mutation.ContentType(); !ok {
		return &ValidationError{Name: "content_type", err: errors.New(`generated: missing required field "File.content_type"`)}
	}
	if _, ok := fc.mutation.StoreKey(); !ok {
		return &ValidationError{Name: "store_key", err: errors.New(`generated: missing required field "File.store_key"`)}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected File.ID type: %T", _spec.ID.Value)
		}
	}
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(file.Table, sqlgraph.NewFieldSpec(file.FieldID, field.TypeString))
	)
	_spec.Schema = fc.schemaConfig.File
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(file.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.CreatedBy(); ok {
		_spec.SetField(file.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := fc.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.SetField(file.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := fc.mutation.DeletedBy(); ok {
		_spec.SetField(file.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := fc.mutation.MappingID(); ok {
		_spec.SetField(file.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := fc.mutation.FileName(); ok {
		_spec.SetField(file.FieldFileName, field.TypeString, value)
		_node.FileName = value
	}
	if value, ok := fc.mutation.FileExtension(); ok {
		_spec.SetField(file.FieldFileExtension, field.TypeString, value)
		_node.FileExtension = value
	}
	if value, ok := fc.mutation.FileSize(); ok {
		_spec.SetField(file.FieldFileSize, field.TypeInt, value)
		_node.FileSize = value
	}
	if value, ok := fc.mutation.ContentType(); ok {
		_spec.SetField(file.FieldContentType, field.TypeString, value)
		_node.ContentType = value
	}
	if value, ok := fc.mutation.StoreKey(); ok {
		_spec.SetField(file.FieldStoreKey, field.TypeString, value)
		_node.StoreKey = value
	}
	if value, ok := fc.mutation.Category(); ok {
		_spec.SetField(file.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := fc.mutation.Annotation(); ok {
		_spec.SetField(file.FieldAnnotation, field.TypeString, value)
		_node.Annotation = value
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.File
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_files = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.OrganizationFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fc.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
