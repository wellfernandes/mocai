package countries

import (
	"fmt"
	"math/rand"
	"time"
)

// Brazilian certificate types
// 1 - for birth certificate
// 2 - for marriage certificate
// 3 - for death certificate
const (
	brazilianBirthCertificateType    = 1
	brazilianMarriageCertificateType = 2
	brazilianDeathCertificateType    = 3
)

// BaseCertificate is the base structure for all brazilian certificates
type BaseCertificate struct {
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

// BirthCertificate represents a brazilian birth certificate
type BirthCertificate struct {
	BaseCertificate
}

// MarriageCertificate represents a brazilian marriage certificate
type MarriageCertificate struct {
	BaseCertificate
}

// DeathCertificate represents a brazilian death certificate
type DeathCertificate struct {
	BaseCertificate
}

// BrazilianCertificates represents all brazilian certificates
type BrazilianCertificates struct {
	BirthCertificate    BirthCertificate
	MarriageCertificate MarriageCertificate
	DeathCertificate    DeathCertificate
}

// GenerateBrazilianCertificates generates a valid brazilian certificates
// If formatted is true, returns the number with separators (-)
// Returns a pointer to BrazilianCertificates and error if any validation fails
func GenerateBrazilianCertificates(formatted bool) (BrazilianCertificates, error) {
	createdBirthCertificate, err := generateBirthCertificate(formatted)
	if err != nil {
		return BrazilianCertificates{}, err
	}

	createdMarriageCertificate, err := generateMarriageCertificate(formatted)
	if err != nil {
		return BrazilianCertificates{}, err
	}

	createdDeathCertificate, err := generateDeathCertificate(formatted)
	if err != nil {
		return BrazilianCertificates{}, err
	}

	createdBrazilianCertificates := BrazilianCertificates{
		BirthCertificate:    createdBirthCertificate,
		MarriageCertificate: createdMarriageCertificate,
		DeathCertificate:    createdDeathCertificate,
	}

	return createdBrazilianCertificates, nil
}

// generateCertificates generates a valid brazilian certificate
// If formatted is true, returns the number with separators (-)
// certificateType is the type of certificate [1 - for birth certificate, 2 - for marriage certificate, 3 - for death certificate]
// Returns a pointer to BaseCertificate and error if any validation fails
func generateCertificate(formatted bool, certificateType int) (BaseCertificate, error) {
	// Registry Office
	// 1. Vital Records Office [6 digits]
	vitalRecordsOffice := generateRandomNumber(100000, 899999)
	if vitalRecordsOffice < 0 {
		return BaseCertificate{}, ErrInvalidVitalRecordsOffice
	}

	// 2. Archive [2 digits]
	archiveCode := 1

	// 3. Civil Registry of Natural Persons [2 digits]
	serviceType := 55

	// 4. Birth Year [4 digits]
	birthYear := generateRandomYear(2010)

	// 5. Certificate type [1 digit]
	//certificateType

	// 6. Book number [5 digits]
	bookNumber := generateRandomNumber(10000, 89999)
	if bookNumber < 0 {
		return BaseCertificate{}, ErrInvalidBookNumber
	}

	// 7. Page number [3 digits]
	pageNumber := generateRandomNumber(100, 899)
	if pageNumber < 0 {
		return BaseCertificate{}, ErrInvalidPageNumber
	}

	// 8. Term number [7 digits]
	termNumber := generateRandomNumber(1000000, 8999999)
	if termNumber < 0 {
		return BaseCertificate{}, ErrInvalidTermNumber
	}

	// Number without check digits [30 digits]
	numberWithoutCheckDigits := fmt.Sprintf("%06d%02d%02d%04d%d%05d%03d%07d",
		vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber)
	if len(numberWithoutCheckDigits) != 30 {
		return BaseCertificate{}, ErrInvalidNumberWithoutCheckDigits
	}

	// 9. Check digits calculation [2 digits]
	checkDigits := calculateCheckDigits(numberWithoutCheckDigits)

	certificateNumber := fmt.Sprintf("%s%02s", numberWithoutCheckDigits, checkDigits)
	if formatted {
		certificateNumber = fmt.Sprintf("%06d %02d %02d %04d %d %05d %03d %07d-%02s",
			vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber, checkDigits)
	}

	if certificateNumber == "" {
		return BaseCertificate{}, ErrInvalidCertificate
	}

	createdBaseCertificate := BaseCertificate{
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

	return createdBaseCertificate, nil
}

// generateBirthCertificate generates a valid Brazilian birth certificate
func generateBirthCertificate(formatted bool) (BirthCertificate, error) {
	base, err := generateCertificate(formatted, brazilianBirthCertificateType)
	if err != nil {
		return BirthCertificate{}, err
	}

	createdBirthCertificate := BirthCertificate{
		BaseCertificate: base,
	}

	return createdBirthCertificate, nil
}

// generateMarriageCertificate generates a valid Brazilian marriage certificate
func generateMarriageCertificate(formatted bool) (MarriageCertificate, error) {
	base, err := generateCertificate(formatted, brazilianMarriageCertificateType)
	if err != nil {
		return MarriageCertificate{}, err
	}

	createdMarriageCertificate := MarriageCertificate{
		BaseCertificate: base,
	}

	return createdMarriageCertificate, nil
}

// generateDeathCertificate generates a valid Brazilian death certificate
func generateDeathCertificate(formatted bool) (DeathCertificate, error) {
	base, err := generateCertificate(formatted, brazilianDeathCertificateType)
	if err != nil {
		return DeathCertificate{}, err
	}

	createdDeathCertificate := DeathCertificate{
		BaseCertificate: base,
	}

	return createdDeathCertificate, nil
}

// generateRandomNumber generates a random number between min and max
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// generateRandomYear generates a random year between startYear and the current year
func generateRandomYear(startYear int) int {
	currentYear := time.Now().Year()
	return rand.Intn(currentYear-startYear+1) + startYear
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
