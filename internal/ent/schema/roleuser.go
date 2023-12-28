package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoleUser holds the schema definition for the RoleUser entity.
type RoleUser struct {
	ent.Schema
}

func (RoleUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "role_id"),
	}
}

// Fields of the RoleUser.
func (RoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id"),
		field.Int("user_id"),
	}
}

// Edges of the RoleUser.
func (RoleUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
