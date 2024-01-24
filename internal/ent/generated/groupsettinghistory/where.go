// Code generated by ent, DO NOT EDIT.

package groupsettinghistory

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/flume/enthistory"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContainsFold(FieldID, id))
}

// HistoryTime applies equality check predicate on the "history_time" field. It's identical to HistoryTimeEQ.
func HistoryTime(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// Ref applies equality check predicate on the "ref" field. It's identical to RefEQ.
func Ref(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldRef, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldDeletedBy, v))
}

// SyncToSlack applies equality check predicate on the "sync_to_slack" field. It's identical to SyncToSlackEQ.
func SyncToSlack(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldSyncToSlack, v))
}

// SyncToGithub applies equality check predicate on the "sync_to_github" field. It's identical to SyncToGithubEQ.
func SyncToGithub(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldSyncToGithub, v))
}

// HistoryTimeEQ applies the EQ predicate on the "history_time" field.
func HistoryTimeEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldHistoryTime, v))
}

// HistoryTimeNEQ applies the NEQ predicate on the "history_time" field.
func HistoryTimeNEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldHistoryTime, v))
}

// HistoryTimeIn applies the In predicate on the "history_time" field.
func HistoryTimeIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldHistoryTime, vs...))
}

// HistoryTimeNotIn applies the NotIn predicate on the "history_time" field.
func HistoryTimeNotIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldHistoryTime, vs...))
}

// HistoryTimeGT applies the GT predicate on the "history_time" field.
func HistoryTimeGT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldHistoryTime, v))
}

// HistoryTimeGTE applies the GTE predicate on the "history_time" field.
func HistoryTimeGTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldHistoryTime, v))
}

// HistoryTimeLT applies the LT predicate on the "history_time" field.
func HistoryTimeLT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldHistoryTime, v))
}

// HistoryTimeLTE applies the LTE predicate on the "history_time" field.
func HistoryTimeLTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldHistoryTime, v))
}

// OperationEQ applies the EQ predicate on the "operation" field.
func OperationEQ(v enthistory.OpType) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldOperation, v))
}

// OperationNEQ applies the NEQ predicate on the "operation" field.
func OperationNEQ(v enthistory.OpType) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldOperation, v))
}

// OperationIn applies the In predicate on the "operation" field.
func OperationIn(vs ...enthistory.OpType) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldOperation, vs...))
}

// OperationNotIn applies the NotIn predicate on the "operation" field.
func OperationNotIn(vs ...enthistory.OpType) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldOperation, vs...))
}

// RefEQ applies the EQ predicate on the "ref" field.
func RefEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldRef, v))
}

// RefNEQ applies the NEQ predicate on the "ref" field.
func RefNEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldRef, v))
}

// RefIn applies the In predicate on the "ref" field.
func RefIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldRef, vs...))
}

// RefNotIn applies the NotIn predicate on the "ref" field.
func RefNotIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldRef, vs...))
}

// RefGT applies the GT predicate on the "ref" field.
func RefGT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldRef, v))
}

// RefGTE applies the GTE predicate on the "ref" field.
func RefGTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldRef, v))
}

// RefLT applies the LT predicate on the "ref" field.
func RefLT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldRef, v))
}

// RefLTE applies the LTE predicate on the "ref" field.
func RefLTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldRef, v))
}

// RefContains applies the Contains predicate on the "ref" field.
func RefContains(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContains(FieldRef, v))
}

// RefHasPrefix applies the HasPrefix predicate on the "ref" field.
func RefHasPrefix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasPrefix(FieldRef, v))
}

// RefHasSuffix applies the HasSuffix predicate on the "ref" field.
func RefHasSuffix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasSuffix(FieldRef, v))
}

// RefIsNil applies the IsNil predicate on the "ref" field.
func RefIsNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIsNull(FieldRef))
}

// RefNotNil applies the NotNil predicate on the "ref" field.
func RefNotNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotNull(FieldRef))
}

// RefEqualFold applies the EqualFold predicate on the "ref" field.
func RefEqualFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEqualFold(FieldRef, v))
}

// RefContainsFold applies the ContainsFold predicate on the "ref" field.
func RefContainsFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContainsFold(FieldRef, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldContainsFold(FieldDeletedBy, v))
}

// VisibilityEQ applies the EQ predicate on the "visibility" field.
func VisibilityEQ(v enums.Visibility) predicate.GroupSettingHistory {
	vc := v
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldVisibility, vc))
}

// VisibilityNEQ applies the NEQ predicate on the "visibility" field.
func VisibilityNEQ(v enums.Visibility) predicate.GroupSettingHistory {
	vc := v
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldVisibility, vc))
}

// VisibilityIn applies the In predicate on the "visibility" field.
func VisibilityIn(vs ...enums.Visibility) predicate.GroupSettingHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSettingHistory(sql.FieldIn(FieldVisibility, v...))
}

// VisibilityNotIn applies the NotIn predicate on the "visibility" field.
func VisibilityNotIn(vs ...enums.Visibility) predicate.GroupSettingHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldVisibility, v...))
}

// JoinPolicyEQ applies the EQ predicate on the "join_policy" field.
func JoinPolicyEQ(v enums.JoinPolicy) predicate.GroupSettingHistory {
	vc := v
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldJoinPolicy, vc))
}

// JoinPolicyNEQ applies the NEQ predicate on the "join_policy" field.
func JoinPolicyNEQ(v enums.JoinPolicy) predicate.GroupSettingHistory {
	vc := v
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldJoinPolicy, vc))
}

// JoinPolicyIn applies the In predicate on the "join_policy" field.
func JoinPolicyIn(vs ...enums.JoinPolicy) predicate.GroupSettingHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSettingHistory(sql.FieldIn(FieldJoinPolicy, v...))
}

// JoinPolicyNotIn applies the NotIn predicate on the "join_policy" field.
func JoinPolicyNotIn(vs ...enums.JoinPolicy) predicate.GroupSettingHistory {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupSettingHistory(sql.FieldNotIn(FieldJoinPolicy, v...))
}

// SyncToSlackEQ applies the EQ predicate on the "sync_to_slack" field.
func SyncToSlackEQ(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldSyncToSlack, v))
}

// SyncToSlackNEQ applies the NEQ predicate on the "sync_to_slack" field.
func SyncToSlackNEQ(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldSyncToSlack, v))
}

// SyncToGithubEQ applies the EQ predicate on the "sync_to_github" field.
func SyncToGithubEQ(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldEQ(FieldSyncToGithub, v))
}

// SyncToGithubNEQ applies the NEQ predicate on the "sync_to_github" field.
func SyncToGithubNEQ(v bool) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.FieldNEQ(FieldSyncToGithub, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupSettingHistory) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupSettingHistory) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupSettingHistory) predicate.GroupSettingHistory {
	return predicate.GroupSettingHistory(sql.NotPredicates(p))
}
