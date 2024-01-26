package hooks

import (
	"context"
	"time"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/generated/hook"
	"github.com/datumforge/datum/internal/ent/generated/invite"
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

			m, err = setRequestorAndToken(ctx, m)
			if err != nil {
				m.Logger.Errorw("error creating verification token", "error", err)

				return nil, err
			}

			// attempt to do the mutation
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				if IsUniqueConstraintError(err) {
					m.Logger.Infow("invitation for user already exists")

					// update invite instead
					retValue, err = updateInvite(ctx, m)
					if err != nil {
						m.Logger.Errorw("unable to update invitation", "error", err)

						return retValue, err
					}
				}

				m.Logger.Errorw("unable to create org invitation", "error", err)

				return retValue, err
			}

			if err := createInviteToSend(ctx, m); err != nil {
				m.Logger.Errorw("error sending email to user", "error", err)

				return nil, err
			}

			return retValue, err
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

func setRequestorAndToken(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
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

func createInviteToSend(ctx context.Context, m *generated.InviteMutation) error {
	orgID, _ := m.OwnerID()
	reqID, _ := m.RequestorID()
	token, _ := m.Token()
	email, _ := m.Recipient()

	org, err := m.Client().Organization.Get(ctx, orgID)
	if err != nil {
		return err
	}

	requestor, err := m.Client().User.Get(ctx, reqID)
	if err != nil {
		return err
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

		return err
	}

	return nil
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

var maxAttempts = 5

// updateInvite if the invite already exists, set a new token, secret, expiration, and increment the attempts
// error at max attempts to resend
func updateInvite(ctx context.Context, m *generated.InviteMutation) (*generated.Invite, error) {
	// get the existing invite by recipient and owner
	rec, _ := m.Recipient()
	ownerID, _ := m.OwnerID()

	invite, err := m.Client().Invite.Query().Where(invite.Recipient(rec)).Where(invite.OwnerID(ownerID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// create update mutation
	if invite.SendAttempts >= maxAttempts {
		return nil, ErrMaxAttempts
	}

	// increment attempts
	invite.SendAttempts++

	// these were already set when the invite was attempted to be added
	// we do not need to create these again
	secret, _ := m.Secret()
	token, _ := m.Token()

	// update the invite
	return m.Client().Invite.
		UpdateOneID(invite.ID).
		SetExpires(time.Now().AddDate(0, 0, 14)). //nolint:gomnd
		SetSendAttempts(invite.SendAttempts).
		SetToken(token).
		SetSecret(secret).
		Save(ctx)
}
