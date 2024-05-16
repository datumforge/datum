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
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/organizationsettinghistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrganizationSettingHistoryUpdate is the builder for updating OrganizationSettingHistory entities.
type OrganizationSettingHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationSettingHistoryMutation
}

// Where appends a list predicates to the OrganizationSettingHistoryUpdate builder.
func (oshu *OrganizationSettingHistoryUpdate) Where(ps ...predicate.OrganizationSettingHistory) *OrganizationSettingHistoryUpdate {
	oshu.mutation.Where(ps...)
	return oshu
}

// SetUpdatedAt sets the "updated_at" field.
func (oshu *OrganizationSettingHistoryUpdate) SetUpdatedAt(t time.Time) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetUpdatedAt(t)
	return oshu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *OrganizationSettingHistoryUpdate {
	if t != nil {
		oshu.SetUpdatedAt(*t)
	}
	return oshu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearUpdatedAt() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearUpdatedAt()
	return oshu
}

// SetUpdatedBy sets the "updated_by" field.
func (oshu *OrganizationSettingHistoryUpdate) SetUpdatedBy(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetUpdatedBy(s)
	return oshu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableUpdatedBy(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetUpdatedBy(*s)
	}
	return oshu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearUpdatedBy() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearUpdatedBy()
	return oshu
}

// SetTags sets the "tags" field.
func (oshu *OrganizationSettingHistoryUpdate) SetTags(s []string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetTags(s)
	return oshu
}

// AppendTags appends s to the "tags" field.
func (oshu *OrganizationSettingHistoryUpdate) AppendTags(s []string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.AppendTags(s)
	return oshu
}

// ClearTags clears the value of the "tags" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearTags() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearTags()
	return oshu
}

// SetDeletedAt sets the "deleted_at" field.
func (oshu *OrganizationSettingHistoryUpdate) SetDeletedAt(t time.Time) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetDeletedAt(t)
	return oshu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableDeletedAt(t *time.Time) *OrganizationSettingHistoryUpdate {
	if t != nil {
		oshu.SetDeletedAt(*t)
	}
	return oshu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearDeletedAt() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearDeletedAt()
	return oshu
}

// SetDeletedBy sets the "deleted_by" field.
func (oshu *OrganizationSettingHistoryUpdate) SetDeletedBy(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetDeletedBy(s)
	return oshu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableDeletedBy(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetDeletedBy(*s)
	}
	return oshu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearDeletedBy() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearDeletedBy()
	return oshu
}

// SetDomains sets the "domains" field.
func (oshu *OrganizationSettingHistoryUpdate) SetDomains(s []string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetDomains(s)
	return oshu
}

// AppendDomains appends s to the "domains" field.
func (oshu *OrganizationSettingHistoryUpdate) AppendDomains(s []string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.AppendDomains(s)
	return oshu
}

// ClearDomains clears the value of the "domains" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearDomains() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearDomains()
	return oshu
}

// SetBillingContact sets the "billing_contact" field.
func (oshu *OrganizationSettingHistoryUpdate) SetBillingContact(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetBillingContact(s)
	return oshu
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableBillingContact(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetBillingContact(*s)
	}
	return oshu
}

// ClearBillingContact clears the value of the "billing_contact" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearBillingContact() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearBillingContact()
	return oshu
}

// SetBillingEmail sets the "billing_email" field.
func (oshu *OrganizationSettingHistoryUpdate) SetBillingEmail(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetBillingEmail(s)
	return oshu
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableBillingEmail(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetBillingEmail(*s)
	}
	return oshu
}

// ClearBillingEmail clears the value of the "billing_email" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearBillingEmail() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearBillingEmail()
	return oshu
}

// SetBillingPhone sets the "billing_phone" field.
func (oshu *OrganizationSettingHistoryUpdate) SetBillingPhone(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetBillingPhone(s)
	return oshu
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableBillingPhone(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetBillingPhone(*s)
	}
	return oshu
}

// ClearBillingPhone clears the value of the "billing_phone" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearBillingPhone() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearBillingPhone()
	return oshu
}

