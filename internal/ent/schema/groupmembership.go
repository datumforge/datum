package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// GroupMembership holds the schema definition for the GroupMembership entity.
type GroupMembership struct {
	ent.Schema
}

// Fields of the GroupMembership.
func (GroupMembership) Fields() []ent.Field {
	return []ent.Field{
		field.Time("joined_at").
			Default(time.Now),
		field.String("user_id"),
		field.String("group_id"),
	}
}

// Edges of the GroupMembership.
func (GroupMembership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
	}
}

// Mixin of the GroupMembership
func (GroupMembership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}

// Annotations of the GroupMembership
func (GroupMembership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
