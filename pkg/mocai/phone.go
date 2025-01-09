package mocai

import (
	enus "github.com/wellfernandes/mocai/pkg/mocai/locale/en-us"
	ptbr "github.com/wellfernandes/mocai/pkg/mocai/locale/pt-br"
)

// Phone represents a mock phone entity.
type Phone struct {
	AreaCode string
	Number   string
}

// GeneratePhone generates a mock phone number with random data.
func GeneratePhone(locale string) Phone {
	var areaCode, number string

	switch locale {
	case "pt-br":
		areaCode = ptbr.GenerateAreaCode()
		number = ptbr.GeneratePhoneNumber()
	case "en-us":
		areaCode = enus.GenerateAreaCode()
		number = enus.GeneratePhoneNumber()
	default:
		areaCode = "000"
		number = "000-0000"
	}

	return Phone{
		AreaCode: areaCode,
		Number:   number,
	}
}
