package heargo

import (
	"context"
	"github.com/defany/govk/api/messages/model"
)

const (
	defaultCommandPayload = ":command"
)

type HandlerEvent interface {
	msgmodel.MessagesNew | msgmodel.MessagesEvent
}

type callback[Event HandlerEvent] func(ctx context.Context, event Event)

type Handler[Event HandlerEvent] struct {
	handler callback[Event]

	matchRules []Matcher[Event]
}

// NewHandler for handling commands from user by payload or text word
func NewHandler[Event HandlerEvent](handler callback[Event]) *Handler[Event] {
	return &Handler[Event]{
		handler: handler,
	}
}

// WithMatchRules provides a way to add rules to ignore some event with pre-validation
//
// Remember, that all rules will be called in order that u added them
func (h *Handler[Event]) WithMatchRules(matchers ...Matcher[Event]) *Handler[Event] {
	if len(h.matchRules) == 0 {
		h.matchRules = make([]Matcher[Event], 0, len(matchers))
	}

	h.matchRules = append(h.matchRules, matchers...)

	return h
}
