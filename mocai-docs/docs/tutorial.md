---
sidebar_position: 1
---

# Basic Tutorial
Welcome to the **Mocaí basic tutorial**! This guide will help you get started with Mocaí in just a few minutes.

## Mocaí 🚀
Welcome to **Mocaí**, an open-source library in Go designed to simplify the generation of mocks for entities such as Person, Address, Phone, and many others. Our goal is to make application development and testing more efficient by providing random yet consistent data that simulates real-world scenarios in a practical and reliable way.

### What is Mocaí used for?
Mocaí is a powerful tool for developers who need to create test data quickly and in an organized manner. With it, you can generate mocks of common entities, saving time and effort in manually creating fictitious data. Whether for unit tests, integration, or developing new features, Mocaí is here to help.

### Curiosity about the Name
The name Mocaí is a tribute to the Brazilian initiative behind the library. It originated from the combination of "mock" (the English term for simulation or fictitious data) with "açaí," a typical fruit from the Brazilian Amazon, known for its energy and versatility. Just as açaí is essential for many Brazilians, Mocaí aims to be an essential tool for developers who need efficient and high-quality test data. 🇧🇷

## Main Features
- **Random Data Generation:** Create mocks of entities with varied and realistic data.

- **Consistency:** Ensure that the generated data is consistent and suitable for testing.

- **Ease of Use:** Simple and intuitive interface for quick integration into your projects.

- **Extensibility:** Add new entities or customize existing ones according to your needs.

- **Open Source:** Collaborate, suggest improvements, and contribute to the growth of the library.

## Why use Mocaí?
- **Productivity:** Reduce the time spent on creating test data.

- **Quality:** Improve the coverage and effectiveness of your tests with realistic data.

- **Flexibility:** Adapt the mocks to the specific needs of your project.

- **Community:** Be part of an open-source community that values collaboration and innovation.

### How to Get Started
To get started with Mocaí, simply follow the steps below:

1. Install the library:
```
  go get github.com/brazzcore/mocai
```

2. Import the translation package:
```
  import "github.com/brazzcore/mocai/pkg/mocai/translations"
```

3. Import the package of the entity you want to generate:
```
   "github.com/brazzcore/mocai/pkg/mocai/entities/person"
```

4. Set the language that will be used to generate the data:
```
  // Set the language to pt-BR
  translations.SetLanguage("ptbr")
```

5. Create the entity, in this example we are creating a person:
```
  // Generate person mock data
  person_mock, err := person.GeneratePerson()
  if err != nil {
    fmt.Print(err)
  }
```

6. To use the generated data, simply select it. In this example, we will only print the person data:
```
  // Print a person's data
  fmt.Print(person_mock.FirstNameMale, person_mock.LastName, person_mock.Gender, person_mock.Age)
```

### How to contribute?
**Mocaí** is an open source project, and your contribution is very welcome! Whether it's reporting bugs, suggesting new features, or submitting pull requests, your participation helps improve the library for everyone.

**Issues:** Report problems or suggest improvements here.

**Pull Requests:** Submit your contributions following the contribution guidelines.

## License
Mocaí is distributed under the **MIT License**. Please refer to the LICENSE file for more details.

**Mocaí: Generating mocks, simplifying tests, accelerating development**.

### GitHub Repository

🔗 Repository on GitHub: https://github.com/brazzcore/mocai

