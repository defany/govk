package updates

import (
	"fmt"
	"github.com/defany/govk/api"
	"reflect"
)

type callback[T any] func(data T)

type Updates struct {
	callbacks map[string]func(data any)

	api *api.API
}

func NewUpdates(api *api.API) *Updates {
	return &Updates{
		api:       api,
		callbacks: make(map[string]func(data any)),
	}
}

func (u *Updates) Check() {

}

// On - function to helps you handle updates from VK
//
// With generics it can be helpful to work without interface{}, just pass it into function and see a magic.
//
// Try to use types only from updates package, but if you trust yourself you can provide your own types. For instance - types for user bots.
// Remember, that you can pass only struct as generic. No pointers, no something else.
//
// All event objects named as in vk, for example: messages_new -> MessagesNew
//
// Keep in mind, that this function will panic if something went wrong
func On[T any](updates *Updates, event string, callback callback[T]) {
	var zero T

	typ := reflect.TypeOf(zero)
	if typ.Kind() == reflect.Pointer {
		panic("pointer has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	if typ.Kind() != reflect.Struct {
		panic("something not a struct has been passed into vk updates handler, recheck your generic types in your `On` functions")
	}

	updates.callbacks[event] = func(data any) {
		pd, ok := data.(T)
		if !ok {
			panic(fmt.Sprintf("failed to cast incoming type %s to %s", reflect.TypeOf(data), reflect.TypeOf(zero)))
		}

		callback(pd)
	}
}
