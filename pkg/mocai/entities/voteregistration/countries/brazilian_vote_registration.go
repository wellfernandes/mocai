package countries

import (
	"fmt"
	"strconv"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
)

// BrazilianVoteRegistration represents a Brazilian vote registration
type BrazilianVoteRegistration struct {
	Section string
	Zone    string
	Number  string
}

// NewBrazilianVoteRegistration generates a new Brazilian vote registration with localized errors
func NewBrazilianVoteRegistration(isFormatted bool) (*BrazilianVoteRegistration, error) {
	return NewBrazilianVoteRegistrationCustom("ptbr", isFormatted, nil)
}

// calculateCheckDigit1 calculates the first check digit.
// it depends on the sequence number.
func calculateCheckDigit1(sequenceNumber string) (string, error) {
	sum := 0
	weights := []int{2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(sequenceNumber); i++ {
		digit, err := strconv.Atoi(string(sequenceNumber[i]))
		if err != nil {
			return "", fmt.Errorf("%s", translations.Get("en_us", "invalid_check_digit_1"))
		}
		sum += digit * weights[i]
	}

	checkDigit1 := sum % 11
	if checkDigit1 == 10 {
		checkDigit1 = 0
	}

	return strconv.Itoa(checkDigit1), nil
}

// calculateCheckDigit2 calculates the second check digit
// it depends on the state code and the first check digit
func calculateCheckDigit2(stateCode, checkDigit1 string, stateCodeInt int) (string, error) {
	sum := 0
	weights := []int{7, 8}
	for i := 0; i < len(stateCode); i++ {
		digit, err := strconv.Atoi(string(stateCode[i]))
		if err != nil {
			return "", fmt.Errorf("%s", translations.Get("en_us", "invalid_check_digit_2"))
		}
		sum += digit * weights[i]
	}
	checkDigit1Int, err := strconv.Atoi(checkDigit1)
	if err != nil {
		return "", fmt.Errorf("%s", translations.Get("en_us", "invalid_check_digit_2"))
	}
	sum += checkDigit1Int * 9

	checkDigit2 := sum % 11
	if checkDigit2 == 10 {
		checkDigit2 = 0
	}

	// special case for states 01 SP and 02 MG
	if (stateCodeInt == 1 || stateCodeInt == 2) && checkDigit2 == 0 {
		checkDigit2 = 1
	}

	return strconv.Itoa(checkDigit2), nil
}

func NewBrazilianVoteRegistrationCustom(lang string, isFormatted bool, rnd translations.RandSource) (*BrazilianVoteRegistration, error) {
	if lang == "" {
		lang = "en_us"
	}

	supported := false
	if translations.Get(lang, "invalid_vote_registration") != "invalid_vote_registration" {
		supported = true
	}
	if !supported {
		lang = "en_us"
	}

	if rnd == nil {
		rnd = translations.DefaultRandSource()
	}

	// generate a random 3-digit section and zone
	section := fmt.Sprintf("%03d", rnd.Intn(1000))
	zone := fmt.Sprintf("%03d", rnd.Intn(1000))

	// generate an 8-digit sequence number
	sequenceNumber := rnd.Intn(99999999) + 1
	sequenceNumberStr := fmt.Sprintf("%08d", sequenceNumber)

	// generate a random state code 01 to 28
	stateCode := rnd.Intn(28) + 1
	stateCodeStr := fmt.Sprintf("%02d", stateCode)

	// calculate the first check digit
	checkDigit1, err := calculateCheckDigit1(sequenceNumberStr)
	if err != nil {
		return nil, fmt.Errorf("%s", translations.Get(lang, "invalid_check_digit_1"))
	}

	// calculate the second check digit
	checkDigit2, err := calculateCheckDigit2(stateCodeStr, checkDigit1, stateCode)
	if err != nil {
		return nil, fmt.Errorf("%s", translations.Get(lang, "invalid_check_digit_2"))
	}

	// combine everything to form the complete number
	number := sequenceNumberStr + stateCodeStr + checkDigit1 + checkDigit2
	if number == "" || len(number) != 12 {
		return nil, fmt.Errorf("%s", translations.Get(lang, "invalid_vote_registration"))
	}

	if isFormatted {
		number = fmt.Sprintf("%s %s %s", number[:4], number[4:8], number[8:])
	}

	return &BrazilianVoteRegistration{
		Section: section,
		Zone:    zone,
		Number:  number,
	}, nil
}
