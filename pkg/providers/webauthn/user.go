package webauthn

import (
	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	ID                  string
	FirstName           string
	LastName            string
	DisplayName         string
	Name                string
	Email               string
	WebauthnCredentials []webauthn.Credential `json:"-"`
}

var Sessions = map[string]*webauthn.SessionData{}
var Users = map[string]*User{}

// WebAuthnID is the user's webauthn ID
func (u *User) WebAuthnID() []byte {
	return []byte(u.ID)
}

// WebAuthnName is the user's webauthn name
func (u *User) WebAuthnName() string {
	return u.Name
}

// WebAuthnDisplayName is the user's webauthn display name
func (u *User) WebAuthnDisplayName() string {
	if u.DisplayName != "" {
		return u.DisplayName
	}

	return u.Name
}

// WebAuthnCredentials is the user's webauthn credentials
func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.WebauthnCredentials
}

// WebAuthnIcon is the user's webauthn icon
func (u *User) WebAuthnIcon() string {
	return ""
}

// CredentialExcludeList returns a list of credentials to exclude from the webauthn credential list
func (u *User) CredentialExcludeList() []protocol.CredentialDescriptor {
	credentialExcludeList := []protocol.CredentialDescriptor{}

	for _, cred := range u.WebauthnCredentials {
		descriptor := protocol.CredentialDescriptor{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: cred.ID,
		}
		credentialExcludeList = append(credentialExcludeList, descriptor)
	}

	return credentialExcludeList
}

// InsertSession adds the session to the Sessions map
func InsertSession(id string, session *webauthn.SessionData) {
	Sessions[id] = session
}

// GetSession returns the SessionData object for the given ID
func GetSession(id string) (*webauthn.SessionData, error) {
	s, ok := Sessions[id]
	if !ok {
		return nil, ErrSessionNotFound
	}

	return s, nil
}

// InsertUser adds the user to the Users map
func InsertUser(id string, user *User) {
	Users[id] = user
}

// GetUser returns the User object for the given name
func GetUser(name string) (*User, error) {
	u, ok := Users[name]
	if !ok {
		return nil, ErrUserNotFound
	}

	return u, nil
}

// GetUserByID returns the User object for the given ID
func GetUserByID(id string) (*User, error) {
	for _, u := range Users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, ErrUserNotFound
}
