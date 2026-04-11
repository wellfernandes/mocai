---
sidebar_position: 11
---

# Gender

Use `NewGender()` when you only need the gender identity value, without generating a full person.

## Generation Method

```go
genderValue, err := mocker.NewGender()
```

## Return Shape

The `Gender` entity exposes:

- `Identity`

The implementation currently defines identities such as:

- `male`
- `female`
- `non-binary`
- `genderfluid`
- `agender`
- `two-spirit`
- `other`

## Example

```go
genderValue, err := mocker.NewGender()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Gender: %s\n", genderValue.Identity)
```

## When To Use It

This entity is useful when:

- you need only the gender field
- you want to avoid generating a full person structure
- you are testing isolated field validation or mapping logic