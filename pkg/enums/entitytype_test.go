package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/pkg/enums"
)

func TestToEntityype(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.EntityType
	}{
		{
			input:    "organization",
			expected: enums.Organization,
		},
		{
			input:    "VENDOR",
			expected: enums.Vendor,
		},
		{
			input:    "UNKNOWN",
			expected: enums.EntityTypeInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to Entity Type", tc.input), func(t *testing.T) {
			result := enums.ToEntityType(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
