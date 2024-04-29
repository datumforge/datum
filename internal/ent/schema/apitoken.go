package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/enthistory"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/pkg/keygen"
)

// APIToken holds the schema definition for the APIToken entity.
type APIToken struct {
	ent.Schema
}

// Fields of the APIToken
func (APIToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name associated with the token").
			NotEmpty(),
		field.String("organization_id").
			Comment("the organization the token is associated with").
			Immutable().
			NotEmpty(),
		field.String("token").
			Unique().
			Immutable().
			Annotations(
				entgql.Skip(^entgql.SkipType),
				entoas.Skip(true),
			).
			DefaultFunc(func() string {
				token := keygen.PrefixedSecret("dtma") // datum api token prefix
				return token
			}),
		field.Time("expires_at").
			Comment("when the token expires").
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			).
			Nillable(),
		field.String("description").
			Comment("a description of the token's purpose").
			Optional().
			Nillable().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.JSON("scopes", []string{}).
			Optional(),
		field.Time("last_used_at").
			Optional().
			Nillable(),
	}
}

// Edges of the APIToken
func (APIToken) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Indexes of the APIToken
func (APIToken) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("token"),
	}
}

// Mixin of the APIToken
func (APIToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
		OrgOwnerMixin{
			Ref: "api_tokens",
		},
	}
}

// Annotations of the APIToken
func (APIToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		enthistory.Annotations{
			Exclude: true,
		},
	}
}

// Hooks of the APIToken
func (APIToken) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookCreateAPIToken(),
		hooks.HookUpdateAPIToken(),
	}
}

// Interceptors of the APIToken
func (APIToken) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		interceptors.InterceptorAPIToken(),
	}
}

// Policy of the APIToken
func (APIToken) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoSubject(),
			rule.AllowMutationAfterApplyingOwnerFilter(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			rule.AllowIfOwnedByViewer(),
			privacy.AlwaysDenyRule(),
		},
	}
}
