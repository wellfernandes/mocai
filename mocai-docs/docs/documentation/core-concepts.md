---
sidebar_position: 2
---

# Core Concepts

This page explains the main concepts behind the current Mocaí API.

## Mocker

`Mocker` is the main runtime object used to generate data.

It stores configuration such as:

- language
- formatting behavior
- random source
- optional custom providers

Instead of relying on global configuration, you create a configured instance and reuse it.

```go
mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithFormatted(true),
)
```

## MockGenerator

`MockGenerator` is the interface implemented by `Mocker`.

This is useful when your code should depend on behavior rather than the concrete implementation.

```go
type MockGenerator interface {
    NewPerson() (*person.Person, error)
    NewGender() (*gender.Gender, error)
    NewCompany() (*company.Company, error)
    NewVoteRegistration() (*voteregistration.VoteRegistration, error)
    NewNationalID() (*nationalid.NationalID, error)
    NewAddress() (*address.Address, error)
    NewCPF() (*cpf.CPF, error)
    NewCertificate() (*certificate.Certificate, error)
    NewPhone() (*phone.Phone, error)
    Language() string
}
```

Use this interface in services, factories, and test helpers when you want easier mocking or dependency injection.

## Functional Options

Mocaí uses functional options to configure a `Mocker`.

Current options include:

- `WithLanguage(lang string)`
- `WithFormatted(formatted bool)`
- `WithRandSource(rnd translations.RandSource)`
- `WithAddressProvider(p AddressProvider)`
- `WithPersonProvider(p PersonProvider)`
- `WithCompanyProvider(p CompanyProvider)`

This pattern keeps the constructor stable even as new configuration knobs are added.

## Formatting

Document generators can return either formatted or unformatted values.

Example with CPF:

- formatted: `123.456.789-00`
- unformatted: `12345678900`

Formatting is controlled through:

```go
mocai.WithFormatted(true)
```

## Random Source

By default, Mocaí uses a safe random source internally.

If you need stable output between test runs, pass your own source:

```go
rnd := translations.NewSafeRandSource(myFixedRand)

mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithRandSource(rnd),
)
```

This is especially useful for snapshot tests, reproducible integration tests, and debugging.

## Language Handling

The `Mocker` is language-aware and exposes the current language through `Language()`.

```go
lang := mocker.Language()
```

At the moment, the implemented locale is `ptbr`.

## Error Handling

Every generation method returns `(*Entity, error)`.

This is important because some generators depend on localized lists or derived values and can fail if required data is unavailable or invalid.

The design encourages explicit error handling instead of silent fallbacks.