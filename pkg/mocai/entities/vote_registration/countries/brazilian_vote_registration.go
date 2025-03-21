package countries

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// BrazilianVoteRegistration represents a Brazilian vote registration.
type BrazilianVoteRegistration struct {
	Section string
	Zone    string
	Number  string
}

// GenerateBrazilianVoteRegistration generates a valid Brazilian vote registration number.
func GenerateBrazilianVoteRegistration(formatted bool) (*BrazilianVoteRegistration, error) {
	section := randomInt3Digits()
	if section == "" {
		return nil, ErrInvalidSection
	}

	zone := randomInt3Digits()
	if zone == "" {
		return nil, ErrInvalidZone
	}

	// Generate an 8 digit sequence number
	sequenceNumber := randomInt(1, 99999999)
	sequenceNumberStr := fmt.Sprintf("%08d", sequenceNumber)

	// Generate a random state code 01 to 28
	stateCode := randomInt(1, 28)
	stateCodeStr := fmt.Sprintf("%02d", stateCode)

	// Calculate the first check digit
	checkDigit1 := calculateCheckDigit1(sequenceNumberStr)
	if checkDigit1 == "" {
		return nil, ErrInvalidCheckDigit1
	}

	// Calculate the second check digit
	checkDigit2 := calculateCheckDigit2(stateCodeStr, checkDigit1, stateCode)
	if checkDigit2 == "" {
		return nil, ErrInvalidCheckDigit2
	}

	// Combine everything to form the complete number
	number := sequenceNumberStr + stateCodeStr + checkDigit1 + checkDigit2
	if number == "" {
		return nil, ErrInvalidNumber
	}

	if formatted {
		number = fmt.Sprintf("%s %s %s", number[:4], number[4:8], number[8:])
	}

	createdBrazilianVoteRegistration := &BrazilianVoteRegistration{
		Section: section,
		Zone:    zone,
		Number:  number,
	}

	return createdBrazilianVoteRegistration, nil
}

// randomInt generates a random integer between min and max
func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// randomInt3Digits generates a random 3 digit number
func randomInt3Digits() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%03d", r.Intn(1000))
}

// calculateCheckDigit1 calculates the first check digit
func calculateCheckDigit1(sequenceNumber string) string {
	sum := 0
	weights := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(sequenceNumber); i++ {
		digit, _ := strconv.Atoi(string(sequenceNumber[i]))
		sum += digit * weights[i]
	}

	checkDigit1 := sum % 11
	if checkDigit1 == 10 {
		checkDigit1 = 0
	}

	return strconv.Itoa(checkDigit1)
}

// calculateCheckDigit2 calculates the second check digit
func calculateCheckDigit2(stateCode, checkDigit1 string, stateCodeInt int) string {
	sum := 0
	weights := []int{7, 8}
	for i := 0; i < len(stateCode); i++ {
		digit, _ := strconv.Atoi(string(stateCode[i]))
		sum += digit * weights[i]
	}
	checkDigit1Int, _ := strconv.Atoi(checkDigit1)
	sum += checkDigit1Int * 9

	checkDigit2 := sum % 11
	if checkDigit2 == 10 {
		checkDigit2 = 0
	}

	// Special case for states 01 SP and 02 MG
	if (stateCodeInt == 1 || stateCodeInt == 2) && checkDigit2 == 0 {
		checkDigit2 = 1
	}

	return strconv.Itoa(checkDigit2)
}
