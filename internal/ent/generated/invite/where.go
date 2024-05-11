// Code generated by ent, DO NOT EDIT.

package invite

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedBy, v))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldMappingID, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedBy, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldOwnerID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldToken, v))
}

// Expires applies equality check predicate on the "expires" field. It's identical to ExpiresEQ.
func Expires(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldExpires, v))
}

// Recipient applies equality check predicate on the "recipient" field. It's identical to RecipientEQ.
func Recipient(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRecipient, v))
}

// SendAttempts applies equality check predicate on the "send_attempts" field. It's identical to SendAttemptsEQ.
func SendAttempts(v int) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldSendAttempts, v))
}

// RequestorID applies equality check predicate on the "requestor_id" field. It's identical to RequestorIDEQ.
func RequestorID(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRequestorID, v))
}

// Secret applies equality check predicate on the "secret" field. It's identical to SecretEQ.
func Secret(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldSecret, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldMappingID, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldDeletedBy, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldOwnerID, v))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldOwnerID, v))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldOwnerID, v))
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.Invite {
	return predicate.Invite(sql.FieldIsNull(FieldOwnerID))
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.Invite {
	return predicate.Invite(sql.FieldNotNull(FieldOwnerID))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldOwnerID, v))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldOwnerID, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldToken, v))
}

// ExpiresEQ applies the EQ predicate on the "expires" field.
func ExpiresEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldExpires, v))
}

// ExpiresNEQ applies the NEQ predicate on the "expires" field.
func ExpiresNEQ(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldExpires, v))
}

// ExpiresIn applies the In predicate on the "expires" field.
func ExpiresIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldExpires, vs...))
}

// ExpiresNotIn applies the NotIn predicate on the "expires" field.
func ExpiresNotIn(vs ...time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldExpires, vs...))
}

// ExpiresGT applies the GT predicate on the "expires" field.
func ExpiresGT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldExpires, v))
}

// ExpiresGTE applies the GTE predicate on the "expires" field.
func ExpiresGTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldExpires, v))
}

// ExpiresLT applies the LT predicate on the "expires" field.
func ExpiresLT(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldExpires, v))
}

// ExpiresLTE applies the LTE predicate on the "expires" field.
func ExpiresLTE(v time.Time) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldExpires, v))
}

// RecipientEQ applies the EQ predicate on the "recipient" field.
func RecipientEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRecipient, v))
}

// RecipientNEQ applies the NEQ predicate on the "recipient" field.
func RecipientNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldRecipient, v))
}

// RecipientIn applies the In predicate on the "recipient" field.
func RecipientIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldRecipient, vs...))
}

// RecipientNotIn applies the NotIn predicate on the "recipient" field.
func RecipientNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldRecipient, vs...))
}

// RecipientGT applies the GT predicate on the "recipient" field.
func RecipientGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldRecipient, v))
}

// RecipientGTE applies the GTE predicate on the "recipient" field.
func RecipientGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldRecipient, v))
}

// RecipientLT applies the LT predicate on the "recipient" field.
func RecipientLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldRecipient, v))
}

// RecipientLTE applies the LTE predicate on the "recipient" field.
func RecipientLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldRecipient, v))
}

// RecipientContains applies the Contains predicate on the "recipient" field.
func RecipientContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldRecipient, v))
}

// RecipientHasPrefix applies the HasPrefix predicate on the "recipient" field.
func RecipientHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldRecipient, v))
}

// RecipientHasSuffix applies the HasSuffix predicate on the "recipient" field.
func RecipientHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldRecipient, v))
}

// RecipientEqualFold applies the EqualFold predicate on the "recipient" field.
func RecipientEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldRecipient, v))
}

// RecipientContainsFold applies the ContainsFold predicate on the "recipient" field.
func RecipientContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldRecipient, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v enums.InviteStatus) predicate.Invite {
	vc := v
	return predicate.Invite(sql.FieldEQ(FieldStatus, vc))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v enums.InviteStatus) predicate.Invite {
	vc := v
	return predicate.Invite(sql.FieldNEQ(FieldStatus, vc))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...enums.InviteStatus) predicate.Invite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Invite(sql.FieldIn(FieldStatus, v...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...enums.InviteStatus) predicate.Invite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Invite(sql.FieldNotIn(FieldStatus, v...))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v enums.Role) predicate.Invite {
	vc := v
	return predicate.Invite(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v enums.Role) predicate.Invite {
	vc := v
	return predicate.Invite(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...enums.Role) predicate.Invite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Invite(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...enums.Role) predicate.Invite {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Invite(sql.FieldNotIn(FieldRole, v...))
}

// SendAttemptsEQ applies the EQ predicate on the "send_attempts" field.
func SendAttemptsEQ(v int) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldSendAttempts, v))
}

// SendAttemptsNEQ applies the NEQ predicate on the "send_attempts" field.
func SendAttemptsNEQ(v int) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldSendAttempts, v))
}

