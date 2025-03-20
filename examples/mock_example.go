package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/birth_certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
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

	company_mock, err := company.GenerateCompany(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	birth_certificate_mock, err := birth_certificate.GenerateBirthCertificate(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(constants.HeaderMain)
	fmt.Println(constants.SubHeader)

	fmt.Printf("Person: %s %s, %s, %d years old, CPF: %s\n",
		person_mock.FirstNameMale, person_mock.LastName, person_mock.Gender, person_mock.Age, person_mock.CPF)

	fmt.Printf("Birth Certificate: %s\n", birth_certificate_mock.BrazilianBirthCertificate.CertificateNumber)

	fmt.Printf("Company: %s, CNPJ: %s\n", company_mock.BrazilianCompany.CompanyName, company_mock.BrazilianCompany.CNPJ)

	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		address_mock.Street, address_mock.Number, address_mock.City, address_mock.State, address_mock.UF, address_mock.ZIP)
	fmt.Printf("Phone: (%s) %s\n", phone_mock.AreaCode, phone_mock.Number)

	fmt.Println(constants.Footer)
}
