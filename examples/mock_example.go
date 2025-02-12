package examples

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/constants"
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/adhar"
	"github.com/brazzcore/mocai/pkg/mocai/entities/pan"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func GenerateMockExample() {
	// set the language
	var languageID string
	fmt.Print("Enter your language : ")
	fmt.Scan(&languageID)
	translations.SetLanguage(languageID) //test for : ptbr,enin

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

	fmt.Println(constants.HeaderMain)
	fmt.Println(constants.SubHeader)

	fmt.Printf("Person: %s %s, %s, %d years old\n",
		person_mock.FirstNameMale, person_mock.LastName, person_mock.Gender, person_mock.Age)
	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		address_mock.Street, address_mock.Number, address_mock.City, address_mock.State, address_mock.UF, address_mock.ZIP)
	fmt.Printf("Phone: (%s) %s\n", phone_mock.AreaCode, phone_mock.Number)

	if languageID == "enin" {
		aadhar_card := adhar.GenerateAadhaarNumber()
		pan_card := pan.GenerateRandomPAN()
		fmt.Printf("Adhar card: %s\n", aadhar_card)
		fmt.Printf("Pan card: %s", pan_card)
	}

	fmt.Println(constants.Footer)
}
