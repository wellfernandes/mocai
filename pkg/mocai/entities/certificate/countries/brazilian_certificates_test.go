package countries

import (
	"testing"
)

func TestGenerateBrazilianCertificates(t *testing.T) {
	certs, err := NewBrazilCertificatesCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if certs.BirthCertificate == nil {
		t.Error("BirthCertificate is nil")
	}

	if certs.MarriageCertificate == nil {
		t.Error("MarriageCertificate is nil")
	}

	if certs.DeathCertificate == nil {
		t.Error("DeathCertificate is nil")
	}
}

func TestBirthCertificateNumber(t *testing.T) {
	certs, err := NewBrazilCertificatesCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(certs.BirthCertificate.Number) != 32 {
		t.Errorf("expected 32 digit birth certificate number, got %d: %s",
			len(certs.BirthCertificate.Number), certs.BirthCertificate.Number)
	}
}

func TestCertificateFormatted(t *testing.T) {
	certs, err := NewBrazilCertificatesCustom("ptbr", true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(certs.BirthCertificate.Number) <= 32 {
		t.Errorf("expected formatted birth certificate to be longer than 32 chars, got %d: %s",
			len(certs.BirthCertificate.Number), certs.BirthCertificate.Number)
	}
}

func TestCertificateTypes(t *testing.T) {
	certs, err := NewBrazilCertificatesCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if certs.BirthCertificate.Type != 1 {
		t.Errorf("expected birth certificate type 1, got %d", certs.BirthCertificate.Type)
	}

	if certs.MarriageCertificate.Type != 2 {
		t.Errorf("expected marriage certificate type 2, got %d", certs.MarriageCertificate.Type)
	}

	if certs.DeathCertificate.Type != 3 {
		t.Errorf("expected death certificate type 3, got %d", certs.DeathCertificate.Type)
	}
}

func TestGenerateMultipleCertificates(t *testing.T) {
	for i := 0; i < 50; i++ {
		certs, err := NewBrazilCertificatesCustom("ptbr", false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if certs.BirthCertificate.Number == "" {
			t.Errorf("birth certificate number is empty on iteration %d", i)
		}
	}
}
