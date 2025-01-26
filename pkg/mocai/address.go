package mocai

import (
	"errors"
	"math/rand"

	ptbr "github.com/brazzcore/mocai/pkg/mocai/locale/pt-br/constants"
)

// Address represents a mock address entity.
type Address struct {
	Street  string
	Number  int
	City    string
	State   string
	UF      string
	ZIPCode string
}

// GenerateAddress generates a mock address with random data.
func GenerateAddress(locale string) (Address, error) {
	street, err := GenerateStreet(locale)
	if err != nil {
		return Address{}, err
	}

	city, err := GenerateCity(locale)
	if err != nil {
		return Address{}, err
	}

	state, err := GenerateState(locale)
	if err != nil {
		return Address{}, err
	}

	uf, err := GenerateUF(locale, state)
	if err != nil {
		return Address{}, err
	}

	zipCode, err := GenerateZIPCode(locale)
	if err != nil {
		return Address{}, err
	}

	addressCredted := Address{
		Street:  street,
		Number:  rand.Intn(1000) + 1,
		City:    city,
		State:   state,
		UF:      uf,
		ZIPCode: zipCode,
	}

	return addressCredted, nil
}

// GenerateStreet generates a random street name in Portuguese.
func GenerateStreet(locale string) (string, error) {
	switch locale {
	case "pt-br":
		streets := ptbr.Streets
		return streets[rand.Intn(len(streets))], nil
	default:
		return "", errors.New("unsupported locale")
	}
}

// GenerateCity generates a random city name in Portuguese.
func GenerateCity(locale string) (string, error) {
	switch locale {
	case "pt-br":
		city := ptbr.Cities
		return city[rand.Intn(len(city))], nil
	default:
		return "", errors.New("unsupported locale")

	}
}

// GenerateState generates a random state abbreviation in Portuguese.
func GenerateState(locale string) (string, error) {
	switch locale {
	case "pt-br":
		states := ptbr.States
		return states[rand.Intn(len(states))], nil
	default:
		return "", errors.New("unsupported locale")
	}
}

// GenerateUF returns the UF (Unidade Federativa) abbreviation for a given state name in Portuguese.
func GenerateUF(locale, state string) (string, error) {
	var uf string

	if locale == "pt-br" {
		uf = ptbr.UF[state]
	} else {
		return "", errors.New("unsupported locale")
	}

	return uf, nil
}

// GenerateZIPCode generates a random ZIP code in Portuguese.
func GenerateZIPCode(locale string) (string, error) {
	switch locale {
	case "pt-br":
		zip_codes := ptbr.ZIPCodes
		return zip_codes[rand.Intn(len(zip_codes))], nil
	default:
		return "", errors.New("unsupported locale")
	}
}
