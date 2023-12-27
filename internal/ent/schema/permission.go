package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("action").
			Unique(),
		field.String("description").
			Optional(),
		field.Bool("is_disabled").
			Default(false),
	}
}

// Edges of the Permission
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("permissions").
			Through("role_permission", RolePermission.Type),
	}
}

// Mixin of the Permission
func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the Permission
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
