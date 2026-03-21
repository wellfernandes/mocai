package person

import (
	"fmt"
	"math/rand"

	"github.com/brazzcore/mocai/pkg/mocai/entities/cpf"
	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Person represents a mock person with first names, last name, and age
type Person struct {
	FirstNameMale   string
	FirstNameFemale string
	LastName        string
	Gender          *gender.Gender
	Age             int
	CPF             *cpf.CPF
}

func NewPerson(lang string, isFormatted bool, rnd translations.RandSource) (*Person, error) {
	return generatePerson(lang, isFormatted, rnd)
}

// GeneratePerson generates a mock person with random data
// It returns a pointer to a Person and an error if the generation fails
func generatePerson(lang string, isFormatted bool, rnd translations.RandSource) (*Person, error) {
	firstNamesMale := translations.GetList(lang, "person_first_name_male")
	firstNamesFemale := translations.GetList(lang, "person_first_name_female")
	lastNames := translations.GetList(lang, "person_last_name")

	if len(firstNamesMale) == 0 || len(firstNamesFemale) == 0 {
		return nil, fmt.Errorf("%w, %s", ErrNoFirstNames, translations.Get(lang, "no_data_available_for_first_names"))
	}
	if len(lastNames) == 0 {
		return nil, fmt.Errorf("%w, %s", ErrNoLastNames, translations.Get(lang, "no_data_available_for_last_names"))
	}

	if rnd == nil {
		rnd = rand.New(rand.NewSource(int64(rand.Int())))
	}

	firstNameMale := firstNamesMale[rnd.Intn(len(firstNamesMale))]
	firstNameFemale := firstNamesFemale[rnd.Intn(len(firstNamesFemale))]
	lastName := lastNames[rnd.Intn(len(lastNames))]

	genderVal, err := gender.NewGender(lang, rnd)
	if err != nil {
		return nil, err
	}

	cpfVal, err := cpf.NewCPF(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}

	if firstNameMale == "" || firstNameFemale == "" || lastName == "" {
		return nil, fmt.Errorf("%w, %s", ErrGeneratingPerson, translations.Get(lang, "error_generating_person"))
	}

	return &Person{
		FirstNameMale:   firstNameMale,
		FirstNameFemale: firstNameFemale,
		LastName:        lastName,
		Gender:          genderVal,
		Age:             rnd.Intn(70) + 18,
		CPF:             cpfVal,
	}, nil
}
