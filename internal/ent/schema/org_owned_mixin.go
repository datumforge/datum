package schema

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/ent/generated"
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
	// Required makes the owner id field required as input
	Required bool
	// AllowEmpty allows the owner id field to be empty
	AllowEmpty bool
	// SkipOASGeneration skips open api spec generation for the field
	SkipOASGeneration bool
	// SkipInterceptor skips the interceptor for that schema for all queries, or specific types,
	// this is useful for tokens, etc
	SkipInterceptor interceptors.SkipMode
}

// Fields of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Fields() []ent.Field {
	ownerIDField := field.String(ownerFieldName)

	if !orgOwned.Required {
		ownerIDField.Optional()

		// if explicitly set to allow empty values, otherwise ensure it is not empty
		if !orgOwned.AllowEmpty {
			ownerIDField.NotEmpty()
		}
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
		Unique()

	if orgOwned.Required {
		ownerEdge.Required()
	}

	return []ent.Edge{
		ownerEdge,
	}
}

// Hooks of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Hooks() []ent.Hook {
	if orgOwned.AllowEmpty {
		// do not add hooks if the field is optional
		return []ent.Hook{}
	}

	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				// set owner on create mutation
				if m.Op() == ent.OpCreate {
					orgID, err := auth.GetOrganizationIDFromContext(ctx)
					if err != nil {
						return nil, fmt.Errorf("failed to get organization id from context: %w", err)
					}

					// set owner on mutation
					if err := m.SetField(ownerFieldName, orgID); err != nil {
						return nil, err
					}
				} else {
					orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
					if err != nil {
						return nil, fmt.Errorf("failed to get organization id from context: %w", err)
					}

					// filter by owner on update and delete mutations
					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *generated.Client
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, ErrUnexpectedMutationType
					}

					orgOwned.P(mx, orgIDs)
				}

				return next.Mutate(ctx, m)
			})
		},
	}
}

// Interceptors of the OrgOwnerMixin
func (orgOwned OrgOwnerMixin) Interceptors() []ent.Interceptor {
	if orgOwned.AllowEmpty {
		// do not add interceptors if the field is optional
		return []ent.Interceptor{}
	}

	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			if orgOwned.SkipInterceptor == interceptors.SkipAll {
				return nil
			}

			ctxQuery := ent.QueryFromContext(ctx)

			// Skip the interceptor if the query is for a single entity
			// and the BypassInterceptor flag is set for Only queries
			if orgOwned.SkipInterceptor == interceptors.SkipOnlyQuery && ctxQuery.Op == "Only" {
				return nil
			}

			orgIDs, err := auth.GetOrganizationIDsFromContext(ctx)
			if err != nil {
				return err
			}

			// sets the owner id on the query for the current organization
			orgOwned.P(q, orgIDs)

			return nil
		}),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (orgOwned OrgOwnerMixin) P(w interface{ WhereP(...func(*sql.Selector)) }, orgIDs []string) {
	w.WhereP(
		sql.FieldIn(ownerFieldName, orgIDs...),
	)
}
