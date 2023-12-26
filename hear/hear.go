package heargo

import (
	"context"
	"github.com/defany/govk/api"
	msgmodel "github.com/defany/govk/api/messages/model"
	"github.com/defany/govk/pkg/cache"
	"time"
)

type TempCache[V any] interface {
	Save(key string, value V)
	Get(key string) (V, bool)
	Stop()
}

type Manager struct {
	api   *api.API
	cache TempCache[Handler[msgmodel.MessagesNew]]
}

func NewManager(api *api.API) *Manager {
	return &Manager{
		api:   api,
		cache: cache.NewTemp[Handler[msgmodel.MessagesNew]](time.Second * 60),
	}
}

// WithTempCache is way to provide your own storage implementation
func (h *Manager) WithTempCache(cache TempCache[Handler[msgmodel.MessagesNew]]) {
	h.cache = cache
}

func (h *Manager) Use(ctx context.Context, event msgmodel.MessagesNew) bool {
	var matchWord string

	payload, err := msgmodel.PayloadValue[string](event.Message, "command")
	if err == nil {
		matchWord = payload
	} else {
		matchWord = event.Message.Text
	}

	handler, ok := h.fromCache(matchWord)
	if ok {
		handler.handler(ctx, event)

		return true
	}

	return true
}

func (h *Manager) processMessagesNewPayload(ctx context.Context, event msgmodel.MessagesNew) bool {
	return false
}

func (h *Manager) processMessagesNewText(ctx context.Context, event msgmodel.MessagesNew) bool {
	return false
}

func (h *Manager) fromCache(matchWord string) (Handler[msgmodel.MessagesNew], bool) {
	return h.cache.Get(matchWord)
}
