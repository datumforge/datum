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

		oID := view.GetOrganizationID()
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
<<<<<<< HEAD:internal/ent/privacy/rule/organization.go
=======

// contains looks for a string within a string slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// HasGroupReadAccess is a rule that returns allow decision if user has view access
func HasGroupReadAccess() privacy.GroupQueryRuleFunc {
	return privacy.GroupQueryRuleFunc(func(ctx context.Context, q *generated.GroupQuery) error {
		// eager load all fields
		if _, err := q.CollectFields(ctx); err != nil {
			return err
		}

		gCtx := graphql.GetFieldContext(ctx)

		// check group id from graphql arg context
		// when all orgs are requested, the interceptor will check group access
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

		// Skip to the next privacy rule (equivalent to return nil)
		return privacy.Skip
	})
}

// HasGroupMutationAccess is a rule that returns allow decision if user has edit or delete access
func HasGroupMutationAccess() privacy.GroupMutationRuleFunc {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *generated.GroupMutation) error {
		m.Logger.Debugw("checking mutation access")

		// No permissions checks on creation of org
		if m.Op().Is(ent.OpCreate) {
			return privacy.Skip
		}

		relation := fga.CanEdit
		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			relation = fga.CanDelete
		}

		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing when checking write access in group")
		}

		gID := view.GetGroupID()
		if gID == "" {
			return privacy.Denyf("missing group ID information in viewer")
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
>>>>>>> b5668ec (group auth with proper groupID name):internal/ent/privacy/rule/rule.go
