package person

import (
	"math/rand"

	mocks "github.com/brazzcore/mocai/pkg/mocai/mocks/pt_br"
)

// GeneratePerson generates a mock person with random data.
func GeneratePerson() interface{} {
	// Choose random values from constants
	firstNameMale := mocks.FirstNamesMale[rand.Intn(len(mocks.FirstNamesMale))]
	firstNameFemale := mocks.FirstNamesFemale[rand.Intn(len(mocks.FirstNamesFemale))]
	lastName := mocks.LastNames[rand.Intn(len(mocks.LastNames))]

	createdPerson := map[string]interface{}{
		"first_name_male":   firstNameMale,
		"first_name_female": firstNameFemale,
		"last_name":         lastName,
		"age":               rand.Intn(80) + 18,
	}

	return createdPerson

}
