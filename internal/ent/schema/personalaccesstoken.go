package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// PersonalAccessToken holds the schema definition for the PersonalAccessToken entity.
type PersonalAccessToken struct {
	ent.Schema
}

// Fields of the PersonalAccessToken.
func (PersonalAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("user_id"),
		field.String("token").Sensitive(),
		field.JSON("abilities", []string{}).
			Optional(),
		field.Time("expiration_at"),
		field.Time("last_used_at").
			Optional().
			Nillable(),
	}
}

// Edges of the PersonalAccessToken.
func (PersonalAccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("personal_access_tokens").
			Unique().
			Required().
			Field("user_id"),
	}
}

func (PersonalAccessToken) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("token"),
	}
}

// Mixin of the RefreshToken
func (PersonalAccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the Organization
func (PersonalAccessToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
