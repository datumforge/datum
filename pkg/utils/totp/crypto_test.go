package totp

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	samples := "abcdefghijklmnopqrstuv"

	for i := 0; i <= 10; i++ {
		ln := 50
		random, err := String(ln, samples)
		require.NoError(t, err, "failed to generate random string")

		assert.Len(t, random, 50, "incorrect character count")

		for _, v := range random {
			s := string(v)
			assert.Contains(t, samples, s, "invalid character used in random string")
		}
	}
}

func TestStringB64(t *testing.T) {
	b64str, err := StringB64(50)
	require.NoError(t, err, "failed to generate random string")

	_, err = base64.StdEncoding.DecodeString(b64str)
	require.NoError(t, err, "failed to decode base64 encoded string")
}

func TestOTPHash(t *testing.T) {
	str := "the quick brown fox"
	hash, err := OTPHash(str)
	require.NoError(t, err, "error generating hash")

	assert.NotEqual(t, str, hash, "string not hashed")

	hash2, err := OTPHash(str)
	require.NoError(t, err, "error generating hash")

	assert.Equal(t, hash, hash2, "hashes do not match")
}
