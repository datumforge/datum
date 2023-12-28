package handlers

import (
	"net/url"

	"github.com/kelseyhightower/envconfig"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

func (h *Handler) SendVerificationEmail(user *User) error {
	// TODO: go back and configure with viper config instead of setting defaults
	conf := &emails.Config{}

	err := envconfig.Process("datum", conf)
	if err != nil {
		return err
	}

	h.sendgrid, err = emails.New(*conf)
	if err != nil {
		return err
	}

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
			Sender: conf.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
		},
		FullName: contact.FullName(),
	}

	// TODO: go back and configure with viper config instead of setting defaults
	urlConf := &URLConfig{}
	if err := envconfig.Process("datum", urlConf); err != nil {
		return nil
	}

	if data.VerifyURL, err = urlConf.VerifyURL(user.GetVerificationToken()); err != nil {
		return err
	}

	msg, err := emails.VerifyEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

// SendPasswordResetRequestEmail Send an email to a user to request them to reset their password
func (h *Handler) SendPasswordResetRequestEmail(user *User) error {
	data := emails.ResetRequestData{
		EmailData: emails.EmailData{
			Sender: h.SendGrid.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: user.Email,
			},
		},
	}
	data.Recipient.ParseName(user.Name)

	var err error
	if data.ResetURL, err = h.EmailURL.ResetURL(user.GetVerificationToken()); err != nil {
		return err
	}

	msg, err := emails.PasswordResetRequestEmail(data)
	if err != nil {
		return err
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

// SendPasswordResetSuccessEmail Send an email to a user to inform them that their password has been reset
func (h *Handler) SendPasswordResetSuccessEmail(user *User) error {
	data := emails.EmailData{
		Sender: h.SendGrid.MustFromContact(),
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
	return h.sendgrid.Send(msg)
}

// URLConfig for the datum registration
// TODO: move this to the same config setup as everything else
type URLConfig struct {
	Base   string `split_words:"true" default:"https://app.datum.net"`
	Verify string `split_words:"true" default:"/verify"`
	Invite string `split_words:"true" default:"/invite"`
	Reset  string `split_words:"true" default:"/reset"`
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

// VerifyURL Construct a verify URL from the token.
func (c URLConfig) VerifyURL(token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Verify, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// ResetURL Construct a reset URL from the token.
func (c URLConfig) ResetURL(token string) (string, error) {
	if token == "" {
		return "", newMissingRequiredFieldError("token")
	}

	base, _ := url.Parse(c.Base)

	url := base.ResolveReference(&url.URL{Path: c.Reset, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

func (h *Handler) createSendGridContact(contact *sendgrid.Contact) error { //nolint:unused
	if err := h.sendgrid.AddContact(contact); err != nil {
		h.Logger.Errorw("unable to add contact to sendgrid", "error", err)
		return err
	}

	return nil
}
