package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func GenerateMockExample() {
	// Set the language to pt-BR
	translations.SetLanguage("pt")

	// Generate mock data
	person_mock := person.GeneratePerson()
	address_mock := address.GenerateAddress()
	phone_mock := phone.GeneratePhone()

	fmt.Println("Pessoa:", person_mock)
	fmt.Println("Endereço:", address_mock)
	fmt.Println("Telefone:", phone_mock)
}
