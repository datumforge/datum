package schema

import (
	"context"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
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
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.String("email").
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
			Annotations(entoas.Annotation{ReadOnly: true}).
			Optional().
			Nillable(),
		field.String("password").
			Comment("user password hash").
			Nillable().
			Sensitive().
			Annotations(entoas.Skip(true)).
			Optional(),
		field.String("sub").
			Comment("the Subject of the user JWT").
			Unique().
			Annotations(entoas.Skip(true)).
			Optional(),
		field.Enum("auth_provider").
			GoType(enums.AuthProvider("")).
			Comment("auth provider used to register the account").
			Default(string(enums.AuthProviderCredentials)),
		field.Enum("role").
			Comment("the user's role").
			GoType(enums.Role("")).
			Default(string(enums.RoleMember)).
			Optional(),
	}
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").
			Unique(), // enforce globally unique ids
		index.Fields("email", "auth_provider").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personal_access_tokens", PersonalAccessToken.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
		edge.To("tfa_settings", TFASetting.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
		edge.To("setting", UserSetting.Type).
			Required().
			Unique().
			Annotations(
				entx.CascadeAnnotationField("User"),
				entoas.Skip(true),
			),
		edge.To("email_verification_tokens", EmailVerificationToken.Type).
			Annotations(
				entx.CascadeAnnotationField("Owner"),
				entoas.Skip(true),
				entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
		edge.To("password_reset_tokens", PasswordResetToken.Type).
			Annotations(
				entx.CascadeAnnotationField("Owner"),
				entoas.Skip(true),
				entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
		edge.To("groups", Group.Type).
			Through("group_memberships", GroupMembership.Type),
		edge.To("organizations", Organization.Type).
			Through("org_memberships", OrgMembership.Type),
		edge.To("webauthn", Webauthn.Type).
			Annotations(entx.CascadeAnnotationField("Owner")),
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		// Delete users from groups and orgs when the user is deleted
		entx.CascadeThroughAnnotationField(
			[]entx.ThroughCleanup{
				{
					Field:   "User",
					Through: "OrgMembership",
				},
				{
					Field:   "User",
					Through: "GroupMembership",
				},
			},
		),
	}
}

// Policy of the User
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfContextHasPrivacyTokenOfType(&token.SignUpToken{}),
					rule.AllowIfContextHasPrivacyTokenOfType(&token.OrgInviteToken{}),
					rule.AllowIfContextHasPrivacyTokenOfType(&token.OauthTooToken{}),
					rule.DenyIfNoViewer(),
					rule.AllowIfSelf(),
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
					// rule.AllowIfAdmin(), // TODO: this currently is always skipped, setup admin policy to get users
					privacy.AlwaysDenyRule(),
				},
				ent.OpUpdateOne|ent.OpUpdate|ent.OpDeleteOne|ent.OpDelete,
			),
		},
		Query: privacy.QueryPolicy{
			// Privacy will be always allow, but interceptors will filter the queries
			privacy.AlwaysAllowRule(),
		},
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookUser(),
		hooks.HookDeleteUser(),
	}
}

// Interceptors of the User.
func (d User) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseUser(func(ctx context.Context, q *generated.UserQuery) error {
			// Filter query based on viewer context
			v := viewer.FromContext(ctx)
			if v != nil {
				// TODO: expand based on viewer settings to
				// obtain users in orgs, groups, etc
				// for now, this will just return self
				viewerID, exists := v.GetID()
				if exists {
					q.Where(user.ID(viewerID))
					return nil
				}
			}

			return nil
		}),
	}
}
