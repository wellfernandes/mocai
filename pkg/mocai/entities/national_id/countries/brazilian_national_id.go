package countries

// RG represents a Brazilian national ID.
type RG struct {
	Number string
	State  string
}

// GenerateBrazilianNationalID generates a valid Brazilian national ID [RG].
func GenerateBrazilianNationalID(formatted bool) (*RG, error) {
	return nil, ErrGeneratingBrazilianNationalID
}
