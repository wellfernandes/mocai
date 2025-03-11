package cnpj

import (
	"strconv"
	"strings"
)

// ValidateCNPJ checks if a CNPJ is valid.
// It accepts both formatted (XX.XXX.XXX/XXXX-XX) and unformatted (XXXXXXXXXXXXXX) CNPJs.
func ValidateCNPJ(cnpj string) bool {
	// Remove formatting (dots, slashes, and dashes)
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")

	// Check if the CNPJ has 14 digits
	if len(cnpj) != 14 {
		return false
	}

	// Convert the CNPJ string to a slice of integers
	digits := make([]int, 14)
	for i, char := range cnpj {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	// Check if all digits are the same (invalid CNPJ)
	allSame := true
	for i := 1; i < len(digits); i++ {
		if digits[i] != digits[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	// Calculate the first check digit
	firstCheckDigit := calculateCNPJCheckDigit(digits[:12])
	if firstCheckDigit != digits[12] {
		return false
	}

	// Calculate the second check digit
	secondCheckDigit := calculateCNPJCheckDigit(digits[:13])
	return secondCheckDigit == digits[13]
}
