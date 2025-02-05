package translations

import (
	"math/rand"

	address_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/address/mocks/ptbr"
	gender_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/gender/mocks/ptbr"
	person_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/person/mocks/ptbr"
	phone_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/phone/mocks/ptbr"
)

func init() {
	// Choose a random state
	state := address_mocks.States[rand.Intn(len(address_mocks.States))]

	// Gets the UF corresponding to the selected state
	uf := address_mocks.UFs[state]

	Register("pt", map[string]string{
		"person_first_name_male":   person_mocks.FirstNamesMale[rand.Intn(len(person_mocks.FirstNamesMale))],
		"person_first_name_female": person_mocks.FirstNamesFemale[rand.Intn(len(person_mocks.FirstNamesFemale))],
		"person_last_name":         person_mocks.LastNames[rand.Intn(len(person_mocks.LastNames))],
		"gender":                   gender_mocks.Genders[rand.Intn(len(gender_mocks.Genders))],
		"address_street":           address_mocks.Streets[rand.Intn(len(address_mocks.Streets))],
		"address_city":             address_mocks.Cities[rand.Intn(len(address_mocks.Cities))],
		"address_state":            state,
		"address_uf":               uf,
		"address_zip":              address_mocks.ZIPCodes[rand.Intn(len(address_mocks.ZIPCodes))],
		"phone_area_code":          phone_mocks.AreaCodes[rand.Intn(len(phone_mocks.AreaCodes))],
	})
}
