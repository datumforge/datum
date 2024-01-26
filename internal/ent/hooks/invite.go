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
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/marionette"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// HookInvite runs on invite mutations
func HookInvite() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.InviteFunc(func(ctx context.Context, m *generated.InviteMutation) (generated.Value, error) {
			if err := personalOrgNoInvite(ctx, m); err != nil {
				m.Logger.Infow("external users cannot be invited to personal organizations")

				return nil, ErrPersonalOrgsNoMembers
			}

			isMember, err := doesUserHaveMembership(ctx, m)
			if err != nil {
				m.Logger.Errorw("error checking membership", "error", err)

				return nil, err
			}

			// already a member, nothing to do here
			if isMember {
				m.Logger.Infow("user is already a member of the organization")

				return nil, ErrUserAlreadyOrgMember
			}

			if IsUniqueConstraintError(err) {
				m.Logger.Infow("invitation for user already exists")
				inv, err := incrementAttempts(ctx, m)
				if err == nil {
					return inv, err
				}

				return nil, err
			}

			m, err = createAndSetToken(ctx, m)
			if err != nil {
				m.Logger.Errorw("error creating verification token", "error", err)

				return nil, err
			}

			return next.Mutate(ctx, m)
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

// personalOrgNoInvite checks if the mutation is for a personal org and denies if true
func personalOrgNoInvite(ctx context.Context, m *generated.InviteMutation) error {
	orgID, ok := m.OwnerID()
	if ok {
		parentOrg, err := m.Client().Organization.Get(ctx, orgID)
		if err != nil {
			return err
		}

		if parentOrg.PersonalOrg {
			return ErrPersonalOrgsNoChildren
		}
	}

	return nil
}

func createAndSetToken(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
	email, _ := m.Recipient()
	orgID, _ := m.OwnerID()

	org, err := m.Client().Organization.Query().Where(organization.ID(orgID)).Only(ctx)
	if err != nil {
		return nil, err
	}

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

	requestor, err := m.Client().User.Query().Where(user.ID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	invite := &emails.Invite{
		OrgName:   org.Name,
		Token:     token,
		Requestor: requestor.FirstName,
		Recipient: email,
	}

	if err := m.Marionette.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return sendOrgInvitationEmail(ctx, m, invite)
	}), marionette.WithErrorf("could not send invitation email to user %s", email),
	); err != nil {
		m.Logger.Errorw("unable to queue email for sending")

		return nil, err
	}

	return m, nil
}

func sendOrgInvitationEmail(ctx context.Context, m *generated.InviteMutation, i *emails.Invite) (err error) {
	data := emails.InviteData{
		InviterName: i.Requestor,
		OrgName:     i.OrgName,
		EmailData: emails.EmailData{
			Sender: m.Emails.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: i.Recipient,
			},
		},
	}

	if data.InviteURL, err = m.Emails.InviteURL(i.Token); err != nil {
		return err
	}

	msg, err := emails.InviteEmail(data)
	if err != nil {
		return err
	}

	return m.Emails.Send(msg)
}

var attempts = 5

func incrementAttempts(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
	curr, err := m.OldSendAttempts(ctx)
	if err != nil {
		return nil, err
	}

	if curr <= attempts {
		m.ResetRequestorID()
		m.ResetToken()
		m.ResetSecret()
		m.ResetExpires()
		m.SetSendAttempts(curr + 1)

		return m, nil
	} else {
		return nil, err
	}
}
