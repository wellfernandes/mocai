---
sidebar_position: 18
---

# Voter Registration

Use `NewVoteRegistration()` to generate a Brazilian voter registration fixture.

## Generation Method

```go
voteRegistration, err := mocker.NewVoteRegistration()
```

## Return Shape

The `VoteRegistration` entity currently exposes:

- `BrazilianVoteRegistration`

And `BrazilianVoteRegistration` includes:

- `Section`
- `Zone`
- `Number`

## Example

```go
voteRegistration, err := mocker.NewVoteRegistration()
if err != nil {
    log.Fatal(err)
}

fmt.Printf(
    "Voter Registration: %s (Section: %s, Zone: %s)\n",
    voteRegistration.BrazilianVoteRegistration.Number,
    voteRegistration.BrazilianVoteRegistration.Section,
    voteRegistration.BrazilianVoteRegistration.Zone,
)
```

## Formatting

The voter registration number may be formatted depending on `WithFormatted(true|false)`.

## When To Use It

This entity is useful for Brazil-specific public-record or identity scenarios where voter-registration data is explicitly required.