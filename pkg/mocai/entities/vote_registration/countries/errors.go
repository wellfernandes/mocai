package countries

import "errors"

var (
	ErrInvalidCheckDigit1 = errors.New("invalid check digit 1")
	ErrInvalidCheckDigit2 = errors.New("invalid check digit 2")
	ErrInvalidSection     = errors.New("invalid section")
	ErrInvalidZone        = errors.New("invalid zone")
	ErrInvalidNumber      = errors.New("invalid number")
)
