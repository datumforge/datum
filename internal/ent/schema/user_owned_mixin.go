package schema

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/ent/generated/intercept"
	"github.com/datumforge/datum/pkg/auth"
)

type UserOwnedMixin struct {
	mixin.Schema
	// Ref table for the id
	Ref string
	// Optional makes the owner id field not required
	Optional bool
	// AllowUpdate allows the owner id field to be updated
	AllowUpdate bool
	// SkipOASGeneration skips open api spec generation for the field
	SkipOASGeneration bool
	// SoftDeleteIndex creates a unique index on the owner id field where deleted_at is null
	SoftDeleteIndex bool
	// AllowWhere includes the owner_id field in gql generated fields
	AllowWhere bool
	// BypassInterceptor skips the interceptor completely for that schema, this is useful for tokens, etc
	// This usually corresponds to schemas that also have the entgql.Skip(entgql.SkipAll) annotation
	BypassInterceptor bool
}

// Fields of the UserOwnedMixin
func (userOwned UserOwnedMixin) Fields() []ent.Field {
	ownerIDField := field.String("owner_id").Annotations(
		entgql.Skip(), entoas.Skip(true),
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

func (userOwned UserOwnedMixin) Interceptors() []ent.Interceptor {
	if userOwned.Optional {
		// do not add interceptors if the field is optional
		return []ent.Interceptor{}
	} else {
		return []ent.Interceptor{
			intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
				// Skip the interceptor for all queries if BypassInterceptor flag is set
				// This is needed for schemas that are never authorized users such as email verification tokens
				if userOwned.BypassInterceptor {
					return nil
				}

				userID, err := auth.GetUserIDFromContext(ctx)
				if err != nil {
					return err
				}

				// sets the owner id on the query for the current organization
				userOwned.P(q, userID)

				return nil
			}),
		}
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (userOwned UserOwnedMixin) P(w interface{ WhereP(...func(*sql.Selector)) }, userID string) {
	w.WhereP(
		sql.FieldEQ(ownerFieldName, userID),
	)
}
