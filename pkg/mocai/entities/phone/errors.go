package phone

import "errors"

// Error constants used for phone number generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingPhone = errors.New("error generating phone")
	ErrNoAreaCodes     = errors.New("no data available for area codes")
)