// SetBillingAddress sets the "billing_address" field.
func (oshu *OrganizationSettingHistoryUpdate) SetBillingAddress(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetBillingAddress(s)
	return oshu
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableBillingAddress(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetBillingAddress(*s)
	}
	return oshu
}

// ClearBillingAddress clears the value of the "billing_address" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearBillingAddress() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearBillingAddress()
	return oshu
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (oshu *OrganizationSettingHistoryUpdate) SetTaxIdentifier(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetTaxIdentifier(s)
	return oshu
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableTaxIdentifier(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetTaxIdentifier(*s)
	}
	return oshu
}

// ClearTaxIdentifier clears the value of the "tax_identifier" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearTaxIdentifier() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearTaxIdentifier()
	return oshu
}

// SetGeoLocation sets the "geo_location" field.
func (oshu *OrganizationSettingHistoryUpdate) SetGeoLocation(e enums.Region) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetGeoLocation(e)
	return oshu
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableGeoLocation(e *enums.Region) *OrganizationSettingHistoryUpdate {
	if e != nil {
		oshu.SetGeoLocation(*e)
	}
	return oshu
}

// ClearGeoLocation clears the value of the "geo_location" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearGeoLocation() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearGeoLocation()
	return oshu
}

// SetOrganizationID sets the "organization_id" field.
func (oshu *OrganizationSettingHistoryUpdate) SetOrganizationID(s string) *OrganizationSettingHistoryUpdate {
	oshu.mutation.SetOrganizationID(s)
	return oshu
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (oshu *OrganizationSettingHistoryUpdate) SetNillableOrganizationID(s *string) *OrganizationSettingHistoryUpdate {
	if s != nil {
		oshu.SetOrganizationID(*s)
	}
	return oshu
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (oshu *OrganizationSettingHistoryUpdate) ClearOrganizationID() *OrganizationSettingHistoryUpdate {
	oshu.mutation.ClearOrganizationID()
	return oshu
}

// Mutation returns the OrganizationSettingHistoryMutation object of the builder.
func (oshu *OrganizationSettingHistoryUpdate) Mutation() *OrganizationSettingHistoryMutation {
	return oshu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oshu *OrganizationSettingHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, oshu.sqlSave, oshu.mutation, oshu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oshu *OrganizationSettingHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := oshu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oshu *OrganizationSettingHistoryUpdate) Exec(ctx context.Context) error {
	_, err := oshu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshu *OrganizationSettingHistoryUpdate) ExecX(ctx context.Context) {
	if err := oshu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oshu *OrganizationSettingHistoryUpdate) check() error {
	if v, ok := oshu.mutation.GeoLocation(); ok {
		if err := organizationsettinghistory.GeoLocationValidator(v); err != nil {
			return &ValidationError{Name: "geo_location", err: fmt.Errorf(`generated: validator failed for field "OrganizationSettingHistory.geo_location": %w`, err)}
		}
	}
	return nil
}

func (oshu *OrganizationSettingHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := oshu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationsettinghistory.Table, organizationsettinghistory.Columns, sqlgraph.NewFieldSpec(organizationsettinghistory.FieldID, field.TypeString))
	if ps := oshu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if oshu.mutation.RefCleared() {
		_spec.ClearField(organizationsettinghistory.FieldRef, field.TypeString)
	}
	if oshu.mutation.CreatedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := oshu.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if oshu.mutation.UpdatedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if oshu.mutation.CreatedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := oshu.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if oshu.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := oshu.mutation.Tags(); ok {
		_spec.SetField(organizationsettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := oshu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsettinghistory.FieldTags, value)
		})
	}
	if oshu.mutation.TagsCleared() {
		_spec.ClearField(organizationsettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := oshu.mutation.DeletedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if oshu.mutation.DeletedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := oshu.mutation.DeletedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if oshu.mutation.DeletedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := oshu.mutation.Domains(); ok {
		_spec.SetField(organizationsettinghistory.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := oshu.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsettinghistory.FieldDomains, value)
		})
	}
	if oshu.mutation.DomainsCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDomains, field.TypeJSON)
	}
	if value, ok := oshu.mutation.BillingContact(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingContact, field.TypeString, value)
	}
	if oshu.mutation.BillingContactCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingContact, field.TypeString)
	}
	if value, ok := oshu.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingEmail, field.TypeString, value)
	}
	if oshu.mutation.BillingEmailCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingEmail, field.TypeString)
	}
	if value, ok := oshu.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingPhone, field.TypeString, value)
	}
	if oshu.mutation.BillingPhoneCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingPhone, field.TypeString)
	}
	if value, ok := oshu.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingAddress, field.TypeString, value)
	}
	if oshu.mutation.BillingAddressCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingAddress, field.TypeString)
	}
	if value, ok := oshu.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsettinghistory.FieldTaxIdentifier, field.TypeString, value)
	}
	if oshu.mutation.TaxIdentifierCleared() {
		_spec.ClearField(organizationsettinghistory.FieldTaxIdentifier, field.TypeString)
	}
	if value, ok := oshu.mutation.GeoLocation(); ok {
		_spec.SetField(organizationsettinghistory.FieldGeoLocation, field.TypeEnum, value)
	}
	if oshu.mutation.GeoLocationCleared() {
		_spec.ClearField(organizationsettinghistory.FieldGeoLocation, field.TypeEnum)
	}
	if value, ok := oshu.mutation.OrganizationID(); ok {
		_spec.SetField(organizationsettinghistory.FieldOrganizationID, field.TypeString, value)
	}
	if oshu.mutation.OrganizationIDCleared() {
		_spec.ClearField(organizationsettinghistory.FieldOrganizationID, field.TypeString)
	}
	_spec.Node.Schema = oshu.schemaConfig.OrganizationSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, oshu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, oshu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationsettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oshu.mutation.done = true
	return n, nil
}

