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
			provider: enums.GitHub,
			expected: ent.CreateUserInput{
				FirstName:         "Walter",
				LastName:          "White",
				Email:             email,
				AuthProvider:      &enums.GitHub,
				LastSeen:          lo.ToPtr(time.Now().UTC()),
				IsWebauthnAllowed: lo.ToPtr(false),
				Oauth:             lo.ToPtr(true),
			},
		},
		{
			testName: "oauth provider - google",
			name:     name,
			email:    email,
			provider: enums.Google,
			expected: ent.CreateUserInput{
				FirstName:         "Walter",
				LastName:          "White",
				Email:             email,
				AuthProvider:      &enums.Google,
				LastSeen:          lo.ToPtr(time.Now().UTC()),
				IsWebauthnAllowed: lo.ToPtr(false),
				Oauth:             lo.ToPtr(true),
			},
		},
		{
			testName: "webauthn provider",
			name:     name,
			email:    email,
			provider: enums.Webauthn,
			expected: ent.CreateUserInput{
				FirstName:         "Walter",
				LastName:          "White",
				Email:             email,
				AuthProvider:      &enums.Webauthn,
				LastSeen:          lo.ToPtr(time.Now().UTC()),
				IsWebauthnAllowed: lo.ToPtr(true),
				Oauth:             lo.ToPtr(false),
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
			assert.Equal(t, *tc.expected.IsWebauthnAllowed, *input.IsWebauthnAllowed)
			assert.Equal(t, *tc.expected.Oauth, *input.Oauth)
		})
	}
}
