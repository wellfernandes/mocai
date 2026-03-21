package cpf

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type CPF struct {
	Number string
}

// NewCPF generates a mock CPF using a custom language and random source
func NewCPF(lang string, isFormatted bool, rnd translations.RandSource) (*CPF, error) {
	if rnd == nil {
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	// generate the first 9 digits
	digits := make([]int, 9)
	for i := range digits {
		digits[i] = rnd.Intn(10)
	}

	// calculate the first check digit
	digits = append(digits, CalculateCheckDigit(digits, 10))

	// calculate the second check digit
	digits = append(digits, CalculateCheckDigit(digits, 11))

	// convert the digits to a string
	cpfNumber := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")
	if isFormatted {
		if len(cpfNumber) != 11 {
			return nil, fmt.Errorf("%w, %s", ErrInvalidCPF, translations.Get(lang, "invalid_cpf"))
		}
		return &CPF{Number: cpfNumber[:3] + "." + cpfNumber[3:6] + "." + cpfNumber[6:9] + "-" + cpfNumber[9:]}, nil
	}
	return &CPF{Number: cpfNumber}, nil
}
func CalculateCheckDigit(digits []int, weight int) int {
	sum := 0
	for _, digit := range digits {
		sum += digit * weight
		weight--
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
