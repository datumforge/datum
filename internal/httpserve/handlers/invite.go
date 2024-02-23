package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/ent/privacy/token"
	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/transaction"
	"github.com/datumforge/datum/internal/rout"
	"github.com/datumforge/datum/internal/tokens"
	"github.com/datumforge/datum/pkg/auth"
)

// InviteReply holds the fields that are sent on a response to an accepted invitation
// Note: there is no InviteRequest as this is handled via our graph interfaces
type InviteReply struct {
	rout.Reply
	ID          string `json:"user_id"`
	Email       string `json:"email"`
	Message     string `json:"message"`
	JoinedOrgID string `json:"joined_org_id"`
	Role        string `json:"role"`
}

// Invite holds the Token, InviteToken references, and the additional user input to //
// complete acceptance of the invitation
type Invite struct {
	Token     string
	UserID    ulid.ULID
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

// OrganizationInviteAccept is responsible for handling the invitation of a user to an organization.
// It receives a request with the user's invitation details, validates the request,
// and creates organization membership for the user
// On success, it returns a response with the organization information
func (h *Handler) OrganizationInviteAccept(ctx echo.Context) error {
	// setup view context
	context := ctx.Request().Context()
	userCtx := viewer.NewContext(context, viewer.NewUserViewerFromSubject(context))

	// get the authenticated user from the context
	userID, err := auth.GetUserIDFromContext(context)
	if err != nil {
		h.Logger.Errorw("unable to get user id from context", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// parse the token out of the context
	inv := &Invite{
		Token: ctx.QueryParam("token"),
	}

	// ensure the user that is logged in, matches the invited user
	if err := inv.validateInviteRequest(); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// set the initial context based on the token
	ctxWithToken := token.NewContextWithOrgInviteToken(userCtx, inv.Token)

	// fetch the recipient and org owner based on token
	invitedUser, err := h.getUserByInviteToken(ctxWithToken, inv.Token)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving invite token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	// add email to the invite
	inv.Email = invitedUser.Recipient

	// get user details for logged in user
	user, err := h.getUserBySub(userCtx, userID)
	if err != nil {
		h.Logger.Errorw("unable to get user for request", "error", err)

		return ctx.JSON(http.StatusUnauthorized, rout.ErrorResponse(err))
	}

	// ensure the user that is logged in, matches the invited user
	if err := inv.validateUser(user.Email); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// string to ulid so we can match the token input
	oid, err := ulid.Parse(invitedUser.OwnerID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// string to ulid so we can match the token input
	uid, err := ulid.Parse(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// construct the invite details but set email to the original recipient, and the joining organization ID as the current owner of the invitation
	invite := &Invite{
		Email:     invitedUser.Recipient,
		UserID:    uid,
		DestOrgID: oid,
		Role:      invitedUser.Role,
	}

	// set tokens for request
	if err := invite.setOrgInviteTokens(invitedUser, inv.Token); err != nil {
		h.Logger.Errorw("unable to set invite token for request", "error", err)

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
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

			out := &InviteReply{
				Message: "invite token is expired, you will need to re-request an invite",
			}

			return ctx.JSON(http.StatusBadRequest, out)
		}

		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	if err := updateInviteStatusAccepted(ctxWithToken, invitedUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	// reply with the relevant details
	out := &InviteReply{
		Reply:       rout.Reply{Success: true},
		ID:          userID,
		Email:       invitedUser.Recipient,
		JoinedOrgID: invitedUser.OwnerID,
		Role:        string(invitedUser.Role),
		Message:     "Welcome to your new organization!",
	}

	return ctx.JSON(http.StatusCreated, out)
}

// validateInviteRequest validates the required fields are set in the user request
func (i *Invite) validateInviteRequest() error {
	// ensure the token is set
	if i.Token == "" {
		return rout.NewMissingRequiredFieldError("token")
	}

	return nil
}

func (i *Invite) validateUser(email string) error {
	// ensure the logged in user is the same as the invite
	if i.Email != email {
		return ErrUnableToVerifyEmail
	}

	return nil
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
