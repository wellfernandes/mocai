package countries

import (
	"fmt"
	"strconv"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type RG struct {
	Number      string
	State       string
	IssuingBody string
}

// NewRGCustom generates a random RG (Brazilian Identity Card) using an injected random source and custom language
func NewRGCustom(isFormatted bool, rnd translations.RandSource, lang ...string) RG {
	l := "pt_br"
	if len(lang) > 0 && lang[0] != "" {
		l = lang[0]
	}
	rg, err := generateBrazilianNationalIDCustom(isFormatted, rnd, l)
	if err != nil {
		return RG{}
	}
	return rg
}

func generateBrazilianNationalIDCustom(formatted bool, rnd translations.RandSource, lang string) (RG, error) {
	rgNumber, err := calculateSPRGDigitCustom(rnd, lang)
	if err != nil {
		return RG{}, err
	}
	if formatted {
		rgNumber = formatRG(rgNumber)
	}
	state := translations.Get(lang, "brazilian_rg_state")
	issuingBody := translations.Get(lang, "brazilian_rg_issuing_body")
	return RG{
		Number:      rgNumber,
		State:       state,
		IssuingBody: issuingBody,
	}, nil
}

func calculateSPRGDigitCustom(rnd translations.RandSource, lang string) (string, error) {
	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}
	base := rnd.Intn(100000000)
	baseStr := fmt.Sprintf("%08d", base)
	d := make([]int, 8)
	for i := range 8 {
		val, err := strconv.Atoi(string(baseStr[i]))
		if err != nil {
			return "", fmt.Errorf("%s", translations.Get(lang, "error_converting_digit"))
		}
		d[i] = val
	}
	wheights := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	for i := range 8 {
		sum += d[7-i] * wheights[i]
	}
	checkDigit := sum % 11
	var dvStr string
	if checkDigit == 10 {
		dvStr = "X"
	} else {
		dvStr = strconv.Itoa(checkDigit)
	}
	full := fmt.Sprintf("%s%s", baseStr, dvStr)
	return full, nil
}

func formatRG(rgNumber string) string {
	if len(rgNumber) != 9 {
		return rgNumber
	}

	return fmt.Sprintf("%s.%s.%s-%s", rgNumber[0:3], rgNumber[3:6], rgNumber[6:8], rgNumber[8:])
}
