package certificate

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate/countries"
)

// Certificate represents all certificates available
type Certificate struct {
	Brazil *countries.BrazilianCertificates
}

// NewCertificate creates a new Certificate with generated data.
func NewCertificate(isFormatted bool) (*Certificate, error) {
	cert, err := (&Certificate{}).generate(isFormatted)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

// GenerateCertificate generates all certificates available.
// If formatted is true, returns the number with separators (-).
// Returns a pointer to Certificate and error if any validation fails
func (c *Certificate) generate(formatted bool) (*Certificate, error) {
	createdBrazilianCertificates, err := countries.NewBrazilCertificates(formatted)
	if err != nil {
		return nil, err
	}
	return &Certificate{Brazil: createdBrazilianCertificates}, nil
}
