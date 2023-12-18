package rule

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// DenyIfNoSubject is a rule that returns deny decision if the subject is missing in the context.
func DenyIfNoSubject() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		sub, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return privacy.Denyf("cannot get subject from context")
		}

		if sub == "" {
			return privacy.Denyf("subject is missing")
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
