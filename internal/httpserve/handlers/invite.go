package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
	"github.com/datumforge/datum/internal/passwd"
	"github.com/datumforge/datum/internal/tokens"
)

// OrganizationInviteAccept is responsible for handling the invitation of a user to an organization.
// It receives a request with the user's invitation details, validates the request,
// and creates organization membership for the user
// On success, it returns a response with the organization information
func (h *Handler) OrganizationInviteAccept(ctx echo.Context) error {
	// parse the token out of the context
	inv := &Invite{
		Token: ctx.QueryParam("token"),
	}

	var in *InviteRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	// set the received input values on the invite
	inv.Password = in.Password
	inv.FirstName = in.FirstName
	inv.LastName = in.LastName

	// ensure we've collected everything we need to confirm the token and create the user
	if err := inv.validateInviteRequest(); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// set the initial context based on the token
	ctxWithToken := token.NewContextWithOrgInviteToken(ctx.Request().Context(), inv.Token)

	// fetch the recipient and org owner based on token
	invitedUser, err := h.getUserByInviteToken(ctxWithToken, inv.Token)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving invite token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	// string to ulid so we can match the token input
	oid, err := ulid.Parse(invitedUser.OwnerID)
	if err != nil {
		return err
	}

	// construct the invite details but set email to the original recipient, and the joining organization ID as the current owner of the invitation
	invite := &Invite{
		Email:     invitedUser.Recipient,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		DestOrgID: oid,
		Role:      invitedUser.Role,
	}

	// set tokens for request
	if err := invite.setOrgInviteTokens(invitedUser, inv.Token); err != nil {
		h.Logger.Errorw("unable to set invite token for request", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// reconstruct the token based on recipient & owning organization so we can compare it to the one were receiving
	t := &tokens.OrgInviteToken{
		Email: invitedUser.Recipient,
		OrgID: oid,
	}

	// check and ensure the token has not expired
	if t.ExpiresAt, err = invite.GetInviteExpires(); err != nil {
		h.Logger.Errorw("unable to parse expiration", "error", err)

		return ctx.JSON(http.StatusInternalServerError, tokens.ErrTokenExpired)
	}

	// Verify the token is valid with the stored secret
	if err = t.Verify(invite.GetInviteToken(), invite.Secret); err != nil {
		if errors.Is(err, tokens.ErrTokenExpired) {
			if err := updateInviteStatusExpired(ctxWithToken, invitedUser); err != nil {
				return err
			}

			out := &Response{
				Message: "invite token is expired, you will need to re-request an invite",
			}

			return ctx.JSON(http.StatusBadRequest, out)
		}

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// create the user, but don't allow email to be different than the one listed as the recipient in the original invitation
	createdUser, err := h.createUser(ctxWithToken, generated.CreateUserInput{
		Email:     invitedUser.Recipient,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Password:  &in.Password,
	})
	if err != nil {
		h.Logger.Errorw("error creating new user", "error", err)

		// this would only hit if the user registered normally after having already received an invite
		if !IsUniqueConstraintError(err) {
			if generated.IsValidationError(err) {
				field := err.(*generated.ValidationError).Name
				return ctx.JSON(http.StatusBadRequest, ErrorResponse(fmt.Sprintf("%s was invalid", field)))
			}

			return err
		}
	}

	// if the user was created, use that user as the createdUser
	if IsUniqueConstraintError(err) {
		createdUser, err = h.getUserByEmail(ctxWithToken, invitedUser.Recipient, enums.Credentials)
		if err != nil {
			return err
		}
	}

	// set new viewer context with user id
	viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(createdUser.ID, true))

	// don't require an additional email verification since the invite was sent to an email
	if err := h.setEmailConfirmed(viewerCtx, createdUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	if err := updateInviteStatusAccepted(viewerCtx, invitedUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	claims := createClaims(createdUser)

	access, refresh, err := h.TM.CreateTokenPair(claims)
	if err != nil {
		h.Logger.Errorw("error creating token pair", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// set cookies on request with the access and refresh token
	auth.SetAuthCookies(ctx.Response().Writer, access, refresh)

	// set sessions in response
	if err := h.SessionConfig.CreateAndStoreSession(ctx, createdUser.ID); err != nil {
		h.Logger.Errorw("unable to save session", "error", err)

		return err
	}

	if err := h.updateUserLastSeen(viewerCtx, createdUser.ID); err != nil {
		h.Logger.Errorw("unable to update last seen", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
	}

	// reply with the relevant details
	out := &InviteReply{
		ID:          createdUser.ID,
		Email:       createdUser.Email,
		JoinedOrgID: invitedUser.OwnerID,
		Role:        string(invitedUser.Role),
		Message:     "Welcome to your new organization!",
	}

	return ctx.JSON(http.StatusCreated, out)
}

// validateInviteRequest validates the required fields are set in the user request
func (i *Invite) validateInviteRequest() error {
	i.FirstName = strings.TrimSpace(i.FirstName)
	i.LastName = strings.TrimSpace(i.LastName)
	i.Password = strings.TrimSpace(i.Password)

	switch {
	case i.Token == "":
		return newMissingRequiredFieldError("token")
	case i.FirstName == "":
		return auth.MissingField("first name")
	case i.LastName == "":
		return auth.MissingField("last name")
	case i.Password == "":
		return auth.MissingField("password")
	case passwd.Strength(i.Password) < passwd.Moderate:
		return auth.ErrPasswordTooWeak
	}

	return nil
}

// InviteReply holds the fields that are sent on a response to an accepted invitation
type InviteReply struct {
	ID          string `json:"user_id"`
	Email       string `json:"email"`
	Message     string `json:"message"`
	JoinedOrgID string `json:"joined_org_id"`
	Role        string `json:"role"`
}

// InviteRequest holds the additional input from the user collected during acceptance
type InviteRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// Invite holds the Token, InviteToken references, and the additional user input to //
// complete acceptance of the invitation
type Invite struct {
	Token     string
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	DestOrgID ulid.ULID
	Role      enums.Role
	InviteToken
}

// InviteToken holds data specific to a future user of the system for invite logic
type InviteToken struct {
	Expires sql.NullString
	Token   sql.NullString
	Secret  []byte
}

// GetInviteToken returns the invitation token if its valid
func (i *Invite) GetInviteToken() string {
	if i.InviteToken.Token.Valid {
		return i.InviteToken.Token.String
	}

	return ""
}

// GetInviteExpires returns the expiration time of invite token
func (i *Invite) GetInviteExpires() (time.Time, error) {
	if i.InviteToken.Expires.Valid {
		return time.Parse(time.RFC3339Nano, i.InviteToken.Expires.String)
	}

	return time.Time{}, nil
}

// setOrgInviteTokens ets the fields of the `Invite` struct to verify the email
// invitation. It takes in an `Invite` object and an invitation token as parameters. If
// the invitation token matches the token stored in the `Invite` object, it sets the
// `Token`, `Secret`, and `Expires` fields of the `InviteToken` struct. This allows the
// token to be verified later when the user accepts the invitation
func (i *Invite) setOrgInviteTokens(inv *generated.Invite, invToken string) error {
	if inv.Token == invToken {
		i.InviteToken.Token = sql.NullString{String: inv.Token, Valid: true}
		i.InviteToken.Secret = *inv.Secret
		i.InviteToken.Expires = sql.NullString{String: inv.Expires.Format(time.RFC3339Nano), Valid: true}

		return nil
	}

	return ErrNotFound
}

// updateInviteStatusAccepted updates the status of an invite to "Accepted"
func updateInviteStatusAccepted(ctx context.Context, i *generated.Invite) error {
	_, err := transaction.FromContext(ctx).Invite.UpdateOneID(i.ID).SetStatus(enums.InvitationAccepted).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// updateInviteStatusAccepted updates the status of an invite to "Expired"
func updateInviteStatusExpired(ctx context.Context, i *generated.Invite) error {
	_, err := transaction.FromContext(ctx).Invite.UpdateOneID(i.ID).SetStatus(enums.InvitationExpired).Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
