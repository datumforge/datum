package mixin

import (
	"context"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"github.com/datumforge/datum/internal/auth"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/intercept"
)

// SoftDeleteMixin implements the soft delete pattern for schemas.
type SoftDeleteMixin struct {
	mixin.Schema
}

// Fields of the SoftDeleteMixin.
func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
		field.String("deleted_by").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
			),
	}
}

// softDeleteSkipKey is used to indicate to allow soft deleted records to be returned in records
// and to skip soft delete on mutations and proceed with a regular delete
type softDeleteSkipKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/hooks.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteSkipKey{}, true)
}

// CheckSkipSoftDelete checks if the softDeleteSkipKey is set in the context
func CheckSkipSoftDelete(ctx context.Context) bool {
	return ctx.Value(softDeleteSkipKey{}) != nil
}

// softDeleteKey is used to indicate a soft delete mutation is in progress
type softDeleteKey struct{}

// IsSoftDelete returns a new context that informs the delete is a soft-delete for interceptor/hooks.
func IsSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// CheckIsSoftDelete checks if the softDeleteKey is set in the context
func CheckIsSoftDelete(ctx context.Context) bool {
	return ctx.Value(softDeleteKey{}) != nil
}

// Interceptors of the SoftDeleteMixin.
func (d SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteSkipKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

// SoftDeleteHook will soft delete records, by changing the delete mutation to an update and setting
// the deleted_at and deleted_by fields, unless the softDeleteSkipKey is set
func (d SoftDeleteMixin) SoftDeleteHook(next ent.Mutator) ent.Mutator {
	type SoftDelete interface {
		SetOp(ent.Op)
		Client() *generated.Client
		SetDeletedAt(time.Time)
		SetDeletedBy(string)
		WhereP(...func(*sql.Selector))
	}

	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		if skip, _ := ctx.Value(softDeleteSkipKey{}).(bool); skip {
			return next.Mutate(ctx, m)
		}

		actor, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			actor = "unknown"
		}

		sd, ok := m.(SoftDelete)
		if !ok {
			return nil, newUnexpectedMutationTypeError(m)
		}

		d.P(sd)
		sd.SetOp(ent.OpUpdate)

		// set that the transaction is a soft-delete
		ctx = IsSoftDelete(ctx)

		sd.SetDeletedAt(time.Now())
		sd.SetDeletedBy(actor)
		return sd.Client().Mutate(ctx, m)
	})
}

// Hooks of the SoftDeleteMixin.
func (d SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			d.SoftDeleteHook,
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d SoftDeleteMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
