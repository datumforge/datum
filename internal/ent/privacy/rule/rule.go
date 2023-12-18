package rule

import (
	"context"

	"entgo.io/ent"
	"github.com/99designs/gqlgen/graphql"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/fga"
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

// HasOrgReadAccess is a rule that returns allow decision if user has view access
func HasOrgReadAccess() privacy.OrganizationQueryRuleFunc {
	return privacy.OrganizationQueryRuleFunc(func(ctx context.Context, q *generated.OrganizationQuery) error {
		gCtx := graphql.GetFieldContext(ctx)

		// check org id from graphql arg context
		// when all orgs are requested, the interceptor will check org access
		oID, ok := gCtx.Args["id"].(string)
		if !ok {
			return privacy.Allowf("nil request, bypassing auth check")
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		q.Logger.Infow("checking relationship tuples", "relation", fga.CanView, "organization_id", oID)

		access, err := q.Authz.CheckOrgAccess(ctx, userID, oID, fga.CanView)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			q.Logger.Infow("access allowed", "relation", fga.CanView, "organization_id", oID)

			return privacy.Allow
		}

		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// HasOrgMutationAccess is a rule that returns allow decision if user has edit or delete access
func HasOrgMutationAccess() privacy.OrganizationMutationRuleFunc {
	return privacy.OrganizationMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationMutation) error {
		m.Logger.Debugw("checking mutation access")

		relation := fga.CanEdit
		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			relation = fga.CanDelete
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		// No permissions checks on creation of org except if this is not a root org
		if m.Op().Is(ent.OpCreate) {
			parentOrgID, ok := m.ParentID()

			if ok {
				access, err := m.Authz.CheckOrgAccess(ctx, userID, parentOrgID, relation)
				if err != nil {
					return privacy.Skipf("unable to check access, %s", err.Error())
				}

				if !access {
					m.Logger.Debugw("access denied to parent org", "relation", relation, "organization_id", parentOrgID)

					return privacy.Deny
				}
			}

			return privacy.Skip
		}

		view := viewer.FromContext(ctx)
		if view == nil {
			m.Logger.Debugw("missing viewer context")

			return privacy.Denyf("viewer-context is missing when checking write access in org")
		}

		oID := view.OrganizationID()
		if oID == "" {
			m.Logger.Debugw("missing expected organization id")

			return privacy.Denyf("missing organization ID information in viewer")
		}

		m.Logger.Infow("checking relationship tuples", "relation", relation, "organization_id", oID)

		access, err := m.Authz.CheckOrgAccess(ctx, userID, oID, relation)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			m.Logger.Debugw("access allowed", "relation", relation, "organization_id", oID)

			return privacy.Allow
		}

		// deny if it was a mutation is not allowed
		return privacy.Deny
	})
}

// contains looks for a string within a string slice
func contains(s []string, e string) bool { //nolint:unused
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
