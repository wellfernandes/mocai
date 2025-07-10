package main

import (
	"fmt"

	"github.com/brazzcore/mocai/internal/cli"
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/national_id"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/entities/vote_registration"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

func main() {
	// Set the language to pt-BR
	translations.SetLanguage("ptbr")

	// Generate mock data
	p, err := person.GeneratePerson()
	if err != nil {
		fmt.Print(err)
		return
	}

	addr, err := address.GenerateAddress()
	if err != nil {
		fmt.Print(err)
		return
	}

	ph, err := phone.GeneratePhone()
	if err != nil {
		fmt.Print(err)
		return
	}

	c, err := company.GenerateCompany(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	cert, err := certificate.GenerateCertificate(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	vr, err := vote_registration.GenerateVoteRegistration(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	nid, err := national_id.GenerateNationalID(false)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Println(cli.HeaderMain)
	fmt.Println(cli.SubHeader)

	fmt.Printf("Person: %s %s, %s, %d years old, CPF: %s\n",
		p.FirstNameMale, p.LastName, *p.Gender, p.Age, p.CPF)

	fmt.Printf("RG: %s, State: %s, Issuing Body: %s\n",
		nid.BrazilianRG.Number, nid.BrazilianRG.State, nid.BrazilianRG.IssuingBody)

	fmt.Printf("Birth Certificate: %s\n",
		cert.BrazilianCertificates.BirthCertificate.CertificateNumber)

	fmt.Printf("Marriage Certificate: %s\n",
		cert.BrazilianCertificates.MarriageCertificate.CertificateNumber)

	fmt.Printf("Death Certificate: %s\n",
		cert.BrazilianCertificates.DeathCertificate.CertificateNumber)

	fmt.Printf("Company: %s, CNPJ: %s\n",
		c.BrazilianCompany.CompanyName, c.BrazilianCompany.CNPJ)

	fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
		addr.Street, addr.Number, addr.City, addr.State, addr.UF, addr.ZIP)

	fmt.Printf("Phone: (%s) %s\n",
		ph.AreaCode, ph.Number)

	fmt.Printf("Voter Registration Card: Section: %s, Zone: %s, Registration: %s\n",
		vr.BrazilianVoteRegistration.Section, vr.BrazilianVoteRegistration.Zone,
		vr.BrazilianVoteRegistration.Number)

	fmt.Println(cli.Footer)
}
