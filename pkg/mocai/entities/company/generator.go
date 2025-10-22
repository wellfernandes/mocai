package company

import "github.com/brazzcore/mocai/pkg/mocai/entities/company/countries"

// Company represents a generic company.
type Company struct {
	BrazilianCompany countries.BrazilianCompany
}

// NewCompany creates a new Company with generated data.
func NewCompany(isFormatted bool) (*Company, error) {
	comp, err := (&Company{}).generateCompany(isFormatted)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

// GenerateCompany generates all companies available.
func (c *Company) generateCompany(formatted bool) (*Company, error) {
	createdCompanyBrazilian, err := countries.GenerateBrazilianCompany(formatted)
	if err != nil {
		return nil, err
	}

	return &Company{BrazilianCompany: createdCompanyBrazilian}, nil
}
