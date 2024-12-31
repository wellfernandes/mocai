package enus

import (
	"math/rand"
)

// GenerateStreet generates a random street name in English.
func GenerateStreet() string {
	streets := []string{"Main St", "Elm St", "Oak St", "Maple St", "Pine St"}
	return streets[rand.Intn(len(streets))]
}

// GenerateCity generates a random city name in English.
func GenerateCity() string {
	cities := []string{"New York", "Los Angeles", "Chicago", "Houston", "Phoenix"}
	return cities[rand.Intn(len(cities))]
}

// GenerateState generates a random state abbreviation in English.
func GenerateState() string {
	states := []string{"NY", "CA", "IL", "TX", "AZ"}
	return states[rand.Intn(len(states))]
}

// GenerateZIPCode generates a random ZIP code in English.
func GenerateZIPCode() string {
	// Simplified ZIP code generation
	return "12345"
}
