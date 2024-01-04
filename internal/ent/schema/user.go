package schema

import (
	"net/mail"
	"net/url"
	"strings"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/entx"
)

const (
	urlMaxLen  = 255
	nameMaxLen = 64
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		mixin.IDMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.String("email").
			Unique().
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}),
		field.String("first_name").
			NotEmpty().
			MaxLen(nameMaxLen).
			Annotations(
				entgql.OrderField("first_name"),
			),
		field.String("last_name").
			NotEmpty().
			MaxLen(nameMaxLen).
			Annotations(
				entgql.OrderField("last_name"),
			),
		field.String("display_name").
			Comment("The user's displayed 'friendly' name").
			MaxLen(nameMaxLen).
			NotEmpty().
			Default("").
			Annotations(
				entgql.OrderField("display_name"),
			).
			Validate(
				func(s string) error {
					if strings.Contains(s, " ") {
						return ErrContainsSpaces
					}
					return nil
				},
			),
		field.String("avatar_remote_url").
			Comment("URL of the user's remote avatar").
			MaxLen(urlMaxLen).
			Validate(func(s string) error {
				_, err := url.Parse(s)
				return err
			}).
			Optional().
			Nillable(),
		field.String("avatar_local_file").
			Comment("The user's local avatar file").
			MaxLen(urlMaxLen).
			Optional().
			Nillable(),
		field.Time("avatar_updated_at").
			Comment("The time the user's (local) avatar was last updated").
			UpdateDefault(time.Now).
			Optional().
			Nillable(),
		field.Time("last_seen").
			Comment("the time the user was last seen").
			UpdateDefault(time.Now).
			Optional().
			Nillable(),
		field.String("password").
			Comment("user password hash").
			Nillable().
			Sensitive().
			Optional(),
		field.String("sub").
			Comment("the Subject of the user JWT").
			Unique().
			Optional(),
		field.Bool("oauth").
			Comment("whether the user uses oauth for login or not").
			Default(false),
	}
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organizations", Organization.Type),
		edge.To("sessions", Session.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
		edge.From("groups", Group.Type).
			Ref("users"),
		edge.To("personal_access_tokens", PersonalAccessToken.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
		edge.To("setting", UserSetting.Type).
			Required().
			Unique().
			Annotations(entx.CascadeAnnotationField("User")),
		edge.To("email_verification_tokens", EmailVerificationToken.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
		edge.To("reset_tokens", PasswordResetToken.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Policy of the User
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfContextHasPrivacyTokenOfType(&token.EmailSignUpToken{}),
					rule.DenyIfNoViewer(),
					rule.AllowIfSelf(),
					rule.AllowIfOwnedByViewer(),
					// rule.AllowIfAdmin(), // TODO: this currently is always skipped, setup admin policy to get users
					privacy.AlwaysDenyRule(),
				},
				// the user hook has update operations on user create so we need to allow email token sign up for update
				// operations as well
				ent.OpCreate|ent.OpUpdateOne,
			),
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.DenyIfNoViewer(),
					rule.AllowIfSelf(),
					rule.AllowIfOwnedByViewer(),
					// rule.AllowIfAdmin(), // TODO: this currently is always skipped, setup admin policy to get users
					privacy.AlwaysDenyRule(),
				},
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
		Query: privacy.QueryPolicy{
			// Required to verify tokens
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.VerifyToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.VerifyToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereHasEmailVerificationTokensWith(emailverificationtoken.Token(actualToken.VerifyToken))
				},
			),
			// Forgot password path
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.ForgotPasswordToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.ForgotPasswordToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereEmail(entql.StringEQ(actualToken.Email))
				},
			),
			// Reset password path
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.PasswordResetToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.PasswordResetToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereHasResetTokensWith(passwordresettoken.Token(actualToken.ResetToken))
				},
			),
			// Login path
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.LoginToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.LoginToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereEmail(entql.StringEQ(actualToken.Email))
				},
			),
			// Register and Resend path
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.EmailSignUpToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.EmailSignUpToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereEmail(entql.StringEQ(actualToken.Email))
				},
			),
			// Refresh path
			rule.AllowAfterApplyingPrivacyTokenFilter(
				&token.RefreshToken{},
				func(t token.PrivacyToken, filter privacy.Filter) {
					actualToken := t.(*token.RefreshToken)
					userFilter := filter.(*generated.UserFilter)
					userFilter.WhereSub(entql.StringEQ(actualToken.Subject))
				},
			),
			rule.DenyIfNoSubject(),
			rule.AllowIfSelf(),
			// rule.AllowIfAdmin(), // TODO: this currently is always skipped, setup admin policy to get users
			privacy.AlwaysDenyRule(),
		},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookUser(),
	}
}
