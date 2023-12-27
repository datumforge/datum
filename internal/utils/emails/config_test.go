package emails_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/emails"
)

func TestSendGrid(t *testing.T) {
	conf := &emails.Config{}
	require.False(t, conf.Enabled(), "sendgrid should be disabled when there is no API key")
	require.NoError(t, conf.Validate(), "no validation error should be returned when sendgrid is disabled")

	conf.APIKey = "SG.testing123"
	require.True(t, conf.Enabled(), "sendgrid should be enabled when there is an API key")

	// FromEmail is required when enabled
	conf.FromEmail = ""
	conf.AdminEmail = "meow@mattthecat.com"
	require.Error(t, conf.Validate(), "expected from email to be required")

	// AdminEmail is required when enabled
	conf.FromEmail = "meow@mattthecat.com"
	conf.AdminEmail = ""
	require.Error(t, conf.Validate(), "expected admin email to be required")

	// Require parsable emails when enabled
	conf.FromEmail = "tacos"
	conf.AdminEmail = "meow@mattthecat.com"
	require.Error(t, conf.Validate())

	conf.FromEmail = "meow@mattthecat.com"
	conf.AdminEmail = "tacos"
	require.Error(t, conf.Validate())

	// Should be valid when enabled and emails are specified
	conf = &emails.Config{
		APIKey:     "testing123",
		FromEmail:  "meow@mattthecat.com",
		AdminEmail: "sarahistheboss@example.com",
	}
	require.NoError(t, conf.Validate(), "expected configuration to be valid")

	// Archive is only supported in testing mode
	conf.Archive = "fixtures/emails"
	require.Error(t, conf.Validate(), "expected error when archive is set in non-testing mode")
	conf.Testing = true
	require.NoError(t, conf.Validate(), "expected configuration to be valid with archive in testing mode")
}
