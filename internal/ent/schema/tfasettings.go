package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// TFASettings holds the schema definition for the TFASettings entity
type TFASettings struct {
	ent.Schema
}

// Fields of the TFASettings
func (TFASettings) Fields() []ent.Field {
	return []ent.Field{
		field.String("tfa_secret").
			Comment("TFA secret for the user").
			Sensitive().
			Annotations(entoas.Skip(true)).
			Optional().
			Nillable(),
		field.Bool("verified").
			Comment("specifies if the TFA device has been verified").
			Annotations(
				entoas.Skip(true),
				entgql.Skip(entgql.SkipAll),
			).
			Default(false),
		field.Strings("recovery_codes").
			Comment("recovery codes for 2fa").
			Annotations(entoas.Skip(true)).
			Optional(),
		field.Bool("phone_otp_allowed").
			Comment("specifies a user may complete authentication by verifying an OTP code delivered through SMS").
			Optional().
			Default(false),
		field.Bool("email_otp_allowed").
			Comment("specifies a user may complete authentication by verifying an OTP code delivered through email").
			Optional().
			Default(false),
		field.Bool("totp_allowed").
			Comment("specifies a user may complete authentication by verifying a TOTP code delivered through an authenticator app").
			Optional().
			Default(false),
	}
}

// Mixin of the TFASettings
func (TFASettings) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		UserOwnedMixin{
			Ref:         "tfa_settings",
			Optional:    true,
			AllowUpdate: false,
			// SoftDeleteIndex: true,
		},
	}
}

// Edges of the TFASettings
func (TFASettings) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges go here
	}
}

// Hooks of the TFASettings
func (TFASettings) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookTFA(),
	}
}

// Annotations of the TFASettings
func (TFASettings) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
