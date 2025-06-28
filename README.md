# Mocai

![mocai](img/mocai.svg)

#### [README-PT](/docs/localization/pt/README-PT.md)

A Go library for generating test data, allowing for the simple and efficient creation of mocks for entities.

## 📖 Description
**Mocaí** is an open-source library in Go designed to simplify the generation of mocks for entities such as Person, Address, Phone, and many others. Our goal is to make the development and testing of applications more efficient by providing random yet consistent data that simulates real-world scenarios in a practical and reliable manner.

## 🌟 Curiosity about the Name
The name Mocaí is a tribute to the Brazilian initiative behind the library. It originated from the combination of "mock" (the English term for simulation or fictitious data) with "açaí," a typical fruit from the Brazilian Amazon, known for its energy and versatility. Just as açaí is essential for many Brazilians, Mocaí aims to be an essential tool for developers who need efficient and high-quality test data. 🇧🇷

## 🛠️ Main Features
- **Random Data Generation:** Create mocks of entities with varied and realistic data.
- **Consistency:** Ensure that the generated data is consistent and suitable for testing.
- **Ease of Use:** Simple and intuitive interface for quick integration into your projects.
- **Extensibility:** Add new entities or customize existing ones according to your needs.
- **Open Source:** Collaborate, suggest improvements, and contribute to the growth of the library.

## 🚀 Why Use Mocaí?
- **Productivity:** Reduce the time spent on creating test data.
- **Quality:** Improve the coverage and effectiveness of your tests with realistic data.
- **Flexibility:** Adapt the mocks to the specific needs of your project.
- **Community:** Be part of an open-source community that values collaboration and innovation.


## 🚀 How to Get Started
### Installation
To start using Mocaí, install the library with the following command:

```
go get github.com/brazzcore/mocai
```

Basic Usage
Import the library into your project and start generating mocks:

```go
	// specific
	translations.SetLanguage("ptbr")
	
	addr, err := address.GenerateAddress()
	if err != nil {
		return err
	}

	addr.Street // returns street name
```

```go
	// mocker
	m := mocai.NewMocker("ptbr")
	
	m.Address().Street // returns street name
```
### Examples
The ***examples*** folder contains samples of how to use the library, organized by approach:

- `specific/`: Examples using direct calls to each of Mocaí's mockable entities (e.g., phone.GeneratePhone, etc).

- `mocker/`: Example using a main entry point `mocai.NewMocker("ptbr")` to generate mocks in a fluent and simplified way.

To run an example, navigate to the desired directory and run:
```
go run main.go
```

## 🤝 Contribute
Mocaí is an open-source project, and your contribution is very welcome! Whether it's reporting bugs, suggesting new features, or submitting pull requests, your participation helps improve the library for everyone.

### How to Contribute
1. **Report Issues:** Found a bug or have a suggestion? Open an issue.

2. **Submit Pull Requests:** Follow the contribution guidelines and submit your improvements.

3. **Discuss Ideas:** Join discussions and share your ideas for the project.

### Contribution Guidelines
1. Follow the project's coding standards.
2. Add tests for new features.
3. If necessary, document your changes in the **README.md**.

## 📄 Licence
Mocaí is distributed under the **MIT License**.

### 🌟 Mocaí: Generating mocks, simplifying tests, accelerating development.
Welcome to the project, and feel free to explore, use, and contribute!

## 🌐 Discord Server
Join the **Mocai** server on Discord:
[Click here and join us!](https://discord.gg/TFRnQBkAMt)
