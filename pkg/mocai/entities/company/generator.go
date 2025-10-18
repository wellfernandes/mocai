package company

import "github.com/brazzcore/mocai/pkg/mocai/entities/company/countries"

// Company represents a generic company.
type Company struct {
	BrazilianCompany countries.BrazilianCompany
}

// NewCompany creates a new Company with generated data.
func NewCompany(isFormatted bool) Company {
	comp, err := (&Company{}).generateCompany(isFormatted)
	if err != nil {
		return Company{}
	}
	return comp
}

// GenerateCompany generates all companies available.
func (c *Company) generateCompany(formatted bool) (Company, error) {
	createdCompanyBrazilian, err := countries.GenerateBrazilianCompany(formatted)
	if err != nil {
		return Company{}, err
	}

	return Company{BrazilianCompany: createdCompanyBrazilian}, nil
}
