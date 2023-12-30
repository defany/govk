package updates

import (
	"errors"
)

var (
	ErrFailedRawCast  = errors.New("failed to cast raw to")
	ErrLongpollFailed = errors.New("longpoll unknown error")
)
