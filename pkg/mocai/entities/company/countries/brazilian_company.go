package countries

import (
	"errors"
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Sentinel errors for Brazilian company generation.
var (
	ErrNoCompanyNames     = errors.New("company: no company names available")
	ErrGeneratingCompany  = errors.New("company: error generating brazilian company")
	ErrInvalidCompanyCNPJ = errors.New("company: invalid CNPJ")
)

// BrazilianCompany represents a Brazilian company.
type BrazilianCompany struct {
	Name string
	CNPJ string
}

// GenerateBrazilianCompany generates a mock Brazilian company using customizable lists, language, and random source.
func GenerateBrazilianCompany(lang string, formatted bool, rnd translations.RandSource) (BrazilianCompany, error) {
	companyNames := translations.GetList(lang, "company_name")
	if len(companyNames) == 0 {
		return BrazilianCompany{}, fmt.Errorf("%w: %s", ErrNoCompanyNames, translations.Get(lang, "no_company_names_available"))
	}

	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}

	companyName := companyNames[rnd.Intn(len(companyNames))]

	cnpjVal, err := cnpj.GenerateCNPJ(formatted, rnd)
	if err != nil {
		return BrazilianCompany{}, err
	}

	if companyName == "" {
		return BrazilianCompany{}, fmt.Errorf("%w: %s", ErrGeneratingCompany, translations.Get(lang, "error_generating_brazilian_company"))
	}

	if cnpjVal == "" {
		return BrazilianCompany{}, fmt.Errorf("%w: %s", ErrInvalidCompanyCNPJ, translations.Get(lang, "invalid_cnpj"))
	}

	return BrazilianCompany{
		Name: companyName,
		CNPJ: cnpjVal,
	}, nil
}
