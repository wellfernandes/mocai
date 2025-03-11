package company

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// BrazilianCompany represents a mock company with a name and CNPJ.
type BrazilianCompany struct {
	CompanyName string
	CNPJ        string
}

// GenerateCompany generates a mock Brazilian company with a valid CNPJ.
func GenerateCompany() (*BrazilianCompany, error) {
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
