package handlers

import (
	"net/http"

	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/rout"
)

type PublishRequest struct {
	Tags    map[string]string `json:"tags" koanf:"tags"`
	Topic   string            `json:"topic" koanf:"topic"`
	Message string            `json:"message" koanf:"message"`
}

func (h *Handler) EventPublisher(ctx echo.Context) error {
	var in PublishRequest
	if err := ctx.Bind(&in); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	h.EventManager.Topic = "mitb"
	if err := h.EventManager.Publish([]byte(in.Message)); err != nil {
		return ctx.JSON(http.StatusConflict, rout.ErrorResponse(err))
	}

	return ctx.JSON(http.StatusOK, "success!")
}
