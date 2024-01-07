package schema

import (
	"context"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
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
			// Privacy willl be always allow, but interceptors will filter the queries
			privacy.AlwaysAllowRule(),
		},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookUser(),
	}
}

// Interceptors of the User.
func (d User) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseUser(func(ctx context.Context, q *generated.UserQuery) error {
			// skip filters on non-authorized user
			err, allow := privacy.DecisionFromContext(ctx)
			if err == nil && allow {
				return nil
			}

			v := viewer.FromContext(ctx)
			if v == nil {
				t := token.EmailSignUpTokenFromContext(ctx)
				if t != nil {
					q.Where(user.Email(t.GetEmail()))
					return nil
				}

				vt := token.VerifyTokenFromContext(ctx)
				if vt != nil {
					q.Filter().WhereHasEmailVerificationTokensWith(emailverificationtoken.Token(vt.VerifyToken))
					return nil
				}

				rt := token.PasswordResetTokenFromContext(ctx)
				if rt != nil {
					q.Filter().WhereHasResetTokensWith(passwordresettoken.Token(rt.ResetToken))
					return nil
				}

				fp := token.ForgotPasswordTokenFromContext(ctx)
				if fp != nil {
					q.Where(user.Email(fp.Email))
					return nil
				}

				ref := token.RefreshTokenFromContext(ctx)
				if ref != nil {
					q.Where(user.Sub(ref.Subject))
					return nil
				}

				// block request
				return privacy.Denyf("anonymous viewer with no valid token")
			}

			viewerID, exists := v.GetID()
			if exists {
				q.Where(user.ID(viewerID))
				return nil
			}

			return nil
		}),
	}
}
