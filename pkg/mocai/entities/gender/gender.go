package gender

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Gender represents a person's gender.
type Gender struct {
	Identity
}

// Identity represents a person's gender identity.
type Identity string

// Constants for different gender identities.
const (
	Male        Identity = "male"        // Male gender identity
	Female      Identity = "female"      // Female gender identity
	NonBinary   Identity = "non-binary"  // Non-binary gender identity
	GenderFluid Identity = "genderfluid" // Genderfluid identity
	Agender     Identity = "agender"     // Agender identity
	TwoSpirit   Identity = "two-spirit"  // Two-Spirit identity (used by some Indigenous peoples)
	Other       Identity = "other"       // Other gender identity
)

// NewGender creates a new Gender instance with a randomly generated gender.
func NewGender() (*Gender, error) {
	g, err := (&Gender{}).generateRandomGender()
	if err != nil {
		return nil, err
	}
	return g, nil
}

// GenerateRandomGender generates a random gender based on the current language.
func (g *Gender) generateRandomGender() (*Gender, error) {
	lang := translations.GetLanguage()
	genderStr := translations.Get(lang, "gender")

	if genderStr == "" {
		return nil, fmt.Errorf("%w, %s", ErrNoGenders, translations.Translate("no_data_available_for_genders")+lang)
	}
	return &Gender{Identity: Identity(genderStr)}, nil
}
