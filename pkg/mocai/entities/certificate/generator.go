package certificate

import "github.com/brazzcore/mocai/pkg/mocai/entities/certificate/countries"

// Certificate represents all certificates available
type Certificate struct {
	BrazilianCertificates *countries.BrazilianCertificates
}

// GenerateCertificate generates all certificates available.
// If formatted is true, returns the number with separators (-).
// Returns a pointer to Certificate and error if any validation fails
func GenerateCertificate(formatted bool) (*Certificate, error) {
	createdBrazilianCertificates, err := countries.GenerateBrazilianCertificates(formatted)
	if err != nil {
		return nil, err
	}

	return &Certificate{BrazilianCertificates: createdBrazilianCertificates}, nil
}
