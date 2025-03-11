package cnpj

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateCNPJ generates a valid CNPJ number.
// If formatted is true, the CNPJ will be returned in the format XX.XXX.XXX/XXXX-XX.
// If formatted is false, the CNPJ will be returned as a plain string of 14 digits.
func GenerateCNPJ(formatted bool) (string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate the first 12 digits
	digits := make([]int, 12)
	for i := 0; i < 8; i++ {
		digits[i] = r.Intn(10)
	}
	// Set the branch identifier to 0001 (common for new companies)
	digits[8], digits[9], digits[10], digits[11] = 0, 0, 0, 1

	// Calculate the first check digit
	digits = append(digits, calculateCNPJCheckDigit(digits))

	// Calculate the second check digit
	digits = append(digits, calculateCNPJCheckDigit(digits))

	// Convert the digits to a string
	cnpj := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")

	// Format the CNPJ if requested
	if formatted {
		if len(cnpj) != 14 {
			return "", errors.New("invalid CNPJ length")
		}
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
