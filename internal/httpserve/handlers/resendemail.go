package handlers

import (
	"errors"
	"net/http"

	echo "github.com/datumforge/echox"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/datumforge/datum/internal/ent/enums"
	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/pkg/auth"
	"github.com/datumforge/datum/pkg/models"
	"github.com/datumforge/datum/pkg/rout"
)

// ResendEmail will resend an email verification email if the provided
// email exists
func (h *Handler) ResendEmail(ctx echo.Context) error {
	out := &models.ResendReply{
		Reply:   rout.Reply{Success: true},
		Message: "We've received your request to be resent an email to complete verification. Please check your email.",
	}

	var in models.ResendRequest
	if err := ctx.Bind(&in); err != nil {
		return h.BadRequest(ctx, err)
	}

	if err := validateResendRequest(in); err != nil {
		h.Logger.Errorw("error validating request", "error", err)

		return h.BadRequest(ctx, err)
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), in.Email)

	// email verifications only come to users that were created with username/password logins
	entUser, err := h.getUserByEmail(ctxWithToken, in.Email, enums.AuthProviderCredentials)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 200 response even if user is not found to avoid
			// exposing confidential information
			return h.Success(ctx, out)
		}

		h.Logger.Errorf("error retrieving user email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	// check to see if user is already confirmed
	if entUser.Edges.Setting.EmailConfirmed {
		out.Message = "email is already confirmed"

		return h.Success(ctx, out)
	}

	// setup user context
	userCtx := auth.AddAuthenticatedUserContext(ctx, &auth.AuthenticatedUser{
		SubjectID: entUser.ID,
	})

	// create email verification token
	user := &User{
		FirstName: entUser.FirstName,
		LastName:  entUser.LastName,
		Email:     entUser.Email,
		ID:        entUser.ID,
	}

	if _, err = h.storeAndSendEmailVerificationToken(userCtx, user); err != nil {
		h.Logger.Errorw("error storing token", "error", err)

		if errors.Is(err, ErrMaxAttempts) {
			return ctx.JSON(http.StatusTooManyRequests, rout.ErrorResponse(err))
		}

		return ctx.JSON(http.StatusInternalServerError, rout.ErrorResponse(ErrProcessingRequest))
	}

	return h.Success(ctx, out)
}

// validateResendRequest validates the required fields are set in the user request
func validateResendRequest(req models.ResendRequest) error {
	if req.Email == "" {
		return rout.NewMissingRequiredFieldError("email")
	}

	return nil
}

// BindResendEmail binds the resend email verification endpoint to the OpenAPI schema
func (h *Handler) BindResendEmailHandler() *openapi3.Operation {
	resendEmail := openapi3.NewOperation()
	resendEmail.Description = "ResendEmail accepts an email address via a POST request and always returns a 200 Status OK response, no matter the input or result of the processing. This is to ensure that no secure information is leaked from this unauthenticated endpoint. If the email address belongs to a user who has not been verified, another verification email is sent. If the post request contains an orgID and the user is invited to that organization but hasn't accepted the invite, then the invite is resent."
	resendEmail.OperationID = "ResendEmail"

	h.AddRequestBody("ResendEmail", models.ResendRequest{}, resendEmail)
	h.AddResponse("ResendReply", "success", models.ResendReply{}, resendEmail, http.StatusOK)
	resendEmail.AddResponse(http.StatusInternalServerError, internalServerError())
	resendEmail.AddResponse(http.StatusBadRequest, badRequest())

	return resendEmail
}
