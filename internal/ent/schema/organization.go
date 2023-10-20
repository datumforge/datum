package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Organization holds the schema definition for the Organization entity - organizations are the top level tenancy construct in the system
type Organization struct {
	ent.Schema
}

// Fields of the Organization
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
	}
}

// Edges of the Organization
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		// an org can have and belong to many users
		edge.To("memberships", Membership.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("integrations", Integration.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}

// Annotations of the Organization
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Organization
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
