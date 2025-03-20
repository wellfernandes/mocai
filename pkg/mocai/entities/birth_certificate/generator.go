package birth_certificate

import "github.com/brazzcore/mocai/pkg/mocai/entities/birth_certificate/countries"

type BirthCertificate struct {
	BrazilianBirthCertificate *countries.BrazilianBirthCertificate
}

func GenerateBirthCertificate(formatted bool) (*BirthCertificate, error) {
	createdBrazilianBirthCertificate, err := countries.GenerateBirthCertificate(formatted)
	if err != nil {
		return nil, err
	}

	return &BirthCertificate{BrazilianBirthCertificate: createdBrazilianBirthCertificate}, nil
}
