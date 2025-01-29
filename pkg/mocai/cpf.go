package mocai

import (
	cryptoRand "crypto/rand"
	"fmt"
	"strconv"
)

type CpfNumber struct {
	Number string
}

// GenerateCPF generates a valid CPF
func GenerateValidCPF(local string) (CpfNumber, error) {
	// Generate the first 9 digits using crypto/rand
	cpfDigits := make([]int, 9)
	randomBytes := make([]byte, 9)
	if _, err := cryptoRand.Read(randomBytes); err != nil {
		return CpfNumber{}, fmt.Errorf("failed to generate random digits: %w", err)
	}
	for i, b := range randomBytes {
		cpfDigits[i] = int(b) % 10
	}

	// Validate that not all digits are the same
	allSame := true
	for i := 1; i < len(cpfDigits); i++ {
		if cpfDigits[i] != cpfDigits[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return GenerateValidCPF(local) // Recursively try again
	}

	//fmt.Println("cpfDigits: ", cpfDigits)

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
