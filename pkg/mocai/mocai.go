package mocai

// Mocai represents a complete mock entity.
type Mocai struct {
	Person  Person
	Address Address
	Phone   Phone
}

// GenerateMocai generates a complete mock entity with random data.
func GenerateMocai(locale string) Mocai {
	return Mocai{
		Person:  GeneratePerson(locale),
		Address: GenerateAddress(locale),
		Phone:   GeneratePhone(locale),
	}
}
