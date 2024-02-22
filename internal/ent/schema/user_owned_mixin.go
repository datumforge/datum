package schema

import (
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UserOwnedMixin struct {
	mixin.Schema
	Ref         string
	Optional    bool
	AllowUpdate bool
}

// Fields of the UserOwnedMixin
func (userOwned UserOwnedMixin) Fields() []ent.Field {
	ownerIDField := field.String("owner_id").Annotations(
		entgql.Skip(),
	)
	if userOwned.Optional {
		ownerIDField.Optional()
	}

	return []ent.Field{
		ownerIDField,
	}
}

// Edges of the UserOwnedMixin
func (userOwned UserOwnedMixin) Edges() []ent.Edge {
	if userOwned.Ref == "" {
		panic(errors.New("ref must be non-empty string")) // nolint: goerr113
	}

	ownerEdge := edge.
		From("owner", User.Type).
		Field("owner_id").
		Ref(userOwned.Ref).
		Unique()

	if !userOwned.Optional {
		ownerEdge.Required()
	}

	if !userOwned.AllowUpdate {
		ownerEdge.Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
		)
	}

	return []ent.Edge{
		ownerEdge,
	}
}
