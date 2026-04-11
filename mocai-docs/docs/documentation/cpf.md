---
sidebar_position: 15
---

# CPF

Use `NewCPF()` when you only need a CPF value and do not want to generate a full person entity.

## Generation Method

```go
cpfValue, err := mocker.NewCPF()
```

## Return Shape

The `CPF` entity exposes:

- `Number`

## Example

```go
cpfValue, err := mocker.NewCPF()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("CPF: %s\n", cpfValue.Number)
```

## Formatting

Formatting is controlled by the mocker configuration.

- formatted: `123.456.789-00`
- unformatted: `12345678900`

## When To Use It

Use this entity when a CPF is the only field required by the scenario, such as validation, masking, parsing, or persistence tests.