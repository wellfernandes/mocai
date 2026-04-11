---
sidebar_position: 4
---

# Customization And Extensibility

Mocaí is not limited to its default static generators. The current API already supports important extension points.

## Custom Providers

You can replace the default generation logic for selected entities by injecting custom providers.

Current provider interfaces:

```go
type AddressProvider interface {
    GenerateAddress(lang string) (*address.Address, error)
}

type PersonProvider interface {
    GeneratePerson(lang string, formatted bool) (*person.Person, error)
}

type CompanyProvider interface {
    GenerateCompany(lang string, formatted bool) (*company.Company, error)
}
```

This is useful when you want to integrate with:

- internal company fixtures
- external APIs
- seeded datasets
- domain-specific fake data sources

## Injecting Providers

```go
mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithAddressProvider(myAddressProvider),
    mocai.WithPersonProvider(myPersonProvider),
    mocai.WithCompanyProvider(myCompanyProvider),
)
```

When a provider is present, the matching `New...()` method delegates to it instead of using the default generator.

## Deterministic Generation

For tests that must be reproducible, use a custom random source.

```go
rnd := translations.NewSafeRandSource(myFixedRand)

mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithRandSource(rnd),
)
```

This is especially valuable for:

- CI pipelines
- golden-file tests
- debugging failing tests
- repeatable examples in documentation

## Choosing Between Default Generators And Providers

Use the built-in generators when:

- you need speed and simplicity
- static randomized data is enough
- you want low setup cost

Use custom providers when:

- your tests depend on domain-specific constraints
- you need data from an external system
- your team wants centralized control of fixture generation

## Dependency Injection With MockGenerator

For application code, the best extension point is often the `MockGenerator` interface itself.

```go
func CreateUser(generator mocai.MockGenerator) (*User, error) {
    person, err := generator.NewPerson()
    if err != nil {
        return nil, err
    }

    return &User{Name: person.FirstNameMale + " " + person.LastName}, nil
}
```

This keeps your code independent from the concrete `Mocker` type while still using the default implementation in production or tests.