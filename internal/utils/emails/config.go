package emails

import (
	"net/mail"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// Config is a struct for sending emails via SendGrid and managing marketing contacts
type Config struct {
	// APIKey is the sendgrid API key
	APIKey string `split_words:"true" required:"false"`
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

// Validate the from and admin emails are present if the SendGrid API is enabled
func (c Config) Validate() (err error) {
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
func (c Config) Enabled() bool {
	return c.APIKey != ""
}

// FromContact parses the FromEmail and returns a sendgrid contact
func (c Config) FromContact() (sendgrid.Contact, error) {
	return parseEmail(c.FromEmail)
}

// AdminContact parses the AdminEmail and returns a sendgrid contact
func (c Config) AdminContact() (sendgrid.Contact, error) {
	return parseEmail(c.AdminEmail)
}

func (c Config) MustFromContact() sendgrid.Contact {
	contact, err := c.FromContact()

	if err != nil {
		panic(err)
	}

	return contact
}

func (c Config) MustAdminContact() sendgrid.Contact {
	contact, err := c.AdminContact()

	if err != nil {
		panic(err)
	}

	return contact
}

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
