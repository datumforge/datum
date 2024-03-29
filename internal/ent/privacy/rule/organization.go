package rule

import (
	"context"

	"entgo.io/ent"
	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
)

// HasOrgMutationAccess is a rule that returns allow decision if user has edit or delete access
func HasOrgMutationAccess() privacy.OrganizationMutationRuleFunc {
	return privacy.OrganizationMutationRuleFunc(func(ctx context.Context, m *generated.OrganizationMutation) error {
		m.Logger.Debugw("checking mutation access")

		relation := fgax.CanEdit
		if m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
			relation = fgax.CanDelete
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
			var exists bool
			oID, exists = m.ID()
			// if its still empty fail
			if !exists {
				m.Logger.Debugw("missing expected organization id")

				return privacy.Denyf("missing organization ID information in viewer")
			}
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
