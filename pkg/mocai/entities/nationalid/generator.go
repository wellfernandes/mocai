package nationalid

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/nationalid/countries"
)

type NationalID struct {
	BrazilianRG countries.RG
}

// NewNationalId creates a new NationalID entity with a generated Brazilian national ID [RG].
func NewNationalId(isFormatted bool) NationalID {
	n, err := (&NationalID{}).generateNationalID(isFormatted)
	if err != nil {
		return NationalID{}
	}
	return n
}

// GenerateNationalID generates a Brazilian national ID [RG].
// If formatted is true, the Brazilian national ID will be returned in the format XX.XXX.XXX-X.
// if formatted is false, the Brazilian national ID will be returned as a plain string.
func (n *NationalID) generateNationalID(formatted bool) (NationalID, error) {
	rg := countries.NewRG(formatted)
	return NationalID{BrazilianRG: rg}, nil
}
