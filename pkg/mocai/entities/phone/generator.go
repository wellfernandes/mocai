package phone

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Phone represents a phone number with area code and number.
type Phone struct {
	AreaCode string
	Number   string
}

// GeneratePhone generates a mock phone number with random data.
// It returns a pointer to a Phone and an error if the generation fails.
func GeneratePhone() (*Phone, error) {
	lang := translations.GetLanguage()

	// Get the list of area codes
	areaCodeStr := translations.Get(lang, "phone_area_code")
	if areaCodeStr == "" {
		return nil, fmt.Errorf("%s for: %s", ErrNoAreaCodes, lang)
	}
	areaCodes := strings.Split(areaCodeStr, ",")

	// Validate data
	if len(areaCodes) == 0 {
		return nil, fmt.Errorf("%s for: %s", ErrNoAreaCodes, lang)
	}
	areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// Generate a random phone number
	number := fmt.Sprintf("9%08d", rand.Intn(100000000))

	if areaCode == "" || number == "" {
		return nil, fmt.Errorf("%s: missing required data (areaCode: %s, number: %s)",
			ErrGeneratingPhone, areaCode, number)
	}

	return &Phone{
		AreaCode: areaCode,
		Number:   number,
	}, nil
}
