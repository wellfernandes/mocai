package countries

import (
	"fmt"
	"math/rand"
	"time"
)

// BrazilianBirthCertificate represents the structure of a Brazilian birth certificate
// containing all its components and the final certificate number
type BrazilianBirthCertificate struct {
	VitalRecordsOffice int
	ArchiveCode        int
	ServiceType        int
	BirthYear          int
	CertificateType    int
	BookNumber         int
	PageNumber         int
	TermNumber         int
	CheckDigits        string
	CertificateNumber  string
}

// GenerateBirthCertificate generates a valid Brazilian birth certificate number
// If formatted is true, returns the number with separators (-)
// Returns a pointer to BrazilianBirthCertificate and error if any validation fails
func GenerateBirthCertificate(formatted bool) (*BrazilianBirthCertificate, error) {

	// Create a new random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Registry Office
	// 1. Vital Records Office [6 digits]
	vitalRecordsOffice := r.Intn(899999) + 100000
	if vitalRecordsOffice < 0 {
		return nil, ErrInvalidVitalRecordsOffice
	}

	// 2. Archive [2 digits]
	archiveCode := 1

	// 3. Civil Registry of Natural Persons [2 digits]
	serviceType := 55

	// 4. Birth Year [4 digits]
	currentYear := time.Now().Year()
	birthYear := r.Intn(currentYear-2010+1) + 2010

	// 5. Certificate type [1 digit]
	certificateType := 1

	// 6. Book number [5 digits]
	bookNumber := r.Intn(89999) + 10000
	if bookNumber < 0 {
		return nil, ErrInvalidBookNumber
	}

	// 7. Page number [3 digits]
	pageNumber := r.Intn(899) + 100
	if pageNumber < 0 {
		return nil, ErrInvalidPageNumber
	}

	// 8. Term number [7 digits]
	termNumber := r.Intn(8999999) + 1000000
	if termNumber < 0 {
		return nil, ErrInvalidTermNumber
	}

	// Number without check digits [30 digits]
	numberWithoutCheckDigits := fmt.Sprintf("%06d%02d%02d%04d%d%05d%03d%07d",
		vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber)
	if len(numberWithoutCheckDigits) != 30 {
		print(numberWithoutCheckDigits)
		return nil, ErrInvalidNumberWithoutCheckDigits
	}

	// 9. Check digits calculation [2 digits]
	checkDigits := calculateCheckDigits(numberWithoutCheckDigits)
	print("checkDigits: ", checkDigits)

	certificateNumber := fmt.Sprintf("%s%02s", numberWithoutCheckDigits, checkDigits)
	if formatted {
		certificateNumber = fmt.Sprintf("%06d %02d %02d %04d %d %05d %03d %07d-%02s",
			vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber, checkDigits)
	}

	if certificateNumber == "" {
		return nil, ErrInvalidBirthCertificate
	}

	brazilianCertificate := &BrazilianBirthCertificate{
		VitalRecordsOffice: vitalRecordsOffice,
		ArchiveCode:        archiveCode,
		ServiceType:        serviceType,
		BirthYear:          birthYear,
		CertificateType:    certificateType,
		BookNumber:         bookNumber,
		PageNumber:         pageNumber,
		TermNumber:         termNumber,
		CheckDigits:        checkDigits,
		CertificateNumber:  certificateNumber,
	}

	return brazilianCertificate, nil
}

// calculateCheckDigits calculates the check digits for the birth certificate number
// using a weight-based algorithm / using Mod 11
func calculateCheckDigits(number string) string {
	dv1Weights := [30]int{
		9, 8, 7, 6, 5, 4, 3, 2, 1, 0,
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 10, 9, 8, 7, 6, 5, 4, 3, 2,
	}

	sumDV1 := 0
	for i := 0; i < 30; i++ {
		// reverse access: b1 = last digit [position 29], b30 = first [position 0]
		digit := int(number[29-i] - '0')
		sumDV1 += digit * dv1Weights[i]
	}

	dv1 := sumDV1 % 11
	if dv1 == 10 {
		dv1 = 1
	}

	dv2Weights := [31]int{
		9, // first term: dv1 * 9
		8, 7, 6, 5, 4, 3, 2, 1, 0,
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
		0, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}

	sumDV2 := 0
	// first term: dv1 * 9
	sumDV2 += dv1 * dv2Weights[0]

	// remaining terms: b1 to b30 with remaining weights
	for i := 0; i < 30; i++ {
		digit := int(number[29-i] - '0')
		sumDV2 += digit * dv2Weights[i+1]
	}

	dv2 := sumDV2 % 11
	if dv2 == 10 {
		dv2 = 1
	}

	return fmt.Sprintf("%d%d", dv1, dv2)
}
