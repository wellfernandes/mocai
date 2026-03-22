package countries

import (
	"fmt"
	"math/rand"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// BrazilianCompany represents a Brazilian company
type BrazilianCompany struct {
	Name string
	CNPJ string
}

// GenerateBrazilianCompany generates a mock Brazilian company using customizable lists, language, and random source
func GenerateBrazilianCompany(lang string, formatted bool, rnd translations.RandSource) (BrazilianCompany, error) {
	companyNames := translations.GetList(lang, "company_name")
	if len(companyNames) == 0 {
		return BrazilianCompany{}, fmt.Errorf("%s", translations.Get(lang, "no_company_names_available"))
	}
	if rnd == nil {
		rnd = defaultRandSource()
	}
	companyName := companyNames[rnd.Intn(len(companyNames))]

	cnpjVal, err := cnpj.GenerateCNPJ(formatted, rnd)
	if err != nil {
		return BrazilianCompany{}, err
	}
	if companyName == "" {
		return BrazilianCompany{}, fmt.Errorf("%s", translations.Get(lang, "error_generating_brazilian_company"))
	}
	if cnpjVal == "" {
		return BrazilianCompany{}, fmt.Errorf("%s", translations.Get(lang, "invalid_cnpj"))
	}
	createdCompany := BrazilianCompany{
		Name: companyName,
		CNPJ: cnpjVal,
	}
	return createdCompany, nil
}

func defaultRandSource() translations.RandSource {
	return rand.New(rand.NewSource(int64(rand.Int())))
}
