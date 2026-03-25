package mocai

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
)

// AddressProvider generates address data.
// The default implementation uses static mock data.
// Implement this interface to integrate with external APIs.
type AddressProvider interface {
	GenerateAddress(lang string) (*address.Address, error)
}

// PersonProvider generates person data.
// Implement this interface to provide custom person generation logic.
type PersonProvider interface {
	GeneratePerson(lang string, formatted bool) (*person.Person, error)
}

// CompanyProvider generates company data.
// Implement this interface to provide custom company generation logic.
type CompanyProvider interface {
	GenerateCompany(lang string, formatted bool) (*company.Company, error)
}
