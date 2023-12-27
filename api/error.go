package api

import "errors"

var (
	ErrTokensNotProvided = errors.New("api tokens not provided")
)

var (
	ErrUnknownParamType = errors.New("unknown param type")
)
