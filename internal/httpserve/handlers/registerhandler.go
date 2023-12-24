package handlers

import (
	"encoding/json"

	echo "github.com/datumforge/echox"
	"github.com/oklog/ulid/v2"

	"github.com/datumforge/datum/internal/ent/generated"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
)

func (h *Handler) Register(ctx echo.Context, input generated.CreateUserInput) (*User, error) {
	var u User

	// parse request body
	if err := json.NewDecoder(ctx.Request().Body).Decode(&u); err != nil {
		auth.Unauthorized(ctx) //nolint:errcheck
		return nil, ErrBadRequest
	}

	input = generated.CreateUserInput{}
	err := h.DBClient.User.Create().SetInput(input).Exec(ctx.Request().Context())

	if err != nil {
		return &u, err
	}

	return &u, err
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
