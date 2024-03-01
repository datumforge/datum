package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	echo "github.com/datumforge/echox"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/utils/ulids"
)

// WebauthnRegistrationRequest is the request to begin a webauthn login
type WebauthnRegistrationRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// WebauthnRegistrationResponse is the response to begin a webauthn login
type WebauthnRegistrationResponse struct {
	protocol.CredentialCreationResponse
	Email string `json:"email"`
}

// BeginWebauthnRegistration is the request to begin a webauthn login
func (h *Handler) BeginWebauthnRegistration(ctx echo.Context) error {
	var r WebauthnRegistrationRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&r); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	// to register a new passkey, the user needs to be logged in first
	// once the the passkey is added to the user's account, they can use it to login
	context := ctx.Request().Context()

	entUser, err := h.ConfirmOrCreateUser(context, r.Name, r.Email, enums.AuthProvider(strings.ToUpper("webauthn")))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	//	userCtx := viewer.NewContext(context, viewer.NewUserViewerFromSubject(context))

	user := &User{
		ID:    entUser.ID,
		Email: entUser.Email,
		Name:  entUser.FirstName + " " + entUser.LastName,
	}

	// options is the object that needs to be returned for the front end to open the creation dialog for the user
	options, session, err := h.WebAuthn.BeginRegistration(user,
		webauthn.WithResidentKeyRequirement(protocol.ResidentKeyRequirementRequired),
		webauthn.WithExclusions(user.CredentialExcludeList()),
	)
	if err != nil {
		return err
	}

	id := ulids.New().String()

	setSessionMap := map[string]any{}
	setSessionMap[id] = session

	setSessionMap[sessions.ExternalUserIDKey] = id
	setSessionMap[sessions.UsernameKey] = user.Name
	setSessionMap[sessions.UserTypeKey] = strings.ToUpper("webauthn")
	setSessionMap[sessions.EmailKey] = r.Email
	setSessionMap[sessions.UserIDKey] = user.ID

	if err := h.SessionConfig.SaveAndStoreSession(ctx.Request().Context(), ctx.Response().Writer, setSessionMap, user.ID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	ctx.SetCookie(&http.Cookie{
		Name:  "registration",
		Value: id,
		Path:  "/",
	})

	return ctx.JSON(http.StatusOK, options)
}

var Sessions = map[string]*webauthn.SessionData{}
var Users = map[string]*User{}

func InsertSession(id string, session *webauthn.SessionData) {
	Sessions[id] = session
}

func GetSession(id string) (*webauthn.SessionData, error) {
	s, ok := Sessions[id]
	if !ok {
		return nil, ErrNoAuthUser
	}

	return s, nil
}

// FinishWebauthnRegistration is the request to finish a webauthn login
func (h *Handler) FinishWebauthnRegistration(ctx echo.Context) error {
	var r WebauthnRegistrationResponse

	if err := json.NewDecoder(ctx.Request().Body).Decode(&r); err != nil {
		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(err))
	}

	id, err := ctx.Cookie("session_id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	response, err := protocol.ParseCredentialCreationResponseBody(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	user, err := GetUser(id.Value)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	session, err := GetSession(id.Value)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	credential, err := h.WebAuthn.CreateCredential(user, *session, response)
	if err != nil {
		return err
	}

	user.WebauthnCredentials = append(user.WebauthnCredentials, *credential)
	InsertUser(id.Value, user)

	return nil
}

// BeginWebauthnLogin is the request to begin a webauthn login
func (h *Handler) BeginWebauthnLogin(ctx echo.Context) error {
	credential, session, err := h.WebAuthn.BeginDiscoverableLogin()
	if err != nil {
		return err
	}

	id := ulids.New().String()

	InsertSession(id, session)
	ctx.SetCookie(&http.Cookie{
		Name:  "login_session_id",
		Value: id,
	})

	return ctx.JSON(http.StatusOK, credential)
}

// FinishWebauthnLogin is the request to finish a webauthn login
func (h *Handler) FinishWebauthnLogin(ctx echo.Context) error {
	id, err := ctx.Cookie("login_session_id")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	session, err := GetSession(id.Value)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	response, err := protocol.ParseCredentialRequestResponseBody(ctx.Request().Body)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	var loggedInUser *User = nil

	handler := func(rawID, userHandle []byte) (user webauthn.User, err error) {
		u, err := GetUserByID(userHandle)

		if err != nil {
			fmt.Printf("user not found: %v", err)
			return nil, err
		}

		loggedInUser = u

		return u, nil
	}

	_, err = h.WebAuthn.ValidateDiscoverableLogin(handler, *session, response)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, loggedInUser)
}

func InsertUser(id string, user *User) {
	Users[id] = user
}

func GetUser(name string) (*User, error) {
	u, ok := Users[name]
	if !ok {
		return nil, ErrNoAuthUser
	}

	return u, nil
}

func GetUserByID(id []byte) (*User, error) {
	for _, u := range Users {
		if string(u.ID) == string(id) {
			return u, nil
		}
	}

	return nil, ErrNoAuthUser
}

func (u *User) WebAuthnID() []byte {
	return []byte(u.ID)
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	if u.DisplayName != "" {
		return u.DisplayName
	}

	return u.Name
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.WebauthnCredentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}

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

func (h *Handler) ConfirmOrCreateUser(ctx context.Context, name, email string, provider enums.AuthProvider) (*ent.User, error) {
	// check if users exists
	entUser, err := h.getUserByEmail(ctx, email, provider)
	if err != nil {
		// if the user is not found, create now
		if ent.IsNotFound(err) {
			isWebAuthnAllowed := true
			lastSeen := time.Now()

			input := parseName(name)
			input.Email = email
			input.LastSeen = &lastSeen
			input.IsWebauthnAllowed = &isWebAuthnAllowed

			entUser, err = h.createUser(ctx, input)
			if err != nil {
				h.Logger.Errorw("error creating new user", "error", err)
				return nil, err
			}

			return entUser, nil
		}

		return nil, err
	}

	if err := h.updateUserLastSeen(ctx, entUser.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return nil, err
	}

	return entUser, nil
}
