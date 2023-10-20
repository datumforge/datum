package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Integration maps configured integrations (github, slack, etc.) to organizations
type Integration struct {
	ent.Schema
}

// Fields of the Integration
func (Integration) Fields() []ent.Field {
	return []ent.Field{
		// NOTE: the created_at and updated_at fields are automatically created by the AuditMixin, you do not need to re-declare / add them in these fields
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique().Immutable(),
		field.String("kind").Immutable(),
		field.String("description").Optional(),
		field.String("secret_name").Immutable(),
	}
}

// Edges of the Integration
func (Integration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Ref("integrations").Unique().Required(),
	}
}

// Annotations of the Integration
func (Integration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), (entgql.MutationUpdate())),
	}
}

// Mixin of the Integration
func (Integration) Mixin() []ent.Mixin {
	return []ent.Mixin{
		AuditMixin{},
	}
}
