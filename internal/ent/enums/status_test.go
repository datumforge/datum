package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/ent/enums"
)

func TestToUserStatus(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.UserStatus
	}{
		{
			input:    "active",
			expected: enums.Active,
		},
		{
			input:    "inactive",
			expected: enums.Inactive,
		},
		{
			input:    "DEACTIVATED",
			expected: enums.Deactivated,
		},
		{
			input:    "suspended",
			expected: enums.Suspended,
		},
		{
			input:    "UNKNOWN",
			expected: enums.StatusInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to UserStatus", tc.input), func(t *testing.T) {
			result := enums.ToUserStatus(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
