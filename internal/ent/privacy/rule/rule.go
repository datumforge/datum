package rule

import (
	"context"

	"entgo.io/ent"
	ofgaclient "github.com/openfga/go-sdk/client"

	"github.com/datumforge/datum/internal/echox"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/fga"
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

		// view := viewer.FromContext(ctx)
		// if view == nil {
		// 	return privacy.Denyf("viewer-context is missing")
		// }

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// HasOrgReadAccess is a rule that returns allow decision if user has view access
func HasOrgReadAccess() privacy.OrganizationQueryRuleFunc {
	return privacy.OrganizationQueryRuleFunc(func(ctx context.Context, q *generated.OrganizationQuery) error {
		userID, err := echox.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		sub := fga.Entity{
			Kind:       "user",
			Identifier: userID,
		}

		view := viewer.FromContext(ctx)
		objID := view.GetObjectID()

		obj := fga.Entity{
			Kind:       "organization",
			Identifier: objID,
		}

		q.Logger.Infow("checking relationship tuples", "relation", fga.CanView, "object", obj.String())

		checkReq := ofgaclient.ClientCheckRequest{
			User:     sub.String(),
			Relation: fga.CanView,
			Object:   obj.String(),
		}

		access, err := q.Authz.CheckTuple(ctx, checkReq)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			return privacy.Allow
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// HasOrgDeleteAccess is a rule that returns allow decision if user has delete access
func HasOrgDeleteAccess() privacy.OrganizationMutationRuleFunc {
	return privacy.OrganizationMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationMutation) error {
		m.Logger.Debugw("checking mutation access")

		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			userID, err := echox.GetUserIDFromContext(ctx)
			if err != nil {
				return err
			}

			sub := fga.Entity{
				Kind:       "user",
				Identifier: userID,
			}

			view := viewer.FromContext(ctx)
			objID := view.GetObjectID()

			obj := fga.Entity{
				Kind:       "organization",
				Identifier: objID,
			}

			m.Logger.Infow("checking relationship tuples", "relation", fga.CanDelete, "object", obj.String())

			checkReq := ofgaclient.ClientCheckRequest{
				User:     sub.String(),
				Relation: fga.CanView,
				Object:   obj.String(),
			}

			access, err := m.Authz.CheckTuple(ctx, checkReq)
			if err != nil {
				return privacy.Skipf("unable to check access, %s", err.Error())
			}

			if access {
				m.Logger.Debugw("access allowed", "relation", fga.CanDelete, "object", obj.String())

				return privacy.Allow
			}

			// deny if it was a delete mutation and user is not allowed
			return privacy.Deny
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
