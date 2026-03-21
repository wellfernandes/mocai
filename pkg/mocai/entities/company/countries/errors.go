package countries

import "errors"

// Errors for Company data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingBrazilianCompany = errors.New("company: error generating brazilian company")
	ErrGeneratingCNPJ             = errors.New("company: error generating CNPJ")
	ErrNoCompanyNamesAvailable    = errors.New("company: no company names available")
)
