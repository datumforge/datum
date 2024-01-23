package emails

import (
	"github.com/sendgrid/rest"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

// EmailManager allows a server to send rich emails using the SendGrid service
type EmailManager struct {
	conf   Config
	client SendGridClient
}

// SendGridClient is an interface that can be implemented by live email clients to send
// real emails or by mock clients for testing
type SendGridClient interface {
	Send(email *sgmail.SGMailV3) (*rest.Response, error)
}

// New email manager with the specified configuration
func New(conf Config) (m *EmailManager, err error) {
	// conf.Valdate checks presence of admin, from email, and testing flags
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	m = &EmailManager{conf: conf}

	return m, nil
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
