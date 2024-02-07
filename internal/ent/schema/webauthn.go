package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/entx"
)

// Webauthn holds the schema definition for the Webauthn entity
type Webauthn struct {
	ent.Schema
}

// Fields of the Webauthn
func (Webauthn) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("name of the credential"),
		field.String("user_id").
			Comment("the userid").
			Unique().
			NotEmpty(),
		field.String("credential_id").
			Comment("the userid ").
			Unique().
			Optional(),
		field.Bytes("public_key").
			Optional().
			Comment("pub key"),
		field.String("attestation_type").
			Optional().
			Comment("type"),
		field.String("aaguid").
			Optional().
			Comment("aaguid"),
		field.Int("sign_count").
			Optional().
			Comment("sign count"),
		field.Strings("transports").
			Optional().
			Comment("transport"),
		field.Strings("flags").
			Optional().
			Comment("flags"),
		field.Strings("authenticator").
			Optional().
			Comment("auth"),
		field.Bool("backup_eligible").
			Optional().
			Comment("backup?"),
		field.Bool("backup_state").
			Optional().
			Comment("backup?"),
	}
}

// Edges of the Webauthn
func (Webauthn) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Webauthn
func (Webauthn) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		UserOwnedMixin{
			Ref: "webauthn",
		},
	}
}

// Annotations of the Webauthn
func (Webauthn) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
		entx.SchemaGenSkip(true),
	}
}
