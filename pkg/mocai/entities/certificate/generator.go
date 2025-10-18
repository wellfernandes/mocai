package certificate

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate/countries"
)

// Certificate represents all certificates available
type Certificate struct {
	BrazilianCertificates countries.BrazilianCertificates
}

// NewCertificate creates a new Certificate with generated data.
func NewCertificate(isFormatted bool) Certificate {
	cert, err := (&Certificate{}).generateCertificate(isFormatted)
	if err != nil {
		return Certificate{}
	}
	return cert
}

// GenerateCertificate generates all certificates available.
// If formatted is true, returns the number with separators (-).
// Returns a pointer to Certificate and error if any validation fails
func (c *Certificate) generateCertificate(formatted bool) (Certificate, error) {
	createdBrazilianCertificates, err := countries.GenerateBrazilianCertificates(formatted)
	if err != nil {
		return Certificate{}, err
	}

	return Certificate{BrazilianCertificates: createdBrazilianCertificates}, nil
}
