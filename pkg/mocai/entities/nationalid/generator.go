package nationalid

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/nationalid/countries"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// NationalID represents a national identity document.
type NationalID struct {
	BrazilianRG countries.RG
}

// NewNationalID generates a mock RG (Brazilian Identity Card) using a custom language and random source.
func NewNationalID(lang string, isFormatted bool, rnd translations.RandSource) (*NationalID, error) {
	return generateNationalID(lang, isFormatted, rnd)
}

func generateNationalID(lang string, isFormatted bool, rnd translations.RandSource) (*NationalID, error) {
	rg, err := countries.NewRGCustom(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}
	return &NationalID{BrazilianRG: rg}, nil
}
