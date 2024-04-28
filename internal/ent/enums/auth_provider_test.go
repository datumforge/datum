package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/ent/enums"
)

func TestToAuthProvider(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.AuthProvider
	}{
		{
			input:    "credentials",
			expected: enums.AuthProviderCredentials,
		},
		{
			input:    "google",
			expected: enums.AuthProviderGoogle,
		},
		{
			input:    "github",
			expected: enums.AuthProviderGitHub,
		},
		{
			input:    "webauthn",
			expected: enums.AuthProviderWebauthn,
		},
		{
			input:    "UNKNOWN",
			expected: enums.AuthProviderInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to Region", tc.input), func(t *testing.T) {
			result := enums.ToAuthProvider(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
