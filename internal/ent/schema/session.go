package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Session holds authentication sessions. They can either be first-party web auth sessions or OAuth sessions. Sessions should persist in the database for some time duration after expiration, but with the "disabled" boolean set to true.
type Session struct {
	ent.Schema
}

// Fields of the Session
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("session_token").
			Comment("token is a string token issued to users that has a limited lifetime").
			Unique().
			Immutable(),
		field.Time("issued_at").
			UpdateDefault(time.Now),
		field.Time("expires_at"),
		field.String("organization_id").
			Comment("organization ID of the organization the user is accessing"),
		field.String("user_id").
			Comment("the user the session is associated with"),
	}
}

// Indexes of the Session
func (Session) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("session_token").
			Unique(),
	}
}

// Edges of the Session
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("sessions").
			Field("user_id").
			Unique().
			Required().
			Comment("Sessions belong to users"),
	}
}

// Annotations of the Session
func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Session
func (Session) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Hooks of the Session
func (Session) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookSession(),
	}
}
