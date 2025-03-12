package models

// BrazilianCompany represents a mock company with a name and CNPJ.
type BrazilianCompany struct {
	CompanyName string
	CNPJ        string
}

func (b *BrazilianCompany) Name() string {
	return b.CompanyName
}

func (b *BrazilianCompany) RegistrationNumber() string {
	return b.CNPJ
}
