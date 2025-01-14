package ptbr

import (
	"math/rand"

	"github.com/wellfernandes/mocai/pkg/mocai/locale/pt-br/constants"
)

// GenerateStreet generates a random street name in Portuguese.
func GenerateStreet() string {
	streets := constants.Streets
	return streets[rand.Intn(len(streets))]
}

// GenerateCity generates a random city name in Portuguese.
func GenerateCity() string {
	cities := []string{"São Paulo", "Rio de Janeiro", "Belo Horizonte", "Curitiba", "Porto Alegre"}
	return cities[rand.Intn(len(cities))]
}

// GenerateState generates a random state abbreviation in Portuguese.
func GenerateState() string {
	states := []string{"SP", "RJ", "MG", "PR", "RS"}
	return states[rand.Intn(len(states))]
}

// GenerateZIPCode generates a random ZIP code in Portuguese.
func GenerateZIPCode() string {
	// Simplified ZIP code generation
	return "12345-678"
}
