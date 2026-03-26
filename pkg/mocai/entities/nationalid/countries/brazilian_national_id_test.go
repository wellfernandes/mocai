package countries

import (
	"errors"
	"testing"
)

func TestGenerateRG(t *testing.T) {
	rg, err := NewRGCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rg.Number == "" {
		t.Error("RG number is empty")
	}

	if len(rg.Number) != 9 {
		t.Errorf("expected 9 char RG number, got %d: %s", len(rg.Number), rg.Number)
	}
}

func TestGenerateRGFormatted(t *testing.T) {
	rg, err := NewRGCustom("ptbr", true, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(rg.Number) != 12 {
		t.Errorf("expected 12 char formatted RG, got %d: %s", len(rg.Number), rg.Number)
	}
}

func TestRGHasStateAndIssuingBody(t *testing.T) {
	rg, err := NewRGCustom("ptbr", false, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if rg.State == "" {
		t.Error("RG state is empty")
	}

	if rg.IssuingBody == "" {
		t.Error("RG issuing body is empty")
	}
}

func TestRGReturnsError(t *testing.T) {
	_, err := NewRGCustom("ptbr", false, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRGErrorIsSentinel(t *testing.T) {
	err := ErrConvertingDigit
	if !errors.Is(err, ErrConvertingDigit) {
		t.Error("expected ErrConvertingDigit to match with errors.Is")
	}
}

func TestGenerateMultipleRGs(t *testing.T) {
	for i := 0; i < 50; i++ {
		rg, err := NewRGCustom("ptbr", false, nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if rg.Number == "" {
			t.Errorf("RG number is empty on iteration %d", i)
		}
	}
}
