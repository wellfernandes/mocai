package mocai

import (
	"errors"
	"fmt"
	"math/rand"

	ptbr "github.com/brazzcore/mocai/pkg/mocai/locale/pt-br/constants"
)

// Phone represents a mock phone entity.
type Phone struct {
	AreaCode string
	Number   string
}

// GeneratePhone generates a mock phone number with random data.
func GeneratePhone(locale string) (Phone, error) {
	areaCode, _ := GenerateAreaCode(locale)
	number, _ := GeneratePhoneNumber(locale)

	phoneCredted := Phone{
		AreaCode: areaCode,
		Number:   number,
	}

	return phoneCredted, nil
}

// GenerateAreaCode generates a random area code in Portuguese.
func GenerateAreaCode(locale string) (string, error) {
	switch locale {
	case "pt-br":
		areaCode := ptbr.AreaCodes
		return areaCode[rand.Intn(len(areaCode))], nil
	default:
		return "", errors.New("unsupported locale")
	}
}

// GeneratePhoneNumber generates a random phone number in Portuguese.
func GeneratePhoneNumber(locale string) (string, error) {
	switch locale {
	case "pt-br":
		phoneNumber := fmt.Sprintf("9%08d", rand.Intn(100000000))
		return phoneNumber, nil
	default:
		return "", errors.New("unsupported locale")
	}
}
