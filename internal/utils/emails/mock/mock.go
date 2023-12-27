package mock

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/sendgrid/rest"
	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

// EmailLog combines email data with a mutex to ensure that tests can safely access the mock concurrently for that email blastertons
type EmailLog struct {
	sync.Mutex
	Data [][]byte
}

// Emails contains all emails sent by the mock which tests can use to verify which emails were sent
var Emails EmailLog

// ResetEmailMock tests that send emails should call ResetEmailMock as part of their cleanup to ensure that other tests can depend on the state of the mock
func ResetEmailMock() {
	Emails.Lock()
	defer Emails.Unlock()
	Emails.Data = nil
}

// EmailMetadata makes it easier for tests to verify that the correct emails were sent without comparing the entire email body since that's 'spensive
type EmailMetadata struct {
	To          string
	From        string
	Subject     string
	Timestamp   time.Time
	Attachments int
}

// CheckEmails verifies that the provided email messages exist on the mock. This method
// should only be called from a test context and checks to make sure that the given emails
// were "sent" (not actually sent hehe)
func CheckEmails(t *testing.T, messages []*EmailMetadata) {
	var sentEmails []*sgmail.SGMailV3

	// Check total number of emails sent
	Emails.Lock()
	defer Emails.Unlock()
	require.Len(t, Emails.Data, len(messages), "incorrect number of emails sent")

	// Get emails from the mock
	for _, data := range Emails.Data {
		msg := &sgmail.SGMailV3{}
		require.NoError(t, json.Unmarshal(data, msg), "could not unmarshal email from mock")
		sentEmails = append(sentEmails, msg)
	}

	// Assert that all emails were sent
	for i, msg := range messages {
		expectedRecipient, err := mail.ParseAddress(msg.To)
		require.NoError(t, err, "could not parse expected recipient address")

		// Search for the sent email in the mock and check the metadata
		found := false

		for _, sent := range sentEmails {
			recipient, err := sendgrid.GetFirstRecipient(sent)
			require.NoError(t, err, "could not parse recipient address")

			if recipient == expectedRecipient.Address {
				found = true
				sender, err := mail.ParseAddress(msg.From)
				require.NoError(t, err, "could not parse expected sender address")
				require.Equal(t, sender.Address, sentEmails[i].From.Address)
				require.Equal(t, msg.Subject, sentEmails[i].Subject)

				// Disable attachments checking by setting EmailMetadata.Attachments = -1
				if msg.Attachments >= 0 {
					require.Len(t, sentEmails[i].Attachments, msg.Attachments)
				}

				break
			}
		}

		require.True(t, found, "email not sent for recipient %s", msg.To)
	}
}

// SendGridClient is the duct tape method which allows for local storage output
type SendGridClient struct {
	Storage string
}

// Send is responsible for sending an email using the SendGrid API
func (c *SendGridClient) Send(msg *sgmail.SGMailV3) (rep *rest.Response, err error) {
	// Marshal the email struct into those bytes
	data := sgmail.GetRequestBody(msg)
	if data == nil {
		return &rest.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "invalid email data",
		}, errors.New("could not marshal email") // nolint: goerr113
	}

	// Email needs to contain a From address
	if msg.From.Address == "" {
		return &rest.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "no From address",
		}, errors.New("requires From address") // nolint: goerr113
	}

	// Validate From address
	if _, err := sgmail.ParseEmail(msg.From.Address); err != nil {
		return &rest.Response{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("invalid From address: %s", msg.From.Address),
		}, err
	}

	// Email recipients are stored in Personalizations
	if len(msg.Personalizations) == 0 {
		return &rest.Response{
			StatusCode: http.StatusBadRequest,
			Body:       "no Personalization info",
		}, errors.New("requires Personalization info") // nolint: goerr113
	}

	var toAddress string

	for _, p := range msg.Personalizations {
		// Email needs to contain at least one To address
		if len(p.To) == 0 {
			return &rest.Response{
				StatusCode: http.StatusBadRequest,
				Body:       "no To addresses",
			}, errors.New("requires To address") // nolint: goerr113
		}

		for _, t := range p.To {
			// Validate To address
			if t.Address == "" {
				return &rest.Response{
					StatusCode: http.StatusBadRequest,
					Body:       "no To address",
				}, errors.New("empty To address") // nolint: goerr113
			}

			var mail *sgmail.Email

			if mail, err = sgmail.ParseEmail(t.Address); err != nil {
				return &rest.Response{
					StatusCode: http.StatusBadRequest,
					Body:       fmt.Sprintf("invalid To address: %s", t.Address),
				}, err
			}

			toAddress = mail.Address
		}
	}

	// "Send" the email
	Emails.Lock()
	Emails.Data = append(Emails.Data, data)
	Emails.Unlock()

	if c.Storage != "" {
		// Save the email to disk for manual inspection
		dir := filepath.Join(c.Storage, toAddress)
		if err = os.MkdirAll(dir, 0755); err != nil { // nolint: gomnd
			return &rest.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("could not create archive directory at %s", dir),
			}, err
		}

		// Generate unique filename to avoid overwriting
		ts := time.Now().Format(time.RFC3339)
		h := fnv.New32()
		h.Write(data)
		path := filepath.Join(dir, fmt.Sprintf("%s-%d.mim", ts, h.Sum32()))

		if err = sendgrid.WriteMIME(msg, path); err != nil {
			return &rest.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       fmt.Sprintf("could not archive email to %s", path),
			}, err
		}
	}

	return &rest.Response{StatusCode: http.StatusOK}, nil
}
