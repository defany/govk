package updates

import (
	"errors"
)

var (
	ErrFailedToCastRawToFloat64 = errors.New("failed to cast raw to float64")
	ErrLongpollFailed           = errors.New("longpoll unknown error")
)
