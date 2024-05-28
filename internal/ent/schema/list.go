package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// List holds the schema definition for the List entity
type List struct {
	ent.Schema
}

// Fields of the List
func (List) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("name of the list"),
		field.String("description").
			Comment("description of the list"),
	}
}

// Mixin of the List
func (List) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.TagMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		OrgOwnerMixin{
			Ref:             "lists",
			AllowWhere:      true,
			SkipInterceptor: interceptors.SkipOnlyQuery,
		},
	}
}

// Edges of the List
func (List) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type),
	}
}

// Annotations of the List
func (List) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			OrgOwnedField:   true,
			NillableIDField: true,
			IDField:         "OwnerID",
		},
	}
}
