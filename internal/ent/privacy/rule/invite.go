package rule

import (
	"context"
	"strings"

	"github.com/datumforge/fgax"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/privacy"
	"github.com/datumforge/datum/pkg/auth"
)

const (
	inviteMemberRelation = "can_invite_members"
	inviteAdminRelation  = "can_invite_admins"
)

// CanInviteMembers is a rule that returns allow decision if user has access to invite members to the organization
func CanInviteMembers() privacy.InviteMutationRuleFunc {
	return privacy.InviteMutationRuleFunc(func(ctx context.Context, m *generated.InviteMutation) error {
		oID, ok := m.OwnerID()
		if !ok || oID == "" {
			// get organization from the auth context
			var err error

			oID, err = auth.GetOrganizationIDFromContext(ctx)
			if err != nil || oID == "" {
				return privacy.Skipf("no owner set on request, cannot check access")
			}
		}

		m.Logger.Debugw("checking mutation access")

		relation := inviteMemberRelation

		role, ok := m.Role()
		if ok && !strings.EqualFold(role.String(), fgax.MemberRelation) {
			relation = inviteAdminRelation
		}

		userID, err := auth.GetUserIDFromContext(ctx)
		if err != nil {
			return err
		}

		m.Logger.Infow("checking relationship tuples", "relation", relation, "organization_id", oID)

		ac := fgax.AccessCheck{
			SubjectID:   userID,
			SubjectType: auth.GetAuthzSubjectType(ctx),
			ObjectID:    oID,
			Relation:    relation,
		}

		access, err := m.Authz.CheckOrgAccess(ctx, ac)
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
