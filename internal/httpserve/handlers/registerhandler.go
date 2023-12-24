package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
)

func (h *Handler) RegisterHandler(ctx echo.Context) error {
	var (
		err error
		in  *RegisterRequest
		// out *RegisterReply
	)

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return err
	}

	if err = in.Validate(); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, auth.ErrorResponse(err))
	}

	user := generated.CreateUserInput{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  &in.Password,
		// AgreePrivacy: in.AgreePrivacy,
		// AgreeToS:     in.AgreeToS,
	}

	meowerr := h.DBClient.User.Create().SetInput(user).Exec(ctx.Request().Context())

	if meowerr != nil {
		return meowerr
	}

	return ctx.JSON(http.StatusOK, Response{Message: "success"})
}

type RegisterReply struct {
	ID        ulid.ULID `json:"user_id"`
	OrgID     ulid.ULID `json:"org_id"`
	Email     string    `json:"email"`
	OrgName   string    `json:"org_name"`
	OrgDomain string    `json:"org_domain"`
	Message   string    `json:"message"`
	Role      string    `json:"role"`
	Created   string    `json:"created"`
}

type RegisterRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PwCheck      string `json:"pwcheck"`
	Organization string `json:"organization"`
	Domain       string `json:"domain"`
	AgreeToS     bool   `json:"terms_agreement"`
	AgreePrivacy bool   `json:"privacy_agreement"`
}

// Validate the register request ensuring that the required fields are available and
// that the password is valid - an error is returned if the request is not correct. This
// method also performs some basic data cleanup, trimming whitespace.
func (r *RegisterRequest) Validate() error {
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)
	r.PwCheck = strings.TrimSpace(r.PwCheck)
	r.Organization = strings.TrimSpace(r.Organization)
	r.Domain = strings.ToLower(strings.TrimSpace(r.Domain))

	// Required for all requests
	switch {
	case r.Email == "":
		return auth.MissingField("email")
	case r.Password == "":
		return auth.MissingField("password")
	case r.Password != r.PwCheck:
		return auth.ErrPasswordMismatch
	case passwd.Strength(r.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
	case !r.AgreeToS:
		return auth.MissingField("terms_agreement")
	case !r.AgreePrivacy:
		return auth.MissingField("privacy_agreement")
	}

	return nil
}
