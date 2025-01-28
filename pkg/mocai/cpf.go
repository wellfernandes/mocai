package mocai

import (
	"math/rand"
	"strconv"
)

type CpfNumber struct {
	Number string
}

// GenerateCPF generates a valid CPF
func GenerateValidCPF(local string) (CpfNumber, error) {
	// Generate the first 9 digits randomly
	cpfDigits := make([]int, 9)
	for i := range cpfDigits {
		cpfDigits[i] = rand.Intn(10)
	}

	// Calculate the first check digit
	checkDigit, err := calculateCheckDigit(cpfDigits, 10)
	if err != nil {
		return CpfNumber{}, err
	}
	cpfDigits = append(cpfDigits, checkDigit)

	// Calculate the second check digit
	secondCheckDigit, err := calculateCheckDigit(cpfDigits, 11)
	if err != nil {
		return CpfNumber{}, err
	}
	cpfDigits = append(cpfDigits, secondCheckDigit)

	// Convert the slice of digits to a string
	cpf := ""
	for _, digit := range cpfDigits {
		cpf += strconv.Itoa(digit)
	}

	return CpfNumber{Number: cpf}, nil
}

// calculateCheckDigit calculates a single CPF check digit
func calculateCheckDigit(digits []int, weight int) (int, error) {
	sum := 0
	for _, digit := range digits {
		sum += digit * weight
		weight--
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0, nil
	}
	return 11 - remainder, nil
}
