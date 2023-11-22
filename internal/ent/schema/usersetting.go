package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
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
	}
}

// Fields of the UserSetting
func (UserSetting) Fields() []ent.Field {
	return []ent.Field{
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
			Nillable(),
		field.String("recovery_code").
			Comment("local user password recovery code generated during account creation - does not exist for oauth'd users").
			Sensitive().
			Nillable().
			Optional(),
		field.Enum("status").
			NamedValues(
				"Active", "ACTIVE",
				"Inactive", "INACTIVE",
				"Deactivated", "DEACTIVATED",
				"Suspended", "SUSPENDED",
			).
			Default("ACTIVE"),
		field.Enum("role").
			NamedValues(
				"User", "USER",
				"Admin", "ADMIN",
				"Owner", "OWNER",
			).
			Default("USER"),
		field.JSON("permissions", []string{}).Default([]string{}),
		field.Bool("email_confirmed").Default(false),
		field.JSON("tags", []string{}).
			Comment("tags associated with the object").
			Default([]string{}),
	}
}

// Edges of the UserSetting
func (UserSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("setting").Unique(),
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
