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
)

// PasswordResetToken holds the schema definition for the PasswordResetToken entity
type PasswordResetToken struct {
	ent.Schema
}

// Fields of the PasswordResetToken
func (PasswordResetToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Comment("the reset token sent to the user via email which should only be provided to the /forgot-password endpoint + handler").
			Unique().
			NotEmpty(),
		field.Time("ttl").
			Comment("the ttl of the reset token which defaults to 15 minutes").
			Nillable(),
		field.String("email").
			Comment("the email used as input to generate the reset token; this is used to verify that the token when regenerated within the server matches the token emailed").
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

// Edges of the PasswordResetToken
func (PasswordResetToken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("reset_tokens").
			Required().
			Unique(),
	}
}

// Mixin of the PasswordResetToken
func (PasswordResetToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Indexes of the PasswordResetToken
func (PasswordResetToken) Indexes() []ent.Index {
	return []ent.Index{
		// PasswordResetTokens should be unique, but ignore deleted PasswordResetTokens
		index.Fields("token").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the PasswordResetToken
func (PasswordResetToken) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Skip(entgql.SkipAll),
	}
}

// Hooks of the PasswordResetToken
func (PasswordResetToken) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookPasswordResetToken(),
	}
}
