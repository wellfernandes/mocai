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
	Address          address.Address
	Certificate      certificate.Certificate
	Company          company.Company
	Gender           gender.Gender
	NationalID       nationalid.NationalID
	Person           person.Person
	Phone            phone.Phone
	VoteRegistration voteregistration.VoteRegistration
}

func NewMocker(lang string, isFormatted bool) (*Mocker, error) {
	err := translations.SetLanguage(lang)
	if err != nil {
		return nil, err
	}

	return &Mocker{
		Address:          address.NewAddress(),
		Certificate:      certificate.NewCertificate(isFormatted),
		Company:          company.NewCompany(isFormatted),
		Gender:           gender.NewGender(),
		NationalID:       nationalid.NewNationalId(isFormatted),
		Person:           person.NewPerson(isFormatted),
		Phone:            phone.NewPhone(),
		VoteRegistration: voteregistration.NewVoteRegistration(isFormatted),
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
