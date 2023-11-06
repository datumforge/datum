package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Integration maps configured integrations (github, slack, etc.) to organizations
type Integration struct {
	ent.Schema
}

// Fields of the Integration
func (Integration) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
			),
		field.String("kind").
			Immutable().
			Annotations(
				entgql.OrderField("kind"),
			),
		field.String("description").Optional(),
		field.String("secret_name").Immutable(),
	}
}

// Edges of the Integration
func (Integration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Organization.Type).Ref("integrations").Unique(),
	}
}

// Annotations of the Integration
func (Integration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Integration
func (Integration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
