package rule

import (
	"context"

	"entgo.io/ent/entql"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// AllowIfSelf determines whether a query or mutation operation should be allowed based on whether the requested data is for the viewer
func AllowIfSelf() privacy.QueryMutationRule {
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		type UserFilter interface {
			WhereID(entql.StringP)
		}

		v := viewer.FromContext(ctx)
		if v == nil {
			return privacy.Skipf("missing viewer in context")
		}

		viewerID, exists := v.GetID()
		if !exists {
			return privacy.Skipf("anonymous viewer")
		}

		actualFilter, ok := f.(UserFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}

		actualFilter.WhereID(entql.StringEQ(viewerID))

		return privacy.Allow
	},
	)
}
