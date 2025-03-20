package countries

import "errors"

// Error constants for Company data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingBrazilianCompany = errors.New("error generating brazilian company")
)
