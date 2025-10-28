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
	m, err := mocai.NewMocker("ptbr", true)
	if err != nil {
		log.Println(err)
	}
	// Generate mock data
	// You can select which entity attributes to use
	// by accessing the entity and its fields
	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		m.Address.Street, m.Address.Number, m.Address.City, m.Address.State, m.Address.UF, m.Address.ZIP)

	fmt.Println(cli.Footer)
}
