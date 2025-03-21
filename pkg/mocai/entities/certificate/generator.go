package certificate

import "github.com/brazzcore/mocai/pkg/mocai/entities/certificate/countries"

type Certificate struct {
	BrazilianCertificates *countries.BrazilianCertificates
}

func GenerateCertificate(formatted bool) (*Certificate, error) {
	createdBrazilianCertificates, err := countries.GenerateBrazilianCertificates(formatted)
	if err != nil {
		return nil, err
	}

	return &Certificate{BrazilianCertificates: createdBrazilianCertificates}, nil
}
