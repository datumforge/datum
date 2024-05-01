package schema

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/internal/ent/interceptors"
	"github.com/datumforge/datum/pkg/auth"
)

const (
	ownerFieldName = "owner_id"
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
	// SkipInterceptor skips the interceptor for that schema for all queries, or specific types,
	// this is useful for tokens, etc
	SkipInterceptor interceptors.SkipMode
}

// Fields of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Fields() []ent.Field {
	ownerIDField := field.String(ownerFieldName).Annotations(entoas.Skip(true))

	if !orgOwned.AllowWhere {
		ownerIDField.Annotations(entgql.Skip(), entoas.Skip(true))
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
		Field(ownerFieldName).
		Ref(orgOwned.Ref).
		Annotations(entoas.Skip(true)).
		Unique()

	if !orgOwned.Optional {
		ownerEdge.Required()
	}

	if orgOwned.SkipOASGeneration {
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

func (orgOwned OrgOwnerMixin) Interceptors() []ent.Interceptor {
	if orgOwned.Optional {
		// do not add interceptors if the field is optional
		return []ent.Interceptor{}
	} else {
		return []ent.Interceptor{
			intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
				if orgOwned.SkipInterceptor == interceptors.SkipAll {
					return nil
				}

				orgID, err := auth.GetOrganizationIDFromContext(ctx)
				if err != nil {
					ctxQuery := ent.QueryFromContext(ctx)

					// Skip the interceptor if the query is for a single entity
					// and the BypassInterceptor flag is set for Only queries
					if orgOwned.SkipInterceptor == interceptors.SkipOnlyQuery && ctxQuery.Op == "Only" {
						return nil
					}

					return err
				}

				// sets the owner id on the query for the current organization
				orgOwned.P(q, orgID)

				return nil
			}),
		}
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (orgOwned OrgOwnerMixin) P(w interface{ WhereP(...func(*sql.Selector)) }, orgID string) {
	w.WhereP(
		sql.FieldEQ(ownerFieldName, orgID),
	)
}
