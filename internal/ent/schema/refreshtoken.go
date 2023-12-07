package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// RefreshToken holds the schema definition for the RefreshToken entity
type RefreshToken struct {
	ent.Schema
}

// Fields of the RefreshToken
func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("refresh_token").Sensitive().
			Unique().
			Immutable(),
		field.Time("expires_at").
			Default(func() time.Time { return time.Now().Add(time.Hour * 24 * 7) }), // nolint: gomnd
		field.Time("issued_at").
			Default(time.Now()),
		field.String("organization_id").
			Comment("organization ID of the organization the user is accessing"),
		field.String("user_id").
			Comment("the user the session is associated with"),
	}
}

// Edges of the RefreshToken
func (RefreshToken) Edges() []ent.Edge {
	return nil
}

// Mixin of the RefreshToken
func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}

// Annotations of the RefreshToken
func (RefreshToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
