package translations

import (
	"math/rand"

	mocks "github.com/brazzcore/mocai/pkg/mocai/mocks/pt_br"
)

func init() {
	// Choose a random state
	state := mocks.States[rand.Intn(len(mocks.States))]

	// Gets the UF corresponding to the selected state
	uf := mocks.UFs[state]

	Register("pt", map[string]string{
		"person_first_name_male":   mocks.FirstNamesMale[rand.Intn(len(mocks.FirstNamesMale))],
		"person_first_name_female": mocks.FirstNamesFemale[rand.Intn(len(mocks.FirstNamesFemale))],
		"person_last_name":         mocks.LastNames[rand.Intn(len(mocks.LastNames))],
		"address_street":           mocks.Streets[rand.Intn(len(mocks.Streets))],
		"address_city":             mocks.Cities[rand.Intn(len(mocks.Cities))],
		"address_state":            state,
		"address_uf":               uf,
		"address_zip":              mocks.ZIPCodes[rand.Intn(len(mocks.ZIPCodes))],
		"phone_area_code":          mocks.AreaCodes[rand.Intn(len(mocks.AreaCodes))],
	})
}
