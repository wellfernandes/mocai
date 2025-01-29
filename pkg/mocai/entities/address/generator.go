package address

import (
	"math/rand"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// GenerateAddress generates a mock address with random data.
func GenerateAddress() interface{} {
	lang := translations.GetLanguage()

	// Get the list of streets, cities, states, and ZIP codes
	streets := strings.Split(translations.Get(lang, "address_street"), ",")
	cities := strings.Split(translations.Get(lang, "address_city"), ",")
	states := strings.Split(translations.Get(lang, "address_state"), ",")
	zips := strings.Split(translations.Get(lang, "address_zip"), ",")

	// Choose random values
	street := streets[rand.Intn(len(streets))]
	city := cities[rand.Intn(len(cities))]
	state := states[rand.Intn(len(states))]
	zip := zips[rand.Intn(len(zips))]

	createdAddress := map[string]interface{}{
		"street": street,
		"number": rand.Intn(1000) + 1,
		"city":   city,
		"state":  state,
		"zip":    zip,
	}

	return createdAddress
}
