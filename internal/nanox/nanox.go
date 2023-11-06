// Package nanox provides a ID interface based on go-nanoid
package nanox

import (
	"github.com/jaevor/go-nanoid"
)

const (
	idLength = 21
)

// GetNewID returns an ID based on go-nanoid
func GetNewID() (string, error) {
	canonicID, err := nanoid.Standard(idLength)
	if err != nil {
		return "", err
	}

	return canonicID(), nil
}

// MustGetNewID returns an ID
func MustGetNewID() string {
	v, err := GetNewID()
	if err != nil {
		panic(err)
	}

	return v
}
