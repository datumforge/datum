// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/organizationhistory"
	"github.com/datumforge/enthistory"
)

// OrganizationHistoryCreate is the builder for creating a OrganizationHistory entity.
type OrganizationHistoryCreate struct {
	config
	mutation *OrganizationHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ohc *OrganizationHistoryCreate) SetHistoryTime(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetHistoryTime(t)
	return ohc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableHistoryTime(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetHistoryTime(*t)
	}
	return ohc
}

// SetOperation sets the "operation" field.
func (ohc *OrganizationHistoryCreate) SetOperation(et enthistory.OpType) *OrganizationHistoryCreate {
	ohc.mutation.SetOperation(et)
	return ohc
}

// SetRef sets the "ref" field.
func (ohc *OrganizationHistoryCreate) SetRef(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetRef(s)
	return ohc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableRef(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetRef(*s)
	}
	return ohc
}

// SetCreatedAt sets the "created_at" field.
func (ohc *OrganizationHistoryCreate) SetCreatedAt(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetCreatedAt(t)
	return ohc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableCreatedAt(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetCreatedAt(*t)
	}
	return ohc
}

// SetUpdatedAt sets the "updated_at" field.
func (ohc *OrganizationHistoryCreate) SetUpdatedAt(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetUpdatedAt(t)
	return ohc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetUpdatedAt(*t)
	}
	return ohc
}

// SetCreatedBy sets the "created_by" field.
func (ohc *OrganizationHistoryCreate) SetCreatedBy(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetCreatedBy(s)
	return ohc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableCreatedBy(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetCreatedBy(*s)
	}
	return ohc
}

// SetUpdatedBy sets the "updated_by" field.
func (ohc *OrganizationHistoryCreate) SetUpdatedBy(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetUpdatedBy(s)
	return ohc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableUpdatedBy(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetUpdatedBy(*s)
	}
	return ohc
}

// SetMappingID sets the "mapping_id" field.
func (ohc *OrganizationHistoryCreate) SetMappingID(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetMappingID(s)
	return ohc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableMappingID(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetMappingID(*s)
	}
	return ohc
}

// SetTags sets the "tags" field.
func (ohc *OrganizationHistoryCreate) SetTags(s []string) *OrganizationHistoryCreate {
	ohc.mutation.SetTags(s)
	return ohc
}

// SetDeletedAt sets the "deleted_at" field.
func (ohc *OrganizationHistoryCreate) SetDeletedAt(t time.Time) *OrganizationHistoryCreate {
	ohc.mutation.SetDeletedAt(t)
	return ohc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableDeletedAt(t *time.Time) *OrganizationHistoryCreate {
	if t != nil {
		ohc.SetDeletedAt(*t)
	}
	return ohc
}

// SetDeletedBy sets the "deleted_by" field.
func (ohc *OrganizationHistoryCreate) SetDeletedBy(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetDeletedBy(s)
	return ohc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableDeletedBy(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetDeletedBy(*s)
	}
	return ohc
}

// SetName sets the "name" field.
func (ohc *OrganizationHistoryCreate) SetName(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetName(s)
	return ohc
}

// SetDisplayName sets the "display_name" field.
func (ohc *OrganizationHistoryCreate) SetDisplayName(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetDisplayName(s)
	return ohc
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableDisplayName(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetDisplayName(*s)
	}
	return ohc
}

// SetDescription sets the "description" field.
func (ohc *OrganizationHistoryCreate) SetDescription(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetDescription(s)
	return ohc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableDescription(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetDescription(*s)
	}
	return ohc
}

// SetParentOrganizationID sets the "parent_organization_id" field.
func (ohc *OrganizationHistoryCreate) SetParentOrganizationID(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetParentOrganizationID(s)
	return ohc
}

// SetNillableParentOrganizationID sets the "parent_organization_id" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableParentOrganizationID(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetParentOrganizationID(*s)
	}
	return ohc
}

