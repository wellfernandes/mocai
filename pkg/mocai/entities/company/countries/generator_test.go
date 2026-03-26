package countries

import (
	"errors"
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

func TestCompanyFormattedCNPJ(t *testing.T) {
	company, err := GenerateBrazilianCompany("ptbr", true, nil)
	if err != nil {
		t.Errorf("Failed to generate company: %v", err)
	}
	// Formatted CNPJ: XX.XXX.XXX/XXXX-XX (18 chars)
	if len(company.CNPJ) != 18 {
		t.Errorf("expected 18 char formatted CNPJ, got %d: %s", len(company.CNPJ), company.CNPJ)
	}
	if !cnpj.ValidateCNPJ(company.CNPJ) {
		t.Errorf("Generated formatted CNPJ is invalid: %s", company.CNPJ)
	}
}

func TestCompanyUnsupportedLang(t *testing.T) {
	_, err := GenerateBrazilianCompany("xyz", false, nil)
	if err == nil {
		t.Error("expected error for unsupported language, got nil")
	}
	if !errors.Is(err, ErrNoCompanyNames) {
		t.Errorf("expected ErrNoCompanyNames, got: %v", err)
	}
}

func TestGenerateMultipleCompanies(t *testing.T) {
	for i := 0; i < 50; i++ {
		company, err := GenerateBrazilianCompany("ptbr", false, nil)
		if err != nil {
			t.Errorf("unexpected error on iteration %d: %v", i, err)
		}
		if company.Name == "" || company.CNPJ == "" {
			t.Errorf("company has empty fields on iteration %d: %+v", i, company)
		}
	}
}
