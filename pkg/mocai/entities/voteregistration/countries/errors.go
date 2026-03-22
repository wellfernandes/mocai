package countries

import "errors"

var (
	ErrInvalidVoteRegistration = errors.New("vote registration: invalid vote registration")
	ErrInvalidCheckDigit1      = errors.New("vote registration: invalid check digit 1")
	ErrInvalidCheckDigit2      = errors.New("vote registration: invalid check digit 2")
)
