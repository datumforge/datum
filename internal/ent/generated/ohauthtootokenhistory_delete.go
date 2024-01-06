// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
	"github.com/datumforge/datum/internal/ent/generated/ohauthtootokenhistory"
)

// OhAuthTooTokenHistoryDelete is the builder for deleting a OhAuthTooTokenHistory entity.
type OhAuthTooTokenHistoryDelete struct {
	config
	hooks    []Hook
	mutation *OhAuthTooTokenHistoryMutation
}

// Where appends a list predicates to the OhAuthTooTokenHistoryDelete builder.
func (oatthd *OhAuthTooTokenHistoryDelete) Where(ps ...predicate.OhAuthTooTokenHistory) *OhAuthTooTokenHistoryDelete {
	oatthd.mutation.Where(ps...)
	return oatthd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (oatthd *OhAuthTooTokenHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, oatthd.sqlExec, oatthd.mutation, oatthd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (oatthd *OhAuthTooTokenHistoryDelete) ExecX(ctx context.Context) int {
	n, err := oatthd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (oatthd *OhAuthTooTokenHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ohauthtootokenhistory.Table, sqlgraph.NewFieldSpec(ohauthtootokenhistory.FieldID, field.TypeString))
	_spec.Node.Schema = oatthd.schemaConfig.OhAuthTooTokenHistory
	ctx = internal.NewSchemaConfigContext(ctx, oatthd.schemaConfig)
	if ps := oatthd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, oatthd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	oatthd.mutation.done = true
	return affected, err
}

// OhAuthTooTokenHistoryDeleteOne is the builder for deleting a single OhAuthTooTokenHistory entity.
type OhAuthTooTokenHistoryDeleteOne struct {
	oatthd *OhAuthTooTokenHistoryDelete
}

// Where appends a list predicates to the OhAuthTooTokenHistoryDelete builder.
func (oatthdo *OhAuthTooTokenHistoryDeleteOne) Where(ps ...predicate.OhAuthTooTokenHistory) *OhAuthTooTokenHistoryDeleteOne {
	oatthdo.oatthd.mutation.Where(ps...)
	return oatthdo
}

// Exec executes the deletion query.
func (oatthdo *OhAuthTooTokenHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := oatthdo.oatthd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ohauthtootokenhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oatthdo *OhAuthTooTokenHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := oatthdo.Exec(ctx); err != nil {
		panic(err)
	}
}
