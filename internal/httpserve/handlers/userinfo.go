package handlers

import (
	"net/http"

	"github.com/datumforge/datum/internal/ent/privacy/viewer"
	"github.com/datumforge/datum/internal/httpserve/middleware/auth"
	echo "github.com/datumforge/echox"
)

func (h *Handler) UserInfo(ctx echo.Context) error {
	// setup view context
	context := ctx.Request().Context()
	userCtx := viewer.NewContext(context, viewer.NewUserViewerFromSubject(context))

	userId, err := auth.GetUserIDFromContext(context)
	if err != nil {
		return err
	}

	// get user
	user, err := h.getUserBySub(userCtx, userId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}
