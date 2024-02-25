package updates

import (
	"context"
	"fmt"
	"github.com/goccy/go-json"
	"reflect"
)

// Use - function to helps you handle updates from VK
//
// With generics it can be helpful to work without interface{}, just pass it into function and see a magic.
//
// Try to use types only from <any>/model package, but if you trust yourself you can provide your own types. For instance - types for user bots.
// Remember, that you can pass only struct as generic. No pointers, no something else.
//
// All event objects named as in vk, for example: messages_new -> MessagesNew
//
// Keep in mind, that this function will panic if something went wrong
func Use[T any](updates *Updates, event string, middlewares ...middleware[T]) {
	var zero T

	typ := reflect.TypeOf(zero)
	if typ.Kind() == reflect.Pointer {
		panic("pointer has been passed into vk updates handler, recheck your generic types in your `Use` functions")
	}

	if typ.Kind() != reflect.Struct {
		panic("something not a struct has been passed into vk updates handler, recheck your generic types in your `Use` functions")
	}

	// FIXME: add growing event middlewares slice cap in future
	eventMiddlewares, ok := updates.middlewares[event]
	if !ok {
		eventMiddlewares = make([]middleware[json.RawMessage], 0, len(middlewares))
	}

	for _, middleware := range middlewares {
		eventMiddlewares = append(eventMiddlewares, func(ctx context.Context, data json.RawMessage) bool {
			var marshaled T

			if err := json.Unmarshal(data, &marshaled); err != nil {
				panic(fmt.Sprintf("failed to unmarshal incoming type %s to %s", reflect.TypeOf(data), typ))
			}

			return middleware(ctx, marshaled)
		})
	}

	updates.middlewares[event] = eventMiddlewares
}
