package handlers

import (
	"encoding/json"
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
)

const (
	webauthnProvider     = "WEBAUTHN"
	webauthnRegistration = "WEBAUTHN_REGISTRATION"
	webauthnLogin        = "WEBAUTHN_LOGIN"
)

// WebauthnRegistrationRequest is the request to begin a webauthn login
type WebauthnRegistrationRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// WebauthnRegistrationResponse is the response to begin a webauthn login
type WebauthnRegistrationResponse struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// BeginWebauthnRegistration is the request to begin a webauthn login
func (h *Handler) BeginWebauthnRegistration(ctx echo.Context) error {
	var r WebauthnRegistrationRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&r); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	ctxWithToken := token.NewContextWithOauthTooToken(ctx.Request().Context(), r.Email)

	// to register a new passkey, the user needs to be created + logged in first
	// once the the passkey is added to the user's account, they can use it to login
	// we treat this verify similar to the oauth or basic registration flow
	// user is created first, no credential method is set / they are unable to login until the credential flow is finished
	entUser, err := h.CheckAndCreateUser(ctxWithToken, r.Name, r.Email, enums.AuthProvider(webauthnProvider))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set context for remaining request based on logged in user
	userCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(entUser.ID, true))

	if err := h.addDefaultOrgToUserQuery(userCtx, entUser); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// create claims for verified user
	claims := createClaims(entUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set cookies for the user
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
		Name:  entUser.FirstName + " " + entUser.LastName,
	}

	// options is the object that needs to be returned for the front end to open the creation dialog for the user to create the passkey
	options, session, err := h.WebAuthn.BeginRegistration(user,
		webauthn.WithResidentKeyRequirement(protocol.ResidentKeyRequirementRequired),
		webauthn.WithExclusions(user.CredentialExcludeList()),
	)
	if err != nil {
		return err
	}

	// we have to set not just a regular datum session for the user but also capture the return of the webauthn session
	setSessionMap := map[string]any{}
	setSessionMap[sessions.WebAuthnKey] = session
	setSessionMap[sessions.UsernameKey] = r.Name
	setSessionMap[sessions.UserTypeKey] = webauthnRegistration
	setSessionMap[sessions.EmailKey] = r.Email
	setSessionMap[sessions.UserIDKey] = user.ID

	if err := h.SessionConfig.SaveAndStoreSession(userCtx, ctx.Response().Writer, setSessionMap, user.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, options)
}

// FinishWebauthnRegistration is the request to finish a webauthn registration - this is where we get the credential created by the user back
func (h *Handler) FinishWebauthnRegistration(ctx echo.Context) error {
	// lookup userID in cache to ensure cookie and tokens match
	session, err := h.SessionConfig.SessionManager.Get(ctx.Request(), h.SessionConfig.CookieConfig.Name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	sessionID := h.SessionConfig.SessionManager.GetSessionIDFromCookie(session)
	sessionData := h.SessionConfig.SessionManager.GetSessionDataFromCookie(session)

	userIDFromCookie := sessionData.(map[string]any)[sessions.UserIDKey]

	userID, err := h.SessionConfig.RedisStore.GetSession(ctx.Request().Context(), sessionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// ensure the user is the same as the one who started the registration
	if userIDFromCookie != userID {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// get user from the database
	entUser, err := h.getUserByID(ctx.Request().Context(), userID, enums.AuthProvider(webauthnProvider))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	userCtx := viewer.NewContext(ctx.Request().Context(), viewer.NewUserViewerFromID(entUser.ID, true))

	// follows https://www.w3.org/TR/webauthn/#sctn-registering-a-new-credential
	response, err := protocol.ParseCredentialCreationResponseBody(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	webauthnData := sessionData.(map[string]any)[sessions.WebAuthnKey]

	wd, ok := webauthnData.(webauthn.SessionData)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(ErrNoAuthUser))
	}

	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
		Name:  entUser.FirstName + " " + entUser.LastName,
	}

	credential, err := h.WebAuthn.CreateCredential(user, wd, response)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// save the credential to the database
	if err := h.addCredentialToUser(userCtx, entUser, *credential); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// create claims for verified user
	claims := createClaims(entUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// set cookies for the user
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	out := &WebauthnRegistrationResponse{
		Reply:   rout.Reply{Success: true},
		Message: "passkey successfully created",
	}

	return ctx.JSON(http.StatusOK, out)
}

// BeginWebauthnLogin is the request to begin a webauthn login
func (h *Handler) BeginWebauthnLogin(ctx echo.Context) error {
	credential, session, err := h.WebAuthn.BeginDiscoverableLogin()
	if err != nil {
		return err
	}

	setSessionMap := map[string]any{}
	setSessionMap[sessions.WebAuthnKey] = session
	setSessionMap[sessions.UserTypeKey] = webauthnLogin

	if err := h.SessionConfig.SaveAndStoreSession(ctx.Request().Context(), ctx.Response().Writer, setSessionMap, ""); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, credential)
}

// FinishWebauthnLogin is the request to finish a webauthn login
func (h *Handler) FinishWebauthnLogin(ctx echo.Context) error {
	session, err := h.SessionConfig.SessionManager.Get(ctx.Request(), h.SessionConfig.CookieConfig.Name)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	sessionData := h.SessionConfig.SessionManager.GetSessionDataFromCookie(session)

	webauthnData := sessionData.(map[string]any)[sessions.WebAuthnKey]

	wd, ok := webauthnData.(webauthn.SessionData)
	if !ok {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(ErrNoAuthUser))
	}

	response, err := protocol.ParseCredentialRequestResponseBody(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	handler := func(rawID, userHandle []byte) (user webauthn.User, err error) {
		u, err := h.getUserByID(ctx.Request().Context(), string(userHandle), enums.AuthProvider(webauthnProvider))
		if err != nil {
			return nil, err
		}

		authnUser := &User{
			ID:                  u.ID,
			Email:               u.Email,
			Name:                u.FirstName + " " + u.LastName,
			WebauthnCredentials: []webauthn.Credential{},
		}

		for _, cred := range u.Edges.Webauthn {
			authnCred := webauthn.Credential{
				ID:              cred.CredentialID,
				PublicKey:       cred.PublicKey,
				AttestationType: cred.AttestationType,
			}

			for _, t := range cred.Transports {
				authnCred.Transport = append(authnCred.Transport, protocol.AuthenticatorTransport(t))
			}

			authnUser.WebauthnCredentials = append(authnUser.WebauthnCredentials, authnCred)
		}

		return authnUser, nil
	}

	if _, err = h.WebAuthn.ValidateDiscoverableLogin(handler, wd, response); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, rout.Reply{Success: true})
}

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

var Sessions = map[string]*webauthn.SessionData{}
var Users = map[string]*User{}

// InsertSession adds the session to the Sessions map
func InsertSession(id string, session *webauthn.SessionData) {
	Sessions[id] = session
}

// GetSession returns the SessionData object for the given ID
func GetSession(id string) (*webauthn.SessionData, error) {
	s, ok := Sessions[id]
	if !ok {
		return nil, ErrNoAuthUser
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
		return nil, ErrNoAuthUser
	}

	return u, nil
}

// GetUserByID returns the User object for the given ID
func GetUserByID(id []byte) (*User, error) {
	for _, u := range Users {
		if string(u.ID) == string(id) {
			return u, nil
		}
	}

	return nil, ErrNoAuthUser
}
