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
	supportedLang := lang
	if translations.GetUFMap(lang) == nil {
		supportedLang = "ptbr"
	}
	streets := translations.GetList(supportedLang, "address_street")
	cities := translations.GetList(supportedLang, "address_city")
	states := translations.GetList(supportedLang, "address_state")
	zips := translations.GetList(supportedLang, "address_zip")

	slicesToCheck := []SlicesToCheck{
		{streets, ErrNoStreets},
		{cities, ErrNoCities},
		{states, ErrNoStates},
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

	// select the index for state and UF in a paired manner
	stateIdx := rnd.Intn(len(states))
	state := states[stateIdx]

	var uf string

	ufMap := translations.GetUFMap(supportedLang)
	if ufMap != nil {
		uf = ufMap[state]
	}
	if uf == "" {
		ufs := translations.GetList(supportedLang, "address_uf")
		if len(ufs) > stateIdx {
			uf = ufs[stateIdx]
		} else if len(ufs) > 0 {
			uf = ufs[rnd.Intn(len(ufs))]
		} else {
			uf = ""
		}
	}

	street := streets[rnd.Intn(len(streets))]
	city := cities[rnd.Intn(len(cities))]
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
