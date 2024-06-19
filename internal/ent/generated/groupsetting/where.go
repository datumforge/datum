// Code generated by ent, DO NOT EDIT.

package groupsetting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/pkg/enums"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// MappingID applies equality check predicate on the "mapping_id" field. It's identical to MappingIDEQ.
func MappingID(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldMappingID, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldDeletedBy, v))
}

// SyncToSlack applies equality check predicate on the "sync_to_slack" field. It's identical to SyncToSlackEQ.
func SyncToSlack(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldSyncToSlack, v))
}

// SyncToGithub applies equality check predicate on the "sync_to_github" field. It's identical to SyncToGithubEQ.
func SyncToGithub(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldSyncToGithub, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldGroupID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldUpdatedAt))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// MappingIDEQ applies the EQ predicate on the "mapping_id" field.
func MappingIDEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldMappingID, v))
}

// MappingIDNEQ applies the NEQ predicate on the "mapping_id" field.
func MappingIDNEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldMappingID, v))
}

// MappingIDIn applies the In predicate on the "mapping_id" field.
func MappingIDIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldMappingID, vs...))
}

// MappingIDNotIn applies the NotIn predicate on the "mapping_id" field.
func MappingIDNotIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldMappingID, vs...))
}

// MappingIDGT applies the GT predicate on the "mapping_id" field.
func MappingIDGT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldMappingID, v))
}

// MappingIDGTE applies the GTE predicate on the "mapping_id" field.
func MappingIDGTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldMappingID, v))
}

// MappingIDLT applies the LT predicate on the "mapping_id" field.
func MappingIDLT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldMappingID, v))
}

// MappingIDLTE applies the LTE predicate on the "mapping_id" field.
func MappingIDLTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldMappingID, v))
}

// MappingIDContains applies the Contains predicate on the "mapping_id" field.
func MappingIDContains(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContains(FieldMappingID, v))
}

// MappingIDHasPrefix applies the HasPrefix predicate on the "mapping_id" field.
func MappingIDHasPrefix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasPrefix(FieldMappingID, v))
}

// MappingIDHasSuffix applies the HasSuffix predicate on the "mapping_id" field.
func MappingIDHasSuffix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasSuffix(FieldMappingID, v))
}

// MappingIDEqualFold applies the EqualFold predicate on the "mapping_id" field.
func MappingIDEqualFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldMappingID, v))
}

// MappingIDContainsFold applies the ContainsFold predicate on the "mapping_id" field.
func MappingIDContainsFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldMappingID, v))
}

// TagsIsNil applies the IsNil predicate on the "tags" field.
func TagsIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldTags))
}

// TagsNotNil applies the NotNil predicate on the "tags" field.
func TagsNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldTags))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldDeletedBy, v))
}

// VisibilityEQ applies the EQ predicate on the "visibility" field.
func VisibilityEQ(v enums.Visibility) predicate.GroupSetting {
	vc := v
	return predicate.GroupSetting(sql.FieldEQ(FieldVisibility, vc))
}

// VisibilityNEQ applies the NEQ predicate on the "visibility" field.
func VisibilityNEQ(v enums.Visibility) predicate.GroupSetting {
	vc := v
	return predicate.GroupSetting(sql.FieldNEQ(FieldVisibility, vc))
}

// VisibilityIn applies the In predicate on the "visibility" field.
func VisibilityIn(vs ...enums.Visibility) predicate.GroupSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSetting(sql.FieldIn(FieldVisibility, v...))
}

// VisibilityNotIn applies the NotIn predicate on the "visibility" field.
func VisibilityNotIn(vs ...enums.Visibility) predicate.GroupSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSetting(sql.FieldNotIn(FieldVisibility, v...))
}

// JoinPolicyEQ applies the EQ predicate on the "join_policy" field.
func JoinPolicyEQ(v enums.JoinPolicy) predicate.GroupSetting {
	vc := v
	return predicate.GroupSetting(sql.FieldEQ(FieldJoinPolicy, vc))
}

// JoinPolicyNEQ applies the NEQ predicate on the "join_policy" field.
func JoinPolicyNEQ(v enums.JoinPolicy) predicate.GroupSetting {
	vc := v
	return predicate.GroupSetting(sql.FieldNEQ(FieldJoinPolicy, vc))
}

// JoinPolicyIn applies the In predicate on the "join_policy" field.
func JoinPolicyIn(vs ...enums.JoinPolicy) predicate.GroupSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSetting(sql.FieldIn(FieldJoinPolicy, v...))
}

// JoinPolicyNotIn applies the NotIn predicate on the "join_policy" field.
func JoinPolicyNotIn(vs ...enums.JoinPolicy) predicate.GroupSetting {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSetting(sql.FieldNotIn(FieldJoinPolicy, v...))
}

// SyncToSlackEQ applies the EQ predicate on the "sync_to_slack" field.
func SyncToSlackEQ(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldSyncToSlack, v))
}

// SyncToSlackNEQ applies the NEQ predicate on the "sync_to_slack" field.
func SyncToSlackNEQ(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldSyncToSlack, v))
}

// SyncToSlackIsNil applies the IsNil predicate on the "sync_to_slack" field.
func SyncToSlackIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldSyncToSlack))
}

// SyncToSlackNotNil applies the NotNil predicate on the "sync_to_slack" field.
func SyncToSlackNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldSyncToSlack))
}

// SyncToGithubEQ applies the EQ predicate on the "sync_to_github" field.
func SyncToGithubEQ(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldSyncToGithub, v))
}

// SyncToGithubNEQ applies the NEQ predicate on the "sync_to_github" field.
func SyncToGithubNEQ(v bool) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldSyncToGithub, v))
}

// SyncToGithubIsNil applies the IsNil predicate on the "sync_to_github" field.
func SyncToGithubIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldSyncToGithub))
}

// SyncToGithubNotNil applies the NotNil predicate on the "sync_to_github" field.
func SyncToGithubNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldSyncToGithub))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldLTE(FieldGroupID, v))
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContains(FieldGroupID, v))
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasPrefix(FieldGroupID, v))
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldHasSuffix(FieldGroupID, v))
}

// GroupIDIsNil applies the IsNil predicate on the "group_id" field.
func GroupIDIsNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldIsNull(FieldGroupID))
}

// GroupIDNotNil applies the NotNil predicate on the "group_id" field.
func GroupIDNotNil() predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldNotNull(FieldGroupID))
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldEqualFold(FieldGroupID, v))
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.GroupSetting {
	return predicate.GroupSetting(sql.FieldContainsFold(FieldGroupID, v))
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.GroupSetting {
	return predicate.GroupSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, GroupTable, GroupColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.GroupSetting
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.GroupSetting {
	return predicate.GroupSetting(func(s *sql.Selector) {
		step := newGroupStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.GroupSetting
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupSetting) predicate.GroupSetting {
	return predicate.GroupSetting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupSetting) predicate.GroupSetting {
	return predicate.GroupSetting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupSetting) predicate.GroupSetting {
	return predicate.GroupSetting(sql.NotPredicates(p))
}
