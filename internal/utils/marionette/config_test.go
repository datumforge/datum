package marionette_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/datumforge/datum/internal/utils/marionette"
)

func TestConfig(t *testing.T) {
	testCases := []struct {
		conf Config
		err  error
	}{
		{Config{}, ErrNoWorkers},
		{Config{Workers: 4}, ErrNoServerName},
		{Config{Workers: 4, ServerName: "marionette"}, nil},
	}

	for i, tc := range testCases {
		err := tc.conf.Validate()
		require.ErrorIs(t, err, tc.err, "test case %d failed", i)
	}
}
