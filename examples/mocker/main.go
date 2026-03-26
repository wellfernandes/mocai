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

	// Create a Mocker using functional options
	m := mocai.NewMocker(
		mocai.WithLanguage("ptbr"),
		mocai.WithFormatted(true),
	)

	// Generate a mock person
	p, err := m.NewPerson()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Person: %s %s, Gender: %s, Age: %d, CPF: %s\n",
		p.FirstNameMale, p.LastName, p.Gender.Identity, p.Age, p.CPF.Number)

	// Generate a mock address
	addr, err := m.NewAddress()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		addr.Street, addr.Number, addr.City, addr.State, addr.UF, addr.ZIP)

	// Generate a mock company
	comp, err := m.NewCompany()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Company: %s, CNPJ: %s\n", comp.BrazilianCompany.Name, comp.BrazilianCompany.CNPJ)

	// Generate a mock CPF
	cpfVal, err := m.NewCPF()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CPF: %s\n", cpfVal.Number)

	// Generate a mock certificate
	cert, err := m.NewCertificate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Birth Certificate: %s\n", cert.Brazil.BirthCertificate.Number)

	// Generate a mock national ID (RG)
	nid, err := m.NewNationalID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("RG: %s - %s/%s\n", nid.BrazilianRG.Number, nid.BrazilianRG.IssuingBody, nid.BrazilianRG.State)

	// Generate a mock voter registration
	vr, err := m.NewVoteRegistration()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Voter Registration: %s (Section: %s, Zone: %s)\n",
		vr.BrazilianVoteRegistration.Number, vr.BrazilianVoteRegistration.Section, vr.BrazilianVoteRegistration.Zone)

	// Generate a mock phone
	ph, err := m.NewPhone()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Phone: (%s) %s\n", ph.AreaCode, ph.Number)

	// You can also use the MockGenerator interface for dependency injection:
	printLanguage(m)

	fmt.Println(cli.Footer)
}

func printLanguage(m mocai.MockGenerator) {
	fmt.Printf("\nLanguage: %s\n", m.Language())
}
