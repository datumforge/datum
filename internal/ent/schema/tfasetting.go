package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/enthistory"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// TFASetting holds the schema definition for the TFASetting entity
type TFASetting struct {
	ent.Schema
}

// Fields of the TFASetting
func (TFASetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("tfa_secret").
			Comment("TFA secret for the user").
			Annotations(
				entoas.Skip(true),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			Optional().
			Nillable(),
		field.Bool("verified").
			Comment("specifies if the TFA device has been verified").
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput),
			).
			Default(false),
		field.Strings("recovery_codes").
			Comment("recovery codes for 2fa").
			Annotations(
				entoas.Skip(true),
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			).
			Optional(),
		field.Bool("phone_otp_allowed").
			Comment("specifies a user may complete authentication by verifying an OTP code delivered through SMS").
			Optional().
			Annotations(
				// skip until feature is implemented
				entoas.Skip(true),
				entgql.Skip(entgql.SkipAll),
			).
			Default(false),
		field.Bool("email_otp_allowed").
			Comment("specifies a user may complete authentication by verifying an OTP code delivered through email").
			Optional().
			Annotations(
				// skip until feature is implemented
				entoas.Skip(true),
				entgql.Skip(entgql.SkipAll),
			).
			Default(false),
		field.Bool("totp_allowed").
			Comment("specifies a user may complete authentication by verifying a TOTP code delivered through an authenticator app").
			Optional().
			Default(false),
	}
}

// Mixin of the TFASetting
func (TFASetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		UserOwnedMixin{
			Ref:             "tfa_settings",
			Optional:        true,
			AllowUpdate:     false,
			SoftDeleteIndex: true,
		},
	}
}

// Hooks of the TFASetting
func (TFASetting) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookEnableTFA(), // sets 2fa on user settings and stores recovery codes
	}
}

// Annotations of the TFASetting
func (TFASetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entoas.Skip(true),
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		enthistory.Annotations{
			Exclude: true,
		},
	}
}
