package countries

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// BrazilianCompany represents a Brazilian company.
type BrazilianCompany struct {
	CompanyName string
	CNPJ        string
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateBrazilianCompany(formatted bool) (BrazilianCompany, error) {
	lang := translations.GetLanguage()

	// Get the list of company names
	companyNames := strings.Split(translations.Get(lang, "company_name"), ",")

	// Validate data
	if len(companyNames) == 0 {
		return BrazilianCompany{}, fmt.Errorf("%w, %s", ErrNoCompanyNamesAvailable, translations.Translate("no_company_names_available"))
	}

	// Choose a random company name
	companyName := companyNames[rng.Intn(len(companyNames))]

	// Generate a random CNPJ without a mask
	cnpj, err := cnpj.GenerateCNPJ(formatted)
	if err != nil {
		return BrazilianCompany{}, err
	}

	// Validate required fields
	if companyName == "" {
		return BrazilianCompany{}, fmt.Errorf("%w, %s", ErrGeneratingBrazilianCompany, translations.Translate("error_generating_brazilian_company"))
	}

	if cnpj == "" {
		return BrazilianCompany{}, fmt.Errorf("%w, %s", ErrGeneratingCNPJ, translations.Translate("invalid_cnpj"))
	}

	createdCompany := BrazilianCompany{
		CompanyName: companyName,
		CNPJ:        cnpj,
	}

	return createdCompany, nil
}
