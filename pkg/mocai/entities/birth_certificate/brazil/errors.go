package birth_certificate

import "errors"

// Package birth_certificate defines common errors used during birth certificate generation and validation.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrInvalidBirthCertificate         = errors.New("invalid birth certificate")
	ErrInvalidVitalRecordsOffice       = errors.New("invalid vital records office number")
	ErrInvalidArchive                  = errors.New("invalid archive number")
	ErrInvalidVitalRecordsService      = errors.New("invalid vital records service number")
	ErrInvalidBirthYear                = errors.New("invalid birth year")
	ErrInvalidBookNumber               = errors.New("invalid book number")
	ErrInvalidPageNumber               = errors.New("invalid page number")
	ErrInvalidTermNumber               = errors.New("invalid term number")
	ErrInvalidNumberWithoutCheckDigits = errors.New("invalid number without check digits")
)
