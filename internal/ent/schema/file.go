package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"

	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/pkg/gqlplugin/searchgen"
)

// File defines the file schema.
type File struct {
	ent.Schema
}

// Fields returns file fields.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_name").
			Annotations(
				searchgen.FieldSearchable(),
			),
		field.String("file_extension"),
		field.Int("file_size").
			NonNegative().
			Optional(),
		field.String("content_type"),
		field.String("store_key"),
		field.String("category").
			Optional(),
		field.String("annotation").
			Optional().
			Annotations(
				searchgen.FieldSearchable(),
			),
	}
}

// Edges of the File
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("files").
			Unique(),
		edge.From("organization", Organization.Type).
			Ref("files"),
		edge.From("group", Group.Type).
			Ref("files"),
	}
}

// Mixin of the File
func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
		emixin.TagMixin{},
	}
}

// Annotations of the File
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entx.SchemaSearchable(true),
	}
}
