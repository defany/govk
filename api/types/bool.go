package types

import (
	"bytes"
	"github.com/goccy/go-json"
	"reflect"
)

type ApiBool bool

func (a *ApiBool) UnmarshalJSON(data []byte) (err error) {
	switch {
	case bytes.Equal(data, []byte("1")), bytes.Equal(data, []byte("true")):
		*a = true
	case bytes.Equal(data, []byte("0")), bytes.Equal(data, []byte("false")):
		*a = false
	default:
		// return json error
		err = &json.UnmarshalTypeError{
			Value: string(data),
			Type:  reflect.TypeOf((*ApiBool)(nil)),
		}
	}

	return
}
