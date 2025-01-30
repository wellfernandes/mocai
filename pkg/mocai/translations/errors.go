// Package translations provides standardized error messages and translations
// used across the mocai package. These constants ensure consistency and reusability
// throughout the application, especially in error handling and mock data generation.
//
// Example usage:
//
//	if !isSupported(lang) {
//	  return nil, errors.New(translations.ERROR_UNSUPPORTED_LANGUAGE)
//	}
package translations

import "errors"

// ErrUnsupportedLanguage indicates that the requested language is not supported.
var ErrUnsupportedLanguage = errors.New("unsupported language")
