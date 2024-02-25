package hear

import (
	"github.com/buger/jsonparser"
	"github.com/defany/govk/api/messages/model"
	"github.com/goccy/go-json"
	"slices"
	"strings"
)

type Matcher[Event msgmodel.MessagesNew | msgmodel.MessagesEvent] func(event Event) bool

func WithMatchWord(word string) Matcher[msgmodel.MessagesNew] {
	return func(event msgmodel.MessagesNew) bool {
		return strings.Contains(event.Message.Text, word) && event.Message.Text != "" && word != ""
	}
}

func WithWordEqualTo(word string) Matcher[msgmodel.MessagesNew] {
	return func(event msgmodel.MessagesNew) bool {
		return event.Message.Text == word
	}
}

func WithWordsIn(words ...string) Matcher[msgmodel.MessagesNew] {
	return func(event msgmodel.MessagesNew) bool {
		return slices.ContainsFunc(words, func(word string) bool {
			return event.Message.Text == word
		})
	}
}

func WithMatchPayload[Event HandlerEvent](payload string, commandName ...string) Matcher[Event] {
	return func(event Event) bool {
		command := defaultCommandPayload
		if len(commandName) > 0 {
			command = commandName[0]
		}

		gotPayload, err := jsonparser.GetString(marshal(event), command)
		if err != nil {
			return false
		}

		return strings.Contains(gotPayload, payload) && gotPayload != "" && command != ""
	}
}

func WithEqualPayloadTo[Event HandlerEvent](payload string, commandName ...string) Matcher[Event] {
	return func(event Event) bool {
		command := defaultCommandPayload
		if len(commandName) > 0 {
			command = commandName[0]
		}

		gotPayload, err := jsonparser.GetString(marshal(event), command)
		if err != nil {
			return false
		}

		return gotPayload == payload
	}
}

// And is true if all the matchers are true
func And[Event HandlerEvent](matchers ...Matcher[Event]) Matcher[Event] {
	return func(event Event) bool {
		for _, p := range matchers {
			if !p(event) {
				return false
			}
		}
		return true
	}
}

func marshal[Event HandlerEvent](event Event) []byte {
	m, _ := json.Marshal(event)

	return m
}
