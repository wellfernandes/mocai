package countries

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

type RG struct {
	Number      string
	State       string
	IssuingBody string
}

// GenerateBrazilianNationalID generates a valid Brazilian national ID [RG] for São Paulo.
func GenerateBrazilianNationalID(formatted bool) (*RG, error) {
	rgNumber, err := calculateSPRGDigit()
	if err != nil {
		return nil, err
	}
	if formatted {
		rgNumber = formatRG(rgNumber)
	}
	return &RG{
		Number:      rgNumber,
		State:       "SP",
		IssuingBody: "SSP - Secretaria de Seguranca Publica",
	}, nil
}

func calculateSPRGDigit() (string, error) {
	// generate a random base until 8 digits
	base := globalRand.Intn(100000000)
	baseStr := fmt.Sprintf("%08d", base)

	// slice contains the digits
	d := make([]int, 8)
	for i := range 8 {
		val, err := strconv.Atoi(string(baseStr[i]))
		if err != nil {
			return "", ErrToConvertDigit
		}
		d[i] = val
	}

	// wheights from right to left: 9, 8, 7, 6, 5, 4, 3, 2
	wheights := []int{2, 3, 4, 5, 6, 7, 8, 9}

	// calculate sum
	sum := 0
	for i := 0; i < 8; i++ {
		sum += d[7-i] * wheights[i]
	}

	// calculate check digit
	checkDigit := sum % 11
	var dvStr string
	if checkDigit == 10 {
		dvStr = "X"
	} else {
		dvStr = strconv.Itoa(checkDigit)
	}

	full := fmt.Sprintf("%s%s", baseStr, dvStr)
	return full, nil
}

func formatRG(input string) string {
	if len(input) != 9 {
		return input
	}
	return fmt.Sprintf("%s.%s.%s-%s", input[0:3], input[3:6], input[6:8], input[8:])
}
