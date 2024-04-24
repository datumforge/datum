package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	echo "github.com/datumforge/echox"

	"github.com/datumforge/datum/pkg/rout"
)

type PublishRequest struct {
	Topic   string      `json:"topic" koanf:"topic"`
	Message interface{} `json:"message" koanf:"message"`
}

func (h *Handler) PublishEvent(ctx echo.Context) error {
	req := new(PublishRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, rout.ErrorResponse(err))
	}

	var payload []byte

	switch val := req.Message.(type) {
	case string:
		payload = []byte(val)
	default:
		b, _err := json.Marshal(val)
		if _err != nil {
			return fmt.Errorf("generate payload error: %q", _err)
		}

		payload = b
	}

	id := watermill.NewUUID()
	msg := message.NewMessage(id, payload)

	if err := h.EventPublisher.Publish(req.Topic, msg); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "meowmeow")
	}

	if err := h.Publisher.Publish(req.Topic, msg); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "oops")
	}

	return ctx.JSON(http.StatusOK, map[string]string{"id": id})
}
