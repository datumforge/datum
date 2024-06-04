package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/golang-jwt/jwt/v5"
	ph "github.com/posthog/posthog-go"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/passwd"
	"github.com/datumforge/datum/pkg/rout"
	"github.com/datumforge/datum/pkg/sessions"
	"github.com/datumforge/datum/pkg/tokens"
)

// LoginHandler validates the user credentials and returns a valid cookie
// this handler only supports username password login
func (h *Handler) LoginHandler(ctx echo.Context) error {
	user, err := h.verifyUserPassword(ctx)
	if err != nil {
		return h.BadRequest(ctx, err)
	}

	// set context for remaining request based on logged in user
	userCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		SubjectID: user.ID,
	})

	if err := h.addDefaultOrgToUserQuery(userCtx, user); err != nil {
		return h.InternalServerError(ctx, err)
	}

	claims := createClaims(user)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		return h.InternalServerError(ctx, err)
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, user.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return h.InternalServerError(ctx, err)
	}

	if err := h.updateUserLastSeen(userCtx, user.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return h.InternalServerError(ctx, err)
	}

	props := ph.NewProperties().
		Set("user_id", user.ID).
		Set("email", user.Email).
		Set("organization_id", user.Edges.Setting.Edges.DefaultOrg.ID). // user is logged into their default org
		Set("auth_provider", user.AuthProvider)

	h.AnalyticsClient.Event("user_authenticated", props)
	h.AnalyticsClient.UserProperties(user.ID, props)

	// return the session value for the UI to use
	// the UI will need to set the cookie because authentication is handled
	// server side
	s, err := sessions.SessionToken(ctx.Request().Context())
	if err != nil {
		return h.InternalServerError(ctx, err)
	}

	out := models.LoginReply{
		Reply:        rout.Reply{Success: true},
		Message:      "success",
		AccessToken:  access,
		RefreshToken: refresh,
		Session:      s,
		ExpiresIn:    claims.ExpiresAt.Unix(),
	}

	return h.Success(ctx, out)
}

// createClaims creates the claims for the JWT token using the mapping ids for the user and organization
func createClaims(u *generated.User) *tokens.Claims {
	orgID := ""
	if u.Edges.Setting.Edges.DefaultOrg != nil {
		orgID = u.Edges.Setting.Edges.DefaultOrg.MappingID
	}

	return &tokens.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: u.MappingID,
		},
		UserID: u.MappingID,
		OrgID:  orgID,
	}
}

// verifyUserPassword verifies the username and password are valid
func (h *Handler) verifyUserPassword(ctx echo.Context) (*generated.User, error) {
	var l models.LoginRequest
	if err := ctx.Bind(&l); err != nil {
		return nil, ErrBadRequest
	}

	if l.Username == "" || l.Password == "" {
		return nil, ErrMissingRequiredFields
	}

	// check user in the database, username == email and ensure only one record is returned
	user, err := h.getUserByEmail(ctx.Request().Context(), l.Username, enums.AuthProviderCredentials)
	if err != nil {
		return nil, ErrNoAuthUser
	}

	if user.Edges.Setting.Status != "ACTIVE" {
		return nil, ErrNoAuthUser
	}

	// verify the password is correct
	valid, err := passwd.VerifyDerivedKey(*user.Password, l.Password)
	if err != nil || !valid {
		return nil, ErrInvalidCredentials
	}

	// verify email is verified
	if !user.Edges.Setting.EmailConfirmed {
		return nil, ErrUnverifiedUser
	}

	return user, nil
}

// BindLoginHandler binds the login request to the OpenAPI schema
func (h *Handler) BindLoginHandler() *openapi3.Operation {
	login := openapi3.NewOperation()
	login.Description = "Login is oriented towards human users who use their email and password for authentication (whereas authenticate is used for machine access using API keys). Login verifies the password submitted for the user is correct by looking up the user by email and using the argon2 derived key verification process to confirm the password matches. Upon authentication an access token and a refresh token with the authorized claims of the user are returned. The user can use the access token to authenticate to Datum systems. The access token has an expiration and the refresh token can be used with the refresh endpoint to get a new access token without the user having to log in again. The refresh token overlaps with the access token to provide a seamless authentication experience and the user can refresh their access token so long as the refresh token is valid"
	login.OperationID = "LoginHandler"

	h.AddRequestBody("LoginRequest", models.ExampleLoginSuccessRequest, login)
	h.AddResponse("LoginReply", "success", models.ExampleLoginSuccessResponse, login, http.StatusOK)
	login.AddResponse(http.StatusInternalServerError, internalServerError())
	login.AddResponse(http.StatusBadRequest, badRequest())

	return login
}
