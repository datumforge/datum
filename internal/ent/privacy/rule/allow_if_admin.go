package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// AllowIfAdmin is used to determine whether a query or mutation should be allowed or skipped based on the user's admin status
// TODO: implement setting admin, this will currently always return a skip
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		v := viewer.FromContext(ctx)
		if v != nil && v.IsAdmin() {
			return privacy.Allow
		}
		return privacy.Skip
	})
}
