package schema

import (
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type UserOwnedMixin struct {
	mixin.Schema
	Ref               string
	Optional          bool
	AllowUpdate       bool
	SkipOASGeneration bool
	SoftDeleteIndex   bool
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
		Annotations(entoas.Skip(true)).
		Unique()

	if !userOwned.Optional {
		ownerEdge.Required()
	}

	if !userOwned.AllowUpdate {
		ownerEdge.Annotations(
			entgql.Skip(entgql.SkipMutationUpdateInput),
			entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		)
	}

	if userOwned.SkipOASGeneration {
		ownerEdge.Annotations(
			entoas.Skip(true),
			entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		)
	}

	return []ent.Edge{
		ownerEdge,
	}
}

// Indexes of the UserOwnedMixin
func (userOwned UserOwnedMixin) Indexes() []ent.Index {
	if !userOwned.SoftDeleteIndex {
		return []ent.Index{}
	}

	return []ent.Index{
		index.Fields("owner_id").
			Unique().Annotations(entsql.IndexWhere("deleted_at is NULL")),
	}
}
