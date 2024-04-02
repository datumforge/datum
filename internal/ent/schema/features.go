package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Feature holds the schema definition for the Feature entity
type Feature struct {
	ent.Schema
}

// Fields of the Feature
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the feature").
			NotEmpty(),
		field.String("description").
			Comment("a description of the feature").
			Optional(),
		field.Strings("tiers").
			Comment("the tiers that own the feature").
			Optional(),
	}
}

// Mixin of the Feature
func (Feature) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Edges of the Feature
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tier", Tier.Type),
	}
}

func (Feature) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Indexes of the Feature
func (Feature) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "tier_id").
			Unique().
			Annotations(
				entsql.IndexWhere("deleted_at is NULL"),
			),
	}
}

// Annotations of the Feature
func (Feature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Interceptors of the Feature
func (Feature) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{}
}

// Policy of the Feature
func (Feature) Policy() ent.Policy {
	return privacy.Policy{
		// Mutation: privacy.MutationPolicy{
		// 	privacy.FeatureMutationRuleFunc(func(ctx context.Context, m *generated.FeatureMutation) error {
		// 		return m.CheckAccessForEdit(ctx)
		// 	}),
		// 	privacy.AlwaysDenyRule(),
		// },
		// Query: privacy.QueryPolicy{
		// 	privacy.FeatureQueryRuleFunc(func(ctx context.Context, q *generated.FeatureQuery) error {
		// 		return q.CheckAccess(ctx)
		// 	}),
		// 	privacy.AlwaysDenyRule(),
		// },
	}
}
