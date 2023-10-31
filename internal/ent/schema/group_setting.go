package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"

	"github.com/google/uuid"
)

// GroupSettings holds the schema definition for the GroupSettings entity.
type GroupSettings struct {
	ent.Schema
}

// Fields of the GroupSettings.
func (GroupSettings) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.Enum("visibility").NamedValues("public", "PUBLIC", "private", "PRIVATE").Default("PUBLIC"),
		field.Enum("join_policy").NamedValues(
			"open", "OPEN",
			"invite_only", "INVITE_ONLY",
			"application_only", "APPLICATION_ONLY",
			"invite_or_application", "INVITE_OR_APPLICATION",
		).Default("OPEN"),
	}
}

// Edges of the GroupSettings.
func (GroupSettings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("setting").Unique().Annotations(
			entgql.Skip(entgql.SkipAll),
		),
	}
}

// Annotations of the GroupSettings
func (GroupSettings) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the GroupSettings
func (GroupSettings) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
	}
}
