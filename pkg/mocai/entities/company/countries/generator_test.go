package countries

import (
	"testing"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cnpj"
)

func TestGenerateCompany(t *testing.T) {
	company, err := GenerateBrazilianCompany("ptbr", false, nil)
	if err != nil {
		t.Errorf("Failed to generate company: %v", err)
	}

	if company.Name == "" || company.CNPJ == "" {
		t.Errorf("Generated company has empty fields: %+v", company)
	}
}

func TestIfCompanyHasAValidCNPJ(t *testing.T) {
	company, err := GenerateBrazilianCompany("ptbr", false, nil)
	if err != nil {
		t.Errorf("Failed to generate company: %v", err)
	}

	if !cnpj.ValidateCNPJ(company.CNPJ) {
		t.Errorf("Generated company has an invalid CNPJ: %s", company.CNPJ)
	}
}
