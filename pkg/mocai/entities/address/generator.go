package address

import (
	"fmt"
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

// NewAddress generates a mock address using a custom language and a random source
func NewAddress(lang string, rnd translations.RandSource) (*Address, error) {
	addr, err := generateAddress(lang, rnd)
	if err != nil {
		return nil, err
	}
	return addr, nil
}

func generateAddress(lang string, rnd translations.RandSource) (*Address, error) {
	streets := translations.GetList(lang, "address_street")
	cities := translations.GetList(lang, "address_city")
	states := translations.GetList(lang, "address_state")
	ufs := translations.GetList(lang, "address_uf")
	zips := translations.GetList(lang, "address_zip")

	slicesToCheck := []SlicesToCheck{
		{streets, ErrNoStreets},
		{cities, ErrNoCities},
		{states, ErrNoStates},
		{ufs, ErrNoUFs},
		{zips, ErrNoZips},
	}

	for _, s := range slicesToCheck {
		if len(s.data) == 0 {
			return nil, s.err
		}
		for _, item := range s.data {
			if strings.TrimSpace(item) == "" {
				return nil, fmt.Errorf("%w: empty value in slice", s.err)
			}
		}
	}

	// performs a random selection using a custom random source
	street := streets[rnd.Intn(len(streets))]
	city := cities[rnd.Intn(len(cities))]
	state := states[rnd.Intn(len(states))]
	uf := ufs[rnd.Intn(len(ufs))]
	zip := zips[rnd.Intn(len(zips))]

	return &Address{
		Street: street,
		Number: rnd.Intn(9999),
		City:   city,
		State:  state,
		UF:     uf,
		ZIP:    zip,
	}, nil
}
