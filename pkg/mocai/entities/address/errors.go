package address

import "errors"

// Package address defines common errors used during address generation and validation.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrNoStreets = errors.New("no data available for streets")
	ErrNoCities  = errors.New("no data available for cities")
	ErrNoStates  = errors.New("no data available for states")
	ErrNoUFs     = errors.New("no data available for UFs")
	ErrNoZips    = errors.New("no data available for zips")
)
