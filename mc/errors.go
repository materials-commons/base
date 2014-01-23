package mc

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("Not found")
	ErrInvalidRequest = errors.New("Invalid Request")
)
