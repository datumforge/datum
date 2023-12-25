package emails

import (
	"fmt"
	"net/mail"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/utils/emails/mock"
)

// New email manager with the specified configuration
func New(conf Config) (m *EmailManager, err error) {
	m = &EmailManager{conf: conf}

	if conf.Testing {
		// there's an additional Storage field in the SendGridClient within mock
		m.client = &mock.SendGridClient{
			Storage: conf.Archive,
		}
	} else {
		if conf.APIKey == "" {
			return nil, ErrFailedToCreateEmailClient
		}
		m.client = sendgrid.NewSendClient(conf.APIKey)
	}

	// Parse the from and admin emails from the configuration
	if m.fromEmail, err = mail.ParseAddress(conf.FromEmail); err != nil {
		return nil, fmt.Errorf("could not parse 'from' email %q: %s", conf.FromEmail, err) // nolint: goerr113
	}

	return m, nil
}

// EmailManager allows a server to send rich emails using the SendGrid service
type EmailManager struct {
	conf      Config
	client    SendGridClient
	fromEmail *mail.Address
}

func (m *EmailManager) Send(message *sgmail.SGMailV3) (err error) {
	var rep *rest.Response

	if rep, err = m.client.Send(message); err != nil {
		return err
	}

	if rep.StatusCode < 200 || rep.StatusCode >= 300 {
		return auth.ErrorResponse(rep.Body)
	}

	return nil
}
