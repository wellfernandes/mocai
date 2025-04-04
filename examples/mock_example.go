package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/national_id"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/entities/vote_registration"
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

	certificate_mock, err := certificate.GenerateCertificate(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	vote_registration_mock, err := vote_registration.GenerateVoteRegistration(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	nationalID_mock, err := national_id.GenerateNationalID(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(constants.HeaderMain)
	fmt.Println(constants.SubHeader)

	fmt.Printf("Person: %s %s, %s, %d years old, CPF: %s\n",
		person_mock.FirstNameMale, person_mock.LastName, person_mock.Gender, person_mock.Age, person_mock.CPF)

	fmt.Printf("RG: %s, State: %s, Issuing Body: %s\n",
		nationalID_mock.BrazilianRG.Number, nationalID_mock.BrazilianRG.State, nationalID_mock.BrazilianRG.IssuingBody)

	fmt.Printf("Birth Certificate: %s\n",
		certificate_mock.BrazilianCertificates.BirthCertificate.CertificateNumber)

	fmt.Printf("Marriage Certificate: %s\n",
		certificate_mock.BrazilianCertificates.MarriageCertificate.CertificateNumber)

	fmt.Printf("Death Certificate: %s\n",
		certificate_mock.BrazilianCertificates.DeathCertificate.CertificateNumber)

	fmt.Printf("Company: %s, CNPJ: %s\n",
		company_mock.BrazilianCompany.CompanyName, company_mock.BrazilianCompany.CNPJ)

	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		address_mock.Street, address_mock.Number, address_mock.City, address_mock.State, address_mock.UF, address_mock.ZIP)

	fmt.Printf("Phone: (%s) %s\n",
		phone_mock.AreaCode, phone_mock.Number)

	fmt.Printf("Voter Registration Card: Section: %s, Zone: %s, Registration: %s\n",
		vote_registration_mock.BrazilianVoteRegistration.Section, vote_registration_mock.BrazilianVoteRegistration.Zone,
		vote_registration_mock.BrazilianVoteRegistration.Number)

	fmt.Println(constants.Footer)
}
