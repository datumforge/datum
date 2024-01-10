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

// GroupMembership holds the schema definition for the GroupMembership entity
type GroupMembership struct {
	ent.Schema
}

// Fields of the GroupMembership
func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			GoType(enums.RoleMember).
			Default(string(enums.RoleMember)),
		field.String("group_id"),
		field.String("user_id"),
	}
}

// Edges of the GroupMembership
func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).
			Field("group_id").
			Required().
			Unique(),
		edge.To("user", User.Type).
			Field("user_id").
			Required().
			Unique(),
	}
}

// Annotations of the GroupMembership
func (GroupMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the GroupMembership
func (GroupMembership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
		mixin.SoftDeleteMixin{},
	}
}
