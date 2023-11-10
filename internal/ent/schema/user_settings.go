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
)


// UserSettings holds the schema definition for the User entity.
type UserSettings struct {
	ent.Schema
}

// Mixin of the UserSettings
func (UserSettings) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Fields of the UserSettings
func (UserSettings) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("locked").
			Comment("user account is locked if unconfirmed or explicitly locked").
			Default(false),
		field.Time("silenced_at").
			Comment("The time the user was silenced").
			Optional().
			Nillable(),
		field.Time("suspended_at").
			Comment("The time the user was suspended").
			Optional().
			Nillable(),
		field.String("recovery_code").
			Comment("local Actor password recovery code generated during account creation").
			Sensitive().
			Nillable().
			Optional(),

	}
}

// Edges of the UserSettings
func (UserSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", Group.Type).Ref("setting").Unique()

	}
}

// Annotations of the UserSettings
func (UserSettings) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
