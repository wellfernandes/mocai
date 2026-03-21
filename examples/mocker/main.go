package main

import (
	"fmt"

	"github.com/brazzcore/mocai/internal/cli"
	"github.com/brazzcore/mocai/pkg/mocai"
)

func main() {

	fmt.Println(cli.HeaderMain)
	fmt.Println(cli.SubHeader)

	m := mocai.NewMocker("ptbr", true, nil)

	address, err := m.NewAddress()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n",
			address.Street, address.Number, address.City, address.State, address.UF, address.ZIP)
	}

	certificate, _ := m.NewCertificate()
	fmt.Println("Brazilian birth certificate:", certificate.Brazil.BirthCertificate.Number)

	fmt.Println(cli.Footer)
}
