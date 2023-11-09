package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("tier").Values("free", "pro", "enterprise").Default("free"),
		field.String("stripe_customer_id").Optional(),
		field.String("stripe_subscription_id").Optional(),
		field.Time("expires_at").Optional(),
		field.Bool("cancelled").Default(false),
	}
}

// Edges of the Subscription
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the RefreshToken
func (Subscription) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}
