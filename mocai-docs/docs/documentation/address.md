---
sidebar_position: 12
---

# Address

Use `NewAddress()` to generate a Brazilian-style address fixture.

## Generation Method

```go
address, err := mocker.NewAddress()
```

## Return Shape

The `Address` entity includes:

- `Street`
- `Number`
- `City`
- `State`
- `UF`
- `ZIP`

## Example

```go
address, err := mocker.NewAddress()
if err != nil {
    log.Fatal(err)
}

fmt.Printf(
    "Address: %s, %d - %s, %s (%s) - %s\n",
    address.Street,
    address.Number,
    address.City,
    address.State,
    address.UF,
    address.ZIP,
)
```

## Important Notes

- The current implementation is built around Brazilian data.
- `UF` is paired with the state when possible.
- You can replace the default generator with `WithAddressProvider(...)`.

## When To Use It

`Address` is useful for:

- checkout tests
- registration forms
- shipping or billing flows
- company or user profile fixtures