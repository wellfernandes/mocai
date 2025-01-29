package person

import (
	"fmt"
	"math/rand"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
	mocks "github.com/brazzcore/mocai/pkg/mocai/mocks/pt_br"
)

// Person represents a mock person with first names, last name, and age.
type Person struct {
	FirstNameMale   string
	FirstNameFemale string
	LastName        string
	Age             int
}

// GeneratePerson generates a mock person with random data.
// It returns a pointer to a Person and an error if the generation fails.
func GeneratePerson() (*Person, error) {
	// Validate data
	if len(mocks.FirstNamesMale) == 0 || len(mocks.FirstNamesFemale) == 0 {
		return nil, fmt.Errorf(constants.ERROR_NO_FIRST_NAMES)
	}

	if len(mocks.LastNames) == 0 {
		return nil, fmt.Errorf(constants.ERROR_NO_LAST_NAMES)
	}

	// Choose random values from constants
	firstNameMale := mocks.FirstNamesMale[rand.Intn(len(mocks.FirstNamesMale))]
	firstNameFemale := mocks.FirstNamesFemale[rand.Intn(len(mocks.FirstNamesFemale))]
	lastName := mocks.LastNames[rand.Intn(len(mocks.LastNames))]

	if firstNameMale == "" || firstNameFemale == "" || lastName == "" {
		return nil, fmt.Errorf("%s: missing required data (firstNameMale: %s, firstNameFemale: %s, lastName: %s)",
			constants.ERROR_GENERATING_PERSON, firstNameMale, firstNameFemale, lastName)
	}

	createdPerson := &Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastName,
		Age:             rand.Intn(80) + 18,
	}

	return createdPerson, nil

}
