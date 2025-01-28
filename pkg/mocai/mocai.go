package mocai

// Mocai represents a complete mock entity.
type Mocai struct {
	Person         Person
	Address        Address
	Phone          Phone
	RegistrationID RegistrationID
	CpfNumber      CpfNumber
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

	registrationID, err := GenerateRegistrationID(locale)
	if err != nil {
		panic(err)
	}

	cpf, err := GenerateValidCPF(locale)
	if err != nil {
		panic(err)
	}

	mocai := Mocai{
		Person:         person,
		Address:        address,
		Phone:          phone,
		RegistrationID: registrationID,
		CpfNumber:      cpf,
	}

	return mocai
}
