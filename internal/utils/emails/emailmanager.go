package emails

import (
	"github.com/sendgrid/rest"
	sg "github.com/sendgrid/sendgrid-go"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/utils/emails/mock"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// EmailManager allows a server to send rich emails using the SendGrid service
type EmailManager struct {
	conf   Config
	client SendGridClient
	URLConfig
}

// SendGridClient is an interface that can be implemented by live email clients to send
// real emails or by mock clients for testing
type SendGridClient interface {
	Send(email *sgmail.SGMailV3) (*rest.Response, error)
}

// New email manager with the specified configuration
func New(conf Config) (m *EmailManager, err error) {
	// conf.Valdate checks presence of admin, from email, and testing flags
	m = &EmailManager{conf: conf}

	if err := m.Validate(); err != nil {
		return nil, err
	}

	if conf.Testing {
		// there's an additional Storage field in the SendGridClient within mock
		m.client = &mock.SendGridClient{
			Storage: conf.Archive,
		}
	} else {
		if conf.SendGridAPIKey == "" {
			return nil, ErrFailedToCreateEmailClient
		}

		m.client = sg.NewSendClient(conf.SendGridAPIKey)
	}

	return m, nil
}

func (m *EmailManager) Send(message *sgmail.SGMailV3) (err error) {
	var rep *rest.Response

	if rep, err = m.client.Send(message); err != nil {
		return err
	}

	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return rout.HTTPErrorResponse(rep.Body)
	}

	return nil
}

// MustFromContact function is a helper function that returns the
// `sendgrid.Contact` for the `FromEmail` field in the `Config` struct
func (m *EmailManager) MustFromContact() sendgrid.Contact {
	contact, err := m.FromContact()
	if err != nil {
		panic(err)
	}

	return contact
}

// Enabled returns true if there is a SendGrid API key available
func (m *EmailManager) Enabled() bool {
	return m.conf.SendGridAPIKey != ""
}

// FromContact parses the FromEmail and returns a sendgrid contact
func (m *EmailManager) FromContact() (sendgrid.Contact, error) {
	return parseEmail(m.conf.FromEmail)
}

// AdminContact parses the AdminEmail and returns a sendgrid contact
func (m *EmailManager) AdminContact() (sendgrid.Contact, error) {
	return parseEmail(m.conf.AdminEmail)
}

// MustAdminContact is a helper function that returns the
// `sendgrid.Contact` for the `AdminEmail` field in the `Config` struct. It first calls the
// `AdminContact` function to parse the `AdminEmail` and return a `sendgrid.Contact`. If there is an
// error parsing the email, it will panic and throw an error. Otherwise, it will return the parsed
// `sendgrid.Contact`
func (m *EmailManager) MustAdminContact() sendgrid.Contact {
	contact, err := m.AdminContact()
	if err != nil {
		panic(err)
	}

	return contact
}

// Validate the from and admin emails are present if the SendGrid API is enabled
func (m *EmailManager) Validate() (err error) {
	if m.Enabled() {
		if m.conf.AdminEmail == "" || m.conf.FromEmail == "" {
			return ErrBothAdminAndFromRequired
		}

		if _, err = m.AdminContact(); err != nil {
			return ErrEmailNotParseable
		}

		if _, err = m.FromContact(); err != nil {
			return ErrAdminEmailNotParseable
		}

		if !m.conf.Testing && m.conf.Archive != "" {
			return ErrEmailArchiveOnlyInTestMode
		}
	}

	return nil
}

// SendOrgInvitationEmail sends an email inviting a user to join Datum and an existing organization
func (m *EmailManager) SendOrgInvitationEmail(i *Invite) error {
	data := InviteData{
		InviterName: i.Requestor,
		OrgName:     i.OrgName,
		EmailData: EmailData{
			Sender: m.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: i.Recipient,
			},
		},
	}

	var err error

	if data.InviteURL, err = m.URLConfig.InviteURL(i.Token); err != nil {
		return err
	}

	msg, err := InviteEmail(data)

	if err != nil {
		return err
	}

	return m.Send(msg)
}

// SendAddedtoOrgEmail sends an email notifying the user they've been added to an org
func (m *EmailManager) SendAddedtoOrgEmail(i *Invite) error {
	data := InviteData{
		InviterName: i.Requestor,
		OrgName:     i.OrgName,
		EmailData: EmailData{
			Sender: m.MustFromContact(),
			Recipient: sendgrid.Contact{
				Email: i.Recipient,
			},
		},
	}

	msg, err := InviteAccepted(data)

	if err != nil {
		return err
	}

	return m.Send(msg)
}
