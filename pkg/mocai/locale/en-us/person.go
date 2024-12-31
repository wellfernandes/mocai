package enus

import (
	"math/rand"
)

// GenerateFirstName generates a random first name in English.
func GenerateFirstName() string {
	firstNames := []string{"John", "Jane", "Michael", "Emily", "David"}
	return firstNames[rand.Intn(len(firstNames))]
}

// GenerateLastName generates a random last name in English.
func GenerateLastName() string {
	lastNames := []string{"Smith", "Johnson", "Williams", "Brown", "Jones"}
	return lastNames[rand.Intn(len(lastNames))]
}
