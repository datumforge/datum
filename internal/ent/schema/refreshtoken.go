package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
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
		field.Text("client_id").
			NotEmpty(),
		field.JSON("scopes", []string{}).
			Optional(),
		field.Text("nonce").
			NotEmpty(),
		field.Text("claims_user_id").
			NotEmpty(),
		field.Text("claims_username").
			NotEmpty(),
		field.Text("claims_email").
			NotEmpty(),
		field.Bool("claims_email_verified"),
		field.JSON("claims_groups", []string{}).
			Optional(),
		field.Text("claims_preferred_username"),
		field.Text("connector_id").
			NotEmpty(),
		field.JSON("connector_data", []string{}).
			Optional(),
		field.Text("token"),
		field.Text("obsolete_token"),
		field.Time("last_used").
			Default(time.Now),
	}
}

// Edges of the RefreshToken
func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("refreshtoken").Unique(),
	}
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
