package handlers

import (
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
)

func TestCreateUserInput(t *testing.T) {
	name := "Walter White"
	email := "ww@datum.net"

	testCases := []struct {
		testName string
		name     string
		email    string
		provider enums.AuthProvider
		expected ent.CreateUserInput
	}{
		{
			testName: "oauth provider - github",
			name:     name,
			email:    email,
			provider: enums.AuthProviderGitHub,
			expected: ent.CreateUserInput{
				FirstName:    "Walter",
				LastName:     "White",
				Email:        email,
				AuthProvider: &enums.AuthProviderGitHub,
				LastSeen:     lo.ToPtr(time.Now().UTC()),
			},
		},
		{
			testName: "oauth provider - google",
			name:     name,
			email:    email,
			provider: enums.AuthProviderGoogle,
			expected: ent.CreateUserInput{
				FirstName:    "Walter",
				LastName:     "White",
				Email:        email,
				AuthProvider: &enums.AuthProviderGoogle,
				LastSeen:     lo.ToPtr(time.Now().UTC()),
			},
		},
		{
			testName: "webauthn provider",
			name:     name,
			email:    email,
			provider: enums.AuthProviderWebauthn,
			expected: ent.CreateUserInput{
				FirstName:    "Walter",
				LastName:     "White",
				Email:        email,
				AuthProvider: &enums.AuthProviderWebauthn,
				LastSeen:     lo.ToPtr(time.Now().UTC()),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			input := createUserInput(tc.name, tc.email, tc.provider)
			assert.Equal(t, tc.expected.FirstName, input.FirstName)
			assert.Equal(t, tc.expected.LastName, input.LastName)
			assert.Equal(t, tc.expected.Email, input.Email)
			assert.Equal(t, tc.expected.AuthProvider, input.AuthProvider)
			assert.WithinDuration(t, *tc.expected.LastSeen, *input.LastSeen, 1*time.Minute) // allow for a reasonable drift while tests are running
		})
	}
}
