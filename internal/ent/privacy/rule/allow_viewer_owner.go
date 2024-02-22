package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// AllowIfOwnedByViewer determines whether a query or mutation operation should be allowed based on whether the requested data is owned by the viewer
func AllowIfOwnedByViewer() privacy.QueryMutationRule {
	type UserOwnedFilter interface {
		WhereHasOwnerWith(...predicate.User)
	}

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		v := viewer.FromContext(ctx)

		if v == nil {
			return privacy.Skipf("missing viewer in context")
		}

		viewerID, exists := v.GetID()
		if !exists {
			return privacy.Skipf("anonymous viewer")
		}

		actualFilter, ok := f.(UserOwnedFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		actualFilter.WhereHasOwnerWith(user.ID(viewerID))

		return privacy.Allow
	},
	)
}
