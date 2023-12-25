package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/mcuadros/go-defaults"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/utils/emails"
	"github.com/datumforge/datum/internal/utils/sendgrid"
)

func (h *Handler) RegisterHandler(ctx echo.Context) error {
	var (
		err error
		in  *RegisterRequest
		out *RegisterReply
	)

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return err
	}

	if err = in.Validate(); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, auth.ErrorResponse(err))

		return err
	}

	//	u := &User{}

	//	u.SetAgreement(in.AgreeToS, in.AgreePrivacy)

	input := generated.CreateUserInput{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  &in.Password,
		//		AgreeTos:     &in.AgreeToS,
		//		AgreePrivacy: &in.AgreePrivacy,
	}

	meowuser, err := h.DBClient.User.Create().SetInput(input).Save(ctx.Request().Context())
	if err != nil {
		return err
	}

	out = &RegisterReply{
		ID:      meowuser.ID,
		Email:   meowuser.Email,
		Message: "Welcome to Datum!",
	}

	//	user := &User{
	//		Email:     input.Email,
	//		FirstName: input.FirstName,
	//		LastName:  input.LastName,
	//	}

	//	if err := SendVerificationNoContact(); err != nil {
	//		return fmt.Errorf("it's your shitty code man")
	//	}

	//h.tasks.Queue(marionette.TaskFunc(func(ctx context.Context) error {
	//	return h.SendVerificationEmail(user)
	//}), marionette.WithRetries(3),
	//	marionette.WithBackoff(backoff.NewExponentialBackOff()),
	//	marionette.WithErrorf("could not send verification email to user %s", meowuser.ID),
	//)

	err = SendVerificationNoContact()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, out)
}

type RegisterReply struct {
	ID      string `json:"user_id"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	// PwCheck   string `json:"pwcheck"`
	// Organization string `json:"organization"`
	// Domain       string `json:"domain"`
	// AgreeToS     bool   `json:"terms_agreement"`
	// AgreePrivacy bool   `json:"privacy_agreement"`
}

// Validate the register request ensuring that the required fields are available and
// that the password is valid - an error is returned if the request is not correct. This
// method also performs some basic data cleanup, trimming whitespace.
func (r *RegisterRequest) Validate() error {
	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)
	r.Email = strings.TrimSpace(r.Email)
	r.Password = strings.TrimSpace(r.Password)
	// r.PwCheck = strings.TrimSpace(r.PwCheck)
	//	r.Organization = strings.TrimSpace(r.Organization)
	//	r.Domain = strings.ToLower(strings.TrimSpace(r.Domain))

	// Required for all requests
	switch {
	case r.Email == "":
		return auth.MissingField("email")
	case r.Password == "":
		return auth.MissingField("password")
	// case r.Password != r.PwCheck:
	// 	return auth.ErrPasswordMismatch
	case passwd.Strength(r.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
		//	case !r.AgreeToS:
		//		return auth.MissingField("terms_agreement")
		//	case !r.AgreePrivacy:
		//		return auth.MissingField("privacy_agreement")
	}

	return nil
}

func (u *User) SetAgreement(agreeTos, agreePrivacy bool) {
	u.AgreeToS = sql.NullBool{Valid: true, Bool: agreeTos}
	u.AgreePrivacy = sql.NullBool{Valid: true, Bool: agreePrivacy}
}

func SendVerificationNoContact() (err error) {
	sender := sendgrid.Contact{
		Email: "no-reply@datum.net",
	}
	recipient := sendgrid.Contact{
		FirstName: "Matt",
		LastName:  "Anderson",
		Email:     "manderson@datum.net",
	}
	//	data := emails.EmailData{
	//		Sender:    sender,
	//		Recipient: recipient,
	//	}

	data := emails.VerifyEmailData{
		EmailData: emails.EmailData{
			Sender:    sender,
			Recipient: recipient,
		},
		FullName:  "Matt Anderson",
		VerifyURL: "https://datum.net/token",
	}

	fmt.Sprintf("message data", data)
	// data.Recipient.ParseName(user.FirstName)

	//	token, err := tokens.NewVerificationToken(recipient.Email)
	//	if err != nil {
	//		return fmt.Errorf("HERE XXXXXXXX", token)
	//	}
	//
	//	signature, secret, err := token.Sign()
	//	if err != nil {
	//		return fmt.Errorf("HEREHEREHEREHERE", signature, secret)
	//	}

	var msg *mail.SGMailV3

	if msg, err = emails.VerifyEmail(data); err != nil {
		return fmt.Errorf("SHITBROKEHERE", err)
	}

	fmt.Sprintf("email", msg)

	// Send the email
	conf := &emails.Config{}
	defaults.SetDefaults(conf)

	em, err := emails.New(*conf)
	if err != nil {
		return err
	}

	return em.Send(msg)
}
