package handlers

import (
	"encoding/json"
	"net/http"

	echo "github.com/datumforge/echox"

	ent "github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
)

// ResendRequest contains fields for a resend email verification request
type ResendRequest struct {
	Email string `json:"email"`
}

// ResendReply holds the fields that are sent on a response to the `/resend` endpoint
type ResendReply struct {
	Message string `json:"message"`
}

// ResendEmail will resend an email verification email if the provided
// email exists
func (h *Handler) ResendEmail(ctx echo.Context) error {
	var in *ResendRequest

	out := &ResendReply{
		Message: "We've received your request to be resend an email to complete verification. If your email exists in our system, you should receive it shortly",
	}

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	if err := validateResendRequest(in); err != nil {
		h.Logger.Errorw("error validating request", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// set viewer context
	ctxWithToken := token.NewContextWithSignUpToken(ctx.Request().Context(), in.Email)

	entUser, err := h.getUserByEmail(ctxWithToken, in.Email)
	if err != nil {
		if ent.IsNotFound(err) {
			// return a 200 response even if user is not found to avoid
			// exposing confidential information
			return ctx.JSON(http.StatusOK, out)
		}

		h.Logger.Errorf("error retrieving user email", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// check to see if user is already confirmed
	if entUser.Edges.Setting.EmailConfirmed {
		out.Message = "email is already confirmed"

		return ctx.JSON(http.StatusOK, out)
	}

	viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(entUser.ID, true))

	// create email verification token
	user := &User{
		FirstName: entUser.FirstName,
		LastName:  entUser.LastName,
		Email:     entUser.Email,
		ID:        entUser.ID,
	}

	if _, err = h.storeAndSendEmailVerificationToken(viewerCtx, user); err != nil {
		h.Logger.Errorw("error storing token", "error", err)
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	return ctx.JSON(http.StatusOK, out)
}

// validateResendRequest validates the required fields are set in the user request
func validateResendRequest(req *ResendRequest) error {
	if req.Email == "" {
		return newMissingRequiredFieldError("email")
	}

	return nil
}
