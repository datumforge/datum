package sendgrid

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"

	sgmail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

// WriteMIME is used to write the contents of an email message in MIME format to a file.
// It takes a `SGMailV3` object, which represents the email message, and a file path as input
func WriteMIME(msg *sgmail.SGMailV3, path string) error {
	type EmailMetadata struct {
		From    string   `json:"from"`
		To      []string `json:"to"`
		Subject string   `json:"subject"`
	}

	// Create a buffer to store the MIME data
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	defer writer.Close()

	// Create the metadata header
	header := textproto.MIMEHeader{}
	header.Set(http.CanonicalHeaderKey("Content-Type"), "application/json")

	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}

	// Construct the metadata header
	metadata := EmailMetadata{
		From:    msg.From.Address,
		Subject: msg.Subject,
	}

	for _, p := range msg.Personalizations {
		for _, r := range p.To {
			metadata.To = append(metadata.To, r.Address)
		}
	}

	// Write the metadata header
	var b []byte

	if b, err = json.Marshal(metadata); err != nil {
		return err
	}

	if _, err = part.Write(b); err != nil {
		return err
	}

	// Write the email content sections
	for _, c := range msg.Content {
		header := textproto.MIMEHeader{}

		header.Set(http.CanonicalHeaderKey("Content-Type"), c.Type)
		part, err := writer.CreatePart(header)

		if err != nil {
			return err
		}

		if _, err = part.Write([]byte(c.Value)); err != nil {
			return err
		}
	}

	// Write the attachment sections
	for _, a := range msg.Attachments {
		header := textproto.MIMEHeader{}
		header.Set(http.CanonicalHeaderKey("Content-Type"), a.Type)
		header.Set(http.CanonicalHeaderKey("Content-Disposition"), a.Disposition)
		part, err := writer.CreatePart(header)

		if err != nil {
			return err
		}

		if _, err = part.Write([]byte(a.Content)); err != nil {
			return err
		}
	}

	if err = os.WriteFile(path, body.Bytes(), 0644); err != nil { // nolint: all
		return err
	}

	return nil
}

// GetFirstRecipient is used to extract the recipient email address from a `SGMailV3` object.
// It iterates through the personalizations and recipients of the email and returns the first
// recipient's email address. If no recipient is found, it returns an error
func GetFirstRecipient(msg *sgmail.SGMailV3) (recipient string, err error) {
	for _, p := range msg.Personalizations {
		for _, t := range p.To {
			recipient = t.Address

			return recipient, nil
		}
	}

	return "", ErrNoRecipientFound
}