// OrganizationSettingHistoryUpdateOne is the builder for updating a single OrganizationSettingHistory entity.
type OrganizationSettingHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationSettingHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetUpdatedAt(t time.Time) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetUpdatedAt(t)
	return oshuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrganizationSettingHistoryUpdateOne {
	if t != nil {
		oshuo.SetUpdatedAt(*t)
	}
	return oshuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearUpdatedAt() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearUpdatedAt()
	return oshuo
}

// SetUpdatedBy sets the "updated_by" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetUpdatedBy(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetUpdatedBy(s)
	return oshuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableUpdatedBy(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetUpdatedBy(*s)
	}
	return oshuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearUpdatedBy() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearUpdatedBy()
	return oshuo
}

// SetTags sets the "tags" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetTags(s []string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetTags(s)
	return oshuo
}

// AppendTags appends s to the "tags" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) AppendTags(s []string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.AppendTags(s)
	return oshuo
}

// ClearTags clears the value of the "tags" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearTags() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearTags()
	return oshuo
}

// SetDeletedAt sets the "deleted_at" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetDeletedAt(t time.Time) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetDeletedAt(t)
	return oshuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *OrganizationSettingHistoryUpdateOne {
	if t != nil {
		oshuo.SetDeletedAt(*t)
	}
	return oshuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearDeletedAt() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearDeletedAt()
	return oshuo
}

// SetDeletedBy sets the "deleted_by" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetDeletedBy(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetDeletedBy(s)
	return oshuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableDeletedBy(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetDeletedBy(*s)
	}
	return oshuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearDeletedBy() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearDeletedBy()
	return oshuo
}

// SetDomains sets the "domains" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetDomains(s []string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetDomains(s)
	return oshuo
}

// AppendDomains appends s to the "domains" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) AppendDomains(s []string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.AppendDomains(s)
	return oshuo
}

// ClearDomains clears the value of the "domains" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearDomains() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearDomains()
	return oshuo
}

// SetBillingContact sets the "billing_contact" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetBillingContact(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetBillingContact(s)
	return oshuo
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableBillingContact(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetBillingContact(*s)
	}
	return oshuo
}

// ClearBillingContact clears the value of the "billing_contact" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearBillingContact() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearBillingContact()
	return oshuo
}

