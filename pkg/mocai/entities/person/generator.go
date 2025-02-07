package person

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Person represents a mock person with first names, last name, and age.
type Person struct {
	FirstNameMale   string
	FirstNameFemale string
	LastName        string
	Gender          gender.Gender
	Age             int
}

// GeneratePerson generates a mock person with random data.
// It returns a pointer to a Person and an error if the generation fails.
func GeneratePerson() (*Person, error) {
	lang := translations.GetLanguage()

	// Get the list of first names and last names
	firstNamesMale := strings.Split(translations.Get(lang, "person_first_name_male"), ",")
	firstNamesFemale := strings.Split(translations.Get(lang, "person_first_name_female"), ",")
	lastNames := strings.Split(translations.Get(lang, "person_last_name"), ",")

	// Validate data
	if len(firstNamesMale) == 0 || len(firstNamesFemale) == 0 {
		return nil, errors.New(ERROR_NO_FIRST_NAMES)
	}

	if len(lastNames) == 0 {
		return nil, errors.New(ERROR_NO_LAST_NAMES)
	}

	// Choose random values
	firstNameMale := firstNamesMale[rand.Intn(len(firstNamesMale))]
	firstNameFemale := firstNamesFemale[rand.Intn(len(firstNamesFemale))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	// Generate a random gender
	gender, err := gender.GenerateRandomGender()
	if err != nil {
		return nil, err
	}

	// Validate required fields
	if firstNameMale == "" || firstNameFemale == "" || lastName == "" {
		return nil, fmt.Errorf("%s: missing required data (firstNameMale: %s, firstNameFemale: %s, lastName: %s)",
			ERROR_GENERATING_PERSON, firstNameMale, firstNameFemale, lastName)
	}

	createdPerson := &Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastName,
		Gender:          gender,
		Age:             rand.Intn(80) + 18,
	}

	return createdPerson, nil
}
