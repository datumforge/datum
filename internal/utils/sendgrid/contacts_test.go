package sendgrid_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/sendgrid"
)

func TestContact(t *testing.T) {
	// Test constructing the full name
	contact := &sendgrid.Contact{}
	require.Equal(t, "", contact.FullName())

	contact.FirstName = "Rusty"
	require.Equal(t, "Rusty", contact.FullName())

	contact.LastName = "Shackleford"
	require.Equal(t, "Rusty Shackleford", contact.FullName())

	contact.FirstName = ""
	require.Equal(t, "Shackleford", contact.FullName())

	// Test parsing a full name into the contact
	contact.ParseName("")
	require.Empty(t, contact.FirstName)
	require.Empty(t, contact.LastName)

	contact.ParseName("Rusty")
	require.Equal(t, "Rusty", contact.FirstName)
	require.Empty(t, contact.LastName)

	contact.ParseName("Rusty Shackleford")
	require.Equal(t, "Rusty", contact.FirstName)
	require.Equal(t, "Shackleford", contact.LastName)

	contact.ParseName("Rusty Shackleford Smith")
	require.Equal(t, "Rusty", contact.FirstName)
	require.Equal(t, "Shackleford Smith", contact.LastName)

	// Test creating an email object from the contact
	contact = &sendgrid.Contact{
		FirstName: "Rusty",
		LastName:  "Shackleford",
		Email:     "Rusty@example.com",
	}
	email := contact.NewEmail()
	require.Equal(t, "Rusty Shackleford", email.Name)
	require.Equal(t, "Rusty@example.com", email.Address)
}
