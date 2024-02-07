// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
	"github.com/datumforge/datum/internal/ent/generated/webauthn"
)

// WebauthnDelete is the builder for deleting a Webauthn entity.
type WebauthnDelete struct {
	config
	hooks    []Hook
	mutation *WebauthnMutation
}

// Where appends a list predicates to the WebauthnDelete builder.
func (wd *WebauthnDelete) Where(ps ...predicate.Webauthn) *WebauthnDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WebauthnDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WebauthnDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WebauthnDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(webauthn.Table, sqlgraph.NewFieldSpec(webauthn.FieldID, field.TypeString))
	_spec.Node.Schema = wd.schemaConfig.Webauthn
	ctx = internal.NewSchemaConfigContext(ctx, wd.schemaConfig)
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WebauthnDeleteOne is the builder for deleting a single Webauthn entity.
type WebauthnDeleteOne struct {
	wd *WebauthnDelete
}

// Where appends a list predicates to the WebauthnDelete builder.
func (wdo *WebauthnDeleteOne) Where(ps ...predicate.Webauthn) *WebauthnDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WebauthnDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{webauthn.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WebauthnDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
