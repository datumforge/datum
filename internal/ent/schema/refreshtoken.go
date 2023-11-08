package schema

import (
	"time"

	"entgo.io/ent"
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
		// TO DO figure out why JSON doesn't work with oas / ogent
		field.Text("scopes").Optional(),
		field.Text("nonce").
			NotEmpty(),
		field.Text("claims_user_id").
			NotEmpty(),
		field.Text("claims_username").
			NotEmpty(),
		field.Text("claims_email").
			NotEmpty(),
		field.Bool("claims_email_verified"),
		// TO DO figure out why JSON doesn't work with oas / ogent
		field.Text("claims_groups").Optional(),
		//			Annotations(entoas.Schema(&ogen.Schema(AsArray))),
		field.Text("claims_preferred_username"),
		field.Text("connector_id").
			NotEmpty(),
		// TO DO figure out why Bytes doesn't work with oas / ogent
		field.Text("connector_data").Nillable().Optional(),
		field.Text("token"),
		field.Text("obsolete_token"),
		field.Time("last_used").
			Default(time.Now),
	}
}

// Edges of the RefreshToken
func (RefreshToken) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the RefreshToken
func (RefreshToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.IDMixin{},
	}
}
