package company

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/company/interfaces"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func GenerateCompany() (interfaces.Company, error) {
	lang := translations.GetLanguage()

	factory := CompanyFactory{}
	generator, err := factory.GetCompanyGenerator(lang)
	if err != nil {
		return nil, err
	}

	return generator.GenerateCompany()
}
