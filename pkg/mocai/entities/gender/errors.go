package gender

import "errors"

// Package gender defines common errors used during gender generation and validation.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrNoGenders     = errors.New("no data available for genders")
	ErrInvalidGender = errors.New("invalid gender")
)
