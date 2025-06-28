package mocai

import (
	"fmt"

	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/entities/national_id"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/entities/vote_registration"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type Mocker struct {
	lang string
}

func NewMocker(lang string) *Mocker {

	err := translations.SetLanguage(lang)
	if err != nil {
		fmt.Println(err)
	}

	return &Mocker{
		lang: lang,
	}
}

func (m *Mocker) SetLanguage(lang string) error {
	return translations.SetLanguage(lang)
}

func (m *Mocker) GetLanguage() string {
	return translations.GetLanguage()
}

func (m *Mocker) Address() *address.Address {
	a, _ := address.GenerateAddress()

	return a
}

func (m *Mocker) Certificate() *certificate.Certificate {
	c, _ := certificate.GenerateCertificate(false)

	return c
}

func (m *Mocker) Company() *company.Company {
	c, _ := company.GenerateCompany(false)

	return c
}

func (m *Mocker) Gender() *gender.Gender {
	g, _ := gender.GenerateRandomGender()

	return g
}

func (m *Mocker) NationalID() *national_id.NationalID {
	n, _ := national_id.GenerateNationalID(false)

	return n
}

func (m *Mocker) Person() *person.Person {
	p, _ := person.GeneratePerson()

	return p
}

func (m *Mocker) Phone() *phone.Phone {
	p, _ := phone.GeneratePhone()

	return p
}

func (m *Mocker) VoteRegistration() *vote_registration.VoteRegistration {
	v, _ := vote_registration.GenerateVoteRegistration(false)

	return v
}
