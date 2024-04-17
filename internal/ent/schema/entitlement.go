package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/enthistory"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Entitlement holds the schema definition for the Entitlement entity.
type Entitlement struct {
	ent.Schema
}

// Fields of the Entitlement.
func (Entitlement) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("tier").
			GoType(enums.Tier("")).
			Default(string(enums.Free)),
		field.String("external_customer_id").
			Comment("used to store references to external systems, e.g. Stripe").
			Optional(),
		field.String("external_subscription_id").
			Comment("used to store references to external systems, e.g. Stripe").
			Optional(),
		field.Bool("expires").
			Comment("whether or not the customers entitlement expires - expires_at will show the time").
			Default(false),
		field.Time("expires_at").
			Comment("the time at which a customer's entitlement will expire, e.g. they've cancelled but paid through the end of the month").
			Optional().
			Nillable(),
		field.Bool("cancelled").
			Comment("whether or not the customer has cancelled their entitlement - usually used in conjunction with expires and expires at").
			Default(false),
	}
}

// Edges of the Entitlement
func (Entitlement) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Annotations of the Entitlement
func (Entitlement) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		enthistory.Annotations{
			Exclude: true,
		},
		entsql.Schema("datum"),
	}
}

// Mixin of the Entitlement
func (Entitlement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{
			Ref: "entitlements",
		},
	}
}
