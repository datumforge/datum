package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// Entitlement holds the schema definition for the Entitlement entity.
type Entitlement struct {
	ent.Schema
}

// Fields of the Entitlement.
func (Entitlement) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("tier").Values("free", "pro", "enterprise").Default("free"),
		field.String("external_customer_id").Optional().Comment("used to store references to external systems, e.g. Stripe"),
		field.String("external_subscription_id").Comment("used to store references to external systems, e.g. Stripe").Optional(),
		field.Time("expires_at").Optional(),
		field.Time("upgraded_at").Optional(),
		field.String("upgraded_tier").Comment("the tier the customer upgraded from").Optional(),
		field.Time("downgraded_at").Optional(),
		field.String("downgraded_tier").Comment("the tier the customer downgraded from").Optional(),
		field.Bool("cancelled").Default(false),
	}
}

// Edges of the Entitlement
func (Entitlement) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Organization.Type).Ref("entitlements").Unique(),
	}
}

// Annotations of the Entitlement
func (Entitlement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Entitlement
func (Entitlement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
