package enums

import (
	"fmt"
	"io"
)

type AuthProvider string

var (
	// Credentials provider is when the user authenticates with a username and password
	Credentials AuthProvider = "CREDENTIALS"
	// Google oauth2 provider for authentication
	Google AuthProvider = "GOOGLE"
	// Github oauth2 provider for authentication
	GitHub AuthProvider = "GITHUB"
)

// Values returns a slice of strings that represents all the possible values of the AuthProvider enum.
// Possible default values are "CREDENTIALS", "GOOGLE", and "GITHUB"
func (AuthProvider) Values() (kinds []string) {
	for _, s := range []AuthProvider{Credentials, Google, GitHub} {
		kinds = append(kinds, string(s))
	}

	return
}

// String returns the AuthProvider as a string
func (r AuthProvider) String() string {
	return string(r)
}

// MarshalGQL implement the Marshaler interface for gqlgen
func (r AuthProvider) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(`"` + r.String() + `"`))
}

// UnmarshalGQL implement the Unmarshaler interface for gqlgen
func (r *AuthProvider) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong type for AuthProvider, got: %T", v) //nolint:goerr113
	}

	*r = AuthProvider(str)

	return nil
}
