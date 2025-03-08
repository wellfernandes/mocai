package cpf

import (
	"strconv"
	"strings"
)

// ValidateCPF checks if a CPF is valid.
// It accepts both formatted (xxx.xxx.xxx-xx) and unformatted (xxxxxxxxxxx) CPFs.
func ValidateCPF(cpf string) bool {
	// Remove formatting (dots and dashes)
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Check if the CPF has 11 digits
	if len(cpf) != 11 {
		return false
	}

	// Convert the CPF string to a slice of integers
	digits := make([]int, 11)
	for i, char := range cpf {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	// Check if all digits are the same (invalid CPF)
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
	firstCheckDigit := calculateCheckDigit(digits[:9], 10)
	if firstCheckDigit != digits[9] {
		return false
	}

	// Calculate the second check digit
	secondCheckDigit := calculateCheckDigit(digits[:10], 11)
	return secondCheckDigit == digits[10]
}
