package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	company "github.com/brazzcore/mocai/pkg/mocai/entities/company/brazil"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func GenerateMockExample() {
	// Set the language to pt-BR
	translations.SetLanguage("ptbr")

	// Generate mock data
	person_mock, err := person.GeneratePerson()
	if err != nil {
		fmt.Print(err)
	}

	address_mock, err := address.GenerateAddress()
	if err != nil {
		fmt.Print(err)
	}

	phone_mock, err := phone.GeneratePhone()
	if err != nil {
		fmt.Print(err)
	}

	company_mock, err := company.GenerateCompany()
	if err != nil {
		fmt.Println("Error generating company:", err)
		return
	}

	fmt.Println(constants.HeaderMain)
	fmt.Println(constants.SubHeader)

	fmt.Printf("Person: %s %s, %s, %d years old, CPF: %s\n",
		person_mock.FirstNameMale, person_mock.LastName, person_mock.Gender, person_mock.Age, person_mock.CPF)

	fmt.Printf("Company: %s, CNPJ: %s\n", company_mock.CompanyName, company_mock.CNPJ)

	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		address_mock.Street, address_mock.Number, address_mock.City, address_mock.State, address_mock.UF, address_mock.ZIP)
	fmt.Printf("Phone: (%s) %s\n", phone_mock.AreaCode, phone_mock.Number)

	fmt.Println(constants.Footer)
}
