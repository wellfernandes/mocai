package mocai

import "github.com/brazzcore/mocai/pkg/mocai/translations"

// Option configures a Mocker instance.
type Option func(*Mocker)

// WithLanguage sets the language for mock data generation.
// Default: "ptbr"
func WithLanguage(lang string) Option {
	return func(m *Mocker) {
		m.lang = lang
	}
}

// WithFormatted controls whether documents are returned formatted.
// Default: false
func WithFormatted(formatted bool) Option {
	return func(m *Mocker) {
		m.formatted = formatted
	}
}

// WithRandSource sets a custom random source for deterministic generation.
// Default: translations.DefaultRandSource()
func WithRandSource(rnd translations.RandSource) Option {
	return func(m *Mocker) {
		m.rnd = rnd
	}
}

// WithAddressProvider sets a custom address provider.
// When set, NewAddress() delegates to this provider instead of the default static generator.
func WithAddressProvider(p AddressProvider) Option {
	return func(m *Mocker) {
		m.addressProvider = p
	}
}

// WithPersonProvider sets a custom person provider.
// When set, NewPerson() delegates to this provider instead of the default static generator.
func WithPersonProvider(p PersonProvider) Option {
	return func(m *Mocker) {
		m.personProvider = p
	}
}

// WithCompanyProvider sets a custom company provider.
// When set, NewCompany() delegates to this provider instead of the default static generator.
func WithCompanyProvider(p CompanyProvider) Option {
	return func(m *Mocker) {
		m.companyProvider = p
	}
}
