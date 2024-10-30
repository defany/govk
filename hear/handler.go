package hear

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
	callback callback[Event]

	children []*Handler[Event]

	matchRules []Matcher[Event]
}

func NewHandler[Event HandlerEvent]() *Handler[Event] {
	return &Handler[Event]{}
}

func (h *Handler[Event]) WithCallback(callback callback[Event]) *Handler[Event] {
	h.callback = callback

	return h
}

func (h *Handler[Event]) Group(children ...*Handler[Event]) *Handler[Event] {
	if len(h.children) == 0 {
		h.children = make([]*Handler[Event], 0, len(children))
	}

	h.children = append(h.children, children...)

	return h
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

func (h *Handler[Event]) IsMatch(event Event) bool {
	for _, matcher := range h.matchRules {
		if ok := matcher(event); ok {
			return true
		}
	}

	return false
}