// SetPersonalOrg sets the "personal_org" field.
func (ohc *OrganizationHistoryCreate) SetPersonalOrg(b bool) *OrganizationHistoryCreate {
	ohc.mutation.SetPersonalOrg(b)
	return ohc
}

// SetNillablePersonalOrg sets the "personal_org" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillablePersonalOrg(b *bool) *OrganizationHistoryCreate {
	if b != nil {
		ohc.SetPersonalOrg(*b)
	}
	return ohc
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (ohc *OrganizationHistoryCreate) SetAvatarRemoteURL(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetAvatarRemoteURL(s)
	return ohc
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableAvatarRemoteURL(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetAvatarRemoteURL(*s)
	}
	return ohc
}

// SetDedicatedDb sets the "dedicated_db" field.
func (ohc *OrganizationHistoryCreate) SetDedicatedDb(b bool) *OrganizationHistoryCreate {
	ohc.mutation.SetDedicatedDb(b)
	return ohc
}

// SetNillableDedicatedDb sets the "dedicated_db" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableDedicatedDb(b *bool) *OrganizationHistoryCreate {
	if b != nil {
		ohc.SetDedicatedDb(*b)
	}
	return ohc
}

// SetID sets the "id" field.
func (ohc *OrganizationHistoryCreate) SetID(s string) *OrganizationHistoryCreate {
	ohc.mutation.SetID(s)
	return ohc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ohc *OrganizationHistoryCreate) SetNillableID(s *string) *OrganizationHistoryCreate {
	if s != nil {
		ohc.SetID(*s)
	}
	return ohc
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohc *OrganizationHistoryCreate) Mutation() *OrganizationHistoryMutation {
	return ohc.mutation
}

// Save creates the OrganizationHistory in the database.
func (ohc *OrganizationHistoryCreate) Save(ctx context.Context) (*OrganizationHistory, error) {
	ohc.defaults()
	return withHooks(ctx, ohc.sqlSave, ohc.mutation, ohc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ohc *OrganizationHistoryCreate) SaveX(ctx context.Context) *OrganizationHistory {
	v, err := ohc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohc *OrganizationHistoryCreate) Exec(ctx context.Context) error {
	_, err := ohc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohc *OrganizationHistoryCreate) ExecX(ctx context.Context) {
	if err := ohc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ohc *OrganizationHistoryCreate) defaults() {
	if _, ok := ohc.mutation.HistoryTime(); !ok {
		v := organizationhistory.DefaultHistoryTime()
		ohc.mutation.SetHistoryTime(v)
	}
	if _, ok := ohc.mutation.CreatedAt(); !ok {
		v := organizationhistory.DefaultCreatedAt()
		ohc.mutation.SetCreatedAt(v)
	}
	if _, ok := ohc.mutation.UpdatedAt(); !ok {
		v := organizationhistory.DefaultUpdatedAt()
		ohc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ohc.mutation.MappingID(); !ok {
		v := organizationhistory.DefaultMappingID()
		ohc.mutation.SetMappingID(v)
	}
	if _, ok := ohc.mutation.Tags(); !ok {
		v := organizationhistory.DefaultTags
		ohc.mutation.SetTags(v)
	}
	if _, ok := ohc.mutation.DisplayName(); !ok {
		v := organizationhistory.DefaultDisplayName
		ohc.mutation.SetDisplayName(v)
	}
	if _, ok := ohc.mutation.PersonalOrg(); !ok {
		v := organizationhistory.DefaultPersonalOrg
		ohc.mutation.SetPersonalOrg(v)
	}
	if _, ok := ohc.mutation.DedicatedDb(); !ok {
		v := organizationhistory.DefaultDedicatedDb
		ohc.mutation.SetDedicatedDb(v)
	}
	if _, ok := ohc.mutation.ID(); !ok {
		v := organizationhistory.DefaultID()
		ohc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ohc *OrganizationHistoryCreate) check() error {
	if _, ok := ohc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "OrganizationHistory.history_time"`)}
	}
	if _, ok := ohc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "OrganizationHistory.operation"`)}
	}
	if v, ok := ohc.mutation.Operation(); ok {
		if err := organizationhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "OrganizationHistory.operation": %w`, err)}
		}
	}
	if _, ok := ohc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "OrganizationHistory.mapping_id"`)}
	}
	if _, ok := ohc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`generated: missing required field "OrganizationHistory.name"`)}
	}
	if _, ok := ohc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`generated: missing required field "OrganizationHistory.display_name"`)}
	}
	if _, ok := ohc.mutation.DedicatedDb(); !ok {
		return &ValidationError{Name: "dedicated_db", err: errors.New(`generated: missing required field "OrganizationHistory.dedicated_db"`)}
	}
	return nil
}

