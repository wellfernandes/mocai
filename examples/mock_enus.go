package examples

import (
	"fmt"

	"github.com/wellfernandes/mocai/pkg/mocai"
)

// GenerateMockENUS generates a mock in English.
func GenerateMockENUS() {
	mock := mocai.GenerateMocai("en-us")
	fmt.Println("=== Mock in English (en-us) ===")
	fmt.Println("Person:")
	fmt.Printf("  Name: %s %s\n", mock.Person.FirstName, mock.Person.LastName)
	fmt.Printf("  Age: %d\n", mock.Person.Age)
	fmt.Printf("  CPF: %s\n", mock.Person.CPF)

	fmt.Println("\nAddress:")
	fmt.Printf("  Street: %s, %d\n", mock.Address.Street, mock.Address.Number)
	fmt.Printf("  City: %s, %s\n", mock.Address.City, mock.Address.State)
	fmt.Printf("  ZIP Code: %s\n", mock.Address.ZIPCode)

	fmt.Println("\nPhone:")
	fmt.Printf("  Area Code: %s\n", mock.Phone.AreaCode)
	fmt.Printf("  Number: %s\n", mock.Phone.Number)
}
