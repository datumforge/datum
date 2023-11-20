package rule

import (
	"context"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		ec, err := echox.EchoContextFromContext(ctx)
		if err != nil {
			return err
		}

		sub, err := echox.GetActorSubject(*ec)
		if err != nil {
			return err
		}

		if sub == "" {
			return privacy.Denyf("subject is missing")
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin(ctx) {
			return privacy.Allow
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
