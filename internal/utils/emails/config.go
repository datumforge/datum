package emails

import (
	"net/mail"

	"github.com/sendgrid/rest"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// Config is a struct for sending emails via SendGrid and managing marketing contacts
// TODO: migrate these configs into the default httpserve/config struct and add local yaml / viper configs
type Config struct {
	// SendGridAPIKey is the sendgrid API key
	SendGridAPIKey string `split_words:"true" required:"false"`
	// FromEmail is the default email we'll send from and is safe to configure by default as our emails and domain are signed
	FromEmail string `split_words:"true" default:"no-reply@datum.net"`
	// Testing is a bool flag to indicate we shouldn't be sending live emails and defaults to true so needs to be specifically changed to send live emails
	Testing bool `split_words:"true" default:"false"`
	// Archive is only supported in testing mode and is what is tied through the mock to write out fixtures
	Archive string `split_words:"true" default:"false"`
	// DatumListID is the UUID sendgrid spits out when you create marketing lists
	DatumListID string `split_words:"true" required:"false" default:"f5459563-8a46-44ef-9066-e96124d30e52"`
	// AdminEmail is an internal group email configured within datum for email testing and visibility
	AdminEmail string `split_words:"true" default:"admins@datum.net"`
}

// Email subject lines
const (
	WelcomeRE              = "Welcome to Datum!"
	VerifyEmailRE          = "Please verify your email address to login to Datum"
	InviteRE               = "Join Your Teammate %s on Datum!"
	PasswordResetRequestRE = "Datum Password Reset - Action Required"
	PasswordResetSuccessRE = "Datum Password Reset Confirmation"
)

// Validate the from and admin emails are present if the SendGrid API is enabled
func (c *Config) Validate() (err error) {
	if c.Enabled() {
		if c.AdminEmail == "" || c.FromEmail == "" {
			return ErrBothAdminAndFromRequired
		}

		if _, err = c.AdminContact(); err != nil {
			return ErrEmailNotParseable
		}

		if _, err = c.FromContact(); err != nil {
			return ErrAdminEmailNotParseable
		}

		if !c.Testing && c.Archive != "" {
			return ErrEmailArchiveOnlyInTestMode
		}
	}

	return nil
}

// Enabled returns true if there is a SendGrid API key available
func (c *Config) Enabled() bool {
	return c.SendGridAPIKey != ""
}

// FromContact parses the FromEmail and returns a sendgrid contact
func (c *Config) FromContact() (sendgrid.Contact, error) {
	return parseEmail(c.FromEmail)
}

// AdminContact parses the AdminEmail and returns a sendgrid contact
func (c Config) AdminContact() (sendgrid.Contact, error) {
	return parseEmail(c.AdminEmail)
}

// MustFromContact function is a helper function that returns the
// `sendgrid.Contact` for the `FromEmail` field in the `Config` struct
func (c *Config) MustFromContact() sendgrid.Contact {
	contact, err := c.FromContact()
	if err != nil {
		panic(err)
	}

	return contact
}

// MustAdminContact is a helper function that returns the
// `sendgrid.Contact` for the `AdminEmail` field in the `Config` struct. It first calls the
// `AdminContact` function to parse the `AdminEmail` and return a `sendgrid.Contact`. If there is an
// error parsing the email, it will panic and throw an error. Otherwise, it will return the parsed
// `sendgrid.Contact`
func (c *Config) MustAdminContact() sendgrid.Contact {
	contact, err := c.AdminContact()
	if err != nil {
		panic(err)
	}

	return contact
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

// SendGridClient is an interface that can be implemented by live email clients to send
// real emails or by mock clients for testing
type SendGridClient interface {
	Send(email *sgmail.SGMailV3) (*rest.Response, error)
}
