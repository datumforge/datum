package gravatar_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/datumforge/datum/internal/utils/gravatar"
)

func TestGravatar(t *testing.T) {
	email := "sfunk@datum.net"
	url := gravatar.New(email, nil)
	require.Equal(t, "https://www.gravatar.com/avatar/241e2e069287a36e9c902e55a3b6af8a?d=robohash&r=pg&s=80", url)
}

func TestHash(t *testing.T) {
	// Test case from: https://en.gravatar.com/site/implement/hash/
	input := "sfunk@datum.net"
	expected := "241e2e069287a36e9c902e55a3b6af8a"
	require.Equal(t, expected, gravatar.Hash(input))
}
