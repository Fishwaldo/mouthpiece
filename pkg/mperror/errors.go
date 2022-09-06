package mperror

import (
	"errors"
)

var (
	ErrAppExists                   = errors.New("App Already Exists")
	ErrAppNotFound                 = errors.New("App Not Found")
	ErrUserNotFound                = errors.New("User Not Found")
	ErrFilterNotFound              = errors.New("Filter Not Found")
	ErrTransportInstanceNotFound   = errors.New("Transport Instance Not Found")
	ErrTransportReciepiantNotFound = errors.New("Transport Recipient Not Found")
)
