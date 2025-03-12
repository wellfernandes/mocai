package generators

import (
	"errors"
	"math/rand"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"

	"github.com/brazzcore/mocai/pkg/mocai/entities/company/interfaces"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company/mocks/ptbr"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company/models"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

type BrazilianCompanyGenerator struct{}

func (b *BrazilianCompanyGenerator) GenerateCompany() (interfaces.Company, error) {
	if len(ptbr.CompanyNames) == 0 {
		return nil, errors.New("no company names available")
	}

	companyName := ptbr.CompanyNames[rng.Intn(len(ptbr.CompanyNames))]
	cnpj, err := cnpj.GenerateCNPJ(false)
	if err != nil {
		return nil, err
	}

	brazilianCompany := &models.BrazilianCompany{
		CompanyName: companyName,
		CNPJ:        cnpj,
	}

	return brazilianCompany, nil
}
