package emails

import (
	"embed"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

const (
	// Email templates must be provided in this directory and are loaded at compile time
	templatesDir = "templates"

	// Partials are included when rendering templates for composability and reuse - includes footer, header, etc.
	partialsDir = "partials"
)

var (
	//go:embed templates/*.html templates/*.txt templates/partials/*html
	files     embed.FS
	templates map[string]*template.Template
)

// Load templates when the package is imported
func init() {
	var (
		err           error
		templateFiles []fs.DirEntry
	)

	templates = make(map[string]*template.Template)

	if templateFiles, err = fs.ReadDir(files, templatesDir); err != nil {
		panic(err)
	}

	// Each template needs to be parsed independently to ensure that define directives
	// are not overwritten if they have the same name; e.g. to use the base template
	for _, file := range templateFiles {
		if file.IsDir() {
			continue
		}

		// Each template will be accessible by its base name in the global map
		patterns := make([]string, 0, 2) // nolint: gomnd
		patterns = append(patterns, filepath.Join(templatesDir, file.Name()))

		if filepath.Ext(file.Name()) == ".html" {
			patterns = append(patterns, filepath.Join(templatesDir, partialsDir, "*.html"))
		}

		templates[file.Name()] = template.Must(template.ParseFS(files, patterns...))
	}
}

const (
	UnknownDate = "unknown date"
	DateFormat  = "Monday, January 27, 1987"
)

// EmailData includes data fields that are common to all the email builders
type EmailData struct {
	Subject   string           `json:"-"`
	Sender    sendgrid.Contact `json:"-"`
	Recipient sendgrid.Contact `json:"-"`
}

// Validate that all required data is present to assemble a sendable email
func (e EmailData) Validate() error {
	switch {
	case e.Subject == "":
		return ErrMissingSubject
	case e.Sender.Email == "":
		return ErrMissingSender
	case e.Recipient.Email == "":
		return ErrMissingRecipient
	}

	return nil
}

// Build creates a new email from pre-rendered templates
func (e EmailData) Build(text, html string) (msg *mail.SGMailV3, err error) {
	if err = e.Validate(); err != nil {
		return nil, err
	}

	return mail.NewSingleEmail(
		e.Sender.NewEmail(),
		e.Subject,
		e.Recipient.NewEmail(),
		text,
		html,
	), nil
}

// TODO: implement handlers / sending logic for the below structs; things such as "Role" may change over time and require updates as the stack evolves.

// WelcomeData is used to complete the welcome email template
type WelcomeData struct {
	EmailData
	FirstName    string
	LastName     string
	Email        string
	Organization string
	Domain       string
}

// VerifyEmailData is used to complete the verify email template
type VerifyEmailData struct {
	EmailData
	FullName  string
	VerifyURL string
}

// InviteData is used to complete the invite email template
type InviteData struct {
	EmailData
	Email       string
	InviterName string
	OrgName     string
	Role        string
	InviteURL   string
}

// ResetRequestData is used to complete the password reset request email template
type ResetRequestData struct {
	EmailData
	ResetURL string
}

// WelcomeEmail creates a welcome email for a new user
func WelcomeEmail(data WelcomeData) (message *mail.SGMailV3, err error) {
	var text, html string

	if text, html, err = Render("welcome", data); err != nil {
		return nil, err
	}

	data.Subject = WelcomeRE

	return data.Build(text, html)
}

// VerifyEmail creates an email to verify a user's email address
func VerifyEmail(data VerifyEmailData) (message *mail.SGMailV3, err error) {
	var text, html string

	if text, html, err = Render("verify_email", data); err != nil {
		return nil, err
	}

	data.Subject = VerifyEmailRE

	return data.Build(text, html)
}

// InviteEmail creates an email to invite a user to join an organization
func InviteEmail(data InviteData) (message *mail.SGMailV3, err error) {
	var text, html string

	if text, html, err = Render("invite", data); err != nil {
		return nil, err
	}

	data.Subject = fmt.Sprintf(InviteRE, data.InviterName)

	return data.Build(text, html)
}

// PasswordResetRequestEmail creates a password reset email
func PasswordResetRequestEmail(data ResetRequestData) (message *mail.SGMailV3, err error) {
	var text, html string

	if text, html, err = Render("password_reset_request", data); err != nil {
		return nil, err
	}

	data.Subject = PasswordResetRequestRE

	return data.Build(text, html)
}

// PasswordResetSuccessEmail creates an email to send to users once their password has been reset
func PasswordResetSuccessEmail(data EmailData) (message *mail.SGMailV3, err error) {
	var text, html string

	if text, html, err = Render("password_reset_success", data); err != nil {
		return nil, err
	}

	data.Subject = PasswordResetSuccessRE

	return data.Build(text, html)
}

// Render returns the text and html executed templates for the specified name and data - while Sendgrid does have "rich" templates that could have been used instead, it seemed better not to explicitly reference SG ID's specific to Datum's account wherever possible so these files are rendered exports of templates created within that system and can be customized to suite purpose
func Render(name string, data interface{}) (text, html string, err error) {
	if text, err = render(name+".txt", data); err != nil {
		return "", "", err
	}

	if html, err = render(name+".html", data); err != nil {
		return "", "", err
	}

	return text, html, nil
}

func render(name string, data interface{}) (_ string, err error) {
	t, ok := templates[name]
	if !ok {
		return "", fmt.Errorf("could not find %q in templates", name) // nolint: goerr113
	}

	buf := &strings.Builder{}
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// AttachData onto an email as a file with the specified mimetype
func AttachData(message *mail.SGMailV3, data []byte, filename, mimetype string) error {
	// Encode the data to attach to the email
	encoded := base64.StdEncoding.EncodeToString(data)

	// Create the attachment
	attach := mail.NewAttachment()
	attach.SetContent(encoded)
	attach.SetType(mimetype)
	attach.SetFilename(filename)
	attach.SetDisposition("attachment")
	message.AddAttachment(attach)

	return nil
}

// LoadAttachment onto email from a file on disk - the intent here was the most "common types" we could envision sending out of the system rather than support every file type under the sun
func LoadAttachment(message *mail.SGMailV3, attachmentPath string) (err error) {
	var data []byte

	if data, err = os.ReadFile(attachmentPath); err != nil {
		return err
	}

	var mimetype string

	switch filepath.Ext(attachmentPath) {
	case ".zip":
		mimetype = "application/zip"
	case ".json":
		mimetype = "application/json"
	case ".csv":
		mimetype = "text/csv"
	case ".pdf":
		mimetype = "application/pdf"
	case ".tgz", ".gz":
		mimetype = "application/gzip"
	default:
		mimetype = http.DetectContentType(data)
	}

	// Create the attachment
	return AttachData(message, data, filepath.Base(attachmentPath), mimetype)
}

// AttachJSON by marshaling the specified data into human-readable data and encode and
// attach it to the email as a file
func AttachJSON(message *mail.SGMailV3, data []byte, filename string) (err error) {
	return AttachData(message, data, filename, "application/json")
}

// AttachCSV by encoding the csv data and attaching it to the email as a file - we don't have any CSV usecases today but likely will be generating reports and emailing them at some point in the future so adding it now
func AttachCSV(message *mail.SGMailV3, data []byte, filename string) (err error) {
	return AttachData(message, data, filename, "text/csv")
}
