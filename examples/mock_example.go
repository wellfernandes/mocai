package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
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

	fmt.Println(constants.HeaderMain)
	fmt.Println(constants.SubHeader)

	fmt.Println("Person:", person_mock)
	fmt.Println("Address:", address_mock)
	fmt.Println("Phone:", phone_mock)

	fmt.Println(constants.Footer)
}
