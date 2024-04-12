// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/documentdata"
	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// DocumentDataDelete is the builder for deleting a DocumentData entity.
type DocumentDataDelete struct {
	config
	hooks    []Hook
	mutation *DocumentDataMutation
}

// Where appends a list predicates to the DocumentDataDelete builder.
func (ddd *DocumentDataDelete) Where(ps ...predicate.DocumentData) *DocumentDataDelete {
	ddd.mutation.Where(ps...)
	return ddd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ddd *DocumentDataDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ddd.sqlExec, ddd.mutation, ddd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ddd *DocumentDataDelete) ExecX(ctx context.Context) int {
	n, err := ddd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ddd *DocumentDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(documentdata.Table, sqlgraph.NewFieldSpec(documentdata.FieldID, field.TypeString))
	_spec.Node.Schema = ddd.schemaConfig.DocumentData
	ctx = internal.NewSchemaConfigContext(ctx, ddd.schemaConfig)
	if ps := ddd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ddd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ddd.mutation.done = true
	return affected, err
}

// DocumentDataDeleteOne is the builder for deleting a single DocumentData entity.
type DocumentDataDeleteOne struct {
	ddd *DocumentDataDelete
}

// Where appends a list predicates to the DocumentDataDelete builder.
func (dddo *DocumentDataDeleteOne) Where(ps ...predicate.DocumentData) *DocumentDataDeleteOne {
	dddo.ddd.mutation.Where(ps...)
	return dddo
}

// Exec executes the deletion query.
func (dddo *DocumentDataDeleteOne) Exec(ctx context.Context) error {
	n, err := dddo.ddd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{documentdata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dddo *DocumentDataDeleteOne) ExecX(ctx context.Context) {
	if err := dddo.Exec(ctx); err != nil {
		panic(err)
	}
}
