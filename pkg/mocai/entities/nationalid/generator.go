package nationalid

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/nationalid/countries"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type NationalID struct {
	BrazilianRG countries.RG
}

// NewNationalID generates a mock RG (Brazilian Identity Card) using a custom language and random source
func NewNationalID(lang string, isFormatted bool, rnd translations.RandSource) (*NationalID, error) {
	return generateNationalID(lang, isFormatted, rnd)
}

func generateNationalID(lang string, isFormatted bool, rnd translations.RandSource) (*NationalID, error) {
	rg := countries.NewRGCustom(isFormatted, rnd, lang)
	return &NationalID{BrazilianRG: rg}, nil
}
