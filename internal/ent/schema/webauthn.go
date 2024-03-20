package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"
)

// Webauthn holds the schema definition for the Webauthn entity
type Webauthn struct {
	ent.Schema
}

// Fields of the Webauthn
func (Webauthn) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("credential_id").
			Comment("A probabilistically-unique byte sequence identifying a public key credential source and its authentication assertions").
			Unique().
			Optional(),
		field.Bytes("public_key").
			Comment("The public key portion of a Relying Party-specific credential key pair, generated by an authenticator and returned to a Relying Party at registration time").
			Optional(),
		field.String("attestation_type").
			Comment("The attestation format used (if any) by the authenticator when creating the credential").
			Optional(),
		field.Bytes("aaguid").
			Comment("The AAGUID of the authenticator; AAGUID is defined as an array containing the globally unique identifier of the authenticator model being sought").
			Unique().
			Immutable(),
		field.Int32("sign_count").
			Comment("SignCount -Upon a new login operation, the Relying Party compares the stored signature counter value with the new signCount value returned in the assertions authenticator data. If this new signCount value is less than or equal to the stored value, a cloned authenticator may exist, or the authenticator may be malfunctioning"),
		field.Strings("transports").
			Comment("transport"),
		field.Bool("backup_eligible").
			Comment("Flag backup eligible indicates the credential is able to be backed up and/or sync'd between devices. This should NEVER change").
			Default(false).
			Immutable(),
		field.Bool("backup_state").
			Comment("Flag backup state indicates the credential has been backed up and/or sync'd").
			Default(false),
		field.Bool("user_present").
			Comment("Flag user present indicates the users presence").
			Default(false),
		field.Bool("user_verified").
			Comment("Flag user verified indicates the user performed verification").
			Default(false),
	}
}

// Edges of the Webauthn
func (Webauthn) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Webauthn
func (Webauthn) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		UserOwnedMixin{
			Ref:               "webauthn",
			SkipOASGeneration: true,
		},
	}
}

// Annotations of the Webauthn
func (Webauthn) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
		entx.SchemaGenSkip(true),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}
