// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/entitlementplanfeature"
	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// EntitlementPlanFeatureDelete is the builder for deleting a EntitlementPlanFeature entity.
type EntitlementPlanFeatureDelete struct {
	config
	hooks    []Hook
	mutation *EntitlementPlanFeatureMutation
}

// Where appends a list predicates to the EntitlementPlanFeatureDelete builder.
func (epfd *EntitlementPlanFeatureDelete) Where(ps ...predicate.EntitlementPlanFeature) *EntitlementPlanFeatureDelete {
	epfd.mutation.Where(ps...)
	return epfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (epfd *EntitlementPlanFeatureDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, epfd.sqlExec, epfd.mutation, epfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (epfd *EntitlementPlanFeatureDelete) ExecX(ctx context.Context) int {
	n, err := epfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (epfd *EntitlementPlanFeatureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(entitlementplanfeature.Table, sqlgraph.NewFieldSpec(entitlementplanfeature.FieldID, field.TypeString))
	_spec.Node.Schema = epfd.schemaConfig.EntitlementPlanFeature
	ctx = internal.NewSchemaConfigContext(ctx, epfd.schemaConfig)
	if ps := epfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, epfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	epfd.mutation.done = true
	return affected, err
}

// EntitlementPlanFeatureDeleteOne is the builder for deleting a single EntitlementPlanFeature entity.
type EntitlementPlanFeatureDeleteOne struct {
	epfd *EntitlementPlanFeatureDelete
}

// Where appends a list predicates to the EntitlementPlanFeatureDelete builder.
func (epfdo *EntitlementPlanFeatureDeleteOne) Where(ps ...predicate.EntitlementPlanFeature) *EntitlementPlanFeatureDeleteOne {
	epfdo.epfd.mutation.Where(ps...)
	return epfdo
}

// Exec executes the deletion query.
func (epfdo *EntitlementPlanFeatureDeleteOne) Exec(ctx context.Context) error {
	n, err := epfdo.epfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{entitlementplanfeature.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (epfdo *EntitlementPlanFeatureDeleteOne) ExecX(ctx context.Context) {
	if err := epfdo.Exec(ctx); err != nil {
		panic(err)
	}
}