package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GroupSettings holds the schema definition for the GroupSettings entity.
type GroupSettings struct {
	ent.Schema
}

// Fields of the GroupSettings.
func (GroupSettings) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
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
		edge.From("group", Group.Type).Ref("settings").Unique().Annotations(
			entgql.Skip(entgql.SkipAll),
		),
	}
}
