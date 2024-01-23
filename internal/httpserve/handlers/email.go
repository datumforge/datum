package handlers

import (
	"context"

	"github.com/cenkalti/backoff/v4"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/marionette"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

var EmailURL = &emails.URLConfig{}
var EM = &emails.Config{}

func (h *Handler) SendVerificationEmail(user *User) error {
	contact := &sendgrid.Contact{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	data := emails.VerifyEmailData{
		EmailData: emails.EmailData{
			Sender: EM.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
		FullName: contact.FullName(),
	}

	var err error
	if data.VerifyURL, err = EmailURL.VerifyURL(user.GetVerificationToken()); err != nil {
		return err
	}

	msg, err := emails.VerifyEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.EmailManager.Send(msg)
}

// SendPasswordResetRequestEmail Send an email to a user to request them to reset their password
func (h *Handler) SendPasswordResetRequestEmail(user *User) error {
	data := emails.ResetRequestData{
		EmailData: emails.EmailData{
			Sender: EM.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
	}
	data.Recipient.ParseName(user.Name)

	var err error
	if data.ResetURL, err = EmailURL.ResetURL(user.GetPasswordResetToken()); err != nil {
		return err
	}

	msg, err := emails.PasswordResetRequestEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.EmailManager.Send(msg)
}

// SendPasswordResetSuccessEmail Send an email to a user to inform them that their password has been reset
func (h *Handler) SendPasswordResetSuccessEmail(user *User) error {
	data := emails.EmailData{
		Sender: EM.MustFromContact(),
		Recipient: sendgrid.Contact{
			Email: user.Email,
		},
	}
	data.Recipient.ParseName(user.Name)

	msg, err := emails.PasswordResetSuccessEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.EmailManager.Send(msg)
}

// SendOrgInvitationEmail sends an email inviting a user to join Datum and an existing organization
func (h *Handler) SendOrgInvitationEmail(i *emails.Invite) error {
	data := emails.InviteData{
		InviterName: i.Requestor,
		OrgName:     i.OrgName,
		EmailData: emails.EmailData{
			Sender: EM.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: i.Recipient,
			},
		},
	}

	var err error
	if data.InviteURL, err = EmailURL.InviteURL(i.Token); err != nil {
		return err
	}

	msg, err := emails.InviteEmail(data)
	if err != nil {
		return err
	}

	return h.EmailManager.Send(msg)
}

func (h *Handler) SendInvitationEmail(i *emails.Invite) error {
	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendOrgInvitationEmail(i)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
	); err != nil {
		return err
	}

	return nil
}
