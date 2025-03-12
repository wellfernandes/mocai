package company

import (
	"errors"

	"github.com/brazzcore/mocai/pkg/mocai/entities/company/generators"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company/interfaces"
)

type CompanyFactory struct{}

func (f *CompanyFactory) GetCompanyGenerator(lang string) (interfaces.CompanyStrategy, error) {
	switch lang {
	case "ptbr":
		return &generators.BrazilianCompanyGenerator{}, nil
	default:
		return nil, errors.New(ErrUnsupportedLanguageToGenerateCompany.Error())
	}
}
