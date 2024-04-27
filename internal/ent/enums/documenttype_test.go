package enums_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/datumforge/datum/internal/ent/enums"
)

func TestToDocumentType(t *testing.T) {
	testCases := []struct {
		input    string
		expected enums.DocumentType
	}{
		{
			input:    "roottemplate",
			expected: enums.RootTemplate,
		},
		{
			input:    "document",
			expected: enums.Document,
		},
		{
			input:    "UNKNOWN",
			expected: enums.DocumentTypeInvalid,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Convert %s to Region", tc.input), func(t *testing.T) {
			result := enums.ToDocumentType(tc.input)
			assert.Equal(t, tc.expected, *result)
		})
	}
}
