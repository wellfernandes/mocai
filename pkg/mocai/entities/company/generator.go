package company

import "github.com/brazzcore/mocai/pkg/mocai/entities/company/countries"

// Company represents a generic company.
type Company struct {
	BrazilianCompany *countries.BrazilianCompany
}

// GenerateCompany generates all companies available.
func GenerateCompany() (*Company, error) {
	createdCompanyBrazilian, err := countries.GenerateBrazilianCompany()

	if err != nil {
		return nil, err
	}

	return &Company{BrazilianCompany: createdCompanyBrazilian}, nil
}