// SendAttemptsIn applies the In predicate on the "send_attempts" field.
func SendAttemptsIn(vs ...int) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldSendAttempts, vs...))
}

// SendAttemptsNotIn applies the NotIn predicate on the "send_attempts" field.
func SendAttemptsNotIn(vs ...int) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldSendAttempts, vs...))
}

// SendAttemptsGT applies the GT predicate on the "send_attempts" field.
func SendAttemptsGT(v int) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldSendAttempts, v))
}

// SendAttemptsGTE applies the GTE predicate on the "send_attempts" field.
func SendAttemptsGTE(v int) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldSendAttempts, v))
}

// SendAttemptsLT applies the LT predicate on the "send_attempts" field.
func SendAttemptsLT(v int) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldSendAttempts, v))
}

// SendAttemptsLTE applies the LTE predicate on the "send_attempts" field.
func SendAttemptsLTE(v int) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldSendAttempts, v))
}

// RequestorIDEQ applies the EQ predicate on the "requestor_id" field.
func RequestorIDEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldRequestorID, v))
}

// RequestorIDNEQ applies the NEQ predicate on the "requestor_id" field.
func RequestorIDNEQ(v string) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldRequestorID, v))
}

// RequestorIDIn applies the In predicate on the "requestor_id" field.
func RequestorIDIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldRequestorID, vs...))
}

// RequestorIDNotIn applies the NotIn predicate on the "requestor_id" field.
func RequestorIDNotIn(vs ...string) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldRequestorID, vs...))
}

// RequestorIDGT applies the GT predicate on the "requestor_id" field.
func RequestorIDGT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldRequestorID, v))
}

// RequestorIDGTE applies the GTE predicate on the "requestor_id" field.
func RequestorIDGTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldRequestorID, v))
}

// RequestorIDLT applies the LT predicate on the "requestor_id" field.
func RequestorIDLT(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldRequestorID, v))
}

// RequestorIDLTE applies the LTE predicate on the "requestor_id" field.
func RequestorIDLTE(v string) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldRequestorID, v))
}

// RequestorIDContains applies the Contains predicate on the "requestor_id" field.
func RequestorIDContains(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContains(FieldRequestorID, v))
}

// RequestorIDHasPrefix applies the HasPrefix predicate on the "requestor_id" field.
func RequestorIDHasPrefix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasPrefix(FieldRequestorID, v))
}

// RequestorIDHasSuffix applies the HasSuffix predicate on the "requestor_id" field.
func RequestorIDHasSuffix(v string) predicate.Invite {
	return predicate.Invite(sql.FieldHasSuffix(FieldRequestorID, v))
}

// RequestorIDEqualFold applies the EqualFold predicate on the "requestor_id" field.
func RequestorIDEqualFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldEqualFold(FieldRequestorID, v))
}

// RequestorIDContainsFold applies the ContainsFold predicate on the "requestor_id" field.
func RequestorIDContainsFold(v string) predicate.Invite {
	return predicate.Invite(sql.FieldContainsFold(FieldRequestorID, v))
}

// SecretEQ applies the EQ predicate on the "secret" field.
func SecretEQ(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldEQ(FieldSecret, v))
}

// SecretNEQ applies the NEQ predicate on the "secret" field.
func SecretNEQ(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldNEQ(FieldSecret, v))
}

// SecretIn applies the In predicate on the "secret" field.
func SecretIn(vs ...[]byte) predicate.Invite {
	return predicate.Invite(sql.FieldIn(FieldSecret, vs...))
}

// SecretNotIn applies the NotIn predicate on the "secret" field.
func SecretNotIn(vs ...[]byte) predicate.Invite {
	return predicate.Invite(sql.FieldNotIn(FieldSecret, vs...))
}

// SecretGT applies the GT predicate on the "secret" field.
func SecretGT(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldGT(FieldSecret, v))
}

// SecretGTE applies the GTE predicate on the "secret" field.
func SecretGTE(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldGTE(FieldSecret, v))
}

// SecretLT applies the LT predicate on the "secret" field.
func SecretLT(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldLT(FieldSecret, v))
}

// SecretLTE applies the LTE predicate on the "secret" field.
func SecretLTE(v []byte) predicate.Invite {
	return predicate.Invite(sql.FieldLTE(FieldSecret, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Invite
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Organization) predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Invite
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEvents applies the HasEdge predicate on the "events" edge.
func HasEvents() predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, EventsTable, EventsPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.InviteEvents
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventsWith applies the HasEdge predicate on the "events" edge with a given conditions (other predicates).
func HasEventsWith(preds ...predicate.Event) predicate.Invite {
	return predicate.Invite(func(s *sql.Selector) {
		step := newEventsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Event
		step.Edge.Schema = schemaConfig.InviteEvents
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Invite) predicate.Invite {
	return predicate.Invite(sql.NotPredicates(p))
}
