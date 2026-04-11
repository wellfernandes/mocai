---
sidebar_position: 10
---

# Person

Use `NewPerson()` when you need a richer personal fixture instead of isolated pieces of data.

## Generation Method

```go
person, err := mocker.NewPerson()
```

## Return Shape

The generated `Person` includes:

- `FirstNameMale`
- `FirstNameFemale`
- `LastName`
- `Gender`
- `Age`
- `CPF`

This means a generated person already composes two other entity types:

- `Gender`
- `CPF`

## Example

```go
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
```

## Important Notes

- `Age` is currently generated in the adult range used by the library implementation.
- The entity stores both male and female first-name variants.
- CPF formatting follows the mocker configuration through `WithFormatted(true|false)`.

## When To Use It

`Person` is a good default entity for:

- user creation tests
- onboarding flows
- form population
- identity-heavy business scenarios