// Code generated by ent, DO NOT EDIT.

package entitlement

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldDeletedBy, v))
}

// ExternalCustomerID applies equality check predicate on the "external_customer_id" field. It's identical to ExternalCustomerIDEQ.
func ExternalCustomerID(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExternalCustomerID, v))
}

// ExternalSubscriptionID applies equality check predicate on the "external_subscription_id" field. It's identical to ExternalSubscriptionIDEQ.
func ExternalSubscriptionID(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExternalSubscriptionID, v))
}

// Expires applies equality check predicate on the "expires" field. It's identical to ExpiresEQ.
func Expires(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExpires, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExpiresAt, v))
}

// Cancelled applies equality check predicate on the "cancelled" field. It's identical to CancelledEQ.
func Cancelled(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCancelled, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldDeletedBy, v))
}

// TierEQ applies the EQ predicate on the "tier" field.
func TierEQ(v Tier) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldTier, v))
}

// TierNEQ applies the NEQ predicate on the "tier" field.
func TierNEQ(v Tier) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldTier, v))
}

// TierIn applies the In predicate on the "tier" field.
func TierIn(vs ...Tier) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldTier, vs...))
}

// TierNotIn applies the NotIn predicate on the "tier" field.
func TierNotIn(vs ...Tier) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldTier, vs...))
}

// ExternalCustomerIDEQ applies the EQ predicate on the "external_customer_id" field.
func ExternalCustomerIDEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExternalCustomerID, v))
}

// ExternalCustomerIDNEQ applies the NEQ predicate on the "external_customer_id" field.
func ExternalCustomerIDNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldExternalCustomerID, v))
}

// ExternalCustomerIDIn applies the In predicate on the "external_customer_id" field.
func ExternalCustomerIDIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldExternalCustomerID, vs...))
}

// ExternalCustomerIDNotIn applies the NotIn predicate on the "external_customer_id" field.
func ExternalCustomerIDNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldExternalCustomerID, vs...))
}

// ExternalCustomerIDGT applies the GT predicate on the "external_customer_id" field.
func ExternalCustomerIDGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldExternalCustomerID, v))
}

// ExternalCustomerIDGTE applies the GTE predicate on the "external_customer_id" field.
func ExternalCustomerIDGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldExternalCustomerID, v))
}

// ExternalCustomerIDLT applies the LT predicate on the "external_customer_id" field.
func ExternalCustomerIDLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldExternalCustomerID, v))
}

// ExternalCustomerIDLTE applies the LTE predicate on the "external_customer_id" field.
func ExternalCustomerIDLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldExternalCustomerID, v))
}

// ExternalCustomerIDContains applies the Contains predicate on the "external_customer_id" field.
func ExternalCustomerIDContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldExternalCustomerID, v))
}

// ExternalCustomerIDHasPrefix applies the HasPrefix predicate on the "external_customer_id" field.
func ExternalCustomerIDHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldExternalCustomerID, v))
}

// ExternalCustomerIDHasSuffix applies the HasSuffix predicate on the "external_customer_id" field.
func ExternalCustomerIDHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldExternalCustomerID, v))
}

// ExternalCustomerIDIsNil applies the IsNil predicate on the "external_customer_id" field.
func ExternalCustomerIDIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldExternalCustomerID))
}

// ExternalCustomerIDNotNil applies the NotNil predicate on the "external_customer_id" field.
func ExternalCustomerIDNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldExternalCustomerID))
}

// ExternalCustomerIDEqualFold applies the EqualFold predicate on the "external_customer_id" field.
func ExternalCustomerIDEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldExternalCustomerID, v))
}

// ExternalCustomerIDContainsFold applies the ContainsFold predicate on the "external_customer_id" field.
func ExternalCustomerIDContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldExternalCustomerID, v))
}

// ExternalSubscriptionIDEQ applies the EQ predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDNEQ applies the NEQ predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDIn applies the In predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldExternalSubscriptionID, vs...))
}

// ExternalSubscriptionIDNotIn applies the NotIn predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldExternalSubscriptionID, vs...))
}

// ExternalSubscriptionIDGT applies the GT predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDGTE applies the GTE predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDLT applies the LT predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDLTE applies the LTE predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDContains applies the Contains predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDHasPrefix applies the HasPrefix predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDHasSuffix applies the HasSuffix predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDIsNil applies the IsNil predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldExternalSubscriptionID))
}

// ExternalSubscriptionIDNotNil applies the NotNil predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldExternalSubscriptionID))
}

// ExternalSubscriptionIDEqualFold applies the EqualFold predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldExternalSubscriptionID, v))
}

// ExternalSubscriptionIDContainsFold applies the ContainsFold predicate on the "external_subscription_id" field.
func ExternalSubscriptionIDContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldExternalSubscriptionID, v))
}

// ExpiresEQ applies the EQ predicate on the "expires" field.
func ExpiresEQ(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExpires, v))
}

// ExpiresNEQ applies the NEQ predicate on the "expires" field.
func ExpiresNEQ(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldExpires, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldExpiresAt, v))
}

// ExpiresAtIsNil applies the IsNil predicate on the "expires_at" field.
func ExpiresAtIsNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIsNull(FieldExpiresAt))
}

// ExpiresAtNotNil applies the NotNil predicate on the "expires_at" field.
func ExpiresAtNotNil() predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotNull(FieldExpiresAt))
}

// CancelledEQ applies the EQ predicate on the "cancelled" field.
func CancelledEQ(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldCancelled, v))
}

// CancelledNEQ applies the NEQ predicate on the "cancelled" field.
func CancelledNEQ(v bool) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldCancelled, v))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// OrganizationIDGT applies the GT predicate on the "organization_id" field.
func OrganizationIDGT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGT(FieldOrganizationID, v))
}

// OrganizationIDGTE applies the GTE predicate on the "organization_id" field.
func OrganizationIDGTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldGTE(FieldOrganizationID, v))
}

// OrganizationIDLT applies the LT predicate on the "organization_id" field.
func OrganizationIDLT(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLT(FieldOrganizationID, v))
}

// OrganizationIDLTE applies the LTE predicate on the "organization_id" field.
func OrganizationIDLTE(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldLTE(FieldOrganizationID, v))
}

// OrganizationIDContains applies the Contains predicate on the "organization_id" field.
func OrganizationIDContains(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContains(FieldOrganizationID, v))
}

// OrganizationIDHasPrefix applies the HasPrefix predicate on the "organization_id" field.
func OrganizationIDHasPrefix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasPrefix(FieldOrganizationID, v))
}

// OrganizationIDHasSuffix applies the HasSuffix predicate on the "organization_id" field.
func OrganizationIDHasSuffix(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldHasSuffix(FieldOrganizationID, v))
}

// OrganizationIDEqualFold applies the EqualFold predicate on the "organization_id" field.
func OrganizationIDEqualFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldEqualFold(FieldOrganizationID, v))
}

// OrganizationIDContainsFold applies the ContainsFold predicate on the "organization_id" field.
func OrganizationIDContainsFold(v string) predicate.Entitlement {
	return predicate.Entitlement(sql.FieldContainsFold(FieldOrganizationID, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Entitlement {
	return predicate.Entitlement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Entitlement
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Organization) predicate.Entitlement {
	return predicate.Entitlement(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Entitlement
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Entitlement) predicate.Entitlement {
	return predicate.Entitlement(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Entitlement) predicate.Entitlement {
	return predicate.Entitlement(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Entitlement) predicate.Entitlement {
	return predicate.Entitlement(sql.NotPredicates(p))
}
