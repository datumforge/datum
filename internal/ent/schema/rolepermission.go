package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// RolePermission holds the schema definition for the RolePermission entity.
type RolePermission struct {
	ent.Schema
}

// Fields of the RolePermission.
func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_id"),
		field.String("permission_id"),
	}
}

// Edges of the RolePermission.
func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
		edge.To("permission", Permission.Type).
			Unique().
			Required().
			Field("permission_id"),
	}
}

// Mixin of the RolePermission
func (RolePermission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the RolePermission
func (RolePermission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
