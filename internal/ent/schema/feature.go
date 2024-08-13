package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/pkg/gqlplugin/searchgen"
)

// Feature defines the feature schema
type Feature struct {
	ent.Schema
}

// Fields returns feature fields
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the unique name of the feature").
			NotEmpty().
			Annotations(
				searchgen.FieldSearchable(),
			).
			Immutable(),
		field.String("display_name").
			Comment("the displayed 'friendly' name of the feature").
			Annotations(
				searchgen.FieldSearchable(),
			).
			Optional(),
		field.Bool("enabled").
			Comment("enabled features are available for use").
			Default(false),
		field.String("description").
			Comment("a description of the feature").
			Nillable().
			Optional(),
		field.JSON("metadata", map[string]interface{}{}).
			Comment("metadata for the feature").
			Optional(),
	}
}

// Edges returns feature edges
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plans", EntitlementPlan.Type).
			Through("features", EntitlementPlanFeature.Type),
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
		OrgOwnerMixin{
			Ref: "features",
		},
	}
}

func (Feature) Indexes() []ent.Index {
	return []ent.Index{
		// names should be unique, but ignore deleted names
		index.Fields("name", "owner_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Feature
func (Feature) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			NillableIDField: true,
			OrgOwnedField:   true,
			IDField:         "OwnerID",
		},
		entx.SchemaSearchable(true),
	}
}

// Hooks of the Feature
func (Feature) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookFeature(),
	}
}

// Interceptors of the Feature
func (Feature) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{}
}

// Policy of the Feature
func (Feature) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.FeatureMutationRuleFunc(func(ctx context.Context, m *generated.FeatureMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),

			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.FeatureQueryRuleFunc(func(ctx context.Context, q *generated.FeatureQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
