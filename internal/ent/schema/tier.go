package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Tier holds the schema definition for the Tier entity
type Tier struct {
	ent.Schema
}

// Fields of the Tier
func (Tier) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the tier").
			NotEmpty(),
		field.String("description").
			Comment("a description of the tier").
			Optional(),
		field.String("organization_id").
			Comment("the organization that owns the tier").
			Optional(),
	}
}

// Mixin of the Tier
func (Tier) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{ // empty org means Datum system Tier
			Ref:        "tiers",
			Optional:   true,
			AllowWhere: true,
		},
	}
}

// Edges of the Tier
func (Tier) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges go here
	}
}

func (Tier) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookTier(),
	}
}

// Indexes of the Tier
func (Tier) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "owner_id").
			Unique().
			Annotations(
				entsql.IndexWhere("deleted_at is NULL"),
			),
	}
}

// Annotations of the Tier
func (Tier) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			IDField:         "OwnerID",
			NillableIDField: true,
		},
	}
}

// Interceptors of the Tier
func (Tier) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorTier(),
	}
}

// Policy of the Tier
func (Tier) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.TierMutationRuleFunc(func(ctx context.Context, m *generated.TierMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.TierQueryRuleFunc(func(ctx context.Context, q *generated.TierQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
