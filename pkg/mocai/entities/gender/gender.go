package gender

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Gender represents a person's gender identity.
type Gender string

// Constants for different gender identities.
const (
	Male        Gender = "male"        // Male gender identity
	Female      Gender = "female"      // Female gender identity
	NonBinary   Gender = "non-binary"  // Non-binary gender identity
	GenderFluid Gender = "genderfluid" // Genderfluid identity
	Agender     Gender = "agender"     // Agender identity
	TwoSpirit   Gender = "two-spirit"  // Two-Spirit identity (used by some Indigenous peoples)
	Other       Gender = "other"       // Other gender identity
)

// GenerateRandomGender generates a random gender based on the current language.
func GenerateRandomGender() (*Gender, error) {
	lang := translations.GetLanguage()
	genderStr := translations.Get(lang, "gender")

	if genderStr == "" {
		return nil, fmt.Errorf("%s for: %s", ErrNoGenders, lang)
	}

	return (*Gender)(&genderStr), nil
}