func (ohc *OrganizationHistoryCreate) sqlSave(ctx context.Context) (*OrganizationHistory, error) {
	if err := ohc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ohc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ohc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrganizationHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ohc.mutation.id = &_node.ID
	ohc.mutation.done = true
	return _node, nil
}

func (ohc *OrganizationHistoryCreate) createSpec() (*OrganizationHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationHistory{config: ohc.config}
		_spec = sqlgraph.NewCreateSpec(organizationhistory.Table, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeString))
	)
	_spec.Schema = ohc.schemaConfig.OrganizationHistory
	if id, ok := ohc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ohc.mutation.HistoryTime(); ok {
		_spec.SetField(organizationhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ohc.mutation.Operation(); ok {
		_spec.SetField(organizationhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ohc.mutation.Ref(); ok {
		_spec.SetField(organizationhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ohc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ohc.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ohc.mutation.CreatedBy(); ok {
		_spec.SetField(organizationhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ohc.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ohc.mutation.MappingID(); ok {
		_spec.SetField(organizationhistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := ohc.mutation.Tags(); ok {
		_spec.SetField(organizationhistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := ohc.mutation.DeletedAt(); ok {
		_spec.SetField(organizationhistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ohc.mutation.DeletedBy(); ok {
		_spec.SetField(organizationhistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ohc.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ohc.mutation.DisplayName(); ok {
		_spec.SetField(organizationhistory.FieldDisplayName, field.TypeString, value)
		_node.DisplayName = value
	}
	if value, ok := ohc.mutation.Description(); ok {
		_spec.SetField(organizationhistory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ohc.mutation.ParentOrganizationID(); ok {
		_spec.SetField(organizationhistory.FieldParentOrganizationID, field.TypeString, value)
		_node.ParentOrganizationID = value
	}
	if value, ok := ohc.mutation.PersonalOrg(); ok {
		_spec.SetField(organizationhistory.FieldPersonalOrg, field.TypeBool, value)
		_node.PersonalOrg = value
	}
	if value, ok := ohc.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(organizationhistory.FieldAvatarRemoteURL, field.TypeString, value)
		_node.AvatarRemoteURL = &value
	}
	if value, ok := ohc.mutation.DedicatedDb(); ok {
		_spec.SetField(organizationhistory.FieldDedicatedDb, field.TypeBool, value)
		_node.DedicatedDb = value
	}
	return _node, _spec
}

// OrganizationHistoryCreateBulk is the builder for creating many OrganizationHistory entities in bulk.
type OrganizationHistoryCreateBulk struct {
	config
	err      error
	builders []*OrganizationHistoryCreate
}

// Save creates the OrganizationHistory entities in the database.
func (ohcb *OrganizationHistoryCreateBulk) Save(ctx context.Context) ([]*OrganizationHistory, error) {
	if ohcb.err != nil {
		return nil, ohcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ohcb.builders))
	nodes := make([]*OrganizationHistory, len(ohcb.builders))
	mutators := make([]Mutator, len(ohcb.builders))
	for i := range ohcb.builders {
		func(i int, root context.Context) {
			builder := ohcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ohcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ohcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ohcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ohcb *OrganizationHistoryCreateBulk) SaveX(ctx context.Context) []*OrganizationHistory {
	v, err := ohcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ohcb *OrganizationHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ohcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohcb *OrganizationHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ohcb.Exec(ctx); err != nil {
		panic(err)
	}
}
