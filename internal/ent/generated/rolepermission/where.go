// Code generated by ent, DO NOT EDIT.

package rolepermission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedBy, v))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldRoleID, vs...))
}

// RoleIDGT applies the GT predicate on the "role_id" field.
func RoleIDGT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldRoleID, v))
}

// RoleIDGTE applies the GTE predicate on the "role_id" field.
func RoleIDGTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldRoleID, v))
}

// RoleIDLT applies the LT predicate on the "role_id" field.
func RoleIDLT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldRoleID, v))
}

// RoleIDLTE applies the LTE predicate on the "role_id" field.
func RoleIDLTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldRoleID, v))
}

// RoleIDContains applies the Contains predicate on the "role_id" field.
func RoleIDContains(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContains(FieldRoleID, v))
}

// RoleIDHasPrefix applies the HasPrefix predicate on the "role_id" field.
func RoleIDHasPrefix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasPrefix(FieldRoleID, v))
}

// RoleIDHasSuffix applies the HasSuffix predicate on the "role_id" field.
func RoleIDHasSuffix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasSuffix(FieldRoleID, v))
}

// RoleIDEqualFold applies the EqualFold predicate on the "role_id" field.
func RoleIDEqualFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEqualFold(FieldRoleID, v))
}

// RoleIDContainsFold applies the ContainsFold predicate on the "role_id" field.
func RoleIDContainsFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContainsFold(FieldRoleID, v))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldNotIn(FieldPermissionID, vs...))
}

// PermissionIDGT applies the GT predicate on the "permission_id" field.
func PermissionIDGT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGT(FieldPermissionID, v))
}

// PermissionIDGTE applies the GTE predicate on the "permission_id" field.
func PermissionIDGTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldGTE(FieldPermissionID, v))
}

// PermissionIDLT applies the LT predicate on the "permission_id" field.
func PermissionIDLT(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLT(FieldPermissionID, v))
}

// PermissionIDLTE applies the LTE predicate on the "permission_id" field.
func PermissionIDLTE(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldLTE(FieldPermissionID, v))
}

// PermissionIDContains applies the Contains predicate on the "permission_id" field.
func PermissionIDContains(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContains(FieldPermissionID, v))
}

// PermissionIDHasPrefix applies the HasPrefix predicate on the "permission_id" field.
func PermissionIDHasPrefix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasPrefix(FieldPermissionID, v))
}

// PermissionIDHasSuffix applies the HasSuffix predicate on the "permission_id" field.
func PermissionIDHasSuffix(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldHasSuffix(FieldPermissionID, v))
}

// PermissionIDEqualFold applies the EqualFold predicate on the "permission_id" field.
func PermissionIDEqualFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldEqualFold(FieldPermissionID, v))
}

// PermissionIDContainsFold applies the ContainsFold predicate on the "permission_id" field.
func PermissionIDContainsFold(v string) predicate.RolePermission {
	return predicate.RolePermission(sql.FieldContainsFold(FieldPermissionID, v))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoleTable, RoleColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RolePermission
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newRoleStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Role
		step.Edge.Schema = schemaConfig.RolePermission
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PermissionTable, PermissionColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Permission
		step.Edge.Schema = schemaConfig.RolePermission
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.RolePermission {
	return predicate.RolePermission(func(s *sql.Selector) {
		step := newPermissionStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Permission
		step.Edge.Schema = schemaConfig.RolePermission
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RolePermission) predicate.RolePermission {
	return predicate.RolePermission(sql.NotPredicates(p))
}
