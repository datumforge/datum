package schema

import (
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OrgOwnerMixin struct {
	mixin.Schema
	// Ref table for the id
	Ref string
	// Optional makes the owner id field not required
	Optional bool
	// SkipOASGeneration skips open api spec generation for the field
	SkipOASGeneration bool
	// AllowWhere includes the owner_id field in gql generated fields
	AllowWhere bool
}

// Fields of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Fields() []ent.Field {
	ownerIDField := field.String("owner_id")

	if !orgOwned.AllowWhere {
		ownerIDField.Annotations(entgql.Skip())
	}

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

	if orgOwned.SkipOASGeneration {
		ownerEdge.Annotations(
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
