---
sidebar_position: 5
---

# Examples

This page collects compact examples based on the current `develop` branch usage style.

## Basic Mocker

```go
mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithFormatted(true),
)
```

## Generate a Person

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

## Generate an Address

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

## Generate a Company

```go
company, err := mocker.NewCompany()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Company: %s, CNPJ: %s\n", company.BrazilianCompany.Name, company.BrazilianCompany.CNPJ)
```

## Generate Documents

```go
cpfValue, err := mocker.NewCPF()
if err != nil {
    log.Fatal(err)
}

nationalID, err := mocker.NewNationalID()
if err != nil {
    log.Fatal(err)
}

voteRegistration, err := mocker.NewVoteRegistration()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("CPF: %s\n", cpfValue.Number)
fmt.Printf("RG: %s\n", nationalID.BrazilianRG.Number)
fmt.Printf("Voter Registration: %s\n", voteRegistration.BrazilianVoteRegistration.Number)
```

## Generate a Phone

```go
phone, err := mocker.NewPhone()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Phone: (%s) %s\n", phone.AreaCode, phone.Number)
```

## Depend On The Interface

```go
func printLanguage(generator mocai.MockGenerator) {
    fmt.Println(generator.Language())
}
```

## Use a Deterministic Random Source

```go
rnd := translations.NewSafeRandSource(myFixedRand)

mocker := mocai.NewMocker(
    mocai.WithLanguage("ptbr"),
    mocai.WithRandSource(rnd),
)
```