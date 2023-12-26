package parser

import (
	"errors"
	"reflect"
)

var (
	ErrWrongType = errors.New("wrong type was passed")
	ErrNotValid  = errors.New("not valid")
)

func UnpackArray[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))

	for i, e := range s {
		r[i] = e
	}

	return r
}

func StructAttr(obj interface{}, fieldName string) (reflect.Value, error) {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()

	if curStruct.Kind() != reflect.Struct {
		return curStruct, ErrWrongType
	}

	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		return curField, ErrNotValid
	}

	return curField, nil
}
