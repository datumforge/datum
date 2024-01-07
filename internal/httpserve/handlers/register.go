package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/cenkalti/backoff/v4"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/utils/marionette"
)

// RegisterRequest holds the fields that should be included on a request to the `/register` endpoint
type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// RegisterReply holds the fields that are sent on a response to the `/register` endpoint
type RegisterReply struct {
	ID      string `json:"user_id"`
	Email   string `json:"email"`
	Message string `json:"message"`
	// TODO: remove this before go live, we shouldn't actually return the token here
	Token string `json:"token"`
}

// RegisterHandler handles the registration of a new datum user, creating the user, personal organization
// and sending an email verification to the email address in the request
// the user will not be able to authenticate until the email is verified
func (h *Handler) RegisterHandler(ctx echo.Context) error {
	var in *RegisterRequest

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	if err := in.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// create user
	input := generated.CreateUserInput{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  &in.Password,
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), in.Email)

	meowuser, err := h.createUser(ctxWithToken, input)
	if err != nil {
		h.Logger.Errorw("error creating new user", "error", err)

		if IsUniqueConstraintError(err) {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse("user already exists"))
		}

		if generated.IsValidationError(err) {
			field := err.(*generated.ValidationError).Name
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(fmt.Sprintf("%s was invalid", field)))
		}

		return err
	}

	// setup viewer context
	viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(meowuser.ID, true))

	// create email verification token
	user := &User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		ID:        meowuser.ID,
	}

	meowtoken, err := h.storeAndSendEmailVerificationToken(viewerCtx, user)
	if err != nil {
		h.Logger.Errorw("error storing token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	out := &RegisterReply{
		ID:      meowuser.ID,
		Email:   meowuser.Email,
		Message: "Welcome to Datum!",
		Token:   meowtoken.Token,
	}

	return ctx.JSON(http.StatusCreated, out)
}

func (h *Handler) storeAndSendEmailVerificationToken(ctx context.Context, user *User) (*generated.EmailVerificationToken, error) {
	if err := h.expireAllVerificationTokensUserByEmail(ctx, user.Email); err != nil {
		h.Logger.Errorw("error expiring existing tokens", "error", err)

		return nil, err
	}

	if err := user.CreateVerificationToken(); err != nil {
		h.Logger.Errorw("unable to create verification token", "error", err)

		return nil, err
	}

	meowtoken, err := h.createEmailVerificationToken(ctx, user)
	if err != nil {
		return nil, err
	}

	// send emails via TaskMan as to not create blocking operations in the server
	if err := h.TaskMan.Queue(marionette.TaskFunc(func(ctx context.Context) error {
		return h.SendVerificationEmail(user)
	}), marionette.WithRetries(3), // nolint: gomnd
		marionette.WithBackoff(backoff.NewExponentialBackOff()),
		marionette.WithErrorf("could not send verification email to user %s", user.ID),
	); err != nil {
		return nil, err
	}

	return meowtoken, nil
}

// Validate the register request ensuring that the required fields are available and
// that the password is valid - an error is returned if the request is not correct. This
// method also performs some basic data cleanup, trimming whitespace
func (r *RegisterRequest) Validate() error {
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)

	// Required for all requests
	switch {
	case r.Email == "":
		return auth.MissingField("email")
	case r.FirstName == "":
		return auth.MissingField("first name")
	case r.LastName == "":
		return auth.MissingField("last name")
	case r.Password == "":
		return auth.MissingField("password")
	case passwd.Strength(r.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
	}

	return nil
}
