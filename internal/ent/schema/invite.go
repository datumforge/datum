package schema

import (
	"context"
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/fgax/entfga"
)

// Invite holds the schema definition for the Invite entity
type Invite struct {
	ent.Schema
}

// Fields of the Invite
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Comment("the invitation token sent to the user via email which should only be provided to the /verify endpoint + handler").
			Unique().
			Sensitive().
			Annotations(
				entgql.Skip(),
				entoas.Skip(true),
			).
			NotEmpty(),
		field.Time("expires").
			Comment("the expiration date of the invitation token which defaults to 14 days in the future from creation").
			Annotations(entoas.Annotation{ReadOnly: true}).
			Nillable(),
		field.String("recipient").
			Comment("the email used as input to generate the invitation token and is the destination person the invitation is sent to who is required to accept to join the organization").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			Annotations(
				entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			).
			NotEmpty(),
		field.Enum("status").
			Comment("the status of the invitation").
			Annotations(entoas.Annotation{ReadOnly: true}).
			GoType(enums.InvitationSent).
			Default(string(enums.InvitationSent)),
		field.Enum("role").
			GoType(enums.Role("")).
			Default(string(enums.RoleMember)).
			Annotations(
				entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
				entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			).
			Values(string(enums.RoleOwner)),
		field.Int("send_attempts").
			Comment("the number of attempts made to perform email send of the invitation, maximum of 5").
			Annotations(entoas.Annotation{ReadOnly: true}).
			Default(0),
		field.String("requestor_id").
			Comment("the user who initiated the invitation").
			Annotations(entoas.Annotation{ReadOnly: true}).
			Immutable().
			NotEmpty(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Nillable().
			Annotations(entgql.Skip(), entoas.Skip(true)).
			Sensitive(),
	}
}

// Mixin of the Invite
func (Invite) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{
			Ref:        "invites",
			AllowWhere: true,
		},
	}
}

// Indexes of the Invite
func (Invite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("recipient", "owner_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Edges of the Invite
func (Invite) Edges() []ent.Edge {
	return []ent.Edge{
		// Edges go here
	}
}

// Annotations of the Invite
func (Invite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:   "organization",
			IncludeHooks: false,
			IDField:      "OwnerID",
		},
	}
}

// Hooks of the Invite
func (Invite) Hooks() []ent.Hook {
	return []ent.Hook{
		hooks.HookInvite(),
		hooks.HookInviteAccepted(),
	}
}

// Policy of the Invite
func (Invite) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowIfAdmin(),
			rule.AllowIfContextHasPrivacyTokenOfType(&token.OrgInviteToken{}),
			privacy.InviteMutationRuleFunc(func(ctx context.Context, m *generated.InviteMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),
			rule.AllowMutationAfterApplyingOwnerFilter(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			rule.AllowIfContextHasPrivacyTokenOfType(&token.OrgInviteToken{}),
			privacy.InviteQueryRuleFunc(func(ctx context.Context, q *generated.InviteQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
