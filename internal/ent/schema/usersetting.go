package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/usersetting"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// UserSetting holds the schema definition for the User entity.
type UserSetting struct {
	ent.Schema
}

// Mixin of the UserSetting
func (UserSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Fields of the UserSetting
func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Optional(),
		field.Bool("locked").
			Comment("user account is locked if unconfirmed or explicitly locked").
			Default(false),
		field.Time("silenced_at").
			Comment("The time notifications regarding the user were silenced").
			Optional().
			Nillable(),
		field.Time("suspended_at").
			Comment("The time the user was suspended").
			Optional().
			Nillable().
			Annotations(entoas.Annotation{ReadOnly: true}),
		field.Enum("status").
			GoType(enums.UserStatus("")).
			Default(string(enums.Active)),
		field.Bool("email_confirmed").Default(false).
			Annotations(entoas.Annotation{ReadOnly: true}),
		field.JSON("tags", []string{}).
			Comment("tags associated with the user").
			Default([]string{}),
		field.Bool("is_webauthn_allowed").
			Comment("specifies a user may complete authentication by verifying a WebAuthn capable device").
			Optional().
			Default(false),
		field.Bool("is_tfa_enabled").
			Comment("whether the user has two factor authentication enabled").
			Optional().
			Default(false),
		field.String("phone_number").
			Comment("phone number associated with the account, used 2factor SMS authentication").
			Optional().
			Annotations(
				// skip until SMS 2fa feature is implemented
				entoas.Skip(true),
				entgql.Skip(entgql.SkipAll),
			).
			Nillable(),
	}
}

// Edges of the UserSetting
func (UserSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("setting").Unique().Field("user_id"),
		edge.To("default_org", Organization.Type).
			Unique().
			Comment("organization to load on user login"),
	}
}

// Annotations of the UserSetting
func (UserSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Hooks of the UserSetting.
func (UserSetting) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookUserSetting(),
	}
}

// Interceptors of the UserSetting.
func (d UserSetting) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseUserSetting(func(ctx context.Context, q *generated.UserSettingQuery) error {
			// Filter query based on viewer context
			v := viewer.FromContext(ctx)
			if v != nil {
				viewerID, exists := v.GetID()
				if exists {
					q.Where(usersetting.UserID(viewerID))
					return nil
				}
			}

			return nil
		}),
	}
}

func (UserSetting) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfContextHasPrivacyTokenOfType(&token.VerifyToken{}),
					rule.DenyIfNoViewer(),
					rule.AllowIfSelf(),
					privacy.AlwaysDenyRule(),
				},
				// only resolvers exist for update operations
				ent.OpUpdateOne|ent.OpUpdate,
			),
		},
		Query: privacy.QueryPolicy{
			// Privacy will be always allow, but interceptors will filter the queries
			privacy.AlwaysAllowRule(),
		},
	}
}
