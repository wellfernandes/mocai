package person

import "errors"

// Error constants for person-related data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingPerson = errors.New("error generating person")
	ErrNoFirstNames     = errors.New("no data available for first names")
	ErrNoLastNames      = errors.New("no data available for last names")
)
