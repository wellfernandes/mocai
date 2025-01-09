package examples

import (
	"fmt"

	"github.com/wellfernandes/mocai/pkg/mocai"
)

// GenerateMockPTBR generates a mock in Portuguese.
func GenerateMockPTBR() {
	mock := mocai.GenerateMocai("pt-br")
	fmt.Println("=== Mock em Português (pt-br) ===")
	fmt.Println("Pessoa:")
	fmt.Printf("  Nome: %s %s\n", mock.Person.FirstName, mock.Person.LastName)
	fmt.Printf("  Idade: %d\n", mock.Person.Age)
	fmt.Printf("  CPF: %s\n", mock.Person.CPF)

	fmt.Println("\nEndereço:")
	fmt.Printf("  Rua: %s, %d\n", mock.Address.Street, mock.Address.Number)
	fmt.Printf("  Cidade: %s, %s\n", mock.Address.City, mock.Address.State)
	fmt.Printf("  CEP: %s\n", mock.Address.ZIPCode)

	fmt.Println("\nTelefone:")
	fmt.Printf("  Código de Área: %s\n", mock.Phone.AreaCode)
	fmt.Printf("  Número: %s\n", mock.Phone.Number)
}
