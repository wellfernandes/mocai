package birth_certificate

import (
	"regexp"
	"strings"
)

// ValidateBirthCertificate checks if a Brazilian birth certificate number is valid
// It validates the format, length, and check digits of the certificate number
func ValidateBirthCertificate(certificate string) bool {
	// Remove any formatting
	certificate = strings.ReplaceAll(certificate, "-", "")
	certificate = strings.ReplaceAll(certificate, " ", "")

	// Check if the certificate number has exactly 32 digits
	if len(certificate) != 32 {
		return false
	}

	// Verify that the certificate contains only numeric characters
	match, _ := regexp.MatchString("^[0-9]+$", certificate)
	if !match {
		return false
	}

	// Extract the number without check digits
	number := certificate[:30]

	// Get the informed check digits (last 2 digits)
	informedCheckDigit := certificate[30:]

	// Calculate the expected check digits
	calculatedCheckDigit := calculateCheckDigits(number)

	// Compare the informed check digits with the calculated ones
	return informedCheckDigit == calculatedCheckDigit
}
