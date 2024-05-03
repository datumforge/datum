// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettinghistory"
	"github.com/datumforge/enthistory"
)

// OrganizationSettingHistoryCreate is the builder for creating a OrganizationSettingHistory entity.
type OrganizationSettingHistoryCreate struct {
	config
	mutation *OrganizationSettingHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (oshc *OrganizationSettingHistoryCreate) SetHistoryTime(t time.Time) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetHistoryTime(t)
	return oshc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableHistoryTime(t *time.Time) *OrganizationSettingHistoryCreate {
	if t != nil {
		oshc.SetHistoryTime(*t)
	}
	return oshc
}

// SetOperation sets the "operation" field.
func (oshc *OrganizationSettingHistoryCreate) SetOperation(et enthistory.OpType) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetOperation(et)
	return oshc
}

// SetRef sets the "ref" field.
func (oshc *OrganizationSettingHistoryCreate) SetRef(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetRef(s)
	return oshc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableRef(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetRef(*s)
	}
	return oshc
}

// SetCreatedAt sets the "created_at" field.
func (oshc *OrganizationSettingHistoryCreate) SetCreatedAt(t time.Time) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetCreatedAt(t)
	return oshc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableCreatedAt(t *time.Time) *OrganizationSettingHistoryCreate {
	if t != nil {
		oshc.SetCreatedAt(*t)
	}
	return oshc
}

// SetUpdatedAt sets the "updated_at" field.
func (oshc *OrganizationSettingHistoryCreate) SetUpdatedAt(t time.Time) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetUpdatedAt(t)
	return oshc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableUpdatedAt(t *time.Time) *OrganizationSettingHistoryCreate {
	if t != nil {
		oshc.SetUpdatedAt(*t)
	}
	return oshc
}

// SetCreatedBy sets the "created_by" field.
func (oshc *OrganizationSettingHistoryCreate) SetCreatedBy(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetCreatedBy(s)
	return oshc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableCreatedBy(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetCreatedBy(*s)
	}
	return oshc
}

// SetUpdatedBy sets the "updated_by" field.
func (oshc *OrganizationSettingHistoryCreate) SetUpdatedBy(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetUpdatedBy(s)
	return oshc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableUpdatedBy(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetUpdatedBy(*s)
	}
	return oshc
}

// SetMappingID sets the "mapping_id" field.
func (oshc *OrganizationSettingHistoryCreate) SetMappingID(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetMappingID(s)
	return oshc
}

// SetNillableMappingID sets the "mapping_id" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableMappingID(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetMappingID(*s)
	}
	return oshc
}

// SetDeletedAt sets the "deleted_at" field.
func (oshc *OrganizationSettingHistoryCreate) SetDeletedAt(t time.Time) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetDeletedAt(t)
	return oshc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableDeletedAt(t *time.Time) *OrganizationSettingHistoryCreate {
	if t != nil {
		oshc.SetDeletedAt(*t)
	}
	return oshc
}

// SetDeletedBy sets the "deleted_by" field.
func (oshc *OrganizationSettingHistoryCreate) SetDeletedBy(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetDeletedBy(s)
	return oshc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableDeletedBy(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetDeletedBy(*s)
	}
	return oshc
}

// SetDomains sets the "domains" field.
func (oshc *OrganizationSettingHistoryCreate) SetDomains(s []string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetDomains(s)
	return oshc
}

// SetBillingContact sets the "billing_contact" field.
func (oshc *OrganizationSettingHistoryCreate) SetBillingContact(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetBillingContact(s)
	return oshc
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableBillingContact(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetBillingContact(*s)
	}
	return oshc
}

// SetBillingEmail sets the "billing_email" field.
func (oshc *OrganizationSettingHistoryCreate) SetBillingEmail(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetBillingEmail(s)
	return oshc
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableBillingEmail(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetBillingEmail(*s)
	}
	return oshc
}

// SetBillingPhone sets the "billing_phone" field.
func (oshc *OrganizationSettingHistoryCreate) SetBillingPhone(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetBillingPhone(s)
	return oshc
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableBillingPhone(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetBillingPhone(*s)
	}
	return oshc
}

