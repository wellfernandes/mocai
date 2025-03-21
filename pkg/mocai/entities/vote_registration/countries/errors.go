package countries

import "errors"

// Errors for Vote Registration data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrInvalidVoteRegistration = errors.New("invalid vote registration")
)
