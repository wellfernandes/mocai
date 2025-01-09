package ptbr

import (
	"math/rand"
)

// GenerateAreaCode generates a random area code in Portuguese.
func GenerateAreaCode() string {
	areaCodes := []string{"11", "21", "31", "41", "51"}
	return areaCodes[rand.Intn(len(areaCodes))]
}

// GeneratePhoneNumber generates a random phone number in Portuguese.
func GeneratePhoneNumber() string {
	// Simplified phone number generation
	return "91234-5678"
}
