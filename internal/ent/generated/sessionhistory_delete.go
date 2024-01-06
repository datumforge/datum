// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
	"github.com/datumforge/datum/internal/ent/generated/sessionhistory"
)

// SessionHistoryDelete is the builder for deleting a SessionHistory entity.
type SessionHistoryDelete struct {
	config
	hooks    []Hook
	mutation *SessionHistoryMutation
}

// Where appends a list predicates to the SessionHistoryDelete builder.
func (shd *SessionHistoryDelete) Where(ps ...predicate.SessionHistory) *SessionHistoryDelete {
	shd.mutation.Where(ps...)
	return shd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (shd *SessionHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, shd.sqlExec, shd.mutation, shd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (shd *SessionHistoryDelete) ExecX(ctx context.Context) int {
	n, err := shd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (shd *SessionHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sessionhistory.Table, sqlgraph.NewFieldSpec(sessionhistory.FieldID, field.TypeString))
	_spec.Node.Schema = shd.schemaConfig.SessionHistory
	ctx = internal.NewSchemaConfigContext(ctx, shd.schemaConfig)
	if ps := shd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, shd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	shd.mutation.done = true
	return affected, err
}

// SessionHistoryDeleteOne is the builder for deleting a single SessionHistory entity.
type SessionHistoryDeleteOne struct {
	shd *SessionHistoryDelete
}

// Where appends a list predicates to the SessionHistoryDelete builder.
func (shdo *SessionHistoryDeleteOne) Where(ps ...predicate.SessionHistory) *SessionHistoryDeleteOne {
	shdo.shd.mutation.Where(ps...)
	return shdo
}

// Exec executes the deletion query.
func (shdo *SessionHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := shdo.shd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sessionhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (shdo *SessionHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := shdo.Exec(ctx); err != nil {
		panic(err)
	}
}
