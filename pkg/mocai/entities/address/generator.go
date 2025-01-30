package address

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"

	address_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/address/address_mocks/pt_br"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// Address represents a mock address with street, number, city, state, UF, and ZIP code.
type Address struct {
	Street string
	Number int
	City   string
	State  string
	UF     string
	ZIP    string
}

// GenerateAddress generates a mock address with random data.
// It returns a pointer to an Address and an error if the generation fails.
func GenerateAddress() (*Address, error) {
	lang := translations.GetLanguage()

	// Get the list of streets, cities, states, and ZIP codes
	streets := strings.Split(translations.Get(lang, "address_street"), ",")
	cities := strings.Split(translations.Get(lang, "address_city"), ",")
	states := strings.Split(translations.Get(lang, "address_state"), ",")
	zips := strings.Split(translations.Get(lang, "address_zip"), ",")

	// Validate data
	for _, slice := range [][]string{streets, cities, states, zips} {
		if len(slice) == 0 {
			return nil, fmt.Errorf("%s: empty data slice", ERROR_GENERATING_ADDRESS)
		}
	}

	if len(streets) == 0 {
		return nil, errors.New(ERROR_NO_STREETS)
	}

	if len(cities) == 0 {
		return nil, errors.New(ERROR_NO_CITIES)
	}

	if len(states) == 0 {
		return nil, errors.New(ERROR_NO_STATES)
	}

	if len(zips) == 0 {
		return nil, errors.New(ERROR_NO_ZIPS)
	}

	// Choose random values
	street := streets[rand.Intn(len(streets))]
	city := cities[rand.Intn(len(cities))]
	state := states[rand.Intn(len(states))]
	zip := zips[rand.Intn(len(zips))]

	// Get the UF from the state name
	uf, exists := address_mocks.UFs[state]
	if !exists {
		return nil, errors.New(ERROR_NO_UFS)
	}

	if street == "" || city == "" || state == "" || zip == "" || uf == "" {
		return nil, fmt.Errorf("%s: missing required data (street: %s, city: %s, state: %s, zip: %s, uf: %s)",
			ERROR_GENERATING_ADDRESS, street, city, state, zip, uf)
	}

	createdAddress := &Address{
		Street: street,
		Number: rand.Intn(9999),
		City:   city,
		State:  state,
		UF:     uf,
		ZIP:    zip,
	}

	return createdAddress, nil
}
