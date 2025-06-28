package person

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cpf"
	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Person represents a mock person with first names, last name, and age.
type Person struct {
	FirstNameMale   string
	FirstNameFemale string
	LastName        string
	Gender          *gender.Gender
	Age             int
	CPF             string
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
		return nil, ErrNoFirstNames
	}

	if len(lastNames) == 0 {
		return nil, ErrNoLastNames
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

	// Generate a random CPF without a mask
	cpf, err := cpf.GenerateCPF(false)
	if err != nil {
		return nil, err
	}

	// Validate required fields
	if firstNameMale == "" || firstNameFemale == "" || lastName == "" {
		return nil, fmt.Errorf("%s: missing required data (firstNameMale: %s, firstNameFemale: %s, lastName: %s)",
			ErrGeneratingPerson, firstNameMale, firstNameFemale, lastName)
	}

	return &Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastName,
		Gender:          gender,
		Age:             rand.Intn(80) + 18,
		CPF:             cpf,
	}, nil
}
