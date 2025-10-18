package voteregistration

import "github.com/brazzcore/mocai/pkg/mocai/entities/voteregistration/countries"

// VoteRegistration represents a Brazilian vote registration.
type VoteRegistration struct {
	BrazilianVoteRegistration countries.BrazilianVoteRegistration
}

// NewVoteRegistration generates a Brazilian vote registration.
func NewVoteRegistration(isFormatted bool) VoteRegistration {
	vr, err := (&VoteRegistration{}).generateVoteRegistration(isFormatted)
	if err != nil {
		return VoteRegistration{}
	}
	return vr
}

// GenerateVoteRegistration generates a Brazilian vote registration.
// If formatted is true, the Brazilian vote registration will be returned in the format XXX XXX XXX.
// If formatted is false, the Brazilian vote registration will be returned as a plain string.
func (v *VoteRegistration) generateVoteRegistration(isFormatted bool) (VoteRegistration, error) {
	brazilianVoteRegistration := countries.NewBrazilianVoteRegistration(isFormatted)

	return VoteRegistration{BrazilianVoteRegistration: brazilianVoteRegistration}, nil
}
