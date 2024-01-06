package schema

import (
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/entx"
)

// EmailVerificationToken holds the schema definition for the EmailVerificationToken entity
type EmailVerificationToken struct {
	ent.Schema
}

// Fields of the EmailVerificationToken
func (EmailVerificationToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Comment("the verification token sent to the user via email which should only be provided to the /verify endpoint + handler").
			Unique().
			NotEmpty(),
		field.Time("ttl").
			Comment("the ttl of the verification token which defaults to 7 days").
			Nillable(),
		field.String("email").
			Comment("the email used as input to generate the verification token; this is used to verify that the token when regenerated within the server matches the token emailed").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			NotEmpty(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Nillable(),
	}
}

// Edges of the EmailVerificationToken
func (EmailVerificationToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("email_verification_tokens").
			Required().
			Unique(),
	}
}

// Mixin of the EmailVerificationToken
func (EmailVerificationToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Indexes of the EmailVerificationToken
func (EmailVerificationToken) Indexes() []ent.Index {
	return []ent.Index{
		// EmailVerificationTokens should be unique, but ignore deleted EmailVerificationTokens
		index.Fields("token").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the EmailVerificationToken
func (EmailVerificationToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
		entx.SchemaGenSkip(true),
	}
}

// Hooks of the EmailVerificationToken
func (EmailVerificationToken) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookEmailVerificationToken(),
	}
}
