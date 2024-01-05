package handlers

import (
	"net/url"

	"github.com/kelseyhightower/envconfig"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// NewEmailManager is responsible for initializing and configuring the email manager used for sending emails
func (h *Handler) NewEmailManager() error {
	h.SendGridConfig = &emails.Config{}

	err := envconfig.Process("datum_email", h.SendGridConfig)
	if err != nil {
		return err
	}

	h.emailManager, err = emails.New(h.SendGridConfig)
	if err != nil {
		return err
	}

	h.EmailURL = &URLConfig{}
	if err := envconfig.Process("datum_email_url", h.EmailURL); err != nil {
		return err
	}

	return nil
}

// NewTestEmailManager is responsible for initializing and configuring the email manager used for sending emails but does not
// send emails, should be used in tests
func (h *Handler) NewTestEmailManager() error {
	h.SendGridConfig = &emails.Config{}

	err := envconfig.Process("datum", h.SendGridConfig)
	if err != nil {
		return err
	}

	h.SendGridConfig.Testing = true

	h.emailManager, err = emails.New(h.SendGridConfig)
	if err != nil {
		return err
	}

	h.EmailURL = &URLConfig{}
	if err := envconfig.Process("datum", h.EmailURL); err != nil {
		return err
	}

	return nil
}

func (h *Handler) SendVerificationEmail(user *User) error {
	contact := &sendgrid.Contact{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	// TODO: this current returns a 403 error, come back and figure out why
	// if err := h.createSendGridContact(contact); err != nil {
	// 	return err
	// }

	data := emails.VerifyEmailData{
		EmailData: emails.EmailData{
			Sender: h.SendGridConfig.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
		FullName: contact.FullName(),
	}

	var err error
	if data.VerifyURL, err = h.EmailURL.VerifyURL(user.GetVerificationToken()); err != nil {
		return err
	}

	msg, err := emails.VerifyEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.emailManager.Send(msg)
}

// SendPasswordResetRequestEmail Send an email to a user to request them to reset their password
// TODO: implement handler to use this and send password reset email
func (h *Handler) SendPasswordResetRequestEmail(user *User) error {
	data := emails.ResetRequestData{
		EmailData: emails.EmailData{
			Sender: h.SendGridConfig.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
	}
	data.Recipient.ParseName(user.Name)

	var err error
	if data.ResetURL, err = h.EmailURL.ResetURL(user.GetPasswordResetToken()); err != nil {
		return err
	}

	msg, err := emails.PasswordResetRequestEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.emailManager.Send(msg)
}

// SendPasswordResetSuccessEmail Send an email to a user to inform them that their password has been reset
func (h *Handler) SendPasswordResetSuccessEmail(user *User) error {
	data := emails.EmailData{
		Sender: h.SendGridConfig.MustFromContact(),
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
	return h.emailManager.Send(msg)
}

// URLConfig for the datum registration
type URLConfig struct {
	Base   string `split_words:"true" default:"https://app.datum.net"`
	Verify string `split_words:"true" default:"/v1/verify"`
	Invite string `split_words:"true" default:"/v1/invite"`
	Reset  string `split_words:"true" default:"/v1/reset"`
}

func (c URLConfig) Validate() error {
	if c.Base == "" {
		return newInvalidEmailConfigError("base URL")
	}

	if c.Invite == "" {
		return newInvalidEmailConfigError("invite path")
	}

	if c.Verify == "" {
		return newInvalidEmailConfigError("verify path")
	}

	if c.Reset == "" {
		return newInvalidEmailConfigError("reset path")
	}

	return nil
}

// InviteURL Construct an invite URL from the token.
func (c URLConfig) InviteURL(token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Invite, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// VerifyURL constructs a verify URL from the token.
func (c URLConfig) VerifyURL(token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Verify, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// ResetURL constructs a reset URL from the token.
func (c URLConfig) ResetURL(token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, _ := url.Parse(c.Base)

	url := base.ResolveReference(&url.URL{Path: c.Reset, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// createSendGridContact creates a contact in sendgrid to be used later
// TODO: this current returns a 403 error, come back and figure out why
func (h *Handler) createSendGridContact(contact *sendgrid.Contact) error { //nolint:unused
	if err := h.emailManager.AddContact(contact); err != nil {
		h.Logger.Errorw("unable to add contact to sendgrid", "error", err)
		return err
	}

	return nil
}
