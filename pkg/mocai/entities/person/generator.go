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
	CPF             *cpf.CPF
}

func NewPerson(isFormatted bool) (*Person, error) {
	p, err := (&Person{}).generatePerson(isFormatted)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// GeneratePerson generates a mock person with random data.
// It returns a pointer to a Person and an error if the generation fails.
func (p *Person) generatePerson(isFormatted bool) (*Person, error) {
	lang := translations.GetLanguage()

	// Get the list of first names and last names
	firstNamesMale := strings.Split(translations.Get(lang, "person_first_name_male"), ",")
	firstNamesFemale := strings.Split(translations.Get(lang, "person_first_name_female"), ",")
	lastNames := strings.Split(translations.Get(lang, "person_last_name"), ",")

	// Validate data
	if len(firstNamesMale) == 0 || len(firstNamesFemale) == 0 {
		return nil, fmt.Errorf("%w, %s", ErrNoFirstNames, translations.Translate("no_data_available_for_first_names"))
	}

	if len(lastNames) == 0 {
		return nil, fmt.Errorf("%w, %s", ErrNoLastNames, translations.Translate("no_data_available_for_last_names"))
	}

	// Choose random values
	firstNameMale := firstNamesMale[rand.Intn(len(firstNamesMale))]
	firstNameFemale := firstNamesFemale[rand.Intn(len(firstNamesFemale))]
	lastName := lastNames[rand.Intn(len(lastNames))]

	// Generate a random gender
	gender, err := gender.NewGender()
	if err != nil {
		return nil, err
	}

	// Generate a random CPF without a mask
	cpf, err := cpf.NewCPF(isFormatted)
	if err != nil {
		return nil, err
	}

	// Validate required fields
	if firstNameMale == "" || firstNameFemale == "" || lastName == "" {
		return nil, fmt.Errorf("%w, %s", ErrGeneratingPerson, translations.Translate("error_generating_person"))
	}

	return &Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastName,
		Gender:          gender,
		Age:             rand.Intn(70) + 18,
		CPF:             cpf,
	}, nil
}
