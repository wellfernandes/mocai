package gender

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type Gender struct {
	gender
}

// Gender represents a person's gender identity.
type gender string

// Constants for different gender identities.
const (
	Male        gender = "male"        // Male gender identity
	Female      gender = "female"      // Female gender identity
	NonBinary   gender = "non-binary"  // Non-binary gender identity
	GenderFluid gender = "genderfluid" // Genderfluid identity
	Agender     gender = "agender"     // Agender identity
	TwoSpirit   gender = "two-spirit"  // Two-Spirit identity (used by some Indigenous peoples)
	Other       gender = "other"       // Other gender identity
)

// NewGender creates a new Gender instance with a randomly generated gender.
func NewGender() Gender {
	g, err := (&Gender{}).generateRandomGender()
	if err != nil {
		return Gender{}
	}
	return g
}

// GenerateRandomGender generates a random gender based on the current language.
func (g *Gender) generateRandomGender() (Gender, error) {
	lang := translations.GetLanguage()
	genderStr := translations.Get(lang, "gender")

	if genderStr == "" {
		return Gender{}, fmt.Errorf("%s for: %s", ErrNoGenders, lang)
	}

	return Gender{gender: gender(genderStr)}, nil
}
