package countries

import (
	"errors"
	"testing"
)

func TestGenerateVoteRegistration(t *testing.T) {
	vr, err := NewBrazilianVoteRegistrationCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if vr.Number == "" {
		t.Error("vote registration number is empty")
	}

	if len(vr.Number) != 12 {
		t.Errorf("expected 12 digit number, got %d: %s", len(vr.Number), vr.Number)
	}

	if vr.Section == "" {
		t.Error("section is empty")
	}

	if vr.Zone == "" {
		t.Error("zone is empty")
	}
}

func TestGenerateVoteRegistrationFormatted(t *testing.T) {
	vr, err := NewBrazilianVoteRegistrationCustom("ptbr", true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(vr.Number) != 14 {
		t.Errorf("expected 14 char formatted number, got %d: %s", len(vr.Number), vr.Number)
	}
}

func TestGenerateVoteRegistrationDefault(t *testing.T) {
	vr, err := NewBrazilianVoteRegistration(false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if vr.Number == "" {
		t.Error("vote registration number is empty")
	}
}

func TestVoteRegistrationErrorIsSentinel(t *testing.T) {
	if !errors.Is(ErrInvalidVoteRegistration, ErrInvalidVoteRegistration) {
		t.Error("expected ErrInvalidVoteRegistration to match with errors.Is")
	}

	if !errors.Is(ErrInvalidCheckDigit1, ErrInvalidCheckDigit1) {
		t.Error("expected ErrInvalidCheckDigit1 to match with errors.Is")
	}

	if !errors.Is(ErrInvalidCheckDigit2, ErrInvalidCheckDigit2) {
		t.Error("expected ErrInvalidCheckDigit2 to match with errors.Is")
	}
}

func TestGenerateMultipleVoteRegistrations(t *testing.T) {
	for i := 0; i < 50; i++ {
		vr, err := NewBrazilianVoteRegistrationCustom("ptbr", false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if vr.Number == "" || len(vr.Number) != 12 {
			t.Errorf("invalid vote registration on iteration %d: %+v", i, vr)
		}
	}
}
