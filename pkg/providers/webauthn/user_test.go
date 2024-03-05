package webauthn_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/go-webauthn/webauthn/protocol"
	gowebauthn "github.com/go-webauthn/webauthn/webauthn"

	"github.com/datumforge/datum/pkg/providers/webauthn"
)

func TestUserWebAuthnID(t *testing.T) {
	// Create a user instance
	user := &webauthn.User{
		ID: "exampleID",
	}

	// Call the WebAuthnID method
	webAuthnID := user.WebAuthnID()

	// Check if the returned value is correct
	expectedWebAuthnID := []byte("exampleID")
	assert.Equal(t, expectedWebAuthnID, webAuthnID)
}

func TestUserWebAuthnName(t *testing.T) {
	// Create a user instance
	user := &webauthn.User{
		Name: "example",
	}

	// Call the WebAuthnID method
	webAuthnName := user.WebAuthnName()

	// Check if the returned value is correct
	assert.Equal(t, "example", webAuthnName)
}

func TestWebAuthnDisplayName(t *testing.T) {
	testCases := []struct {
		testName    string
		name        string
		displayName string
		expected    string
	}{
		{
			testName:    "display name is set",
			name:        "Noah Kahan",
			displayName: "Noah",
			expected:    "Noah",
		},
		{
			testName:    "display name is empty",
			name:        "Noah Kahan",
			displayName: "",
			expected:    "Noah Kahan",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			// Create a user instance
			user := &webauthn.User{
				DisplayName: tc.displayName,
				Name:        tc.name,
			}

			result := user.WebAuthnDisplayName()
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestWebAuthnCredentials(t *testing.T) {
	// Create a user instance
	user := &webauthn.User{
		WebauthnCredentials: []gowebauthn.Credential{
			{
				ID: []byte("exampleID"),
			},
		},
	}

	// Call the WebAuthnCredentials method
	creds := user.WebAuthnCredentials()

	// Check if the returned value is correct
	assert.NotEmpty(t, creds)
	assert.Equal(t, user.WebauthnCredentials, creds)
}

func TestUserCredentialExcludeList(t *testing.T) {
	// Create a user instance
	user := &webauthn.User{
		WebauthnCredentials: []gowebauthn.Credential{
			{
				ID: []byte("exampleID"),
			},
		},
	}

	// Call the CredentialExcludeList method
	excludeList := user.CredentialExcludeList()

	// Check if the returned value is correct
	expectedExcludeList := []protocol.CredentialDescriptor{
		{
			Type:         protocol.PublicKeyCredentialType,
			CredentialID: []byte("exampleID"),
		},
	}
	assert.Equal(t, expectedExcludeList, excludeList)
}

func TestInsertSession(t *testing.T) {
	// Create a session data instance
	session := &gowebauthn.SessionData{
		Challenge: "challenge id",
		UserID:    []byte("01HR7NNG2PVCQCHYHCEQGP43CQ"),
	}

	// Call the InsertSession function
	webauthn.InsertSession("exampleID", session)

	// Check if the session is inserted correctly
	assert.Equal(t, session, webauthn.Sessions["exampleID"])
}

func TestGetSession(t *testing.T) {
	// Create a session data instance
	webauthn.Sessions = make(map[string]*gowebauthn.SessionData)
	webauthn.Sessions["01HR7NNG2PVCQCHYHCEQGP43CQ"] = &gowebauthn.SessionData{
		Challenge: "q5aQn9c231vZM9CnyBbz88KFAs4efMMaKdBx7OCUBeQ",
		UserID:    []byte("01HR5ZWH1H7G4PF3S8HAKXXFKH"),
	}

	testCases := []struct {
		name      string
		sessionID string
		expected  *gowebauthn.SessionData
		err       error
	}{
		{
			name:      "happy path",
			sessionID: "01HR7NNG2PVCQCHYHCEQGP43CQ",
			expected: &gowebauthn.SessionData{
				Challenge: "q5aQn9c231vZM9CnyBbz88KFAs4efMMaKdBx7OCUBeQ",
				UserID:    []byte("01HR5ZWH1H7G4PF3S8HAKXXFKH"),
			},
		},
		{
			name:      "session not found",
			sessionID: "NOTFOUNDID",
			expected:  nil,
			err:       webauthn.ErrSessionNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := webauthn.GetSession(tc.sessionID)
			if tc.err != nil {
				require.Error(t, err)
				assert.Equal(t, tc.err, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInsertUser(t *testing.T) {
	// Create a user data instance
	user := &webauthn.User{
		ID:          "01HR7NNG2PVCQCHYHCEQGP43CQ",
		DisplayName: "Eminem",
		Name:        "Marshall Mathers",
	}

	// Call the InsertUser function
	webauthn.InsertUser("01HR7NNG2PVCQCHYHCEQGP43CQ", user)

	// Check if the user is inserted correctly
	assert.Equal(t, user, webauthn.Users["01HR7NNG2PVCQCHYHCEQGP43CQ"])
}

func TestGetUser(t *testing.T) {
	// Create a session data instance
	webauthn.Users = make(map[string]*webauthn.User)
	webauthn.Users["Marshall Mathers"] = &webauthn.User{
		ID:          "01HR7NNG2PVCQCHYHCEQGP43CQ",
		DisplayName: "Eminem",
		Name:        "Marshall Mathers",
	}

	testCases := []struct {
		name     string
		userName string
		expected *webauthn.User
		err      error
	}{
		{
			name:     "happy path",
			userName: "Marshall Mathers",
			expected: &webauthn.User{
				ID:          "01HR7NNG2PVCQCHYHCEQGP43CQ",
				DisplayName: "Eminem",
				Name:        "Marshall Mathers",
			},
		},
		{
			name:     "user not found",
			userName: "Jayz",
			expected: nil,
			err:      webauthn.ErrUserNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := webauthn.GetUser(tc.userName)
			if tc.err != nil {
				require.Error(t, err)
				assert.Equal(t, tc.err, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	// Create a session data instance
	webauthn.Users = make(map[string]*webauthn.User)
	webauthn.Users["Marshall Mathers"] = &webauthn.User{
		ID:          "01HR7NNG2PVCQCHYHCEQGP43CQ",
		DisplayName: "Eminem",
		Name:        "Marshall Mathers",
	}

	testCases := []struct {
		name     string
		userID   string
		expected *webauthn.User
		err      error
	}{
		{
			name:   "happy path",
			userID: "01HR7NNG2PVCQCHYHCEQGP43CQ",
			expected: &webauthn.User{
				ID:          "01HR7NNG2PVCQCHYHCEQGP43CQ",
				DisplayName: "Eminem",
				Name:        "Marshall Mathers",
			},
		},
		{
			name:     "user not found",
			userID:   "NOTFOUNDID",
			expected: nil,
			err:      webauthn.ErrUserNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := webauthn.GetUserByID(tc.userID)
			if tc.err != nil {
				require.Error(t, err)
				assert.Equal(t, tc.err, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expected, result)
		})
	}
}
