package person

import "errors"

// Error constants for person-related data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingPerson = errors.New("person: error generating person")
	ErrNoFirstNames     = errors.New("person: no data available for first names")
	ErrNoLastNames      = errors.New("person: no data available for last names")
)
