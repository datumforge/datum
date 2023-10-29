package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/datumforge/datum/internal/ent/mixin"

	"github.com/google/uuid"
)

// Membership maps users belonging to logical structures, firstly to organizations but in the future to projects, groups, etc.
type Membership struct {
	ent.Schema
}

// Fields of the Membership
func (Membership) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Bool("current").Default(false),
	}
}

// Edges of the Membership
func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("memberships").Unique().Required(),
		edge.From("user", User.Type).Ref("memberships").Unique().Required(),
		edge.From("group", Group.Type).Ref("memberships").Unique().Required(),
	}
}

// Indexes of the Membership
func (Membership) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("organization", "user", "group").Unique(),
	}
}

// Annotations of the Membership
func (Membership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Membership
func (Membership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}
