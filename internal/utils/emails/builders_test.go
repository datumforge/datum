package emails_test

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/stretchr/testify/require"

	emails "github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

func TestEmailBuilders(t *testing.T) {
	setupMIMEDir(t)

	sender := sendgrid.Contact{
		FirstName: "George",
		LastName:  "Costanza",
		Email:     "gcostanza@example.com",
	}
	recipient := sendgrid.Contact{
		FirstName: "Jerry",
		LastName:  "Seinfeld",
		Email:     "littlejerry@kramerica.com",
	}
	data := emails.EmailData{
		Sender:    sender,
		Recipient: recipient,
	}

	welcomeData := emails.WelcomeData{
		EmailData:    data,
		FirstName:    "Sarah",
		LastName:     "Funkhouser",
		Email:        "sfunk@example.com",
		Organization: "Kids R Us",
		Domain:       "kidsrus.com",
	}
	mail, err := emails.WelcomeEmail(welcomeData)
	require.NoError(t, err, "expected no error when building welcome email")
	require.Equal(t, emails.WelcomeRE, mail.Subject, "expected welcome email subject to match")
	generateMIME(t, mail, "welcome.mime")

	verifyEmail := emails.VerifyEmailData{
		EmailData: data,
		FullName:  "Rusty Shackleford",
		VerifyURL: "https://datum.net/verify?token=1234567890",
	}
	mail, err = emails.VerifyEmail(verifyEmail)
	require.NoError(t, err, "expected no error when building verify email")
	require.Equal(t, emails.VerifyEmailRE, mail.Subject, "expected verify email subject to match")
	generateMIME(t, mail, "verify_email.mime")

	inviteData := emails.InviteData{
		EmailData:   data,
		Email:       "mitb@example.com",
		InviterName: "Matt Anderson",
		OrgName:     "Kids R Us",
		Role:        "Member",
		InviteURL:   "https://datum.net/invite?token=1234567890",
	}
	mail, err = emails.InviteEmail(inviteData)
	require.NoError(t, err, "expected no error when building invite email")
	require.Equal(t, fmt.Sprintf(emails.InviteRE, "Matt Anderson"), mail.Subject, "expected invite email subject to match")
	generateMIME(t, mail, "invite.mime")

	resetData := emails.ResetRequestData{
		EmailData: data,
		ResetURL:  "https://datum.net/reset?token=1234567890",
	}
	mail, err = emails.PasswordResetRequestEmail(resetData)
	require.NoError(t, err, "expected no error when building password reset email")
	require.Equal(t, emails.PasswordResetRequestRE, mail.Subject, "expected password reset email subject to match")
	generateMIME(t, mail, "reset_request.mime")

	mail, err = emails.PasswordResetSuccessEmail(data)
	require.NoError(t, err, "expected no error when building password reset success email")
	require.Equal(t, emails.PasswordResetSuccessRE, mail.Subject, "expected password reset success email subject to match")
	generateMIME(t, mail, "reset_success.mime")
}

func TestEmailData(t *testing.T) {
	sender := sendgrid.Contact{
		FirstName: "Art",
		LastName:  "Vandelay",
		Email:     "rvandelay@example.com",
	}
	recipient := sendgrid.Contact{
		FirstName: "Cosmo",
		LastName:  "Kramer",
		Email:     "ckramer@example.com",
	}
	data := emails.EmailData{
		Sender:    sender,
		Recipient: recipient,
	}

	// Email is not valid without a subject
	require.EqualError(t, data.Validate(), emails.ErrMissingSubject.Error(), "email subject should be required")

	// Email is not valid without a sender
	data.Subject = "Subject Line"
	data.Sender.Email = ""
	require.EqualError(t, data.Validate(), emails.ErrMissingSender.Error(), "email sender should be required")

	// Email is not valid without a recipient
	data.Sender.Email = sender.Email
	data.Recipient.Email = ""
	require.EqualError(t, data.Validate(), emails.ErrMissingRecipient.Error(), "email recipient should be required")

	// Successful validation
	data.Recipient.Email = recipient.Email
	require.NoError(t, data.Validate(), "expected no error when validating email data")
}

func TestLoadAttachment(t *testing.T) {
	// Add an attachment to a new email
	msg := mail.NewV3Mail()
	err := emails.LoadAttachment(msg, filepath.Join("testdata", "foo.zip"))
	require.NoError(t, err, "expected no error when adding attachment")

	// Ensure the attachment was added
	require.Len(t, msg.Attachments, 1, "expected one attachment")
	require.Equal(t, "foo.zip", msg.Attachments[0].Filename, "expected attachment to have correct filename")
	require.Equal(t, "application/zip", msg.Attachments[0].Type, "expected attachment to have correct type")
	require.Equal(t, "attachment", msg.Attachments[0].Disposition, "expected attachment to have correct disposition")

	// Ensure that we can decode the attachment data
	var data []byte
	data, err = base64.StdEncoding.DecodeString(msg.Attachments[0].Content)
	require.NoError(t, err, "expected no error when decoding attachment data")
	require.NotEmpty(t, data, "attachment has no data")
}

func TestAttachJSON(t *testing.T) {
	foo := map[string]string{"foo": "bar"}
	data, err := json.Marshal(foo)
	require.NoError(t, err, "expected no error when marshaling JSON")

	// Add an attachment to a new email
	msg := mail.NewV3Mail()
	err = emails.AttachJSON(msg, data, "foo.json")
	require.NoError(t, err, "expected no error when adding attachment")

	// Ensure the attachment was added
	require.Len(t, msg.Attachments, 1, "expected one attachment")
	require.Equal(t, "foo.json", msg.Attachments[0].Filename, "expected attachment to have correct filename")
	require.Equal(t, "application/json", msg.Attachments[0].Type, "expected attachment to have correct type")
	require.Equal(t, "attachment", msg.Attachments[0].Disposition, "expected attachment to have correct disposition")

	// Ensure that we can decode the attachment data
	var decoded []byte
	decoded, err = base64.StdEncoding.DecodeString(msg.Attachments[0].Content)
	require.NoError(t, err, "expected no error when decoding attachment data")

	actual := make(map[string]string)
	err = json.Unmarshal(decoded, &actual)

	require.NoError(t, err, "expected no error when unmarshaling JSON attachment")
	require.Equal(t, foo, actual, "expected JSON to match")
}

func TestAttachCSV(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 0, 10))
	w := csv.NewWriter(buf)

	w.Write([]string{"foo", "bar"}) // nolint: errcheck
	w.Flush()

	data := buf.Bytes()

	// Add an attachment to a new email
	msg := mail.NewV3Mail()
	err := emails.AttachCSV(msg, data, "foo.csv")
	require.NoError(t, err, "expected no error when adding attachment")

	// Ensure the attachment was added
	require.Len(t, msg.Attachments, 1, "expected one attachment")
	require.Equal(t, "foo.csv", msg.Attachments[0].Filename, "expected attachment to have correct filename")
	require.Equal(t, "text/csv", msg.Attachments[0].Type, "expected attachment to have correct type")
	require.Equal(t, "attachment", msg.Attachments[0].Disposition, "expected attachment to have correct disposition")

	// Ensure that we can decode the attachment data
	var decoded []byte
	decoded, err = base64.StdEncoding.DecodeString(msg.Attachments[0].Content)
	require.NoError(t, err, "expected no error when decoding attachment data")

	r := csv.NewReader(bytes.NewReader(decoded))
	actual, err := r.ReadAll()
	require.NoError(t, err, "expected no error when reading CSV attachment")
	require.Len(t, actual, 1)
	require.Equal(t, actual[0], []string{"foo", "bar"})
}
