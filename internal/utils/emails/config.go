package emails

import (
	"net/mail"
	"net/url"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// Config is a struct for sending emails via SendGrid and managing marketing contacts
type Config struct {
	// SendGridAPIKey is the sendgrid API key
	SendGridAPIKey string `split_words:"true" required:"false"`
	// FromEmail is the default email we'll send from and is safe to configure by default as our emails and domain are signed
	FromEmail string `split_words:"true" default:"no-reply@datum.net"`
	// Testing is a bool flag to indicate we shouldn't be sending live emails and defaults to true so needs to be specifically changed to send live emails
	Testing bool `split_words:"true" default:"false"`
	// Archive is only supported in testing mode and is what is tied through the mock to write out fixtures
	Archive string `split_words:"true" default:"fixtures/emails"`
	// DatumListID is the UUID sendgrid spits out when you create marketing lists
	DatumListID string `split_words:"true" required:"false" default:"f5459563-8a46-44ef-9066-e96124d30e52"`
	// AdminEmail is an internal group email configured within datum for email testing and visibility
	AdminEmail string `split_words:"true" default:"admins@datum.net"`
}

// URLConfig for the datum registration
type URLConfig struct {
	Base   string `split_words:"true" default:"https://api.datum.net"`
	Verify string `split_words:"true" default:"/v1/verify"`
	Invite string `split_words:"true" default:"/v1/invite"`
	Reset  string `split_words:"true" default:"/v1/reset-password"`
}

// SetSendGridAPIKey to provided key
func (m *EmailManager) SetSendGridAPIKey(key string) {
	m.conf.SendGridAPIKey = key
}

// GetSendGridAPIKey from the email manager config
func (m *EmailManager) GetSendGridAPIKey() string {
	return m.conf.SendGridAPIKey
}

// SetFromEmail to provided email
func (m *EmailManager) SetFromEmail(email string) {
	m.conf.FromEmail = email
}

// GetFromEmail from the email manager config
func (m *EmailManager) GetFromEmail() string {
	return m.conf.FromEmail
}

// SetAdminEmail to provided email
func (m *EmailManager) SetAdminEmail(email string) {
	m.conf.AdminEmail = email
}

// GetAdminEmail from the email manager config
func (m *EmailManager) GetAdminEmail() string {
	return m.conf.AdminEmail
}

// SetTesting to true/false to enable testing settings
func (m *EmailManager) SetTesting(testing bool) {
	m.conf.Testing = testing
}

// GetTesting from the email manager config
func (m *EmailManager) GetTesting() bool {
	return m.conf.Testing
}

// SetArchive location of email fixtures
func (m *EmailManager) SetArchive(archive string) {
	m.conf.Archive = archive
}

// GetArchive from the email manager config
func (m *EmailManager) GetArchive() string {
	return m.conf.Archive
}

// SetDatumListID to provided uuid
func (m *EmailManager) SetDatumListID(id string) {
	m.conf.DatumListID = id
}

// GetDatumListID from the email manager config
func (m *EmailManager) GetDatumListID() string {
	return m.conf.DatumListID
}

// parseEmail takes an email string as input and parses it into a `sendgrid.Contact`
// struct. It uses the `mail.ParseAddress` function from the `net/mail` package to parse the email
// address and name from the string. If the parsing is successful, it creates a `sendgrid.Contact`
// struct with the parsed email address and name (if available). If the parsing fails, it returns an
// error
func parseEmail(email string) (contact sendgrid.Contact, err error) {
	if email == "" {
		return contact, ErrEmailUnparseable
	}

	var addr *mail.Address

	if addr, err = mail.ParseAddress(email); err != nil {
		return contact, ErrEmailUnparseable
	}

	contact = sendgrid.Contact{
		Email: addr.Address,
	}
	contact.ParseName(addr.Name)

	return contact, nil
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
