package datumclient

import (
	"context"

	"github.com/datumforge/datum/pkg/models"
)

// A Reauthenticator generates new access and refresh pair given a valid refresh token
type Reauthenticator interface {
	Refresh(context.Context, *models.RefreshRequest) (*models.RefreshReply, error)
}
