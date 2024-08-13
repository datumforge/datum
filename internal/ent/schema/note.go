package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	emixin "github.com/datumforge/entx/mixin"
	"github.com/datumforge/fgax/entfga"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/mixin"
)

// Note holds the schema definition for the Note entity
type Note struct {
	ent.Schema
}

// Fields of the Note
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("text").
			Comment("the text of the note").
			NotEmpty(),
	}
}

// Mixin of the Note
func (Note) Mixin() []ent.Mixin {
	return []ent.Mixin{
		emixin.AuditMixin{},
		emixin.IDMixin{},
		mixin.SoftDeleteMixin{},
		emixin.TagMixin{},
		OrgOwnerMixin{
			Ref: "notes",
		},
	}
}

// Edges of the Note
func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("entity", Entity.Type).
			Unique().
			Ref("notes"),
	}
}

// Indexes of the Note
func (Note) Indexes() []ent.Index {
	return []ent.Index{}
}

// Annotations of the Note
func (Note) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
		entfga.Annotations{
			ObjectType:      "organization",
			IncludeHooks:    false,
			NillableIDField: true,
			OrgOwnedField:   true,
			IDField:         "OwnerID",
		},
	}
}

// Hooks of the Note
func (Note) Hooks() []ent.Hook {
	return []ent.Hook{}
}

// Interceptors of the Note
func (Note) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{}
}

// Policy of the Note
func (Note) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.NoteMutationRuleFunc(func(ctx context.Context, m *generated.NoteMutation) error {
				return m.CheckAccessForEdit(ctx)
			}),

			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.NoteQueryRuleFunc(func(ctx context.Context, q *generated.NoteQuery) error {
				return q.CheckAccess(ctx)
			}),
			privacy.AlwaysDenyRule(),
		},
	}
}
