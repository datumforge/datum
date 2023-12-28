package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/mixin"
)

type GroupMembership struct {
	ent.Schema
}

func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("group_role").GoType(enums.RoleMember),
	}
}

func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("members").
			Required().
			Unique(),
		edge.From("user", User.Type).
			Ref("group_memberships").
			Required().
			Unique(),
		edge.From("role", Role.Type).Ref("group_roles").Unique(),
	}
}

func (GroupMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}

// Mixin of the Permission
func (GroupMembership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
