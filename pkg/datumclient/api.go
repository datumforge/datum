package datumclient

import (
	"context"

	api "github.com/datumforge/datum/pkg/models"
)

type DatumClient interface {
	Register(context.Context, *api.RegisterRequest) (*api.RegisterReply, error)
	Login(context.Context, *api.LoginRequest) (*api.LoginReply, error)
	Refresh(context.Context, *api.RefreshRequest) (*api.RefreshReply, error)
	Switch(context.Context, *api.SwitchOrganizationRequest) (*api.SwitchOrganizationReply, error)
	VerifyEmail(context.Context, *api.VerifyRequest) (*api.VerifyReply, error)
	ResendEmail(context.Context, *api.ResendRequest) (*api.ResendReply, error)
	ForgotPassword(context.Context, *api.ForgotPasswordRequest) (*api.ForgotPasswordReply, error)
	ResetPassword(context.Context, *api.ResetPasswordRequest) (*api.ResetPasswordReply, error)
	Status(context.Context) (*api.StatusReply, error)
	Invite(context.Context, *api.InviteRequest) (*api.InviteReply, error)
	WaitForReady(context.Context) error
}
