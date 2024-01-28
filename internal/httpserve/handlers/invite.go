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

func (h *Handler) OrganizationInvite(ctx echo.Context) error {
	inv := &Invite{
		Token: ctx.QueryParam("token"),
	}

	var in *InviteInput

	if err := json.NewDecoder(ctx.Request().Body).Decode(&in); err != nil {
		h.Logger.Errorw("error parsing request", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrProcessingRequest))
	}

	inv.Password = in.Password
	inv.FirstName = in.FirstName
	inv.LastName = in.LastName

	if err := inv.validateInviteRequest(); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	ctxWithToken := token.NewContextWithOrgInviteToken(ctx.Request().Context(), inv.Token)

	invitedUser, err := h.getUserByInviteToken(ctxWithToken, inv.Token)
	if err != nil {
		if generated.IsNotFound(err) {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

		h.Logger.Errorf("error retrieving invite token", "error", err)

		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	oid, err := ulid.Parse(invitedUser.OwnerID)
	if err != nil {
		return err
	}

	invite := &Invite{
		Email:     invitedUser.Recipient,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		DestOrgID: oid,
		Role:      invitedUser.Role,
	}

	// set tokens for request
	if err := invite.setOrgInviteTokens(invitedUser.Edges.Owner, inv.Token); err != nil {
		h.Logger.Errorw("unable to set invite token for request", "error", err)

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	// reconstruct the token based on recipient & owning organization
	token := &tokens.OrgInviteToken{
		Email: invitedUser.Recipient,
		OrgID: oid,
	}

	if token.ExpiresAt, err = invite.GetInviteExpires(); err != nil {
		h.Logger.Errorw("unable to parse expiration", "error", err)

		return ctx.JSON(http.StatusInternalServerError, ErrUnableToVerifyEmail)
	}

	// Verify the token is valid with the stored secret
	if err = token.Verify(invite.GetInviteToken(), invite.Secret); err != nil {
		if errors.Is(err, tokens.ErrTokenExpired) {
			out := &Response{
				Message: "invite token is expired, you will need to re-request an invite",
			}

			return ctx.JSON(http.StatusBadRequest, out)
		}

		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	createdUser, err := h.createUser(ctxWithToken, generated.CreateUserInput{
		Email:     invitedUser.Recipient,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Password:  &in.Password,
	})
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

	viewerCtx := viewer.NewContext(ctxWithToken, viewer.NewUserViewerFromID(createdUser.ID, true))

	if err := h.setEmailConfirmed(viewerCtx, createdUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	mem, err := h.addUserToOrganization(viewerCtx, invitedUser, createdUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	out := &InviteReply{
		ID:          createdUser.ID,
		Email:       createdUser.Email,
		JoinedOrgID: mem.OrgID,
		Role:        string(mem.Role),
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

// InviteRequest holds the fields that should be included on a request to the `/invite` endpoint
type InviteRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	DestOrgID string `json:"destination_org_id"`
	Role      enums.Role
}

// InviteReply holds the fields that are sent on a response to the `/invite` endpoint
type InviteReply struct {
	ID            string `json:"user_id"`
	Email         string `json:"email"`
	Message       string `json:"message"`
	PersonalOrgID string `json:"personal_org_id"`
	JoinedOrgID   string `json:"joined_org_id"`
	Role          string `json:"role"`
}

type InviteInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

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

// GetInviteToken returns the verification token if its valid
func (i *Invite) GetInviteToken() string {
	if i.InviteToken.Token.Valid {
		return i.InviteToken.Token.String
	}

	return ""
}

// GetInviteExpires returns the expiration time of email verification token
func (i *Invite) GetInviteExpires() (time.Time, error) {
	if i.InviteToken.Expires.Valid {
		return time.Parse(time.RFC3339Nano, i.InviteToken.Expires.String)
	}

	return time.Time{}, nil
}

// CreateInviteToken creates a new email verification token for the user
func (i *Invite) CreateInviteToken() error {
	// Create a unique token from the user's email address
	verify, err := tokens.NewOrgInvitationToken(i.Email, i.DestOrgID)
	if err != nil {
		return err
	}

	// Sign the token to ensure that we can verify it later
	token, secret, err := verify.Sign()
	if err != nil {
		return err
	}

	i.InviteToken.Token = sql.NullString{Valid: true, String: token}
	i.InviteToken.Expires = sql.NullString{Valid: true, String: verify.ExpiresAt.Format(time.RFC3339Nano)}
	i.InviteToken.Secret = secret

	return nil
}

// setUserTokens sets the fields to verify the email
func (i *Invite) setOrgInviteTokens(inv *generated.Organization, invToken string) error {
	tokens := inv.Edges.Invites
	for _, t := range tokens {
		if t.Token == invToken {
			i.InviteToken.Token = sql.NullString{String: t.Token, Valid: true}
			i.InviteToken.Secret = *t.Secret
			i.InviteToken.Expires = sql.NullString{String: t.Expires.Format(time.RFC3339Nano), Valid: true}

			return nil
		}
	}

	return ErrNotFound
}

func (h *Handler) addUserToOrganization(ctx context.Context, i *generated.Invite, user *generated.User) (*generated.OrgMembership, error) {
	input := generated.CreateOrgMembershipInput{
		UserID: user.ID,
		OrgID:  i.OwnerID,
		Role:   &i.Role,
	}

	mem, err := transaction.FromContext(ctx).OrgMembership.Create().SetInput(input).Save(ctx)

	if err != nil {
		h.Logger.Errorw("error creating org membership for owner", "error", err)

		return nil, err
	}

	return mem, nil
}
