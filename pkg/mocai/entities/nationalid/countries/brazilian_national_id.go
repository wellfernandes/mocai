package countries

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// ErrConvertingDigit is returned when a digit conversion fails during RG generation.
var ErrConvertingDigit = errors.New("nationalid: error converting digit")

// RG represents a Brazilian identity card (Registro Geral).
type RG struct {
	Number      string
	State       string
	IssuingBody string
}

// NewRGCustom generates a random RG (Brazilian Identity Card) using an injected random source and language.
// Returns an error if generation fails, instead of silently returning an empty RG.
func NewRGCustom(lang string, isFormatted bool, rnd translations.RandSource) (RG, error) {
	if lang == "" {
		lang = "ptbr"
	}
	rg, err := generateBrazilianNationalIDCustom(lang, isFormatted, rnd)
	if err != nil {
		return RG{}, err
	}
	return rg, nil
}

func generateBrazilianNationalIDCustom(lang string, formatted bool, rnd translations.RandSource) (RG, error) {
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
			return "", fmt.Errorf("%w: %s", ErrConvertingDigit, translations.Get(lang, "error_converting_digit"))
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
