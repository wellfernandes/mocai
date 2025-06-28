package address

import (
	"fmt"
	"math/rand"
	"strings"

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

type SlicesToCheck struct {
	data []string
	err  error
}

// GenerateAddress generates a mock address with random data.
// It returns a pointer to an Address and an error if the generation fails.
func GenerateAddress() (*Address, error) {
	lang := translations.GetLanguage()

	// Get the list of streets, cities, states, and ZIP codes
	streets := strings.Split(translations.Get(lang, "address_street"), ",")
	cities := strings.Split(translations.Get(lang, "address_city"), ",")
	states := strings.Split(translations.Get(lang, "address_state"), ",")
	uf := translations.Get(lang, "address_uf")
	zips := strings.Split(translations.Get(lang, "address_zip"), ",")

	// slicesToCheck defines validation groups for related address components.
	// Each entry contains the data slice to validate and its specific error.
	slicesToCheck := []SlicesToCheck{
		{streets, ErrNoStreets},
		{cities, ErrNoCities},
		{states, ErrNoStates},
		{[]string{uf}, ErrNoUFs},
		{zips, ErrNoZips},
	}

	// Validate all required data slices
	// - Checks for empty slices
	// - Verifies each item isn't just whitespace
	for _, s := range slicesToCheck {
		if len(s.data) == 0 {
			return nil, s.err
		}

		// Individual item validation (prevent empty values)
		for _, item := range s.data {
			if strings.TrimSpace(item) == "" {
				return nil, fmt.Errorf("%w: empty value in slice", s.err)
			}
		}
	}

	// Choose random values
	street := streets[rand.Intn(len(streets))]
	city := cities[rand.Intn(len(cities))]
	state := states[rand.Intn(len(states))]
	zip := zips[rand.Intn(len(zips))]

	return &Address{
		Street: street,
		Number: rand.Intn(9999),
		City:   city,
		State:  state,
		UF:     uf,
		ZIP:    zip,
	}, nil
}
