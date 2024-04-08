package schema

import (
	"net/url"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/ogen-go/ogen"

	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/internal/ent/schematype"
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
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
			).
			Optional(),
		field.Other("otherconfig", &schematype.TemplateConfig{}).
			SchemaType(map[string]string{
				dialect.SQLite: "json",
			}).
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
			).
			Optional(),
		field.Bytes("pair").
			GoType(customtypes.Pair{}).
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
			).
			DefaultFunc(func() customtypes.Pair {
				return customtypes.Pair{K: []byte("K"), V: []byte("V")}
			}),
		field.String("url").
			GoType(&url.URL{}).
			Annotations(
				entgql.Type("JSON"),
				entoas.Schema(ogen.String().AsArray()),
				entgql.Skip(entgql.SkipMutationUpdateInput, entgql.SkipType),
			).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}),
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
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}
