package hear

import (
	"context"
	"github.com/buger/jsonparser"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/pkg/cache"
	govk "github.com/defany/govk/vk"
	"time"
)

type TempCache[V any] interface {
	Save(key string, value V)
	Get(key string) (V, bool)
	Stop()
}

type Manager struct {
	vk       *govk.VK
	commands []*Handler[msgmodel.MessagesNew]
	cache    TempCache[*Handler[msgmodel.MessagesNew]]
}

type EventManager struct {
	vk       *govk.VK
	commands []*Handler[msgmodel.MessagesEvent]
	cache    TempCache[*Handler[msgmodel.MessagesEvent]]
}

func NewManager(vk *govk.VK) *Manager {
	return &Manager{
		vk:    vk,
		cache: cache.NewTemp[*Handler[msgmodel.MessagesNew]](time.Second*60, true),
	}
}

func NewEventManager(vk *govk.VK) *EventManager {
	return &EventManager{
		vk:    vk,
		cache: cache.NewTemp[*Handler[msgmodel.MessagesEvent]](time.Second*60, true),
	}
}

// WithTempCache is way to provide your own storage implementation
func (h *Manager) WithTempCache(cache TempCache[*Handler[msgmodel.MessagesNew]]) {
	h.cache = cache
}

// WithTempCache is way to provide your own storage implementation
func (h *EventManager) WithTempCache(cache TempCache[*Handler[msgmodel.MessagesEvent]]) {
	h.cache = cache
}

func (h *Manager) Middleware(ctx context.Context, event msgmodel.MessagesNew) bool {
	var matchWord string

	payload, err := jsonparser.GetString([]byte(event.Message.Payload), "command")
	if err == nil {
		matchWord = payload
	} else {
		matchWord = event.Message.Text
	}

	if matchWord == "" {
		return true
	}

	handler, ok := h.fromCache(matchWord)
	if ok {
		if handler.callback == nil {
			return true
		}
		handler.callback(ctx, event)

		return false
	}

	for _, command := range h.commands {
		if command.IsMatch(event) {
			h.inCache(matchWord, command)

			if command.callback == nil {
				return true
			}

			command.callback(ctx, event)

			return false
		}
	}

	return true
}

func (h *EventManager) Middleware(ctx context.Context, event msgmodel.MessagesEvent) bool {
	var matchWord string

	payload, err := jsonparser.GetString(event.Payload, "command")
	if err != nil {
		return true
	}

	matchWord = payload

	handler, ok := h.fromCache(matchWord)
	if ok {
		if handler.callback == nil {
			return true
		}

		handler.callback(ctx, event)

		return false
	}

	for _, command := range h.commands {
		if command.IsMatch(event) {
			h.inCache(matchWord, command)

			if command.callback == nil {
				return true
			}

			command.callback(ctx, event)

			return false
		}
	}

	return true
}

func (h *Manager) AddHandlers(handlers ...*Handler[msgmodel.MessagesNew]) {
	for _, handler := range handlers {
		h.commands = append(h.commands, handler)
		for _, child := range handler.children {
			matchRules := make([]Matcher[msgmodel.MessagesNew], 0, len(handler.matchRules)+len(child.matchRules))

			matchRules = append(matchRules, append(handler.matchRules, child.matchRules...)...)

			child.matchRules = []Matcher[msgmodel.MessagesNew]{}
			child.WithMatchRules(And(matchRules...))

			h.commands = append(h.commands, child)
		}
	}
	return
}

func (h *EventManager) AddHandlers(handlers ...*Handler[msgmodel.MessagesEvent]) {
	for _, handler := range handlers {
		h.commands = append(h.commands, handler)
		for _, child := range handler.children {
			matchRules := make([]Matcher[msgmodel.MessagesEvent], 0, len(handler.matchRules)+len(child.matchRules))

			matchRules = append(matchRules, append(handler.matchRules, child.matchRules...)...)

			child.matchRules = []Matcher[msgmodel.MessagesEvent]{}
			child.WithMatchRules(And(matchRules...))

			h.commands = append(h.commands, child)
		}
	}
	return
}

func (h *Manager) fromCache(matchWord string) (*Handler[msgmodel.MessagesNew], bool) {
	return h.cache.Get(matchWord)
}

func (h *Manager) inCache(matchWord string, handler *Handler[msgmodel.MessagesNew]) {
	h.cache.Save(matchWord, handler)
}

func (h *EventManager) fromCache(matchWord string) (*Handler[msgmodel.MessagesEvent], bool) {
	return h.cache.Get(matchWord)
}

func (h *EventManager) inCache(matchWord string, handler *Handler[msgmodel.MessagesEvent]) {
	h.cache.Save(matchWord, handler)
}
