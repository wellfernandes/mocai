---
sidebar_position: 1
---

# Introduction

Mocaí is a Go library for generating realistic mock data for tests, prototypes, and local development workflows.

Its main goal is to reduce the cost of creating valid-looking fixture data while keeping the API small and practical. Instead of manually assembling people, addresses, companies, documents, and phone numbers in every test, you configure a `Mocker` and generate them on demand.

## What the Library Solves

Mocaí helps when you need:

- realistic fake data for automated tests
- deterministic generation for reproducible test runs
- a simple API that can be injected behind an interface
- extension points for custom data providers
- Brazilian Portuguese mock data for common business entities

## Current Public API Shape

The current API is centered around:

- `mocai.NewMocker(opts ...Option)`
- the `MockGenerator` interface
- functional options such as `WithLanguage`, `WithFormatted`, and `WithRandSource`

This keeps the entry point small while allowing more advanced customization when needed.

## Supported Entities

The library currently generates:

- Person
- Gender
- Address
- Phone
- Company
- CPF
- Certificate
- National ID (RG)
- Voter Registration

## Current Locale Support

At the moment, Mocaí is focused on `ptbr` data generation.

The architecture already uses a language-aware configuration model, so the project is prepared to grow into additional locales later.

For the current country-specific context, see [Brazil](./countries/brazil).

## Practical Design Goals

The codebase in `develop` reflects a few clear goals:

- keep mock generation simple to adopt
- avoid global mutable configuration in normal usage
- support dependency injection with `MockGenerator`
- allow provider-based customization for selected entities
- support deterministic test data through custom random sources

## Recommended Reading Order

If you are new to the library, this is the best reading path in this site:

1. Introduction
2. Quickstart
3. Get started
4. Core concepts
5. Supported entities
6. Brazil
7. Customization and extensibility

## Repository

- [GitHub repository](https://github.com/brazzcore/mocai)
- [MIT License](https://github.com/brazzcore/mocai/blob/develop/LICENSE)