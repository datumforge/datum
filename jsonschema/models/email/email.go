package email

type Config struct {
	EmailTemplate  EmailTemplate  `json:"emailTemplate"`
	EmailTemplates EmailTemplates `json:"emailTemplates"`
}

type EmailTemplate struct {
	// Html description: Template for HTML body
	Html string `json:"html"`
	// Subject description: Template for email subject header
	Subject string `json:"subject"`
	// Text description: Optional template for plain-text body. If not provided, a plain-text body will be automatically generated from the HTML template.
	Text string `json:"text"`
}

// EmailTemplates description: Configurable templates for some email types sent by Sourcegraph.
type EmailTemplates struct {
	// ResetPassword description: Email sent on password resets. Available template variables: {{.Host}}, {{.Username}}, {{.URL}}
	ResetPassword EmailTemplate `json:"resetPassword"`
	// SetPassword description: Email sent on account creation, if a password reset URL is created. Available template variables: {{.Host}}, {{.Username}}, {{.URL}}
	SetPassword EmailTemplate `json:"setPassword"`
}
