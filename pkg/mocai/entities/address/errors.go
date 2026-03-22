package address

import "errors"

var (
	ErrNoStreets = errors.New("address: no streets available")
	ErrNoCities  = errors.New("address: no cities available")
	ErrNoStates  = errors.New("address: no states available")
	ErrNoZips    = errors.New("address: no zip codes available")
)
