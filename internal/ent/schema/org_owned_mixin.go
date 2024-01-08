package schema

import (
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OrgOwnerMixin struct {
	mixin.Schema
	Ref      string
	Optional bool
}

// Fields of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Fields() []ent.Field {
	ownerIDField := field.String("owner_id").Annotations(
		entgql.Skip(),
	)

	if orgOwned.Optional {
		ownerIDField.Optional()
	}

	return []ent.Field{
		ownerIDField,
	}
}

// Edges of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Edges() []ent.Edge {
	if orgOwned.Ref == "" {
		panic(errors.New("ref must be non-empty string")) // nolint: goerr113
	}

	ownerEdge := edge.
		From("owner", Organization.Type).
		Field("owner_id").
		Ref(orgOwned.Ref).
		Unique()

	if !orgOwned.Optional {
		ownerEdge.Required()
	}

	return []ent.Edge{
		ownerEdge,
	}
}
