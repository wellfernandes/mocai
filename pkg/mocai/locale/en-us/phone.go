package enus

import (
	"math/rand"
)

// GenerateAreaCode generates a random area code in English.
func GenerateAreaCode() string {
	areaCodes := []string{"212", "310", "312", "713", "602"}
	return areaCodes[rand.Intn(len(areaCodes))]
}

// GeneratePhoneNumber generates a random phone number in English.
func GeneratePhoneNumber() string {
	// Simplified phone number generation
	return "555-1234"
}
