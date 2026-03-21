package certificate

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate/countries"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Certificate represents all certificates available
type Certificate struct {
	Brazil *countries.BrazilianCertificates
}

// NewCertificate generates a mock certificate using a custom language and random source.
func NewCertificate(lang string, isFormatted bool, rnd translations.RandSource) (*Certificate, error) {
	return generateCertificate(lang, isFormatted, rnd)
}

func generateCertificate(lang string, formatted bool, rnd translations.RandSource) (*Certificate, error) {
	createdBrazilianCertificates, err := countries.NewBrazilCertificatesCustom(lang, formatted, rnd)
	if err != nil {
		return nil, err
	}
	return &Certificate{Brazil: createdBrazilianCertificates}, nil
}
