package hooks

import (
	"context"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/authsessions"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/middleware/echocontext"
)

func updateUserSession(ctx context.Context, as authsessions.Auth, newOrgID string) error {
	au, err := auth.GetAuthenticatedUserContext(ctx)
	if err != nil {
		return err
	}

	user, err := generated.FromContext(ctx).User.Get(ctx, au.SubjectID)
	if err != nil {
		return err
	}

	ec, err := echocontext.EchoContextFromContext(ctx)
	if err != nil {
		return err
	}

	as.Logger.Debugw("updating user session", "user_id", user.ID, "org_id", newOrgID)

	_, err = as.GenerateUserAuthSessionWithOrg(ec, user, newOrgID)

	return err
}
