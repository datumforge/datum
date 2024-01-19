// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
	"github.com/datumforge/datum/internal/ent/generated/invite"
)

// InviteDelete is the builder for deleting a Invite entity.
type InviteDelete struct {
	config
	hooks    []Hook
	mutation *InviteMutation
}

// Where appends a list predicates to the InviteDelete builder.
func (id *InviteDelete) Where(ps ...predicate.Invite) *InviteDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InviteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InviteDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InviteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(invite.Table, sqlgraph.NewFieldSpec(invite.FieldID, field.TypeString))
	_spec.Node.Schema = id.schemaConfig.Invite
	ctx = internal.NewSchemaConfigContext(ctx, id.schemaConfig)
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// InviteDeleteOne is the builder for deleting a single Invite entity.
type InviteDeleteOne struct {
	id *InviteDelete
}

// Where appends a list predicates to the InviteDelete builder.
func (ido *InviteDeleteOne) Where(ps ...predicate.Invite) *InviteDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *InviteDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{invite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InviteDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
