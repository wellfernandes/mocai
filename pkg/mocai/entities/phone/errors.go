package phone

// Error constants used for phone number generation failures.
// These errors represent specific failure scenarios and should be wrapped with additional
// context when returned.
const (
	ERROR_GENERATING_PHONE = "error generating phone"
	ERROR_NO_AREA_CODES    = "no data available for area codes"
)
