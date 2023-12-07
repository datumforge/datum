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
)

// AccessToken holds the schema definition for the AccessToken entity.
type AccessToken struct {
	ent.Schema
}

// Fields of the AccessToken
func (AccessToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("access_token").Sensitive().
			Unique().
			Immutable(),
		field.Time("expires_at").
			Default(func() time.Time { return time.Now().Add(time.Hour * 24 * 7) }), // nolint: gomnd
		field.Time("issued_at").
			Default(time.Now()),
		field.Time("last_used_at").
			UpdateDefault(time.Now).
			Optional().
			Nillable(),
		field.String("organization_id").
			Comment("organization ID of the organization the user is accessing"),
		field.String("user_id").
			Comment("the user the session is associated with"),
	}
}

// Edges of the AccessToken
func (AccessToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("access_token").
			Field("user_id").
			Required().
			Unique(),
	}
}

// Indexes of the AccessToken
func (AccessToken) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("access_token"),
	}
}

// Mixin of the AccessToken
func (AccessToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the AccessToken
func (AccessToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
