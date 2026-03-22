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

// Mocker represents an immutable collection of mock data in a specific language.
type Mocker struct {
	lang      string
	formatted bool
	rnd       translations.RandSource
}

// NewMocker creates a new Mocker instance with customizable language, formatting, and random source
func NewMocker(lang string, isFormatted bool, rnd translations.RandSource) *Mocker {
	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}
	return &Mocker{
		lang:      lang,
		formatted: isFormatted,
		rnd:       rnd,
	}
}

// NewPerson generates a mock person using a custom language and random source
func (m *Mocker) NewPerson() (*person.Person, error) {
	return person.NewPerson(m.lang, m.formatted, m.rnd)
}

// NewGender generates a mock gender using a custom language and random source
func (m *Mocker) NewGender() (*gender.Gender, error) {
	return gender.NewGender(m.lang, m.rnd)
}

// NewCompany generates a mock company using a custom language and random source
func (m *Mocker) NewCompany() (*company.Company, error) {
	return company.NewCompany(m.lang, m.formatted, m.rnd)
}

// NewVoteRegistration generates a mock Voter ID using a custom language and random source
func (m *Mocker) NewVoteRegistration() (*voteregistration.VoteRegistration, error) {
	return voteregistration.NewVoteRegistration(m.lang, m.formatted, m.rnd)
}

// NewNationalID generates a mock identity document using a custom language and random source
func (m *Mocker) NewNationalID() (*nationalid.NationalID, error) {
	return nationalid.NewNationalID(m.lang, m.formatted, m.rnd)
}

// NewAddress generates a mock address using a custom language and random source
func (m *Mocker) NewAddress() (*address.Address, error) {
	return address.NewAddress(m.lang, m.rnd)
}

// NewCPF generates a mock CPF using a custom language and random source
func (m *Mocker) NewCPF() (*cpf.CPF, error) {
	return cpf.NewCPF(m.lang, m.formatted, m.rnd)
}

// NewCertificate generates a mock certificate using a custom language and random source
func (m *Mocker) NewCertificate() (*certificate.Certificate, error) {
	return certificate.NewCertificate(m.lang, m.formatted, m.rnd)
}

// NewPhone generates a mock phone using a custom language and random source
func (m *Mocker) NewPhone() (*phone.Phone, error) {
	return phone.NewPhone(m.lang, m.rnd)
}

// GetLanguage returns the language used in this instance
func (m *Mocker) GetLanguage() string {
	return m.lang
}
