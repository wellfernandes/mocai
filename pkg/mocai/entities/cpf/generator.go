package cpf

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateCPF generates a valid CPF number.
// If formatted is true, the CPF will be returned in the format xxx.xxx.xxx-xx.
// If formatted is false, the CPF will be returned as a plain string of 11 digits.
func GenerateCPF(formatted bool) (string, error) {
	// Create a local random generator with a unique seed
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	// Generate the first 9 digits
	digits := make([]int, 9)
	for i := range digits {
		digits[i] = r.Intn(10)
	}

	// Calculate the first check digit
	digits = append(digits, calculateCheckDigit(digits, 10))

	// Calculate the second check digit
	digits = append(digits, calculateCheckDigit(digits, 11))

	// Convert the digits to a string
	cpf := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(digits)), ""), "[]")

	// Format the CPF if requested
	if formatted {
		if len(cpf) != 11 {
			return "", errors.New(ErrInvalidCPF.Error())
		}
		return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:], nil
	}

	return cpf, nil
}

// calculateCheckDigit calculates the check digit for a CPF.
func calculateCheckDigit(digits []int, weight int) int {
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
