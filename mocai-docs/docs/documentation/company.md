---
sidebar_position: 14
---

# Company

Use `NewCompany()` to generate a company fixture with a Brazilian company payload.

## Generation Method

```go
company, err := mocker.NewCompany()
```

## Return Shape

The top-level `Company` entity currently exposes:

- `BrazilianCompany`

And `BrazilianCompany` includes:

- `Name`
- `CNPJ`

## Example

```go
company, err := mocker.NewCompany()
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Company: %s, CNPJ: %s\n", company.BrazilianCompany.Name, company.BrazilianCompany.CNPJ)
```

## Important Notes

- CNPJ formatting follows `WithFormatted(true|false)`.
- You can override company generation with `WithCompanyProvider(...)`.

## When To Use It

`Company` is a good fit for:

- B2B workflows
- supplier/customer fixtures
- document-heavy back-office tests
- onboarding and account creation flows