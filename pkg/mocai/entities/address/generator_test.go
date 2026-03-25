package address

import (
	"testing"
)

func TestGenerateAddress(t *testing.T) {
	addr, err := NewAddress("ptbr", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if addr.Street == "" {
		t.Error("Street is empty")
	}

	if addr.City == "" {
		t.Error("City is empty")
	}

	if addr.State == "" {
		t.Error("State is empty")
	}

	if addr.UF == "" {
		t.Error("UF is empty")
	}

	if addr.ZIP == "" {
		t.Error("ZIP is empty")
	}

	if addr.Number <= 0 {
		t.Errorf("Number should be positive, got %d", addr.Number)
	}
}

func TestGenerateAddressFallbackLanguage(t *testing.T) {
	addr, err := NewAddress("xyz", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if addr.Street == "" || addr.City == "" {
		t.Errorf("address has empty fields with fallback: %+v", addr)
	}
}

func TestGenerateMultipleAddresses(t *testing.T) {
	for i := 0; i < 50; i++ {
		addr, err := NewAddress("ptbr", nil)
		if err != nil {
			t.Fatalf("unexpected error on iteration %d: %v", i, err)
		}

		if addr.Street == "" || addr.City == "" || addr.State == "" {
			t.Errorf("address has empty fields on iteration %d: %+v", i, addr)
		}
	}
}
