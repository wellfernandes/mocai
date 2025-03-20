package countries

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type BrazilianCompany struct {
	CompanyName string
	CNPJ        string
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateBrazilianCompany() (*BrazilianCompany, error) {
	lang := translations.GetLanguage()

	// Get the list of company names
	companyNames := strings.Split(translations.Get(lang, "company_name"), ",")

	// Validate data
	if len(companyNames) == 0 {
		return nil, errors.New("no company names available")
	}

	// Choose a random company name
	companyName := companyNames[rng.Intn(len(companyNames))]

	// Generate a random CNPJ without a mask
	cnpj, err := cnpj.GenerateCNPJ(false)
	if err != nil {
		return nil, err
	}

	// Validate required fields
	if companyName == "" {
		return nil, fmt.Errorf("%s: error generating company: missing required data (companyName: %s)", ERROR_GENERATING_BRAZILIAN_COMPANY, companyName)
	}

	createdCompany := &BrazilianCompany{
		CompanyName: companyName,
		CNPJ:        cnpj,
	}

	return createdCompany, nil
}
