package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Entitlement holds the schema definition for the Entitlement entity.
type Entitlement struct {
	ent.Schema
}

// Fields of the Subscription.
func (Entitlement) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("tier").Values("free", "pro", "enterprise").Default("free"),
		field.String("stripe_customer_id").Optional(),
		field.String("stripe_subscription_id").Optional(),
		field.Time("expires_at").Optional(),
		field.Bool("cancelled").Default(false),
	}
}

// Edges of the Subscription
func (Entitlement) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the Entitlement
func (Entitlement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll, entgql.SkipMutationUpdateInput),
	}
}

// Mixin of the RefreshToken
func (Entitlement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}
