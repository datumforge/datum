// Code generated by ent, DO NOT EDIT.

package personalaccesstoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDeletedBy, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldOwnerID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldName, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldToken, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldExpiresAt, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDescription, v))
}

// LastUsedAt applies equality check predicate on the "last_used_at" field. It's identical to LastUsedAtEQ.
func LastUsedAt(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldDeletedBy, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldOwnerID, v))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldOwnerID, v))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldOwnerID, v))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldOwnerID, v))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldOwnerID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldName, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldToken, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldExpiresAt, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldContainsFold(FieldDescription, v))
}

// ScopesIsNil applies the IsNil predicate on the "scopes" field.
func ScopesIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldScopes))
}

// ScopesNotNil applies the NotNil predicate on the "scopes" field.
func ScopesNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldScopes))
}

// LastUsedAtEQ applies the EQ predicate on the "last_used_at" field.
func LastUsedAtEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldEQ(FieldLastUsedAt, v))
}

// LastUsedAtNEQ applies the NEQ predicate on the "last_used_at" field.
func LastUsedAtNEQ(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNEQ(FieldLastUsedAt, v))
}

// LastUsedAtIn applies the In predicate on the "last_used_at" field.
func LastUsedAtIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIn(FieldLastUsedAt, vs...))
}

// LastUsedAtNotIn applies the NotIn predicate on the "last_used_at" field.
func LastUsedAtNotIn(vs ...time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotIn(FieldLastUsedAt, vs...))
}

// LastUsedAtGT applies the GT predicate on the "last_used_at" field.
func LastUsedAtGT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGT(FieldLastUsedAt, v))
}

// LastUsedAtGTE applies the GTE predicate on the "last_used_at" field.
func LastUsedAtGTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldGTE(FieldLastUsedAt, v))
}

// LastUsedAtLT applies the LT predicate on the "last_used_at" field.
func LastUsedAtLT(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLT(FieldLastUsedAt, v))
}

// LastUsedAtLTE applies the LTE predicate on the "last_used_at" field.
func LastUsedAtLTE(v time.Time) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldLTE(FieldLastUsedAt, v))
}

// LastUsedAtIsNil applies the IsNil predicate on the "last_used_at" field.
func LastUsedAtIsNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldIsNull(FieldLastUsedAt))
}

// LastUsedAtNotNil applies the NotNil predicate on the "last_used_at" field.
func LastUsedAtNotNil() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.FieldNotNull(FieldLastUsedAt))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.PersonalAccessToken
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(func(s *sql.Selector) {
		step := newOwnerStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.PersonalAccessToken
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganizations applies the HasEdge predicate on the "organizations" edge.
func HasOrganizations() predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, OrganizationsTable, OrganizationsPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.OrganizationPersonalAccessTokens
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationsWith applies the HasEdge predicate on the "organizations" edge with a given conditions (other predicates).
func HasOrganizationsWith(preds ...predicate.Organization) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(func(s *sql.Selector) {
		step := newOrganizationsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.OrganizationPersonalAccessTokens
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PersonalAccessToken) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PersonalAccessToken) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PersonalAccessToken) predicate.PersonalAccessToken {
	return predicate.PersonalAccessToken(sql.NotPredicates(p))
}
