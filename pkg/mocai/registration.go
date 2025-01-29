package mocai

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type RegistrationID struct {
	ID string
}

func GenerateRegistrationID(locale string) (RegistrationID, error) {
	id, err := GenerateID(locale, 8)
	if err != nil {
		return RegistrationID{}, err
	}

	registrationID := RegistrationID{
		ID: id,
	}

	return registrationID, nil
}

// generates a random registration number in Portuguese.
func GenerateID(locale string, length int) (string, error) {
	switch locale {
	case "pt-br":
		if length < 2 {
			return "", fmt.Errorf("length must be at least 2 to include the checksum")
		}

		// Create a local random source
		source := rand.NewSource(time.Now().UnixNano())
		rng := rand.New(source)

		// Generate random digits for the ID, leaving room for a checksum digit
		id := make([]int, length-1)
		for i := 0; i < length-1; i++ {
			id[i] = rng.Intn(10) // Random digit between 0-9
		}

		// Calculate checksum (mod 10 of sum of digits for simplicity)
		sum := 0
		for _, digit := range id {
			sum += digit
		}
		checksum := (10 - (sum % 10)) % 10

		// Append checksum to ID
		id = append(id, checksum)

		// Convert ID to string
		idStr := ""
		for _, digit := range id {
			idStr += strconv.Itoa(digit)
		}

		return idStr, nil
	default:
		return "", errors.New("unsupported locale")
	}
}
