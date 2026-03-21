package countries

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/brazzcore/mocai/pkg/mocai/translations"
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
	Type               int
	BookNumber         int
	PageNumber         int
	TermNumber         int
	CheckDigits        string
	Number             string
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
	BirthCertificate    *BirthCertificate
	MarriageCertificate *MarriageCertificate
	DeathCertificate    *DeathCertificate
}

func NewBrazilCertificatesCustom(lang string, isFormatted bool, rnd translations.RandSource) (*BrazilianCertificates, error) {
	cert, err := generateBrazilianCertificatesCustom(lang, isFormatted, rnd)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func generateBrazilianCertificatesCustom(lang string, formatted bool, rnd translations.RandSource) (*BrazilianCertificates, error) {
	createdBirthCertificate, err := generateBirthCertificateCustom(lang, formatted, rnd)
	if err != nil {
		return nil, err
	}
	createdMarriageCertificate, err := generateMarriageCertificateCustom(lang, formatted, rnd)
	if err != nil {
		return nil, err
	}
	createdDeathCertificate, err := generateDeathCertificateCustom(lang, formatted, rnd)
	if err != nil {
		return nil, err
	}
	createdBrazilianCertificates := &BrazilianCertificates{
		BirthCertificate:    createdBirthCertificate,
		MarriageCertificate: createdMarriageCertificate,
		DeathCertificate:    createdDeathCertificate,
	}
	return createdBrazilianCertificates, nil
}

func generateCertificateCustom(rnd translations.RandSource, formatted bool, certificateType int, lang ...string) (*BaseCertificate, error) {
	l := "pt_br"
	if len(lang) > 0 && lang[0] != "" {
		l = lang[0]
	}
	if rnd == nil {
		rnd = rand.New(rand.NewSource(int64(rand.Int())))
	}
	vitalRecordsOffice := rnd.Intn(899999-100000+1) + 100000
	if vitalRecordsOffice < 0 {
		return nil, fmt.Errorf("%w: %s", ErrInvalidVitalRecordsOffice, translations.Get(l, "invalid_vital_records_office_number"))
	}
	archiveCode := 1
	serviceType := 55
	birthYear := rnd.Intn(time.Now().Year()-2010+1) + 2010
	bookNumber := rnd.Intn(89999-10000+1) + 10000
	if bookNumber < 0 {
		return nil, fmt.Errorf("%w, %s", ErrInvalidBookNumber, translations.Get(l, "invalid_book_number"))
	}
	pageNumber := rnd.Intn(899-100+1) + 100
	if pageNumber < 0 {
		return nil, ErrInvalidPageNumber
	}
	termNumber := rnd.Intn(8999999-1000000+1) + 1000000
	if termNumber < 0 {
		return nil, fmt.Errorf("%w, %s", ErrInvalidTermNumber, translations.Get(l, "invalid_term_number"))
	}
	numberWithoutCheckDigits := fmt.Sprintf("%06d%02d%02d%04d%d%05d%03d%07d",
		vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber)
	if len(numberWithoutCheckDigits) != 30 {
		return nil, fmt.Errorf("%w, %s", ErrInvalidNumberWithoutCheckDigits, translations.Get(l, "invalid_number_without_check_digits"))
	}
	checkDigits := calculateCheckDigits(numberWithoutCheckDigits)
	certificateNumber := fmt.Sprintf("%s%02s", numberWithoutCheckDigits, checkDigits)
	if formatted {
		certificateNumber = fmt.Sprintf("%06d %02d %02d %04d %d %05d %03d %07d-%02s",
			vitalRecordsOffice, archiveCode, serviceType, birthYear, certificateType, bookNumber, pageNumber, termNumber, checkDigits)
	}
	if certificateNumber == "" {
		return nil, fmt.Errorf("%w: %s", ErrInvalidCertificate, translations.Get(l, "invalid_certificate"))
	}
	createdBaseCertificate := &BaseCertificate{
		VitalRecordsOffice: vitalRecordsOffice,
		ArchiveCode:        archiveCode,
		ServiceType:        serviceType,
		BirthYear:          birthYear,
		Type:               certificateType,
		BookNumber:         bookNumber,
		PageNumber:         pageNumber,
		TermNumber:         termNumber,
		CheckDigits:        checkDigits,
		Number:             certificateNumber,
	}
	return createdBaseCertificate, nil
}

func generateBirthCertificateCustom(lang string, formatted bool, rnd translations.RandSource) (*BirthCertificate, error) {
	base, err := generateCertificateCustom(rnd, formatted, brazilianBirthCertificateType)
	if err != nil {
		return nil, err
	}
	createdBirthCertificate := &BirthCertificate{
		BaseCertificate: *base,
	}
	return createdBirthCertificate, nil
}

func generateMarriageCertificateCustom(lang string, formatted bool, rnd translations.RandSource) (*MarriageCertificate, error) {
	base, err := generateCertificateCustom(rnd, formatted, brazilianMarriageCertificateType)
	if err != nil {
		return nil, err
	}
	createdMarriageCertificate := &MarriageCertificate{
		BaseCertificate: *base,
	}
	return createdMarriageCertificate, nil
}

func generateDeathCertificateCustom(lang string, formatted bool, rnd translations.RandSource) (*DeathCertificate, error) {
	base, err := generateCertificateCustom(rnd, formatted, brazilianDeathCertificateType)
	if err != nil {
		return nil, err
	}
	createdDeathCertificate := &DeathCertificate{
		BaseCertificate: *base,
	}
	return createdDeathCertificate, nil
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
