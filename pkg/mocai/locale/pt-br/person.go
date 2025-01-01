package ptbr

import (
	"math/rand"
)

// GenerateFirstName generates a random first name in Portuguese.
func GenerateFirstName() string {
	firstNames := []string{"João", "Maria", "Pedro", "Ana", "Carlos"}
	return firstNames[rand.Intn(len(firstNames))]
}

// GenerateLastName generates a random last name in Portuguese.
func GenerateLastName() string {
	lastNames := []string{"Silva", "Santos", "Oliveira", "Souza", "Pereira"}
	return lastNames[rand.Intn(len(lastNames))]
}
