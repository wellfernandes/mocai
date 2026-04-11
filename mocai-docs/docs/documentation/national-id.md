---
sidebar_position: 17
---

# National ID

Use `NewNationalID()` to generate a Brazilian RG fixture.

## Generation Method

```go
nationalID, err := mocker.NewNationalID()
```

## Return Shape

The `NationalID` entity currently exposes:

- `BrazilianRG`

And `BrazilianRG` includes:

- `Number`
- `State`
- `IssuingBody`

## Example

```go
nationalID, err := mocker.NewNationalID()
if err != nil {
    log.Fatal(err)
}

fmt.Printf(
    "RG: %s - %s/%s\n",
    nationalID.BrazilianRG.Number,
    nationalID.BrazilianRG.IssuingBody,
    nationalID.BrazilianRG.State,
)
```

## Formatting

RG formatting follows the `WithFormatted(true|false)` option when applicable.

## When To Use It

`NationalID` is relevant for identity workflows, document validation tests, and Brazil-specific registration scenarios.