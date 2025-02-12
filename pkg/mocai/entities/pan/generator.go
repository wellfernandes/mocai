package pan

import (
	"math/rand"
	"time"
)

// Function to generate a random PAN number
func GenerateRandomPAN() string {
	var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var digits = "0123456789"
	// Create a new source and random generator
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)

	// Generate first 5 letters
	pan := ""
	for i := 0; i < 5; i++ {
		pan += string(letters[rng.Intn(len(letters))])
	}

	// Generate 4 digits
	for i := 0; i < 4; i++ {
		pan += string(digits[rng.Intn(len(digits))])
	}

	// Generate last letter
	pan += string(letters[rng.Intn(len(letters))])

	return pan
}
