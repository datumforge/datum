package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/datumforge/enthistory"
	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Template holds the schema definition for the Template entity
type Template struct {
	ent.Schema
}

// Mixin of the Template
func (Template) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		mixin.SoftDeleteMixin{},
		emixin.IDMixin{},
		OrgOwnerMixin{
			Ref: "templates",
		},
	}
}

// Fields of the Template
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the template").
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
			),
		field.Enum("type").
			Comment("the type of the template, either a provided template or an implementation (document)").
			GoType(enums.DocumentType("")).
			Default(string(enums.Document)),
		field.String("description").
			Comment("the description of the template").
			Optional(),
		field.JSON("jsonconfig", customtypes.JSONObject{}).
			Comment("the jsonschema object of the template").
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
			),
		field.JSON("uischema", customtypes.JSONObject{}).
			Comment("the uischema for the template to render in the UI").
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
			).
			Optional(),
	}
}

// Edges of the Template
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("documents", DocumentData.Type).
			Annotations(
				entx.CascadeAnnotationField("Template"),
			),
	}
}

// Indexes of the Template
func (Template) Indexes() []ent.Index {
	return []ent.Index{
		// names should be unique, but ignore deleted names
		index.Fields("name").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Template
func (Template) Annotations() []schema.Annotation {
	return []schema.Annotation{
		enthistory.Annotations{
			Exclude: true,
		},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
