package cnpj

import (
	"errors"
	"fmt"
	"strings"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// ErrInvalidCNPJLength is returned when the generated CNPJ does not have 14 digits.
var ErrInvalidCNPJLength = errors.New("cnpj: invalid CNPJ length")

// GenerateCNPJ generates a valid CNPJ number using a custom random source.
// If formatted is true, the CNPJ will be returned in the format XX.XXX.XXX/XXXX-XX.
// If formatted is false, the CNPJ will be returned as a plain string of 14 digits.
// If rnd is nil, a default random source will be used.
func GenerateCNPJ(formatted bool, rnd translations.RandSource) (string, error) {
	// Fallback to default if rnd is nil
	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}

	// Generate the first 12 digits
	digits := make([]int, 12)
	for i := 0; i < 8; i++ {
		digits[i] = rnd.Intn(10)
	}
	// Set the branch identifier to 0001 (common for new companies)
	digits[8], digits[9], digits[10], digits[11] = 0, 0, 0, 1

	// Calculate the first check digit
	digits = append(digits, calculateCNPJCheckDigit(digits))

	// Calculate the second check digit
	digits = append(digits, calculateCNPJCheckDigit(digits))

	// Convert the digits to a string
	cnpj := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")

	if len(cnpj) != 14 {
		return "", ErrInvalidCNPJLength
	}

	// Format the CNPJ if requested
	if formatted {
		return fmt.Sprintf("%s.%s.%s/%s-%s", cnpj[:2], cnpj[2:5], cnpj[5:8], cnpj[8:12], cnpj[12:]), nil
	}

	return cnpj, nil
}

// calculateCNPJCheckDigit calculates the check digit for a CNPJ.
func calculateCNPJCheckDigit(digits []int) int {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	if len(digits) == 13 {
		weights = append([]int{6}, weights...)
	}

	sum := 0
	for i, digit := range digits {
		sum += digit * weights[i]
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
