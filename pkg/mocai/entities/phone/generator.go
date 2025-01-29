package phone

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// GeneratePhone generates a mock phone number with random data.
func GeneratePhone() interface{} {
	lang := translations.GetLanguage()

	// Get the list of area codes
	areaCodes := strings.Split(translations.Get(lang, "phone_area_code"), ",")

	// Choose a random area code
	areaCode := areaCodes[rand.Intn(len(areaCodes))]

	// Generate a random phone number
	number := fmt.Sprintf("9%08d", rand.Intn(100000000))

	createdPhone := map[string]interface{}{
		"area_code": areaCode,
		"number":    number,
	}

	return createdPhone

}
