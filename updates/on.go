package updates

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"reflect"
)

// On - function to helps you handle updates from VK
//
// With generics it can be helpful to work without interface{}, just pass it into function and see a magic.
//
// Try to use types only from <any>/model package, but if you trust yourself you can provide your own types. For instance - types for user bots.
// Remember, that you can pass only struct as generic. No pointers, no something else.
//
// All event objects named as in vk, for example: messages_new -> MessagesNew
//
// Keep in mind, that this function will panic if something went wrong
func On[T any](updates *Updates, event string, callbacks ...callback[T]) {
	var zero T

	typ := reflect.TypeOf(zero)
	if typ.Kind() == reflect.Pointer {
		panic("pointer has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	if typ.Kind() != reflect.Struct {
		panic("something not a struct has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	// FIXME: add growing event callbacks slice cap in future
	eventCallbacks, ok := updates.callbacks[event]
	if !ok {
		eventCallbacks = make([]callback[json.RawMessage], 0, len(callbacks))
	}

	for _, callback := range callbacks {
		eventCallbacks = append(eventCallbacks, func(ctx context.Context, data json.RawMessage) {
			var marshaled T

			if err := json.Unmarshal(data, &marshaled); err != nil {
				panic(fmt.Sprintf("failed to unmarshal incoming type %s to %s", reflect.TypeOf(data), typ))
			}

			callback(ctx, marshaled)
		})
	}

	updates.callbacks[event] = eventCallbacks
}
