---
sidebar_position: 3
---

# Roadmap

This page summarizes the current direction of Mocaí based on the code and public API in the `develop` branch.

## Current Status

Today, Mocaí already provides:

- A unified public entry point through `mocai.NewMocker(...)`
- A `MockGenerator` interface for dependency injection
- Support for realistic mock generation in Brazilian Portuguese (`ptbr`)
- Generators for person, gender, address, phone, company, CPF, certificate, national ID, and voter registration
- Functional options for formatting, language selection, random source control, and custom providers

## Direction of the Project

The current architecture suggests three clear directions for evolution:

### 1. Broader Language Support

The `Mocker` is configured by language, and the translation layer already isolates locale-specific behavior. The current implementation is focused on `ptbr`, but the design leaves room for new languages in the future.

### 2. More Flexible Data Sources

The library already supports custom providers for:

- Address
- Person
- Company

This means Mocaí can evolve beyond static in-project data and integrate with external services or internal company fixtures without changing the public API.

### 3. Better Test Control

The option to inject a custom random source makes the library suitable for reproducible automated tests. This is an important direction for teams that need deterministic fixtures while keeping the same fluent API.

## What Is Stable Today

If you are adopting Mocaí now, these are the parts that should be considered the current usage model:

- `mocai.NewMocker(opts ...Option)` as the main entry point
- `WithLanguage("ptbr")` for locale selection
- `WithFormatted(true|false)` for document formatting control
- `MockGenerator` as the preferred interface boundary in application code
- Provider injection for custom integrations

## What to Expect When Contributing

The codebase is moving toward a more configurable and extensible mocking library rather than a collection of isolated generators. Contributions that align with this direction are especially relevant:

- New entities with the same `Mocker`-based usage style
- New language implementations
- Better translation coverage for errors and messages
- Additional provider-based integrations
- More examples and real-world usage documentation

## Practical Reading of the Roadmap

Mocaí is already useful in its current state. The near-term value is not in redefining the API again, but in expanding the coverage around the architecture that already exists:

- keep the API simple
- extend supported data safely
- preserve deterministic testing support
- improve documentation and examples