// SetBillingAddress sets the "billing_address" field.
func (oshc *OrganizationSettingHistoryCreate) SetBillingAddress(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetBillingAddress(s)
	return oshc
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableBillingAddress(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetBillingAddress(*s)
	}
	return oshc
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (oshc *OrganizationSettingHistoryCreate) SetTaxIdentifier(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetTaxIdentifier(s)
	return oshc
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableTaxIdentifier(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetTaxIdentifier(*s)
	}
	return oshc
}

// SetTags sets the "tags" field.
func (oshc *OrganizationSettingHistoryCreate) SetTags(s []string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetTags(s)
	return oshc
}

// SetGeoLocation sets the "geo_location" field.
func (oshc *OrganizationSettingHistoryCreate) SetGeoLocation(e enums.Region) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetGeoLocation(e)
	return oshc
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableGeoLocation(e *enums.Region) *OrganizationSettingHistoryCreate {
	if e != nil {
		oshc.SetGeoLocation(*e)
	}
	return oshc
}

// SetOrganizationID sets the "organization_id" field.
func (oshc *OrganizationSettingHistoryCreate) SetOrganizationID(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetOrganizationID(s)
	return oshc
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableOrganizationID(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetOrganizationID(*s)
	}
	return oshc
}

// SetID sets the "id" field.
func (oshc *OrganizationSettingHistoryCreate) SetID(s string) *OrganizationSettingHistoryCreate {
	oshc.mutation.SetID(s)
	return oshc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (oshc *OrganizationSettingHistoryCreate) SetNillableID(s *string) *OrganizationSettingHistoryCreate {
	if s != nil {
		oshc.SetID(*s)
	}
	return oshc
}

// Mutation returns the OrganizationSettingHistoryMutation object of the builder.
func (oshc *OrganizationSettingHistoryCreate) Mutation() *OrganizationSettingHistoryMutation {
	return oshc.mutation
}

// Save creates the OrganizationSettingHistory in the database.
func (oshc *OrganizationSettingHistoryCreate) Save(ctx context.Context) (*OrganizationSettingHistory, error) {
	oshc.defaults()
	return withHooks(ctx, oshc.sqlSave, oshc.mutation, oshc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oshc *OrganizationSettingHistoryCreate) SaveX(ctx context.Context) *OrganizationSettingHistory {
	v, err := oshc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oshc *OrganizationSettingHistoryCreate) Exec(ctx context.Context) error {
	_, err := oshc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshc *OrganizationSettingHistoryCreate) ExecX(ctx context.Context) {
	if err := oshc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oshc *OrganizationSettingHistoryCreate) defaults() {
	if _, ok := oshc.mutation.HistoryTime(); !ok {
		v := organizationsettinghistory.DefaultHistoryTime()
		oshc.mutation.SetHistoryTime(v)
	}
	if _, ok := oshc.mutation.CreatedAt(); !ok {
		v := organizationsettinghistory.DefaultCreatedAt()
		oshc.mutation.SetCreatedAt(v)
	}
	if _, ok := oshc.mutation.UpdatedAt(); !ok {
		v := organizationsettinghistory.DefaultUpdatedAt()
		oshc.mutation.SetUpdatedAt(v)
	}
	if _, ok := oshc.mutation.MappingID(); !ok {
		v := organizationsettinghistory.DefaultMappingID()
		oshc.mutation.SetMappingID(v)
	}
	if _, ok := oshc.mutation.Tags(); !ok {
		v := organizationsettinghistory.DefaultTags
		oshc.mutation.SetTags(v)
	}
	if _, ok := oshc.mutation.GeoLocation(); !ok {
		v := organizationsettinghistory.DefaultGeoLocation
		oshc.mutation.SetGeoLocation(v)
	}
	if _, ok := oshc.mutation.ID(); !ok {
		v := organizationsettinghistory.DefaultID()
		oshc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oshc *OrganizationSettingHistoryCreate) check() error {
	if _, ok := oshc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "OrganizationSettingHistory.history_time"`)}
	}
	if _, ok := oshc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "OrganizationSettingHistory.operation"`)}
	}
	if v, ok := oshc.mutation.Operation(); ok {
		if err := organizationsettinghistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "OrganizationSettingHistory.operation": %w`, err)}
		}
	}
	if _, ok := oshc.mutation.MappingID(); !ok {
		return &ValidationError{Name: "mapping_id", err: errors.New(`generated: missing required field "OrganizationSettingHistory.mapping_id"`)}
	}
	if v, ok := oshc.mutation.GeoLocation(); ok {
		if err := organizationsettinghistory.GeoLocationValidator(v); err != nil {
			return &ValidationError{Name: "geo_location", err: fmt.Errorf(`generated: validator failed for field "OrganizationSettingHistory.geo_location": %w`, err)}
		}
	}
	return nil
}

func (oshc *OrganizationSettingHistoryCreate) sqlSave(ctx context.Context) (*OrganizationSettingHistory, error) {
	if err := oshc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oshc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oshc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected OrganizationSettingHistory.ID type: %T", _spec.ID.Value)
		}
	}
	oshc.mutation.id = &_node.ID
	oshc.mutation.done = true
	return _node, nil
}

func (oshc *OrganizationSettingHistoryCreate) createSpec() (*OrganizationSettingHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationSettingHistory{config: oshc.config}
		_spec = sqlgraph.NewCreateSpec(organizationsettinghistory.Table, sqlgraph.NewFieldSpec(organizationsettinghistory.FieldID, field.TypeString))
	)
	_spec.Schema = oshc.schemaConfig.OrganizationSettingHistory
	if id, ok := oshc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oshc.mutation.HistoryTime(); ok {
		_spec.SetField(organizationsettinghistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := oshc.mutation.Operation(); ok {
		_spec.SetField(organizationsettinghistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := oshc.mutation.Ref(); ok {
		_spec.SetField(organizationsettinghistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := oshc.mutation.CreatedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oshc.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oshc.mutation.CreatedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := oshc.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := oshc.mutation.MappingID(); ok {
		_spec.SetField(organizationsettinghistory.FieldMappingID, field.TypeString, value)
		_node.MappingID = value
	}
	if value, ok := oshc.mutation.DeletedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oshc.mutation.DeletedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := oshc.mutation.Domains(); ok {
		_spec.SetField(organizationsettinghistory.FieldDomains, field.TypeJSON, value)
		_node.Domains = value
	}
	if value, ok := oshc.mutation.BillingContact(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingContact, field.TypeString, value)
		_node.BillingContact = value
	}
	if value, ok := oshc.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingEmail, field.TypeString, value)
		_node.BillingEmail = value
	}
	if value, ok := oshc.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingPhone, field.TypeString, value)
		_node.BillingPhone = value
	}
	if value, ok := oshc.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingAddress, field.TypeString, value)
		_node.BillingAddress = value
	}
	if value, ok := oshc.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsettinghistory.FieldTaxIdentifier, field.TypeString, value)
		_node.TaxIdentifier = value
	}
	if value, ok := oshc.mutation.Tags(); ok {
		_spec.SetField(organizationsettinghistory.FieldTags, field.TypeJSON, value)
		_node.Tags = value
	}
	if value, ok := oshc.mutation.GeoLocation(); ok {
		_spec.SetField(organizationsettinghistory.FieldGeoLocation, field.TypeEnum, value)
		_node.GeoLocation = value
	}
	if value, ok := oshc.mutation.OrganizationID(); ok {
		_spec.SetField(organizationsettinghistory.FieldOrganizationID, field.TypeString, value)
		_node.OrganizationID = value
	}
	return _node, _spec
}

// OrganizationSettingHistoryCreateBulk is the builder for creating many OrganizationSettingHistory entities in bulk.
type OrganizationSettingHistoryCreateBulk struct {
	config
	err      error
	builders []*OrganizationSettingHistoryCreate
}

// Save creates the OrganizationSettingHistory entities in the database.
func (oshcb *OrganizationSettingHistoryCreateBulk) Save(ctx context.Context) ([]*OrganizationSettingHistory, error) {
	if oshcb.err != nil {
		return nil, oshcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oshcb.builders))
	nodes := make([]*OrganizationSettingHistory, len(oshcb.builders))
	mutators := make([]Mutator, len(oshcb.builders))
	for i := range oshcb.builders {
		func(i int, root context.Context) {
			builder := oshcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationSettingHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, oshcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oshcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oshcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oshcb *OrganizationSettingHistoryCreateBulk) SaveX(ctx context.Context) []*OrganizationSettingHistory {
	v, err := oshcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oshcb *OrganizationSettingHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := oshcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshcb *OrganizationSettingHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := oshcb.Exec(ctx); err != nil {
		panic(err)
	}
}
