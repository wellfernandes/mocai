package mocai

// Mocai represents a complete mock entity.
type Mocai struct {
	Person  Person
	Address Address
	Phone   Phone
}

// GenerateMocai generates a complete mock entity with random data.
func GenerateMocai(locale string) Mocai {

	person, err := GeneratePerson(locale)
	if err != nil {
		panic(err)
	}

	address, err := GenerateAddress(locale)
	if err != nil {
		panic(err)
	}

	phone, err := GeneratePhone(locale)
	if err != nil {
		panic(err)
	}

	mocai := Mocai{
		Person:  person,
		Address: address,
		Phone:   phone,
	}

	return mocai
}
