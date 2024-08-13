package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/datumforge/entx"
	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
	"github.com/datumforge/datum/pkg/enums"
	"github.com/datumforge/datum/pkg/gqlplugin/searchgen"
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
		emixin.TagMixin{},
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
				searchgen.FieldSearchable(),
				entgql.OrderField("name"),
			),
		field.Enum("template_type").
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
			),
		field.JSON("uischema", customtypes.JSONObject{}).
			Comment("the uischema for the template to render in the UI").
			Annotations(
				entgql.Type("JSON"),
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
		index.Fields("name", "owner_id", "template_type").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}

// Annotations of the Template
func (Template) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			NillableIDField: true,
			OrgOwnedField:   true,
			IDField:         "OwnerID",
		},
		entx.SchemaSearchable(true),
	}
}

// Policy of the Template
func (Template) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.TemplateMutationRuleFunc(func(ctx context.Context, m *generated.TemplateMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),

			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.TemplateQueryRuleFunc(func(ctx context.Context, q *generated.TemplateQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
