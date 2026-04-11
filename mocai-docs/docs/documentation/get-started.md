---
sidebar_position: 1
---

# Get started

This guide focuses on the practical first-use flow for Mocaí.

## Installation

```bash
go get github.com/brazzcore/mocai
```

Module path:

```go
module github.com/brazzcore/mocai
```

Minimum Go version:

```text
go 1.23.4
```

## Recommended Entry Point

Use `github.com/brazzcore/mocai/pkg/mocai` and instantiate a mocker with functional options.

```go
package main

import (
	"log"

	"github.com/brazzcore/mocai/pkg/mocai"
)

func main() {
	mocker := mocai.NewMocker(
		mocai.WithLanguage("ptbr"),
		mocai.WithFormatted(true),
	)

	_, err := mocker.NewPerson()
	if err != nil {
		log.Fatal(err)
	}
}
```

## Current Language Support

Mocaí currently supports:

- `ptbr`

The API is structured so more languages can be added later, but the current implementation is centered on Brazilian Portuguese data.

## Formatting Behavior

When `WithFormatted(true)` is enabled, formatted documents are returned where applicable.

Examples:

- CPF formatted: `123.456.789-00`
- CPF unformatted: `12345678900`

## Reproducible Tests

If your tests require stable output, provide a custom random source:

```go
rnd := translations.NewSafeRandSource(myFixedRand)

mocker := mocai.NewMocker(
	mocai.WithLanguage("ptbr"),
	mocai.WithRandSource(rnd),
)
```

## Example

```go
person, err := mocker.NewPerson()
if err != nil {
	log.Fatal(err)
}

company, err := mocker.NewCompany()
if err != nil {
	log.Fatal(err)
}

cpfValue, err := mocker.NewCPF()
if err != nil {
	log.Fatal(err)
}

log.Printf("%s %s", person.FirstNameMale, person.LastName)
log.Printf("%s", company.BrazilianCompany.Name)
log.Printf("%s", cpfValue.Number)
```

## Where To Go Next

For deeper material, continue with:

- [Core concepts](./core-concepts)
- [Supported entities](./supported-entities)
- [Customization and extensibility](./customization-and-extensibility)
- [Examples](./examples)
