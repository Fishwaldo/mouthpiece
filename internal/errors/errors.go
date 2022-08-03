package mperror

import (
	"errors"
)

var (
	ErrAppExists = errors.New("App Already Exists")
	ErrAppNotFound = errors.New("App Not Found")
)