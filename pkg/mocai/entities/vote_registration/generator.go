package vote_registration

import "github.com/brazzcore/mocai/pkg/mocai/entities/vote_registration/countries"

// VoteRegistration represents a Brazilian vote registration.
type VoteRegistration struct {
	BrazilianVoteRegistration *countries.BrazilianVoteRegistration
}

// GenerateVoteRegistration generates a Brazilian vote registration.
// If formatted is true, the Brazilian vote registration will be returned in the format XXX XXX XXX.
// If formatted is false, the Brazilian vote registration will be returned as a plain string.
func GenerateVoteRegistration(formatted bool) (*VoteRegistration, error) {
	brazilianVoteRegistration, err := countries.GenerateBrazilianVoteRegistration(formatted)
	if err != nil {
		return nil, err
	}

	return &VoteRegistration{BrazilianVoteRegistration: brazilianVoteRegistration}, nil
}
