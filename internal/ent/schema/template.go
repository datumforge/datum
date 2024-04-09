package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/enthistory"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/mixin"
)

type Template struct {
	ent.Schema
}

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

func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("the name of the template").
			NotEmpty().
			Annotations(
				entgql.OrderField("name"),
			),
		field.String("description").
			Comment("the description of the template").
			Optional(),
		field.JSON("jsonconfig", customtypes.JSONObject{}).
			Comment("the jsonschema object of the template").
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
			).
			Optional(),
		field.JSON("thatjsonbaby", map[string]interface{}{}).
			Comment("the jsonschema object of the template").
			Annotations(
				//				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
			).
			Optional(),
	}
}

func (Template) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (Template) Indexes() []ent.Index {
	return []ent.Index{}
}

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
