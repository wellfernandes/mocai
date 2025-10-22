package mocai

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/address"
	"github.com/brazzcore/mocai/pkg/mocai/entities/certificate"
	"github.com/brazzcore/mocai/pkg/mocai/entities/company"
	"github.com/brazzcore/mocai/pkg/mocai/entities/gender"
	"github.com/brazzcore/mocai/pkg/mocai/entities/nationalid"
	"github.com/brazzcore/mocai/pkg/mocai/entities/person"
	"github.com/brazzcore/mocai/pkg/mocai/entities/phone"
	"github.com/brazzcore/mocai/pkg/mocai/entities/voteregistration"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type Mocker struct {
	Address          *address.Address
	Certificate      *certificate.Certificate
	Company          *company.Company
	Gender           *gender.Gender
	NationalID       *nationalid.NationalID
	Person           *person.Person
	Phone            *phone.Phone
	VoteRegistration *voteregistration.VoteRegistration
}

// NewMocker initializes a Mocai instance with the specified language and formatting options
func NewMocker(lang string, isFormatted bool) (*Mocker, error) {
	err := translations.SetLanguage(lang)
	if err != nil {
		return nil, err
	}

	address, err := address.NewAddress()
	if err != nil {
		return nil, err
	}

	certificate, err := certificate.NewCertificate(isFormatted)
	if err != nil {
		return nil, err
	}

	company, err := company.NewCompany(isFormatted)
	if err != nil {
		return nil, err
	}

	gender, err := gender.NewGender()
	if err != nil {
		return nil, err
	}

	nationalID, err := nationalid.NewNationalId(isFormatted)
	if err != nil {
		return nil, err
	}

	person, err := person.NewPerson(isFormatted)
	if err != nil {
		return nil, err
	}

	phone, err := phone.NewPhone()
	if err != nil {
		return nil, err
	}

	voteRegistration, err := voteregistration.NewVoteRegistration(isFormatted)
	if err != nil {
		return nil, err
	}

	return &Mocker{
		Address:          address,
		Certificate:      certificate,
		Company:          company,
		Gender:           gender,
		NationalID:       nationalID,
		Person:           person,
		Phone:            phone,
		VoteRegistration: voteRegistration,
	}, nil

}

// SetLanguage changes the language used
func (m *Mocker) SetLanguage(lang string) error {
	err := translations.SetLanguage(lang)
	if err != nil {
		return err
	}
	return err
}

func (m *Mocker) GetLanguage() string {
	return translations.GetLanguage()
}
