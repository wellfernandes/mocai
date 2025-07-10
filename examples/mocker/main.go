package main

import (
	"fmt"
	"log"

	"github.com/brazzcore/mocai/internal/cli"
	"github.com/brazzcore/mocai/pkg/mocai"
)

func main() {

	fmt.Println(cli.HeaderMain)
	fmt.Println(cli.SubHeader)

	// Initialize mocai and set a supported language
	// If the language is not available, the default language pt-BR will be set
	m, err := mocai.NewMocker("ptbr")
	if err != nil {
		log.Println(err)
	}
	// Generate mock data
	// You can select which entity attributes to use
	// by accessing the entity and its fields
	m.Person()  // m.Person().Age, m.Person().CPF...
	m.Address() // m.Address().City, m.Address().Number...
	m.Phone()   // m.Phone().AreaCode, m.Phone().Number

	// Printing mocks example
	addr := m.Address()
	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		addr.Street, addr.Number, addr.City, addr.State, addr.UF, addr.ZIP)

	fmt.Println(cli.Footer)
}
