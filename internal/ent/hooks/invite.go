package hooks

import (
	"context"

	"entgo.io/ent"

	"github.com/datumforge/datum/internal/ent/enums"
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
	"github.com/datumforge/datum/internal/utils/ulids"
)

// HookInvite runs on invite create mutations
func HookInvite() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.InviteFunc(func(ctx context.Context, m *generated.InviteMutation) (generated.Value, error) {
			m, err := setRequestor(ctx, m)
			if err != nil {
				m.Logger.Errorw("unable to determine requestor")

				return nil, err
			}

			// check that the invite isn't to a personal organization
			if err := personalOrgNoInvite(ctx, m); err != nil {
				m.Logger.Infow("unable to add user to specified organization", "error", err)

				return nil, err
			}

			// check if user exists
			email, _ := m.Recipient()
			inviteUser, err := getUserByEmail(ctx, m, email)
			if err != nil {
				// if error is anything other than not found, return now
				if !generated.IsNotFound(err) {
					return nil, err
				}
			}

			// generate token based on recipient + target org ID
			m, err = setRecipientAndToken(ctx, m)
			if err != nil {
				m.Logger.Errorw("error creating verification token", "error", err)

				return nil, err
			}

			// user exists, so automatically add them to the organization but record the invite in the database
			if inviteUser != nil {
				// check to see if user already has membership in the organization (or someone with the provided email)
				isMember, err := doesUserHaveMembership(ctx, m, inviteUser)
				if err != nil {
					m.Logger.Errorw("error checking membership", "error", err)

					return nil, err
				}

				// already a member, nothing to do here
				if isMember {
					m.Logger.Infow("user is already a member of the organization")

					return nil, ErrUserAlreadyOrgMember
				}

				// set status to accepted
				m.SetStatus(enums.InvitationAccepted)

				// run the mutation
				retValue, err := next.Mutate(ctx, m)
				if err != nil {
					m.Logger.Errorw("unable to create invitation", "error", err)
				}

				return retValue, nil
			}

			// attempt to do the mutation for a new user invite
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				if IsUniqueConstraintError(err) {
					m.Logger.Infow("invitation for user already exists")

					// update invite instead
					retValue, err = UpdateInvite(ctx, m)
					if err != nil {
						m.Logger.Errorw("unable to update invitation", "error", err)
					}

					return retValue, err
				}

				m.Logger.Errorw("unable to create org invitation", "error", err)

				return retValue, err
			}

			// non-blocking queued email
			if err := createInviteToSend(ctx, m); err != nil {
				m.Logger.Errorw("error sending email to user", "error", err)
			}

			return retValue, err
		})
	}, ent.OpCreate)
}

// HookInviteAccepted adds the user to the organization when the status is accepted
func HookInviteAccepted() ent.Hook {
	return hook.On(func(next ent.Mutator) ent.Mutator {
		return hook.InviteFunc(func(ctx context.Context, m *generated.InviteMutation) (generated.Value, error) {
			status, ok := m.Status()
			if !ok || status != enums.InvitationAccepted {
				// nothing to do here
				return next.Mutate(ctx, m)
			}

			ownerID, ownerOK := m.OwnerID()
			role, roleOK := m.Role()
			recipient, recipientOK := m.Recipient()

			// if we are missing any, get them from the db
			// this should happen on an update mutation
			if !ownerOK || !roleOK || !recipientOK {
				id, _ := m.ID()

				invite, err := m.Client().Invite.Get(ctx, id)
				if err != nil {
					m.Logger.Errorw("unable to get existing invite", "error", err)

					return nil, err
				}

				ownerID = invite.OwnerID
				role = invite.Role
				recipient = invite.Recipient
			}

			user, err := getUserByEmail(ctx, m, recipient)
			if err != nil {
				m.Logger.Errorw("unable to get user", "error", err)

				return nil, err
			}

			input := generated.CreateOrgMembershipInput{
				UserID: user.ID,
				OrgID:  ownerID,
				Role:   &role,
			}

			// add user to the inviting org
			if _, err := m.Client().OrgMembership.Create().SetInput(input).Save(ctx); err != nil {
				m.Logger.Errorw("unable to add user to organization", "error", err)

				return nil, err
			}

			// finish the mutation
			retValue, err := next.Mutate(ctx, m)
			if err != nil {
				return nil, err
			}

			// fetch org details to pass the name in the email
			org, err := m.Client().Organization.Query().Where(organization.ID(ownerID)).Only(ctx)
			if err != nil {
				m.Logger.Errorw("unable to get organization", "error", err)

				return retValue, err
			}

			invite := &emails.Invite{
				OrgName:   org.Name,
				Recipient: recipient,
				Role:      string(role),
			}

			// send an email to recipient notifying them they've been added to a datum organization
			if err := m.Marionette.Queue(marionette.TaskFunc(func(ctx context.Context) error {
				return sendOrgAccepted(ctx, m, invite)
			}), marionette.WithErrorf("could not send invitation email to user %s", recipient),
			); err != nil {
				m.Logger.Errorw("unable to queue email for sending")

				return retValue, err
			}

			// delete the invite that has been accepted
			if err := deleteInvite(ctx, m); err != nil {
				return retValue, err
			}

			return retValue, err
		})
	}, ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne)
}

