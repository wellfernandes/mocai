package birth_certificate

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// BrazilianBirthCertificate represents the structure of a Brazilian birth certificate
// containing all its components and the final certificate number
type BrazilianBirthCertificate struct {
	VitalRecordsOffice     int
	ArchiveCode            int
	ServieType             int
	BirthYear              int
	CertificateType        int
	BookNumber             int
	PageNumber             int
	TermNumber             int
	CheckDigits            int
	BirthCertificateNumber string
}

// GenerateBirthCertificate generates a valid Brazilian birth certificate number
// If formatted is true, returns the number with separators (-)
// Returns a pointer to BrazilianBirthCertificate and error if any validation fails
func GenerateBirthCertificate(formatted bool) (*BrazilianBirthCertificate, error) {

	// Create a new random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Registry Office
	// 1. Vital Records Office [6 digits]
	vitalRecordsOffice := r.Intn(999999) + 1
	if vitalRecordsOffice < 0 {
		return nil, ErrInvalidVitalRecordsOffice
	}

	// 2. Archive [2 digits]
	archiveCode := 1

	// 3. Civil Registry of Natural Persons [2 digits]
	servieType := 55

	// 4. Birth Year [4 digits]
	currentYear := time.Now().Year()
	birthYear := r.Intn(currentYear-2010+1) + 2010

	// 5. Certificate type [1 digit]
	certificateType := 1

	// 6. Book number [5 digits]
	bookNumber := r.Intn(99999) + 1
	if bookNumber < 0 {
		return nil, ErrInvalidBookNumber
	}

	// 7. Page number [3 digits]
	pageNumber := r.Intn(999) + 1
	if pageNumber < 0 {
		return nil, ErrInvalidPageNumber
	}

	// 8. Term number [7 digits]
	termNumber := r.Intn(9999999) + 1
	if termNumber < 0 {
		return nil, ErrInvalidTermNumber
	}

	// Number without check digits [30 digits]
	numberWithoutCheckDigits := fmt.Sprintf("%06d%02d%02d%04d%d%05d%03d%07d",
		vitalRecordsOffice, archiveCode, servieType, birthYear, certificateType, bookNumber, pageNumber, termNumber)
	if len(numberWithoutCheckDigits) != 30 {
		return nil, ErrInvalidNumberWithoutCheckDigits
	}

	// 9. Check digits calculation [2 digits]
	checkDigits := calculateCheckDigits(numberWithoutCheckDigits)

	birthCertificateNumber := fmt.Sprintf("%s%02d", numberWithoutCheckDigits, checkDigits)
	if formatted {
		birthCertificateNumber = fmt.Sprintf("%06d %02d %02d %04d %d %05d %03d %07d-%02d",
			vitalRecordsOffice, archiveCode, servieType, birthYear, certificateType, bookNumber, pageNumber, termNumber, checkDigits)
	}

	brazilianCertificate := &BrazilianBirthCertificate{
		VitalRecordsOffice:     vitalRecordsOffice,
		ArchiveCode:            archiveCode,
		ServieType:             servieType,
		BirthYear:              birthYear,
		CertificateType:        certificateType,
		BookNumber:             bookNumber,
		PageNumber:             pageNumber,
		TermNumber:             termNumber,
		CheckDigits:            checkDigits,
		BirthCertificateNumber: birthCertificateNumber,
	}

	return brazilianCertificate, nil
}

// calculateCheckDigits calculates the check digits for the birth certificate number
// using a weight-based algorithm
func calculateCheckDigits(number string) int {
	sum := 0
	weight := 2

	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))
		sum += digit * weight
		weight++
		if weight > 9 {
			weight = 2
		}
	}

	remainder := sum % 11
	checkDigit1 := 11 - remainder
	if checkDigit1 >= 10 {
		checkDigit1 = 1
	}

	numberWithFirstCheckDigit := number + strconv.Itoa(checkDigit1)
	sum = 0
	weight = 3

	for i := len(numberWithFirstCheckDigit) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(numberWithFirstCheckDigit[i]))
		sum += digit * weight
		weight++
		if weight > 9 {
			weight = 2
		}
	}

	remainder = sum % 11
	checkDigit2 := 11 - remainder
	if checkDigit2 >= 10 {
		checkDigit2 = 1
	}

	return checkDigit1*10 + checkDigit2
}
