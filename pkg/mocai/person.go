package mocai

import (
	"math/rand"

	enus "github.com/wellfernandes/mocai/pkg/mocai/locale/en-us"
	ptbr "github.com/wellfernandes/mocai/pkg/mocai/locale/pt-br"
)

// Person represents a mock person entity.
type Person struct {
	FirstName string
	LastName  string
	Age       int
	CPF       string
}

// GeneratePerson generates a mock person with random data.
func GeneratePerson(locale string) Person {
	var firstName, lastName string

	switch locale {
	case "pt-br":
		firstName = ptbr.GenerateFirstName()
		lastName = ptbr.GenerateLastName()
	case "en-us":
		firstName = enus.GenerateFirstName()
		lastName = enus.GenerateLastName()
	default:
		firstName = "Unknown"
		lastName = "Unknown"
	}

	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       rand.Intn(80) + 18,
		CPF:       GenerateCPF(),
	}
}

// GenerateCPF generates a random CPF (Brazilian ID number).
func GenerateCPF() string {
	// Simplified CPF generation
	return "123.456.789-00"
}
