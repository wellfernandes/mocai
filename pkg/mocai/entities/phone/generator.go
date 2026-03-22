package phone

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Phone represents a phone number with area code and number.
type Phone struct {
	AreaCode string
	Number   string
}

// NewPhone generates a mock phone number using customizable lists, language, and random source
func NewPhone(lang string, rnd translations.RandSource) (*Phone, error) {
	return generatePhone(lang, rnd)
}

func generatePhone(lang string, rnd translations.RandSource) (*Phone, error) {
	areaCodes := translations.GetList(lang, "phone_area_code")
	if len(areaCodes) == 0 {
		return nil, fmt.Errorf("%s", translations.Get(lang, "no_data_available_for_area_codes"))
	}
	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}

	areaCode := areaCodes[rnd.Intn(len(areaCodes))]
	number := fmt.Sprintf("9%08d", rnd.Intn(100000000))
	if areaCode == "" {
		return nil, fmt.Errorf("%s", translations.Get(lang, "error_generating_phone"))
	}
	return &Phone{
		AreaCode: areaCode,
		Number:   number,
	}, nil
}
