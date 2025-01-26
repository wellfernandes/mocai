package mocai

import (
	"fmt"
	"math/rand"

	ptbr "github.com/brazzcore/mocai/pkg/mocai/locale/pt-br/constants"
)

// Person represents a mock person entity.
type Person struct {
	FirstNameMale   string
	FirstNameFemale string
	LastName        string
	Age             int
	CPF             string
}

// GeneratePerson generates a mock person with random data.
func GeneratePerson(locale string) (Person, error) {
	firstNameMale, err := GenerateFirstNameMale(locale)
	if err != nil {
		return Person{}, err
	}

	firstNameFemale, err := GenerateFirstNameMale(locale)
	if err != nil {
		return Person{}, err
	}

	lastNameFemale, err := GenerateLastName(locale)
	if err != nil {
		return Person{}, err
	}

	personCredted := Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastNameFemale,
		Age:             rand.Intn(80) + 18,
		CPF:             GenerateCPF(),
	}

	return personCredted, nil
}

// GenerateFirstName generates a random first name in Portuguese.
func GenerateFirstNameMale(locale string) (string, error) {
	switch locale {
	case "pt-br":
		firstNameMale := ptbr.FirstNamesMale
		return firstNameMale[rand.Intn(len(firstNameMale))], nil
	default:
		return "", fmt.Errorf("unsupported locale")
	}
}

// GenerateLastName generates a random last name in Portuguese.
func GenerateLastName(locale string) (string, error) {
	switch locale {
	case "pt-br":
		lastNames := ptbr.FirstNamesMale
		return lastNames[rand.Intn(len(lastNames))], nil
	default:
		return "", fmt.Errorf("unsupported locale")
	}
}

// GenerateCPF generates a random CPF (Brazilian ID number).
func GenerateCPF() string {
	return "not implemented"
}
