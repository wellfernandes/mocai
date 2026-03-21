package voteregistration

import (
	"github.com/brazzcore/mocai/pkg/mocai/entities/voteregistration/countries"
	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

type VoteRegistration struct {
	BrazilianVoteRegistration countries.BrazilianVoteRegistration
}

// NewVoteRegistration generates a Voter ID using a custom language and random source
func NewVoteRegistration(lang string, isFormatted bool, rnd translations.RandSource) (*VoteRegistration, error) {
	reg, err := generateVoteRegistration(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}
	return reg, nil
}

func generateVoteRegistration(lang string, isFormatted bool, rnd translations.RandSource) (*VoteRegistration, error) {
	reg, err := countries.NewBrazilianVoteRegistrationCustom(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}
	return &VoteRegistration{BrazilianVoteRegistration: *reg}, nil
}
