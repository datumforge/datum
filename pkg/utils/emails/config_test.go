package emails_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/pkg/utils/emails"
)

const adminEmail = "meow@mattthecat.com"

func TestSendGrid(t *testing.T) {
	em, err := emails.New(emails.Config{Testing: true})
	require.NoError(t, err)

	require.False(t, em.Enabled(), "sendgrid should be disabled when there is no API key")
	require.NoError(t, em.Validate(), "no validation error should be returned when sendgrid is disabled")

	em.SetSendGridAPIKey("SG.testing123")

	require.True(t, em.Enabled(), "sendgrid should be enabled when there is an API key")

	// FromEmail is required when enabled
	em.SetFromEmail("")
	em.SetAdminEmail(adminEmail)

	require.Error(t, em.Validate(), "expected from email to be required")

	// AdminEmail is required when enabled
	em.SetFromEmail(adminEmail)
	em.SetAdminEmail("")

	require.Error(t, em.Validate(), "expected admin email to be required")

	// Require parsable emails when enabled
	em.SetFromEmail("tacos")
	em.SetAdminEmail(adminEmail)

	require.Error(t, em.Validate())

	em.SetFromEmail(adminEmail)
	em.SetAdminEmail("tacos")

	require.Error(t, em.Validate())

	// Should be valid when enabled and emails are specified
	em.SetSendGridAPIKey("testing123")
	em.SetFromEmail(adminEmail)
	em.SetAdminEmail("sarahistheboss@example.com")

	require.NoError(t, em.Validate(), "expected configuration to be valid")

	// Archive is only supported in testing mode
	em.SetArchive("fixtures/emails")
	em.SetTesting(false)

	require.Error(t, em.Validate(), "expected error when archive is set in non-testing mode")

	em.SetTesting(true)

	require.NoError(t, em.Validate(), "expected configuration to be valid with archive in testing mode")
}
