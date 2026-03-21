go get github.com/brazzcore/mocai
go run main.go

# Mocai

![mocai](img/mocai.svg)

#### [README-PT](/docs/localization/pt/README-PT.md)

A Go library for generating test data, allowing for the simple and efficient creation of mocks for entities.

## 📖 Description
**Mocaí** is an open-source Go library designed to simplify the generation of realistic mock data for entities such as Person, Address, Phone, Company, CPF, CNPJ, Certificates, National ID (RG), and Voter Registration. The goal is to make development and testing more efficient by providing random yet consistent data that simulates real-world scenarios in a practical and reliable way.

## 🌟 Curiosity about the Name
The name Mocaí is a tribute to the Brazilian initiative behind the library. It originated from the combination of "mock" (simulation or fake data) with "açaí," a typical fruit from the Brazilian Amazon, known for its energy and versatility. Just as açaí is essential for many Brazilians, Mocaí aims to be an essential tool for developers who need efficient and high-quality test data. 🇧🇷

## 🛠️ Main Features
- **Random Data Generation:** Generate mocks for entities with varied and realistic data (Person, Address, Phone, Company, CPF, CNPJ, Certificates, National ID, Voter Registration, and more).
- **Consistency:** Ensures generated data is consistent and suitable for testing.
- **Ease of Use:** Simple and intuitive API for quick integration.
- **Extensibility:** Easily add new entities or languages. The architecture is modular and ready for expansion.
- **Open Source:** Collaborate, suggest improvements, and contribute to the growth of the library.

## 📦 Supported Entities
- Person (with gender, age, CPF)
- Address (street, number, city, state, UF, ZIP)
- Phone (area code, number)
- Company (name, CNPJ)
- CPF (Brazilian individual taxpayer registry)
- CNPJ (Brazilian company registry)
- Certificates (Birth, Marriage, Death)
- National ID (RG)
- Voter Registration (Título de Eleitor)

## 🌐 Language Support
- **ptbr** (Brazilian Portuguese) is currently supported. The structure allows for easy addition of new languages in the future.

## 🚀 Why Use Mocaí?
- **Productivity:** Reduce the time spent creating test data.
- **Quality:** Improve test coverage and effectiveness with realistic data.
- **Flexibility:** Adapt mocks to your project's specific needs.
- **Community:** Be part of an open-source community that values collaboration and innovation.

## ⚙️ Requirements
- Go 1.23.4 or higher

## 🚀 How to Get Started
### Installation
To start using Mocaí, install the library with:

```sh
go get github.com/brazzcore/mocai
```

### Basic Usage
Import the library and generate mocks using the `Mocker` struct methods:

```go
package main

import (
    "fmt"
    "log"
    "github.com/brazzcore/mocai/pkg/mocai"
)

func main() {
    // Create a Mocker instance for Brazilian Portuguese
    mocker := mocai.NewMocker("ptbr", true, nil) // isFormatted: true for formatted docs (e.g., CPF/CNPJ)

    // Generate a mock address
    address, err := mocker.NewAddress()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Address: %s, %d - %s, %s (%s) - %s\n", address.Street, address.Number, address.City, address.State, address.UF, address.ZIP)

    // Generate a mock person
    person, err := mocker.NewPerson()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Person: %s %s, Gender: %s, Age: %d, CPF: %s\n", person.FirstNameMale, person.LastName, person.Gender.Identity, person.Age, person.CPF.Number)

    // Generate a mock company
    company, err := mocker.NewCompany()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Company: %s, CNPJ: %s\n", company.BrazilianCompany.Name, company.BrazilianCompany.CNPJ)
}
```

> **Note:** Each call to a method like `NewPerson()` or `NewAddress()` generates a new mock with random data. The `Mocker` instance is immutable regarding its configuration (language, formatting, random source).

#### About Languages
Currently, only "ptbr" is implemented. To support other languages, contribute with new translation and mock data files.

#### About Formatting
The `isFormatted` parameter controls whether documents like CPF/CNPJ are returned formatted (e.g., `123.456.789-00`) or as plain numbers (`12345678900`).

### Examples
The ***examples*** folder contains usage samples:

- `mocker/`: Example using the main entry point `mocai.NewMocker(lang string, isFormatted bool, rnd RandSource)` to generate mocks in a fluent and simplified way.

To run an example:
```sh
cd examples/mocker
go run main.go
```

## 🧪 Running Tests
To run all tests:
```sh
go test ./...
```

## 🤝 Contribute
Mocaí is an open-source project, and your contribution is very welcome! Whether it's reporting bugs, suggesting new features, or submitting pull requests, your participation helps improve the library for everyone.

### How to Contribute
1. **Report Issues:** Found a bug or have a suggestion? Open an issue.
2. **Submit Pull Requests:** Follow the contribution guidelines and submit your improvements. Always run tests before submitting.
3. **Discuss Ideas:** Join discussions and share your ideas for the project.

### Contribution Guidelines
1. Follow the project's coding standards.
2. Add tests for new features.
3. If necessary, document your changes in the **README.md**.

## 📄 License
Mocaí is distributed under the **MIT License**.

### 🌟 Mocaí: Generating mocks, simplifying tests, accelerating development.
Welcome to the project, and feel free to explore, use, and contribute!

## 🌐 Discord Server
Join the **Mocai** server on Discord:
[Click here and join us!](https://discord.gg/TFRnQBkAMt)
