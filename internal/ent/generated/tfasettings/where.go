// Code generated by ent, DO NOT EDIT.

package tfasettings

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldDeletedBy, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldOwnerID, v))
}

// TfaSecret applies equality check predicate on the "tfa_secret" field. It's identical to TfaSecretEQ.
func TfaSecret(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldTfaSecret, v))
}

// Verified applies equality check predicate on the "verified" field. It's identical to VerifiedEQ.
func Verified(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldVerified, v))
}

// PhoneOtpAllowed applies equality check predicate on the "phone_otp_allowed" field. It's identical to PhoneOtpAllowedEQ.
func PhoneOtpAllowed(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldPhoneOtpAllowed, v))
}

// EmailOtpAllowed applies equality check predicate on the "email_otp_allowed" field. It's identical to EmailOtpAllowedEQ.
func EmailOtpAllowed(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldEmailOtpAllowed, v))
}

// TotpAllowed applies equality check predicate on the "totp_allowed" field. It's identical to TotpAllowedEQ.
func TotpAllowed(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldTotpAllowed, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldDeletedBy, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContains(FieldOwnerID, v))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasPrefix(FieldOwnerID, v))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasSuffix(FieldOwnerID, v))
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldOwnerID))
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldOwnerID))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldOwnerID, v))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldOwnerID, v))
}

// TfaSecretEQ applies the EQ predicate on the "tfa_secret" field.
func TfaSecretEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldTfaSecret, v))
}

// TfaSecretNEQ applies the NEQ predicate on the "tfa_secret" field.
func TfaSecretNEQ(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldTfaSecret, v))
}

// TfaSecretIn applies the In predicate on the "tfa_secret" field.
func TfaSecretIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIn(FieldTfaSecret, vs...))
}

// TfaSecretNotIn applies the NotIn predicate on the "tfa_secret" field.
func TfaSecretNotIn(vs ...string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotIn(FieldTfaSecret, vs...))
}

// TfaSecretGT applies the GT predicate on the "tfa_secret" field.
func TfaSecretGT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGT(FieldTfaSecret, v))
}

// TfaSecretGTE applies the GTE predicate on the "tfa_secret" field.
func TfaSecretGTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldGTE(FieldTfaSecret, v))
}

// TfaSecretLT applies the LT predicate on the "tfa_secret" field.
func TfaSecretLT(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLT(FieldTfaSecret, v))
}

// TfaSecretLTE applies the LTE predicate on the "tfa_secret" field.
func TfaSecretLTE(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldLTE(FieldTfaSecret, v))
}

// TfaSecretContains applies the Contains predicate on the "tfa_secret" field.
func TfaSecretContains(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContains(FieldTfaSecret, v))
}

// TfaSecretHasPrefix applies the HasPrefix predicate on the "tfa_secret" field.
func TfaSecretHasPrefix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasPrefix(FieldTfaSecret, v))
}

// TfaSecretHasSuffix applies the HasSuffix predicate on the "tfa_secret" field.
func TfaSecretHasSuffix(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldHasSuffix(FieldTfaSecret, v))
}

// TfaSecretIsNil applies the IsNil predicate on the "tfa_secret" field.
func TfaSecretIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldTfaSecret))
}

// TfaSecretNotNil applies the NotNil predicate on the "tfa_secret" field.
func TfaSecretNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldTfaSecret))
}

// TfaSecretEqualFold applies the EqualFold predicate on the "tfa_secret" field.
func TfaSecretEqualFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEqualFold(FieldTfaSecret, v))
}

// TfaSecretContainsFold applies the ContainsFold predicate on the "tfa_secret" field.
func TfaSecretContainsFold(v string) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldContainsFold(FieldTfaSecret, v))
}

// VerifiedEQ applies the EQ predicate on the "verified" field.
func VerifiedEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldVerified, v))
}

// VerifiedNEQ applies the NEQ predicate on the "verified" field.
func VerifiedNEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldVerified, v))
}

// RecoveryCodesIsNil applies the IsNil predicate on the "recovery_codes" field.
func RecoveryCodesIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldRecoveryCodes))
}

// RecoveryCodesNotNil applies the NotNil predicate on the "recovery_codes" field.
func RecoveryCodesNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldRecoveryCodes))
}

// PhoneOtpAllowedEQ applies the EQ predicate on the "phone_otp_allowed" field.
func PhoneOtpAllowedEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldPhoneOtpAllowed, v))
}

// PhoneOtpAllowedNEQ applies the NEQ predicate on the "phone_otp_allowed" field.
func PhoneOtpAllowedNEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldPhoneOtpAllowed, v))
}

// PhoneOtpAllowedIsNil applies the IsNil predicate on the "phone_otp_allowed" field.
func PhoneOtpAllowedIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldPhoneOtpAllowed))
}

// PhoneOtpAllowedNotNil applies the NotNil predicate on the "phone_otp_allowed" field.
func PhoneOtpAllowedNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldPhoneOtpAllowed))
}

// EmailOtpAllowedEQ applies the EQ predicate on the "email_otp_allowed" field.
func EmailOtpAllowedEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldEmailOtpAllowed, v))
}

// EmailOtpAllowedNEQ applies the NEQ predicate on the "email_otp_allowed" field.
func EmailOtpAllowedNEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldEmailOtpAllowed, v))
}

// EmailOtpAllowedIsNil applies the IsNil predicate on the "email_otp_allowed" field.
func EmailOtpAllowedIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldEmailOtpAllowed))
}

// EmailOtpAllowedNotNil applies the NotNil predicate on the "email_otp_allowed" field.
func EmailOtpAllowedNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldEmailOtpAllowed))
}

// TotpAllowedEQ applies the EQ predicate on the "totp_allowed" field.
func TotpAllowedEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldEQ(FieldTotpAllowed, v))
}

// TotpAllowedNEQ applies the NEQ predicate on the "totp_allowed" field.
func TotpAllowedNEQ(v bool) predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNEQ(FieldTotpAllowed, v))
}

// TotpAllowedIsNil applies the IsNil predicate on the "totp_allowed" field.
func TotpAllowedIsNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldIsNull(FieldTotpAllowed))
}

// TotpAllowedNotNil applies the NotNil predicate on the "totp_allowed" field.
func TotpAllowedNotNil() predicate.TFASettings {
	return predicate.TFASettings(sql.FieldNotNull(FieldTotpAllowed))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.TFASettings {
	return predicate.TFASettings(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.TFASettings
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.TFASettings {
	return predicate.TFASettings(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.TFASettings
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TFASettings) predicate.TFASettings {
	return predicate.TFASettings(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TFASettings) predicate.TFASettings {
	return predicate.TFASettings(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TFASettings) predicate.TFASettings {
	return predicate.TFASettings(sql.NotPredicates(p))
}
