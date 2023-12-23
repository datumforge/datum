package emails

import (
	"github.com/sendgrid/rest"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridClient is an interface that can be implemented by live email clients to send
// real emails or by mock clients for testing.
type SendGridClient interface {
	Send(email *sgmail.SGMailV3) (*rest.Response, error)
}
