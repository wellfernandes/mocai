package helper

import (
	"math/rand"
	"strconv"
	"time"
)

// Generate a random Aadhaar number
func GenerateAadhaarNumber() string {
	rand.Seed(time.Now().UnixNano())
	aadhaar := ""

	for i := 0; i < 12; i++ {
		aadhaar += strconv.Itoa(rand.Intn(10))
	}
	return aadhaar
}

// Function to generate a random PAN number
func GenerateRandomPAN() string {
	var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var digits = "0123456789"
	rand.Seed(time.Now().UnixNano())

	// Generate first 5 letters
	pan := ""
	for i := 0; i < 5; i++ {
		pan += string(letters[rand.Intn(len(letters))])
	}

	// Generate 4 digits
	for i := 0; i < 4; i++ {
		pan += string(digits[rand.Intn(len(digits))])
	}

	// Generate last letter
	pan += string(letters[rand.Intn(len(letters))])

	return pan
}
