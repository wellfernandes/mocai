package countries

import "errors"

var (
	ErrGeneratingBrazilianNationalID = errors.New("error generating brazilian national id")
	ErrToConvertDigit                = errors.New("error converting digit")
)
