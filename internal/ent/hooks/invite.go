package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/orgmembership"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/tokens"
)

// HookInvite runs on invite mutations
func HookInvite() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.InviteFunc(func(ctx context.Context, mutation *generated.InviteMutation) (generated.Value, error) {
			isMember, err := doesUserHaveMembership(ctx, mutation)
			if err != nil {
				mutation.Logger.Errorw("error checking membership", "error", err)

				return nil, err
			}

			// already a member, nothing to do here
			if isMember {
				mutation.Logger.Infow("user is already a member of the organization")

				return nil, ErrUserAlreadyOrgMember
			}

			// set token and secret for email
			mutation, err = createAndSetToken(ctx, mutation)
			if err != nil {
				mutation.Logger.Errorw("error creating verification token", "error", err)

				return nil, err
			}

			return next.Mutate(ctx, mutation)
		})
	}, ent.OpCreate)
}

// confirmUserExists checks to see if there is an existing user in the system based on the provided email, and returns the user if they do exist or nil if they don't
func confirmUserExists(ctx context.Context, m *generated.InviteMutation) (*generated.User, error) {
	email, _ := m.Recipient()

	user, err := m.Client().User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		m.Logger.Errorw("could not find user by email", "error", err)

		return nil, err
	}

	return user, nil
}

// doesUserHaveMembership checks if the user already has membership to the requested organization; if false user exists, but without requested organization membership
func doesUserHaveMembership(ctx context.Context, m *generated.InviteMutation) (bool, error) {
	orgID, _ := m.OwnerID()

	entUser, err := confirmUserExists(ctx, m)
	if err != nil {
		// if we did not find the user, return now
		if generated.IsNotFound(err) {
			return false, nil
		}

		// any other error, we should error
		return false, err
	}

	return m.Client().OrgMembership.Query().
		Where((orgmembership.HasUserWith(user.ID(entUser.ID)))).
		Where((orgmembership.HasOrgWith((organization.ID(orgID))))).Exist(ctx)
}

func createAndSetToken(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
	email, _ := m.Recipient()

	verify, err := tokens.NewVerificationToken(email)
	if err != nil {
		return nil, err
	}

	token, secret, err := verify.Sign()
	if err != nil {
		return nil, err
	}

	// set values on mutation
	m.SetToken(token)
	m.SetExpires(time.Now().Add(time.Hour * 24 * 14)) //nolint:gomnd
	m.SetSecret(secret)

	// requestor is the authenticated user
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		m.Logger.Errorw("unable to get requestor", "error", err)

		return nil, err
	}

	m.SetRequestorID(userID)

	return m, nil
}
