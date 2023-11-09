package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Taco holds the schema definition for the Subscription entity.
type Taco struct {
	ent.Schema
}

// Fields of the Subscription.
func (Taco) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("tier").Values("free", "pro", "enterprise").Default("free"),
		field.String("stripe_customer_id").Optional(),
		field.String("stripe_subscription_id").Optional(),
		field.Time("expires_at").Optional(),
		field.Bool("cancelled").Default(false),
	}
}

// Edges of the Taco
func (Taco) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Taco
func (Taco) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
		mixin.AuditMixin{},
	}
}

// Annotations of the Taco
func (Taco) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
