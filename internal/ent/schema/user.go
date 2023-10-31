package schema

import (
	"net/mail"
	"net/url"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/mixin"

	"github.com/google/uuid"
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
		mixin.BaseMixin{},
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
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
			Default("unknown").
			Annotations(
				entgql.OrderField("display_name"),
			),
		field.Bool("locked").
			Comment("user account is locked if unconfirmed or explicitly locked").
			Default(false),
		// TO DO figure out the right models to use to pass bytes
		//		field.Bytes("pubkey").
		//			Comment("user public key").
		//			NotEmpty().
		//			Default([]byte{}),
		//		field.Bytes("privkey").
		//			Comment("user private key").
		//			Optional().
		//			Sensitive().
		//			Nillable(),
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
			Optional().
			Nillable(),
		field.Time("silenced_at").
			Comment("The time the user was silenced").
			Optional().
			Nillable(),
		field.Time("suspended_at").
			Comment("The time the user was suspended").
			Optional().
			Nillable(),
		//
		// Fields present in local account types but not remote accounts
		//
		//		field.Bytes("passwordHash").
		//			Comment("user bcrypt password hash").
		//			Sensitive().
		//			Nillable().
		//			Optional(). // only local accounts have hashes. empty hashes will fail bcrypt.ConstantTimeCompare() regardless
		//			MaxLen(60). // All hashes have a len of 60. MinLen not set (for non-local accounts)
		//			Validate(func(b []byte) error {
		//				if !bytes.HasPrefix(b, []byte("$2a$")) {
		//					return fmt.Errorf("invalid bcrypt password hash")
		//				}
		//				return nil
		//			}),
		field.String("recovery_code").
			// BIP words?
			Comment("local Actor password recovery code generated during account creation").
			// TODO: specify len, validate
			Sensitive().
			Nillable().
			Optional(),
		// TO DO figure out what model to use for this
		//		field.Enum("locale").
		//			Comment("local user locale").
		//			Values(
		//				language.AmericanEnglish.String(),
		//				// TODO: additional languages at some point, probably.
		//			),
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
			Annotations(entsql.Annotation{
				// When a user is deleted, delete the sessions
				OnDelete: entsql.Cascade}),
		edge.From("groups", Group.Type).Ref("users"),
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

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	// Privacy policy defined in the BaseMixin and TenantMixin.
	return nil
}
