package company

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/company/countries"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Company represents a generic company.
type Company struct {
	BrazilianCompany countries.BrazilianCompany
}

// NewCompany generates a randomized company entity utilizing a user-defined language and random source
func NewCompany(lang string, isFormatted bool, rnd translations.RandSource) (*Company, error) {
	comp, err := generateCompany(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}
	return comp, nil
}

func generateCompany(lang string, formatted bool, rnd translations.RandSource) (*Company, error) {
	createdCompanyBrazilian, err := countries.GenerateBrazilianCompany(lang, formatted, rnd)
	if err != nil {
		return nil, err
	}
	return &Company{BrazilianCompany: createdCompanyBrazilian}, nil
}
