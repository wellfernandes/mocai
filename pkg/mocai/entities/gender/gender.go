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

// NewGender generates a random gender using customizable lists, language, and random source
func NewGender(lang string, rnd translations.RandSource) (*Gender, error) {
	return generateRandomGender(lang, rnd)
}

func generateRandomGender(lang string, rnd translations.RandSource) (*Gender, error) {
	genders := translations.GetList(lang, "gender")
	if len(genders) == 0 {
		return nil, fmt.Errorf("%s", translations.Get(lang, "no_data_available_for_genders"))
	}
	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}
	genderStr := genders[rnd.Intn(len(genders))]
	return &Gender{Identity: Identity(genderStr)}, nil
}
