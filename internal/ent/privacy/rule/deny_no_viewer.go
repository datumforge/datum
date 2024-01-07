package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// DenyIfNoViewer returns deny if viewer is not present in context
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		v := viewer.FromContext(ctx)
		if v == nil {
			return privacy.Denyf("viewer is missing")
		}

		return privacy.Skip
	})
}
