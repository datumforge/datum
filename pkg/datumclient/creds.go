package datumclient

import "github.com/datumforge/datum/pkg/rout"

// Credentials provides a basic interface for loading an access token
type Credentials interface {
	AccessToken() (string, error)
}

// A Token is just the JWT base64 encoded token string that is obtained from
// Datum either using the authtest server or from a login with the client
type Token string

// Token implements the credentials interface and performs limited validation
func (t Token) AccessToken() (string, error) {
	if string(t) == "" {
		return "", rout.ErrInvalidCredentials
	}

	return string(t), nil
}