// SetBillingEmail sets the "billing_email" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetBillingEmail(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetBillingEmail(s)
	return oshuo
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableBillingEmail(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetBillingEmail(*s)
	}
	return oshuo
}

// ClearBillingEmail clears the value of the "billing_email" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearBillingEmail() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearBillingEmail()
	return oshuo
}

// SetBillingPhone sets the "billing_phone" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetBillingPhone(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetBillingPhone(s)
	return oshuo
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableBillingPhone(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetBillingPhone(*s)
	}
	return oshuo
}

// ClearBillingPhone clears the value of the "billing_phone" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearBillingPhone() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearBillingPhone()
	return oshuo
}

// SetBillingAddress sets the "billing_address" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetBillingAddress(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetBillingAddress(s)
	return oshuo
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableBillingAddress(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetBillingAddress(*s)
	}
	return oshuo
}

// ClearBillingAddress clears the value of the "billing_address" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearBillingAddress() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearBillingAddress()
	return oshuo
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetTaxIdentifier(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetTaxIdentifier(s)
	return oshuo
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableTaxIdentifier(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetTaxIdentifier(*s)
	}
	return oshuo
}

// ClearTaxIdentifier clears the value of the "tax_identifier" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearTaxIdentifier() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearTaxIdentifier()
	return oshuo
}

// SetGeoLocation sets the "geo_location" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetGeoLocation(e enums.Region) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetGeoLocation(e)
	return oshuo
}

// SetNillableGeoLocation sets the "geo_location" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableGeoLocation(e *enums.Region) *OrganizationSettingHistoryUpdateOne {
	if e != nil {
		oshuo.SetGeoLocation(*e)
	}
	return oshuo
}

// ClearGeoLocation clears the value of the "geo_location" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearGeoLocation() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearGeoLocation()
	return oshuo
}

// SetOrganizationID sets the "organization_id" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetOrganizationID(s string) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.SetOrganizationID(s)
	return oshuo
}

// SetNillableOrganizationID sets the "organization_id" field if the given value is not nil.
func (oshuo *OrganizationSettingHistoryUpdateOne) SetNillableOrganizationID(s *string) *OrganizationSettingHistoryUpdateOne {
	if s != nil {
		oshuo.SetOrganizationID(*s)
	}
	return oshuo
}

// ClearOrganizationID clears the value of the "organization_id" field.
func (oshuo *OrganizationSettingHistoryUpdateOne) ClearOrganizationID() *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.ClearOrganizationID()
	return oshuo
}

// Mutation returns the OrganizationSettingHistoryMutation object of the builder.
func (oshuo *OrganizationSettingHistoryUpdateOne) Mutation() *OrganizationSettingHistoryMutation {
	return oshuo.mutation
}

// Where appends a list predicates to the OrganizationSettingHistoryUpdate builder.
func (oshuo *OrganizationSettingHistoryUpdateOne) Where(ps ...predicate.OrganizationSettingHistory) *OrganizationSettingHistoryUpdateOne {
	oshuo.mutation.Where(ps...)
	return oshuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oshuo *OrganizationSettingHistoryUpdateOne) Select(field string, fields ...string) *OrganizationSettingHistoryUpdateOne {
	oshuo.fields = append([]string{field}, fields...)
	return oshuo
}

