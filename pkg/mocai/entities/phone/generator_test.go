package phone

import (
	"errors"
	"strings"
	"testing"
)

func TestGeneratePhone(t *testing.T) {
	ph, err := NewPhone("ptbr", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ph.AreaCode == "" {
		t.Error("AreaCode is empty")
	}

	if ph.Number == "" {
		t.Error("Number is empty")
	}

	if !strings.HasPrefix(ph.Number, "9") {
		t.Errorf("expected number to start with '9', got %s", ph.Number)
	}

	if len(ph.Number) != 9 {
		t.Errorf("expected 9 digit number, got %d: %s", len(ph.Number), ph.Number)
	}
}

func TestGeneratePhoneUnsupportedLang(t *testing.T) {
	_, err := NewPhone("xyz", nil)
	if err == nil {
		t.Fatal("expected error for unsupported language, got nil")
	}
	if !errors.Is(err, ErrNoAreaCodes) {
		t.Errorf("expected ErrNoAreaCodes, got: %v", err)
	}
}

func TestGenerateMultiplePhones(t *testing.T) {
	for i := 0; i < 50; i++ {
		ph, err := NewPhone("ptbr", nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if ph.AreaCode == "" || ph.Number == "" {
			t.Errorf("phone has empty fields on iteration %d: %+v", i, ph)
		}
	}
}
