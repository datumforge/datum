package handlers

import (
	"fmt"
	"net/url"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

func (h *Handler) SendVerificationNoContact() (err error) {
	sender := sendgrid.Contact{
		Email: "no-reply@datum.net",
	}
	recipient := sendgrid.Contact{
		FirstName: "Matt",
		LastName:  "Anderson",
		Email:     "manderson@datum.net",
	}
	//	data := emails.EmailData{
	//		Sender:    sender,
	//		Recipient: recipient,
	//	}

	data := emails.VerifyEmailData{
		EmailData: emails.EmailData{
			Sender:    sender,
			Recipient: recipient,
		},
		FullName:  "Matt Anderson",
		VerifyURL: "https://datum.net/token",
	}
	// data.Recipient.ParseName(user.FirstName)

	//	token, err := tokens.NewVerificationToken(recipient.Email)
	//	if err != nil {
	//		return fmt.Errorf("HERE XXXXXXXX", token)
	//	}
	//
	//	signature, secret, err := token.Sign()
	//	if err != nil {
	//		return fmt.Errorf("HEREHEREHEREHERE", signature, secret)
	//	}

	var msg *mail.SGMailV3

	if msg, err = emails.VerifyEmail(data); err != nil {
		return fmt.Errorf("SHITBROKEHERE", err)
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

func (h *Handler) SendVerificationEmail(user *User) (err error) {
	if err := h.createSendgridContact(user); err != nil {
		return fmt.Errorf("shit went bad")
	}

	data := emails.VerifyEmailData{
		EmailData: emails.EmailData{
			Sender: h.SendGrid.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: user.Email,
			},
		},
		FullName: user.FirstName,
	}
	data.Recipient.ParseName(user.FirstName)

	if data.VerifyURL, err = h.EmailURL.VerifyURL(user.GetVerificationToken()); err != nil {
		return err
	}

	var msg *mail.SGMailV3

	if msg, err = emails.VerifyEmail(data); err != nil {
		return err
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

// SendPasswordResetRequestEmail Send an email to a user to request them to reset their password
func (h *Handler) SendPasswordResetRequestEmail(user *User) (err error) {
	data := emails.ResetRequestData{
		EmailData: emails.EmailData{
			Sender: h.SendGrid.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: user.Email,
			},
		},
	}
	data.Recipient.ParseName(user.Name)

	if data.ResetURL, err = h.EmailURL.ResetURL(user.GetVerificationToken()); err != nil {
		return err
	}

	var msg *mail.SGMailV3

	if msg, err = emails.PasswordResetRequestEmail(data); err != nil {
		return err
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

// SendPasswordResetSuccessEmail Send an email to a user to inform them that their password has been reset
func (h *Handler) SendPasswordResetSuccessEmail(user *User) (err error) {
	data := emails.EmailData{
		Sender: h.SendGrid.MustFromContact(),
		Recipient: sendgrid.Contact{
			Email: user.Email,
		},
	}
	data.Recipient.ParseName(user.Name)

	var msg *mail.SGMailV3

	if msg, err = emails.PasswordResetSuccessEmail(data); err != nil {
		return err
	}

	// Send the email
	return h.sendgrid.Send(msg)
}

// URLConfig is there a better way to do this?
type URLConfig struct {
	Base   string `split_words:"true" default:"https://api.datum.net"`
	Verify string `split_words:"true" default:"/verify"`
	Invite string `split_words:"true" default:"/invite"`
	Reset  string `split_words:"true" default:"/reset"`
}

func (c URLConfig) Validate() error {
	if c.Base == "" {
		return fmt.Errorf("invalid email url configuration: base URL is required") // nolint: goerr113
	} // nolint: goerr113

	if c.Invite == "" {
		return fmt.Errorf("invalid email url configuration: invite path is required") // nolint: goerr113
	} // nolint: goerr113

	if c.Verify == "" {
		return fmt.Errorf("invalid email url configuration: verify path is required") // nolint: goerr113
	} // nolint: goerr113

	if c.Reset == "" {
		return fmt.Errorf("invalid email url configuration: reset path is required") // nolint: goerr113
	} // nolint: goerr113

	return nil
}

// InviteURL Construct an invite URL from the token.
func (c URLConfig) InviteURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required") // nolint: goerr113
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Invite, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// VerifyURL Construct a verify URL from the token.
func (c URLConfig) VerifyURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required") // nolint: goerr113
	}

	base, _ := url.Parse(c.Base)
	url := base.ResolveReference(&url.URL{Path: c.Verify, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

// ResetURL Construct a reset URL from the token.
func (c URLConfig) ResetURL(token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("token is required") // nolint: goerr113
	}

	base, _ := url.Parse(c.Base)

	url := base.ResolveReference(&url.URL{Path: c.Reset, RawQuery: url.Values{"token": []string{token}}.Encode()})

	return url.String(), nil
}

func (h *Handler) createSendgridContact(user *User) error {
	contact := &sendgrid.Contact{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	if err := h.sendgrid.AddContact(contact); err != nil {
		return fmt.Errorf("could not add contact to sendgrid: %w", err)
	}

	return nil
}
