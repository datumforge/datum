package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Feature defines the feature schema
type Feature struct {
	ent.Schema
}

// Fields returns feature fields
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Immutable().
			Unique(),
		field.Bool("global").
			Default(true),
		field.Bool("enabled").
			Default(false),
		field.String("description").
			Nillable().
			Optional(),
	}
}

// Edges returns feature edges
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("features"),
		edge.From("groups", Group.Type).
			Ref("features"),
		edge.From("entitlements", Entitlement.Type).
			Ref("features"),
		edge.From("organizations", Organization.Type).
			Ref("features"),
		edge.To("events", Event.Type),
	}
}

// Mixin of the Feature
func (Feature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
		emixin.TagMixin{},
	}
}

// Annotations of the Feature
func (Feature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
