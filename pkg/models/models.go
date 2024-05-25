package models

import (
	"github.com/go-webauthn/webauthn/protocol"

	"github.com/datumforge/datum/pkg/rout"
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

// LoginRequest to authenticate with the Datum Sever
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	OTPCode  string `json:"otp_code,omitempty"`
}

// LoginReply holds response to successful authentication
type LoginReply struct {
	rout.Reply
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Session      string `json:"session,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Message      string `json:"message"`
}

// RefreshRequest holds the fields that should be included on a request to the `/refresh` endpoint
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshReply holds the fields that are sent on a response to the `/refresh` endpoint
type RefreshReply struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// RegisterRequest holds the fields that should be included on a request to the `/register` endpoint
type RegisterRequest struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// RegisterReply holds the fields that are sent on a response to the `/register` endpoint
type RegisterReply struct {
	rout.Reply
	ID      string `json:"user_id"`
	Email   string `json:"email"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

// SwitchOrganizationRequest contains the target organization ID being switched to
type SwitchOrganizationRequest struct {
	TargetOrganizationID string `json:"target_organization_id"`
}

// SwitchOrganizationReply holds the new authentication and session information for the user for the new organization
type SwitchOrganizationReply struct {
	rout.Reply
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Session      string `json:"session"`
}

// VerifyRequest holds the fields that should be included on a request to the `/verify` endpoint
type VerifyRequest struct {
	Token string `query:"token"`
}

// VerifyReply holds the fields that are sent on a response to the `/verify` endpoint
type VerifyReply struct {
	rout.Reply
	ID           string `json:"user_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	Message      string `json:"message,omitempty"`
}

// ResendRequest contains fields for a resend email verification request
type ResendRequest struct {
	Email string `json:"email"`
}

// ResendReply holds the fields that are sent on a response to the `/resend` endpoint
type ResendReply struct {
	rout.Reply
	Message string `json:"message"`
}

// ForgotPasswordRequest contains fields for a forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

// ForgotPasswordReply contains fields for a forgot password response
type ForgotPasswordReply struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// ResetPasswordRequest contains user input required to reset a user's password
type ResetPasswordRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

// ResetPasswordReply is the response returned from a non-successful password reset request
// on success, no content is returned (204)
type ResetPasswordReply struct {
	rout.Reply
	Message string `json:"message"`
}

// WebauthnRegistrationRequest is the request to begin a webauthn login
type WebauthnRegistrationRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// WebauthnRegistrationResponse is the response to begin a webauthn login
// this includes the credential creation options and the session token
type WebauthnBeginRegistrationResponse struct {
	Reply rout.Reply
	*protocol.CredentialCreation
	Session string `json:"session,omitempty"`
}

// WebauthnRegistrationResponse is the response after a successful webauthn registration
type WebauthnRegistrationResponse struct {
	rout.Reply
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Session      string `json:"session,omitempty"`
	TokenType    string `json:"token_type"`
}

// WebauthnBeginLoginResponse is the response to begin a webauthn login
// this includes the credential assertion options and the session token
type WebauthnBeginLoginResponse struct {
	Reply rout.Reply
	*protocol.CredentialAssertion
	Session string `json:"session,omitempty"`
}

// WebauthnRegistrationResponse is the response after a successful webauthn login
type WebauthnLoginResponse struct {
	rout.Reply
	Message      string `json:"message,omitempty"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Session      string `json:"session,omitempty"`
	TokenType    string `json:"token_type"`
}

// VerifySubscribeRequest holds the fields that should be included on a request to the `/subscribe/verify` endpoint
type VerifySubscribeRequest struct {
	Token string `query:"token"`
}

// VerifySubscribeReply holds the fields that are sent on a response to the `/subscribe/verify` endpoint
type VerifySubscribeReply struct {
	rout.Reply
	Message string `json:"message,omitempty"`
}

// PublishRequest is the request payload for the event publisher
type PublishRequest struct {
	Tags    map[string]string `json:"tags"`
	Topic   string            `json:"topic"`
	Message string            `json:"message"`
}

// PublishReply holds the fields that are sent on a response to the `/event/publish` endpoint
type PublishReply struct {
	rout.Reply
	Message string `json:"message"`
}

// InviteRequest holds the fields that should be included on a request to the `/invite` endpoint
type InviteRequest struct {
	Token string `query:"token"`
}

// Reply contains standard fields that are used for generic API responses and errors
type Reply struct {
	Success    bool   `json:"success"`
	Error      string `json:"error,omitempty"`
	Unverified bool   `json:"unverified,omitempty"`
}

// Returned on status requests
type StatusReply struct {
	Status  string `json:"status"`
	Uptime  string `json:"uptime,omitempty"`
	Version string `json:"version,omitempty"`
}

// SwitchRequest is the request payload for the switch organization endpoint
type SwitchRequest struct {
	OrgID string `json:"org_id"`
}
