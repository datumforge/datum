package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/mattn/go-sqlite3"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
)

type RegisterReply struct {
	ID      string `json:"user_id"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (h *Handler) RegisterHandler(ctx echo.Context) error {
	var (
		err error
		in  *RegisterRequest
		out *RegisterReply
	)

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ErrorResponse(err))
	}

	if err = in.Validate(); err != nil {
		return ctx.JSON(http.StatusBadRequest, auth.ErrorResponse(err))
	}

	// create user
	input := generated.CreateUserInput{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  &in.Password,
	}

	tx, err := h.DBClient.Tx(ctx.Request().Context())
	if err != nil {
		h.Logger.Errorw("error starting transaction", "error", err)
		return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
	}

	meowuser, err := tx.User.Create().
		SetInput(input).
		Save(ctx.Request().Context())
	if err != nil {
		if err := tx.Rollback(); err != nil {
			h.Logger.Errorw("error rolling back transaction", "error", err)
			return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
		}

		if generated.IsConstraintError(err) {
			// TODO: this locks us in more closely to sqlite, should consider parsing the full error instead
			sqliteErr, ok := err.(sqlite3.Error)
			if ok {
				return newConstraintError(sqliteErr)
			}

			return ctx.JSON(http.StatusBadRequest, err)
		}

		return err
	}

	// create email verification token
	user := &User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
	}

	if err = user.CreateVerificationToken(); err != nil {
		h.Logger.Errorw("unable to create verification token", "error", err)
		return ctx.JSON(http.StatusInternalServerError, ErrProcessingRequest)
	}

	ttl, err := time.Parse(time.RFC3339Nano, user.EmailVerificationExpires.String)
	if err != nil {
		return err
	}

	meowtoken, err := tx.EmailVerificationToken.Create().
		SetOwnerID(meowuser.ID).
		SetToken(user.EmailVerificationToken.String).
		SetTTL(ttl).
		SetEmail(user.Email).
		SetSecret(user.EmailVerificationSecret).
		Save(ctx.Request().Context())
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse(err))
	}

	if err = h.SendVerificationEmail(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, auth.ErrorResponse(err))
	}

	// h.tasks.Queue(marionette.TaskFunc(func(ctx context.Context) error {
	//	return h.SendVerificationEmail(user)
	// }), marionette.WithRetries(3),
	//	marionette.WithBackoff(backoff.NewExponentialBackOff()),
	//	marionette.WithErrorf("could not send verification email to user %s", meowuser.ID),
	//)

	out = &RegisterReply{
		ID:      meowuser.ID,
		Email:   meowuser.Email,
		Message: "Welcome to Datum!",
		Token:   meowtoken.Token,
	}

	return ctx.JSON(http.StatusCreated, out)
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
	case r.Password == "":
		return auth.MissingField("password")
	case passwd.Strength(r.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
	}

	return nil
}

func (u *User) SetAgreement(agreeTos, agreePrivacy bool) {
	u.AgreeToS = sql.NullBool{Valid: true, Bool: agreeTos}
	u.AgreePrivacy = sql.NullBool{Valid: true, Bool: agreePrivacy}
}
