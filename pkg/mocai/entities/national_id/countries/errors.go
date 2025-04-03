package countries

import "errors"

var (
	ErrGeneratingBrazilianNationalID = errors.New("error generating brazilian national id")
	ErrBaseNumberExceedsLimit        = errors.New("base number exceeds 8 digits limit")
	ErrToConvertDigit                = errors.New("error converting digit")
)
