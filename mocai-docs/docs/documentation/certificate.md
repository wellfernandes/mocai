---
sidebar_position: 16
---

# Certificate

Use `NewCertificate()` to generate Brazilian civil-record certificate data.

## Generation Method

```go
certificate, err := mocker.NewCertificate()
```

## Return Shape

The top-level `Certificate` entity currently exposes:

- `Brazil`

And the Brazilian payload includes:

- `BirthCertificate`
- `MarriageCertificate`
- `DeathCertificate`

Each certificate variant is based on the same certificate structure and includes fields such as:

- `VitalRecordsOffice`
- `ArchiveCode`
- `ServiceType`
- `BirthYear`
- `Type`
- `BookNumber`
- `PageNumber`
- `TermNumber`
- `CheckDigits`
- `Number`

## Example

```go
certificate, err := mocker.NewCertificate()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Birth Certificate: %s\n", certificate.Brazil.BirthCertificate.Number)
fmt.Printf("Marriage Certificate: %s\n", certificate.Brazil.MarriageCertificate.Number)
fmt.Printf("Death Certificate: %s\n", certificate.Brazil.DeathCertificate.Number)
```

## Formatting

Certificate numbers can be returned formatted or unformatted depending on `WithFormatted(true|false)`.

## When To Use It

This entity is useful for:

- Brazil-specific document scenarios
- registry and civil-record workflows
- document masking and formatting tests
- systems that need multiple certificate variants in one fixture