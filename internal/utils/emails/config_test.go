package emails_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/emails"
)

const adminEmail = "meow@mattthecat.com"

func TestSendGrid(t *testing.T) {
	conf := &emails.Config{}
	em := &emails.EmailManager{}
	require.False(t, em.Enabled(), "sendgrid should be disabled when there is no API key")
	require.NoError(t, em.Validate(), "no validation error should be returned when sendgrid is disabled")

	conf.SendGridAPIKey = "SG.testing123"

	require.True(t, em.Enabled(), "sendgrid should be enabled when there is an API key")

	// FromEmail is required when enabled
	conf.FromEmail = ""
	conf.AdminEmail = adminEmail

	require.Error(t, em.Validate(), "expected from email to be required")

	// AdminEmail is required when enabled
	conf.FromEmail = adminEmail
	conf.AdminEmail = ""

	require.Error(t, em.Validate(), "expected admin email to be required")

	// Require parsable emails when enabled
	conf.FromEmail = "tacos"
	conf.AdminEmail = adminEmail

	require.Error(t, em.Validate())

	conf.FromEmail = adminEmail
	conf.AdminEmail = "tacos"

	require.Error(t, em.Validate())

	// Should be valid when enabled and emails are specified
	conf = &emails.Config{
		SendGridAPIKey: "testing123",
		FromEmail:      adminEmail,
		AdminEmail:     "sarahistheboss@example.com",
	}

	require.NoError(t, em.Validate(), "expected configuration to be valid")

	// Archive is only supported in testing mode
	conf.Archive = "fixtures/emails"

	require.Error(t, em.Validate(), "expected error when archive is set in non-testing mode")

	conf.Testing = true

	require.NoError(t, em.Validate(), "expected configuration to be valid with archive in testing mode")
}
