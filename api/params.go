package api

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

type MethodParams interface {
	Params() Params
}

type ParamMarshaler struct {
	Key       string
	Marshaler func(v any) (any, error)
}

type BaseRequestParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Params map[string]any

func (p Params) Build() (url.Values, error) {
	query := url.Values{}

	for k, v := range p {
		switch k {
		case tokenParam:
			continue
		default:
			pv, err := FmtValue(v, 0)
			if err != nil {
				return nil, err
			}

			query.Set(k, pv)
		}
	}

	return query, nil
}

// Params For compatability
func (p Params) Params() Params {
	return p
}

func (p Params) WithLanguage(v uint) Params {
	p[langParam] = v

	return p
}

func fmtReflectValue(value reflect.Value, depth int) (string, error) {
	switch f := value; value.Kind() {
	case reflect.Invalid:
		return "", nil
	case reflect.Bool:
		return fmtBool(f.Bool()), nil
	case reflect.Array, reflect.Slice:
		var builder strings.Builder

		// TODO: try to find a good size for strings
		builder.Grow(f.Len())

		for i := 0; i < f.Len(); i++ {
			if i > 0 {
				builder.WriteString(",")
			}

			res, err := FmtValue(f.Index(i).Interface(), depth)
			if err != nil {
				return "", err
			}

			builder.WriteString(res)
		}

		return builder.String(), nil
	case reflect.Ptr:
		// pointer to array or slice or struct? ok at top level
		// but not embedded (avoid loops)
		if depth == 0 && f.Pointer() != 0 {
			switch a := f.Elem(); a.Kind() {
			case reflect.Array, reflect.Slice, reflect.Struct, reflect.Map:
				return FmtValue(a.Interface(), depth+1)
			default:
				return "", ErrUnknownParamType
			}
		}
	default:
		return fmt.Sprint(value), nil
	}

	return fmt.Sprint(value), nil
}

func fmtBool(value bool) string {
	if value {
		return "1"
	}

	return "0"
}

// FmtValue return api format string.
func FmtValue(value interface{}, depth int) (string, error) {
	if value == nil {
		return "", nil
	}

	switch f := value.(type) {
	case bool:
		return fmtBool(f), nil
	case reflect.Value:
		return fmtReflectValue(f, depth)
	default:
		return fmtReflectValue(reflect.ValueOf(value), depth)
	}
}

func ParamsOrNil(params []MethodParams) MethodParams {
	if len(params) > 0 {
		return params[0]
	}

	return nil
}
