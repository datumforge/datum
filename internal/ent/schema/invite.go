package schema

import (
	"net/mail"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Invite holds the schema definition for the Invite entity
type Invite struct {
	ent.Schema
}

// Fields of the Invite
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").
			Comment("the verification token sent to the user via email which should only be provided to the /verify endpoint + handler").
			Unique().
			NotEmpty(),
		field.Time("ttl").
			Comment("the ttl of the verification token which defaults to 7 days").
			Nillable(),
		field.String("invited_email").
			Comment("the email used as input to generate the verification token and is the destination person the invitation is sent to who is required to accept to join the organization").
			Validate(func(email string) error {
				_, err := mail.ParseAddress(email)
				return err
			}).
			NotEmpty(),
		field.Enum("invitestatus").
			GoType(enums.InvitationSent).
			Default(string(enums.InvitationSent)),
		field.String("requestor_id").
			Comment("the user who initatied the invitation").
			NotEmpty(),
		field.Bytes("secret").
			Comment("the comparison secret to verify the token's signature").
			NotEmpty().
			Nillable(),
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
