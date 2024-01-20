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
	"github.com/datumforge/datum/internal/tokens"
)

// HookInvite runs on accesstoken mutations and sets expires
func HookInvite() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.InviteFunc(func(ctx context.Context, mutation *generated.InviteMutation) (generated.Value, error) {
			if mutation.Op().Is(ent.OpCreate) {
				create, err := doesUserHaveMembership(ctx, mutation)
				if false {
					if err := createAndSetToken(ctx, mutation); err != nil {
						mutation.Logger.Errorw("failed to create invitation", "error", err)
						return nil, err
					}
				}
				if err != nil {
					return create, err
				}
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
		return false, err
	}

	return m.Client().OrgMembership.Query().
		Where((orgmembership.HasUserWith(user.ID(entUser.ID)))).
		Where((orgmembership.HasOrgWith((organization.ID(orgID))))).Exist(ctx)
}

func createAndSetToken(ctx context.Context, m *generated.InviteMutation) error {

	userID, _ := m.ID()
	email, _ := m.Recipient()
	owner, _ := m.OwnerID()

	user, err := m.Client().User.Get(ctx, userID)
	if err != nil {
		return err
	}

	verify, err := tokens.NewVerificationToken(email)

	if err != nil {
		return err
	}

	token, secret, err := verify.Sign()
	if err != nil {
		return err
	}

	m.SetOwnerID(owner)
	m.SetToken(token)
	m.SetExpires(time.Now().Add(time.Hour * 24 * 14))
	m.SetRecipient(email)
	m.SetSecret(secret)
	m.SetRequestorID(user.ID)

	return err
}
