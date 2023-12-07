package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// OAuth2Token holds the schema definition for the OAuth2Token entity
type OAuth2Token struct {
	ent.Schema
}

// Fields of the OAuth2Token
func (OAuth2Token) Fields() []ent.Field {
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
		field.Time("last_used").
			Default(time.Now),
	}
}

// Edges of the OAuth2Token
func (OAuth2Token) Edges() []ent.Edge {
	return nil
}

// Mixin of the OAuth2Token
func (OAuth2Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}

// Annotations of the OAuth2Token
func (OAuth2Token) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
