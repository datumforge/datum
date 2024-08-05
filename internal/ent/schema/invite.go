package schema

import (
	"context"
	"net/mail"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/enthistory"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/hooks"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/privacy/rule"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/enums"
)

const (
	defaultInviteExpiresDays = 14
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
			).
			NotEmpty(),
		field.Time("expires").
			Comment("the expiration date of the invitation token which defaults to 14 days in the future from creation").
			Default(func() time.Time {
				return time.Now().AddDate(0, 0, defaultInviteExpiresDays)
			}).
			Optional(),
		field.String("recipient").
			Comment("the email used as input to generate the invitation token and is the destination person the invitation is sent to who is required to accept to join the organization").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			Immutable().
			NotEmpty(),
		field.Enum("status").
			Comment("the status of the invitation").
			GoType(enums.InviteStatus("")).
			Default(string(enums.InvitationSent)),
		field.Enum("role").
			GoType(enums.Role("")).
			Default(string(enums.RoleMember)),
		field.Int("send_attempts").
			Comment("the number of attempts made to perform email send of the invitation, maximum of 5").
			Default(0),
		field.String("requestor_id").
			Comment("the user who initiated the invitation").
			Immutable().
			Optional().
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
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{
			Ref: "invites",
			SkipTokenType: []token.PrivacyToken{
				&token.OrgInviteToken{},
			},
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
		edge.To("events", Event.Type),
	}
}

// Annotations of the Invite
func (Invite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			NillableIDField: true,
			OrgOwnedField:   true,
			IDField:         "OwnerID",
		},
		enthistory.Annotations{
			Exclude: true,
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
			rule.AllowIfContextHasPrivacyTokenOfType(&token.OrgInviteToken{}),
			rule.CanInviteMembers(),
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
