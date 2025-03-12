package interfaces

type CompanyStrategy interface {
	GenerateCompany() (Company, error)
}
