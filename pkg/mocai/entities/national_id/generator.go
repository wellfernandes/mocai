package national_id

import "github.com/brazzcore/mocai/pkg/mocai/entities/national_id/countries"

type NationalID struct {
	BrazilianRG *countries.RG
}

// GenerateNationalID generates a Brazilian national ID [RG].
// If formatted is true, the Brazilian national ID will be returned in the format XX.XXX.XXX-X.
// if formatted is false, the Brazilian national ID will be returned as a plain string.
func GenerateNationalID(formatted bool) (*NationalID, error) {
	rg, err := countries.GenerateBrazilianNationalID(formatted)
	if err != nil {
		return nil, err
	}

	return &NationalID{BrazilianRG: rg}, nil
}
