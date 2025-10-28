package countries

import "errors"

// Errors for Vote Registration data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrInvalidVoteRegistration = errors.New("invalid vote registration")
	ErrInvalidCheckDigit1      = errors.New("invalid check digit 1")
	ErrInvalidCheckDigit2      = errors.New("invalid check digit 2")
)