// Save executes the query and returns the updated OrganizationSettingHistory entity.
func (oshuo *OrganizationSettingHistoryUpdateOne) Save(ctx context.Context) (*OrganizationSettingHistory, error) {
	return withHooks(ctx, oshuo.sqlSave, oshuo.mutation, oshuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oshuo *OrganizationSettingHistoryUpdateOne) SaveX(ctx context.Context) *OrganizationSettingHistory {
	node, err := oshuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oshuo *OrganizationSettingHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := oshuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oshuo *OrganizationSettingHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := oshuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oshuo *OrganizationSettingHistoryUpdateOne) check() error {
	if v, ok := oshuo.mutation.GeoLocation(); ok {
		if err := organizationsettinghistory.GeoLocationValidator(v); err != nil {
			return &ValidationError{Name: "geo_location", err: fmt.Errorf(`generated: validator failed for field "OrganizationSettingHistory.geo_location": %w`, err)}
		}
	}
	return nil
}

func (oshuo *OrganizationSettingHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationSettingHistory, err error) {
	if err := oshuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationsettinghistory.Table, organizationsettinghistory.Columns, sqlgraph.NewFieldSpec(organizationsettinghistory.FieldID, field.TypeString))
	id, ok := oshuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrganizationSettingHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oshuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationsettinghistory.FieldID)
		for _, f := range fields {
			if !organizationsettinghistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != organizationsettinghistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oshuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if oshuo.mutation.RefCleared() {
		_spec.ClearField(organizationsettinghistory.FieldRef, field.TypeString)
	}
	if oshuo.mutation.CreatedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := oshuo.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if oshuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldUpdatedAt, field.TypeTime)
	}
	if oshuo.mutation.CreatedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := oshuo.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldUpdatedBy, field.TypeString, value)
	}
	if oshuo.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := oshuo.mutation.Tags(); ok {
		_spec.SetField(organizationsettinghistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := oshuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsettinghistory.FieldTags, value)
		})
	}
	if oshuo.mutation.TagsCleared() {
		_spec.ClearField(organizationsettinghistory.FieldTags, field.TypeJSON)
	}
	if value, ok := oshuo.mutation.DeletedAt(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedAt, field.TypeTime, value)
	}
	if oshuo.mutation.DeletedAtCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := oshuo.mutation.DeletedBy(); ok {
		_spec.SetField(organizationsettinghistory.FieldDeletedBy, field.TypeString, value)
	}
	if oshuo.mutation.DeletedByCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := oshuo.mutation.Domains(); ok {
		_spec.SetField(organizationsettinghistory.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := oshuo.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsettinghistory.FieldDomains, value)
		})
	}
	if oshuo.mutation.DomainsCleared() {
		_spec.ClearField(organizationsettinghistory.FieldDomains, field.TypeJSON)
	}
	if value, ok := oshuo.mutation.BillingContact(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingContact, field.TypeString, value)
	}
	if oshuo.mutation.BillingContactCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingContact, field.TypeString)
	}
	if value, ok := oshuo.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingEmail, field.TypeString, value)
	}
	if oshuo.mutation.BillingEmailCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingEmail, field.TypeString)
	}
	if value, ok := oshuo.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingPhone, field.TypeString, value)
	}
	if oshuo.mutation.BillingPhoneCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingPhone, field.TypeString)
	}
	if value, ok := oshuo.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsettinghistory.FieldBillingAddress, field.TypeString, value)
	}
	if oshuo.mutation.BillingAddressCleared() {
		_spec.ClearField(organizationsettinghistory.FieldBillingAddress, field.TypeString)
	}
	if value, ok := oshuo.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsettinghistory.FieldTaxIdentifier, field.TypeString, value)
	}
	if oshuo.mutation.TaxIdentifierCleared() {
		_spec.ClearField(organizationsettinghistory.FieldTaxIdentifier, field.TypeString)
	}
	if value, ok := oshuo.mutation.GeoLocation(); ok {
		_spec.SetField(organizationsettinghistory.FieldGeoLocation, field.TypeEnum, value)
	}
	if oshuo.mutation.GeoLocationCleared() {
		_spec.ClearField(organizationsettinghistory.FieldGeoLocation, field.TypeEnum)
	}
	if value, ok := oshuo.mutation.OrganizationID(); ok {
		_spec.SetField(organizationsettinghistory.FieldOrganizationID, field.TypeString, value)
	}
	if oshuo.mutation.OrganizationIDCleared() {
		_spec.ClearField(organizationsettinghistory.FieldOrganizationID, field.TypeString)
	}
	_spec.Node.Schema = oshuo.schemaConfig.OrganizationSettingHistory
	ctx = internal.NewSchemaConfigContext(ctx, oshuo.schemaConfig)
	_node = &OrganizationSettingHistory{config: oshuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oshuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationsettinghistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oshuo.mutation.done = true
	return _node, nil
}
