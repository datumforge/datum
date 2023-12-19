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

// HasGroupReadAccess is a rule that returns allow decision if user has view access
func HasGroupReadAccess() privacy.GroupQueryRuleFunc {
	return privacy.GroupQueryRuleFunc(func(ctx context.Context, q *generated.GroupQuery) error {
		gCtx := graphql.GetFieldContext(ctx)

		if gCtx != nil {
			// check group id from graphql arg context
			// when all groups are requested, the interceptor will check group access
			gID, ok := gCtx.Args["id"].(string)
			if !ok {
				return privacy.Allowf("nil request, bypassing auth check")
			}

			userID, err := auth.GetUserIDFromContext(ctx)
			if err != nil {
				return err
			}

			q.Logger.Infow("checking relationship tuples", "relation", fga.CanView, "group_id", gID)

			access, err := q.Authz.CheckGroupAccess(ctx, userID, gID, fga.CanView)
			if err != nil {
				return privacy.Skipf("unable to check access, %s", err.Error())
			}

			if access {
				q.Logger.Infow("access allowed", "relation", fga.CanView, "group_id", gID)

				return privacy.Allow
			}
		}

		// Skip to the next privacy rule (equivalent to return nil)
		return privacy.Skip
	})
}

// HasGroupMutationAccess is a rule that returns allow decision if user has edit or delete access
func HasGroupMutationAccess() privacy.GroupMutationRuleFunc {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *generated.GroupMutation) error {
		gCtx := graphql.GetFieldContext(ctx)

		gID, ok := gCtx.Args["id"].(string)
		if !ok {
			return privacy.Allowf("nil request, bypassing auth check")
		}
		m.Logger.Debugw("checking mutation access")

		relation := fga.CanEdit
		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			relation = fga.CanDelete
		}

		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing when checking write access in group")
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		m.Logger.Infow("checking relationship tuples", "relation", relation, "group_id", gID)

		access, err := m.Authz.CheckGroupAccess(ctx, userID, gID, relation)
		if err != nil {
			return privacy.Skipf("unable to check access, %s", err.Error())
		}

		if access {
			m.Logger.Debugw("access allowed", "relation", relation, "group_id", gID)

			return privacy.Allow
		}

		// deny if it was a mutation is not allowed
		return privacy.Deny
	})
}

// CanCreateGroupsInOrg is a rule that returns allow decision if user has edit access in the organization
func CanCreateGroupsInOrg() privacy.GroupMutationRuleFunc {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *generated.GroupMutation) error {
		oID, ok := m.OwnerID()
		if !ok {
			return privacy.Skipf("no owner set on request, cannot check access")
		}

		m.Logger.Debugw("checking mutation access")

		relation := fga.CanEdit
		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			relation = fga.CanDelete
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
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
