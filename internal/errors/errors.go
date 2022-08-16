package mperror

import (
	"errors"
)

var (
	ErrAppExists = errors.New("App Already Exists")
	ErrAppNotFound = errors.New("App Not Found")
	ErrUserNotFound = errors.New("User Not Found")
)