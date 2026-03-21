package cpf

import (
	"strconv"
	"strings"
)

// ValidateCPF checks if a CPF is valid.
// It accepts both formatted (xxx.xxx.xxx-xx) and unformatted (xxxxxxxxxxx) CPFs.
func ValidateCPF(cpf string) bool {
	// remove formatting (dots and dashes)
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// check if the CPF has 11 digits
	if len(cpf) != 11 {
		return false
	}

	// convert the CPF string to a slice of integers
	digits := make([]int, 11)
	for i, char := range cpf {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return false
		}
		digits[i] = digit
	}

	// check if all digits are the same (invalid CPF)
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

	// calculate the first check digit
	firstCheckDigit := CalculateCheckDigit(digits[:9], 10)
	if firstCheckDigit != digits[9] {
		return false
	}

	// calculate the second check digit
	secondCheckDigit := CalculateCheckDigit(digits[:10], 11)
	return secondCheckDigit == digits[10]
}
