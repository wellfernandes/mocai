package company

import "errors"

// Error constants for person-related data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingBrazilianCompany           = errors.New("error generating brazilian company")
	ErrNoCompanyNamesAvailable              = errors.New("no company names available")
	ErrUnsupportedLanguageToGenerateCompany = errors.New("unsupported language to generate company")
)
