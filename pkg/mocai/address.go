package mocai

import (
	"math/rand"

	enus "github.com/wellfernandes/mocai/pkg/mocai/locale/en-us"
	ptbr "github.com/wellfernandes/mocai/pkg/mocai/locale/pt-br"
)

// Address represents a mock address entity.
type Address struct {
	Street  string
	Number  int
	City    string
	State   string
	ZIPCode string
}

// GenerateAddress generates a mock address with random data.
func GenerateAddress(locale string) Address {
	var street, city, state, zipCode string

	switch locale {
	case "pt-br":
		street = ptbr.GenerateStreet()
		city = ptbr.GenerateCity()
		state = ptbr.GenerateState()
		zipCode = ptbr.GenerateZIPCode()
	case "en-us":
		street = enus.GenerateStreet()
		city = enus.GenerateCity()
		state = enus.GenerateState()
		zipCode = enus.GenerateZIPCode()
	default:
		street = "Unknown"
		city = "Unknown"
		state = "Unknown"
		zipCode = "Unknown"
	}

	return Address{
		Street:  street,
		Number:  rand.Intn(1000) + 1,
		City:    city,
		State:   state,
		ZIPCode: zipCode,
	}
}
