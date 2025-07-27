package countries

import "errors"

// Errors for Company data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrGeneratingBrazilianCompany = errors.New("error generating brazilian company")
	ErrGeneratingCNPJ             = errors.New("error generating CNPJ")
	ErrNoCompanyNamesAvailable    = errors.New("no company names available")
)
