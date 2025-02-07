package person

// Error constants for person-related data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
const (
	ERROR_GENERATING_PERSON = "error generating person"
	ERROR_NO_FIRST_NAMES    = "no data available for first names"
	ERROR_NO_LAST_NAMES     = "no data available for last names"
)
