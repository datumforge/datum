package schema

import (
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
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
			Annotations(entgql.Skip()).
			NotEmpty(),
		field.Time("expires").
			Comment("the expiration date of the invitation token which defaults to 14 days in the future from creation").
			Nillable(),
		field.String("recipient").
			Comment("the email used as input to generate the invitation token and is the destination person the invitation is sent to who is required to accept to join the organization").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			NotEmpty(),
		field.Enum("status").
			Comment("the status of the invitation").
			GoType(enums.InvitationSent).
			Default(string(enums.InvitationSent)),
		field.Enum("role").
			GoType(enums.Role("")).
			Default(string(enums.RoleMember)).
			Values(string(enums.RoleOwner)),
		field.Int("send_attempts").
			Comment("the number of attempts made to perform email send of the invitation, maximum of 5").
			Default(0),
		field.String("requestor_id").
			Comment("the user who initiated the invitation").
			Immutable().
			NotEmpty(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Nillable().
			Annotations(entgql.Skip()).
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
			Ref: "invites",
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
		// TODO: come back and add query + delete policies
		Mutation: privacy.MutationPolicy{
			privacy.OnMutationOperation(
				privacy.MutationPolicy{
					rule.AllowIfAdmin(),
					rule.AllowIfContextHasPrivacyTokenOfType(&token.OrgInviteToken{}),
					rule.HasInviteEditAccess(),
					rule.AllowMutationAfterApplyingOwnerFilter(),
					privacy.AlwaysDenyRule(),
				},
				ent.OpCreate,
			),
		},
	}
}
