package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/datumforge/datum/internal/ent/mixin"
)

// GroupSetting holds the schema definition for the GroupSetting entity.
type GroupSetting struct {
	ent.Schema
}

// Fields of the GroupSetting.
func (GroupSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("visibility").NamedValues("public", "PUBLIC", "private", "PRIVATE").Default("PUBLIC"),
		field.Enum("join_policy").NamedValues(
			"open", "OPEN",
			"invite_only", "INVITE_ONLY",
			"application_only", "APPLICATION_ONLY",
			"invite_or_application", "INVITE_OR_APPLICATION",
		).Default("OPEN"),
		field.JSON("tags", []string{}).
			Comment("tags associated with the object").
			Default([]string{}),
		field.Bool("sync_to_slack").Default(false),
		field.Bool("sync_to_github").Default(false),
	}
}

// Edges of the GroupSetting.
func (GroupSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("setting").Unique(),
	}
}

// Annotations of the GroupSetting
func (GroupSetting) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the GroupSetting
func (GroupSetting) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.AuditMixin{},
		mixin.IDMixin{},
	}
}
