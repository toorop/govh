package govh

import (
	"errors"
)

var (
	ErrInvalidkey        = errors.New("INVALID_KEY")
	ErrInvalidCredential = errors.New("INVALID_CREDENTIAL")
)
