---
sidebar_position: 13
---

# Phone

Use `NewPhone()` to generate a phone fixture with area code and number.

## Generation Method

```go
phone, err := mocker.NewPhone()
```

## Return Shape

The `Phone` entity includes:

- `AreaCode`
- `Number`

## Example

```go
phone, err := mocker.NewPhone()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Phone: (%s) %s\n", phone.AreaCode, phone.Number)
```

## When To Use It

This entity is useful when you only need contact data and do not want the overhead of generating a larger structure.