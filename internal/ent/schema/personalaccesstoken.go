package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/keygen"
)

// PersonalAccessToken holds the schema definition for the PersonalAccessToken entity.
type PersonalAccessToken struct {
	ent.Schema
}

// Fields of the PersonalAccessToken
func (PersonalAccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("token").Sensitive().
			Unique().
			Immutable().
			DefaultFunc(func() string {
				token := keygen.Secret()
				return token
			}),
		field.JSON("abilities", []string{}).
			Optional(),
		field.Time("expiration_at"),
		field.String("description").Default("").Annotations(
			entgql.Skip(entgql.SkipWhereInput),
		),
		field.Time("last_used_at").
			UpdateDefault(time.Now).
			Optional().
			Nillable(),
	}
}

// Edges of the PersonalAccessToken
func (PersonalAccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("personal_access_tokens").
			Required().
			Unique(),
	}
}

// Indexes of the PersonalAccessToken
func (PersonalAccessToken) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("token"),
	}
}

// Mixin of the PersonalAccessToken
func (PersonalAccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the PersonalAccessToken
func (PersonalAccessToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
