---
sidebar_position: 2
---

# Quickstart

This page is the fastest path to using Mocaí correctly with the current API.

## Install

```bash
go get github.com/brazzcore/mocai
```

Mocaí currently requires Go `1.23.4` or newer.

## Create a Mocker

The recommended entry point is `mocai.NewMocker(...)`.

```go
package main

import (
    "fmt"
    "log"

    "github.com/brazzcore/mocai/pkg/mocai"
)

func main() {
    mocker := mocai.NewMocker(
        mocai.WithLanguage("ptbr"),
        mocai.WithFormatted(true),
    )

    person, err := mocker.NewPerson()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf(
        "Person: %s %s, Gender: %s, Age: %d, CPF: %s\n",
        person.FirstNameMale,
        person.LastName,
        person.Gender.Identity,
        person.Age,
        person.CPF.Number,
    )
}
```

## Generate More Than One Entity

The same configured mocker can generate multiple types of data:

```go
company, err := mocker.NewCompany()
if err != nil {
    log.Fatal(err)
}

address, err := mocker.NewAddress()
if err != nil {
    log.Fatal(err)
}

cpfValue, err := mocker.NewCPF()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Company: %s\n", company.BrazilianCompany.Name)
fmt.Printf("Address: %s, %d - %s\n", address.Street, address.Number, address.City)
fmt.Printf("CPF: %s\n", cpfValue.Number)
```

## Main Generation Methods

The current generator surface includes:

- `NewPerson()`
- `NewGender()`
- `NewCompany()`
- `NewVoteRegistration()`
- `NewNationalID()`
- `NewAddress()`
- `NewCPF()`
- `NewCertificate()`
- `NewPhone()`
- `Language()`

## Most Common Options

The most common configuration options are:

- `WithLanguage(lang string)`
- `WithFormatted(formatted bool)`

### Formatting

`WithFormatted(true)` affects documents such as CPF and CNPJ.

- Formatted: `123.456.789-00`
- Unformatted: `12345678900`

## Current Notes

- The configured language is stored in the `Mocker` instance and can be retrieved with `Language()`.
- Each call generates a new random value.
- Error messages are localized when translation coverage exists for the configured language.
- The library is currently focused on Brazilian Portuguese data generation.

## Next Steps

After this quickstart, continue with:

- [Get started](./documentation/get-started)
- [Core concepts](./documentation/core-concepts)
- [Supported entities](./documentation/supported-entities)