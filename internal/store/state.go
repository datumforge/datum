package store

import (
	"crypto/rand"
	"encoding/base64"
)

// NewState generates cryptographically secure random state with base64 URL encoding
func NewState() (string, error) {
	cryptoRandBytes := make([]byte, 16) //nolint:gomnd

	_, err := rand.Read(cryptoRandBytes)
	if err != nil {
		return "", err
	}

	base64State := base64.URLEncoding.EncodeToString(cryptoRandBytes)

	return base64State, nil
}
