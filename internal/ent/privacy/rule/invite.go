package rule

import (
	"context"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
)

// HasInviteEditAccess is a rule that returns allow decision if user has edit access to invite users to an organization
// TODO: this should able to be more generic
func HasInviteEditAccess() privacy.InviteMutationRuleFunc {
	return privacy.InviteMutationRuleFunc(func(ctx context.Context, m *generated.InviteMutation) error {
		m.Logger.Debugw("checking mutation access")

		relation := fgax.CanEdit

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		oID, ok := m.OwnerID()
		if !ok {
			return privacy.Skipf("owner not set in mutation, skipping")
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

		return privacy.Skip
	})
}
