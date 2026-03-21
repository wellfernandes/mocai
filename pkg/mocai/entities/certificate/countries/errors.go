package countries

import "errors"

// Errors for Certificate data generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
var (
	ErrInvalidCertificate              = errors.New("certificate: invalid certificate")
	ErrInvalidVitalRecordsOffice       = errors.New("certificate: invalid vital records office number")
	ErrInvalidArchive                  = errors.New("certificate: invalid archive number")
	ErrInvalidVitalRecordsService      = errors.New("certificate: invalid vital records service number")
	ErrInvalidBirthYear                = errors.New("certificate: invalid birth year")
	ErrInvalidBookNumber               = errors.New("certificate: invalid book number")
	ErrInvalidPageNumber               = errors.New("certificate: invalid page number")
	ErrInvalidTermNumber               = errors.New("certificate: invalid term number")
	ErrInvalidNumberWithoutCheckDigits = errors.New("certificate: invalid number without check digits")
)