// getUserByEmail checks to see if there is an existing user in the system based on the provided email,
// and returns the user if they do exist or nil if they don't
func getUserByEmail(ctx context.Context, m *generated.InviteMutation, email string) (*generated.User, error) {
	user, err := m.Client().User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		m.Logger.Errorw("could not find user by email", "error", err)

		return nil, err
	}

	return user, nil
}

// doesUserHaveMembership checks if the user already has membership to the requested organization;
func doesUserHaveMembership(ctx context.Context, m *generated.InviteMutation, entUser *generated.User) (bool, error) {
	orgID, _ := m.OwnerID()

	return m.Client().OrgMembership.Query().
		Where((orgmembership.HasUserWith(user.ID(entUser.ID)))).
		Where((orgmembership.HasOrgWith((organization.ID(orgID))))).Exist(ctx)
}

// personalOrgNoInvite checks if the mutation is for a personal org and denies if true or
// if the user does not hav access to that organization
func personalOrgNoInvite(ctx context.Context, m *generated.InviteMutation) error {
	orgID, ok := m.OwnerID()
	if ok {
		org, err := m.Client().Organization.Get(ctx, orgID)
		if err != nil {
			return err
		}

		if org.PersonalOrg {
			return ErrPersonalOrgsNoChildren
		}
	}

	return nil
}

// setRecipientAndToken function is responsible for generating a invite token based on the
// recipient's email and the target organization ID
func setRecipientAndToken(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
	email, _ := m.Recipient()
	owner, _ := m.OwnerID()

	oid, err := ulids.Parse(owner)
	if err != nil {
		return nil, err
	}

	verify, err := tokens.NewOrgInvitationToken(email, oid)
	if err != nil {
		return nil, err
	}

	token, secret, err := verify.Sign()
	if err != nil {
		return nil, err
	}

	// set values on mutation
	m.SetToken(token)
	m.SetExpires(verify.ExpiresAt)
	m.SetSecret(secret)

	return m, nil
}

// setRequestor sets the requestor on the mutation
func setRequestor(ctx context.Context, m *generated.InviteMutation) (*generated.InviteMutation, error) {
	userID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		m.Logger.Errorw("unable to get requestor", "error", err)

		return m, err
	}

	m.SetRequestorID(userID)

	return m, nil
}

// createInviteToSend sets the necessary data to send invite email + token
func createInviteToSend(ctx context.Context, m *generated.InviteMutation) error {
	// these are all required fields on create so should be found
	orgID, _ := m.OwnerID()
	reqID, _ := m.RequestorID()
	token, _ := m.Token()
	email, _ := m.Recipient()
	role, _ := m.Role()

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
		Role:      string(role),
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

// sendOrgInvitationEmail composes the email metadata and sends via email manager
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

// sendOrgAccepted composes the email metadata to notify the user they've been joined to the org
func sendOrgAccepted(ctx context.Context, m *generated.InviteMutation, i *emails.Invite) (err error) {
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

	msg, err := emails.InviteAccepted(data)
	if err != nil {
		return err
	}

	return m.Emails.Send(msg)
}

var maxAttempts = 5

// UpdateInvite if the invite already exists, set a new token, secret, expiration, and increment the attempts
// error at max attempts to resend
func UpdateInvite(ctx context.Context, m *generated.InviteMutation) (*generated.Invite, error) {
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

	m.SetSendAttempts(invite.SendAttempts)

	// these were already set when the invite was attempted to be added
	// we do not need to create these again
	secret, _ := m.Secret()
	token, _ := m.Token()
	expiresAt, _ := m.Expires()

	// update the invite
	return m.Client().Invite.
		UpdateOneID(invite.ID).
		SetSendAttempts(invite.SendAttempts).
		SetToken(token).
		SetExpires(expiresAt).
		SetSecret(secret).
		Save(ctx)
}

// deleteInvite deletes an invite from the database
func deleteInvite(ctx context.Context, m *generated.InviteMutation) error {
	id, _ := m.ID()

	if err := m.Client().Invite.DeleteOneID(id).Exec(ctx); err != nil {
		m.Logger.Errorw("unable to delete invite", "error", err)

		return err
	}

	return nil
}
