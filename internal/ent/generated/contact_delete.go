// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/contact"
	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ContactDelete is the builder for deleting a Contact entity.
type ContactDelete struct {
	config
	hooks    []Hook
	mutation *ContactMutation
}

// Where appends a list predicates to the ContactDelete builder.
func (cd *ContactDelete) Where(ps ...predicate.Contact) *ContactDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ContactDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ContactDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ContactDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(contact.Table, sqlgraph.NewFieldSpec(contact.FieldID, field.TypeString))
	_spec.Node.Schema = cd.schemaConfig.Contact
	ctx = internal.NewSchemaConfigContext(ctx, cd.schemaConfig)
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// ContactDeleteOne is the builder for deleting a single Contact entity.
type ContactDeleteOne struct {
	cd *ContactDelete
}

// Where appends a list predicates to the ContactDelete builder.
func (cdo *ContactDeleteOne) Where(ps ...predicate.Contact) *ContactDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *ContactDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{contact.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ContactDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}
