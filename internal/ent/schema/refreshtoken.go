package schema

import (
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
		field.String("client_id").
			NotEmpty(),
		//		field.JSON("scopes", []string{}).
		//			Optional(),
		//		field.Text("nonce").
		//			NotEmpty(),
		//		field.Text("claims_user_id").
		//			NotEmpty(),
		//		field.Text("claims_username").
		//			NotEmpty(),
		//		field.Text("claims_email").
		//			NotEmpty(),
		//		field.Bool("claims_email_verified"),
		//		field.JSON("claims_groups", []string{}).
		//			Optional(),
		//		field.Text("claims_preferred_username"),
		//		field.Text("connector_id").
		//			NotEmpty(),
		//		field.Bytes("connector_data").
		//			Nillable().
		//			Optional(),
		//		field.Text("token"),
		//		field.Text("obsolete_token"),
		//		field.Time("last_used").
		//			Default(time.Now),
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
