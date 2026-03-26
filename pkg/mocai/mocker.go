package mocai

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/cpf"
	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/entities/nationalid"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/entities/voteregistration"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// MockGenerator defines the contract for generating mock data.
// Consumers should depend on this interface rather than the concrete Mocker struct,
// enabling easier testing and custom implementations.
type MockGenerator interface {
	NewPerson() (*person.Person, error)
	NewGender() (*gender.Gender, error)
	NewCompany() (*company.Company, error)
	NewVoteRegistration() (*voteregistration.VoteRegistration, error)
	NewNationalID() (*nationalid.NationalID, error)
	NewAddress() (*address.Address, error)
	NewCPF() (*cpf.CPF, error)
	NewCertificate() (*certificate.Certificate, error)
	NewPhone() (*phone.Phone, error)
	Language() string
}

// Compile-time check that Mocker implements MockGenerator.
var _ MockGenerator = (*Mocker)(nil)

// Mocker generates mock data in a specific language with configurable formatting and randomness.
type Mocker struct {
	lang      string
	formatted bool
	rnd       translations.RandSource

	// Optional providers for custom data generation (e.g., external APIs).
	addressProvider AddressProvider
	personProvider  PersonProvider
	companyProvider CompanyProvider
}

// NewMocker creates a new Mocker instance using functional options.
// By default, it uses "ptbr" language, no formatting, and a thread-safe random source.
//
// Example:
//
//	m := mocai.NewMocker(
//	    mocai.WithLanguage("ptbr"),
//	    mocai.WithFormatted(true),
//	)
func NewMocker(opts ...Option) MockGenerator {
	m := &Mocker{
		lang:      "ptbr",
		formatted: false,
		rnd:       translations.DefaultRandSource(),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// NewPerson generates a mock person using the configured language and random source.
// If a PersonProvider is set, it delegates to the provider.
func (m *Mocker) NewPerson() (*person.Person, error) {
	if m.personProvider != nil {
		return m.personProvider.GeneratePerson(m.lang, m.formatted)
	}
	return person.NewPerson(m.lang, m.formatted, m.rnd)
}

// NewGender generates a mock gender using the configured language and random source.
func (m *Mocker) NewGender() (*gender.Gender, error) {
	return gender.NewGender(m.lang, m.rnd)
}

// NewCompany generates a mock company using the configured language and random source.
// If a CompanyProvider is set, it delegates to the provider.
func (m *Mocker) NewCompany() (*company.Company, error) {
	if m.companyProvider != nil {
		return m.companyProvider.GenerateCompany(m.lang, m.formatted)
	}
	return company.NewCompany(m.lang, m.formatted, m.rnd)
}

// NewVoteRegistration generates a mock Voter ID using the configured language and random source.
func (m *Mocker) NewVoteRegistration() (*voteregistration.VoteRegistration, error) {
	return voteregistration.NewVoteRegistration(m.lang, m.formatted, m.rnd)
}

// NewNationalID generates a mock identity document using the configured language and random source.
func (m *Mocker) NewNationalID() (*nationalid.NationalID, error) {
	return nationalid.NewNationalID(m.lang, m.formatted, m.rnd)
}

// NewAddress generates a mock address using the configured language and random source.
// If an AddressProvider is set, it delegates to the provider.
func (m *Mocker) NewAddress() (*address.Address, error) {
	if m.addressProvider != nil {
		return m.addressProvider.GenerateAddress(m.lang)
	}
	return address.NewAddress(m.lang, m.rnd)
}

// NewCPF generates a mock CPF using the configured language and random source.
func (m *Mocker) NewCPF() (*cpf.CPF, error) {
	return cpf.NewCPF(m.lang, m.formatted, m.rnd)
}

// NewCertificate generates a mock certificate using the configured language and random source.
func (m *Mocker) NewCertificate() (*certificate.Certificate, error) {
	return certificate.NewCertificate(m.lang, m.formatted, m.rnd)
}

// NewPhone generates a mock phone using the configured language and random source.
func (m *Mocker) NewPhone() (*phone.Phone, error) {
	return phone.NewPhone(m.lang, m.rnd)
}

// Language returns the language used in this instance.
func (m *Mocker) Language() string {
	return m.lang
}
